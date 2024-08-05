package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func dataSourceManagementAccessRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessRuleRead,

		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
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
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
			},
			"action_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Action settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_identity_captive_portal": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
						},
						"limit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"content": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of processed file types that this rule applies on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"content_direction": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "On which direction the file types processing is applied.",
			},
			"content_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for data.",
			},
			"custom_fields": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Custom fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_1": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "First custom field.",
						},
						"field_2": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Second custom field.",
						},
						"field_3": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Third custom field.",
						},
					},
				},
			},
			"destination": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for destination.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"inline_layer": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for service.",
			},
			"source": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for source.",
			},
			"time": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"track": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Track Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accounting": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Turns accounting for track on and off.",
						},
						"alert": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of alert for the track.",
						},
						"enable_firewall_session": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determine whether to generate session log to firewall only connections.",
						},
						"per_connection": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether to perform the log per connection.",
						},
						"per_session": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether to perform the log per session.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "\"Log\", \"Extended Log\", \"Detailed Log\", \"None\".",
						},
					},
				},
			},
			"user_check": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "User check settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"confirm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"custom_frequency": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"every": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "N/A",
									},
									"unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"frequency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"interaction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Valid values \"Any\", \"All_GwToGw\" or VPN community name",
			},
			"vpn_communities": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "VPN communities (used for multiple VPNs, otherwise, use \"vpn\" field)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn_directional": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VPN directional",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "From VPN community",
						},
						"to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "To VPN community",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"fields_with_uid_identifier": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of resource fields that will use object UIDs as object identifiers. Default is object name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementAccessRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"layer": d.Get("layer"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAccessRuleRes, err := client.ApiCall("show-access-rule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessRuleRes.Success {
		return fmt.Errorf(showAccessRuleRes.ErrorMsg)
	}

	accessRule := showAccessRuleRes.GetData()

	log.Println("Read Access Rule - Show JSON = ", accessRule)

	if v := accessRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
	if v := accessRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessRule["action"]; v != nil {
		actionId := resolveObjectIdentifier("action", accessRule["action"], d)
		if actionId == "Inner Layer" {
			actionId = "Apply Layer"
		}
		_ = d.Set("action", actionId)
	}

	if accessRule["action-settings"] != nil {

		actionSettingsMap := accessRule["action-settings"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v, _ := actionSettingsMap["enable-identity-captive-portal"]; v != nil {
			actionSettingsMapToReturn["enable_identity_captive_portal"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := actionSettingsMap["limit"]; v != nil {
			actionSettingsMapToReturn["limit"] = v
		}

		_ = d.Set("action_settings", actionSettingsMapToReturn)
	} else {
		_ = d.Set("action_settings", nil)
	}

	if accessRule["content"] != nil {
		contentIds := resolveListOfIdentifiers("content", accessRule["content"], d)
		_ = d.Set("content", contentIds)
	} else {
		_ = d.Set("content", nil)
	}

	if v := accessRule["content-direction"]; v != nil {
		_ = d.Set("content_direction", v)
	}

	if v := accessRule["content-negate"]; v != nil {
		_ = d.Set("content_negate", v)
	}

	if accessRule["custom-fields"] != nil {

		customFieldsMap := accessRule["custom-fields"].(map[string]interface{})

		customFieldsMapToReturn := make(map[string]interface{})

		if v, _ := customFieldsMap["field-1"]; v != nil {
			customFieldsMapToReturn["field_1"] = v
		}

		if v, _ := customFieldsMap["field-2"]; v != nil {
			customFieldsMapToReturn["field_2"] = v
		}

		if v, _ := customFieldsMap["field-3"]; v != nil {
			customFieldsMapToReturn["field_3"] = v
		}
		_ = d.Set("custom_fields", customFieldsMapToReturn)
	} else {
		_ = d.Set("custom_fields", nil)
	}

	if accessRule["destination"] != nil {
		destinationIds := resolveListOfIdentifiers("destination", accessRule["destination"], d)
		_ = d.Set("destination", destinationIds)
	}

	if v := accessRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if v := accessRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := accessRule["inline-layer"]; v != nil {
		_ = d.Set("inline_layer", v)
	}

	if accessRule["install-on"] != nil {
		installOnIds := resolveListOfIdentifiers("install-on", accessRule["install-on"], d)
		_ = d.Set("install_on", installOnIds)
	}

	if accessRule["service"] != nil {
		serviceIds := resolveListOfIdentifiers("service", accessRule["service"], d)
		_ = d.Set("service", serviceIds)
	}

	if v := accessRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if accessRule["source"] != nil {
		sourceIds := resolveListOfIdentifiers("source", accessRule["source"], d)
		_ = d.Set("source", sourceIds)
	}

	if v := accessRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if accessRule["time"] != nil {
		timeIds := resolveListOfIdentifiers("time", accessRule["time"], d)
		_ = d.Set("time", timeIds)
	}
	if accessRule["track"] != nil {

		trackMap := accessRule["track"].(map[string]interface{})

		trackMapToReturn := make(map[string]interface{})
		if v := trackMap["accounting"]; v != nil {
			trackMapToReturn["accounting"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["alert"]; v != nil {
			trackMapToReturn["alert"] = v.(string)
		}

		if v := trackMap["enable-firewall-session"]; v != nil {
			trackMapToReturn["enable_firewall_session"] = strconv.FormatBool(v.(bool))
		}

		if v := trackMap["per-connection"]; v != nil {
			trackMapToReturn["per_connection"] = strconv.FormatBool(v.(bool))
		}

		if v := trackMap["per-session"]; v != nil {
			trackMapToReturn["per_session"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["type"]; v != nil {
			trackMapToReturn["type"] = v.(map[string]interface{})["name"].(string)
		}
		err = d.Set("track", trackMapToReturn)

	} else {
		_ = d.Set("track", nil)
	}

	if accessRule["user-check"] != nil {

		userCheckMap := accessRule["user-check"].(map[string]interface{})

		userCheckMapToReturn := make(map[string]interface{})

		if v, _ := userCheckMap["confirm"]; v != nil {
			userCheckMapToReturn["confirm"] = v
		}

		if v, ok := userCheckMap["custom-frequency"]; ok {

			userCheckConfigMap := v.(map[string]interface{})
			userCheckConfigMapToReturn := make(map[string]interface{})

			if v, _ := userCheckConfigMap["every"]; v != nil {
				userCheckConfigMapToReturn["every"] = v
			}

			if v, _ := userCheckConfigMap["unit"]; v != nil {
				userCheckConfigMapToReturn["unit"] = v
			}
			userCheckMapToReturn["custom_frequency"] = []interface{}{userCheckConfigMapToReturn}
		}

		if v, _ := userCheckMap["frequency"]; v != nil {
			userCheckMapToReturn["frequency"] = v
		}

		if v, _ := userCheckMap["interaction"]; v != nil {
			userCheckMapToReturn["interaction"] = v.(map[string]interface{})["name"]
		}

		_ = d.Set("user_check", []interface{}{userCheckMapToReturn})
	} else {
		_ = d.Set("user_check", nil)
	}

	if v := accessRule["vpn"]; v != nil {
		vpnList := v.([]interface{})
		if len(vpnList) > 0 {
			vpnType := vpnList[0].(map[string]interface{})["type"].(string)
			if len(vpnList) == 1 && vpnType != "VpnDirectionalElement" { // BC
				vpnId := resolveObjectIdentifier("vpn", v.([]interface{})[0], d)
				_ = d.Set("vpn", vpnId)
				_ = d.Set("vpn_communities", nil)
				_ = d.Set("vpn_directional", nil)
			} else if vpnType != "VpnDirectionalElement" {
				vpnIds := resolveListOfIdentifiers("vpn", vpnList, d)
				_ = d.Set("vpn_communities", vpnIds)
				_ = d.Set("vpn", nil)
				_ = d.Set("vpn_directional", nil)
			} else if vpnType == "VpnDirectionalElement" {
				var vpnDirectionalListState []map[string]interface{}
				for i := range vpnList {
					vpnDirectionalObj := vpnList[i].(map[string]interface{})
					if v, _ := vpnDirectionalObj["name"]; v != nil {
						vpnDirectionalNames := strings.Split(v.(string), "->")
						vpnDirectionalState := make(map[string]interface{})
						vpnDirectionalState["from"] = vpnDirectionalNames[0]
						vpnDirectionalState["to"] = vpnDirectionalNames[1]
						vpnDirectionalListState = append(vpnDirectionalListState, vpnDirectionalState)
					}
				}
				_ = d.Set("vpn_directional", vpnDirectionalListState)
				_ = d.Set("vpn_communities", nil)
				_ = d.Set("vpn", nil)
			} else {
				return fmt.Errorf("Cannot read invalid VPN type [" + vpnType + "]")
			}
		}
	}

	if v := accessRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	return nil
}
