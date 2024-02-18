package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func resourceManagementAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAccessRule,
		Read:   readManagementAccessRule,
		Update: updateManagementAccessRule,
		Delete: deleteManagementAccessRule,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				arr := strings.Split(d.Id(), ";")
				if len(arr) != 2 {
					return nil, fmt.Errorf("invalid unique identifier format. UID format: <LAYER_IDENTIFIER>;<RULE_UID>")
				}
				_ = d.Set("layer", arr[0])
				d.SetId(arr[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"layer": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
			"position": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Position in the rulebase.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule on top of specific section identified by uid or name. Select value 'top' for entire rule base.",
						},
						"above": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule above specific section/rule identified by uid or name.",
						},
						"below": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule below specific section/rule identified by uid or name.",
						},
						"bottom": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule in the bottom of specific section identified by uid or name. Select value 'bottom' for entire rule base.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"action": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
				Default:     "Drop",
			},
			"action_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Action settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_identity_captive_portal": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
						},
						"limit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"content": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of processed file types that this rule applies on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"content_direction": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "On which direction the file types processing is applied.",
				Default:     "any",
			},
			"content_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for data.",
				Default:     false,
			},
			"custom_fields": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Custom fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "First custom field.",
						},
						"field_2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Second custom field.",
						},
						"field_3": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Third custom field.",
						},
					},
				},
			},
			"destination": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for destination.",
				Default:     false,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/Disable the rule.",
				Default:     true,
			},
			"inline_layer": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for service.",
				Default:     false,
			},
			"source": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for source.",
				Default:     false,
			},
			"time": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"track": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Track Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accounting": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Turns accounting for track on and off.",
						},
						"alert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of alert for the track.",
						},
						"enable_firewall_session": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determine whether to generate session log to firewall only connections.",
						},
						"per_connection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines whether to perform the log per connection.",
						},
						"per_session": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines whether to perform the log per session.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "\"Log\", \"Extended Log\", \"Detailed Log\", \"None\".",
						},
					},
				},
			},
			"user_check": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "User check settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"confirm": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"custom_frequency": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"every": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "N/A",
									},
									"unit": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
								},
							},
						},
						"frequency": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"interaction": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"vpn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Communities or Directional.",
				Default:     "Any",
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
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
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

func createManagementAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	accessRule := make(map[string]interface{})

	if v, ok := d.GetOk("layer"); ok {
		accessRule["layer"] = v.(string)
	}
	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				accessRule["position"] = "top" // entire rule-base
			} else {
				accessRule["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			accessRule["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			accessRule["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				accessRule["position"] = "bottom" // entire rule-base
			} else {
				accessRule["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	if v, ok := d.GetOk("name"); ok {
		accessRule["name"] = v.(string)
	}
	if v, ok := d.GetOk("action"); ok {
		accessRule["action"] = v.(string)
	}
	if _, ok := d.GetOk("action_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("action_settings.enable_identity_captive_portal"); ok {
			res["enable-identity-captive-portal"] = v
		}
		if v, ok := d.GetOk("action_settings.limit"); ok {
			res["limit"] = v.(string)
		}
		accessRule["action-settings"] = res
	}
	if v, ok := d.GetOk("content"); ok {
		accessRule["content"] = v.(*schema.Set).List()
	}
	if v, ok := d.GetOk("content_direction"); ok {
		accessRule["content-direction"] = v.(string)
	}
	if v, ok := d.GetOk("content_negate"); ok {
		accessRule["content-negate"] = v.(bool)
	}
	if _, ok := d.GetOk("custom_fields"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("custom_fields.field_1"); ok {
			res["field-1"] = v.(string)
		}
		if v, ok := d.GetOk("custom_fields.field_2"); ok {
			res["field-2"] = v.(string)
		}
		if v, ok := d.GetOk("custom_fields.field_3"); ok {
			res["field-3"] = v.(string)
		}
		accessRule["custom-fields"] = res
	}
	if val, ok := d.GetOk("destination"); ok {
		accessRule["destination"] = val.(*schema.Set).List()
	}
	if v, ok := d.GetOk("destination_negate"); ok {
		accessRule["destination-negate"] = v.(bool)
	}
	if v, ok := d.GetOk("enabled"); ok {
		accessRule["enabled"] = v.(bool)
	}
	if val, ok := d.GetOk("inline_layer"); ok {
		accessRule["inline-layer"] = val.(string)
	}
	if val, ok := d.GetOk("install_on"); ok {
		accessRule["install-on"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("service"); ok {
		accessRule["service"] = val.(*schema.Set).List()
	}
	if v, ok := d.GetOk("service_negate"); ok {
		accessRule["service-negate"] = v.(bool)
	}
	if val, ok := d.GetOk("source"); ok {
		accessRule["source"] = val.(*schema.Set).List()
	}
	if v, ok := d.GetOk("source_negate"); ok {
		accessRule["source-negate"] = v.(bool)
	}
	if v, ok := d.GetOk("time"); ok {
		accessRule["time"] = v.(*schema.Set).List()
	}
	if _, ok := d.GetOk("track"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("track.accounting"); ok {
			res["accounting"] = v
		}
		if v, ok := d.GetOk("track.alert"); ok {
			res["alert"] = v.(string)
		}
		if v, ok := d.GetOk("track.enable_firewall_session"); ok {
			res["enable-firewall-session"] = v
		}
		if v, ok := d.GetOk("track.per_connection"); ok {
			res["per-connection"] = v
		}
		if v, ok := d.GetOk("track.per_session"); ok {
			res["per-session"] = v
		}
		if v, ok := d.GetOk("track.type"); ok {
			res["type"] = v.(string)
		}

		accessRule["track"] = res
	}

	if v, ok := d.GetOk("user_check"); ok {

		userCheckList := v.([]interface{})

		if len(userCheckList) > 0 {

			userCheckPayload := make(map[string]interface{})

			if v, ok := d.GetOk("user_check.0.confirm"); ok {
				userCheckPayload["confirm"] = v.(string)
			}
			if _, ok := d.GetOk("user_check.0.custom_frequency"); ok {

				customFrequencyPayLoad := make(map[string]interface{})

				if v, ok := d.GetOk("user_check.0.custom_frequency.0.every"); ok {
					customFrequencyPayLoad["every"] = v
				}
				if v, ok := d.GetOk("user_check.0.custom_frequency.0.unit"); ok {
					customFrequencyPayLoad["unit"] = v.(string)
				}
				userCheckPayload["custom-frequency"] = customFrequencyPayLoad
			}
			if v, ok := d.GetOk("user_check.0.frequency"); ok {
				userCheckPayload["frequency"] = v.(string)
			}
			if v, ok := d.GetOk("user_check.0.interaction"); ok {
				userCheckPayload["interaction"] = v.(string)
			}
			accessRule["user-check"] = userCheckPayload
		}
	}
	if v, ok := d.GetOk("vpn"); ok {
		accessRule["vpn"] = v.(string)
	}
	if val, ok := d.GetOk("comments"); ok {
		accessRule["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		accessRule["color"] = val.(string)
	}
	if val, ok := d.GetOk("ignore_errors"); ok {
		accessRule["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_warnings"); ok {
		accessRule["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Access Rule - Map = ", accessRule)

	addAccessRuleRes, err := client.ApiCall("add-access-rule", accessRule, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addAccessRuleRes.Success {
		if addAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addAccessRuleRes.GetData()["uid"].(string))

	return readManagementAccessRule(d, m)
}

func readManagementAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	showAccessRuleRes, err := client.ApiCall("show-access-rule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessRuleRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showAccessRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAccessRuleRes.ErrorMsg)
	}

	accessRule := showAccessRuleRes.GetData()

	log.Println("Read Access Rule - Show JSON = ", accessRule)

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
		defaultTrack := map[string]interface{}{
			"accounting":              "false",
			"alert":                   "none",
			"enable-firewall-session": "false",
			"per-connection":          "false",
			"per-session":             "false",
			"type":                    "None"}
		if v := trackMap["accounting"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "track.accounting", defaultTrack["accounting"].(string)) {
			trackMapToReturn["accounting"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["alert"]; v != nil && isArgDefault(v.(string), d, "track.alert", defaultTrack["alert"].(string)) {
			trackMapToReturn["alert"] = v.(string)
		}

		if v := trackMap["enable-firewall-session"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "track.enable_firewall_session", defaultTrack["enable-firewall-session"].(string)) {
			trackMapToReturn["enable_firewall_session"] = strconv.FormatBool(v.(bool))
		}

		if v := trackMap["per-connection"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "track.per_connection", defaultTrack["per-connection"].(string)) {
			trackMapToReturn["per_connection"] = strconv.FormatBool(v.(bool))
		}

		if v := trackMap["per-session"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "track.per_session", defaultTrack["per-session"].(string)) {
			trackMapToReturn["per_session"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := trackMap["type"]; v != nil && isArgDefault(v.(map[string]interface{})["name"].(string), d, "track.type", defaultTrack["type"].(string)) {
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
		vpnId := resolveObjectIdentifier("vpn", v.([]interface{})[0], d)
		_ = d.Set("vpn", vpnId)
	}

	if v := accessRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	return nil
}

func updateManagementAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	accessRule := make(map[string]interface{})

	accessRule["uid"] = d.Id()
	accessRule["layer"] = d.Get("layer")

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					accessRule["new-position"] = "top" // entire rule-base
				} else {
					accessRule["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				accessRule["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				accessRule["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					accessRule["new-position"] = "bottom" // entire rule-base
				} else {
					accessRule["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	if d.HasChange("name") {
		accessRule["new-name"] = d.Get("name")
	}

	if d.HasChange("action") {
		accessRule["action"] = d.Get("action")
		if val, ok := d.GetOk("inline_layer"); ok {
			accessRule["inline-layer"] = val.(string)
		}
	}

	if d.HasChange("action_settings") {

		if _, ok := d.GetOk("action_settings"); ok {

			res := make(map[string]interface{})

			if d.HasChange("action_settings.enable_identity_captive_portal") {
				if v, ok := d.GetOk("action_settings.enable_identity_captive_portal"); ok {
					res["enable-identity-captive-portal"] = v
				}
			}
			if d.HasChange("action_settings.limit") {
				if v, ok := d.GetOk("action_settings.limit"); ok {
					res["limit"] = v.(string)
				}
			}

			accessRule["action-settings"] = res
		} else { //argument deleted - go back to defaults
			accessRule["action-settings"] = map[string]interface{}{"enable-identity-captive-portal": "false"}
		}
	}

	if d.HasChange("content") {
		if v, ok := d.GetOk("content"); ok {
			accessRule["content"] = v.(*schema.Set).List()
		} else {
			oldContent, _ := d.GetChange("content")
			accessRule["content"] = map[string]interface{}{"remove": oldContent.(*schema.Set).List()}
		}
	}

	if d.HasChange("content_direction") {
		accessRule["content-direction"] = d.Get("content_direction")
	}

	if d.HasChange("content_negate") {
		accessRule["content-negate"] = d.Get("content_negate")
	}

	if d.HasChange("custom_fields") {

		if _, ok := d.GetOk("custom_fields"); ok {

			res := make(map[string]interface{})

			if d.HasChange("custom_fields.field_1") {
				res["field-1"] = d.Get("custom_fields.field_1")
			}
			if d.HasChange("custom_fields.field_2") {
				res["field-2"] = d.Get("custom_fields.field_2")
			}
			if d.HasChange("custom_fields.field_3") {
				res["field-3"] = d.Get("custom_fields.field_3")
			}
			accessRule["custom-fields"] = res
		} else {
			defaultCustomField := map[string]interface{}{"field-1": "", "field-2": "", "field-3": ""}
			accessRule["custom-fields"] = defaultCustomField
		}
	}

	if d.HasChange("destination") {
		if v, ok := d.GetOk("destination"); ok {
			accessRule["destination"] = v.(*schema.Set).List()
		} else {
			oldDestination, _ := d.GetChange("destination")
			accessRule["destination"] = map[string]interface{}{"remove": oldDestination.(*schema.Set).List()}
		}
	}

	if d.HasChange("destination_negate") {
		accessRule["destination-negate"] = d.Get("destination_negate")
	}

	if d.HasChange("enabled") {
		accessRule["enabled"] = d.Get("enabled")
	}

	if d.HasChange("inline_layer") {
		if v, ok := d.GetOk("inline_layer"); ok {
			accessRule["inline-layer"] = v.(string)
		}
	}

	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			accessRule["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstallOn, _ := d.GetChange("install_on")
			accessRule["install-on"] = map[string]interface{}{"remove": oldInstallOn.(*schema.Set).List()}
		}
	}

	if d.HasChange("service") {
		if v, ok := d.GetOk("service"); ok {
			accessRule["service"] = v.(*schema.Set).List()
		} else {
			oldService, _ := d.GetChange("service")
			accessRule["service"] = map[string]interface{}{"remove": oldService.(*schema.Set).List()}
		}
	}

	if d.HasChange("service_negate") {
		accessRule["service-negate"] = d.Get("service_negate")
	}

	if d.HasChange("source") {
		if v, ok := d.GetOk("source"); ok {
			accessRule["source"] = v.(*schema.Set).List()
		} else {
			oldSource, _ := d.GetChange("source")
			accessRule["source"] = map[string]interface{}{"remove": oldSource.(*schema.Set).List()}
		}
	}

	if d.HasChange("source_negate") {
		accessRule["source-negate"] = d.Get("source_negate")
	}

	if d.HasChange("time") {
		if v, ok := d.GetOk("time"); ok {
			accessRule["time"] = v.(*schema.Set).List()
		}
	} else {
		oldTime, _ := d.GetChange("time")
		oldTimeLst := oldTime.(*schema.Set).List()
		if len(oldTimeLst) > 0 {
			accessRule["time"] = map[string]interface{}{"remove": oldTimeLst}
		}
	}

	if d.HasChange("track") {
		defaultTrack := map[string]interface{}{
			"accounting":              "false",
			"alert":                   "none",
			"enable-firewall-session": "false",
			"per-connection":          "false",
			"per-session":             "false",
			"type":                    "None"}
		if v, ok := d.GetOk("track"); ok {
			res := make(map[string]interface{})
			logsSettingsJson := v.(map[string]interface{})
			if val, ok := logsSettingsJson["accounting"]; ok {
				res["accounting"] = val
			} else {
				res["accounting"] = defaultTrack["accounting"]
			}
			if val, ok := logsSettingsJson["alert"]; ok {
				res["alert"] = val
			} else {
				res["alert"] = defaultTrack["alert"]
			}
			if val, ok := logsSettingsJson["enable_firewall_session"]; ok {
				res["enable-firewall-session"] = val
			} else {
				res["enable-firewall-session"] = defaultTrack["enable-firewall-session"]
			}
			if val, ok := logsSettingsJson["per_connection"]; ok {
				res["per-connection"] = val
			} else {
				res["per-connection"] = defaultTrack["per-connection"]
			}
			if val, ok := logsSettingsJson["per_session"]; ok {
				res["per-session"] = val
			} else {
				res["per-session"] = defaultTrack["per-session"]
			}
			if val, ok := logsSettingsJson["type"]; ok {
				res["type"] = val
			} else {
				res["type"] = defaultTrack["type"]
			}
			accessRule["track"] = res

		} else {
			accessRule["track"] = defaultTrack
		}
	}

	if d.HasChange("user_check") {
		if v, ok := d.GetOk("user_check"); ok {
			userCheckList := v.([]interface{})
			if len(userCheckList) > 0 {
				userCheckPayload := make(map[string]interface{})

				if d.HasChange("user_check.0.confirm") {
					userCheckPayload["confirm"] = d.Get("user_check.0.confirm").(string)
				}
				if d.HasChange("user_check.0.custom_frequency") {

					customFrequencyConfigPayLoad := make(map[string]interface{})

					if d.HasChange("user_check.0.custom_frequency.0.every") {
						customFrequencyConfigPayLoad["every"] = d.Get("user_check.0.custom_frequency.0.every")
					}
					if d.HasChange("user_check.0.custom_frequency.0.unit") {
						customFrequencyConfigPayLoad["unit"] = d.Get("user_check.0.custom_frequency.0.unit").(string)
					}

					userCheckPayload["custom-frequency"] = customFrequencyConfigPayLoad
				}
				if d.HasChange("user_check.0.frequency") {
					userCheckPayload["frequency"] = d.Get("user_check.0.frequency").(string)
				}
				if d.HasChange("user_check.0.interaction") {
					userCheckPayload["interaction"] = d.Get("user_check.0.interaction").(string)
				}

				accessRule["user-check"] = userCheckPayload

			}
		}
	}

	if d.HasChange("vpn") {
		accessRule["vpn"] = d.Get("vpn")
	}

	if v, ok := d.GetOk("ignore_errors"); ok {
		accessRule["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOk("ignore_warnings"); ok {
		accessRule["ignore-warnings"] = v.(bool)
	}

	if d.HasChange("comments") {
		accessRule["comments"] = d.Get("comments")
	}

	log.Println("Update Access Rule - Map = ", accessRule)

	updateAccessRuleRes, err := client.ApiCall("set-access-rule", accessRule, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateAccessRuleRes.Success {
		if updateAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(updateAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	return readManagementAccessRule(d, m)
}

func deleteManagementAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	accessRulePayload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	deleteAccessRuleRes, err := client.ApiCall("delete-access-rule", accessRulePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAccessRuleRes.Success {
		if deleteAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
