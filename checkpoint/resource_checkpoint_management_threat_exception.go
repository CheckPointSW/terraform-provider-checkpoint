package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceManagementThreatException() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatException,
		Read:   readManagementThreatException,
		Update: updateManagementThreatException,
		Delete: deleteManagementThreatException,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				arr := strings.Split(d.Id(), ";")
				if len(arr) != 3 {
					return nil, fmt.Errorf("invalid unique identifier format. UID format: <LAYER_UID>;exception_group_uid or rule_uid;<EXCEPTION_GROUP_UID> or <PARENT_RULE_UID>;<RULE_UID>")
				}
				if !strings.EqualFold(arr[1], "exception_group_uid") && !strings.EqualFold(arr[1], "rule_uid") {
					return nil, fmt.Errorf("invalid unique identifier 2nd argument. Valid values: exception_group_uid or rule_uid")
				}
				_ = d.Set("layer", arr[0])
				_ = d.Set(arr[1], arr[2]) // Set exception_group_uid or rule_uid
				d.SetId(arr[3])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Rule name.",
			},
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
			"exception_group_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UID of the exception-group.",
			},
			"exception_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the exception-group.",
			},
			"rule_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UID of the parent rule.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the parent rule.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Action-the enforced profile.",
				Default:     "Detect",
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
			"protection_or_site": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Name of the protection or site.",
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
			"track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Packet tracking.",
				Default:     "Log",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner UID.",
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

