package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementThreatException() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatExceptionRead,
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the exception.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
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
			"protection_or_site": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Name of the protection or site.",
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
			"track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Packet tracking.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner UID.",
			},
		},
	}
}

func dataSourceManagementThreatExceptionRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"layer": d.Get("layer"),
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
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
		return fmt.Errorf(showThreatRuleRes.ErrorMsg)
	}

	exceptionRule := showThreatRuleRes.GetData()

	log.Println("Read Threat Exception - Show JSON = ", exceptionRule)

	if v := exceptionRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
