package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementThreatProfileV0 is the V0 schema where overrides[*].default and overrides[*].final were TypeMap.
func ResourceManagementThreatProfileV0() *schema.Resource {
	return &schema.Resource{
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
			"threat_extraction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is Threat-Extraction blade activated.",
			},
			"zero_phishing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is Zero-Phishing blade activated.",
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

// ResourceManagementThreatProfileStateUpgradeV0 converts overrides[*].default and overrides[*].final from TypeMap to TypeList.
func ResourceManagementThreatProfileStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "overrides"), nil
}
