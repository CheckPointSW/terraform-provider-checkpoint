package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementHttpsRule() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementHttpsRule,
        Read:   readManagementHttpsRule,
        Update: updateManagementHttpsRule,
        Delete: deleteManagementHttpsRule,
        Schema: map[string]*schema.Schema{ 
            "layer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Layer that holds the Object. Identified by the Name or UID.",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "HTTPS rule name.",
            },
            "destination": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of Network objects identified by Name or UID that represents connection destination.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "service": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of Network objects identified by Name or UID that represents connection service.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "source": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of Network objects identified by Name or UID that represents connection source.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "action": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Rule inspect level. \"Bypass\" or \"Inspect\".",
            },
            "blade": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for. \"Anti Bot\",\"Anti Virus\",\"Application Control\",\"Data Awareness\",\"DLP\",\"IPS\",\"Threat Emulation\",\"Url Filtering\".",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "certificate": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Internal Server Certificate identified by Name or UID, otherwise, \"Outbound Certificate\" is a default value.",
            },
            "destination_negate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "TRUE if \"negate\" value is set for Destination.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable/Disable the rule.",
            },
            "install_on": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Which Gateways identified by the name or UID to install the policy on.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "service_negate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "TRUE if \"negate\" value is set for Service.",
            },
            "site_category": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of Site Categories objects identified by the name or UID.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "site_category_negate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "TRUE if \"negate\" value is set for Site Category.",
            },
            "source_negate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "TRUE if \"negate\" value is set for Source.",
            },
            "track": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "\"None\",\"Log\",\"Alert\",\"Mail\",\"SNMP trap\",\"Mail\",\"User Alert\", \"User Alert 2\", \"User Alert 3\".",
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Comments string.",
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
       			"position": &schema.Schema{
       				Type:        schema.TypeMap,
       				Required:    true,
       				Description: "Position in the rulebase.",
       				Elem: &schema.Resource{
       					Schema: map[string]*schema.Schema{
       						"top": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"above": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"below": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"bottom": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       					},
       				},
       			},
        },
    }
}

