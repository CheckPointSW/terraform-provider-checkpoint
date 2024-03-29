package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementServiceGtp() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementServiceGtpRead,
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
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GTP version.",
			},
			"access_point_name": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Match by Access Point Name.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"apn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Access Point Name object identified by Name or UID.",
						},
					},
				},
			},
			"allow_usage_of_static_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow usage of static IP addresses.",
			},
			"apply_access_policy_on_user_traffic": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Apply Access Policy on user traffic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"add_imsi_field_to_log": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Add IMSI field to logs generated by user traffic.",
						},
					},
				},
			},
			"cs_fallback_and_srvcc": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "CS Fallback and SRVCC (Relevant for V2 only).",
			},
			"imsi_prefix": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Match by IMSI prefix.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IMSI prefix.",
						},
					},
				},
			},
			"interface_profile": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Match only message types relevant to the given GTP interface. Relevant only for GTP V1 or V2.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Interface Profile object identified by Name or UID.",
						},
						"custom_message_types": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The messages types to match on them for this service. To specify a range, add a hyphen between the lowest and the highest numbers, for example: 32-35. Multiple Ranges can be chosen when separated with comma. This field relevant only when the Interface profile is set to 'Custom'.",
						},
					},
				},
			},
			"ldap_group": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Match by an LDAP Group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Ldap Group object identified by Name or UID.",
						},
						"according_to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "According to MS-ISDN or according to IMSI.",
						},
					},
				},
			},
			"ms_isdn": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Match by an MS-ISDN.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"ms_isdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MS-ISDN.",
						},
					},
				},
			},
			"radio_access_technology": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Match by Radio Access Technology.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"utran": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(1).",
						},
						"geran": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(2).",
						},
						"wlan": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(3).",
						},
						"gan": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(4).",
						},
						"hspa_evolution": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(5).",
						},
						"eutran": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(6).",
						},
						"virtual": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(7).",
						},
						"nb_iot": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "(8).",
						},
						"other_types_range": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "(9-255).",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "",
									},
									"types": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Other RAT Types. To specify other RAT ranges, add a hyphen between the lowest and the highest numbers, for example: 11-15. Multiple Ranges can be chosen when separated with comma.",
									},
								},
							},
						},
					},
				},
			},
			"restoration_and_recovery": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Restoration and Recovery (Relevant for V2 only).",
			},
			"reverse_service": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Accept PDUs from the GGSN/PGW to the SGSN/SGW on a previously established PDP context, even if different ports are used.",
			},
			"selection_mode": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Match by a selection mode.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The mode as integer. [0 - Verified, 1 - MS - Not verified, 2 - Network - Not verified].",
						},
					},
				},
			},
			"trace_management": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Trace Management (Relevant for V2 only).",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
func dataSourceManagementServiceGtpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceGtpRes, err := client.ApiCall("show-service-gtp", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceGtpRes.Success {
		return fmt.Errorf(showServiceGtpRes.ErrorMsg)
	}

	serviceGtp := showServiceGtpRes.GetData()

	log.Println("Read Service Gtp - Show JSON = ", serviceGtp)

	if v := serviceGtp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceGtp["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := serviceGtp["version"]; v != nil {
		_ = d.Set("version", v)
	}
	if v := serviceGtp["access-point-name"]; v != nil {

		accessPointMap := make(map[string]interface{})

		payload := v.(map[string]interface{})

		if v := payload["enable"]; v != nil {
			accessPointMap["enable"] = strconv.FormatBool(v.(bool))
		}
		if v := payload["apn"]; v != nil {
			innerMap := v.(map[string]interface{})

			accessPointMap["apn"] = innerMap["name"].(string)

		}
		log.Println("map is ", accessPointMap)
		_ = d.Set("access_point_name", accessPointMap)

	} else {
		_ = d.Set("access_point_name", nil)
	}
	if v := serviceGtp["allow-usage-of-static-ip"]; v != nil {
		d.Set("allow_usage_of_static_ip", v)
	}

	if v := serviceGtp["apply-access-policy-on-user-traffic"]; v != nil {

		payload := v.(map[string]interface{})

		res := make(map[string]interface{})

		if v := payload["enable"]; v != nil {
			res["enable"] = strconv.FormatBool(v.(bool))
		}
		if v := payload["add-imsi-field-to-log"]; v != nil {
			res["add_imsi_field_to_log"] = strconv.FormatBool(v.(bool))
		}
		d.Set("apply_access_policy_on_user_traffic", res)
	} else {
		d.Set("apply_access_policy_on_user_traffic", nil)
	}

	if v := serviceGtp["cs-fallback-and-srvcc"]; v != nil {
		d.Set("cs_fallback_and_srvcc", v)
	}

	if v := serviceGtp["imsi-prefix"]; v != nil {
		payload := v.(map[string]interface{})

		res := make(map[string]interface{})

		if v := payload["enable"]; v != nil {
			res["enable"] = strconv.FormatBool(v.(bool))
		}
		if v := payload["prefix"]; v != nil {
			res["prefix"] = v
		}

		d.Set("imsi_prefix", res)
	} else {
		d.Set("imsi_prefix", nil)
	}

	if v := serviceGtp["interface-profile"]; v != nil {
		payload := v.(map[string]interface{})

		res := make(map[string]interface{})

		if v := payload["profile"]; v != nil {
			profileMap := v.(map[string]interface{})
			if j, _ := profileMap["name"]; j != nil {
				res["profile"] = j
			}
		}
		if v := payload["custom-message-types"]; v != nil {
			res["custom_message_types"] = v
		}

		d.Set("interface_profile", res)
	} else {
		d.Set("interface_profile", nil)
	}

	if serviceGtp["ldap-group"] != nil {

		ldapGroupMap := serviceGtp["ldap-group"].(map[string]interface{})

		ldapGroupMapToReturn := make(map[string]interface{})

		if v, _ := ldapGroupMap["enable"]; v != nil {
			ldapGroupMapToReturn["enable"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := ldapGroupMap["group"]; v != nil {
			groupMap := v.(map[string]interface{})
			if j, _ := groupMap["name"]; j != nil {
				ldapGroupMapToReturn["group"] = j.(string)
			}
		}
		if v, _ := ldapGroupMap["according-to"]; v != nil {
			ldapGroupMapToReturn["according_to"] = v
		}
		_ = d.Set("ldap_group", ldapGroupMapToReturn)
	} else {
		_ = d.Set("ldap_group", nil)
	}

	if serviceGtp["ms-isdn"] != nil {

		msIsdnMap := serviceGtp["ms-isdn"].(map[string]interface{})

		msIsdnMapToReturn := make(map[string]interface{})

		if v, _ := msIsdnMap["enable"]; v != nil {
			msIsdnMapToReturn["enable"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := msIsdnMap["ms-isdn"]; v != nil {
			msIsdnMapToReturn["ms_isdn"] = v
		}
		_ = d.Set("ms_isdn", msIsdnMapToReturn)
	} else {
		_ = d.Set("ms_isdn", nil)
	}

	if serviceGtp["radio-access-technology"] != nil {

		radioAccessTechnologyMap, ok := serviceGtp["radio-access-technology"].(map[string]interface{})

		if ok {
			radioAccessTechnologyMapToReturn := make(map[string]interface{})

			if v := radioAccessTechnologyMap["utran"]; v != nil {
				radioAccessTechnologyMapToReturn["utran"] = v
			}
			if v := radioAccessTechnologyMap["geran"]; v != nil {
				radioAccessTechnologyMapToReturn["geran"] = v
			}
			if v := radioAccessTechnologyMap["wlan"]; v != nil {
				radioAccessTechnologyMapToReturn["wlan"] = v
			}
			if v := radioAccessTechnologyMap["gan"]; v != nil {
				radioAccessTechnologyMapToReturn["gan"] = v
			}
			if v := radioAccessTechnologyMap["hspa-evolution"]; v != nil {
				radioAccessTechnologyMapToReturn["hspa_evolution"] = v
			}
			if v := radioAccessTechnologyMap["eutran"]; v != nil {
				radioAccessTechnologyMapToReturn["eutran"] = v
			}
			if v := radioAccessTechnologyMap["virtual"]; v != nil {
				radioAccessTechnologyMapToReturn["virtual"] = v
			}
			if v := radioAccessTechnologyMap["nb-iot"]; v != nil {
				radioAccessTechnologyMapToReturn["nb_iot"] = v
			}
			if v, ok := radioAccessTechnologyMap["other-types-range"]; ok {

				otherTypesRangeMap, ok := v.(map[string]interface{})
				if ok {
					otherTypesRangeMapToReturn := make(map[string]interface{})

					if v, _ := otherTypesRangeMap["enable"]; v != nil {
						otherTypesRangeMapToReturn["enable"] = v
					}
					if v, _ := otherTypesRangeMap["types"]; v != nil {
						otherTypesRangeMapToReturn["types"] = v
					}
					radioAccessTechnologyMapToReturn["other_types_range"] = []interface{}{otherTypesRangeMapToReturn}
				}
			}
			_ = d.Set("radio_access_technology", []interface{}{radioAccessTechnologyMapToReturn})

		}
	} else {
		_ = d.Set("radio_access_technology", nil)
	}

	if v := serviceGtp["restoration-and-recovery"]; v != nil {
		_ = d.Set("restoration_and_recovery", v)
	}

	if v := serviceGtp["reverse-service"]; v != nil {
		_ = d.Set("reverse_service", v)
	}

	if serviceGtp["selection-mode"] != nil {

		selectionModeMapToReturn := make(map[string]interface{})

		innerMap := serviceGtp["selection-mode"].(map[string]interface{})

		if v, _ := innerMap["mode"]; v != nil {

			selectionModeMapToReturn["mode"] = v
		}

		if v, _ := innerMap["enable"]; v != nil {
			selectionModeMapToReturn["enable"] = v
		}

		_ = d.Set("selection_mode", []interface{}{selectionModeMapToReturn})
	} else {
		_ = d.Set("selection_mode", nil)
	}

	if v := serviceGtp["trace-management"]; v != nil {
		_ = d.Set("trace_management", v)
	}

	if serviceGtp["tags"] != nil {
		tagsJson, ok := serviceGtp["tags"].([]interface{})
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

	if v := serviceGtp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceGtp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if serviceGtp["groups"] != nil {
		groupsJson, ok := serviceGtp["groups"].([]interface{})
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

	if v := serviceGtp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	return nil
}
