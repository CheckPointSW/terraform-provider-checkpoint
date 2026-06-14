package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowMaestroGateways() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowMaestroGateways,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If true, show pending topology. If false, show deployed topology`,
            },
            "gateways": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "security_group": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "model": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "version": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "major": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "downlink_ports": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "orchestrator_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "port": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "weight": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowMaestroGateways(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    log.Println("Execute show-maestro-gateways - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-maestro-gateways", payload)
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
            "maestro-gateways",        // resource type
            "read",                       // operation
            "show-maestro-gateways",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-maestro-gateways: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["gateways"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "resource_id": func() string { if _v, _ok := m["id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site": func() int { if f, ok := m["site"].(float64); ok { return int(f) }; return 0 }(),
                        "security_group": func() int { if f, ok := m["security-group"].(float64); ok { return int(f) }; return 0 }(),
                        "member_id": func() int { if f, ok := m["member-id"].(float64); ok { return int(f) }; return 0 }(),
                        "model": func() string { if _v, _ok := m["model"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "version": func() []interface{} {
                            if _obj, _ok := m["version"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "major": func() string { if _v, _ok := _obj["major"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "downlink_ports": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["downlink-ports"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "orchestrator_id": func() string { if _v, _ok := _sgm["orchestrator-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "port": func() string { if _v, _ok := _sgm["port"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "description": func() string { if _v, _ok := m["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "weight": func() int { if f, ok := m["weight"].(float64); ok { return int(f) }; return 0 }(),
                    }
                }
            }
            d.Set("gateways", mapped)
        }
    } else {
        d.Set("gateways", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-maestro-gateways-" + acctest.RandString(10)))
    return nil
}

