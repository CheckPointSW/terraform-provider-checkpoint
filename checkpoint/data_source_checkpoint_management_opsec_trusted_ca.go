package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementOpsecTrustedCa() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementOpsecTrustedCaRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate file encoded in base64.",
			},
			"automatic_enrollment": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Certificate automatic enrollment.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatically_enroll_certificate": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to automatically enroll certificate.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol that communicates with the certificate authority. Available only if \"automatically-enroll-certificate\" parameter is set to true.",
						},
						"scep_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Scep protocol settings. Available only if \"protocol\" is set to \"scep\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ca_identifier": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate authority identifier.",
									},
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate authority URL.",
									},
								},
							},
						},
						"cmpv1_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cmpv1 protocol settings. Available only if \"protocol\" is set to \"cmpv1\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direct_tcp_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Direct tcp transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Certificate authority IP address.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Port number.",
												},
											},
										},
									},
								},
							},
						},
						"cmpv2_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cmpv2 protocol settings. Available only if \"protocol\" is set to \"cmpv1\".",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"transport_layer": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Transport layer.",
									},
									"direct_tcp_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Direct tcp transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Certificate authority IP address.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Port number.",
												},
											},
										},
									},
									"http_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Http transport layer settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:        schema.TypeString,
													Computed:    true,
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
				Computed:    true,
				Description: "Whether to retrieve Certificate Revocation List from http servers.",
			},
			"retrieve_crl_from_ldap_servers": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to retrieve Certificate Revocation List from ldap servers.",
			},
			"cache_crl": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Cache Certificate Revocation List on the Security Gateway.",
			},
			"crl_cache_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period.",
			},
			"crl_cache_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "When to fetch new Certificate Revocation List (in minutes).",
			},
			"allow_certificates_from_branches": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow only certificates from listed branches.",
			},
			"branches": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Branches to allow certificates from. Required only if \"allow-certificates-from-branches\" set to \"true\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func dataSourceManagementOpsecTrustedCaRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := opsecTrustedCa["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

					if v, _ := cmpv2SettingsMap["transport-layer"]; v != nil {
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
