package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementThreatProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatProfile,
		Read:   readManagementThreatProfile,
		Update: updateManagementThreatProfile,
		Delete: deleteManagementThreatProfile,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"active_protections_performance_impact": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protections with this performance impact only will be activated in the profile.",
			},
			"active_protections_severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protections with this severity only will be activated in the profile.",
			},
			"confidence_level_high": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Prevent",
				Description: "Action for protections with high confidence level.",
			},
			"confidence_level_low": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Detect",
				Description: "Action for protections with low confidence level.",
			},
			"confidence_level_medium": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Prevent",
				Description: "Action for protections with medium confidence level.",
			},
			"indicator_overrides": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Indicators whose action will be overridden in this profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The indicator's action in this profile.",
						},
						"indicator": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The indicator whose action is to be overriden.",
						},
					},
				},
			},
			"ips_settings": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "IPS blade settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_protection_with_performance_impact": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to exclude protections depending on their level of performance impact.",
						},
						"exclude_protection_with_performance_impact_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Exclude protections with this level of performance impact.",
						},
						"exclude_protection_with_severity": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to exclude protections depending on their level of severity.",
						},
						"exclude_protection_with_severity_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Exclude protections with this level of severity.",
						},
						"newly_updated_protections": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Activation of newly updated protections.",
						},
					},
				},
			},
			"malicious_mail_policy_settings": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Malicious Mail Policy for MTA Gateways.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_customized_text_to_email_body": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Add customized text to the malicious email body.",
						},
						"add_email_subject_prefix": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Add a prefix to the malicious email subject.",
						},
						"add_x_header_to_email": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Add an X-Header to the malicious email.",
						},
						"email_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Block - block the entire malicious email. Allow - pass the malicious email and apply email changes (like: remove attachments and links, add x-header, etc...).",
						},
						"email_body_customized_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Customized text for the malicious email body. Available predefined fields: $verdicts$ - the malicious/error attachments/links verdict.",
						},
						"email_subject_prefix_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Prefix for the malicious email subject.",
						},
						"failed_to_scan_attachments_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Replace attachments that failed to be scanned with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.",
						},
						"malicious_attachments_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Replace malicious attachments with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.",
						},
						"malicious_links_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Replace malicious links with this text. Available predefined fields: $neutralized_url$ - neutralized malicious link.",
						},
						"remove_attachments_and_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Remove attachments and links from the malicious email.",
						},
						"send_copy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Send a copy of the malicious email to the recipient list.",
						},
						"send_copy_list": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Recipient list to send a copy of the malicious email.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"overrides": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Overrides per profile for this protection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protection": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IPS protection identified by name.",
						},
						"action": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Protection action.",
						},
						"capture_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Capture packets.",
						},
						"track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tracking method for protection.",
						},
						"default": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Default settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Protection action.",
									},
									"capture_packets": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Capture packets.",
									},
									"track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tracking method for protection.",
									},
								},
							},
						},
						"final": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Final settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Protection action.",
									},
									"capture_packets": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Capture packets.",
									},
									"track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tracking method for protection.",
									},
								},
							},
						},
						"protection_external_info": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of industry reference (CVE).",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"protection_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPS protection unique identifier.",
						},
					},
				},
			},
			"scan_malicious_links": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Scans malicious links (URLs) inside email messages.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_bytes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Scan links in the first bytes of the mail body.",
						},
						"max_links": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum links to scan in mail body.",
						},
					},
				},
			},
			"use_indicators": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates whether the profile should make use of indicators.",
			},
			"anti_bot": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Is Anti-Bot blade activated.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Is Anti-Virus blade activated.",
			},
			"ips": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Is IPS blade activated.",
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Is Threat Emulation blade activated.",
			},
			"activate_protections_by_extended_attributes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Activate protections by these extended attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag unique identifier.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag name.",
						},
						"category": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag category name.",
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "IPS protection extended attribute values",
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
								},
							},
						},
					},
				},
			},
			"deactivate_protections_by_extended_attributes": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Deactivate protections by these extended attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag unique identifier.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag name.",
						},
						"category": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPS tag category name.",
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "IPS protection extended attribute values",
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
								},
							},
						},
					},
				},
			},
			"use_extended_attributes": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to activate/deactivate IPS protections according to the extended attributes.",
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementThreatProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatProfile := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		threatProfile["name"] = v.(string)
	}

	if v, ok := d.GetOk("active_protections_performance_impact"); ok {
		threatProfile["active-protections-performance-impact"] = v.(string)
	}

	if v, ok := d.GetOk("active_protections_severity"); ok {
		threatProfile["active-protections-severity"] = v.(string)
	}

	if v, ok := d.GetOk("confidence_level_high"); ok {
		threatProfile["confidence-level-high"] = v.(string)
	}

	if v, ok := d.GetOk("confidence_level_low"); ok {
		threatProfile["confidence-level-low"] = v.(string)
	}

	if v, ok := d.GetOk("confidence_level_medium"); ok {
		threatProfile["confidence-level-medium"] = v.(string)
	}

	if v, ok := d.GetOk("indicator_overrides"); ok {
		indicatorOverridesList := v.([]interface{})
		if len(indicatorOverridesList) > 0 {
			var indicatorOverridesPayload []map[string]interface{}
			for i := range indicatorOverridesList {
				indicatorOverride := make(map[string]interface{})
				if v, ok := d.GetOk("indicator_overrides." + strconv.Itoa(i) + ".action"); ok {
					indicatorOverride["action"] = v.(string)
				}
				if v, ok := d.GetOk("indicator_overrides." + strconv.Itoa(i) + ".indicator"); ok {
					indicatorOverride["indicator"] = v.(string)
				}
				indicatorOverridesPayload = append(indicatorOverridesPayload, indicatorOverride)
			}
			threatProfile["indicator-overrides"] = indicatorOverridesPayload
		}
	}

	if _, ok := d.GetOk("ips_settings"); ok {
		ipsSettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOkExists("ips_settings.exclude_protection_with_performance_impact"); ok {
			ipsSettingsPayload["exclude-protection-with-performance-impact"] = v.(bool)
		}
		if v, ok := d.GetOk("ips_settings.exclude_protection_with_performance_impact_mode"); ok {
			ipsSettingsPayload["exclude-protection-with-performance-impact-mode"] = v.(string)
		}
		if v, ok := d.GetOkExists("ips_settings.exclude_protection_with_severity"); ok {
			ipsSettingsPayload["exclude-protection-with-severity"] = v.(bool)
		}
		if v, ok := d.GetOk("ips_settings.exclude_protection_with_severity_mode"); ok {
			ipsSettingsPayload["exclude-protection-with-severity-mode"] = v.(string)
		}
		if v, ok := d.GetOk("ips_settings.newly_updated_protections"); ok {
			ipsSettingsPayload["newly-updated-protections"] = v.(string)
		}

		threatProfile["ips-settings"] = ipsSettingsPayload
	}

	if _, ok := d.GetOk("malicious_mail_policy_settings"); ok {
		maliciousMailPolicySettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_customized_text_to_email_body"); ok {
			maliciousMailPolicySettingsPayload["add-customized-text-to-email-body"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_email_subject_prefix"); ok {
			maliciousMailPolicySettingsPayload["add-email-subject-prefix"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_x_header_to_email"); ok {
			maliciousMailPolicySettingsPayload["add-x-header-to-email"] = v.(bool)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_action"); ok {
			maliciousMailPolicySettingsPayload["email-action"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_body_customized_text"); ok {
			maliciousMailPolicySettingsPayload["email-body-customized-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_subject_prefix_text"); ok {
			maliciousMailPolicySettingsPayload["email-subject-prefix-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.failed_to_scan_attachments_text"); ok {
			maliciousMailPolicySettingsPayload["failed-to-scan-attachments-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.malicious_attachments_text"); ok {
			maliciousMailPolicySettingsPayload["malicious-attachments-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.malicious_links_text"); ok {
			maliciousMailPolicySettingsPayload["malicious-links-text"] = v.(string)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.remove_attachments_and_links"); ok {
			maliciousMailPolicySettingsPayload["remove-attachments-and-links"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.send_copy"); ok {
			maliciousMailPolicySettingsPayload["send-copy"] = v.(bool)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.send_copy_list"); ok {
			maliciousMailPolicySettingsPayload["send-copy-list"] = v.(*schema.Set).List()
		}

		threatProfile["malicious-mail-policy-settings"] = maliciousMailPolicySettingsPayload
	}

	if v, ok := d.GetOk("overrides"); ok {
		overridesList := v.([]interface{})
		if len(overridesList) > 0 {
			var overridesPayload []map[string]interface{}
			for i := range overridesList {
				override := make(map[string]interface{})
				if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".action"); ok {
					override["action"] = v.(string)
				}
				if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".protection"); ok {
					override["protection"] = v.(string)
				}
				if v, ok := d.GetOkExists("overrides." + strconv.Itoa(i) + ".capture_packets"); ok {
					override["capture-packets"] = v.(bool)
				}
				if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".track"); ok {
					override["track"] = v.(string)
				}
				overridesPayload = append(overridesPayload, override)
			}
			threatProfile["overrides"] = overridesPayload
		}
	}

	if _, ok := d.GetOk("scan_malicious_links"); ok {
		scanMaliciousLinksPayload := make(map[string]interface{})
		if v, ok := d.GetOk("scan_malicious_links.max_bytes"); ok {
			scanMaliciousLinksPayload["max-bytes"] = v.(int)
		}
		if v, ok := d.GetOk("profile_overrides.max_links"); ok {
			scanMaliciousLinksPayload["max-links"] = v.(int)
		}
		threatProfile["scan-malicious-links"] = scanMaliciousLinksPayload
	}

	if v, ok := d.GetOkExists("use_indicators"); ok {
		threatProfile["use-indicators"] = v.(bool)
	}

	if v, ok := d.GetOkExists("anti_bot"); ok {
		threatProfile["anti-bot"] = v.(bool)
	}

	if v, ok := d.GetOkExists("anti_virus"); ok {
		threatProfile["anti-virus"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ips"); ok {
		threatProfile["ips"] = v.(bool)
	}

	if v, ok := d.GetOkExists("threat_emulation"); ok {
		threatProfile["threat-emulation"] = v.(bool)
	}

	if v, ok := d.GetOkExists("use_extended_attributes"); ok {
		threatProfile["use-extended-attributes"] = v.(bool)
	}

	if v, ok := d.GetOk("activate_protections_by_extended_attributes"); ok {
		activateProtectionsByExtendedAttributesList := v.([]interface{})
		if len(activateProtectionsByExtendedAttributesList) > 0 {
			var activateProtectionsByExtendedAttributesPayload []map[string]interface{}
			for i := range activateProtectionsByExtendedAttributesList {
				activateProtectionsByExtendedAttributes := make(map[string]interface{})
				if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".uid"); ok {
					activateProtectionsByExtendedAttributes["uid"] = v.(string)
				}
				if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".name"); ok {
					activateProtectionsByExtendedAttributes["name"] = v.(string)
				}
				if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".category"); ok {
					activateProtectionsByExtendedAttributes["category"] = v.(string)
				}
				activateProtectionsByExtendedAttributesPayload = append(activateProtectionsByExtendedAttributesPayload, activateProtectionsByExtendedAttributes)
			}
			threatProfile["activate-protections-by-extended-attributes"] = activateProtectionsByExtendedAttributesPayload
		}
	}

	if v, ok := d.GetOk("deactivate_protections_by_extended_attributes"); ok {
		deactivateProtectionsByExtendedAttributesList := v.([]interface{})
		if len(deactivateProtectionsByExtendedAttributesList) > 0 {
			var deactivateProtectionsByExtendedAttributesPayload []map[string]interface{}
			for i := range deactivateProtectionsByExtendedAttributesList {
				deactivateProtectionsByExtendedAttributes := make(map[string]interface{})
				if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".uid"); ok {
					deactivateProtectionsByExtendedAttributes["uid"] = v.(string)
				}
				if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".name"); ok {
					deactivateProtectionsByExtendedAttributes["name"] = v.(string)
				}
				if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".category"); ok {
					deactivateProtectionsByExtendedAttributes["category"] = v.(string)
				}
				deactivateProtectionsByExtendedAttributesPayload = append(deactivateProtectionsByExtendedAttributesPayload, deactivateProtectionsByExtendedAttributes)
			}
			threatProfile["deactivate-protections-by-extended-attributes"] = deactivateProtectionsByExtendedAttributesPayload
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		threatProfile["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		threatProfile["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		threatProfile["color"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatProfile["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatProfile["ignore-warnings"] = v.(bool)
	}

	log.Println("Create Threat Profile - Map = ", threatProfile)

	threatProfileRes, err := client.ApiCall("add-threat-profile", threatProfile, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !threatProfileRes.Success {
		if threatProfileRes.ErrorMsg != "" {
			return fmt.Errorf(threatProfileRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-threat-profile", threatProfileRes.GetData())
		return fmt.Errorf(msg)
	}

	showThreatProfileRes, err := client.ApiCall("show-threat-profile", map[string]interface{}{"name": d.Get("name")}, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatProfileRes.Success {
		return fmt.Errorf(showThreatProfileRes.ErrorMsg)
	}

	d.SetId(showThreatProfileRes.GetData()["uid"].(string))

	return readManagementThreatProfile(d, m)
}

func readManagementThreatProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showThreatProfileRes, err := client.ApiCall("show-threat-profile", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatProfileRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showThreatProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatProfileRes.ErrorMsg)
	}

	threatProfile := showThreatProfileRes.GetData()

	log.Println("Read Threat Profile - Show JSON = ", threatProfile)

	if v := threatProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatProfile["active-protections-performance-impact"]; v != nil {
		_ = d.Set("active_protections_performance_impact", v)
	}
	if v := threatProfile["active-protections-severity"]; v != nil {
		_ = d.Set("active_protections_severity", v)
	}
	if v := threatProfile["confidence-level-high"]; v != nil {
		_ = d.Set("confidence_level_high", v)
	}
	if v := threatProfile["confidence-level-low"]; v != nil {
		_ = d.Set("confidence_level_low", v)
	}
	if v := threatProfile["confidence-level-medium"]; v != nil {
		_ = d.Set("confidence_level_medium", v)
	}

	if threatProfile["indicator-overrides"] != nil {
		indicatorOverridesList := threatProfile["indicator-overrides"].([]interface{})
		if len(indicatorOverridesList) > 0 {
			var indicatorOverridesListState []map[string]interface{}
			for i := range indicatorOverridesList {
				indicatorOverridesJson := indicatorOverridesList[i].(map[string]interface{})
				indicatorOverride := make(map[string]interface{})
				if v, _ := indicatorOverridesJson["action"]; v != nil {
					indicatorOverride["action"] = v
				}
				if v, _ := indicatorOverridesJson["indicator"]; v != nil {
					indicatorOverride["indicator"] = v
				}
				indicatorOverridesListState = append(indicatorOverridesListState, indicatorOverride)
			}
			_ = d.Set("indicator_overrides", indicatorOverridesListState)
		} else {
			_ = d.Set("indicator_overrides", indicatorOverridesList)
		}
	} else {
		_ = d.Set("indicator_overrides", nil)
	}

	if v := threatProfile["ips-settings"]; v != nil {
		ipsSettingsJson := threatProfile["ips-settings"].(map[string]interface{})
		ipsSettingsState := make(map[string]interface{})
		if v := ipsSettingsJson["exclude-protection-with-performance-impact"]; v != nil {
			ipsSettingsState["exclude_protection_with_performance_impact"] = v
		}
		if v := ipsSettingsJson["exclude-protection-with-performance-impact-mode"]; v != nil {
			ipsSettingsState["exclude_protection_with_performance_impact_mode"] = v
		}
		if v := ipsSettingsJson["exclude-protection-with-severity"]; v != nil {
			ipsSettingsState["exclude_protection_with_severity"] = v
		}
		if v := ipsSettingsJson["exclude-protection-with-severity-mode"]; v != nil {
			ipsSettingsState["exclude_protection_with_severity_mode"] = v
		}
		if v := ipsSettingsJson["newly-updated-protections"]; v != nil {
			ipsSettingsState["newly_updated_protections"] = v
		}
		_, ipsSettingsInConf := d.GetOk("ips_settings")
		defaultIpsSettings := map[string]interface{}{
			"newly_updated_protections":                  "active",
			"exclude_protection_with_performance_impact": false,
			"exclude_protection_with_severity":           false,
		}
		if reflect.DeepEqual(defaultIpsSettings, ipsSettingsState) && !ipsSettingsInConf {
			_ = d.Set("ips_settings", map[string]interface{}{})
		} else {
			_ = d.Set("ips_settings", ipsSettingsState)
		}
	} else {
		_ = d.Set("ips_settings", nil)
	}

	if v := threatProfile["malicious-mail-policy-settings"]; v != nil {
		maliciousMailPolicySettingsJson := threatProfile["malicious-mail-policy-settings"].(map[string]interface{})
		maliciousMailPolicySettingsState := make(map[string]interface{})
		if v := maliciousMailPolicySettingsJson["add-customized-text-to-email-body"]; v != nil {
			maliciousMailPolicySettingsState["add_customized_text_to_email_body"] = v
		}
		if v := maliciousMailPolicySettingsJson["add-email-subject-prefix"]; v != nil {
			maliciousMailPolicySettingsState["add_email_subject_prefix"] = v
		}
		if v := maliciousMailPolicySettingsJson["add-x-header-to-email"]; v != nil {
			maliciousMailPolicySettingsState["add_x_header_to_email"] = v
		}
		if v := maliciousMailPolicySettingsJson["email-action"]; v != nil {
			maliciousMailPolicySettingsState["email_action"] = v
		}
		if v := maliciousMailPolicySettingsJson["email-body-customized-text"]; v != nil {
			maliciousMailPolicySettingsState["email_body_customized_text"] = v
		}
		if v := maliciousMailPolicySettingsJson["email-subject-prefix-text"]; v != nil {
			maliciousMailPolicySettingsState["email_subject_prefix_text"] = v
		}
		if v := maliciousMailPolicySettingsJson["failed-to-scan-attachments-text"]; v != nil {
			maliciousMailPolicySettingsState["failed_to_scan_attachments_text"] = v
		}
		if v := maliciousMailPolicySettingsJson["malicious-attachments-text"]; v != nil {
			maliciousMailPolicySettingsState["malicious_attachments_text"] = v
		}
		if v := maliciousMailPolicySettingsJson["malicious-links-text"]; v != nil {
			maliciousMailPolicySettingsState["malicious_links_text"] = v
		}
		if v := maliciousMailPolicySettingsJson["remove-attachments-and-links"]; v != nil {
			maliciousMailPolicySettingsState["remove_attachments_and_links"] = v
		}
		if v := maliciousMailPolicySettingsJson["send-copy"]; v != nil {
			maliciousMailPolicySettingsState["send_copy"] = v
		}
		if v := maliciousMailPolicySettingsJson["send-copy-list"]; v != nil {
			maliciousMailPolicySettingsState["send_copy_list"] = v
		}

		_, maliciousMailPolicySettingsInConf := d.GetOk("malicious_mail_policy_settings")
		defaultMaliciousMailPolicySettings := map[string]interface{}{
			"email_action":                      "allow",
			"remove_attachments_and_links":      true,
			"malicious_attachments_text":        "Malicious email attachment '$filename$' removed by Check Point.",
			"failed_to_scan_attachments_text":   "Email attachment '$filename$' failed to be scanned and removed by Check Point.",
			"malicious_links_text":              "[Check Point] Malicious link: $neutralized_url$ [Check Point]",
			"add_x_header_to_email":             false,
			"add_email_subject_prefix":          false,
			"email_subject_prefix_text":         "Attachment was found malicious. It is recommended not to open this mail.",
			"add_customized_text_to_email_body": false,
			"email_body_customized_text":        "[Check Point]<BR>The following verdicts were determined by Check Point:<BR>$verdicts$<BR>[Check Point]",
			"send_copy":                         false,
		}
		if reflect.DeepEqual(defaultMaliciousMailPolicySettings, maliciousMailPolicySettingsState) && !maliciousMailPolicySettingsInConf {
			_ = d.Set("malicious_mail_policy_settings", map[string]interface{}{})
		} else {
			_ = d.Set("malicious_mail_policy_settings", maliciousMailPolicySettingsState)
		}
	} else {
		_ = d.Set("malicious_mail_policy_settings", nil)
	}

	if threatProfile["overrides"] != nil {
		overridesList := threatProfile["overrides"].([]interface{})
		if len(overridesList) > 0 {
			var overridesListState []map[string]interface{}
			for i := range overridesList {
				overridesJson := overridesList[i].(map[string]interface{})
				overrideState := make(map[string]interface{})
				if v, _ := overridesJson["protection"]; v != nil {
					overrideState["protection"] = v
				}
				if v, _ := overridesJson["override"]; v != nil {
					overrideObject := v.(map[string]interface{})
					if v, _ = overrideObject["action"]; v != nil {
						overrideState["action"] = v
					}
					if v, _ = overrideObject["capture-packets"]; v != nil {
						overrideState["capture_packets"] = v
					}
					if v, _ = overrideObject["track"]; v != nil {
						overrideState["track"] = v
					}
				}
				if v, _ := overridesJson["protection-external-info"]; v != nil {
					overrideState["protection_external_info"] = v
				}
				if v, _ := overridesJson["protection-uid"]; v != nil {
					overrideState["protection_uid"] = v
				}

				if v, _ := overridesJson["default"]; v != nil {
					defaultJson := v.(map[string]interface{})
					defaultState := make(map[string]interface{})
					if v, _ = defaultJson["action"]; v != nil {
						defaultState["action"] = v
					}
					if v, _ = defaultJson["capture-packets"]; v != nil {
						defaultState["capture_packets"] = v
					}
					if v, _ = defaultJson["track"]; v != nil {
						defaultState["track"] = v
					}
					overrideState["default"] = defaultState
				}

				if v, _ := overridesJson["final"]; v != nil {
					finalJson := v.(map[string]interface{})
					finalState := make(map[string]interface{})
					if v, _ = finalJson["action"]; v != nil {
						finalState["action"] = v
					}
					if v, _ = finalJson["capture-packets"]; v != nil {
						finalState["capture_packets"] = v
					}
					if v, _ = finalJson["track"]; v != nil {
						finalState["track"] = v
					}
					overrideState["final"] = finalState
				}
				overridesListState = append(overridesListState, overrideState)
			}
			_ = d.Set("overrides", overridesListState)
		} else {
			_ = d.Set("overrides", overridesList)
		}
	} else {
		_ = d.Set("overrides", nil)
	}

	if v := threatProfile["scan-malicious-links"]; v != nil {
		scanMaliciousLinksJson := threatProfile["scan-malicious-links"].(map[string]interface{})
		scanMaliciousLinksState := make(map[string]interface{})
		if v := scanMaliciousLinksJson["max-bytes"]; v != nil {
			scanMaliciousLinksState["max_bytes"] = v
		}
		if v := scanMaliciousLinksJson["max-links"]; v != nil {
			scanMaliciousLinksState["max_links"] = v
		}
		_ = d.Set("scan_malicious_links", scanMaliciousLinksState)
	} else {
		_ = d.Set("scan_malicious_links", nil)
	}

	if v := threatProfile["extended-attributes-to-activate"]; v != nil {
		extendedAttributesToActivateList := threatProfile["extended-attributes-to-activate"].([]interface{})
		if len(extendedAttributesToActivateList) > 0 {
			var extendedAttributesToActivateState []interface{}
			for i := range extendedAttributesToActivateList {
				extendedAttributesToActivateJson := extendedAttributesToActivateList[i].(map[string]interface{})
				extendedAttributesToActivate := make(map[string]interface{})
				if v := extendedAttributesToActivateJson["uid"]; v != nil {
					extendedAttributesToActivate["uid"] = v
				}
				if v := extendedAttributesToActivateJson["name"]; v != nil {
					extendedAttributesToActivate["name"] = v
				}
				if v := extendedAttributesToActivateJson["values"]; v != nil {
					extendedAttributesToActivate["values"] = v
				}
				extendedAttributesToActivateState = append(extendedAttributesToActivateState, extendedAttributesToActivate)
			}
			_ = d.Set("activate_protections_by_extended_attributes", extendedAttributesToActivateState)
		} else {
			_ = d.Set("activate_protections_by_extended_attributes", extendedAttributesToActivateList)
		}
	} else {
		_ = d.Set("activate_protections_by_extended_attributes", nil)
	}

	if v := threatProfile["extended-attributes-to-deactivate"]; v != nil {
		extendedAttributesToDeactivateList := threatProfile["extended-attributes-to-deactivate"].([]interface{})
		if len(extendedAttributesToDeactivateList) > 0 {
			var extendedAttributesToDeactivateState []interface{}
			for i := range extendedAttributesToDeactivateList {
				extendedAttributesToActivateJson := extendedAttributesToDeactivateList[i].(map[string]interface{})
				extendedAttributesToDeactivate := make(map[string]interface{})
				if v := extendedAttributesToActivateJson["uid"]; v != nil {
					extendedAttributesToDeactivate["uid"] = v
				}
				if v := extendedAttributesToActivateJson["name"]; v != nil {
					extendedAttributesToDeactivate["name"] = v
				}
				if v := extendedAttributesToActivateJson["values"]; v != nil {
					extendedAttributesToDeactivate["values"] = v
				}
				extendedAttributesToDeactivateState = append(extendedAttributesToDeactivateState, extendedAttributesToDeactivate)
			}
			_ = d.Set("deactivate_protections_by_extended_attributes", extendedAttributesToDeactivateState)
		} else {
			_ = d.Set("deactivate_protections_by_extended_attributes", extendedAttributesToDeactivateList)
		}
	} else {
		_ = d.Set("deactivate_protections_by_extended_attributes", nil)
	}

	if v := threatProfile["use-indicators"]; v != nil {
		_ = d.Set("use_indicators", v)
	}

	if v := threatProfile["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := threatProfile["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := threatProfile["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if v := threatProfile["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := threatProfile["use-extended-attributes"]; v != nil {
		_ = d.Set("use_extended_attributes", v)
	}

	if v := threatProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := threatProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if threatProfile["tags"] != nil {
		tagsJson := threatProfile["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}

func updateManagementThreatProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatProfile := make(map[string]interface{})

	threatProfile["uid"] = d.Id()

	if d.HasChange("name") {
		threatProfile["new-name"] = d.Get("name")
	}

	if ok := d.HasChange("active_protections_performance_impact"); ok {
		threatProfile["active-protections-performance-impact"] = d.Get("active_protections_performance_impact")
	}

	if ok := d.HasChange("active_protections_severity"); ok {
		threatProfile["active-protections-severity"] = d.Get("active_protections_severity")
	}

	if ok := d.HasChange("confidence_level_high"); ok {
		threatProfile["confidence-level-high"] = d.Get("confidence_level_high")
	}

	if ok := d.HasChange("confidence_level_low"); ok {
		threatProfile["confidence-level-low"] = d.Get("confidence_level_low")
	}

	if ok := d.HasChange("confidence_level_medium"); ok {
		threatProfile["confidence-level-medium"] = d.Get("confidence_level_medium")
	}

	if ok := d.HasChange("indicator_overrides"); ok {
		if v, ok := d.GetOk("indicator_overrides"); ok {
			indicatorOverridesList := v.([]interface{})
			if len(indicatorOverridesList) > 0 {
				var indicatorOverridesPayload []map[string]interface{}
				for i := range indicatorOverridesList {
					indicatorOverride := make(map[string]interface{})
					if v, ok := d.GetOk("indicator_overrides." + strconv.Itoa(i) + ".action"); ok {
						indicatorOverride["action"] = v.(string)
					}
					if v, ok := d.GetOk("indicator_overrides." + strconv.Itoa(i) + ".indicator"); ok {
						indicatorOverride["indicator"] = v.(string)
					}
					indicatorOverridesPayload = append(indicatorOverridesPayload, indicatorOverride)
				}
				threatProfile["indicator-overrides"] = indicatorOverridesPayload
			}
		} else {
			oldVal, _ := d.GetChange("indicator_overrides")
			indicatorOverridesList := oldVal.([]interface{})
			if len(indicatorOverridesList) > 0 {
				var indicatorOverridesPayload []interface{}
				for i := range indicatorOverridesList {
					indicatorOverridesPayload = append(indicatorOverridesPayload, indicatorOverridesList[i].(map[string]interface{})["indicator"])
				}
				threatProfile["indicator-overrides"] = map[string]interface{}{"remove": indicatorOverridesPayload}
			}
		}
	}

	if ok := d.HasChange("ips_settings"); ok {
		ipsSettingsPayload := make(map[string]interface{})
		if v, ok := d.GetOkExists("ips_settings.exclude_protection_with_performance_impact"); ok {
			ipsSettingsPayload["exclude-protection-with-performance-impact"] = v.(bool)
		}
		if v, ok := d.GetOk("ips_settings.exclude_protection_with_performance_impact_mode"); ok {
			ipsSettingsPayload["exclude-protection-with-performance-impact-mode"] = v.(string)
		}
		if v, ok := d.GetOkExists("ips_settings.exclude_protection_with_severity"); ok {
			ipsSettingsPayload["exclude-protection-with-severity"] = v.(bool)
		}
		if v, ok := d.GetOk("ips_settings.exclude_protection_with_severity_mode"); ok {
			ipsSettingsPayload["exclude-protection-with-severity-mode"] = v.(string)
		}
		if v, ok := d.GetOk("ips_settings.newly_updated_protections"); ok {
			ipsSettingsPayload["newly-updated-protections"] = v.(string)
		}
		threatProfile["ips-settings"] = ipsSettingsPayload
	}

	if ok := d.HasChange("malicious_mail_policy_settings"); ok {
		maliciousMailPolicySettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_customized_text_to_email_body"); ok {
			maliciousMailPolicySettingsPayload["add-customized-text-to-email-body"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_email_subject_prefix"); ok {
			maliciousMailPolicySettingsPayload["add-email-subject-prefix"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.add_x_header_to_email"); ok {
			maliciousMailPolicySettingsPayload["add-x-header-to-email"] = v.(bool)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_action"); ok {
			maliciousMailPolicySettingsPayload["email-action"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_body_customized_text"); ok {
			maliciousMailPolicySettingsPayload["email-body-customized-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.email_subject_prefix_text"); ok {
			maliciousMailPolicySettingsPayload["email-subject-prefix-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.failed_to_scan_attachments_text"); ok {
			maliciousMailPolicySettingsPayload["failed-to-scan-attachments-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.malicious_attachments_text"); ok {
			maliciousMailPolicySettingsPayload["malicious-attachments-text"] = v.(string)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.malicious_links_text"); ok {
			maliciousMailPolicySettingsPayload["malicious-links-text"] = v.(string)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.remove_attachments_and_links"); ok {
			maliciousMailPolicySettingsPayload["remove-attachments-and-links"] = v.(bool)
		}
		if v, ok := d.GetOkExists("malicious_mail_policy_settings.send_copy"); ok {
			maliciousMailPolicySettingsPayload["send-copy"] = v.(bool)
		}
		if v, ok := d.GetOk("malicious_mail_policy_settings.send_copy_list"); ok {
			maliciousMailPolicySettingsPayload["send-copy-list"] = v.(*schema.Set).List()
		}
		threatProfile["malicious-mail-policy-settings"] = maliciousMailPolicySettingsPayload
	}

	if ok := d.HasChange("overrides"); ok {
		if v, ok := d.GetOk("overrides"); ok {
			overridesList := v.([]interface{})
			if len(overridesList) > 0 {
				var overridesPayload []map[string]interface{}
				for i := range overridesList {
					override := make(map[string]interface{})
					if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".action"); ok {
						override["action"] = v.(string)
					}
					if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".protection"); ok {
						override["protection"] = v.(string)
					}
					if v, ok := d.GetOkExists("overrides." + strconv.Itoa(i) + ".capture_packets"); ok {
						override["capture-packets"] = v.(bool)
					}
					if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".track"); ok {
						override["track"] = v.(string)
					}
					overridesPayload = append(overridesPayload, override)
				}
				threatProfile["overrides"] = overridesPayload
			}
		} else {
			oldVal, _ := d.GetChange("overrides")
			overridesList := oldVal.([]interface{})
			if len(overridesList) > 0 {
				var overridesPayload []interface{}
				for i := range overridesList {
					overridesPayload = append(overridesPayload, overridesList[i].(map[string]interface{})["protection"])
				}
				threatProfile["overrides"] = map[string]interface{}{"remove": overridesPayload}
			}
		}
	}

	if ok := d.HasChange("scan_malicious_links"); ok {
		scanMaliciousLinksPayload := make(map[string]interface{})
		if v, ok := d.GetOk("scan_malicious_links.max_bytes"); ok {
			scanMaliciousLinksPayload["max-bytes"] = v.(int)
		}
		if v, ok := d.GetOk("profile_overrides.max_links"); ok {
			scanMaliciousLinksPayload["max-links"] = v.(int)
		}
		threatProfile["scan-malicious-links"] = scanMaliciousLinksPayload
	}

	if ok := d.HasChange("use_indicators"); ok {
		threatProfile["use-indicators"] = d.Get("use_indicators")
	}

	if ok := d.HasChange("anti_bot"); ok {
		threatProfile["anti-bot"] = d.Get("anti_bot")
	}

	if ok := d.HasChange("anti_virus"); ok {
		threatProfile["anti-virus"] = d.Get("anti_virus")
	}

	if ok := d.HasChange("ips"); ok {
		threatProfile["ips"] = d.Get("ips")
	}

	if ok := d.HasChange("threat_emulation"); ok {
		threatProfile["threat-emulation"] = d.Get("threat_emulation")
	}

	if ok := d.HasChange("use_extended_attributes"); ok {
		threatProfile["use-extended-attributes"] = d.Get("use_extended_attributes")
	}

	if ok := d.HasChange("activate_protections_by_extended_attributes"); ok {
		if v, ok := d.GetOk("activate_protections_by_extended_attributes"); ok {
			activateProtectionsByExtendedAttributesList := v.([]interface{})
			if len(activateProtectionsByExtendedAttributesList) > 0 {
				var activateProtectionsByExtendedAttributesPayload []map[string]interface{}
				for i := range activateProtectionsByExtendedAttributesList {
					activateProtectionsByExtendedAttributes := make(map[string]interface{})
					if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".uid"); ok {
						activateProtectionsByExtendedAttributes["uid"] = v.(string)
					}
					if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".name"); ok {
						activateProtectionsByExtendedAttributes["name"] = v.(string)
					}
					if v, ok := d.GetOk("activate_protections_by_extended_attributes." + strconv.Itoa(i) + ".category"); ok {
						activateProtectionsByExtendedAttributes["category"] = v.(string)
					}
					activateProtectionsByExtendedAttributesPayload = append(activateProtectionsByExtendedAttributesPayload, activateProtectionsByExtendedAttributes)
				}
				threatProfile["activate-protections-by-extended-attributes"] = activateProtectionsByExtendedAttributesPayload
			}
		} else {
			oldVal, _ := d.GetChange("activate_protections_by_extended_attributes")
			activateProtectionsByExtendedAttributesList := oldVal.([]interface{})
			if len(activateProtectionsByExtendedAttributesList) > 0 {
				var activateProtectionsByExtendedAttributesPayload []map[string]interface{}
				for i := range activateProtectionsByExtendedAttributesList {
					activateProtectionsByExtendedAttributes := make(map[string]interface{})
					activateProtectionsByExtendedAttributes["uid"] = activateProtectionsByExtendedAttributesList[i].(map[string]interface{})["uid"]
					activateProtectionsByExtendedAttributesPayload = append(activateProtectionsByExtendedAttributesPayload, activateProtectionsByExtendedAttributes)
				}
				threatProfile["activate-protections-by-extended-attributes"] = map[string]interface{}{"remove": activateProtectionsByExtendedAttributesPayload}
			}
		}
	}

	if ok := d.HasChange("deactivate_protections_by_extended_attributes"); ok {
		if v, ok := d.GetOk("deactivate_protections_by_extended_attributes"); ok {
			deactivateProtectionsByExtendedAttributesList := v.([]interface{})
			if len(deactivateProtectionsByExtendedAttributesList) > 0 {
				var deactivateProtectionsByExtendedAttributesPayload []map[string]interface{}
				for i := range deactivateProtectionsByExtendedAttributesList {
					deactivateProtectionsByExtendedAttributes := make(map[string]interface{})
					if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".uid"); ok {
						deactivateProtectionsByExtendedAttributes["uid"] = v.(string)
					}
					if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".name"); ok {
						deactivateProtectionsByExtendedAttributes["name"] = v.(string)
					}
					if v, ok := d.GetOk("deactivate_protections_by_extended_attributes." + strconv.Itoa(i) + ".category"); ok {
						deactivateProtectionsByExtendedAttributes["category"] = v.(string)
					}
					deactivateProtectionsByExtendedAttributesPayload = append(deactivateProtectionsByExtendedAttributesPayload, deactivateProtectionsByExtendedAttributes)
				}
				threatProfile["deactivate-protections-by-extended-attributes"] = deactivateProtectionsByExtendedAttributesPayload
			}
		} else {
			oldVal, _ := d.GetChange("deactivate_protections_by_extended_attributes")
			deactivateProtectionsByExtendedAttributesList := oldVal.([]interface{})
			if len(deactivateProtectionsByExtendedAttributesList) > 0 {
				var deactivateProtectionsByExtendedAttributesPayload []map[string]interface{}
				for i := range deactivateProtectionsByExtendedAttributesList {
					deactivateProtectionsByExtendedAttributes := make(map[string]interface{})
					deactivateProtectionsByExtendedAttributes["uid"] = deactivateProtectionsByExtendedAttributesList[i].(map[string]interface{})["uid"]
					deactivateProtectionsByExtendedAttributesPayload = append(deactivateProtectionsByExtendedAttributesPayload, deactivateProtectionsByExtendedAttributes)
				}
				threatProfile["deactivate-protections-by-extended-attributes"] = map[string]interface{}{"remove": deactivateProtectionsByExtendedAttributesPayload}
			}
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			threatProfile["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			threatProfile["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		threatProfile["comments"] = d.Get("comments")
	}

	if ok := d.HasChange("color"); ok {
		threatProfile["color"] = d.Get("color")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatProfile["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatProfile["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Threat Profile - Map = ", threatProfile)

	threatProfileRes, err := client.ApiCall("set-threat-profile", threatProfile, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if !threatProfileRes.Success {
		if threatProfileRes.ErrorMsg != "" {
			return fmt.Errorf(threatProfileRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-threat-profile", threatProfileRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementThreatProfile(d, m)
}

func deleteManagementThreatProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	threatProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	deleteThreatProfileRes, err := client.ApiCall("delete-threat-profile", threatProfilePayload, client.GetSessionID(), true, false)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if !deleteThreatProfileRes.Success {
		if deleteThreatProfileRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatProfileRes.ErrorMsg)
		}
		msg := createTaskFailMessage("delete-threat-profile", deleteThreatProfileRes.GetData())
		return fmt.Errorf(msg)
	}

	d.SetId("")
	return nil
}
