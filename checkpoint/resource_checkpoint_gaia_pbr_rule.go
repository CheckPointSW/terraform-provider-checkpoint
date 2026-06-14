package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaPbrRule() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPbrRule,
        Read:   readGaiaPbrRule,
        Update: updateGaiaPbrRule,
        Delete: deleteGaiaPbrRule,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "priority": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `PBR Rule Priority`,
            },
            "match": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `PBR Rule match conditions. These determine what traffic will match the PBR Rule.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interface": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Match traffic on inbound interface`,
                        },
                        "port": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Match traffic by service port`,
                        },
                        "protocol": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Match traffic by protocol`,
                        },
                        "destination": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Match traffic with destination network`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv4 address of network`,
                                    },
                                    "mask_length": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Mask length of network`,
                                    },
                                },
                            },
                        },
                        "source": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Match traffic with source network`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv4 address of network`,
                                    },
                                    "mask_length": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Mask length of network`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "action": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `PBR Rule actions. These specify the action to take if traffic matches the PBR Rule.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "table": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Name of PBR Table used to route matched traffic`,
                        },
                        "main_table": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Use the main routing table to route matched traffic`,
                        },
                        "prohibit": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Mark matched traffic as prohibited`,
                        },
                        "unreachable": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Report matched traffic as having an unreachable destination`,
                        },
                    },
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaPbrRule(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(int)
    }

    if v := d.Get("match"); len(v.([]interface{})) > 0 {
        _ = v
        matchMap := make(map[string]interface{})
        if v, ok := d.GetOk("match.0.interface"); ok {
            matchMap["interface"] = v.(string)
        }
        if v, ok := d.GetOk("match.0.port"); ok {
            matchMap["port"] = v.(int)
        }
        if v, ok := d.GetOk("match.0.protocol"); ok {
            matchMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("match.0.destination"); ok {
            _ = v
            destinationMap := make(map[string]interface{})
            if v, ok := d.GetOk("match.0.destination.0.address"); ok {
                destinationMap["address"] = v.(string)
            }
            if v, ok := d.GetOk("match.0.destination.0.mask_length"); ok {
                destinationMap["mask-length"] = v.(int)
            }
            if len(destinationMap) > 0 {
                matchMap["destination"] = destinationMap
            }
        }
        if v, ok := d.GetOk("match.0.source"); ok {
            _ = v
            sourceMap := make(map[string]interface{})
            if v, ok := d.GetOk("match.0.source.0.address"); ok {
                sourceMap["address"] = v.(string)
            }
            if v, ok := d.GetOk("match.0.source.0.mask_length"); ok {
                sourceMap["mask-length"] = v.(int)
            }
            if len(sourceMap) > 0 {
                matchMap["source"] = sourceMap
            }
        }
        if len(matchMap) > 0 {
            payload["match"] = matchMap
        }
    }

    if v := d.Get("action"); len(v.([]interface{})) > 0 {
        _ = v
        actionMap := make(map[string]interface{})
        if v, ok := d.GetOk("action.0.table"); ok {
            actionMap["table"] = v.(string)
        }
        if v, ok := d.GetOkExists("action.0.main_table"); ok && v.(bool) {
            actionMap["main-table"] = v.(bool)
        }
        if v, ok := d.GetOkExists("action.0.prohibit"); ok && v.(bool) {
            actionMap["prohibit"] = v.(bool)
        }
        if v, ok := d.GetOkExists("action.0.unreachable"); ok && v.(bool) {
            actionMap["unreachable"] = v.(bool)
        }
        if len(actionMap) > 0 {
            payload["action"] = actionMap
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create PbrRule - Map = ", payload)

    addPbrRuleRes, err := client.ApiCallSimple("set-pbr-rule", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPbrRuleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPbrRuleRes.Success {
            errMsg = addPbrRuleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPbrRuleRes.GetData()
        }

        debugLogOperation(
            "pbr-rule",        // resource type
            "create",                       // operation
            "set-pbr-rule",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add pbr-rule: %v", err)
    }
    if !addPbrRuleRes.Success {
        if addPbrRuleRes.ErrorMsg != "" {
            return fmt.Errorf(addPbrRuleRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("pbr-rule-" + acctest.RandString(10)))
    return readGaiaPbrRule(d, m)
}

func readGaiaPbrRule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPbrRuleRes, err := client.ApiCallSimple("show-pbr-rule", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPbrRuleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPbrRuleRes.Success {
            errMsg = showPbrRuleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPbrRuleRes.GetData()
        }

        debugLogOperation(
            "pbr-rule",        // resource type
            "read",                       // operation
            "show-pbr-rule",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show pbr-rule: %v", err)
    }
    if !showPbrRuleRes.Success {
        if data := showPbrRuleRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPbrRuleRes.ErrorMsg)
    }

    pbrRule := showPbrRuleRes.GetData()

    log.Println("Read PbrRule - Show JSON = ", pbrRule)

    if v, exists := pbrRule["priority"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("priority", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("priority", _n)
            }
        }
    }
    if v, exists := pbrRule["match"]; exists {
        d.Set("match", v)
    }
    if v, exists := pbrRule["action"]; exists {
        d.Set("action", v)
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPbrRule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(int)
    }

    if v := d.Get("match"); len(v.([]interface{})) > 0 {
        _ = v
        matchMap := make(map[string]interface{})
        if v, ok := d.GetOk("match.0.interface"); ok {
            matchMap["interface"] = v.(string)
        }
        if v, ok := d.GetOk("match.0.port"); ok {
            matchMap["port"] = v.(int)
        }
        if v, ok := d.GetOk("match.0.protocol"); ok {
            matchMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("match.0.destination"); ok {
            _ = v
            destinationMap := make(map[string]interface{})
            if v, ok := d.GetOk("match.0.destination.0.address"); ok {
                destinationMap["address"] = v.(string)
            }
            if v, ok := d.GetOk("match.0.destination.0.mask_length"); ok {
                destinationMap["mask-length"] = v.(int)
            }
            if len(destinationMap) > 0 {
                matchMap["destination"] = destinationMap
            }
        }
        if v, ok := d.GetOk("match.0.source"); ok {
            _ = v
            sourceMap := make(map[string]interface{})
            if v, ok := d.GetOk("match.0.source.0.address"); ok {
                sourceMap["address"] = v.(string)
            }
            if v, ok := d.GetOk("match.0.source.0.mask_length"); ok {
                sourceMap["mask-length"] = v.(int)
            }
            if len(sourceMap) > 0 {
                matchMap["source"] = sourceMap
            }
        }
        if len(matchMap) > 0 {
            payload["match"] = matchMap
        }
    }

    if v := d.Get("action"); len(v.([]interface{})) > 0 {
        _ = v
        actionMap := make(map[string]interface{})
        if v, ok := d.GetOk("action.0.table"); ok {
            actionMap["table"] = v.(string)
        }
        if v, ok := d.GetOkExists("action.0.main_table"); ok && v.(bool) {
            actionMap["main-table"] = v.(bool)
        }
        if v, ok := d.GetOkExists("action.0.prohibit"); ok && v.(bool) {
            actionMap["prohibit"] = v.(bool)
        }
        if v, ok := d.GetOkExists("action.0.unreachable"); ok && v.(bool) {
            actionMap["unreachable"] = v.(bool)
        }
        if len(actionMap) > 0 {
            payload["action"] = actionMap
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setPbrRuleRes, err := client.ApiCallSimple("set-pbr-rule", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPbrRuleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPbrRuleRes.Success {
            errMsg = setPbrRuleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPbrRuleRes.GetData()
        }

        debugLogOperation(
            "pbr-rule",        // resource type
            "update",                       // operation
            "set-pbr-rule",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set pbr-rule: %v", err)
    }
    if !setPbrRuleRes.Success {
        return fmt.Errorf(setPbrRuleRes.ErrorMsg)
    }

    return readGaiaPbrRule(d, m)
}

func deleteGaiaPbrRule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    deletePbrRuleRes, err := client.ApiCallSimple("delete-pbr-rule", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deletePbrRuleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deletePbrRuleRes.Success {
            errMsg = deletePbrRuleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deletePbrRuleRes.GetData()
        }

        debugLogOperation(
            "pbr-rule",        // resource type
            "delete",                       // operation
            "delete-pbr-rule",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete pbr-rule: %v", err)
    }
    if !deletePbrRuleRes.Success {
        return fmt.Errorf(deletePbrRuleRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

