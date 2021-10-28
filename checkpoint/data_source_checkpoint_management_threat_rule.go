package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func dataSourceManagementThreatRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatRuleRead,
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action-the enforced profile.",
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
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of objects defining Protected Scope identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for Protected Scope.",
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
			"track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Packet tracking.",
			},
			"track_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Threat rule track settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_capture": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Packet capture.",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
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
		},
	}
}

func dataSourceManagementThreatRuleRead(d *schema.ResourceData, m interface{}) error {

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

	showThreatRuleRes, err := client.ApiCall("show-threat-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatRuleRes.Success {
		return fmt.Errorf(showThreatRuleRes.ErrorMsg)
	}

	threatRule := showThreatRuleRes.GetData()

	log.Println("Read Threat Rule - Show JSON = ", threatRule)

	if v := threatRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
