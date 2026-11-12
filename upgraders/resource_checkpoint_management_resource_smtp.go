package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementResourceSmtpV0 is the V0 schema where match was TypeMap.
func ResourceManagementResourceSmtpV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"mail_delivery_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the server to which mail is forwarded.",
			},
			"deliver_messages_using_dns_mx_records": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "MX record resolving is used to set the destination IP address of the connection.",
				Default:     false,
			},
			"check_rulebase_with_new_destination": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The Rule Base will be rechecked with the new resolved IP address for mail delivery.",
				Default:     false,
			},
			"notify_sender_on_error": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable error mail delivery.",
				Default:     false,
			},
			"error_mail_delivery_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Error mail delivery happens if the SMTP security server is unable to deliver the message within the abandon time, and Notify Sender on Error is checked.",
			},
			"error_deliver_messages_using_dns_mx_records": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "MX record resolving will be used to set the source IP address of the connection used to send the error message.",
				Default:     false,
			},
			"error_check_rulebase_with_new_destination": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The Rule Base will be rechecked with the new resolved IP address for error mail delivery.",
				Default:     false,
			},
			"exception_track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines if an action specified in the Action 2 and CVP categories taken as a result of a resource definition is logged.",
				Default:     "None",
			},
			"match": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Set the Match properties for the SMTP resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sender": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Set the Match sender property for the SMTP resource.",
						},
						"recipient": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Set the Match recipient property for the SMTP resource.",
						},
					},
				},
			},
			"action_1": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Use the Rewriting Rules to rewrite Sender and Recipient headers in emails, you can also rewrite other email headers by using the custom header field.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sender": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Rewrite Sender header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Replacement field.",
									},
								},
							},
						},
						"recipient": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Rewrite Recipient header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Replacement field.",
									},
								},
							},
						},
						"custom_field": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The name of the header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Replacement field.",
									},
									"field": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the header.",
									},
								},
							},
						},
					},
				},
			},
			"action_2": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Use this window to configure mail inspection for the SMTP Resource.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strip_mime_of_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the MIME type to strip from the message.",
						},
						"strip_file_by_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Strips file attachments of the specified name from the message.",
						},
						"mail_capacity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Restrict the size (in kb) of incoming email attachments.",
							Default:     10000,
						},
						"allowed_characters": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The MIME email headers can consist of 8 or 7 bit characters (7 ASCII and 8 for sending Binary characters) in order to encode mail data.",
							Default:     "8 bit",
						},
						"strip_script_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip JAVA scripts.",
							Default:     false,
						},
						"strip_applet_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip JAVA applets.",
							Default:     false,
						},
						"strip_activex_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip activeX tags.",
							Default:     false,
						},
						"strip_ftp_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip ftp links.",
							Default:     false,
						},
						"strip_port_strings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip ports.",
							Default:     false,
						},
					},
				},
			},
			"cvp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Configure CVP inspection on mail messages.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
							Default:     false,
						},
						"server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
							Default:     true,
						},
						"reply_order": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
							Default:     "return_data_after_content_is_approved",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

// ResourceManagementResourceSmtpStateUpgradeV0 converts match from TypeMap to TypeList.
func ResourceManagementResourceSmtpStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "match"), nil
}
