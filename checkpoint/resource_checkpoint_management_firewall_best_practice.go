package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementFirewallBestPractice() *schema.Resource {
	return &schema.Resource{
		Create: createManagementFirewallBestPractice,
		Read:   readManagementFirewallBestPractice,
		Update: updateManagementFirewallBestPractice,
		Delete: deleteManagementFirewallBestPractice,
		Schema: map[string]*schema.Schema{
			"best_practice_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Best Practice ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Best Practice Name.",
			},
			"action_item": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "To comply with Best Practice, do this action item.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the Best Practice.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The activation status of the best practice.",
				Default:     true,
			},
			"expiration": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Deactivation expiration settings.<br><font color=\"red\">Required only if</font> enabled is set to false.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The reason for deactivating the best practice.",
						},
						"expire_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When the deactivation expires. Date and time represented in international ISO 8601 format. Relevant only if mode is set to 'expire-on'.",
							Default:     "current date and time",
						},
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether the deactivation never expires or expires on a specific date.",
							Default:     "never",
						},
					},
				},
			},
			"policy_range_percentage": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The percentage of the Rule Base to scan (0-100).",
				Default:     100,
			},
			"policy_range_position": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The direction of the scan.",
				Default:     "top",
			},
			"poor_condition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Visibility of poor-result rules in the Relevant Objects pane.",
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The rule criteria the firewall best practice evaluates against the rule base.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Network objects to match in the rule Source column. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_source": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the source values are negated.",
							Default:     false,
						},
						"destination": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Network objects to match in the rule Destination column. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_destination": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the destination values are negated.",
							Default:     false,
						},
						"vpn": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "VPN communities to match. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_vpn": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the vpn values are negated.",
							Default:     false,
						},
						"services_and_applications": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Services, applications, categories or sites to match. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_services_and_applications": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the services and applications values are negated.",
							Default:     false,
						},
						"install_on": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Security Gateways or Clusters the rule applies to. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_install_on": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the install-on values are negated.",
							Default:     false,
						},
						"time": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Time objects the rule applies to. Identified by name or UID.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_time": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the time values are negated.",
							Default:     false,
						},
						"action": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Rule actions to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_action": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the action values are negated.",
							Default:     false,
						},
						"track": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Tracking methods to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_track": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the track values are negated.",
							Default:     false,
						},
						"hit_count": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Hit-count levels to match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"negate_hit_count": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Shows if the hit-count values are negated.",
							Default:     false,
						},
						"name_condition": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Match the rule name against a text condition.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The condition type.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The condition match string. Relevant only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.",
									},
								},
							},
						},
						"comment_condition": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Match the rule comment against a text condition.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The condition type.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The condition match string. Relevant only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.",
									},
								},
							},
						},
					},
				},
			},
			"secure_condition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Visibility of secure-result rules in the Relevant Objects pane.",
			},
			"tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of matches allowed before a violation is created. Valid values: between 0 and 100.<br><font color=\"red\">Required only if</font> violation-condition is set to 'Rule found'.",
				Default:     0,
			},
			"violation_condition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Define when a violation occurs: 'Rule found' means the criteria match a rule; 'Rule not found' means no rule matches.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementFirewallBestPractice(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	firewallBestPractice := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		firewallBestPractice["name"] = v.(string)
	}

	if v, ok := d.GetOk("action_item"); ok {
		firewallBestPractice["action-item"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		firewallBestPractice["description"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		firewallBestPractice["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("expiration"); ok {

		expirationList := v.([]interface{})

		if len(expirationList) > 0 {

			expirationPayload := make(map[string]interface{})

			if v, ok := d.GetOk("expiration.0.comment"); ok {
				expirationPayload["comment"] = v.(string)
			}
			if v, ok := d.GetOk("expiration.0.expire_on"); ok {
				expirationPayload["expire-on"] = v.(string)
			}
			if v, ok := d.GetOk("expiration.0.mode"); ok {
				expirationPayload["mode"] = v.(string)
			}
			firewallBestPractice["expiration"] = expirationPayload
		}
	}
	if v, ok := d.GetOk("policy_range_percentage"); ok {
		firewallBestPractice["policy-range-percentage"] = v.(int)
	}

	if v, ok := d.GetOk("policy_range_position"); ok {
		firewallBestPractice["policy-range-position"] = v.(string)
	}

	if v, ok := d.GetOk("poor_condition"); ok {
		firewallBestPractice["poor-condition"] = v.(string)
	}

	if v, ok := d.GetOk("rule"); ok {

		ruleList := v.([]interface{})

		if len(ruleList) > 0 {

			rulePayload := make(map[string]interface{})

			if v, ok := d.GetOk("rule.0.source"); ok {
				rulePayload["source"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_source"); ok {
				rulePayload["negate-source"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.destination"); ok {
				rulePayload["destination"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_destination"); ok {
				rulePayload["negate-destination"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.vpn"); ok {
				rulePayload["vpn"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_vpn"); ok {
				rulePayload["negate-vpn"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.services_and_applications"); ok {
				rulePayload["services-and-applications"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_services_and_applications"); ok {
				rulePayload["negate-services-and-applications"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.install_on"); ok {
				rulePayload["install-on"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_install_on"); ok {
				rulePayload["negate-install-on"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.time"); ok {
				rulePayload["time"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_time"); ok {
				rulePayload["negate-time"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.action"); ok {
				rulePayload["action"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_action"); ok {
				rulePayload["negate-action"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.track"); ok {
				rulePayload["track"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_track"); ok {
				rulePayload["negate-track"] = v.(bool)
			}
			if v, ok := d.GetOk("rule.0.hit_count"); ok {
				rulePayload["hit-count"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOkExists("rule.0.negate_hit_count"); ok {
				rulePayload["negate-hit-count"] = v.(bool)
			}
			if _, ok := d.GetOk("rule.0.name_condition"); ok {

				nameConditionPayload := make(map[string]interface{})

				if v, ok := d.GetOk("rule.0.name_condition.0.condition_type"); ok {
					nameConditionPayload["condition-type"] = v.(string)
				}
				if v, ok := d.GetOk("rule.0.name_condition.0.value"); ok {
					nameConditionPayload["value"] = v.(string)
				}
				rulePayload["name-condition"] = nameConditionPayload
			}
			if _, ok := d.GetOk("rule.0.comment_condition"); ok {

				commentConditionPayload := make(map[string]interface{})

				if v, ok := d.GetOk("rule.0.comment_condition.0.condition_type"); ok {
					commentConditionPayload["condition-type"] = v.(string)
				}
				if v, ok := d.GetOk("rule.0.comment_condition.0.value"); ok {
					commentConditionPayload["value"] = v.(string)
				}
				rulePayload["comment-condition"] = commentConditionPayload
			}
			firewallBestPractice["rule"] = rulePayload
		}
	}
	if v, ok := d.GetOk("secure_condition"); ok {
		firewallBestPractice["secure-condition"] = v.(string)
	}

	if v, ok := d.GetOk("tolerance"); ok {
		firewallBestPractice["tolerance"] = v.(int)
	}

	if v, ok := d.GetOk("violation_condition"); ok {
		firewallBestPractice["violation-condition"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		firewallBestPractice["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		firewallBestPractice["ignore-errors"] = v.(bool)
	}

	log.Println("Create FirewallBestPractice - Map = ", firewallBestPractice)

	addFirewallBestPracticeRes, err := client.ApiCall("add-firewall-best-practice", firewallBestPractice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addFirewallBestPracticeRes.Success {
		if addFirewallBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(addFirewallBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addFirewallBestPracticeRes.GetData()["uid"].(string))

	return readManagementFirewallBestPractice(d, m)
}

func readManagementFirewallBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showFirewallBestPracticeRes, err := client.ApiCall("show-firewall-best-practice", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := firewallBestPractice["best-practice-id"]; v != nil {
		_ = d.Set("best_practice_id", v)
	}

	if v := firewallBestPractice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := firewallBestPractice["action-item"]; v != nil {
		_ = d.Set("action_item", v)
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
			expirationMapToReturn["expire_on"] = v
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

			if v, _ := nameConditionMap["condition-type"]; v != nil {
				nameConditionMapToReturn["condition_type"] = v
			}
			if v, _ := nameConditionMap["value"]; v != nil {
				nameConditionMapToReturn["value"] = v
			}
			ruleMapToReturn["name_condition"] = []interface{}{nameConditionMapToReturn}
		}

		if v := ruleMap["comment-condition"]; v != nil {

			commentConditionMap := v.(map[string]interface{})
			commentConditionMapToReturn := make(map[string]interface{})

			if v, _ := commentConditionMap["condition-type"]; v != nil {
				commentConditionMapToReturn["condition_type"] = v
			}
			if v, _ := commentConditionMap["value"]; v != nil {
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

	if v := firewallBestPractice["tolerance"]; v != nil {
		_ = d.Set("tolerance", v)
	}

	if v := firewallBestPractice["violation-condition"]; v != nil {
		_ = d.Set("violation_condition", v)
	}

	if v := firewallBestPractice["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := firewallBestPractice["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementFirewallBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	firewallBestPractice := make(map[string]interface{})

	firewallBestPractice["uid"] = d.Id()

	if ok := d.HasChange("best_practice_id"); ok {
		firewallBestPractice["best-practice-id"] = d.Get("best_practice_id")
	}

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			firewallBestPractice["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("action_item"); ok {
		firewallBestPractice["action-item"] = d.Get("action_item")
	}

	if ok := d.HasChange("description"); ok {
		firewallBestPractice["description"] = d.Get("description")
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		firewallBestPractice["enabled"] = v.(bool)
	}

	if d.HasChange("expiration") {

		if v, ok := d.GetOk("expiration"); ok {

			expirationList := v.([]interface{})

			if len(expirationList) > 0 {

				expirationPayload := make(map[string]interface{})

				if v, ok := d.GetOk("expiration.0.comment"); ok {
					expirationPayload["comment"] = v.(string)
				}
				if v, ok := d.GetOk("expiration.0.expire_on"); ok {
					expirationPayload["expire-on"] = v.(string)
				}
				if v, ok := d.GetOk("expiration.0.mode"); ok {
					expirationPayload["mode"] = v.(string)
				}
				firewallBestPractice["expiration"] = expirationPayload
			}
		}
	}

	if ok := d.HasChange("policy_range_percentage"); ok {
		firewallBestPractice["policy-range-percentage"] = d.Get("policy_range_percentage")
	}

	if ok := d.HasChange("policy_range_position"); ok {
		firewallBestPractice["policy-range-position"] = d.Get("policy_range_position")
	}

	if ok := d.HasChange("poor_condition"); ok {
		firewallBestPractice["poor-condition"] = d.Get("poor_condition")
	}

	if d.HasChange("rule") {

		if v, ok := d.GetOk("rule"); ok {

			ruleList := v.([]interface{})

			if len(ruleList) > 0 {

				rulePayload := make(map[string]interface{})

				if v, ok := d.GetOk("rule.0.source"); ok {
					rulePayload["source"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_source"); ok {
					rulePayload["negate-source"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.destination"); ok {
					rulePayload["destination"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_destination"); ok {
					rulePayload["negate-destination"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.vpn"); ok {
					rulePayload["vpn"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_vpn"); ok {
					rulePayload["negate-vpn"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.services_and_applications"); ok {
					rulePayload["services-and-applications"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_services_and_applications"); ok {
					rulePayload["negate-services-and-applications"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.install_on"); ok {
					rulePayload["install-on"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_install_on"); ok {
					rulePayload["negate-install-on"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.time"); ok {
					rulePayload["time"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_time"); ok {
					rulePayload["negate-time"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.action"); ok {
					rulePayload["action"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_action"); ok {
					rulePayload["negate-action"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.track"); ok {
					rulePayload["track"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_track"); ok {
					rulePayload["negate-track"] = v.(bool)
				}
				if v, ok := d.GetOk("rule.0.hit_count"); ok {
					rulePayload["hit-count"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("rule.0.negate_hit_count"); ok {
					rulePayload["negate-hit-count"] = v.(bool)
				}
				if _, ok := d.GetOk("rule.0.name_condition"); ok {

					nameConditionPayload := make(map[string]interface{})

					if v, ok := d.GetOk("rule.0.name_condition.0.condition_type"); ok {
						nameConditionPayload["condition-type"] = v.(string)
					}
					if v, ok := d.GetOk("rule.0.name_condition.0.value"); ok {
						nameConditionPayload["value"] = v.(string)
					}
					rulePayload["name-condition"] = nameConditionPayload
				}
				if _, ok := d.GetOk("rule.0.comment_condition"); ok {

					commentConditionPayload := make(map[string]interface{})

					if v, ok := d.GetOk("rule.0.comment_condition.0.condition_type"); ok {
						commentConditionPayload["condition-type"] = v.(string)
					}
					if v, ok := d.GetOk("rule.0.comment_condition.0.value"); ok {
						commentConditionPayload["value"] = v.(string)
					}
					rulePayload["comment-condition"] = commentConditionPayload
				}
				firewallBestPractice["rule"] = rulePayload
			}
		}
	}

	if ok := d.HasChange("secure_condition"); ok {
		firewallBestPractice["secure-condition"] = d.Get("secure_condition")
	}

	if ok := d.HasChange("tolerance"); ok {
		firewallBestPractice["tolerance"] = d.Get("tolerance")
	}

	if ok := d.HasChange("violation_condition"); ok {
		firewallBestPractice["violation-condition"] = d.Get("violation_condition")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		firewallBestPractice["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		firewallBestPractice["ignore-errors"] = v.(bool)
	}

	log.Println("Update FirewallBestPractice - Map = ", firewallBestPractice)

	updateFirewallBestPracticeRes, err := client.ApiCall("set-firewall-best-practice", firewallBestPractice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateFirewallBestPracticeRes.Success {
		if updateFirewallBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(updateFirewallBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementFirewallBestPractice(d, m)
}

func deleteManagementFirewallBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	firewallBestPracticePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete FirewallBestPractice")

	deleteFirewallBestPracticeRes, err := client.ApiCall("delete-firewall-best-practice", firewallBestPracticePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteFirewallBestPracticeRes.Success {
		if deleteFirewallBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(deleteFirewallBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
