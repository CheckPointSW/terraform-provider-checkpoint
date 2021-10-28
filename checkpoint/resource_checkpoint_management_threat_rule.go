package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strings"
)

func resourceManagementThreatRule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatRule,
		Read:   readManagementThreatRule,
		Update: updateManagementThreatRule,
		Delete: deleteManagementThreatRule,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				arr := strings.Split(d.Id(), ";")
				if len(arr) != 2 {
					return nil, fmt.Errorf("invalid unique identifier format. UID format: <LAYER_NAME>;<RULE_UID>")
				}
				_ = d.Set("layer", arr[0])
				d.SetId(arr[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
			"position": {
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
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Action-the enforced profile.",
				Default:     "Optimized",
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
			"install_on": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of objects defining Protected Scope identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for Protected Scope.",
				Default:     false,
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
			"track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Packet tracking.",
				Default:     "Log",
			},
			"track_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Threat rule track settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_capture": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Packet capture.",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"exceptions": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of rule's exceptions identified by UID",
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

func createManagementThreatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	threatRule := make(map[string]interface{})

	if v, ok := d.GetOk("layer"); ok {
		threatRule["layer"] = v.(string)
	}
	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				threatRule["position"] = "top" // entire rule-base
			} else {
				threatRule["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			threatRule["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			threatRule["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				threatRule["position"] = "bottom" // entire rule-base
			} else {
				threatRule["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	if v, ok := d.GetOk("name"); ok {
		threatRule["name"] = v.(string)
	}

	if v, ok := d.GetOk("action"); ok {
		threatRule["action"] = v.(string)
	}

	if v, ok := d.GetOk("destination"); ok {
		threatRule["destination"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("destination_negate"); ok {
		threatRule["destination-negate"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		threatRule["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("install_on"); ok {
		threatRule["install-on"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("protected_scope"); ok {
		threatRule["protected-scope"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("protected_scope_negate"); ok {
		threatRule["protected-scope-negate"] = v.(bool)
	}

	if val, ok := d.GetOk("service"); ok {
		threatRule["service"] = val.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("service_negate"); ok {
		threatRule["service-negate"] = v.(bool)
	}

	if v, ok := d.GetOk("source"); ok {
		threatRule["source"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("source_negate"); ok {
		threatRule["source-negate"] = v.(bool)
	}

	if v, ok := d.GetOk("track"); ok {
		threatRule["track"] = v.(string)
	}

	if _, ok := d.GetOk("track_settings"); ok {
		trackSettings := make(map[string]interface{})
		if v, ok := d.GetOkExists("track_settings.packet_capture"); ok {
			trackSettings["packet-capture"] = v.(bool)
		}
		threatRule["track-settings"] = trackSettings
	}

	if v, ok := d.GetOk("comments"); ok {
		threatRule["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatRule["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatRule["ignore-warnings"] = v.(bool)
	}

	log.Println("Create Threat Rule - Map = ", threatRule)

	addThreatRuleRes, err := client.ApiCall("add-threat-rule", threatRule, client.GetSessionID(), true, false)
	if err != nil || !addThreatRuleRes.Success {
		if addThreatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addThreatRuleRes.GetData()["uid"].(string))

	return readManagementThreatRule(d, m)
}

func readManagementThreatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	showThreatRuleRes, err := client.ApiCall("show-threat-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatRuleRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showThreatRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatRuleRes.ErrorMsg)
	}

	threatRule := showThreatRuleRes.GetData()

	log.Println("Read Threat Rule - Show JSON = ", threatRule)

	if v := threatRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if threatRule["source"] != nil {
		sourceJson := threatRule["source"].([]interface{})
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

	if v := threatRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if threatRule["destination"] != nil {
		destinationJson := threatRule["destination"].([]interface{})
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

	if v := threatRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if threatRule["protected-scope"] != nil {
		protectedScopeJson := threatRule["protected-scope"].([]interface{})
		protectedScopeIds := make([]string, 0)
		if len(protectedScopeJson) > 0 {
			for _, protectedScope := range protectedScopeJson {
				protectedScope := protectedScope.(map[string]interface{})
				protectedScopeIds = append(protectedScopeIds, protectedScope["name"].(string))
			}
		}
		_, protectedScopeInConf := d.GetOk("protected_scope")
		if protectedScopeIds[0] == "Any" && !protectedScopeInConf {
			_ = d.Set("protected_scope", []interface{}{})
		} else {
			_ = d.Set("protected_scope", protectedScopeIds)
		}
	}

	if v := threatRule["protected-scope-negate"]; v != nil {
		_ = d.Set("protected_scope_negate", v)
	}

	if threatRule["service"] != nil {
		serviceJson := threatRule["service"].([]interface{})
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

	if v := threatRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if threatRule["install-on"] != nil {
		installOnJson := threatRule["install-on"].([]interface{})
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

	if v := threatRule["action"]; v != nil {
		_ = d.Set("action", v.(map[string]interface{})["name"])
	}

	if v := threatRule["track"]; v != nil {
		_ = d.Set("track", v.(map[string]interface{})["name"])
	}

	if v := threatRule["track-settings"]; v != nil {
		trackSettingsMap := v.(map[string]interface{})
		trackSettingsState := make(map[string]interface{})
		if v := trackSettingsMap["packet-capture"]; v != nil {
			trackSettingsState["packet_capture"] = v.(bool)
		}

		_, trackSettingsInConf := d.GetOk("track_settings")
		defaultTrackSettings := map[string]interface{}{"packet-capture": true}
		if reflect.DeepEqual(defaultTrackSettings, trackSettingsState) && !trackSettingsInConf {
			_ = d.Set("track_settings", map[string]interface{}{})
		} else {
			_ = d.Set("track_settings", trackSettingsState)
		}
	}

	if threatRule["exceptions"] != nil {
		exceptionsJson := threatRule["exceptions"].([]interface{})
		exceptionsIds := make([]string, 0)
		if len(exceptionsJson) > 0 {
			for _, e := range exceptionsJson {
				e := e.(map[string]interface{})
				exceptionsIds = append(exceptionsIds, e["uid"].(string))
			}
		}
		_ = d.Set("exceptions", exceptionsIds)
	}

	if v := threatRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementThreatRule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	threatRule := make(map[string]interface{})

	threatRule["uid"] = d.Id()
	threatRule["layer"] = d.Get("layer")

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					threatRule["new-position"] = "top" // entire rule-base
				} else {
					threatRule["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				threatRule["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				threatRule["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					threatRule["new-position"] = "bottom" // entire rule-base
				} else {
					threatRule["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	if d.HasChange("name") {
		threatRule["new-name"] = d.Get("name")
	}

	if d.HasChange("action") {
		threatRule["action"] = d.Get("action")
	}

	if d.HasChange("destination") {
		if v, ok := d.GetOk("destination"); ok {
			threatRule["destination"] = v.(*schema.Set).List()
		} else {
			oldDestination, _ := d.GetChange("destination")
			threatRule["destination"] = map[string]interface{}{"remove": oldDestination.(*schema.Set).List()}
		}
	}

	if d.HasChange("destination_negate") {
		threatRule["destination-negate"] = d.Get("destination_negate")
	}

	if d.HasChange("enabled") {
		threatRule["enabled"] = d.Get("enabled")
	}
	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			threatRule["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstallOn, _ := d.GetChange("install_on")
			threatRule["install-on"] = map[string]interface{}{"remove": oldInstallOn.(*schema.Set).List()}
		}
	}

	if d.HasChange("service") {
		if v, ok := d.GetOk("service"); ok {
			threatRule["service"] = v.(*schema.Set).List()
		} else {
			oldService, _ := d.GetChange("service")
			threatRule["service"] = map[string]interface{}{"remove": oldService.(*schema.Set).List()}
		}
	}

	if d.HasChange("service_negate") {
		threatRule["service-negate"] = d.Get("service_negate")
	}

	if d.HasChange("source") {
		if v, ok := d.GetOk("source"); ok {
			threatRule["source"] = v.(*schema.Set).List()
		} else {
			oldSource, _ := d.GetChange("source")
			threatRule["source"] = map[string]interface{}{"remove": oldSource.(*schema.Set).List()}
		}
	}

	if d.HasChange("source_negate") {
		threatRule["source-negate"] = d.Get("source_negate")
	}

	if d.HasChange("protected_scope") {
		if v, ok := d.GetOk("protected_scope"); ok {
			threatRule["protected-scope"] = v.(*schema.Set).List()
		} else {
			oldProtectedSrc, _ := d.GetChange("protected_scope")
			threatRule["protected-scope"] = map[string]interface{}{"remove": oldProtectedSrc.(*schema.Set).List()}
		}
	}

	if d.HasChange("protected_scope_negate") {
		threatRule["protected-scope-negate"] = d.Get("protected_scope_negate")
	}

	if d.HasChange("track") {
		threatRule["track"] = d.Get("track")
	}

	if d.HasChange("track_settings") {
		if v, ok := d.GetOkExists("track_settings.packet_capture"); ok {
			threatRule["track-settings"] = map[string]interface{}{"packet-capture": v}
		}
	}

	if d.HasChange("comments") {
		threatRule["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatRule["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatRule["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Threat Rule - Map = ", threatRule)

	updateThreatRuleRes, err := client.ApiCall("set-threat-rule", threatRule, client.GetSessionID(), true, false)
	if err != nil || !updateThreatRuleRes.Success {
		if updateThreatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(updateThreatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	return readManagementThreatRule(d, m)
}

func deleteManagementThreatRule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatRulePayload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	deleteThreatRuleRes, err := client.ApiCall("delete-threat-rule", threatRulePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteThreatRuleRes.Success {
		if deleteThreatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")
	return nil
}