func createManagementThreatException(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	threatException := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		threatException["name"] = v.(string)
	}

	if v, ok := d.GetOk("layer"); ok {
		threatException["layer"] = v.(string)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				threatException["position"] = "top" // entire rule-base
			} else {
				threatException["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			threatException["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			threatException["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				threatException["position"] = "bottom" // entire rule-base
			} else {
				threatException["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	if v, ok := d.GetOk("exception_group_uid"); ok {
		threatException["exception-group-uid"] = v.(string)
	}

	if v, ok := d.GetOk("exception_group_name"); ok {
		threatException["exception-group-name"] = v.(string)
	}

	if v, ok := d.GetOk("rule_uid"); ok {
		threatException["rule-uid"] = v.(string)
	}

	if v, ok := d.GetOk("rule_name"); ok {
		threatException["rule-name"] = v.(string)
	}

	if v, ok := d.GetOk("action"); ok {
		threatException["action"] = v.(string)
	}

	if v, ok := d.GetOk("destination"); ok {
		threatException["destination"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("destination_negate"); ok {
		threatException["destination-negate"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		threatException["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("install_on"); ok {
		threatException["install-on"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("protected_scope"); ok {
		threatException["protected-scope"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("protected_scope_negate"); ok {
		threatException["protected-scope-negate"] = v.(bool)
	}

	if v, ok := d.GetOk("protection_or_site"); ok {
		threatException["protection-or-site"] = v.(*schema.Set).List()
	}

	if val, ok := d.GetOk("service"); ok {
		threatException["service"] = val.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("service_negate"); ok {
		threatException["service-negate"] = v.(bool)
	}

	if v, ok := d.GetOk("source"); ok {
		threatException["source"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("source_negate"); ok {
		threatException["source-negate"] = v.(bool)
	}

	if v, ok := d.GetOk("track"); ok {
		threatException["track"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		threatException["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatException["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatException["ignore-warnings"] = v.(bool)
	}

	log.Println("Create Threat Exception - Map = ", threatException)

	addThreatExceptionRes, err := client.ApiCall("add-threat-exception", threatException, client.GetSessionID(), true, false)
	if err != nil || !addThreatExceptionRes.Success {
		if addThreatExceptionRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatExceptionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addThreatExceptionRes.GetData()["uid"].(string))

	return readManagementThreatException(d, m)
}

func readManagementThreatException(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	if v, ok := d.GetOk("exception_group_uid"); ok {
		payload["exception-group-uid"] = v.(string)
	}

	if v, ok := d.GetOk("exception_group_name"); ok {
		payload["exception-group-name"] = v.(string)
	}

	if v, ok := d.GetOk("rule_uid"); ok {
		payload["rule-uid"] = v.(string)
	}

	if v, ok := d.GetOk("rule_name"); ok {
		payload["rule-name"] = v.(string)
	}

	showThreatRuleRes, err := client.ApiCall("show-threat-exception", payload, client.GetSessionID(), true, false)
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

	exceptionRule := showThreatRuleRes.GetData()

	log.Println("Read Threat Exception - Show JSON = ", exceptionRule)

	if v := exceptionRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := exceptionRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if exceptionRule["source"] != nil {
		sourceJson := exceptionRule["source"].([]interface{})
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

	if v := exceptionRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if exceptionRule["destination"] != nil {
		destinationJson := exceptionRule["destination"].([]interface{})
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

	if v := exceptionRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if exceptionRule["protected-scope"] != nil {
		protectedScopeJson := exceptionRule["protected-scope"].([]interface{})
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

	if v := exceptionRule["protected-scope-negate"]; v != nil {
		_ = d.Set("protected_scope_negate", v)
	}

	if exceptionRule["service"] != nil {
		serviceJson := exceptionRule["service"].([]interface{})
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

	if v := exceptionRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if exceptionRule["install-on"] != nil {
		installOnJson := exceptionRule["install-on"].([]interface{})
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

	if exceptionRule["protection-or-site"] != nil {
		protectionOrSiteJson := exceptionRule["protection-or-site"].([]interface{})
		protectionOrSiteIds := make([]string, 0)
		if len(protectionOrSiteJson) > 0 {
			for _, protectionOrSite := range protectionOrSiteJson {
				protectionOrSite := protectionOrSite.(map[string]interface{})
				protectionOrSiteIds = append(protectionOrSiteIds, protectionOrSite["name"].(string))
			}
		}

		_, protectionOrSiteInConf := d.GetOk("protection_or_site")
		if protectionOrSiteIds[0] == "Any" && !protectionOrSiteInConf {
			_ = d.Set("protection_or_site", []interface{}{})
		} else {
			_ = d.Set("protection_or_site", protectionOrSiteIds)
		}
	}

	if v := exceptionRule["action"]; v != nil {
		_ = d.Set("action", v.(map[string]interface{})["name"])
	}

	if v := exceptionRule["track"]; v != nil {
		_ = d.Set("track", v.(map[string]interface{})["name"])
	}

	if v := exceptionRule["owner"]; v != nil {
		_ = d.Set("owner", v)
	}

	if v := exceptionRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementThreatException(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	threatException := make(map[string]interface{})

	threatException["uid"] = d.Id()
	threatException["layer"] = d.Get("layer")

	if v, ok := d.GetOk("exception_group_uid"); ok {
		threatException["exception-group-uid"] = v
	}

	if v, ok := d.GetOk("exception_group_name"); ok {
		threatException["exception-group-name"] = v
	}

	if v, ok := d.GetOk("rule_uid"); ok {
		threatException["rule-uid"] = v
	}

	if v, ok := d.GetOk("rule_name"); ok {
		threatException["rule-name"] = v
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					threatException["new-position"] = "top" // entire rule-base
				} else {
					threatException["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				threatException["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				threatException["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					threatException["new-position"] = "bottom" // entire rule-base
				} else {
					threatException["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	if d.HasChange("name") {
		threatException["new-name"] = d.Get("name")
	}

	if d.HasChange("action") {
		threatException["action"] = d.Get("action")
	}

	if d.HasChange("destination") {
		if v, ok := d.GetOk("destination"); ok {
			threatException["destination"] = v.(*schema.Set).List()
		} else {
			oldDestination, _ := d.GetChange("destination")
			threatException["destination"] = map[string]interface{}{"remove": oldDestination.(*schema.Set).List()}
		}
	}

	if d.HasChange("destination_negate") {
		threatException["destination-negate"] = d.Get("destination_negate")
	}

	if d.HasChange("enabled") {
		threatException["enabled"] = d.Get("enabled")
	}

	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			threatException["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstallOn, _ := d.GetChange("install_on")
			threatException["install-on"] = map[string]interface{}{"remove": oldInstallOn.(*schema.Set).List()}
		}
	}

	if d.HasChange("service") {
		if v, ok := d.GetOk("service"); ok {
			threatException["service"] = v.(*schema.Set).List()
		} else {
			oldService, _ := d.GetChange("service")
			threatException["service"] = map[string]interface{}{"remove": oldService.(*schema.Set).List()}
		}
	}

	if d.HasChange("service_negate") {
		threatException["service-negate"] = d.Get("service_negate")
	}

	if d.HasChange("source") {
		if v, ok := d.GetOk("source"); ok {
			threatException["source"] = v.(*schema.Set).List()
		} else {
			oldSource, _ := d.GetChange("source")
			threatException["source"] = map[string]interface{}{"remove": oldSource.(*schema.Set).List()}
		}
	}

	if d.HasChange("source_negate") {
		threatException["source-negate"] = d.Get("source_negate")
	}

	if d.HasChange("protected_scope") {
		if v, ok := d.GetOk("protected_scope"); ok {
			threatException["protected-scope"] = v.(*schema.Set).List()
		} else {
			oldProtectedSrc, _ := d.GetChange("protected_scope")
			threatException["protected-scope"] = map[string]interface{}{"remove": oldProtectedSrc.(*schema.Set).List()}
		}
	}

	if d.HasChange("protected_scope_negate") {
		threatException["protected-scope-negate"] = d.Get("protected_scope_negate")
	}

	if d.HasChange("protection_or_site") {
		if v, ok := d.GetOk("protection_or_site"); ok {
			threatException["protection-or-site"] = v.(*schema.Set).List()
		} else {
			oldProtectedOrSite, _ := d.GetChange("protection_or_site")
			threatException["protection-or-site"] = map[string]interface{}{"remove": oldProtectedOrSite.(*schema.Set).List()}
		}
	}

	if d.HasChange("track") {
		threatException["track"] = d.Get("track")
	}

	if d.HasChange("comments") {
		threatException["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatException["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatException["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Threat Exception - Map = ", threatException)

	updateThreatExceptionRes, err := client.ApiCall("set-threat-exception", threatException, client.GetSessionID(), true, false)
	if err != nil || !updateThreatExceptionRes.Success {
		if updateThreatExceptionRes.ErrorMsg != "" {
			return fmt.Errorf(updateThreatExceptionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	return readManagementThreatException(d, m)
}

func deleteManagementThreatException(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatExceptionPayload := map[string]interface{}{
		"uid":   d.Id(),
		"layer": d.Get("layer"),
	}

	if v, ok := d.GetOk("exception_group_uid"); ok {
		threatExceptionPayload["exception-group-uid"] = v
	}

	if v, ok := d.GetOk("exception_group_name"); ok {
		threatExceptionPayload["exception-group-name"] = v
	}

	if v, ok := d.GetOk("rule_uid"); ok {
		threatExceptionPayload["rule-uid"] = v
	}

	if v, ok := d.GetOk("rule_name"); ok {
		threatExceptionPayload["rule-name"] = v
	}

	deleteThreatExceptionRes, err := client.ApiCall("delete-threat-exception", threatExceptionPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteThreatExceptionRes.Success {
		if deleteThreatExceptionRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatExceptionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")
	return nil
}
