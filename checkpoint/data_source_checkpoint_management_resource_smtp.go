package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementResourceSmtp() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementResourceSmtpRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"mail_delivery_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specify the server to which mail is forwarded.",
			},
			"deliver_messages_using_dns_mx_records": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "MX record resolving is used to set the destination IP address of the connection.",
			},
			"check_rulebase_with_new_destination": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The Rule Base will be rechecked with the new resolved IP address for mail delivery.",
			},
			"notify_sender_on_error": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable error mail delivery.",
			},
			"error_mail_delivery_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Error mail delivery happens if the SMTP security server is unable to deliver the message within the abandon time, and Notify Sender on Error is checked.",
			},
			"error_deliver_messages_using_dns_mx_records": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "MX record resolving will be used to set the source IP address of the connection used to send the error message.",
			},
			"error_check_rulebase_with_new_destination": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The Rule Base will be rechecked with the new resolved IP address for error mail delivery.",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Determines if an action specified in the Action 2 and CVP categories taken as a result of a resource definition is logged.",
			},
			"match": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Set the Match properties for the SMTP resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sender": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Set the Match sender property for the SMTP resource.",
						},
						"recipient": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Set the Match recipient property for the SMTP resource.",
						},
					},
				},
			},
			"action_1": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Use the Rewriting Rules to rewrite Sender and Recipient headers in emails, you can also rewrite other email headers by using the custom header field.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sender": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Rewrite Sender header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Replacement field.",
									},
								},
							},
						},
						"recipient": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Rewrite Recipient header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Replacement field.",
									},
								},
							},
						},
						"custom_field": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The name of the header.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"original": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Original field.",
									},
									"rewritten": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Replacement field.",
									},
									"field": {
										Type:        schema.TypeString,
										Computed:    true,
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
				Computed:    true,
				Description: "Use this window to configure mail inspection for the SMTP Resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strip_mime_of_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the MIME type to strip from the message.",
						},
						"strip_file_by_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Strips file attachments of the specified name from the message.",
						},
						"mail_capacity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Restrict the size (in kb) of incoming email attachments.",
						},
						"allowed_characters": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MIME email headers can consist of 8 or 7 bit characters (7 ASCII and 8 for sending Binary characters) in order to encode mail data.",
						},
						"strip_script_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip JAVA scripts.",
						},
						"strip_applet_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip JAVA applets.",
						},
						"strip_activex_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip activeX tags.",
						},
						"strip_ftp_links": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip ftp links.",
						},
						"strip_port_strings": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip ports.",
						},
					},
				},
			},
			"cvp": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configure CVP inspection on mail messages.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
						},
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
						},
						"reply_order": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementResourceSmtpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showResourceSmtpRes, err := client.ApiCall("show-resource-smtp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceSmtpRes.Success {
		if objectNotFound(showResourceSmtpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceSmtpRes.ErrorMsg)
	}

	resourceSmtp := showResourceSmtpRes.GetData()

	log.Println("Read ResourceSmtp - Show JSON = ", resourceSmtp)

	if v := resourceSmtp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceSmtp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceSmtp["mail-delivery-server"]; v != nil {
		_ = d.Set("mail_delivery_server", v)
	}

	if v := resourceSmtp["deliver-messages-using-dns-mx-records"]; v != nil {
		_ = d.Set("deliver_messages_using_dns_mx_records", v)
	}

	if v := resourceSmtp["check-rulebase-with-new-destination"]; v != nil {
		_ = d.Set("check_rulebase_with_new_destination", v)
	}

	if v := resourceSmtp["notify-sender-on-error"]; v != nil {
		_ = d.Set("notify_sender_on_error", v)
	}

	if v := resourceSmtp["error-mail-delivery-server"]; v != nil {
		_ = d.Set("error_mail_delivery_server", v)
	}

	if v := resourceSmtp["error-deliver-messages-using-dns-mx-records"]; v != nil {
		_ = d.Set("error_deliver_messages_using_dns_mx_records", v)
	}

	if v := resourceSmtp["error-check-rulebase-with-new-destination"]; v != nil {
		_ = d.Set("error_check_rulebase_with_new_destination", v)
	}

	if v := resourceSmtp["exception-track"]; v != nil {
		_ = d.Set("exception_track", v)
	}

	if resourceSmtp["match"] != nil {

		matchMap := resourceSmtp["match"].(map[string]interface{})

		matchMapToReturn := make(map[string]interface{})

		if v, _ := matchMap["sender"]; v != nil {
			matchMapToReturn["sender"] = v
		}
		if v, _ := matchMap["recipient"]; v != nil {
			matchMapToReturn["recipient"] = v
		}
		_ = d.Set("match", matchMapToReturn)
	} else {
		_ = d.Set("match", nil)
	}

	if resourceSmtp["action-1"] != nil {

		action1Map, ok := resourceSmtp["action-1"].(map[string]interface{})

		if ok {
			action1MapToReturn := make(map[string]interface{})

			if v, ok := action1Map["sender"]; ok {

				senderMap, ok := v.(map[string]interface{})
				if ok {
					senderMapToReturn := make(map[string]interface{})

					if v, _ := senderMap["original"]; v != nil {
						senderMapToReturn["original"] = v
					}
					if v, _ := senderMap["rewritten"]; v != nil {
						senderMapToReturn["rewritten"] = v
					}
					action1MapToReturn["sender"] = []interface{}{senderMapToReturn}
				}
			}
			if v, ok := action1Map["recipient"]; ok {

				recipientMap, ok := v.(map[string]interface{})
				if ok {
					recipientMapToReturn := make(map[string]interface{})

					if v, _ := recipientMap["original"]; v != nil {
						recipientMapToReturn["original"] = v
					}
					if v, _ := recipientMap["rewritten"]; v != nil {
						recipientMapToReturn["rewritten"] = v
					}
					action1MapToReturn["recipient"] = []interface{}{recipientMapToReturn}
				}
			}
			if v, ok := action1Map["custom-field"]; ok {

				customFieldMap, ok := v.(map[string]interface{})
				if ok {
					customFieldMapToReturn := make(map[string]interface{})

					if v, _ := customFieldMap["original"]; v != nil {
						customFieldMapToReturn["original"] = v
					}
					if v, _ := customFieldMap["rewritten"]; v != nil {
						customFieldMapToReturn["rewritten"] = v
					}
					if v, _ := customFieldMap["field"]; v != nil {
						customFieldMapToReturn["field"] = v
					}
					action1MapToReturn["custom_field"] = []interface{}{customFieldMapToReturn}
				}
			}
			_ = d.Set("action_1", []interface{}{action1MapToReturn})

		}
	} else {
		_ = d.Set("action_1", nil)
	}

	if resourceSmtp["action-2"] != nil {

		action2Map := resourceSmtp["action-2"].(map[string]interface{})

		action2MapToReturn := make(map[string]interface{})

		if v, _ := action2Map["strip-mime-of-type"]; v != nil {
			action2MapToReturn["strip_mime_of_type"] = v
		}
		if v, _ := action2Map["strip-file-by-name"]; v != nil {
			action2MapToReturn["strip_file_by_name"] = v
		}
		if v, _ := action2Map["mail-capacity"]; v != nil {
			action2MapToReturn["mail_capacity"] = v
		}
		if v, _ := action2Map["allowed-characters"]; v != nil {
			action2MapToReturn["allowed_characters"] = v
		}
		if v, _ := action2Map["strip-script-tags"]; v != nil {
			action2MapToReturn["strip_script_tags"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := action2Map["strip-applet-tags"]; v != nil {
			action2MapToReturn["strip_applet_tags"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := action2Map["strip-activex-tags"]; v != nil {
			action2MapToReturn["strip_activex_tags"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := action2Map["strip-ftp-links"]; v != nil {
			action2MapToReturn["strip_ftp_links"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := action2Map["strip-port-strings"]; v != nil {
			action2MapToReturn["strip_port_strings"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("action_2", []interface{}{action2MapToReturn})
	} else {
		_ = d.Set("action_2", nil)
	}

	if resourceSmtp["cvp"] != nil {

		cvpMap := resourceSmtp["cvp"].(map[string]interface{})

		cvpMapToReturn := make(map[string]interface{})

		if v, _ := cvpMap["enable-cvp"]; v != nil {
			cvpMapToReturn["enable_cvp"] = v
		}
		if v, _ := cvpMap["server"]; v != nil {

			objMap := v.(map[string]interface{})
			if v := objMap["server"]; v != nil {
				cvpMapToReturn["server"] = v
			}
		}
		if v, _ := cvpMap["cvp-server-is-allowed-to-modify-content"]; v != nil {
			cvpMapToReturn["allowed_to_modify_content"] = v
		}
		if v, _ := cvpMap["reply-order"]; v != nil {
			cvpMapToReturn["reply_order"] = v
		}
		_ = d.Set("cvp", []interface{}{cvpMapToReturn})
	} else {
		_ = d.Set("cvp", nil)
	}

	if resourceSmtp["tags"] != nil {
		tagsJson, ok := resourceSmtp["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := resourceSmtp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceSmtp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
