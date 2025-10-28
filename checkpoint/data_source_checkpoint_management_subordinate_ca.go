package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementSubordinateCa() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSubordinateCaRead,
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
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate file encoded in base64.",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
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
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Icon name.",
			},
		},
	}
}

func dataSourceManagementSubordinateCaRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showSubordinateCaRes, err := client.ApiCall("show-subordinate-ca", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSubordinateCaRes.Success {
		return fmt.Errorf(showSubordinateCaRes.ErrorMsg)
	}

	subordinateCa := showSubordinateCaRes.GetData()

	log.Println("Read SubordinateCa - Show JSON = ", subordinateCa)

	if v := subordinateCa["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := subordinateCa["name"]; v != nil {
		_ = d.Set("name", v)
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

				if v, _ := directTcpSettingsMap["port"]; v != nil {
					directTcpSettingsMapToReturn["port"] = v
				}
				cmpv1SettingsMapToReturn["direct_tcp_settings"] = []interface{}{directTcpSettingsMapToReturn}
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

				if v, _ := directTcpSettingsMap["port"]; v != nil {
					directTcpSettingsMapToReturn["port"] = v
				}
				cmpv2SettingsMapToReturn["direct_tcp_settings"] = []interface{}{directTcpSettingsMapToReturn}
			}

			if v := cmpv2SettingsMap["http-settings"]; v != nil {

				httpSettingsMap := v.(map[string]interface{})
				httpSettingsMapToReturn := make(map[string]interface{})

				if v, _ := httpSettingsMap["url"]; v != nil {
					httpSettingsMapToReturn["url"] = v
				}
				cmpv2SettingsMapToReturn["http_settings"] = []interface{}{httpSettingsMapToReturn}
			}

			automaticEnrollmentMapToReturn["cmpv2_settings"] = []interface{}{cmpv2SettingsMapToReturn}
		}

		_ = d.Set("automatic_enrollment", []interface{}{automaticEnrollmentMapToReturn})

	} else {
		_ = d.Set("automatic_enrollment", nil)
	}

	if v := subordinateCa["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", v)
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

	if v := subordinateCa["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
