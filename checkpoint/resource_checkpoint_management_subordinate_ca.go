package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSubordinateCa() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSubordinateCa,
		Read:   readManagementSubordinateCa,
		Update: updateManagementSubordinateCa,
		Delete: deleteManagementSubordinateCa,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate file encoded in base64.",
			},
			"automatic_enrollment": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Certificate automatic enrollment.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatically_enroll_certificate": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to automatically enroll certificate.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Protocol that communicates with the certificate authority. Available only if \"automatically-enroll-certificate\" parameter is set to true.",
							Default:     "scep",
						},
						"scep_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Scep protocol settings. Available only if \"protocol\" is set to \"scep\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ca_identifier": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Certificate authority identifier.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Certificate authority URL.",
									},
								},
							},
						},
						"cmpv1_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Cmpv1 protocol settings. Available only if \"protocol\" is set to \"cmpv1\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direct_tcp_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Direct tcp transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Port number.",
												}, "port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Port number.",
													Default:     829,
												},
											},
										},
									},
								},
							},
						},
						"cmpv2_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Cmpv2 protocol settings. Available only if \"protocol\" is set to \"cmpv1\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"transport_layer": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Transport layer.",
									},
									"direct_tcp_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Direct tcp transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Port number.",
												}, "port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Port number.",
													Default:     829,
												},
											},
										},
									},
									"http_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Http transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Certificate authority URL.",
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
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func createManagementSubordinateCa(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	subordinateCa := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		subordinateCa["name"] = v.(string)
	}

	if v, ok := d.GetOk("base64_certificate"); ok {
		subordinateCa["base64-certificate"] = v.(string)
	}

	if v, ok := d.GetOk("automatic_enrollment"); ok {

		automaticEnrollmentList := v.([]interface{})

		if len(automaticEnrollmentList) > 0 {

			automaticEnrollmentPayload := make(map[string]interface{})

			if v, ok := d.GetOk("automatic_enrollment.0.automatically_enroll_certificate"); ok {
				automaticEnrollmentPayload["automatically-enroll-certificate"] = v.(bool)
			}
			if v, ok := d.GetOk("automatic_enrollment.0.protocol"); ok {
				automaticEnrollmentPayload["protocol"] = v.(string)
			}
			if _, ok := d.GetOk("automatic_enrollment.0.scep_settings"); ok {

				scepSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("automatic_enrollment.0.scep_settings.0.ca_identifier"); ok {
					scepSettingsPayload["ca-identifier"] = v.(string)
				}
				if v, ok := d.GetOk("automatic_enrollment.0.scep_settings.0.url"); ok {
					scepSettingsPayload["url"] = v.(string)
				}
				automaticEnrollmentPayload["scep-settings"] = scepSettingsPayload
			}
			if _, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings"); ok {

				cmpv1SettingsPayload := make(map[string]interface{})

				if _, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings"); ok {
					cmpv1SettingsDirectPayload := make(map[string]interface{})

					if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings.0.ip_address"); ok {
						cmpv1SettingsDirectPayload["ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings.0.port"); ok {
						cmpv1SettingsDirectPayload["port"] = v.(int)
					}
					cmpv1SettingsPayload["direct-tcp-settings"] = cmpv1SettingsDirectPayload
				}
				automaticEnrollmentPayload["cmpv1-settings"] = cmpv1SettingsPayload
			}
			if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings"); ok {

				cmpv2SettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.transport_layer"); ok {
					cmpv2SettingsPayload["transport-layer"] = v.(string)
				}
				if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings"); ok {
					cmpv2SettingsDirectPayload := make(map[string]interface{})

					if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings.0.ip_address"); ok {
						cmpv2SettingsDirectPayload["ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings.0.port"); ok {
						cmpv2SettingsDirectPayload["port"] = v.(int)
					}
					cmpv2SettingsPayload["direct-tcp-settings"] = cmpv2SettingsDirectPayload
				}
				if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.http_settings"); ok {
					cmpv2SettingsHttpPayload := make(map[string]interface{})

					if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.http_settings.0.url"); ok {
						cmpv2SettingsHttpPayload["url"] = v.(string)
					}
					cmpv2SettingsPayload["http-settings"] = cmpv2SettingsHttpPayload
				}
				automaticEnrollmentPayload["cmpv2-settings"] = cmpv2SettingsPayload
			}
			subordinateCa["automatic-enrollment"] = automaticEnrollmentPayload
		}
	}

	if v, ok := d.GetOk("color"); ok {
		subordinateCa["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		subordinateCa["comments"] = v.(string)
	}

	if v, ok := d.GetOk("groups"); ok {
		subordinateCa["groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		subordinateCa["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		subordinateCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		subordinateCa["ignore-errors"] = v.(bool)
	}

	log.Println("Create SubordinateCa - Map = ", subordinateCa)

	addSubordinateCaRes, err := client.ApiCall("add-subordinate-ca", subordinateCa, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addSubordinateCaRes.Success {
		if addSubordinateCaRes.ErrorMsg != "" {
			return fmt.Errorf(addSubordinateCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSubordinateCaRes.GetData()["uid"].(string))

	return readManagementSubordinateCa(d, m)
}

func readManagementSubordinateCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSubordinateCaRes, err := client.ApiCall("show-subordinate-ca", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSubordinateCaRes.Success {
		if objectNotFound(showSubordinateCaRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSubordinateCaRes.ErrorMsg)
	}

	subordinateCa := showSubordinateCaRes.GetData()

	log.Println("Read SubordinateCa - Show JSON = ", subordinateCa)

	if v := subordinateCa["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := subordinateCa["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", cleanseCertificate(v.(string)))
	}

	if subordinateCa["automatic-enrollment"] != nil {

		automaticEnrollmentMap := subordinateCa["automatic-enrollment"].(map[string]interface{})

		automaticEnrollmentMapToReturn := make(map[string]interface{})

		if v := automaticEnrollmentMap["automatically-enroll-certificate"]; v != nil {
			automaticEnrollmentMapToReturn["automatically_enroll_certificate"] = v
		}
		if v := automaticEnrollmentMap["protocol"]; v != nil {
			automaticEnrollmentMapToReturn["protocol"] = v
		}
		if v := automaticEnrollmentMap["scep-settings"]; v != nil {

			scepSettingsMap := v.(map[string]interface{})
			scepSettingsMapToReturn := make(map[string]interface{})

			if v, _ := scepSettingsMap["ca-identifier"]; v != nil {
				scepSettingsMapToReturn["ca_identifier"] = v
			}
			if v, _ := scepSettingsMap["url"]; v != nil {
				scepSettingsMapToReturn["url"] = v
			}
			automaticEnrollmentMapToReturn["scep_settings"] = []interface{}{scepSettingsMapToReturn}
		}

		if v := automaticEnrollmentMap["cmpv1-settings"]; v != nil {

			cmpv1SettingsMap := v.(map[string]interface{})
			cmpv1SettingsMapToReturn := make(map[string]interface{})

			if v := cmpv1SettingsMap["direct-tcp-settings"]; v != nil {

				directTcpSettingsMap := v.(map[string]interface{})
				directTcpSettingsMapToReturn := make(map[string]interface{})

				if v, _ := directTcpSettingsMap["ip-address"]; v != nil {
					directTcpSettingsMapToReturn["ip_address"] = v
				}
				if v, _ := directTcpSettingsMap["port"]; v != nil {
					directTcpSettingsMapToReturn["port"] = v
				}
				cmpv1SettingsMapToReturn["direct_tcp_settings"] = directTcpSettingsMapToReturn
			}

			automaticEnrollmentMapToReturn["cmpv1_settings"] = []interface{}{cmpv1SettingsMapToReturn}
		}

		if v := automaticEnrollmentMap["cmpv2-settings"]; v != nil {

			cmpv2SettingsMap := v.(map[string]interface{})
			cmpv2SettingsMapToReturn := make(map[string]interface{})

			if v, _ := cmpv2SettingsMap["transport-layer"]; v != nil {
				cmpv2SettingsMapToReturn["transport_layer"] = v
			}
			if v := cmpv2SettingsMap["direct-tcp-settings"]; v != nil {

				directTcpSettingsMap := v.(map[string]interface{})
				directTcpSettingsMapToReturn := make(map[string]interface{})

				if v, _ := directTcpSettingsMap["ip-address"]; v != nil {
					directTcpSettingsMapToReturn["ip_address"] = v
				}

				if v, _ := directTcpSettingsMap["port"]; v != nil {
					directTcpSettingsMapToReturn["port"] = v
				}
				cmpv2SettingsMapToReturn["direct_tcp_settings"] = directTcpSettingsMapToReturn
			}

			if v := cmpv2SettingsMap["http-settings"]; v != nil {

				httpSettingsMap := v.(map[string]interface{})
				httpSettingsMapToReturn := make(map[string]interface{})

				if v, _ := httpSettingsMap["url"]; v != nil {
					httpSettingsMapToReturn["url"] = v
				}
				cmpv2SettingsMapToReturn["http_settings"] = httpSettingsMapToReturn
			}

			automaticEnrollmentMapToReturn["cmpv2_settings"] = []interface{}{cmpv2SettingsMapToReturn}
		}

		_ = d.Set("automatic_enrollment", []interface{}{automaticEnrollmentMapToReturn})

	} else {
		_ = d.Set("automatic_enrollment", nil)
	}

	if v := subordinateCa["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := subordinateCa["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if subordinateCa["groups"] != nil {
		groupsJson, ok := subordinateCa["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, groups := range groupsJson {
					groups := groups.(map[string]interface{})
					groupsIds = append(groupsIds, groups["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	if subordinateCa["tags"] != nil {
		tagsJson, ok := subordinateCa["tags"].([]interface{})
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

	if v := subordinateCa["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := subordinateCa["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSubordinateCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	subordinateCa := make(map[string]interface{})

	subordinateCa["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			subordinateCa["new-name"] = v.(string)
		}
	}

	if d.HasChange("automatic_enrollment") {

		if v, ok := d.GetOk("automatic_enrollment"); ok {

			automaticEnrollmentList := v.([]interface{})

			if len(automaticEnrollmentList) > 0 {

				automaticEnrollmentPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("automatic_enrollment.0.automatically_enroll_certificate"); ok {
					automaticEnrollmentPayload["automatically-enroll-certificate"] = v.(bool)
				}
				if v, ok := d.GetOk("automatic_enrollment.0.protocol"); ok {
					automaticEnrollmentPayload["protocol"] = v.(string)
				}
				if _, ok := d.GetOk("automatic_enrollment.0.scep_settings"); ok {
					scepSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("automatic_enrollment.0.scep_settings.0.ca_identifier"); ok {
						scepSettingsPayload["ca-identifier"] = v.(string)
					}
					if v, ok := d.GetOk("automatic_enrollment.0.scep_settings.0.url"); ok {
						scepSettingsPayload["url"] = v.(string)
					}
					automaticEnrollmentPayload["scep-settings"] = scepSettingsPayload
				}
				if _, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings"); ok {
					cmpv1SettingsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings"); ok {
						cmpv1SettingsDirectPayload := make(map[string]interface{})

						if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings.0.ip_address"); ok {
							cmpv1SettingsDirectPayload["ip-address"] = v.(string)
						}
						if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings.0.port"); ok {
							cmpv1SettingsDirectPayload["port"] = v.(int)
						}
						cmpv1SettingsPayload["direct-tcp-settings"] = cmpv1SettingsDirectPayload
					}
					automaticEnrollmentPayload["cmpv1-settings"] = cmpv1SettingsPayload
				}
				if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings"); ok {
					cmpv2SettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.transport_layer"); ok {
						cmpv2SettingsPayload["transport-layer"] = v.(string)
					}
					if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings"); ok {
						cmpv2SettingsDirectPayload := make(map[string]interface{})

						if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings.0.ip_address"); ok {
							cmpv2SettingsDirectPayload["ip-address"] = v.(string)
						}
						if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings.0.port"); ok {
							cmpv2SettingsDirectPayload["port"] = v.(int)
						}
						cmpv2SettingsPayload["direct-tcp-settings"] = cmpv2SettingsDirectPayload
					}

					if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.http_settings"); ok {
						cmpv2SettingsHttpPayload := make(map[string]interface{})

						if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.http_settings.0.url"); ok {
							cmpv2SettingsHttpPayload["url"] = v.(string)
						}
						cmpv2SettingsPayload["http-settings"] = cmpv2SettingsHttpPayload
					}
					automaticEnrollmentPayload["cmpv2-settings"] = cmpv2SettingsPayload
				}
				subordinateCa["automatic-enrollment"] = automaticEnrollmentPayload
			}
		}
	}

	if ok := d.HasChange("base64_certificate"); ok {
		subordinateCa["base64-certificate"] = d.Get("base64_certificate")
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			subordinateCa["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			subordinateCa["comments"] = v.(string)
		}
	}

	if d.HasChange("groups") {
		if v, ok := d.GetOk("groups"); ok {
			subordinateCa["groups"] = v.(*schema.Set).List()
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			subordinateCa["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		subordinateCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		subordinateCa["ignore-errors"] = v.(bool)
	}

	log.Println("Update SubordinateCa - Map = ", subordinateCa)

	updateSubordinateCaRes, err := client.ApiCall("set-subordinate-ca", subordinateCa, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateSubordinateCaRes.Success {
		if updateSubordinateCaRes.ErrorMsg != "" {
			return fmt.Errorf(updateSubordinateCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSubordinateCa(d, m)
}

func deleteManagementSubordinateCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	subordinateCaPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete SubordinateCa")

	deleteSubordinateCaRes, err := client.ApiCall("delete-subordinate-ca", subordinateCaPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteSubordinateCaRes.Success {
		if deleteSubordinateCaRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSubordinateCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
