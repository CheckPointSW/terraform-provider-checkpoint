package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
)

func dataSourceManagementHttpsRuleBase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHttpsRuleBaseRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
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
			"package": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the package.",
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
									"destination": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem:        schema.TypeString,
									},
									"destination_negate": {
										Type:        schema.TypeBool,
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
									"service": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem:        schema.TypeString,
									},
									"service_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for service.",
									},
									"source": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem:        schema.TypeString,
									},
									"source_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for source.",
									},
									"site_category": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem:        schema.TypeString,
									},
									"site_category_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for source.",
									},
									"layer": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Comments string.",
									},
									"certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
									},
									"track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
									},
									"blade": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem:        schema.TypeString,
									},
									"comments": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Comments string.",
									},
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
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
												"destination": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
													Elem:        schema.TypeString,
												},
												"destination_negate": {
													Type:        schema.TypeBool,
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
												"service": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
													Elem:        schema.TypeString,
												},
												"service_negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True if negate is set for service.",
												},
												"source": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
													Elem:        schema.TypeString,
												},
												"source_negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True if negate is set for source.",
												},
												"site_category": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
													Elem:        schema.TypeString,
												},
												"site_category_negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True if negate is set for source.",
												},
												"layer": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Comments string.",
												},
												"certificate": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
												},
												"track": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
												},
												"blade": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Collection of Network objects identified by the name or UID.",
													Elem:        schema.TypeString,
												},
												"comments": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Comments string.",
												},
												"action": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
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

func dataSourceManagementHttpsRuleBaseRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}
	payload["name"] = d.Get("name").(string)

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

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
	if v, ok := d.GetOk("show_as_ranges"); ok {
		payload["show-as-ranges"] = v.(bool)
	}
	if v, ok := d.GetOkExists("show_hits"); ok {
		payload["show-hits"] = v.(bool)
	}

	if v, ok := d.GetOk("hits_settings"); ok {
		newPayload := make(map[string]interface{})
		tempPayload := v.(map[string]interface{})

		if val, ok := tempPayload["from_date"]; ok {
			newPayload["from-date"] = val
		}
		if val, ok := tempPayload["target"]; ok {
			newPayload["target"] = val
		}
		if val, ok := tempPayload["to_date"]; ok {
			newPayload["to-date"] = val
		}
		payload["hits-settings"] = newPayload
	}

	if v, ok := d.GetOk("dereference_group_members"); ok {
		payload["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOk("show_membership"); ok {
		payload["show-membership"] = v.(bool)
	}
	showRuleBaseRes, err := client.ApiCall("show-https-rulebase", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := ruleBaseJson["name"]; v != nil {
		_ = d.Set("name", v)
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
		ruleBaseDictToReturn := readHttpsRuleBaseField(ruleBaseJson)
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

func readHttpsRuleBaseField(RuleBase map[string]interface{}) []map[string]interface{} {
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

			if v, _ := ruleBaseMap["destination"]; v != nil {
				tempRulebase["destination"] = v
			}

			if v, _ := ruleBaseMap["destination-negate"]; v != nil {
				tempRulebase["destination_negate"] = v.(bool)
			}

			if v, _ := ruleBaseMap["install-on"]; v != nil {
				tempRulebase["install_on"] = v
			}

			if v, _ := ruleBaseMap["enabled"]; v != nil {
				tempRulebase["enabled"] = v.(bool)
			}

			if v, _ := ruleBaseMap["source"]; v != nil {
				tempRulebase["source"] = v
			}

			if v, _ := ruleBaseMap["source-negate"]; v != nil {
				tempRulebase["source_negate"] = v.(bool)
			}

			if v, _ := ruleBaseMap["service"]; v != nil {
				tempRulebase["service"] = v
			}

			if v, _ := ruleBaseMap["service-negate"]; v != nil {
				tempRulebase["service_negate"] = v.(bool)
			}

			if v, _ := ruleBaseMap["type"]; v != nil {
				tempRulebase["type"] = v
			}
			if v, _ := ruleBaseMap["comments"]; v != nil && v != "" {
				tempRulebase["comments"] = v
			}

			if v, _ := ruleBaseMap["action"]; v != nil {
				tempRulebase["action"] = v
			}

			if v, _ := ruleBaseMap["site-category"]; v != nil {
				tempRulebase["site_category"] = v
			}

			if v, _ := ruleBaseMap["site-category-negate"]; v != nil {
				tempRulebase["site_category_negate"] = v.(bool)
			}

			if v := ruleBaseMap["from"]; v != nil && v != 0 {
				tempRulebase["from"] = int(math.Round(v.(float64)))
			}

			if v, _ := ruleBaseMap["to"]; v != nil {
				tempRulebase["to"] = int(math.Round(v.(float64)))
			}

			if v, _ := ruleBaseMap["track"]; v != nil {
				tempRulebase["track"] = v
			}

			if v := ruleBaseMap["rule-number"]; v != nil {
				tempRulebase["rule_number"] = v
			}

			if v, _ := ruleBaseMap["type"]; v != nil {
				tempRulebase["type"] = v
			}

			if v, _ := ruleBaseMap["blade"]; v != nil {
				tempRulebase["blade"] = v
			}

			if v, _ := ruleBaseMap["certificate"]; v != nil {
				tempRulebase["certificate"] = v
			}

			if v, _ := ruleBaseMap["layer"]; v != nil {
				tempRulebase["layer"] = v
			}

			if v, _ := ruleBaseMap["rulebase"]; v != nil {
				tempRulebase["rulebase"] = readHttpsRuleBaseField(ruleBaseMap)
			}

			ruleBaseDictToReturn = append(ruleBaseDictToReturn, tempRulebase)
		}
	}
	return ruleBaseDictToReturn
}
