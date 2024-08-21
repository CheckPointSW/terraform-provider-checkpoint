package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementOpsecTrustedCa() *schema.Resource {
	return &schema.Resource{
		Create: createManagementOpsecTrustedCa,
		Read:   readManagementOpsecTrustedCa,
		Update: updateManagementOpsecTrustedCa,
		Delete: deleteManagementOpsecTrustedCa,
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
													Description: "Certificate authority IP address.",
												},
												"port": {
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
													Description: "Certificate authority IP address.",
												},
												"port": {
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
			"retrieve_crl_from_http_servers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve Certificate Revocation List from http servers.",
				Default:     true,
			},
			"retrieve_crl_from_ldap_servers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve Certificate Revocation List from ldap servers.",
				Default:     false,
			},
			"cache_crl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Cache Certificate Revocation List on the Security Gateway.",
				Default:     true,
			},
			"crl_cache_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period.",
				Default:     "timeout",
			},
			"crl_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "When to fetch new Certificate Revocation List (in minutes).",
				Default:     1440,
			},
			"allow_certificates_from_branches": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Allow only certificates from listed branches.",
				Default:     false,
			},
			"branches": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Branches to allow certificates from. Required only if \"allow-certificates-from-branches\" set to \"true\".",
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
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

func createManagementOpsecTrustedCa(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	opsecTrustedCa := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		opsecTrustedCa["name"] = v.(string)
	}

	if v, ok := d.GetOk("base64_certificate"); ok {
		opsecTrustedCa["base64-certificate"] = v.(string)
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

				if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings.0.direct_tcp_settings"); ok {

					mapToReturn := make(map[string]interface{})

					directTcpMap := v.([]interface{})[0].(map[string]interface{})

					if v := directTcpMap["ip_address"]; v != nil {
						mapToReturn["ip-address"] = v
					}
					if v := directTcpMap["port"]; v != nil {
						mapToReturn["port"] = v
					}
					cmpv1SettingsPayload["direct-tcp-settings"] = mapToReturn
				}
				automaticEnrollmentPayload["cmpv1-settings"] = cmpv1SettingsPayload
			}
			if _, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings"); ok {

				cmpv2SettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.transport_layer"); ok {
					cmpv2SettingsPayload["transport-layer"] = v.(string)
				}
				if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.direct_tcp_settings"); ok {
					mapToReturn := make(map[string]interface{})

					directTcpMap := v.([]interface{})[0].(map[string]interface{})

					if v := directTcpMap["ip_address"]; v != nil {
						mapToReturn["ip-address"] = v
					}
					if v := directTcpMap["port"]; v != nil {
						mapToReturn["port"] = v
					}
					cmpv2SettingsPayload["direct-tcp-settings"] = mapToReturn
				}
				if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings.0.http_settings"); ok {

					mapToReturn := make(map[string]interface{})

					directTcpMap := v.([]interface{})[0].(map[string]interface{})

					if v := directTcpMap["url"]; v != nil {
						mapToReturn["url"] = v
					}

					cmpv2SettingsPayload["http-settings"] = directTcpMap
				}
				automaticEnrollmentPayload["cmpv2-settings"] = cmpv2SettingsPayload
			}
			opsecTrustedCa["automatic-enrollment"] = automaticEnrollmentPayload
		}
	}
	if v, ok := d.GetOkExists("retrieve_crl_from_http_servers"); ok {
		opsecTrustedCa["retrieve-crl-from-http-servers"] = v.(bool)
	}

	if v, ok := d.GetOkExists("retrieve_crl_from_ldap_servers"); ok {
		opsecTrustedCa["retrieve-crl-from-ldap-servers"] = v.(bool)
	}

	if v, ok := d.GetOkExists("cache_crl"); ok {
		opsecTrustedCa["cache-crl"] = v.(bool)
	}

	if v, ok := d.GetOk("crl_cache_method"); ok {
		opsecTrustedCa["crl-cache-method"] = v.(string)
	}

	if v, ok := d.GetOk("crl_cache_timeout"); ok {
		opsecTrustedCa["crl-cache-timeout"] = v.(int)
	}

	if v, ok := d.GetOkExists("allow_certificates_from_branches"); ok {
		opsecTrustedCa["allow-certificates-from-branches"] = v.(bool)
	}

	if v, ok := d.GetOk("branches"); ok {
		opsecTrustedCa["branches"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		opsecTrustedCa["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		opsecTrustedCa["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		opsecTrustedCa["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		opsecTrustedCa["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		opsecTrustedCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		opsecTrustedCa["ignore-errors"] = v.(bool)
	}

	log.Println("Create OpsecTrustedCa - Map = ", opsecTrustedCa)

	addOpsecTrustedCaRes, err := client.ApiCall("add-opsec-trusted-ca", opsecTrustedCa, client.GetSessionID(), true, false)
	if err != nil || !addOpsecTrustedCaRes.Success {
		if addOpsecTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(addOpsecTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addOpsecTrustedCaRes.GetData()["uid"].(string))

	return readManagementOpsecTrustedCa(d, m)
}

func readManagementOpsecTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showOpsecTrustedCaRes, err := client.ApiCall("show-opsec-trusted-ca", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOpsecTrustedCaRes.Success {
		if objectNotFound(showOpsecTrustedCaRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOpsecTrustedCaRes.ErrorMsg)
	}

	opsecTrustedCa := showOpsecTrustedCaRes.GetData()

	log.Println("Read OpsecTrustedCa - Show JSON = ", opsecTrustedCa)

	if v := opsecTrustedCa["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := opsecTrustedCa["base64-certificate"]; v != nil {

		_ = d.Set("base64_certificate", cleanseCertificate(v.(string)))
	}

	if opsecTrustedCa["automatic-enrollment"] != nil {

		automaticEnrollmentMap, ok := opsecTrustedCa["automatic-enrollment"].(map[string]interface{})

		if ok {
			automaticEnrollmentMapToReturn := make(map[string]interface{})

			if v := automaticEnrollmentMap["automatically-enroll-certificate"]; v != nil {
				automaticEnrollmentMapToReturn["automatically_enroll_certificate"] = v
			}
			if v := automaticEnrollmentMap["protocol"]; v != nil {
				automaticEnrollmentMapToReturn["protocol"] = v
			}
			if v, ok := automaticEnrollmentMap["scep-settings"]; ok {

				scepSettingsMap, ok := v.(map[string]interface{})
				if ok {
					scepSettingsMapToReturn := make(map[string]interface{})

					if v, _ := scepSettingsMap["ca-identifier"]; v != nil {
						scepSettingsMapToReturn["ca_identifier"] = v
					}
					if v, _ := scepSettingsMap["url"]; v != nil {
						scepSettingsMapToReturn["url"] = v
					}
					automaticEnrollmentMapToReturn["scep_settings"] = []interface{}{scepSettingsMapToReturn}
				}
			}
			if v, ok := automaticEnrollmentMap["cmpv1-settings"]; ok {

				cmpv1SettingsMap, ok := v.(map[string]interface{})
				if ok {
					cmpv1SettingsMapToReturn := make(map[string]interface{})

					if v, _ := cmpv1SettingsMap["direct-tcp-settings"]; v != nil {

						directTcpMapToReturn := make(map[string]interface{})

						directTcpMap := v.(map[string]interface{})

						if v := directTcpMap["ip-address"]; v != nil {
							directTcpMapToReturn["ip_address"] = v
						}
						if v := directTcpMap["port"]; v != nil {
							directTcpMapToReturn["port"] = v
						}
						cmpv1SettingsMapToReturn["direct_tcp_settings"] = []interface{}{directTcpMapToReturn}
					}
					automaticEnrollmentMapToReturn["cmpv1_settings"] = []interface{}{cmpv1SettingsMapToReturn}
				}
			}
			if v, ok := automaticEnrollmentMap["cmpv2-settings"]; ok {

				cmpv2SettingsMap, ok := v.(map[string]interface{})
				if ok {
					cmpv2SettingsMapToReturn := make(map[string]interface{})

					if v := cmpv2SettingsMap["transport-layer"]; v != nil {
						cmpv2SettingsMapToReturn["transport_layer"] = v
					}
					if v, _ := cmpv2SettingsMap["direct-tcp-settings"]; v != nil {

						directTcpMapToReturn := make(map[string]interface{})

						directTcpMap := v.(map[string]interface{})

						if v := directTcpMap["ip-address"]; v != nil {
							directTcpMapToReturn["ip_address"] = v
						}
						if v := directTcpMap["port"]; v != nil {
							directTcpMapToReturn["port"] = v
						}
						cmpv2SettingsMapToReturn["direct_tcp_settings"] = []interface{}{directTcpMapToReturn}

					}
					if v, _ := cmpv2SettingsMap["http-settings"]; v != nil {

						httpSettingsMapToReturn := make(map[string]interface{})

						httpSettingsTcpMap := v.(map[string]interface{})

						if v := httpSettingsTcpMap["url"]; v != nil {
							httpSettingsMapToReturn["url"] = v
						}

						cmpv2SettingsMapToReturn["http_settings"] = []interface{}{httpSettingsMapToReturn}
					}
					automaticEnrollmentMapToReturn["cmpv2_settings"] = []interface{}{cmpv2SettingsMapToReturn}
				}
			}
			_ = d.Set("automatic_enrollment", []interface{}{automaticEnrollmentMapToReturn})

		}
	} else {
		_ = d.Set("automatic_enrollment", nil)
	}

	if v := opsecTrustedCa["retrieve-crl-from-http-servers"]; v != nil {
		_ = d.Set("retrieve_crl_from_http_servers", v)
	}

	if v := opsecTrustedCa["retrieve-crl-from-ldap-servers"]; v != nil {
		_ = d.Set("retrieve_crl_from_ldap_servers", v)
	}

	if v := opsecTrustedCa["cache-crl"]; v != nil {
		_ = d.Set("cache_crl", v)
	}

	if v := opsecTrustedCa["crl-cache-method"]; v != nil {
		_ = d.Set("crl_cache_method", v)
	}

	if v := opsecTrustedCa["crl-cache-timeout"]; v != nil {
		_ = d.Set("crl_cache_timeout", v)
	}

	if v := opsecTrustedCa["allow-certificates-from-branches"]; v != nil {
		_ = d.Set("allow_certificates_from_branches", v)
	}

	if opsecTrustedCa["branches"] != nil {
		branchesJson, ok := opsecTrustedCa["branches"].([]interface{})
		if ok {
			branchesIds := make([]string, 0)
			if len(branchesJson) > 0 {
				for _, branches := range branchesJson {
					branches := branches.(map[string]interface{})
					branchesIds = append(branchesIds, branches["name"].(string))
				}
			}
			_ = d.Set("branches", branchesIds)
		}
	} else {
		_ = d.Set("branches", nil)
	}

	if opsecTrustedCa["tags"] != nil {
		tagsJson, ok := opsecTrustedCa["tags"].([]interface{})
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

	if v := opsecTrustedCa["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := opsecTrustedCa["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if opsecTrustedCa["domains_to_process"] != nil {
		domainsToProcessJson, ok := opsecTrustedCa["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	if v := opsecTrustedCa["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := opsecTrustedCa["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementOpsecTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	opsecTrustedCa := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		opsecTrustedCa["name"] = oldName
		opsecTrustedCa["new-name"] = newName
	} else {
		opsecTrustedCa["name"] = d.Get("name")
	}

	if ok := d.HasChange("base64_certificate"); ok {
		opsecTrustedCa["base64-certificate"] = d.Get("base64_certificate")
	}

	if d.HasChange("automatic_enrollment") {

		if v, ok := d.GetOk("automatic_enrollment"); ok {

			automaticEnrollmentList := v.([]interface{})

			if len(automaticEnrollmentList) > 0 {

				automaticEnrollmentPayload := make(map[string]interface{})

				automaticEnrollmentMap := automaticEnrollmentList[0].(map[string]interface{})

				if v := automaticEnrollmentMap["automatically_enroll_certificate"]; v != nil {
					automaticEnrollmentPayload["automatically-enroll-certificate"] = v
				}

				if v := automaticEnrollmentMap["protocol"]; v != nil {
					automaticEnrollmentPayload["protocol"] = v
				}

				if _, ok := d.GetOk("automatic_enrollment.0.scep_settings"); ok {

					scepSettingsPayload := make(map[string]interface{})

					if v := d.Get("automatic_enrollment.0.scep_settings.0.ca_identifier"); v != nil {
						scepSettingsPayload["ca-identifier"] = v
					}
					if v := d.Get("automatic_enrollment.0.scep_settings.0.url"); v != nil {
						scepSettingsPayload["url"] = v
					}
					automaticEnrollmentPayload["scep-settings"] = scepSettingsPayload
				}

				if v, ok := d.GetOk("automatic_enrollment.0.cmpv1_settings"); ok {

					directTcpMapToReturn := make(map[string]interface{})

					existingMap := v.([]interface{})[0].(map[string]interface{})

					if v := existingMap["direct_tcp_settings"]; v != nil {

						if len(v.([]interface{})) > 0 {

							innerMap := v.([]interface{})[0].(map[string]interface{})

							payload := make(map[string]interface{})

							if v := innerMap["ip_address"]; v != nil {
								payload["ip-address"] = v
							}
							if v := innerMap["port"]; v != nil {
								payload["port"] = v
							}

							directTcpMapToReturn["direct-tcp-settings"] = payload

						}
					}

					automaticEnrollmentPayload["cmpv1-settings"] = directTcpMapToReturn
				}

				if v, ok := d.GetOk("automatic_enrollment.0.cmpv2_settings"); ok {

					cmpv2SettingsPayload := make(map[string]interface{})

					cmpv2Map := v.([]interface{})[0].(map[string]interface{})

					if v := cmpv2Map["transport_layer"]; v != nil {
						cmpv2SettingsPayload["transport-layer"] = v
					}

					if v := cmpv2Map["direct_tcp_settings"]; v != nil {

						if len(v.([]interface{})) > 0 {

							innerMap := v.([]interface{})[0].(map[string]interface{})

							payload := make(map[string]interface{})

							if v := innerMap["ip_address"]; v != nil {
								payload["ip-address"] = v
							}
							if v := innerMap["port"]; v != nil {
								payload["port"] = v
							}

							cmpv2SettingsPayload["direct-tcp-settings"] = payload

						}
					}

					if v := cmpv2Map["http_settings"]; v != nil {

						if len(v.([]interface{})) > 0 {

							innerMap := v.([]interface{})[0].(map[string]interface{})

							payload := make(map[string]interface{})

							if v := innerMap["url"]; v != nil {
								payload["url"] = v
							}

							cmpv2SettingsPayload["http-settings"] = payload
						}

					}

					automaticEnrollmentPayload["cmpv2-settings"] = cmpv2SettingsPayload
				}
				opsecTrustedCa["automatic-enrollment"] = automaticEnrollmentPayload
			}
		}
	}

	if v, ok := d.GetOkExists("retrieve_crl_from_http_servers"); ok {
		opsecTrustedCa["retrieve-crl-from-http-servers"] = v.(bool)
	}

	if v, ok := d.GetOkExists("retrieve_crl_from_ldap_servers"); ok {
		opsecTrustedCa["retrieve-crl-from-ldap-servers"] = v.(bool)
	}

	if v, ok := d.GetOkExists("cache_crl"); ok {
		opsecTrustedCa["cache-crl"] = v.(bool)
	}

	if ok := d.HasChange("crl_cache_method"); ok {
		opsecTrustedCa["crl-cache-method"] = d.Get("crl_cache_method")
	}

	if ok := d.HasChange("crl_cache_timeout"); ok {
		opsecTrustedCa["crl-cache-timeout"] = d.Get("crl_cache_timeout")
	}

	if v, ok := d.GetOkExists("allow_certificates_from_branches"); ok {
		opsecTrustedCa["allow-certificates-from-branches"] = v.(bool)
	}

	if d.HasChange("branches") {
		if v, ok := d.GetOk("branches"); ok {
			opsecTrustedCa["branches"] = v.(*schema.Set).List()
		} else {
			oldBranches, _ := d.GetChange("branches")
			opsecTrustedCa["branches"] = map[string]interface{}{"remove": oldBranches.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			opsecTrustedCa["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			opsecTrustedCa["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		opsecTrustedCa["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		opsecTrustedCa["comments"] = d.Get("comments")
	}

	if d.HasChange("domains_to_process") {
		if v, ok := d.GetOk("domains_to_process"); ok {
			opsecTrustedCa["domains_to_process"] = v.(*schema.Set).List()
		} else {
			oldDomains_To_Process, _ := d.GetChange("domains_to_process")
			opsecTrustedCa["domains_to_process"] = map[string]interface{}{"remove": oldDomains_To_Process.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		opsecTrustedCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		opsecTrustedCa["ignore-errors"] = v.(bool)
	}

	log.Println("Update OpsecTrustedCa - Map = ", opsecTrustedCa)

	updateOpsecTrustedCaRes, err := client.ApiCall("set-opsec-trusted-ca", opsecTrustedCa, client.GetSessionID(), true, false)
	if err != nil || !updateOpsecTrustedCaRes.Success {
		if updateOpsecTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(updateOpsecTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementOpsecTrustedCa(d, m)
}

func deleteManagementOpsecTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	opsecTrustedCaPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete OpsecTrustedCa")

	deleteOpsecTrustedCaRes, err := client.ApiCall("delete-opsec-trusted-ca", opsecTrustedCaPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteOpsecTrustedCaRes.Success {
		if deleteOpsecTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(deleteOpsecTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
