package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowMaestroSites() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowMaestroSites,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "sites": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "site_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "descriptions": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "security_group": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
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
                },
            },
        },
    }
}

func readGaiaShowMaestroSites(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    log.Println("Execute show-maestro-sites - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-maestro-sites", payload)
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
            "maestro-sites",        // resource type
            "read",                       // operation
            "show-maestro-sites",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-maestro-sites: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["sites"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "site_id": func() int { if f, ok := m["site-id"].(float64); ok { return int(f) }; return 0 }(),
                        "descriptions": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["descriptions"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "security_group": func() int { if f, ok := _sgm["security-group"].(float64); ok { return int(f) }; return 0 }(),
                                            "description": func() string { if _v, _ok := _sgm["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "gateways": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["gateways"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "resource_id": func() string { if _v, _ok := _sgm["id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "site": func() int { if f, ok := _sgm["site"].(float64); ok { return int(f) }; return 0 }(),
                                            "security_group": func() int { if f, ok := _sgm["security-group"].(float64); ok { return int(f) }; return 0 }(),
                                            "member_id": func() int { if f, ok := _sgm["member-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "model": func() string { if _v, _ok := _sgm["model"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "version": func() []interface{} {
                                                if _dobj, _ok := _sgm["version"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "major": func() string { if _v, _ok := _dobj["major"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                            "downlink_ports": func() []interface{} {
                                                var _dpOut []interface{}
                                                if _dpArr, _dpOk := _sgm["downlink-ports"].([]interface{}); _dpOk {
                                                    for _, _dp := range _dpArr {
                                                        if _dpm, _dpOk := _dp.(map[string]interface{}); _dpOk {
                                                            _dpOut = append(_dpOut, map[string]interface{}{
                                                                "orchestrator_id": func() string { if _v, _ok := _dpm["orchestrator-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                                "port": func() string { if _v, _ok := _dpm["port"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                            })
                                                        }
                                                    }
                                                }
                                                return _dpOut
                                            }(),
                                            "description": func() string { if _v, _ok := _sgm["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "state": func() string { if _v, _ok := _sgm["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "weight": func() int { if f, ok := _sgm["weight"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                    }
                }
            }
            d.Set("sites", mapped)
        }
    } else {
        d.Set("sites", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-maestro-sites-" + acctest.RandString(10)))
    return nil
}

