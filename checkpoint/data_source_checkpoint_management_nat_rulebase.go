package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
)

func dataSourceManagementNatRuleBase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNatRuleBaseRead,

		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. If an operator is not used, the default OR operator applies.",
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
						"expand_group_members": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.",
							Default:     false,
						},
						"expand_group_with_exclusion_members": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the \"include\" part and is not a member of the \"except\" part.",
							Default:     false,
						},
						"match_on_any": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to match on 'Any' object.",
							Default:     true,
						},
						"match_on_group_with_exclusion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to match on a group-with-exclusion.",
							Default:     true,
						},
						"match_on_negate": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to match on a negated cell.",
							Default:     true,
						},
					},
				},
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximal number of returned results.",
				Default:     50,
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of the results to initially skip.",
				Default:     0,
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
			"dereference_group_members": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When true, the source, destination and services & applications parameters are displayed as ranges of IP addresses and port numbers rather than network objects.\nObjects that are not represented using IP addresses or port numbers are presented as objects.\nIn addition, the response of each rule does not contain the parameters: source, source-negate, destination, destination-negate, service and service-negate, but instead it contains the parameters: source-ranges, destination-ranges and service-ranges.\n\nNote: Requesting to show rules as ranges is limited up to 20 rules per request, otherwise an error is returned. If you wish to request more rules, use the offset and limit parameters to limit your request.",
				Default:     false,
			},
			"show_membership": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
				Default:     false,
			},
			"rulebase": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The show rulebase api reply",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"objects_dictionary": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Must be unique in the domain",
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
								},
							},
						},
						"rulebase": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rules uid.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rules name.",
									},
									"original_destination": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
									},
									"translated_destination": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "True if negate is set for destination.",
									},
									"install_on": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Which Gateways identified by the name or UID to install the policy on.",
										Elem:        schema.TypeString,
									},
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable/Disable the rule.",
									},
									"original_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
									},
									"translated_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "True if negate is set for service.",
									},
									"auto_generated": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "service resource.",
									},
									"original_source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
									},
									"translated_source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "True if negate is set for source.",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Communities or Directional.",
									},
									"comments": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Comments string.",
									},
									"rule_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of the rule.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rules type.",
									},
									"rulebase": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uid": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rules uid.",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rules name.",
												},
												"original_destination": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
												},
												"translated_destination": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "True if negate is set for destination.",
												},
												"install_on": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Which Gateways identified by the name or UID to install the policy on.",
													Elem:        schema.TypeString,
												},
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Enable/Disable the rule.",
												},
												"original_service": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
												},
												"translated_service": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "True if negate is set for service.",
												},
												"auto_generated": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "service resource.",
												},
												"original_source": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
												},
												"translated_source": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "True if negate is set for source.",
												},
												"method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Communities or Directional.",
												},
												"comments": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Comments string.",
												},
												"rule_number": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of the rule.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rules type.",
												},
											},
										},
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

func dataSourceManagementNatRuleBaseRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}
	payload["package"] = d.Get("package").(string)
	if v, ok := d.GetOk("filter"); ok {
		payload["filter"] = v.(string)
	}
	if v, ok := d.GetOk("filter_settings"); ok {
		filters, ok := v.(map[string]interface{})
		if ok {

			filtersMapToReturn := make(map[string]interface{})
			packetSearchMap := make(map[string]interface{})

			if val, ok := filters["search_mode"]; ok {
				filtersMapToReturn["search-mode"] = val
			} else {
				filtersMapToReturn["search-mode"] = "general"
			}

			if val, ok := filters["expand_group_members"]; ok {
				packetSearchMap["expand-group-members"] = val
			} else {
				packetSearchMap["expand-group-members"] = false
			}

			if val, ok := filters["expand_group_with_exclusion_members"]; ok {
				packetSearchMap["expand-group-with-exclusion-members"] = val
			} else {
				packetSearchMap["expand-group-with-exclusion-members"] = false
			}

			if val, ok := filters["match_on_any"]; ok {
				packetSearchMap["match-on-any"] = val
			} else {
				packetSearchMap["match-on-any"] = true
			}

			if val, ok := filters["match_on_group_with_exclusion"]; ok {
				packetSearchMap["match-on-group-with-exclusion"] = val
			} else {
				packetSearchMap["match-on-group-with-exclusion"] = true
			}

			if val, ok := filters["match_on_negate"]; ok {
				packetSearchMap["match-on-negate"] = val
			} else {
				packetSearchMap["match-on-negate"] = true
			}

			filtersMapToReturn["packet-search-settings"] = packetSearchMap
			payload["filter-settings"] = filtersMapToReturn
		}
	}
	if v, ok := d.GetOk("limit"); ok {
		payload["limit"] = v.(int)
	}
	if v, ok := d.GetOk("offset"); ok {
		payload["offset"] = v.(int)
	}
	if v, ok := d.GetOk("order"); ok {

		ordersList, ok := v.([]interface{})
		var ordersDictToReturn []map[string]interface{}

		if ok {
			for i := range ordersList {

				objectsMap := ordersList[i].(map[string]interface{})

				tempOrder := make(map[string]interface{})

				if v, _ := objectsMap["asc"]; v != nil && v != "" {
					tempOrder["ASC"] = v
				}

				if v, _ := objectsMap["desc"]; v != nil && v != "" {
					tempOrder["DESC"] = v
				}

				ordersDictToReturn = append(ordersDictToReturn, tempOrder)
			}
			payload["order"] = ordersDictToReturn
		}
	}
	if v, ok := d.GetOk("package"); ok {
		payload["package"] = v.(string)
	}
	if v, ok := d.GetOk("dereference_group_members"); ok {
		payload["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOk("show_membership"); ok {
		payload["show-membership"] = v.(bool)
	}

	showRuleBaseRes, err := client.ApiCall("show-nat-rulebase", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRuleBaseRes.Success {
		return fmt.Errorf(showRuleBaseRes.ErrorMsg)
	}
	ruleBaseJson := showRuleBaseRes.GetData()

	log.Println("Read ruleBaseJson - Show JSON = ", ruleBaseJson)
	var outputRuleBase []interface{}
	ruleBaseToReturn := make(map[string]interface{})
	if v := ruleBaseJson["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := ruleBaseJson["from"]; v != nil {
		ruleBaseToReturn["from"] = int(math.Round(v.(float64)))
	} else {
		ruleBaseToReturn["from"] = 0
	}
	if ruleBaseJson["objects-dictionary"] != nil {

		objectsList, ok := ruleBaseJson["objects-dictionary"].([]interface{})
		var objectDictToReturn []map[string]interface{}

		if ok {
			for i := range objectsList {

				objectsMap := objectsList[i].(map[string]interface{})

				tempObject := make(map[string]interface{})

				if v, _ := objectsMap["name"]; v != nil {
					tempObject["name"] = v
				}

				if v, _ := objectsMap["uid"]; v != nil {
					tempObject["uid"] = v
				}

				if v, _ := objectsMap["type"]; v != nil {
					tempObject["type"] = v
				}

				objectDictToReturn = append(objectDictToReturn, tempObject)
			}
			ruleBaseToReturn["objects_dictionary"] = objectDictToReturn
		}
	} else {
		ruleBaseToReturn["objects_dictionary"] = []map[string]interface{}{}
	}

	if ruleBaseJson["rulebase"] != nil {
		ruleBaseDictToReturn := readNatRuleBaseField(ruleBaseJson)
		ruleBaseToReturn["rulebase"] = ruleBaseDictToReturn
	} else {
		ruleBaseToReturn["rulebase"] = []map[string]interface{}{}
	}

	if v := ruleBaseJson["to"]; v != nil {
		ruleBaseToReturn["to"] = int(math.Round(v.(float64)))
	} else {
		ruleBaseToReturn["to"] = 0
	}
	if v := ruleBaseJson["total"]; v != nil {
		ruleBaseToReturn["total"] = int(math.Round(v.(float64)))
	} else {
		ruleBaseToReturn["total"] = 0
	}
	outputRuleBase = append(outputRuleBase, ruleBaseToReturn)
	_ = d.Set("rulebase", outputRuleBase)
	return nil
}

func readNatRuleBaseField(RuleBase map[string]interface{}) []map[string]interface{} {
	ruleBaseList, ok := RuleBase["rulebase"].([]interface{})
	var ruleBaseDictToReturn []map[string]interface{}

	if ok {
		for i := range ruleBaseList {

			ruleBaseMap := ruleBaseList[i].(map[string]interface{})

			tempRulebase := make(map[string]interface{})
			if v, _ := ruleBaseMap["uid"]; v != nil {
				tempRulebase["uid"] = v
			}

			if v, _ := ruleBaseMap["name"]; v != nil {
				tempRulebase["name"] = v
			}

			if v, _ := ruleBaseMap["original-destination"]; v != nil {
				tempRulebase["original_destination"] = v
			}

			if v, _ := ruleBaseMap["translated-destination"]; v != nil {
				tempRulebase["translated_destination"] = v
			}

			if v, _ := ruleBaseMap["install-on"]; v != nil {
				tempRulebase["install_on"] = v
			}

			if v, _ := ruleBaseMap["enabled"]; v != nil {
				tempRulebase["enabled"] = v.(bool)
			}

			if v, _ := ruleBaseMap["original-source"]; v != nil {
				tempRulebase["original_source"] = v
			}

			if v, _ := ruleBaseMap["translated-source"]; v != nil {
				tempRulebase["translated_source"] = v
			}

			if v, _ := ruleBaseMap["original-service"]; v != nil {
				tempRulebase["original_service"] = v
			}

			if v, _ := ruleBaseMap["translated-service"]; v != nil {
				tempRulebase["translated_service"] = v
			}

			if v, _ := ruleBaseMap["type"]; v != nil {
				tempRulebase["type"] = v
			}
			if v, _ := ruleBaseMap["comments"]; v != nil && v != "" {
				tempRulebase["comments"] = v
			}

			if v, _ := ruleBaseMap["method"]; v != nil {
				tempRulebase["method"] = v
			}

			if v := ruleBaseMap["rule-number"]; v != nil {
				tempRulebase["rule_number"] = v
			}

			if v, _ := ruleBaseMap["type"]; v != nil {
				tempRulebase["type"] = v
			}

			if v, _ := ruleBaseMap["rulebase"]; v != nil {
				tempRulebase["rulebase"] = readNatRuleBaseField(ruleBaseMap)
			}

			ruleBaseDictToReturn = append(ruleBaseDictToReturn, tempRulebase)
		}
	}
	return ruleBaseDictToReturn
}
