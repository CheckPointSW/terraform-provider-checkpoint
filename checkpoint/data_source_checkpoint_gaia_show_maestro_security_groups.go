package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowMaestroSecurityGroups() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowMaestroSecurityGroups,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If true, show pending Security Groups changes. If false, show deployed topology`,
            },
            "security_groups": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interfaces": {
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
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "vlans": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
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
                        "sites": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "resource_id": {
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
                        "ftw_configuration": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hostname": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "is_vsx": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "one_time_password": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                    "admin_password": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "mgmt_connectivity": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv4_mask_length": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_mask_length": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "default_gateway": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_default_gateway": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "mgmt_interface_settings": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "create_mgmt_as_bond": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bond_mode": {
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
                    },
                },
            },
        },
    }
}

func readGaiaShowMaestroSecurityGroups(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    log.Println("Execute show-maestro-security-groups - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-maestro-security-groups", payload)
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
            "maestro-security-groups",        // resource type
            "read",                       // operation
            "show-maestro-security-groups",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-maestro-security-groups: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["security-groups"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "resource_id": func() int { if f, ok := m["id"].(float64); ok { return int(f) }; return 0 }(),
                        "interfaces": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["interfaces"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "resource_id": func() string { if _v, _ok := _sgm["id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "description": func() string { if _v, _ok := _sgm["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "vlans": func() []interface{} {
                                                if _sl, _slOk := _sgm["vlans"].([]interface{}); _slOk {
                                                    out := make([]interface{}, len(_sl))
                                                    for _i, _sv := range _sl {
                                                        out[_i] = fmt.Sprintf("%v", _sv)
                                                    }
                                                    return out
                                                }
                                                return []interface{}{}
                                            }(),
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
                        "sites": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["sites"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "resource_id": func() int { if f, ok := _sgm["id"].(float64); ok { return int(f) }; return 0 }(),
                                            "description": func() string { if _v, _ok := _sgm["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ftw_configuration": func() []interface{} {
                            if _obj, _ok := m["ftw-configuration"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "hostname": func() string { if _v, _ok := _obj["hostname"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "is_vsx": func() bool { if b, ok := _obj["is-vsx"].(bool); ok { return b }; if s, ok := _obj["is-vsx"].(string); ok { return s == "true" }; return false }(),
                                    "one_time_password": func() string { if _v, _ok := _obj["one-time-password"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "admin_password": func() string { if _v, _ok := _obj["admin-password"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "mgmt_connectivity": func() []interface{} {
                            if _obj, _ok := m["mgmt-connectivity"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "ipv4_address": func() string { if _v, _ok := _obj["ipv4-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "ipv6_address": func() string { if _v, _ok := _obj["ipv6-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "ipv4_mask_length": func() int { if f, ok := _obj["ipv4-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_mask_length": func() int { if f, ok := _obj["ipv6-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                                    "default_gateway": func() string { if _v, _ok := _obj["default-gateway"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "ipv6_default_gateway": func() string { if _v, _ok := _obj["ipv6-default-gateway"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "mgmt_interface_settings": func() []interface{} {
                            if _obj, _ok := m["mgmt-interface-settings"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "create_mgmt_as_bond": func() bool { if b, ok := _obj["create-mgmt-as-bond"].(bool); ok { return b }; if s, ok := _obj["create-mgmt-as-bond"].(string); ok { return s == "true" }; return false }(),
                                    "bond_mode": func() string { if _v, _ok := _obj["bond-mode"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "description": func() string { if _v, _ok := m["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("security_groups", mapped)
        }
    } else {
        d.Set("security_groups", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-maestro-security-groups-" + acctest.RandString(10)))
    return nil
}

