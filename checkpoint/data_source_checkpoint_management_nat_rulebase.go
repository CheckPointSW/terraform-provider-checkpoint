package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementNatRulebase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNatRulebaseRead,
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the \"ip-only\" parameter to true.",
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
			"use_object_dictionary": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use object dictionary.",
			},
			"filter_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Sets filter preferences.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.",
						},
						"packet_search_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "When 'search_mode' is set to 'packet', this object allows to set the packet search preferences.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"expand_group_members": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.",
									},
									"expand_group_with_exclusion_members": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on 'Any' object.",
									},
									"match_on_any": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.",
									},
									"match_on_group_with_exclusion": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on a group-with-exclusion.",
									},
									"match_on_negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on a negated cell.",
									},
								},
							},
						},
					},
				},
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
			"rulebase": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "NAT rulebase.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type.",
						},
						"rulebase": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of object unique identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"objects_dictionary": {
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

func dataSourceManagementNatRulebaseRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("package"); ok {
		payload["package"] = v.(string)
	}

	if v, ok := d.GetOk("filter"); ok {
		payload["filter"] = v.(string)
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

	if v, ok := d.GetOkExists("use_object_dictionary"); ok {
		payload["use-object-dictionary"] = v.(int)
	}

	if _, ok := d.GetOk("filter_settings"); ok {
		filterSettings := make(map[string]interface{})

		if v, ok := d.GetOk("filter_settings.search_mode"); ok {
			filterSettings["search-mode"] = v.(string)
		}

		if _, ok := d.GetOk("filter_settings.packet_search_settings"); ok {
			packetSearchSettings := make(map[string]interface{})

			if v, ok := d.GetOkExists("filter_settings.packet_search_settings.expand_group_members"); ok {
				packetSearchSettings["expand-group-members"] = v.(bool)
			}

			if v, ok := d.GetOkExists("filter_settings.packet_search_settings.expand_group_with_exclusion_members"); ok {
				packetSearchSettings["expand-group-with-exclusion-members"] = v.(bool)
			}

			if v, ok := d.GetOkExists("filter_settings.packet_search_settings.match_on_any"); ok {
				packetSearchSettings["match-on-any"] = v.(bool)
			}

			if v, ok := d.GetOkExists("filter_settings.packet_search_settings.match_on_group_with_exclusion"); ok {
				packetSearchSettings["match-on-group-with-exclusion"] = v.(bool)
			}

			if v, ok := d.GetOkExists("filter_settings.packet_search_settings.match_on_negate"); ok {
				packetSearchSettings["match-on-negate"] = v.(bool)
			}

			filterSettings["packet-search-settings"] = packetSearchSettings
		}

		payload["filter_settings"] = filterSettings
	}

	showNatRulebaseRes, err := client.ApiCall("show-nat-rulebase", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNatRulebaseRes.Success {
		return fmt.Errorf(showNatRulebaseRes.ErrorMsg)
	}

	natRulebaseData := showNatRulebaseRes.GetData()

	log.Println("show-nat-rulebase JSON = ", natRulebaseData)

	d.SetId("show-nat-rulebase-" + acctest.RandString(10))

	if v := natRulebaseData["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := natRulebaseData["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := natRulebaseData["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := natRulebaseData["objects-dictionary"]; v != nil {
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
			_ = d.Set("objects_dictionary", objectsListState)
		} else {
			_ = d.Set("objects_dictionary", objectsList)
		}
	} else {
		_ = d.Set("objects_dictionary", nil)
	}

	if v := natRulebaseData["rulebase"]; v != nil {
		rulebaseList := v.([]interface{})
		if len(rulebaseList) > 0 {
			var rulebaseListState []map[string]interface{}
			for i := range rulebaseList {
				ruleMap := rulebaseList[i].(map[string]interface{})
				ruleMapToAdd := make(map[string]interface{})

				if v := ruleMap["name"]; v != nil {
					ruleMapToAdd["name"] = v
				}

				if v := ruleMap["uid"]; v != nil {
					ruleMapToAdd["uid"] = v
				}

				if v := ruleMap["type"]; v != nil {
					ruleMapToAdd["type"] = v
				}

				if v := ruleMap["rulebase"]; v != nil {
					rules := v.([]interface{})
					rulesUids := make([]string, 0)
					if len(rules) > 0 {
						for i := range rules {
							ruleJson := rules[i].(map[string]interface{})
							rulesUids = append(rulesUids, ruleJson["uid"].(string))
						}
					}
					ruleMapToAdd["rulebase"] = rulesUids
				}
				rulebaseListState = append(rulebaseListState, ruleMapToAdd)
			}
			_ = d.Set("rulebase", rulebaseListState)
		} else {
			_ = d.Set("rulebase", rulebaseList)
		}
	} else {
		_ = d.Set("rulebase", nil)
	}

	return nil
}
