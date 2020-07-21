package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
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
				Description: "Communities or Directional.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
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

	showAccessRuleRes, err := client.ApiCall("show-access-rule", payload, client.GetSessionID(), true, false)
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
		_ = d.Set("action", v.(map[string]interface{})["name"])
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
		contentJson := accessRule["content"].([]interface{})
		contentIds := make([]string, 0)
		if len(contentJson) > 0 {
			for _, content := range contentJson {
				content := content.(map[string]interface{})
				contentIds = append(contentIds, content["name"].(string))
			}
		}
		_, contentInConf := d.GetOk("content")
		if contentIds[0] == "Any" && !contentInConf {
			_ = d.Set("content", []interface{}{})
		} else {
			_ = d.Set("content", contentIds)
		}
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

		_, customFieldsInConf := d.GetOk("custom_fields")
		defaultCustomField := map[string]interface{}{"field_1": "", "field_2": "", "field_3": ""}
		if reflect.DeepEqual(defaultCustomField, customFieldsMapToReturn) && !customFieldsInConf {
			_ = d.Set("custom_fields", map[string]interface{}{})
		} else {
			_ = d.Set("custom_fields", customFieldsMapToReturn)
		}
	} else {
		_ = d.Set("custom_fields", nil)
	}

	if accessRule["destination"] != nil {
		destinationJson := accessRule["destination"].([]interface{})
		destinationIds := make([]string, 0)
		if len(destinationJson) > 0 {
			for _, destination := range destinationJson {
				destination := destination.(map[string]interface{})
				destinationIds = append(destinationIds, destination["name"].(string))
			}
		}
		_, destinationInConf := d.GetOk("destination")
		if destinationIds[0] == "Any" && !destinationInConf {
			_ = d.Set("destination", []interface{}{})
		} else {
			_ = d.Set("destination", destinationIds)
		}
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
		installOnJson := accessRule["install-on"].([]interface{})
		installOnJsonIds := make([]string, 0)
		if len(installOnJson) > 0 {
			for _, installOn := range installOnJson {
				installOn := installOn.(map[string]interface{})
				installOnJsonIds = append(installOnJsonIds, installOn["name"].(string))
			}
		}
		_, installOnInConf := d.GetOk("install_on")
		if installOnJsonIds[0] == "Policy Targets" && !installOnInConf {
			_ = d.Set("install_on", []interface{}{})
		} else {
			_ = d.Set("install_on", installOnJsonIds)
		}
	}

	if accessRule["service"] != nil {
		serviceJson := accessRule["service"].([]interface{})
		serviceJsonIds := make([]string, 0)
		if len(serviceJson) > 0 {
			for _, service := range serviceJson {
				service := service.(map[string]interface{})
				serviceJsonIds = append(serviceJsonIds, service["name"].(string))
			}
		}
		_, serviceInConf := d.GetOk("service")
		if serviceJsonIds[0] == "Any" && !serviceInConf {
			_ = d.Set("service", []interface{}{})
		} else {
			_ = d.Set("service", serviceJsonIds)
		}
	}

	if v := accessRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if accessRule["source"] != nil {
		sourceJson := accessRule["source"].([]interface{})
		sourceIds := make([]string, 0)
		if len(sourceJson) > 0 {
			for _, source := range sourceJson {
				source := source.(map[string]interface{})
				sourceIds = append(sourceIds, source["name"].(string))
			}
		}
		_, sourceInConf := d.GetOk("source")
		if sourceIds[0] == "Any" && !sourceInConf {
			_ = d.Set("source", []interface{}{})
		} else {
			_ = d.Set("source", sourceIds)
		}
	}

	if v := accessRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if accessRule["time"] != nil {
		timeJson := accessRule["time"].([]interface{})
		timeIds := make([]string, 0)
		if len(timeJson) > 0 {
			for _, time := range timeJson {
				time := time.(map[string]interface{})
				timeIds = append(timeIds, time["name"].(string))
			}
		}
		_, timeInConf := d.GetOk("time")
		if timeIds[0] == "Any" && !timeInConf {
			_ = d.Set("time", []interface{}{})
		} else {
			_ = d.Set("time", timeIds)
		}
	}

	if accessRule["track"] != nil {

		trackMap := accessRule["track"].(map[string]interface{})

		trackMapToReturn := make(map[string]interface{})

		if v, _ := trackMap["accounting"]; v != nil {
			trackMapToReturn["accounting"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["alert"]; v != nil {
			trackMapToReturn["alert"] = v
		}

		if v, _ := trackMap["enable-firewall-session"]; v != nil {
			trackMapToReturn["enable_firewall_session"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["per-connection"]; v != nil {
			trackMapToReturn["per_connection"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["per-session"]; v != nil {
			trackMapToReturn["per_session"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["type"]; v != nil {
			trackMapToReturn["type"] = v.(map[string]interface{})["name"]
		}

		_, trackInConf := d.GetOk("track")
		defaultTrack := map[string]interface{}{
			"accounting":              "false",
			"alert":                   "none",
			"enable_firewall_session": "false",
			"per_connection":          "false",
			"per_session":             "false",
			"type":                    "None"}

		if reflect.DeepEqual(defaultTrack, trackMapToReturn) && !trackInConf {
			_ = d.Set("track", map[string]interface{}{})
		} else {
			_ = d.Set("track", trackMapToReturn)
		}
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
		_ = d.Set("vpn", v.([]interface{})[0].(map[string]interface{})["name"])
	}

	if v := accessRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
