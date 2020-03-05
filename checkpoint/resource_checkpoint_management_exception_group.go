package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	"strconv"
)

func resourceManagementExceptionGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementExceptionGroup,
        Read:   readManagementExceptionGroup,
        Update: updateManagementExceptionGroup,
        Delete: deleteManagementExceptionGroup,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "applied_profile": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile.",
            },
            "applied_threat_rules": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: "The threat rules to apply this group on in the case of apply-on manually-select-threat-rules.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "layer": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The layer of the threat rule to which the group is to be attached.",
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The name of the threat rule to which the group is to be attached.",
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
                },
            },
            "apply_on": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user.",
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of tag identifiers.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "color": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Color of the object. Should be one of existing colors.",
                Default:     "black",
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
        },
    }
}

func createManagementExceptionGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    exceptionGroup := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        exceptionGroup["name"] = v.(string)
    }

    if v, ok := d.GetOk("applied_profile"); ok {
        exceptionGroup["applied-profile"] = v.(string)
    }

    if v, ok := d.GetOk("applied_threat_rules"); ok {

        appliedThreatRulesList := v.([]interface{})

        if len(appliedThreatRulesList) > 0 {

            var appliedThreatRulesPayload []map[string]interface{}

            for i := range appliedThreatRulesList {

                Payload := make(map[string]interface{})

                if v, ok := d.GetOk("applied_threat_rules." + strconv.Itoa(i) + ".layer"); ok {
                    Payload["layer"] = v.(string)
                }
                if v, ok := d.GetOk("applied_threat_rules." + strconv.Itoa(i) + ".name"); ok {
                    Payload["name"] = v.(string)
                }
                if v, ok := d.GetOk("applied_threat_rules." + strconv.Itoa(i) + ".rule_number"); ok {
                    Payload["rule-number"] = v.(string)
                }
                if v, ok := d.GetOk("applied_threat_rules." + strconv.Itoa(i) + ".position"); ok {
                    Payload["position"] = v.(string)
                }
                appliedThreatRulesPayload = append(appliedThreatRulesPayload, Payload)
            }
            exceptionGroup["appliedThreatRules"] = appliedThreatRulesPayload
        }
    }

    if v, ok := d.GetOk("apply_on"); ok {
        exceptionGroup["apply-on"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        exceptionGroup["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        exceptionGroup["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        exceptionGroup["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        exceptionGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        exceptionGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Create ExceptionGroup - Map = ", exceptionGroup)

    addExceptionGroupRes, err := client.ApiCall("add-exception-group", exceptionGroup, client.GetSessionID(), true, false)
    if err != nil || !addExceptionGroupRes.Success {
        if addExceptionGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addExceptionGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addExceptionGroupRes.GetData()["uid"].(string))

    return readManagementExceptionGroup(d, m)
}

func readManagementExceptionGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showExceptionGroupRes, err := client.ApiCall("show-exception-group", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showExceptionGroupRes.Success {
		if objectNotFound(showExceptionGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showExceptionGroupRes.ErrorMsg)
    }

    exceptionGroup := showExceptionGroupRes.GetData()

    log.Println("Read ExceptionGroup - Show JSON = ", exceptionGroup)

	if v := exceptionGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := exceptionGroup["applied-profile"]; v != nil {
		_ = d.Set("applied_profile", v)
	}

    if exceptionGroup["applied-threat-rules"] != nil {

        appliedThreatRulesList, ok := exceptionGroup["applied-threat-rules"].([]interface{})

        if ok {

            if len(appliedThreatRulesList) > 0 {

                var appliedThreatRulesListToReturn []map[string]interface{}

                for i := range appliedThreatRulesList {

                    appliedThreatRulesMap := appliedThreatRulesList[i].(map[string]interface{})

                    appliedThreatRulesMapToAdd := make(map[string]interface{})

                    if v, _ := appliedThreatRulesMap["layer"]; v != nil {
                        appliedThreatRulesMapToAdd["layer"] = v
                    }
                    if v, _ := appliedThreatRulesMap["name"]; v != nil {
                        appliedThreatRulesMapToAdd["name"] = v
                    }
                    if v, _ := appliedThreatRulesMap["rule-number"]; v != nil {
                        appliedThreatRulesMapToAdd["rule_number"] = v
                    }
                    if v, _ := appliedThreatRulesMap["position"]; v != nil {
                        appliedThreatRulesMapToAdd["position"] = v
                    }
                    appliedThreatRulesListToReturn = append(appliedThreatRulesListToReturn, appliedThreatRulesMapToAdd)
                }
            }
        }
    }

	if v := exceptionGroup["apply-on"]; v != nil {
		_ = d.Set("apply_on", v)
	}

    if exceptionGroup["tags"] != nil {
        tagsJson, ok := exceptionGroup["tags"].([]interface{})
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

	if v := exceptionGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := exceptionGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := exceptionGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := exceptionGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementExceptionGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    exceptionGroup := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        exceptionGroup["name"] = oldName
        exceptionGroup["new-name"] = newName
    } else {
        exceptionGroup["name"] = d.Get("name")
    }

    if ok := d.HasChange("applied_profile"); ok {
	       exceptionGroup["applied-profile"] = d.Get("applied_profile")
    }

    if d.HasChange("applied_threat_rules") {

        if v, ok := d.GetOk("applied_threat_rules"); ok {

            appliedThreatRulesList := v.([]interface{})

            var appliedThreatRulesPayload []map[string]interface{}

            for i := range appliedThreatRulesList {

                Payload := make(map[string]interface{})

                if d.HasChange("applied_threat_rules." + strconv.Itoa(i) + ".layer") {
                    Payload["layer"] = d.Get("applied_threat_rules." + strconv.Itoa(i) + ".layer")
                }
                if d.HasChange("applied_threat_rules." + strconv.Itoa(i) + ".name") {
                    Payload["name"] = d.Get("applied_threat_rules." + strconv.Itoa(i) + ".name")
                }
                if d.HasChange("applied_threat_rules." + strconv.Itoa(i) + ".rule_number") {
                    Payload["rule-number"] = d.Get("applied_threat_rules." + strconv.Itoa(i) + ".rule_number")
                }
                if d.HasChange("applied_threat_rules." + strconv.Itoa(i) + ".position") {
                    Payload["position"] = d.Get("applied_threat_rules." + strconv.Itoa(i) + ".position")
                }
                appliedThreatRulesPayload = append(appliedThreatRulesPayload, Payload)
            }
            exceptionGroup["applied-threat-rules"] = appliedThreatRulesPayload
        } else {
            oldappliedThreatRules, _ := d.GetChange("applied_threat_rules")
            var appliedThreatRulesToDelete []interface{}
            for _, i := range oldappliedThreatRules.([]interface{}) {
                appliedThreatRulesToDelete = append(appliedThreatRulesToDelete, i.(map[string]interface{})["name"].(string))
            }
            exceptionGroup["applied-threat-rules"] = map[string]interface{}{"remove": appliedThreatRulesToDelete}
        }
    }

    if ok := d.HasChange("apply_on"); ok {
	       exceptionGroup["apply-on"] = d.Get("apply_on")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            exceptionGroup["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           exceptionGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       exceptionGroup["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       exceptionGroup["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       exceptionGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       exceptionGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Update ExceptionGroup - Map = ", exceptionGroup)

    updateExceptionGroupRes, err := client.ApiCall("set-exception-group", exceptionGroup, client.GetSessionID(), true, false)
    if err != nil || !updateExceptionGroupRes.Success {
        if updateExceptionGroupRes.ErrorMsg != "" {
            return fmt.Errorf(updateExceptionGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementExceptionGroup(d, m)
}

func deleteManagementExceptionGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    exceptionGroupPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ExceptionGroup")

    deleteExceptionGroupRes, err := client.ApiCall("delete-exception-group", exceptionGroupPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteExceptionGroupRes.Success {
        if deleteExceptionGroupRes.ErrorMsg != "" {
            return fmt.Errorf(deleteExceptionGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

