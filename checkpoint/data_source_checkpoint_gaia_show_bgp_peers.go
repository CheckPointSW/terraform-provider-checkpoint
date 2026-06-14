package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBgpPeers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBgpPeers,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "filter": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the results`,
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
                Description: `Sorts the peers first by their AS, then by their IDs in either ascending or descending order.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
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
                        "peer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "uptime": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "remote_as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "local_as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "peer_capabilities_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "our_capabilities_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "authtype": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "auth_failures": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_multihop": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "multihop_ttl": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "reachability_detection": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "monitoring": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "monitoring_description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "reachable": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "reachable_description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "last_received": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "graceful_restart": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "enabled_description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "state": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "restart_helper_expiry_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "stalepath_expiry_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "stalepath_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "restart_time_advertised": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "restart_time_received": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "keepalives": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "last_received": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "last_sent": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "interval": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "holdtime": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "received": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "routes_received": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "routes_received_active": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_routes_received": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_routes_received_active": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "notifications_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "code": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "sub_code": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "time": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "sent": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "routes_sent": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_routes_sent": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "notifications_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "code": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "sub_code": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "time": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "member_id": {
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

func readGaiaShowBgpPeers(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("filter"); ok {
        payload["filter"] = v.(string)
    }

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

    log.Println("Execute show-bgp-peers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bgp-peers", payload)
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
            "bgp-peers",        // resource type
            "read",                       // operation
            "show-bgp-peers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bgp-peers: %v", err)
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
                        "peer": func() string { if _v, _ok := m["peer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "as": func() string { if _v, _ok := m["as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "uptime": func() int { if f, ok := m["uptime"].(float64); ok { return int(f) }; return 0 }(),
                        "remote_as": func() string { if _v, _ok := m["remote-as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "local_as": func() string { if _v, _ok := m["local-as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "peer_capabilities_list": func() []interface{} {
                            switch _ev := m["peer-capabilities-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "our_capabilities_list": func() []interface{} {
                            switch _ev := m["our-capabilities-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "authtype": func() string { if _v, _ok := m["authtype"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "auth_failures": func() int { if f, ok := m["auth-failures"].(float64); ok { return int(f) }; return 0 }(),
                        "enable_multihop": func() bool { if b, ok := m["enable-multihop"].(bool); ok { return b }; if s, ok := m["enable-multihop"].(string); ok { return s == "true" }; return false }(),
                        "multihop_ttl": func() int { if f, ok := m["multihop-ttl"].(float64); ok { return int(f) }; return 0 }(),
                        "reachability_detection": func() []interface{} {
                            if _obj, _ok := m["reachability-detection"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "monitoring": func() bool { if b, ok := _obj["monitoring"].(bool); ok { return b }; if s, ok := _obj["monitoring"].(string); ok { return s == "true" }; return false }(),
                                    "monitoring_description": func() string { if _v, _ok := _obj["monitoring-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "reachable": func() bool { if b, ok := _obj["reachable"].(bool); ok { return b }; if s, ok := _obj["reachable"].(string); ok { return s == "true" }; return false }(),
                                    "reachable_description": func() string { if _v, _ok := _obj["reachable-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "last_received": func() int { if f, ok := _obj["last-received"].(float64); ok { return int(f) }; return 0 }(),
                                }}
                            }
                            return nil
                        }(),
                        "graceful_restart": func() []interface{} {
                            if _obj, _ok := m["graceful-restart"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                    "enabled_description": func() string { if _v, _ok := _obj["enabled-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "state": func() string { if _v, _ok := _obj["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "restart_helper_expiry_time": func() int { if f, ok := _obj["restart-helper-expiry-time"].(float64); ok { return int(f) }; return 0 }(),
                                    "stalepath_expiry_time": func() int { if f, ok := _obj["stalepath-expiry-time"].(float64); ok { return int(f) }; return 0 }(),
                                    "stalepath_time": func() int { if f, ok := _obj["stalepath-time"].(float64); ok { return int(f) }; return 0 }(),
                                    "restart_time_advertised": func() int { if f, ok := _obj["restart-time-advertised"].(float64); ok { return int(f) }; return 0 }(),
                                    "restart_time_received": func() int { if f, ok := _obj["restart-time-received"].(float64); ok { return int(f) }; return 0 }(),
                                }}
                            }
                            return nil
                        }(),
                        "keepalives": func() []interface{} {
                            if _obj, _ok := m["keepalives"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "last_received": func() int { if f, ok := _obj["last-received"].(float64); ok { return int(f) }; return 0 }(),
                                    "last_sent": func() int { if f, ok := _obj["last-sent"].(float64); ok { return int(f) }; return 0 }(),
                                    "interval": func() int { if f, ok := _obj["interval"].(float64); ok { return int(f) }; return 0 }(),
                                    "holdtime": func() int { if f, ok := _obj["holdtime"].(float64); ok { return int(f) }; return 0 }(),
                                }}
                            }
                            return nil
                        }(),
                        "received": func() []interface{} {
                            if _obj, _ok := m["received"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "routes_received": func() int { if f, ok := _obj["routes-received"].(float64); ok { return int(f) }; return 0 }(),
                                    "routes_received_active": func() int { if f, ok := _obj["routes-received-active"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_routes_received": func() int { if f, ok := _obj["ipv6-routes-received"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_routes_received_active": func() int { if f, ok := _obj["ipv6-routes-received-active"].(float64); ok { return int(f) }; return 0 }(),
                                    "notifications_list": func() []interface{} {
                                        if _gws, _ok := _obj["notifications-list"].([]interface{}); _ok {
                                            _gwOut := make([]interface{}, len(_gws))
                                            for _gi, _gw := range _gws {
                                                if _gwm, _ok := _gw.(map[string]interface{}); _ok {
                                                    _gwOut[_gi] = map[string]interface{}{
                                                        "code": func() string { if _v, _ok := _gwm["code"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "sub_code": func() string { if _v, _ok := _gwm["sub-code"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "time": func() string { if _v, _ok := _gwm["time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }
                                                }
                                            }
                                            return _gwOut
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                        "sent": func() []interface{} {
                            if _obj, _ok := m["sent"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "routes_sent": func() int { if f, ok := _obj["routes-sent"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_routes_sent": func() int { if f, ok := _obj["ipv6-routes-sent"].(float64); ok { return int(f) }; return 0 }(),
                                    "notifications_list": func() []interface{} {
                                        if _gws, _ok := _obj["notifications-list"].([]interface{}); _ok {
                                            _gwOut := make([]interface{}, len(_gws))
                                            for _gi, _gw := range _gws {
                                                if _gwm, _ok := _gw.(map[string]interface{}); _ok {
                                                    _gwOut[_gi] = map[string]interface{}{
                                                        "code": func() string { if _v, _ok := _gwm["code"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "sub_code": func() string { if _v, _ok := _gwm["sub-code"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "time": func() string { if _v, _ok := _gwm["time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }
                                                }
                                            }
                                            return _gwOut
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                        "member_id": func() string { if _v, _ok := m["member-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-bgp-peers-" + acctest.RandString(10)))
    return nil
}

