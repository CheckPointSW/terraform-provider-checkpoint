package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementBestPractice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementBestPracticeRead,
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
			"best_practice_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice ID.",
			},
			"show_regulations": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Show the applicable regulations of the Best Practice.",
				Default:     false,
			},
			"action_item": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Required action item to comply with the Best Practice.",
			},
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Shows if the Best Practice is active.",
			},
			"blade": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Software Blade name of the Best Practice.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Best Practice.",
			},
			"due_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Shows if there is a due date for the action item of this Best Practice.",
			},
			"regulations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The applicable regulations of the Best Practice. Appears only when the value of the 'show-regulations' parameter is set to 'true'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regulation_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the regulation.",
						},
						"requirement_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the requirement.",
						},
						"requirement_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the requirement.",
						},
						"requirement_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the requirement.",
						},
						"requirement_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the requirement.",
						},
					},
				},
			},
			"relevant_objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The applicable objects of the Best Practice.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_rules_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information about the relevant access rules. Appears only when the value of the 'relevant-objects-type' parameter is 'access-rule'.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Shows if the Compliance scan is enabled or not for this object.",
									},
									"layer_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the relevant policy layer.",
									},
									"layer_uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UID of the relevant policy layer.",
									},
									"policy_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the relevant policy.",
									},
									"rule_indexes": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Comma-separated indexes of the relevant rules in the relevant policy and policy layer.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the relevant object.",
									},
								},
							},
						},
						"cpm_relevant_objects_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information about the relevant objects. Appears only when the value of the 'relevant-objects-type' parameter is 'cpm-relevant-object'.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpm_relevant_object_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the relevant object.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Shows if the Compliance scan is enabled or not for this object.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the relevant object.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the relevant object.",
									},
								},
							},
						},
						"ips_protections_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information about the relevant ips-protection objects. Appears only when the value of the 'relevant-objects-type' parameter is 'ips-protection'.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The current action of the Threat Prevention profile.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Shows if the Compliance scan is enabled or not for this object.",
									},
									"profile_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the relevant Threat Prevention profile.",
									},
									"profile_uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UID of the relevant Threat Prevention profile.",
									},
									"protection_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the relevant IPS protection.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the relevant object.",
									},
								},
							},
						},
						"relevant_objects_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the relevant object.",
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the Best Practice.",
			},
			"user_defined": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Shows if the Best Practice is a user-defined Best Practice.",
			},
			"user_defined_firewall": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The definitions of the user-defined Firewall Best Practice. Relevant only for Firewall Best Practices created by the user.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_range_percentage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "User-defined policy range percentage to test.",
						},
						"policy_range_position": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined policy range position.",
						},
						"poor_condition": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined poor condition.",
						},
						"secure_condition": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined secure condition.",
						},
						"tolerance": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "User-defined tolerance. Appears only when the value of the 'violation-condition' parameter is 'Rule found'.",
						},
						"user_defined_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "User-defined Firewall rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined actions.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"comment": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined comment.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"condition_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The condition type.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.",
												},
											},
										},
									},
									"destination": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined destination objects.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"hit_count": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined hit count value.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"install_on": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined \"Install On\" objects.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"name": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined name.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"condition_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The condition type.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.",
												},
											},
										},
									},
									"services_and_applications": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined service and application objects.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"source": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined source objects.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"time": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined time.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"track": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined track actions.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
															},
														},
													},
												},
											},
										},
									},
									"vpn": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "User-defined VPN objects.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"negate": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Shows if the rule is negated.",
												},
												"reference_objects": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reference objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the reference object.",
															},
															"reference_object_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the reference object.",
															},
															"uid": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The UID of the reference object.",
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
						"violation_condition": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined violation condition.",
						},
					},
				},
			},
			"user_defined_gaia_os": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The definitions of the user-defined Gaia OS Best Practice. Relevant only for Gaia OS Best Practices created by the user.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expected_output_base64": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected output of the script in the Base64.",
						},
						"practice_script_base64": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The script in Base64 to run on Gaia Security Gateways during the Compliance scans.",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementBestPracticeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else if v, ok := d.GetOk("best_practice_id"); ok {
		payload["best-practice-id"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid or best_practice_id must be specified")
	}

	if v, ok := d.GetOkExists("show_regulations"); ok {
		payload["show-regulations"] = v.(bool)
	}

	showBestPracticeRes, err := client.ApiCall("show-best-practice", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showBestPracticeRes.Success {
		return fmt.Errorf(showBestPracticeRes.ErrorMsg)
	}

	bestPractice := showBestPracticeRes.GetData()

	log.Println("Read BestPractice - Show JSON = ", bestPractice)

	if v := bestPractice["uid"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("uid", v)
	}

	if v := bestPractice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := bestPractice["best-practice-id"]; v != nil {
		_ = d.Set("best_practice_id", v)
	}

	if v := bestPractice["action-item"]; v != nil {
		_ = d.Set("action_item", v)
	}

	if v := bestPractice["active"]; v != nil {
		_ = d.Set("active", v)
	}

	if v := bestPractice["blade"]; v != nil {
		_ = d.Set("blade", v)
	}

	if v := bestPractice["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := bestPractice["due-date"]; v != nil {
		_ = d.Set("due_date", v)
	}

	if bestPractice["regulations"] != nil {

		regulationsList := bestPractice["regulations"].([]interface{})

		if len(regulationsList) > 0 {

			var regulationsListToReturn []map[string]interface{}

			for i := range regulationsList {

				regulationsMap := regulationsList[i].(map[string]interface{})

				regulationsMapToAdd := make(map[string]interface{})

				if v, _ := regulationsMap["regulation-name"]; v != nil {
					regulationsMapToAdd["regulation_name"] = v
				}
				if v, _ := regulationsMap["requirement-description"]; v != nil {
					regulationsMapToAdd["requirement_description"] = v
				}
				if v, _ := regulationsMap["requirement-id"]; v != nil {
					regulationsMapToAdd["requirement_id"] = v
				}
				if v, _ := regulationsMap["requirement-status"]; v != nil {
					regulationsMapToAdd["requirement_status"] = v
				}
				if v, _ := regulationsMap["requirement-uid"]; v != nil {
					regulationsMapToAdd["requirement_uid"] = v
				}
				regulationsListToReturn = append(regulationsListToReturn, regulationsMapToAdd)
			}

			_ = d.Set("regulations", regulationsListToReturn)
		} else {
			_ = d.Set("regulations", regulationsList)
		}
	} else {
		_ = d.Set("regulations", nil)
	}

	if bestPractice["relevant-objects"] != nil {

		relevantObjectsMap := bestPractice["relevant-objects"].(map[string]interface{})

		relevantObjectsMapToReturn := make(map[string]interface{})

		if v := relevantObjectsMap["access-rules-info"]; v != nil {

			accessRulesInfoList := v.([]interface{})

			if len(accessRulesInfoList) > 0 {

				var accessRulesInfoListToReturn []map[string]interface{}

				for i := range accessRulesInfoList {

					accessRulesInfoMap := accessRulesInfoList[i].(map[string]interface{})

					accessRulesInfoMapToAdd := make(map[string]interface{})

					if v, _ := accessRulesInfoMap["enabled"]; v != nil {
						accessRulesInfoMapToAdd["enabled"] = v
					}
					if v, _ := accessRulesInfoMap["layer-name"]; v != nil {
						accessRulesInfoMapToAdd["layer_name"] = v
					}
					if v, _ := accessRulesInfoMap["layer-uid"]; v != nil {
						accessRulesInfoMapToAdd["layer_uid"] = v
					}
					if v, _ := accessRulesInfoMap["policy-name"]; v != nil {
						accessRulesInfoMapToAdd["policy_name"] = v
					}
					if v, _ := accessRulesInfoMap["rule-indexes"]; v != nil {
						accessRulesInfoMapToAdd["rule_indexes"] = v
					}
					if v, _ := accessRulesInfoMap["status"]; v != nil {
						accessRulesInfoMapToAdd["status"] = v
					}
					accessRulesInfoListToReturn = append(accessRulesInfoListToReturn, accessRulesInfoMapToAdd)
				}
				relevantObjectsMapToReturn["access_rules_info"] = accessRulesInfoListToReturn
			} else {
				relevantObjectsMapToReturn["access_rules_info"] = accessRulesInfoList
			}
		} else {
			relevantObjectsMapToReturn["access_rules_info"] = nil
		}

		if v := relevantObjectsMap["cpm-relevant-objects-info"]; v != nil {

			cpmRelevantObjectsInfoList := v.([]interface{})

			if len(cpmRelevantObjectsInfoList) > 0 {

				var cpmRelevantObjectsInfoListToReturn []map[string]interface{}

				for i := range cpmRelevantObjectsInfoList {

					cpmRelevantObjectsInfoMap := cpmRelevantObjectsInfoList[i].(map[string]interface{})

					cpmRelevantObjectsInfoMapToAdd := make(map[string]interface{})

					if v, _ := cpmRelevantObjectsInfoMap["cpm-relevant-object-type"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["cpm_relevant_object_type"] = v
					}
					if v, _ := cpmRelevantObjectsInfoMap["enabled"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["enabled"] = v
					}
					if v, _ := cpmRelevantObjectsInfoMap["name"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["name"] = v
					}
					if v, _ := cpmRelevantObjectsInfoMap["status"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["status"] = v
					}
					cpmRelevantObjectsInfoListToReturn = append(cpmRelevantObjectsInfoListToReturn, cpmRelevantObjectsInfoMapToAdd)
				}
				relevantObjectsMapToReturn["cpm_relevant_objects_info"] = cpmRelevantObjectsInfoListToReturn
			} else {
				relevantObjectsMapToReturn["cpm_relevant_objects_info"] = cpmRelevantObjectsInfoList
			}
		} else {
			relevantObjectsMapToReturn["cpm_relevant_objects_info"] = nil
		}

		if v := relevantObjectsMap["ips-protections-info"]; v != nil {

			ipsProtectionsInfoList := v.([]interface{})

			if len(ipsProtectionsInfoList) > 0 {

				var ipsProtectionsInfoListToReturn []map[string]interface{}

				for i := range ipsProtectionsInfoList {

					ipsProtectionsInfoMap := ipsProtectionsInfoList[i].(map[string]interface{})

					ipsProtectionsInfoMapToAdd := make(map[string]interface{})

					if v, _ := ipsProtectionsInfoMap["action"]; v != nil {
						ipsProtectionsInfoMapToAdd["action"] = v
					}
					if v, _ := ipsProtectionsInfoMap["enabled"]; v != nil {
						ipsProtectionsInfoMapToAdd["enabled"] = v
					}
					if v, _ := ipsProtectionsInfoMap["profile-name"]; v != nil {
						ipsProtectionsInfoMapToAdd["profile_name"] = v
					}
					if v, _ := ipsProtectionsInfoMap["profile-uid"]; v != nil {
						ipsProtectionsInfoMapToAdd["profile_uid"] = v
					}
					if v, _ := ipsProtectionsInfoMap["protection-name"]; v != nil {
						ipsProtectionsInfoMapToAdd["protection_name"] = v
					}
					if v, _ := ipsProtectionsInfoMap["status"]; v != nil {
						ipsProtectionsInfoMapToAdd["status"] = v
					}
					ipsProtectionsInfoListToReturn = append(ipsProtectionsInfoListToReturn, ipsProtectionsInfoMapToAdd)
				}
				relevantObjectsMapToReturn["ips_protections_info"] = ipsProtectionsInfoListToReturn
			} else {
				relevantObjectsMapToReturn["ips_protections_info"] = ipsProtectionsInfoList
			}
		} else {
			relevantObjectsMapToReturn["ips_protections_info"] = nil
		}

		if v := relevantObjectsMap["relevant-objects-type"]; v != nil {
			relevantObjectsMapToReturn["relevant_objects_type"] = v
		}
		_ = d.Set("relevant_objects", []interface{}{relevantObjectsMapToReturn})

	} else {
		_ = d.Set("relevant_objects", nil)
	}

	if v := bestPractice["status"]; v != nil {
		_ = d.Set("status", v)
	}

	if v := bestPractice["user-defined"]; v != nil {
		_ = d.Set("user_defined", v)
	}

	if bestPractice["user-defined-firewall"] != nil {

		userDefinedFirewallMap := bestPractice["user-defined-firewall"].(map[string]interface{})

		userDefinedFirewallMapToReturn := make(map[string]interface{})

		if v := userDefinedFirewallMap["policy-range-percentage"]; v != nil {
			userDefinedFirewallMapToReturn["policy_range_percentage"] = v
		}
		if v := userDefinedFirewallMap["policy-range-position"]; v != nil {
			userDefinedFirewallMapToReturn["policy_range_position"] = v
		}
		if v := userDefinedFirewallMap["poor-condition"]; v != nil {
			userDefinedFirewallMapToReturn["poor_condition"] = v
		}
		if v := userDefinedFirewallMap["secure-condition"]; v != nil {
			userDefinedFirewallMapToReturn["secure_condition"] = v
		}
		if v := userDefinedFirewallMap["tolerance"]; v != nil {
			userDefinedFirewallMapToReturn["tolerance"] = v
		}
		if v := userDefinedFirewallMap["user-defined-rules"]; v != nil {

			userDefinedRulesList := v.([]interface{})

			if len(userDefinedRulesList) > 0 {

				var userDefinedRulesListToReturn []map[string]interface{}

				for i := range userDefinedRulesList {

					userDefinedRulesMap := userDefinedRulesList[i].(map[string]interface{})

					userDefinedRulesMapToAdd := make(map[string]interface{})

					if v := userDefinedRulesMap["action"]; v != nil {

						actionMap := v.(map[string]interface{})
						actionMapToReturn := make(map[string]interface{})

						if v, _ := actionMap["negate"]; v != nil {
							actionMapToReturn["negate"] = v
						}
						if v := actionMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								actionMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								actionMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							actionMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["action"] = actionMapToReturn
					}

					if v := userDefinedRulesMap["comment"]; v != nil {

						commentMap := v.(map[string]interface{})
						commentMapToReturn := make(map[string]interface{})

						if v, _ := commentMap["condition-type"]; v != nil {
							commentMapToReturn["condition_type"] = v
						}
						if v, _ := commentMap["value"]; v != nil {
							commentMapToReturn["value"] = v
						}
						userDefinedRulesMapToAdd["comment"] = commentMapToReturn
					}

					if v := userDefinedRulesMap["destination"]; v != nil {

						destinationMap := v.(map[string]interface{})
						destinationMapToReturn := make(map[string]interface{})

						if v, _ := destinationMap["negate"]; v != nil {
							destinationMapToReturn["negate"] = v
						}
						if v := destinationMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								destinationMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								destinationMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							destinationMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["destination"] = destinationMapToReturn
					}

					if v := userDefinedRulesMap["hit-count"]; v != nil {

						hitCountMap := v.(map[string]interface{})
						hitCountMapToReturn := make(map[string]interface{})

						if v, _ := hitCountMap["negate"]; v != nil {
							hitCountMapToReturn["negate"] = v
						}
						if v := hitCountMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								hitCountMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								hitCountMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							hitCountMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["hit_count"] = hitCountMapToReturn
					}

					if v := userDefinedRulesMap["install-on"]; v != nil {

						installOnMap := v.(map[string]interface{})
						installOnMapToReturn := make(map[string]interface{})

						if v, _ := installOnMap["negate"]; v != nil {
							installOnMapToReturn["negate"] = v
						}
						if v := installOnMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								installOnMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								installOnMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							installOnMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["install_on"] = installOnMapToReturn
					}

					if v := userDefinedRulesMap["name"]; v != nil {

						nameMap := v.(map[string]interface{})
						nameMapToReturn := make(map[string]interface{})

						if v, _ := nameMap["condition-type"]; v != nil {
							nameMapToReturn["condition_type"] = v
						}
						if v, _ := nameMap["value"]; v != nil {
							nameMapToReturn["value"] = v
						}
						userDefinedRulesMapToAdd["name"] = nameMapToReturn
					}

					if v := userDefinedRulesMap["services-and-applications"]; v != nil {

						servicesAndApplicationsMap := v.(map[string]interface{})
						servicesAndApplicationsMapToReturn := make(map[string]interface{})

						if v, _ := servicesAndApplicationsMap["negate"]; v != nil {
							servicesAndApplicationsMapToReturn["negate"] = v
						}
						if v := servicesAndApplicationsMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								servicesAndApplicationsMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								servicesAndApplicationsMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							servicesAndApplicationsMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["services_and_applications"] = servicesAndApplicationsMapToReturn
					}

					if v := userDefinedRulesMap["source"]; v != nil {

						sourceMap := v.(map[string]interface{})
						sourceMapToReturn := make(map[string]interface{})

						if v, _ := sourceMap["negate"]; v != nil {
							sourceMapToReturn["negate"] = v
						}
						if v := sourceMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								sourceMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								sourceMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							sourceMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["source"] = sourceMapToReturn
					}

					if v := userDefinedRulesMap["time"]; v != nil {

						timeMap := v.(map[string]interface{})
						timeMapToReturn := make(map[string]interface{})

						if v, _ := timeMap["negate"]; v != nil {
							timeMapToReturn["negate"] = v
						}
						if v := timeMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								timeMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								timeMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							timeMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["time"] = timeMapToReturn
					}

					if v := userDefinedRulesMap["track"]; v != nil {

						trackMap := v.(map[string]interface{})
						trackMapToReturn := make(map[string]interface{})

						if v, _ := trackMap["negate"]; v != nil {
							trackMapToReturn["negate"] = v
						}
						if v := trackMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								trackMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								trackMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							trackMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["track"] = trackMapToReturn
					}

					if v := userDefinedRulesMap["vpn"]; v != nil {

						vpnMap := v.(map[string]interface{})
						vpnMapToReturn := make(map[string]interface{})

						if v, _ := vpnMap["negate"]; v != nil {
							vpnMapToReturn["negate"] = v
						}
						if v := vpnMap["reference-objects"]; v != nil {

							referenceObjectsList := v.([]interface{})

							if len(referenceObjectsList) > 0 {

								var referenceObjectsListToReturn []map[string]interface{}

								for i := range referenceObjectsList {

									referenceObjectsMap := referenceObjectsList[i].(map[string]interface{})

									referenceObjectsMapToAdd := make(map[string]interface{})

									if v, _ := referenceObjectsMap["name"]; v != nil {
										referenceObjectsMapToAdd["name"] = v
									}
									if v, _ := referenceObjectsMap["reference-object-type"]; v != nil {
										referenceObjectsMapToAdd["reference_object_type"] = v
									}
									if v, _ := referenceObjectsMap["uid"]; v != nil {
										referenceObjectsMapToAdd["uid"] = v
									}
									referenceObjectsListToReturn = append(referenceObjectsListToReturn, referenceObjectsMapToAdd)
								}
								vpnMapToReturn["reference_objects"] = referenceObjectsListToReturn
							} else {
								vpnMapToReturn["reference_objects"] = referenceObjectsList
							}
						} else {
							vpnMapToReturn["reference_objects"] = nil
						}

						userDefinedRulesMapToAdd["vpn"] = vpnMapToReturn
					}

					userDefinedRulesListToReturn = append(userDefinedRulesListToReturn, userDefinedRulesMapToAdd)
				}
				userDefinedFirewallMapToReturn["user_defined_rules"] = userDefinedRulesListToReturn
			} else {
				userDefinedFirewallMapToReturn["user_defined_rules"] = userDefinedRulesList
			}
		} else {
			userDefinedFirewallMapToReturn["user_defined_rules"] = nil
		}

		if v := userDefinedFirewallMap["violation-condition"]; v != nil {
			userDefinedFirewallMapToReturn["violation_condition"] = v
		}
		_ = d.Set("user_defined_firewall", []interface{}{userDefinedFirewallMapToReturn})

	} else {
		_ = d.Set("user_defined_firewall", nil)
	}

	if bestPractice["user-defined-gaia-os"] != nil {

		userDefinedGaiaOsMap := bestPractice["user-defined-gaia-os"].(map[string]interface{})

		userDefinedGaiaOsMapToReturn := make(map[string]interface{})

		if v := userDefinedGaiaOsMap["expected-output-base64"]; v != nil {
			userDefinedGaiaOsMapToReturn["expected_output_base64"] = v
		}
		if v := userDefinedGaiaOsMap["practice-script-base64"]; v != nil {
			userDefinedGaiaOsMapToReturn["practice_script_base64"] = v
		}
		_ = d.Set("user_defined_gaia_os", []interface{}{userDefinedGaiaOsMapToReturn})

	} else {
		_ = d.Set("user_defined_gaia_os", nil)
	}

	if v := bestPractice["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if bestPractice["meta-info"] != nil {

		metaInfoMap := bestPractice["meta-info"].(map[string]interface{})

		metaInfoMapToReturn := make(map[string]interface{})

		if v := metaInfoMap["creation-time"]; v != nil {

			creationTimeMap := v.(map[string]interface{})
			creationTimeMapToReturn := make(map[string]interface{})

			if v, _ := creationTimeMap["iso-8601"]; v != nil {
				creationTimeMapToReturn["iso_8601"] = v
			}
			if v, _ := creationTimeMap["posix"]; v != nil {
				creationTimeMapToReturn["posix"] = v
			}
			metaInfoMapToReturn["creation_time"] = []interface{}{creationTimeMapToReturn}
		}

		if v := metaInfoMap["creator"]; v != nil {
			metaInfoMapToReturn["creator"] = v
		}
		if v := metaInfoMap["last-modifier"]; v != nil {
			metaInfoMapToReturn["last_modifier"] = v
		}
		if v := metaInfoMap["last-modify-time"]; v != nil {

			lastModifyTimeMap := v.(map[string]interface{})
			lastModifyTimeMapToReturn := make(map[string]interface{})

			if v, _ := lastModifyTimeMap["iso-8601"]; v != nil {
				lastModifyTimeMapToReturn["iso_8601"] = v
			}
			if v, _ := lastModifyTimeMap["posix"]; v != nil {
				lastModifyTimeMapToReturn["posix"] = v
			}
			metaInfoMapToReturn["last_modify_time"] = []interface{}{lastModifyTimeMapToReturn}
		}

		if v := metaInfoMap["lock"]; v != nil {
			metaInfoMapToReturn["lock"] = v
		}
		if v := metaInfoMap["locking-admin"]; v != nil {
			metaInfoMapToReturn["locking_admin"] = v
		}
		if v := metaInfoMap["locking-session-id"]; v != nil {
			metaInfoMapToReturn["locking_session_id"] = v
		}
		if v := metaInfoMap["validation-state"]; v != nil {
			metaInfoMapToReturn["validation_state"] = v
		}
		_ = d.Set("meta_info", []interface{}{metaInfoMapToReturn})

	} else {
		_ = d.Set("meta_info", nil)
	}

	if v := bestPractice["read-only"]; v != nil {
		_ = d.Set("read_only", v)
	}

	if bestPractice["available-actions"] != nil {

		availableActionsMap := bestPractice["available-actions"].(map[string]interface{})

		availableActionsMapToReturn := make(map[string]interface{})

		if v := availableActionsMap["clone"]; v != nil {
			availableActionsMapToReturn["clone"] = v
		}
		if v := availableActionsMap["delete"]; v != nil {
			availableActionsMapToReturn["delete"] = v
		}
		if v := availableActionsMap["edit"]; v != nil {
			availableActionsMapToReturn["edit"] = v
		}
		_ = d.Set("available_actions", []interface{}{availableActionsMapToReturn})

	} else {
		_ = d.Set("available_actions", nil)
	}

	if v := bestPractice["show-regulations"]; v != nil {
		_ = d.Set("show_regulations", v)
	}

	return nil

}
