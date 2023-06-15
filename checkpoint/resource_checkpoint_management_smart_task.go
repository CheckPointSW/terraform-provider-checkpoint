package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementSmartTask() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSmartTask,
		Read:   readManagementSmartTask,
		Update: updateManagementSmartTask,
		Delete: deleteManagementSmartTask,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"action": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The action to be run when the trigger is fired.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_web_request": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "When the trigger is fired, sends an HTTPS POST web request to the configured URL.<br>The trigger data will be passed along with the SmartTask's custom data in the request's payload.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "URL used for the web request.",
									},
									"fingerprint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The SHA1 fingerprint of the URL's SSL certificate. Used to trust servers with self-signed SSL certificates.",
									},
									"override_proxy": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Option to send to the web request via a proxy other than the Management's Server proxy (if defined).",
									},
									"proxy_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "URL of the proxy used to send the request.",
									},
									"shared_secret": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Shared secret that can be used by the target server to identify the Management Server.<br>The value will be sent as part of the request in the \"X-chkp-shared-secret\" header.",
									},
									"time_out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Web Request time-out in seconds.",
										Default:     30,
									},
								},
							},
						},
						"run_script": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "When the trigger is fired, runs the configured Repository Script on the defined targets.<br>The trigger data is then passed to the script as the first parameter. The parameter is JSON encoded in Base64 format.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository_script": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Repository script that is executed when the trigger is fired.,  identified by the name or UID.",
									},
									"targets": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Targets to execute the script on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Script execution time-out in seconds.",
										Default:     30,
									},
								},
							},
						},
						"send_mail": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "When the trigger is fired, sends the configured email to the defined recipients.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mail_settings": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "The required settings to send the mail by.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"recipients": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "A comma separated list of recipient mail addresses.",
												},
												"sender_email": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "An email address to send the mail from.",
												},
												"subject": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The email subject.",
												},
												"body": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The email body.",
												},
												"attachment": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "What file should be attached to the mail.",
												},
												"bcc_recipients": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "A comma separated list of bcc recipient mail addresses.",
												},
												"cc_recipients": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "A comma separated list of cc recipient mail addresses.",
												},
											},
										},
									},
									"smtp_server": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The UID or the name a preconfigured SMTP server object.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Object name. Must be unique in the domain",
												},
												"port": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The SMTP port to use.",
												},
												"server": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The SMTP server address.",
												},
												"authentication": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Does the mail server requires authentication.",
												},
												"encryption": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Encryption type.",
												},
												"username": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A username for the SMTP server.",
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
			"trigger": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Trigger type associated with the SmartTask.",
			},
			"custom_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Per SmartTask custom data in JSON format.<br>When the trigger is fired, the trigger data is converted to JSON. The custom data is then concatenated to the trigger data JSON.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the SmartTask's functionality and options.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the SmartTask is enabled and will run when triggered.",
				Default:     false,
			},
			"fail_open": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If the action fails to execute, whether to treat the execution failure as an error, or continue.",
				Default:     true,
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

