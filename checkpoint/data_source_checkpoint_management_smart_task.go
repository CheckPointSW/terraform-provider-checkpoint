package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSmartTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSmartTaskRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"action": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The action to be run when the trigger is fired.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_web_request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "When the trigger is fired, sends an HTTPS POST web request to the configured URL.<br>The trigger data will be passed along with the SmartTask's custom data in the request's payload.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URL used for the web request.",
									},
									"fingerprint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SHA1 fingerprint of the URL's SSL certificate. Used to trust servers with self-signed SSL certificates.",
									},
									"override_proxy": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Option to send to the web request via a proxy other than the Management's Server proxy (if defined).",
									},
									"proxy_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URL of the proxy used to send the request.",
									},
									"shared_secret": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Shared secret that can be used by the target server to identify the Management Server.<br>The value will be sent as part of the request in the \"X-chkp-shared-secret\" header.",
									},
									"time_out": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Web Request time-out in seconds.",
									},
								},
							},
						},
						"run_script": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "When the trigger is fired, runs the configured Repository Script on the defined targets.<br>The trigger data is then passed to the script as the first parameter. The parameter is JSON encoded in Base64 format.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository_script": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository script that is executed when the trigger is fired.,  identified by the name or UID.",
									},
									"targets": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Targets to execute the script on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_out": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Script execution time-out in seconds.",
									},
								},
							},
						},
						"send_mail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "When the trigger is fired, sends the configured email to the defined recipients.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mail_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The required settings to send the mail by.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"recipients": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A comma separated list of recipient mail addresses.",
												},
												"sender_email": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "An email address to send the mail from.",
												},
												"subject": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The email subject.",
												},
												"body": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The email body.",
												},
												"attachment": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "What file should be attached to the mail.",
												},
												"bcc_recipients": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A comma separated list of bcc recipient mail addresses.",
												},
												"cc_recipients": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A comma separated list of cc recipient mail addresses.",
												},
											},
										},
									},
									"smtp_server": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The UID or the name a preconfigured SMTP server object.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Object name. Must be unique in the domain",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The SMTP port to use.",
												},
												"server": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The SMTP server address.",
												},
												"authentication": {
													Type:        schema.TypeBool,
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
				Computed:    true,
				Description: "Trigger type associated with the SmartTask.",
			},
			"custom_data": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Per SmartTask custom data in JSON format.<br>When the trigger is fired, the trigger data is converted to JSON. The custom data is then concatenated to the trigger data JSON.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the SmartTask's functionality and options.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the SmartTask is enabled and will run when triggered.",
			},
			"fail_open": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If the action fails to execute, whether to treat the execution failure as an error, or continue.",
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
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}

}
func dataSourceManagementSmartTaskRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showSmartTaskRes, err := client.ApiCall("show-smart-task", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSmartTaskRes.Success {
		return fmt.Errorf(showSmartTaskRes.ErrorMsg)
	}

	smartTask := showSmartTaskRes.GetData()

	log.Println("smart task is ", smartTask)

	if v := smartTask["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

			err = d.Set("action", []interface{}{actionMapToReturn})
			if err != nil {
				return fmt.Errorf(err.Error())
			}
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
