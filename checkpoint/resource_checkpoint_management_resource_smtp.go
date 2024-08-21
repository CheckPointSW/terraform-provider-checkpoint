package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResourceSmtp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceSmtp,
		Read:   readManagementResourceSmtp,
		Update: updateManagementResourceSmtp,
		Delete: deleteManagementResourceSmtp,
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

func createManagementResourceSmtp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceSmtp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceSmtp["name"] = v.(string)
	}

	if v, ok := d.GetOk("mail_delivery_server"); ok {
		resourceSmtp["mail-delivery-server"] = v.(string)
	}

	if v, ok := d.GetOkExists("deliver_messages_using_dns_mx_records"); ok {
		resourceSmtp["deliver-messages-using-dns-mx-records"] = v.(bool)
	}

	if v, ok := d.GetOkExists("check_rulebase_with_new_destination"); ok {
		resourceSmtp["check-rulebase-with-new-destination"] = v.(bool)
	}

	if v, ok := d.GetOkExists("notify_sender_on_error"); ok {
		resourceSmtp["notify-sender-on-error"] = v.(bool)
	}

	if v, ok := d.GetOk("error_mail_delivery_server"); ok {
		resourceSmtp["error-mail-delivery-server"] = v.(string)
	}

	if v, ok := d.GetOkExists("error_deliver_messages_using_dns_mx_records"); ok {
		resourceSmtp["error-deliver-messages-using-dns-mx-records"] = v.(bool)
	}

	if v, ok := d.GetOkExists("error_check_rulebase_with_new_destination"); ok {
		resourceSmtp["error-check-rulebase-with-new-destination"] = v.(bool)
	}

	if v, ok := d.GetOk("exception_track"); ok {
		resourceSmtp["exception-track"] = v
	}

	if _, ok := d.GetOk("match"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("match.sender"); ok {
			res["sender"] = v.(string)
		}
		if v, ok := d.GetOk("match.recipient"); ok {
			res["recipient"] = v.(string)
		}
		resourceSmtp["match"] = res
	}

	if v, ok := d.GetOk("action_1"); ok {

		action1List := v.([]interface{})

		if len(action1List) > 0 {

			action1Payload := make(map[string]interface{})

			if _, ok := d.GetOk("action_1.0.sender"); ok {

				senderPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action_1.0.sender.0.original"); ok {
					senderPayload["original"] = v.(string)
				}
				if v, ok := d.GetOk("action_1.0.sender.0.rewritten"); ok {
					senderPayload["rewritten"] = v.(string)
				}
				action1Payload["sender"] = senderPayload
			}
			if _, ok := d.GetOk("action_1.0.recipient"); ok {

				recipientPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action_1.0.recipient.0.original"); ok {
					recipientPayload["original"] = v.(string)
				}
				if v, ok := d.GetOk("action_1.0.recipient.0.rewritten"); ok {
					recipientPayload["rewritten"] = v.(string)
				}
				action1Payload["recipient"] = recipientPayload
			}
			if _, ok := d.GetOk("action_1.0.custom_field"); ok {

				customFieldPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action_1.0.custom_field.0.original"); ok {
					customFieldPayload["original"] = v.(string)
				}
				if v, ok := d.GetOk("action_1.0.custom_field.0.rewritten"); ok {
					customFieldPayload["rewritten"] = v.(string)
				}
				if v, ok := d.GetOk("action_1.0.custom_field.0.field"); ok {
					customFieldPayload["field"] = v.(string)
				}
				action1Payload["custom-field"] = customFieldPayload
			}
			resourceSmtp["action-1"] = action1Payload
		}
	}
	if v, ok := d.GetOk("action_2"); ok {

		res := make(map[string]interface{})

		v := v.([]interface{})

		action2Map := v[0].(map[string]interface{})

		if v := action2Map["strip_mime_of_type"]; v != nil {
			res["strip-mime-of-type"] = v
		}
		if v := action2Map["strip_file_by_name"]; v != nil {
			res["strip-file-by-name"] = v
		}
		if v := action2Map["mail_capacity"]; v != nil {
			res["mail-capacity"] = v
		}
		if v := action2Map["allowed_characters"]; v != nil {
			res["allowed-characters"] = v
		}
		if v := action2Map["strip_script_tags"]; v != nil {
			res["strip-script-tags"] = v
		}
		if v := action2Map["strip_mime_of_type"]; v != nil {
			res["strip-mime-of-type"] = v
		}
		if v := action2Map["strip_applet_tags"]; v != nil {
			res["strip-applet-tags"] = v
		}
		if v := action2Map["trip_activex_tags"]; v != nil {
			res["trip-activex-tags"] = v
		}
		if v := action2Map["strip_ftp_links"]; v != nil {
			res["strip-ftp-links"] = v
		}
		if v := action2Map["strip_port_strings"]; v != nil {
			res["strip-port-strings"] = v
		}

		resourceSmtp["action-2"] = res
	}

	if v, ok := d.GetOk("cvp"); ok {

		res := make(map[string]interface{})

		v := v.([]interface{})

		cvpMap := v[0].(map[string]interface{})

		if v := cvpMap["enable_cvp"]; v != nil {
			res["enable-cvp"] = v
		}
		if v := cvpMap["server"]; v != nil {
			if len(v.(string)) > 0 {
				res["server"] = v
			}

		}
		if v := cvpMap["allowed_to_modify_content"]; v != nil {
			res["allowed-to-modify-content"] = v
		}
		if v := cvpMap["reply_order"]; v != nil {
			res["reply-order"] = v
		}

		resourceSmtp["cvp"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceSmtp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceSmtp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceSmtp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceSmtp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceSmtp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceSmtp - Map = ", resourceSmtp)

	addResourceSmtpRes, err := client.ApiCall("add-resource-smtp", resourceSmtp, client.GetSessionID(), true, false)
	if err != nil || !addResourceSmtpRes.Success {
		if addResourceSmtpRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceSmtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceSmtpRes.GetData()["uid"].(string))

	return readManagementResourceSmtp(d, m)
}

func readManagementResourceSmtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

		objMap := v.(map[string]interface{})

		if v := objMap["name"]; v != nil {
			_ = d.Set("exception_track", v)
		}

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
			action2MapToReturn["strip_script_tags"] = v
		}
		if v, _ := action2Map["strip-applet-tags"]; v != nil {
			action2MapToReturn["strip_applet_tags"] = v
		}
		if v, _ := action2Map["strip-activex-tags"]; v != nil {
			action2MapToReturn["strip_activex_tags"] = v
		}
		if v, _ := action2Map["strip-ftp-links"]; v != nil {
			action2MapToReturn["strip_ftp_links"] = v
		}
		if v, _ := action2Map["strip-port-strings"]; v != nil {
			action2MapToReturn["strip_port_strings"] = v
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

	if v := resourceSmtp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceSmtp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceSmtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceSmtp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceSmtp["name"] = oldName
		resourceSmtp["new-name"] = newName
	} else {
		resourceSmtp["name"] = d.Get("name")
	}

	if ok := d.HasChange("mail_delivery_server"); ok {
		resourceSmtp["mail-delivery-server"] = d.Get("mail_delivery_server")
	}

	if v, ok := d.GetOkExists("deliver_messages_using_dns_mx_records"); ok {
		resourceSmtp["deliver-messages-using-dns-mx-records"] = v.(bool)
	}

	if v, ok := d.GetOkExists("check_rulebase_with_new_destination"); ok {
		resourceSmtp["check-rulebase-with-new-destination"] = v.(bool)
	}

	if v, ok := d.GetOkExists("notify_sender_on_error"); ok {
		resourceSmtp["notify-sender-on-error"] = v.(bool)
	}

	if ok := d.HasChange("error_mail_delivery_server"); ok {
		resourceSmtp["error-mail-delivery-server"] = d.Get("error_mail_delivery_server")
	}

	if v, ok := d.GetOkExists("error_deliver_messages_using_dns_mx_records"); ok {
		resourceSmtp["error-deliver-messages-using-dns-mx-records"] = v.(bool)
	}

	if v, ok := d.GetOkExists("error_check_rulebase_with_new_destination"); ok {
		resourceSmtp["error-check-rulebase-with-new-destination"] = v.(bool)
	}

	if ok := d.HasChange("exception_track"); ok {
		resourceSmtp["exception-track"] = d.Get("exception_track")
	}

	if d.HasChange("match") {

		if _, ok := d.GetOk("match"); ok {

			res := make(map[string]interface{})

			if d.HasChange("match.sender") {
				res["sender"] = d.Get("match.sender")
			}
			if d.HasChange("match.recipient") {
				res["recipient"] = d.Get("match.recipient")
			}
			resourceSmtp["match"] = res
		}
	}

	if d.HasChange("action_1") {

		if v, ok := d.GetOk("action_1"); ok {

			action1List := v.([]interface{})

			if len(action1List) > 0 {

				action1Payload := make(map[string]interface{})

				if d.HasChange("action_1.0.sender") {

					senderPayload := make(map[string]interface{})

					if d.HasChange("action_1.0.sender.0.original") {
						senderPayload["original"] = d.Get("action_1.0.sender.0.original").(string)
					}
					if d.HasChange("action_1.0.sender.0.rewritten") {
						senderPayload["rewritten"] = d.Get("action_1.0.sender.0.rewritten").(string)
					}
					action1Payload["sender"] = senderPayload
				}
				if d.HasChange("action_1.0.recipient") {

					recipientPayload := make(map[string]interface{})

					if d.HasChange("action_1.0.recipient.0.original") {
						recipientPayload["original"] = d.Get("action_1.0.recipient.0.original").(string)
					}
					if d.HasChange("action_1.0.recipient.0.rewritten") {
						recipientPayload["rewritten"] = d.Get("action_1.0.recipient.0.rewritten").(string)
					}
					action1Payload["recipient"] = recipientPayload
				}
				if d.HasChange("action_1.0.custom_field") {

					customFieldPayload := make(map[string]interface{})

					if d.HasChange("action_1.0.custom_field.0.original") {
						customFieldPayload["original"] = d.Get("action_1.0.custom_field.0.original").(string)
					}
					if d.HasChange("action_1.0.custom_field.0.rewritten") {
						customFieldPayload["rewritten"] = d.Get("action_1.0.custom_field.0.rewritten").(string)
					}
					if d.HasChange("action_1.0.custom_field.0.field") {
						customFieldPayload["field"] = d.Get("action_1.0.custom_field.0.field").(string)
					}
					action1Payload["custom-field"] = customFieldPayload
				}
				resourceSmtp["action-1"] = action1Payload
			}
		}
	}

	if d.HasChange("action_2") {

		if v, ok := d.GetOk("action_2"); ok {

			res := make(map[string]interface{})

			v := v.([]interface{})

			action2Map := v[0].(map[string]interface{})

			if v := action2Map["strip_mime_of_type"]; v != nil {
				res["strip-mime-of-type"] = v
			}
			if v := action2Map["strip_file_by_name"]; v != nil {
				res["strip-file-by-name"] = v
			}
			if v := action2Map["mail_capacity"]; v != nil {
				res["mail-capacity"] = v
			}
			if v := action2Map["allowed_characters"]; v != nil {
				res["allowed-characters"] = v
			}
			if v := action2Map["strip_script_tags"]; v != nil {
				res["strip-script-tags"] = v
			}
			if v := action2Map["strip_mime_of_type"]; v != nil {
				res["strip-mime-of-type"] = v
			}
			if v := action2Map["strip_applet_tags"]; v != nil {
				res["strip-applet-tags"] = v
			}
			if v := action2Map["trip_activex_tags"]; v != nil {
				res["trip-activex-tags"] = v
			}
			if v := action2Map["strip_ftp_links"]; v != nil {
				res["strip-ftp-links"] = v
			}
			if v := action2Map["strip_port_strings"]; v != nil {
				res["strip-port-strings"] = v
			}

			resourceSmtp["action-2"] = res
		}
	}

	if d.HasChange("cvp") {

		if v, ok := d.GetOk("cvp"); ok {

			res := make(map[string]interface{})

			v := v.([]interface{})

			cvpMap := v[0].(map[string]interface{})

			if v := cvpMap["enable_cvp"]; v != nil {
				res["enable-cvp"] = v
			}
			if v := cvpMap["server"]; v != nil {
				if len(v.(string)) > 0 {
					res["server"] = v
				}
			}
			if v := cvpMap["allowed_to_modify_content"]; v != nil {
				res["allowed-to-modify-content"] = v
			}
			if v := cvpMap["reply_order"]; v != nil {
				res["reply-order"] = v
			}

			resourceSmtp["cvp"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceSmtp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			resourceSmtp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceSmtp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceSmtp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceSmtp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceSmtp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceSmtp - Map = ", resourceSmtp)

	updateResourceSmtpRes, err := client.ApiCall("set-resource-smtp", resourceSmtp, client.GetSessionID(), true, false)
	if err != nil || !updateResourceSmtpRes.Success {
		if updateResourceSmtpRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceSmtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceSmtp(d, m)
}

func deleteManagementResourceSmtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceSmtpPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceSmtp")

	deleteResourceSmtpRes, err := client.ApiCall("delete-resource-smtp", resourceSmtpPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteResourceSmtpRes.Success {
		if deleteResourceSmtpRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceSmtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
