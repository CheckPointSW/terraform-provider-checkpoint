package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementObjects() *schema.Resource {
	return &schema.Resource{
		Read:               dataSourceManagementObjectsRead,
		DeprecationMessage: "This data source is deprecated. Use instead data source 'checkpoint_management_show_objects'",
		Schema: map[string]*schema.Schema{
			"uids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of UIDs of the objects to retrieve.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the \"ip-only\" parameter to true.",
			},
			"ip_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If using \"filter\", use this field to search objects by their IP address only, without involving the textual search.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     50,
				Description: "The maximal number of returned results.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Number of the results to initially skip.",
			},
			"order": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in ascending order.",
						},
						"desc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in descending order.",
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "object",
				Description: "The objects' type, e.g.: host, service-tcp, network, address-range...",
			},
			"dereference_group_members": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to dereference \"members\" field by details level for every object in reply.",
			},
			"show_membership": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"groups\" field for every object in reply.",
			},
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
			},
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. In case of details-level is set to full, all unsupported objects will be shown with a warning saying that the response could change when support for these object is added.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type.",
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
						},
					},
				},
			},
			"to": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "To which element number the query was done.",
			},
			"total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of elements returned by the query.",
			},
		},
	}
}

func dataSourceManagementObjectsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("uids"); ok {
		payload["uids"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("filter"); ok {
		payload["filter"] = v.(string)
	}

	if v, ok := d.GetOkExists("ip_only"); ok {
		payload["ip-only"] = v.(bool)
	}

	if v, ok := d.GetOk("limit"); ok {
		payload["limit"] = v.(int)
	}

	if v, ok := d.GetOk("offset"); ok {
		payload["offset"] = v.(int)
	}

	if v, ok := d.GetOk("order"); ok {
		orderList := v.([]interface{})

		if len(orderList) > 0 {
			var orderPayload []map[string]interface{}

			for i := range orderList {
				payload := make(map[string]interface{})

				if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".asc"); ok {
					payload["ASC"] = v.(string)
				}
				if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".desc"); ok {
					payload["DESC"] = v.(string)
				}

				orderPayload = append(orderPayload, payload)
			}

			payload["order"] = orderPayload
		}
	}

	if v, ok := d.GetOk("type"); ok {
		payload["type"] = v.(string)
	}

	if v, ok := d.GetOkExists("dereference_group_members"); ok {
		payload["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_membership"); ok {
		payload["show-membership"] = v.(bool)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		payload["domains-to-process"] = v.(*schema.Set).List()
	}

	showObjectsRes, err := client.ApiCall("show-objects", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showObjectsRes.Success {
		return fmt.Errorf(showObjectsRes.ErrorMsg)
	}

	objects := showObjectsRes.GetData()

	log.Println("Read Objects - Show JSON = ", objects)

	d.SetId("show-objects-" + acctest.RandString(10))

	if v := objects["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if objects["objects"] != nil {
		objectsList := objects["objects"].([]interface{})

		if len(objectsList) > 0 {
			var objectsListToReturn []map[string]interface{}

			for i := range objectsList {
				objectsMap := objectsList[i].(map[string]interface{})

				objectsMapToAdd := make(map[string]interface{})

				if v, _ := objectsMap["name"]; v != nil {
					objectsMapToAdd["name"] = v
				}
				if v, _ := objectsMap["uid"]; v != nil {
					objectsMapToAdd["uid"] = v
				}
				if v, _ := objectsMap["type"]; v != nil {
					objectsMapToAdd["type"] = v
				}
				if v, _ := objectsMap["color"]; v != nil {
					objectsMapToAdd["color"] = v
				}

				objectsListToReturn = append(objectsListToReturn, objectsMapToAdd)
			}

			_ = d.Set("objects", objectsListToReturn)
		} else {
			_ = d.Set("objects", objectsList)
		}
	} else {
		_ = d.Set("objects", nil)
	}

	if v := objects["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := objects["total"]; v != nil {
		_ = d.Set("total", v)
	}

	return nil
}
