package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPbrRules() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPbrRules,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "limit": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The maximum number of returned results`,
            },
            "offset": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The number of results to initially skip`,
            },
            "order": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Sorts the rules by priority in either ascending or descending order.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "from": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "to": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "total": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "objects": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "priority": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "match": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "interface": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "port": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "protocol": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "destination": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "mask_length": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "source": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "mask_length": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "action": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "table": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "main_table": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "prohibit": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "unreachable": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowPbrRules(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("limit"); ok {
        payload["limit"] = v.(int)
    }

    if v, ok := d.GetOk("offset"); ok {
        payload["offset"] = v.(int)
    }

    if v, ok := d.GetOk("order"); ok {
        payload["order"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-pbr-rules - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-pbr-rules", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "pbr-rules",        // resource type
            "read",                       // operation
            "show-pbr-rules",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-pbr-rules: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["from"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("from", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["to"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("to", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["total"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("total", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["objects"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "priority": func() int { if f, ok := m["priority"].(float64); ok { return int(f) }; return 0 }(),
                        "match": func() []interface{} {
                            if _obj, _ok := m["match"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "interface": func() string { if _v, _ok := _obj["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "port": func() int { if f, ok := _obj["port"].(float64); ok { return int(f) }; return 0 }(),
                                    "protocol": func() string { if _v, _ok := _obj["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "destination": func() []interface{} {
                                        if _d2, _ok := _obj["destination"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "address": func() string { if _v, _ok := _d2["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "mask_length": func() int { if f, ok := _d2["mask-length"].(float64); ok { return int(f) }; return 0 }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "source": func() []interface{} {
                                        if _d2, _ok := _obj["source"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "address": func() string { if _v, _ok := _d2["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "mask_length": func() int { if f, ok := _d2["mask-length"].(float64); ok { return int(f) }; return 0 }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                        "action": func() []interface{} {
                            if _obj, _ok := m["action"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "table": func() string { if _v, _ok := _obj["table"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "main_table": func() bool { if b, ok := _obj["main-table"].(bool); ok { return b }; if s, ok := _obj["main-table"].(string); ok { return s == "true" }; return false }(),
                                    "prohibit": func() bool { if b, ok := _obj["prohibit"].(bool); ok { return b }; if s, ok := _obj["prohibit"].(string); ok { return s == "true" }; return false }(),
                                    "unreachable": func() bool { if b, ok := _obj["unreachable"].(bool); ok { return b }; if s, ok := _obj["unreachable"].(string); ok { return s == "true" }; return false }(),
                                }}
                            }
                            return nil
                        }(),
                    }
                }
            }
            d.Set("objects", mapped)
        }
    } else {
        d.Set("objects", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-pbr-rules-" + acctest.RandString(10)))
    return nil
}

