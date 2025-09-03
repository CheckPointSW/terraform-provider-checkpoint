package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementNetworks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNetworksRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by.",
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
			"fetch_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, fetches all results.",
				Default:     false,
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
						"subnet4": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"subnet_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask.",
						},
						"nat_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "NAT settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_rule": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to add automatic address translation rules.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 address.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 address.",
									},
									"hide_behind": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
									},
									"install_on": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Which gateway should apply the NAT translation.",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "NAT translation method.",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of tag identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"broadcast": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Allow broadcast address inclusion.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments string.",
						},
						"groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of group identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
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
						"icon": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object icon.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementNetworksRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	showNetworksRes := checkpoint.APIResponse{}
	var err error
	fetchAll, _ := d.GetOkExists("fetch_all")

	if fetchAll.(bool) {
		showNetworksRes, err = client.ApiQuery("show-networks", "full", "objects", true, map[string]interface{}{})
	} else {
		payload := make(map[string]interface{})

		if val, ok := d.GetOk("filter"); ok {
			payload["filter"] = val.(string)
		}

		if val, ok := d.GetOk("limit"); ok {
			payload["limit"] = val.(int)
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

		if val, ok := d.GetOk("offset"); ok {
			payload["offset"] = val.(int)
		}

		payload["details-level"] = "full"
		showNetworksRes, err = client.ApiCallSimple("show-networks", payload)
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNetworksRes.Success {
		return fmt.Errorf(showNetworksRes.ErrorMsg)
	}

	networks := showNetworksRes.GetData()

	log.Println("Read Networks - Show JSON = ", networks)

	d.SetId("show-networks-" + acctest.RandString(10))

	if v := networks["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := networks["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := networks["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := networks["objects"]; v != nil {
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

				if v := objectMap["subnet4"]; v != nil {
					objectMapToAdd["subnet4"] = v
				}

				if v := objectMap["subnet6"]; v != nil {
					objectMapToAdd["subnet6"] = v
				}

				if v := objectMap["mask-length4"]; v != nil {
					objectMapToAdd["mask_length4"] = v
				}

				if v := objectMap["mask-length6"]; v != nil {
					objectMapToAdd["mask_length6"] = v
				}

				if v := objectMap["subnet-mask"]; v != nil {
					objectMapToAdd["subnet_mask"] = v
				}

				if v := objectMap["broadcast"]; v != nil {
					objectMapToAdd["broadcast"] = v
				}

				if v := objectMap["comments"]; v != nil {
					objectMapToAdd["comments"] = v
				}

				if v := objectMap["color"]; v != nil {
					objectMapToAdd["color"] = v
				}

				if objectMap["nat-settings"] != nil {

					natSettingsMap := objectMap["nat-settings"].(map[string]interface{})

					natSettingsMapToReturn := make(map[string]interface{})

					if v, _ := natSettingsMap["auto-rule"]; v != nil {
						natSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
					}

					if v, _ := natSettingsMap["ipv4-address"]; v != "" && v != nil {
						natSettingsMapToReturn["ipv4_address"] = v
					}

					if v, _ := natSettingsMap["ipv6-address"]; v != "" && v != nil {
						natSettingsMapToReturn["ipv6_address"] = v
					}

					if v, _ := natSettingsMap["hide-behind"]; v != nil {
						natSettingsMapToReturn["hide_behind"] = v
					}

					if v, _ := natSettingsMap["install-on"]; v != nil {
						natSettingsMapToReturn["install_on"] = v
					}

					if v, _ := natSettingsMap["method"]; v != nil {
						natSettingsMapToReturn["method"] = v
					}

					objectMapToAdd["nat_settings"] = natSettingsMapToReturn

				} else {
					objectMapToAdd["nat_settings"] = nil
				}

				if objectMap["groups"] != nil {
					groupsJson := objectMap["groups"].([]interface{})
					groupsIds := make([]string, 0)
					if len(groupsJson) > 0 {
						// Create slice of group names
						for _, group := range groupsJson {
							group := group.(map[string]interface{})
							groupsIds = append(groupsIds, group["name"].(string))
						}
					}
					objectMapToAdd["groups"] = groupsIds
				} else {
					objectMapToAdd["groups"] = nil
				}

				if objectMap["tags"] != nil {
					tagsJson := objectMap["tags"].([]interface{})
					var tagsIds = make([]string, 0)
					if len(tagsJson) > 0 {
						// Create slice of tag names
						for _, tag := range tagsJson {
							tag := tag.(map[string]interface{})
							tagsIds = append(tagsIds, tag["name"].(string))
						}
					}
					objectMapToAdd["tags"] = tagsIds
				} else {
					objectMapToAdd["tags"] = nil
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

				if v := objectMap["icon"]; v != nil {
					objectMapToAdd["icon"] = v
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