func createManagementSmartTask(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	smartTask := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		smartTask["name"] = v.(string)
	}

	if v, ok := d.GetOk("action"); ok {

		actionList := v.([]interface{})

		if len(actionList) > 0 {

			actionPayload := make(map[string]interface{})

			if _, ok := d.GetOk("action.0.send_web_request"); ok {

				sendWebRequestPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action.0.send_web_request.0.url"); ok {
					sendWebRequestPayload["url"] = v.(string)
				}
				if v, ok := d.GetOk("action.0.send_web_request.0.fingerprint"); ok {
					sendWebRequestPayload["fingerprint"] = v.(string)
				}
				if v, ok := d.GetOk("action.0.send_web_request.0.override_proxy"); ok {
					sendWebRequestPayload["override-proxy"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("action.0.send_web_request.0.proxy_url"); ok {
					sendWebRequestPayload["proxy-url"] = v.(string)
				}
				if v, ok := d.GetOk("action.0.send_web_request.0.shared_secret"); ok {
					sendWebRequestPayload["shared-secret"] = v.(string)
				}
				if v, ok := d.GetOk("action.0.send_web_request.0.time_out"); ok {
					sendWebRequestPayload["time-out"] = v
				}
				actionPayload["send-web-request"] = sendWebRequestPayload
			}
			if _, ok := d.GetOk("action.0.run_script"); ok {

				runScriptPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action.0.run_script.0.repository_script"); ok {
					runScriptPayload["repository-script"] = v.(string)
				}
				if v, ok := d.GetOk("action.0.run_script.0.targets"); ok {
					runScriptPayload["targets"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("action.0.run_script.0.time_out"); ok {
					runScriptPayload["time-out"] = v
				}
				actionPayload["run-script"] = runScriptPayload
			}
			if _, ok := d.GetOk("action.0.send_mail"); ok {

				sendMailPayload := make(map[string]interface{})

				if v, ok := d.GetOk("action.0.send_mail.0.mail_settings"); ok {

					mailSettingsMap := v.([]interface{})[0].(map[string]interface{})

					payload := make(map[string]interface{})

					if v := mailSettingsMap["recipients"]; v != nil {
						payload["recipients"] = v
					}
					if v := mailSettingsMap["sender_email"]; v != nil {
						payload["sender-email"] = v
					}
					if v := mailSettingsMap["subject"]; v != nil {
						payload["subject"] = v
					}
					if v := mailSettingsMap["body"]; v != nil {
						payload["body"] = v
					}
					if v := mailSettingsMap["attachment"]; v != nil {
						if len(v.(string)) > 0 {
							payload["attachment"] = v
						}
					}
					if v := mailSettingsMap["bcc_recipients"]; v != nil {
						payload["bcc-recipients"] = v
					}
					if v := mailSettingsMap["cc_recipients"]; v != nil {
						payload["cc-recipients"] = v
					}
					sendMailPayload["mail-settings"] = payload
				}
				if v, ok := d.GetOk("action.0.send_mail.0.smtp_server.0.name"); ok {
					sendMailPayload["smtp-server"] = v.(string)
				}
				actionPayload["send-mail"] = sendMailPayload
			}
			smartTask["action"] = actionPayload
		}
	}
	if v, ok := d.GetOk("trigger"); ok {
		smartTask["trigger"] = v.(string)
	}

	if v, ok := d.GetOk("custom_data"); ok {
		smartTask["custom-data"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		smartTask["description"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		smartTask["enabled"] = v.(bool)
	}

	if v, ok := d.GetOkExists("fail_open"); ok {
		smartTask["fail-open"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		smartTask["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		smartTask["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		smartTask["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		smartTask["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		smartTask["ignore-errors"] = v.(bool)
	}

	log.Println("Create SmartTask - Map = ", smartTask)

	addSmartTaskRes, err := client.ApiCall("add-smart-task", smartTask, client.GetSessionID(), true, false)
	if err != nil || !addSmartTaskRes.Success {
		if addSmartTaskRes.ErrorMsg != "" {
			return fmt.Errorf(addSmartTaskRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSmartTaskRes.GetData()["uid"].(string))

	return readManagementSmartTask(d, m)
}

func readManagementSmartTask(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSmartTaskRes, err := client.ApiCall("show-smart-task", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSmartTaskRes.Success {
		if objectNotFound(showSmartTaskRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSmartTaskRes.ErrorMsg)
	}

	smartTask := showSmartTaskRes.GetData()

	log.Println("Read SmartTask - Show JSON = ", smartTask)

	if v := smartTask["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if smartTask["action"] != nil {

		actionMap, ok := smartTask["action"].(map[string]interface{})

		if ok {
			actionMapToReturn := make(map[string]interface{})

			if v, ok := actionMap["send-web-request"]; ok {

				sendWebRequestMap, ok := v.(map[string]interface{})
				if ok {
					sendWebRequestMapToReturn := make(map[string]interface{})

					if v, _ := sendWebRequestMap["url"]; v != nil {
						sendWebRequestMapToReturn["url"] = v
					}
					if v, _ := sendWebRequestMap["fingerprint"]; v != nil {
						sendWebRequestMapToReturn["fingerprint"] = v
					}
					if v, _ := sendWebRequestMap["override-proxy"]; v != nil {
						sendWebRequestMapToReturn["override_proxy"] = v
					}
					if v, _ := sendWebRequestMap["proxy-url"]; v != nil {
						sendWebRequestMapToReturn["proxy_url"] = v
					}
					if v, _ := sendWebRequestMap["shared-secret"]; v != nil {
						sendWebRequestMapToReturn["shared_secret"] = v
					}
					if v, _ := sendWebRequestMap["time-out"]; v != nil {
						sendWebRequestMapToReturn["time_out"] = v
					}
					actionMapToReturn["send_web_request"] = []interface{}{sendWebRequestMapToReturn}
				}
			}
			if v, ok := actionMap["run-script"]; ok {

				runScriptMap, ok := v.(map[string]interface{})
				if ok {
					runScriptMapToReturn := make(map[string]interface{})

					if v, _ := runScriptMap["repository-script"]; v != nil {

						payload := v.(map[string]interface{})

						if v := payload["name"]; v != nil {
							runScriptMapToReturn["repository_script"] = v.(string)
						}
					}
					if v, _ := runScriptMap["targets"]; v != nil {
						runScriptMapToReturn["targets"] = v
					}
					if v, _ := runScriptMap["time-out"]; v != nil {
						runScriptMapToReturn["time_out"] = v
					}
					actionMapToReturn["run_script"] = []interface{}{runScriptMapToReturn}
				}
			}
			if v, ok := actionMap["send-mail"]; ok {

				sendMailMap, ok := v.(map[string]interface{})
				if ok {
					sendMailMapToReturn := make(map[string]interface{})

					if v, _ := sendMailMap["mail-settings"]; v != nil {

						innerMap := v.(map[string]interface{})

						res := make(map[string]interface{})

						if v := innerMap["recipients"]; v != nil {
							res["recipients"] = v
						}
						if v := innerMap["sender-email"]; v != nil {
							res["sender_email"] = v
						}
						if v := innerMap["subject"]; v != nil {
							res["subject"] = v
						}
						if v := innerMap["body"]; v != nil {
							res["body"] = v
						}
						if v := innerMap["attachment"]; v != nil {
							res["attachment"] = v
						}
						if v := innerMap["bcc-recipients"]; v != nil {
							res["bcc_recipients"] = v
						}
						if v := innerMap["cc-recipients"]; v != nil {
							res["cc_recipients"] = v
						}
						sendMailMapToReturn["mail_settings"] = []interface{}{res}
					}
					if v, _ := sendMailMap["smtp-server"]; v != nil {

						innerMap := v.(map[string]interface{})

						res := make(map[string]interface{})

						if v := innerMap["name"]; v != nil {
							res["name"] = v
						}
						if v := innerMap["port"]; v != nil {
							res["port"] = v
						}
						if v := innerMap["server"]; v != nil {
							res["server"] = v
						}
						if v := innerMap["authentication"]; v != nil {
							res["authentication"] = v
						}
						if v := innerMap["encryption"]; v != nil {
							res["encryption"] = v
						}
						if v := innerMap["username"]; v != nil {
							res["username"] = v
						}
						sendMailMapToReturn["smtp_server"] = []interface{}{res}
					}
					actionMapToReturn["send_mail"] = []interface{}{sendMailMapToReturn}
				}
			}
			_ = d.Set("action", []interface{}{actionMapToReturn})
		}
	} else {
		_ = d.Set("action", nil)
	}

	if v := smartTask["trigger"]; v != nil {
		payload := v.(map[string]interface{})

		if v := payload["name"]; v != nil {
			_ = d.Set("trigger", v.(string))
		}

	}

	if v := smartTask["custom-data"]; v != nil {
		_ = d.Set("custom_data", v)
	}

	if v := smartTask["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := smartTask["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := smartTask["fail-open"]; v != nil {
		_ = d.Set("fail_open", v)
	}

	if smartTask["tags"] != nil {
		tagsJson, ok := smartTask["tags"].([]interface{})
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

	if v := smartTask["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := smartTask["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := smartTask["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := smartTask["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSmartTask(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	smartTask := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		smartTask["name"] = oldName
		smartTask["new-name"] = newName
	} else {
		smartTask["name"] = d.Get("name")
	}

	if d.HasChange("action") {

		if v, ok := d.GetOk("action"); ok {

			actionList := v.([]interface{})

			if len(actionList) > 0 {

				actionPayload := make(map[string]interface{})

				if _, ok := d.GetOk("action.0.send_web_request"); ok {

					sendWebRequestPayload := make(map[string]interface{})

					if v, ok := d.GetOk("action.0.send_web_request.0.url"); ok {
						sendWebRequestPayload["url"] = v.(string)
					}
					if v, ok := d.GetOk("action.0.send_web_request.0.fingerprint"); ok {
						sendWebRequestPayload["fingerprint"] = v.(string)
					}
					if v, ok := d.GetOk("action.0.send_web_request.0.override_proxy"); ok {
						sendWebRequestPayload["override-proxy"] = v
					}
					if v, ok := d.GetOk("action.0.send_web_request.0.proxy_url"); ok {
						sendWebRequestPayload["proxy-url"] = v.(string)
					}
					if v, ok := d.GetOk("action.0.send_web_request.0.shared_secret"); ok {
						sendWebRequestPayload["shared-secret"] = v.(string)
					}
					if v, ok := d.GetOk("action.0.send_web_request.0.time_out"); ok {
						sendWebRequestPayload["time-out"] = v.(int)
					}
					actionPayload["send-web-request"] = sendWebRequestPayload
				}
				if _, ok := d.GetOk("action.0.run_script"); ok {

					runScriptPayload := make(map[string]interface{})

					if v, ok := d.GetOk("action.0.run_script.0.repository_script"); ok {
						runScriptPayload["repository-script"] = v.(string)
					}
					if v, ok := d.GetOk("action.0.run_script.0.targets"); ok {
						runScriptPayload["targets"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("action.0.run_script.0.time_out"); ok {
						runScriptPayload["time-out"] = v
					}
					actionPayload["run-script"] = runScriptPayload
				}
				if _, ok := d.GetOk("action.0.send_mail"); ok {

					sendMailPayload := make(map[string]interface{})

					if v, ok := d.GetOk("action.0.send_mail.0.mail_settings"); ok {

						mailSettingsMap := v.([]interface{})[0].(map[string]interface{})

						payload := make(map[string]interface{})

						if v := mailSettingsMap["recipients"]; v != nil {
							payload["recipients"] = v
						}
						if v := mailSettingsMap["sender_email"]; v != nil {
							payload["sender-email"] = v
						}
						if v := mailSettingsMap["subject"]; v != nil {
							payload["subject"] = v
						}
						if v := mailSettingsMap["body"]; v != nil {
							payload["body"] = v
						}
						if v := mailSettingsMap["attachment"]; v != nil {
							if len(v.(string)) > 0 {
								payload["attachment"] = v
							}
						}
						if v := mailSettingsMap["bcc_recipients"]; v != nil {
							payload["bcc-recipients"] = v
						}
						if v := mailSettingsMap["cc_recipients"]; v != nil {
							payload["cc-recipients"] = v
						}
						sendMailPayload["mail-settings"] = payload
					}
					if v, ok := d.GetOk("action.0.send_mail.0.smtp_server"); ok {
						smtp := v.([]interface{})[0].(map[string]interface{})
						if j := smtp["name"]; j != nil {
							sendMailPayload["smtp-server"] = j
						}
					}
					actionPayload["send-mail"] = sendMailPayload
				}
				smartTask["action"] = actionPayload
			}
		}
	}

	if ok := d.HasChange("trigger"); ok {
		smartTask["trigger"] = d.Get("trigger")
	}

	if ok := d.HasChange("custom_data"); ok {
		smartTask["custom-data"] = d.Get("custom_data")
	}

	if ok := d.HasChange("description"); ok {
		smartTask["description"] = d.Get("description")
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		smartTask["enabled"] = v.(bool)
	}

	if v, ok := d.GetOkExists("fail_open"); ok {
		smartTask["fail-open"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			smartTask["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			smartTask["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		smartTask["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		smartTask["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		smartTask["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		smartTask["ignore-errors"] = v.(bool)
	}

	log.Println("Update SmartTask - Map = ", smartTask)

	updateSmartTaskRes, err := client.ApiCall("set-smart-task", smartTask, client.GetSessionID(), true, false)
	if err != nil || !updateSmartTaskRes.Success {
		if updateSmartTaskRes.ErrorMsg != "" {
			return fmt.Errorf(updateSmartTaskRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSmartTask(d, m)
}

func deleteManagementSmartTask(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	smartTaskPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		smartTaskPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		smartTaskPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete SmartTask")

	deleteSmartTaskRes, err := client.ApiCall("delete-smart-task", smartTaskPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteSmartTaskRes.Success {
		if deleteSmartTaskRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSmartTaskRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
