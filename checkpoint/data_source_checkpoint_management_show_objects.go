package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementShowObjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementShowObjectsRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the \"ip-only\" parameter to true.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The objects' type, e.g.: host, service-tcp, network, address-range...",
				Default:     "object",
			},
			"ip_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If using \"filter\", use this field to search objects by their IP address only, without involving the textual search.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximal number of returned results.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
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
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
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
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Objects list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name.",
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
						"domain": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Information about the domain that holds the Object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name.",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object unique identifier.",
									},
									"domain_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Domain type.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementShowObjectsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

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

	if v, ok := d.GetOk("type"); ok {
		payload["type"] = v.(string)
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

	showObjectsRes, err := client.ApiCall("show-objects", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showObjectsRes.Success {
		return fmt.Errorf(showObjectsRes.ErrorMsg)
	}

	objectsData := showObjectsRes.GetData()

	log.Println("show-objects JSON = ", objectsData)

	d.SetId("show-objects-" + acctest.RandString(10))

	if v := objectsData["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := objectsData["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := objectsData["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := objectsData["objects"]; v != nil {
		objectsList := v.([]interface{})
		if len(objectsList) > 0 {
			var objectsListState []map[string]interface{}
			for i := range objectsList {
				objectMap := objectsList[i].(map[string]interface{})
				objectMapToAdd := make(map[string]interface{})

				if v := objectMap["name"]; v != nil {
					objectMapToAdd["name"] = v
				}

				if v := objectMap["uid"]; v != nil {
					objectMapToAdd["uid"] = v
				}

				if v := objectMap["type"]; v != nil {
					objectMapToAdd["type"] = v
				}

				if v := objectMap["domain"]; v != nil {
					domainMap := v.(map[string]interface{})
					domainMapToAdd := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToAdd["name"] = v
					}

					if v := domainMap["uid"]; v != nil {
						domainMapToAdd["uid"] = v
					}

					if v := domainMap["domain-type"]; v != nil {
						domainMapToAdd["domain_type"] = v
					}
					objectMapToAdd["domain"] = domainMapToAdd
				}
				objectsListState = append(objectsListState, objectMapToAdd)
			}
			_ = d.Set("objects", objectsListState)
		} else {
			_ = d.Set("objects", objectsList)
		}
	} else {
		_ = d.Set("objects", nil)
	}

	return nil
}