func createManagementHttpsRule(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    httpsRule := make(map[string]interface{})

    if v, ok := d.GetOk("rule_number"); ok {
        httpsRule["rule-number"] = v.(string)
    }

    if v, ok := d.GetOk("layer"); ok {
        httpsRule["layer"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        httpsRule["name"] = v.(string)
    }

    if v, ok := d.GetOk("destination"); ok {
        httpsRule["destination"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("service"); ok {
        httpsRule["service"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("source"); ok {
        httpsRule["source"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("action"); ok {
        httpsRule["action"] = v.(string)
    }

    if v, ok := d.GetOk("blade"); ok {
        httpsRule["blade"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("certificate"); ok {
        httpsRule["certificate"] = v.(string)
    }

    if v, ok := d.GetOkExists("destination_negate"); ok {
        httpsRule["destination-negate"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        httpsRule["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("install_on"); ok {
        httpsRule["install-on"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("service_negate"); ok {
        httpsRule["service-negate"] = v.(bool)
    }

    if v, ok := d.GetOk("site_category"); ok {
        httpsRule["site-category"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("site_category_negate"); ok {
        httpsRule["site-category-negate"] = v.(bool)
    }

    if v, ok := d.GetOkExists("source_negate"); ok {
        httpsRule["source-negate"] = v.(bool)
    }

    if v, ok := d.GetOk("track"); ok {
        httpsRule["track"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        httpsRule["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        httpsRule["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        httpsRule["ignore-errors"] = v.(bool)
    }

    if _, ok := d.GetOk("position"); ok {
		if _, ok := d.GetOk("position.top"); ok {
            httpsRule["position"] = "top"
        }
        if v, ok := d.GetOk("position.above"); ok {
            httpsRule["position"] = map[string]interface{}{"above": v.(string)}
        }
        if v, ok := d.GetOk("position.bottom"); ok {
            httpsRule["position"] = map[string]interface{}{"bottom": v.(string)}
        }
        if _, ok := d.GetOk("position.bottom"); ok {
            httpsRule["position"] = "bottom"
        }
    }
    log.Println("Create HttpsRule - Map = ", httpsRule)

    addHttpsRuleRes, err := client.ApiCall("add-https-rule", httpsRule, client.GetSessionID(), true, false)
    if err != nil || !addHttpsRuleRes.Success {
        if addHttpsRuleRes.ErrorMsg != "" {
            return fmt.Errorf(addHttpsRuleRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addHttpsRuleRes.GetData()["uid"].(string))

    return readManagementHttpsRule(d, m)
}

func readManagementHttpsRule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
		"layer": d.Get("layer"),
    }

    showHttpsRuleRes, err := client.ApiCall("show-https-rule", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showHttpsRuleRes.Success {
		if objectNotFound(showHttpsRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showHttpsRuleRes.ErrorMsg)
    }

    httpsRule := showHttpsRuleRes.GetData()

    log.Println("Read HttpsRule - Show JSON = ", httpsRule)

	if v := httpsRule["rule-number"]; v != nil {
		_ = d.Set("rule_number", v)
	}

	if v := httpsRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if httpsRule["destination"] != nil {
        destinationJson, ok := httpsRule["destination"].([]interface{})
        if ok {
            destinationIds := make([]string, 0)
            if len(destinationJson) > 0 {
                for _, destination := range destinationJson {
                    destination := destination.(map[string]interface{})
                    destinationIds = append(destinationIds, destination["name"].(string))
                }
            }
        _ = d.Set("destination", destinationIds)
        }
    } else {
        _ = d.Set("destination", nil)
    }

    if httpsRule["service"] != nil {
        serviceJson, ok := httpsRule["service"].([]interface{})
        if ok {
            serviceIds := make([]string, 0)
            if len(serviceJson) > 0 {
                for _, service := range serviceJson {
                    service := service.(map[string]interface{})
                    serviceIds = append(serviceIds, service["name"].(string))
                }
            }
        _ = d.Set("service", serviceIds)
        }
    } else {
        _ = d.Set("service", nil)
    }

    if httpsRule["source"] != nil {
        sourceJson, ok := httpsRule["source"].([]interface{})
        if ok {
            sourceIds := make([]string, 0)
            if len(sourceJson) > 0 {
                for _, source := range sourceJson {
                    source := source.(map[string]interface{})
                    sourceIds = append(sourceIds, source["name"].(string))
                }
            }
        _ = d.Set("source", sourceIds)
        }
    } else {
        _ = d.Set("source", nil)
    }

	if v := httpsRule["action"]; v != nil {
		_ = d.Set("action", v)
	}

    if httpsRule["blade"] != nil {
        bladeJson, ok := httpsRule["blade"].([]interface{})
        if ok {
            bladeIds := make([]string, 0)
            if len(bladeJson) > 0 {
                for _, blade := range bladeJson {
                    blade := blade.(map[string]interface{})
                    bladeIds = append(bladeIds, blade["name"].(string))
                }
            }
        _ = d.Set("blade", bladeIds)
        }
    } else {
        _ = d.Set("blade", nil)
    }

	if v := httpsRule["certificate"]; v != nil {
		_ = d.Set("certificate", v)
	}

	if v := httpsRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if v := httpsRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

    if httpsRule["install_on"] != nil {
        installOnJson, ok := httpsRule["install_on"].([]interface{})
        if ok {
            installOnIds := make([]string, 0)
            if len(installOnJson) > 0 {
                for _, install_on := range installOnJson {
                    install_on := install_on.(map[string]interface{})
                    installOnIds = append(installOnIds, install_on["name"].(string))
                }
            }
        _ = d.Set("install_on", installOnIds)
        }
    } else {
        _ = d.Set("install_on", nil)
    }

	if v := httpsRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

    if httpsRule["site_category"] != nil {
        siteCategoryJson, ok := httpsRule["site_category"].([]interface{})
        if ok {
            siteCategoryIds := make([]string, 0)
            if len(siteCategoryJson) > 0 {
                for _, site_category := range siteCategoryJson {
                    site_category := site_category.(map[string]interface{})
                    siteCategoryIds = append(siteCategoryIds, site_category["name"].(string))
                }
            }
        _ = d.Set("site_category", siteCategoryIds)
        }
    } else {
        _ = d.Set("site_category", nil)
    }

	if v := httpsRule["site-category-negate"]; v != nil {
		_ = d.Set("site_category_negate", v)
	}

	if v := httpsRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if v := httpsRule["track"]; v != nil {
		_ = d.Set("track", v)
	}

	if v := httpsRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := httpsRule["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := httpsRule["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementHttpsRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    httpsRule := make(map[string]interface{})

    if ok := d.HasChange("rule_number"); ok {
	       httpsRule["rule-number"] = d.Get("rule_number")
    }

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        httpsRule["name"] = oldName
        httpsRule["new-name"] = newName
    } else {
        httpsRule["name"] = d.Get("name")
    }

    if d.HasChange("destination") {
        if v, ok := d.GetOk("destination"); ok {
            httpsRule["destination"] = v.(*schema.Set).List()
        } else {
            oldDestination, _ := d.GetChange("destination")
	           httpsRule["destination"] = map[string]interface{}{"remove": oldDestination.(*schema.Set).List()}
        }
    }

    if d.HasChange("service") {
        if v, ok := d.GetOk("service"); ok {
            httpsRule["service"] = v.(*schema.Set).List()
        } else {
            oldService, _ := d.GetChange("service")
	           httpsRule["service"] = map[string]interface{}{"remove": oldService.(*schema.Set).List()}
        }
    }

    if d.HasChange("source") {
        if v, ok := d.GetOk("source"); ok {
            httpsRule["source"] = v.(*schema.Set).List()
        } else {
            oldSource, _ := d.GetChange("source")
	           httpsRule["source"] = map[string]interface{}{"remove": oldSource.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("action"); ok {
	       httpsRule["action"] = d.Get("action")
    }

    if d.HasChange("blade") {
        if v, ok := d.GetOk("blade"); ok {
            httpsRule["blade"] = v.(*schema.Set).List()
        } else {
            oldBlade, _ := d.GetChange("blade")
	           httpsRule["blade"] = map[string]interface{}{"remove": oldBlade.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("certificate"); ok {
	       httpsRule["certificate"] = d.Get("certificate")
    }

    if v, ok := d.GetOkExists("destination_negate"); ok {
	       httpsRule["destination-negate"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
	       httpsRule["enabled"] = v.(bool)
    }

    if d.HasChange("install_on") {
        if v, ok := d.GetOk("install_on"); ok {
            httpsRule["install_on"] = v.(*schema.Set).List()
        } else {
            oldInstall_On, _ := d.GetChange("install_on")
	           httpsRule["install_on"] = map[string]interface{}{"remove": oldInstall_On.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("service_negate"); ok {
	       httpsRule["service-negate"] = v.(bool)
    }

    if d.HasChange("site_category") {
        if v, ok := d.GetOk("site_category"); ok {
            httpsRule["site_category"] = v.(*schema.Set).List()
        } else {
            oldSite_Category, _ := d.GetChange("site_category")
	           httpsRule["site_category"] = map[string]interface{}{"remove": oldSite_Category.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("site_category_negate"); ok {
	       httpsRule["site-category-negate"] = v.(bool)
    }

    if v, ok := d.GetOkExists("source_negate"); ok {
	       httpsRule["source-negate"] = v.(bool)
    }

    if ok := d.HasChange("track"); ok {
	       httpsRule["track"] = d.Get("track")
    }

    if ok := d.HasChange("comments"); ok {
	       httpsRule["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       httpsRule["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       httpsRule["ignore-errors"] = v.(bool)
    }

    if ok := d.HasChange("position"); ok {
		if _, ok := d.GetOk("position"); ok {
			if _, ok := d.GetOk("position.top"); ok {
                httpsRule["new-position"] = "top"
            }
            if v, ok := d.GetOk("position.above"); ok {
                httpsRule["new-position"] = map[string]interface{}{"above": v.(string)}
            }
            if v, ok := d.GetOk("position.below"); ok {
                httpsRule["new-position"] = map[string]interface{}{"below": v.(string)}
            }
            if _, ok := d.GetOk("position.bottom"); ok {
                httpsRule["new-position"] = "bottom"
            }
        }
    }

    log.Println("Update HttpsRule - Map = ", httpsRule)

    updateHttpsRuleRes, err := client.ApiCall("set-https-rule", httpsRule, client.GetSessionID(), true, false)
    if err != nil || !updateHttpsRuleRes.Success {
        if updateHttpsRuleRes.ErrorMsg != "" {
            return fmt.Errorf(updateHttpsRuleRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementHttpsRule(d, m)
}

func deleteManagementHttpsRule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    httpsRulePayload := map[string]interface{}{
        "uid": d.Id(),
        "layer": d.Get("layer"),
    }

    log.Println("Delete HttpsRule")

    deleteHttpsRuleRes, err := client.ApiCall("delete-https-rule", httpsRulePayload , client.GetSessionID(), true, false)
    if err != nil || !deleteHttpsRuleRes.Success {
        if deleteHttpsRuleRes.ErrorMsg != "" {
            return fmt.Errorf(deleteHttpsRuleRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

