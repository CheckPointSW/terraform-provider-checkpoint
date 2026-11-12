package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementFirewallBestPractice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementFirewallBestPracticeRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice Name.",
			},
			"best_practice_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice ID.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"show_regulations": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Show the applicable regulations of the Best Practice.",
			},
			"show_relevant_objects": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Show the relevant objects of the Best Practice.",
			},
			"action_item": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Required action item to comply with the Best Practice.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Best Practice.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The activation status of the best practice.",
			},
			"expiration": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Deactivation expiration settings. Present only when the best practice is disabled.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason the best practice was deactivated.",
						},
						"expire_on": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "When the deactivation expires. Date and time represented in international ISO 8601 format.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the deactivation never expires or expires on a specific date.",
						},
					},
				},
			},
			"policy_range_percentage": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Percentage of the rule base to scan, 0-100.",
			},
			"policy_range_position": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The direction of the scan.",
			},
			"poor_condition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility of poor-result rules in the Relevant Objects pane.",
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
				Description: "The applicable objects of the Best Practice. Appears only when the value of the 'show-relevant-objects' parameter is set to 'true'.",
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
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The UID of the relevant object.",
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
			"rule": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rule criteria the firewall best practice evaluates against the rule base.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Network objects to match in the rule Source column.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_source": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the source values are negated.",
						},
						"destination": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Network objects to match in the rule Destination column.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_destination": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the destination values are negated.",
						},
						"vpn": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "VPN communities to match.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_vpn": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the vpn values are negated.",
						},
						"services_and_applications": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Services, applications, categories or sites to match.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_services_and_applications": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the services and applications values are negated.",
						},
						"install_on": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Security Gateways or Clusters the rule applies to.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_install_on": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the install-on values are negated.",
						},
						"time": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Time objects the rule applies to.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_time": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the time values are negated.",
						},
						"action": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Rule actions to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_action": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the action values are negated.",
						},
						"track": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Tracking methods to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_track": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the track values are negated.",
						},
						"hit_count": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Hit-count levels to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_hit_count": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Shows if the hit-count values are negated.",
						},
						"name_condition": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Match the rule name against a text condition.",
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
						"comment_condition": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Match the rule comment against a text condition.",
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
					},
				},
			},
			"secure_condition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility of secure-result rules in the Relevant Objects pane.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the Best Practice.",
			},
			"tolerance": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of matches allowed before a violation is created. Relevant only when violation-condition is set to 'Rule found'.",
			},
			"violation_condition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Define when a violation occurs: 'Rule found' means the criteria match a rule; 'Rule not found' means no rule matches.",
			},
		},
	}
}

func dataSourceManagementFirewallBestPracticeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	if v, ok := d.GetOk("best_practice_id"); ok {
		payload["best-practice-id"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_regulations"); ok {
		payload["show-regulations"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_relevant_objects"); ok {
		payload["show-relevant-objects"] = v.(bool)
	}

	showFirewallBestPracticeRes, err := client.ApiCall("show-firewall-best-practice", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showFirewallBestPracticeRes.Success {
		if objectNotFound(showFirewallBestPracticeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showFirewallBestPracticeRes.ErrorMsg)
	}

	firewallBestPractice := showFirewallBestPracticeRes.GetData()

	log.Println("Read FirewallBestPractice - Show JSON = ", firewallBestPractice)

	if v := firewallBestPractice["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := firewallBestPractice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := firewallBestPractice["action-item"]; v != nil {
		_ = d.Set("action_item", v)
	}

	if v := firewallBestPractice["best-practice-id"]; v != nil {
		_ = d.Set("best_practice_id", v)
	}

	if v := firewallBestPractice["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := firewallBestPractice["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if firewallBestPractice["expiration"] != nil {

		expirationMap := firewallBestPractice["expiration"].(map[string]interface{})

		expirationMapToReturn := make(map[string]interface{})

		if v := expirationMap["comment"]; v != nil {
			expirationMapToReturn["comment"] = v
		}
		if v := expirationMap["expire-on"]; v != nil {

			expireOnMap := v.(map[string]interface{})

			expireOnMapToReturn := make(map[string]interface{})

			if v := expireOnMap["iso-8601"]; v != nil {
				expireOnMapToReturn["iso_8601"] = v
			}
			if v := expireOnMap["posix"]; v != nil {
				expireOnMapToReturn["posix"] = v
			}

			expirationMapToReturn["expire_on"] = []interface{}{expireOnMapToReturn}
		}

		if v := expirationMap["mode"]; v != nil {
			expirationMapToReturn["mode"] = v
		}

		_ = d.Set("expiration", []interface{}{expirationMapToReturn})

	} else {
		_ = d.Set("expiration", nil)
	}

	if v := firewallBestPractice["policy-range-percentage"]; v != nil {
		_ = d.Set("policy_range_percentage", v)
	}

	if v := firewallBestPractice["policy-range-position"]; v != nil {
		_ = d.Set("policy_range_position", v)
	}

	if v := firewallBestPractice["poor-condition"]; v != nil {
		_ = d.Set("poor_condition", v)
	}

	if firewallBestPractice["regulations"] != nil {

		regulationsList := firewallBestPractice["regulations"].([]interface{})

		if len(regulationsList) > 0 {

			var regulationsListToReturn []map[string]interface{}

			for i := range regulationsList {

				regulationsMap := regulationsList[i].(map[string]interface{})

				regulationsMapToAdd := make(map[string]interface{})

				if v := regulationsMap["regulation-name"]; v != nil {
					regulationsMapToAdd["regulation_name"] = v
				}
				if v := regulationsMap["requirement-description"]; v != nil {
					regulationsMapToAdd["requirement_description"] = v
				}
				if v := regulationsMap["requirement-id"]; v != nil {
					regulationsMapToAdd["requirement_id"] = v
				}
				if v := regulationsMap["requirement-status"]; v != nil {
					regulationsMapToAdd["requirement_status"] = v
				}
				if v := regulationsMap["requirement-uid"]; v != nil {
					regulationsMapToAdd["requirement_uid"] = v
				}

				regulationsListToReturn = append(regulationsListToReturn, regulationsMapToAdd)
			}

			_ = d.Set("regulations", regulationsListToReturn)
		}
	} else {
		_ = d.Set("regulations", nil)
	}

	if firewallBestPractice["relevant-objects"] != nil {

		relevantObjectsMap := firewallBestPractice["relevant-objects"].(map[string]interface{})

		relevantObjectsMapToReturn := make(map[string]interface{})

		if v := relevantObjectsMap["access-rules-info"]; v != nil {

			accessRulesInfoList := v.([]interface{})

			if len(accessRulesInfoList) > 0 {

				var accessRulesInfoListToReturn []map[string]interface{}

				for i := range accessRulesInfoList {

					accessRulesInfoMap := accessRulesInfoList[i].(map[string]interface{})

					accessRulesInfoMapToAdd := make(map[string]interface{})

					if v := accessRulesInfoMap["enabled"]; v != nil {
						accessRulesInfoMapToAdd["enabled"] = v
					}
					if v := accessRulesInfoMap["layer-name"]; v != nil {
						accessRulesInfoMapToAdd["layer_name"] = v
					}
					if v := accessRulesInfoMap["layer-uid"]; v != nil {
						accessRulesInfoMapToAdd["layer_uid"] = v
					}
					if v := accessRulesInfoMap["policy-name"]; v != nil {
						accessRulesInfoMapToAdd["policy_name"] = v
					}
					if v := accessRulesInfoMap["rule-indexes"]; v != nil {
						accessRulesInfoMapToAdd["rule_indexes"] = v
					}
					if v := accessRulesInfoMap["status"]; v != nil {
						accessRulesInfoMapToAdd["status"] = v
					}

					accessRulesInfoListToReturn = append(accessRulesInfoListToReturn, accessRulesInfoMapToAdd)
				}

				relevantObjectsMapToReturn["access_rules_info"] = accessRulesInfoListToReturn
			}
		}

		if v := relevantObjectsMap["cpm-relevant-objects-info"]; v != nil {

			cpmRelevantObjectsInfoList := v.([]interface{})

			if len(cpmRelevantObjectsInfoList) > 0 {

				var cpmRelevantObjectsInfoListToReturn []map[string]interface{}

				for i := range cpmRelevantObjectsInfoList {

					cpmRelevantObjectsInfoMap := cpmRelevantObjectsInfoList[i].(map[string]interface{})

					cpmRelevantObjectsInfoMapToAdd := make(map[string]interface{})

					if v := cpmRelevantObjectsInfoMap["cpm-relevant-object-type"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["cpm_relevant_object_type"] = v
					}
					if v := cpmRelevantObjectsInfoMap["enabled"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["enabled"] = v
					}
					if v := cpmRelevantObjectsInfoMap["name"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["name"] = v
					}
					if v := cpmRelevantObjectsInfoMap["status"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["status"] = v
					}
					if v := cpmRelevantObjectsInfoMap["uid"]; v != nil {
						cpmRelevantObjectsInfoMapToAdd["uid"] = v
					}

					cpmRelevantObjectsInfoListToReturn = append(cpmRelevantObjectsInfoListToReturn, cpmRelevantObjectsInfoMapToAdd)
				}

				relevantObjectsMapToReturn["cpm_relevant_objects_info"] = cpmRelevantObjectsInfoListToReturn
			}
		}

		if v := relevantObjectsMap["ips-protections-info"]; v != nil {

			ipsProtectionsInfoList := v.([]interface{})

			if len(ipsProtectionsInfoList) > 0 {

				var ipsProtectionsInfoListToReturn []map[string]interface{}

				for i := range ipsProtectionsInfoList {

					ipsProtectionsInfoMap := ipsProtectionsInfoList[i].(map[string]interface{})

					ipsProtectionsInfoMapToAdd := make(map[string]interface{})

					if v := ipsProtectionsInfoMap["action"]; v != nil {
						ipsProtectionsInfoMapToAdd["action"] = v
					}
					if v := ipsProtectionsInfoMap["enabled"]; v != nil {
						ipsProtectionsInfoMapToAdd["enabled"] = v
					}
					if v := ipsProtectionsInfoMap["profile-name"]; v != nil {
						ipsProtectionsInfoMapToAdd["profile_name"] = v
					}
					if v := ipsProtectionsInfoMap["profile-uid"]; v != nil {
						ipsProtectionsInfoMapToAdd["profile_uid"] = v
					}
					if v := ipsProtectionsInfoMap["protection-name"]; v != nil {
						ipsProtectionsInfoMapToAdd["protection_name"] = v
					}
					if v := ipsProtectionsInfoMap["status"]; v != nil {
						ipsProtectionsInfoMapToAdd["status"] = v
					}

					ipsProtectionsInfoListToReturn = append(ipsProtectionsInfoListToReturn, ipsProtectionsInfoMapToAdd)
				}

				relevantObjectsMapToReturn["ips_protections_info"] = ipsProtectionsInfoListToReturn
			}
		}

		if v := relevantObjectsMap["relevant-objects-type"]; v != nil {
			relevantObjectsMapToReturn["relevant_objects_type"] = v
		}

		_ = d.Set("relevant_objects", []interface{}{relevantObjectsMapToReturn})

	} else {
		_ = d.Set("relevant_objects", nil)
	}

	if firewallBestPractice["rule"] != nil {

		ruleMap := firewallBestPractice["rule"].(map[string]interface{})

		ruleMapToReturn := make(map[string]interface{})

		if v := ruleMap["source"]; v != nil {
			sourceList := v.([]interface{})
			sourceNames := make([]interface{}, 0, len(sourceList))
			for _, item := range sourceList {
				sourceNames = append(sourceNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["source"] = sourceNames
		}

		if v := ruleMap["negate-source"]; v != nil {
			ruleMapToReturn["negate_source"] = v
		}
		if v := ruleMap["destination"]; v != nil {
			destinationList := v.([]interface{})
			destinationNames := make([]interface{}, 0, len(destinationList))
			for _, item := range destinationList {
				destinationNames = append(destinationNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["destination"] = destinationNames
		}

		if v := ruleMap["negate-destination"]; v != nil {
			ruleMapToReturn["negate_destination"] = v
		}
		if v := ruleMap["vpn"]; v != nil {
			vpnList := v.([]interface{})
			vpnNames := make([]interface{}, 0, len(vpnList))
			for _, item := range vpnList {
				vpnNames = append(vpnNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["vpn"] = vpnNames
		}

		if v := ruleMap["negate-vpn"]; v != nil {
			ruleMapToReturn["negate_vpn"] = v
		}
		if v := ruleMap["services-and-applications"]; v != nil {
			servicesList := v.([]interface{})
			servicesNames := make([]interface{}, 0, len(servicesList))
			for _, item := range servicesList {
				servicesNames = append(servicesNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["services_and_applications"] = servicesNames
		}

		if v := ruleMap["negate-services-and-applications"]; v != nil {
			ruleMapToReturn["negate_services_and_applications"] = v
		}
		if v := ruleMap["install-on"]; v != nil {
			installOnList := v.([]interface{})
			installOnNames := make([]interface{}, 0, len(installOnList))
			for _, item := range installOnList {
				installOnNames = append(installOnNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["install_on"] = installOnNames
		}

		if v := ruleMap["negate-install-on"]; v != nil {
			ruleMapToReturn["negate_install_on"] = v
		}
		if v := ruleMap["time"]; v != nil {
			timeValList := v.([]interface{})
			timeValNames := make([]interface{}, 0, len(timeValList))
			for _, item := range timeValList {
				timeValNames = append(timeValNames, item.(map[string]interface{})["name"])
			}
			ruleMapToReturn["time"] = timeValNames
		}

		if v := ruleMap["negate-time"]; v != nil {
			ruleMapToReturn["negate_time"] = v
		}
		if v := ruleMap["action"]; v != nil {
			ruleMapToReturn["action"] = v
		}
		if v := ruleMap["negate-action"]; v != nil {
			ruleMapToReturn["negate_action"] = v
		}
		if v := ruleMap["track"]; v != nil {
			ruleMapToReturn["track"] = v
		}
		if v := ruleMap["negate-track"]; v != nil {
			ruleMapToReturn["negate_track"] = v
		}
		if v := ruleMap["hit-count"]; v != nil {
			ruleMapToReturn["hit_count"] = v
		}
		if v := ruleMap["negate-hit-count"]; v != nil {
			ruleMapToReturn["negate_hit_count"] = v
		}
		if v := ruleMap["name-condition"]; v != nil {

			nameConditionMap := v.(map[string]interface{})

			nameConditionMapToReturn := make(map[string]interface{})

			if v := nameConditionMap["condition-type"]; v != nil {
				nameConditionMapToReturn["condition_type"] = v
			}
			if v := nameConditionMap["value"]; v != nil {
				nameConditionMapToReturn["value"] = v
			}

			ruleMapToReturn["name_condition"] = []interface{}{nameConditionMapToReturn}
		}

		if v := ruleMap["comment-condition"]; v != nil {

			commentConditionMap := v.(map[string]interface{})

			commentConditionMapToReturn := make(map[string]interface{})

			if v := commentConditionMap["condition-type"]; v != nil {
				commentConditionMapToReturn["condition_type"] = v
			}
			if v := commentConditionMap["value"]; v != nil {
				commentConditionMapToReturn["value"] = v
			}

			ruleMapToReturn["comment_condition"] = []interface{}{commentConditionMapToReturn}
		}

		_ = d.Set("rule", []interface{}{ruleMapToReturn})

	} else {
		_ = d.Set("rule", nil)
	}

	if v := firewallBestPractice["secure-condition"]; v != nil {
		_ = d.Set("secure_condition", v)
	}

	if v := firewallBestPractice["status"]; v != nil {
		_ = d.Set("status", v)
	}

	if v := firewallBestPractice["tolerance"]; v != nil {
		_ = d.Set("tolerance", v)
	}

	if v := firewallBestPractice["violation-condition"]; v != nil {
		_ = d.Set("violation_condition", v)
	}

	if v := firewallBestPractice["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
