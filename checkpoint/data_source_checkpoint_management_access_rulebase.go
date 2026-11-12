package checkpoint

import (
	"fmt"
	"log"
	"math"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceManagementAccessRuleBase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessRuleBaseRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"show_as_ranges": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When true, the source, destination and services & applications parameters are displayed as ranges of IP addresses and port numbers rather than network objects.\nObjects that are not represented using IP addresses or port numbers are presented as objects.\nIn addition, the response of each rule does not contain the parameters: source, source-negate, destination, destination-negate, service and service-negate, but instead it contains the parameters: source-ranges, destination-ranges and service-ranges.\n\nNote: Requesting to show rules as ranges is limited up to 20 rules per request, otherwise an error is returned. If you wish to request more rules, use the offset and limit parameters to limit your request.",
				Default:     false,
			},
			"show_hits": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
			},
			"hits_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from_date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.",
						},
						"target": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target gateway name or UID.",
						},
						"to_date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.",
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
									"service_resource": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "service resource.",
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
									"vpn": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Communities or Directional.",
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
									"action_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Action settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable_identity_captive_portal": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "N/A",
												},
												"limit": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "N/A",
												},
											},
										},
									},
									"content": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of processed file types that this rule applies on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"content_direction": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "On which direction the file types processing is applied.",
									},
									"content_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for data.",
									},
									"custom_fields": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Custom fields.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"field_1": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "First custom field.",
												},
												"field_2": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Second custom field.",
												},
												"field_3": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Third custom field.",
												},
											},
										},
									},
									"rule_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of the rule.",
									},
									"inline_layer": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".",
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
									"time": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"track": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Track Settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accounting": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Turns accounting for track on and off.",
												},
												"alert": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Type of alert for the track.",
												},
												"enable_firewall_session": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Determine whether to generate session log to firewall only connections.",
												},
												"per_connection": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Determines whether to perform the log per connection.",
												},
												"per_session": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Determines whether to perform the log per session.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "\"Log\", \"Extended Log\", \"Detailed Log\", \"None\".",
												},
											},
										},
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
												"service_resource": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "service resource.",
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
												"vpn": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Communities or Directional.",
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
												"action_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Action settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enable_identity_captive_portal": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "N/A",
															},
															"limit": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "N/A",
															},
														},
													},
												},
												"content": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "List of processed file types that this rule applies on.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"content_direction": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "On which direction the file types processing is applied.",
												},
												"content_negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True if negate is set for data.",
												},
												"custom_fields": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Custom fields.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"field_1": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First custom field.",
															},
															"field_2": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second custom field.",
															},
															"field_3": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Third custom field.",
															},
														},
													},
												},
												"rule_number": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of the rule.",
												},
												"inline_layer": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".",
												},
												"time": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"track": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Track Settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"accounting": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Turns accounting for track on and off.",
															},
															"alert": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of alert for the track.",
															},
															"enable_firewall_session": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Determine whether to generate session log to firewall only connections.",
															},
															"per_connection": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Determines whether to perform the log per connection.",
															},
															"per_session": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Determines whether to perform the log per session.",
															},
															"type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "\"Log\", \"Extended Log\", \"Detailed Log\", \"None\".",
															},
														},
													},
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

func dataSourceManagementAccessRuleBaseRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{}
	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	if v, ok := d.GetOk("filter"); ok {
		payload["filter"] = v.(string)
	}
	if v, ok := d.GetOk("filter_settings"); ok {
		fsList := v.([]interface{})
		if len(fsList) > 0 {
			fsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("filter_settings.0.search_mode"); ok {
				fsPayload["search-mode"] = v.(string)
			}

			if v, ok := d.GetOk("filter_settings.0.packet_search_settings"); ok {
				pssList := v.([]interface{})
				if len(pssList) > 0 {
					pssPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("filter_settings.0.packet_search_settings.0.expand_group_members"); ok {
						pssPayload["expand-group-members"] = v.(bool)
					}
					if v, ok := d.GetOkExists("filter_settings.0.packet_search_settings.0.expand_group_with_exclusion_members"); ok {
						pssPayload["expand-group-with-exclusion-members"] = v.(bool)
					}
					if v, ok := d.GetOk("filter_settings.0.packet_search_settings.0.intersection_mode_dst"); ok {
						pssPayload["intersection-mode-dst"] = v.(string)
					}
					if v, ok := d.GetOk("filter_settings.0.packet_search_settings.0.intersection_mode_src"); ok {
						pssPayload["intersection-mode-src"] = v.(string)
					}
					if v, ok := d.GetOkExists("filter_settings.0.packet_search_settings.0.match_on_any"); ok {
						pssPayload["match-on-any"] = v.(bool)
					}
					if v, ok := d.GetOkExists("filter_settings.0.packet_search_settings.0.match_on_group_with_exclusion"); ok {
						pssPayload["match-on-group-with-exclusion"] = v.(bool)
					}
					if v, ok := d.GetOkExists("filter_settings.0.packet_search_settings.0.match_on_negate"); ok {
						pssPayload["match-on-negate"] = v.(bool)
					}

					fsPayload["packet-search-settings"] = pssPayload
				}
			}

			payload["filter-settings"] = fsPayload
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
	showRuleBaseRes, err := client.ApiCall("show-access-rulebase", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRuleBaseRes.Success {
		return fmt.Errorf("%s", showRuleBaseRes.ErrorMsg)
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
		ruleBaseDictToReturn := readAccessRuleBaseField(ruleBaseJson)
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

func readAccessRuleBaseField(RuleBase map[string]interface{}) []map[string]interface{} {
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

			if v, _ := ruleBaseMap["service-resource"]; v != nil {
				tempRulebase["service_resource"] = v
			}

			if v, _ := ruleBaseMap["vpn"]; v != nil {
				tempRulebase["vpn"] = v
			}

			if v, _ := ruleBaseMap["action"]; v != nil {
				tempRulebase["action"] = v
			}

			if ruleBaseMap["action-settings"] != nil {

				actionSettingsMap := ruleBaseMap["action-settings"].(map[string]interface{})

				actionSettingsMapToReturn := make(map[string]interface{})

				if v := actionSettingsMap["enable-identity-captive-portal"]; v != nil {
					actionSettingsMapToReturn["enable_identity_captive_portal"] = v
				}
				if v := actionSettingsMap["limit"]; v != nil {
					actionSettingsMapToReturn["limit"] = v
				}
				tempRulebase["action_settings"] = []interface{}{actionSettingsMapToReturn}
			}
			if v, _ := ruleBaseMap["content"]; v != nil {
				tempRulebase["content"] = v
			}

			if v, _ := ruleBaseMap["content-negate"]; v != nil {
				tempRulebase["content_negate"] = v.(bool)
			}

			if v, _ := ruleBaseMap["content-direction"]; v != nil {
				tempRulebase["content_direction"] = v
			}

			if v, _ := ruleBaseMap["time"]; v != nil {
				tempRulebase["time"] = v
			}

			if v := ruleBaseMap["from"]; v != nil && v != 0 {
				tempRulebase["from"] = int(math.Round(v.(float64)))
			}

			if v, _ := ruleBaseMap["to"]; v != nil {
				tempRulebase["to"] = int(math.Round(v.(float64)))
			}

			if ruleBaseMap["track"] != nil {

				trackMap := ruleBaseMap["track"].(map[string]interface{})

				trackMapToReturn := make(map[string]interface{})
				if v := trackMap["accounting"]; v != nil {
					trackMapToReturn["accounting"] = v.(bool)
				}

				if v, _ := trackMap["alert"]; v != nil {
					trackMapToReturn["alert"] = v.(string)
				}

				if v := trackMap["enable-firewall-session"]; v != nil {
					trackMapToReturn["enable_firewall_session"] = v.(bool)
				}

				if v := trackMap["per-connection"]; v != nil {
					trackMapToReturn["per_connection"] = v.(bool)
				}

				if v := trackMap["per-session"]; v != nil {
					trackMapToReturn["per_session"] = v.(bool)
				}

				if v, _ := trackMap["type"]; v != nil {
					trackMapToReturn["type"] = v.(string)
				}

				tempRulebase["track"] = []interface{}{trackMapToReturn}
			}

			if ruleBaseMap["custom-fields"] != nil {

				customFieldsMap := ruleBaseMap["custom-fields"].(map[string]interface{})

				customFieldsMapToReturn := make(map[string]interface{})

				if v, _ := customFieldsMap["field-1"]; v != nil {
					customFieldsMapToReturn["field_1"] = v
				}

				if v, _ := customFieldsMap["field-2"]; v != nil {
					customFieldsMapToReturn["field_2"] = v
				}

				if v, _ := customFieldsMap["field-3"]; v != nil {
					customFieldsMapToReturn["field_3"] = v
				}
				tempRulebase["custom_fields"] = []interface{}{customFieldsMapToReturn}
			}

			if v := ruleBaseMap["rule-number"]; v != nil {
				tempRulebase["rule_number"] = v
			}

			if v := ruleBaseMap["inline-layer"]; v != nil {
				tempRulebase["inline_layer"] = v
			}

			if v, _ := ruleBaseMap["type"]; v != nil {
				tempRulebase["type"] = v
			}

			if v, _ := ruleBaseMap["rulebase"]; v != nil {
				tempRulebase["rulebase"] = readAccessRuleBaseField(ruleBaseMap)
			}

			ruleBaseDictToReturn = append(ruleBaseDictToReturn, tempRulebase)
		}
	}
	return ruleBaseDictToReturn
}
