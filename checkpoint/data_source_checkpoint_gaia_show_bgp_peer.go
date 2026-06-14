package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBgpPeer() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBgpPeer,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "peer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The peer to be queried`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
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
        },
    }
}

func readGaiaShowBgpPeer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bgp-peer - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bgp-peer", payload)
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
            "bgp-peer",        // resource type
            "read",                       // operation
            "show-bgp-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bgp-peer: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["peer"]; exists {
        d.Set("peer", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["type"]; exists {
        d.Set("type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["as"]; exists {
        d.Set("as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["uptime"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("uptime", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["remote-as"]; exists {
        d.Set("remote_as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["local-as"]; exists {
        d.Set("local_as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["peer-capabilities-list"]; exists {
        d.Set("peer_capabilities_list", v.([]interface{}))
    } else {
        d.Set("peer_capabilities_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["our-capabilities-list"]; exists {
        d.Set("our_capabilities_list", v.([]interface{}))
    } else {
        d.Set("our_capabilities_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["authtype"]; exists {
        d.Set("authtype", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["auth-failures"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("auth_failures", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["enable-multihop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_multihop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_multihop", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["multihop-ttl"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("multihop_ttl", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["reachability-detection"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("reachability_detection", []interface{}{map[string]interface{}{
                "type": func() string { if _v, _ok := _m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "monitoring": func() bool { if b, ok := _m["monitoring"].(bool); ok { return b }; if s, ok := _m["monitoring"].(string); ok { return s == "true" }; return false }(),
                "monitoring_description": func() string { if _v, _ok := _m["monitoring-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "reachable": func() bool { if b, ok := _m["reachable"].(bool); ok { return b }; if s, ok := _m["reachable"].(string); ok { return s == "true" }; return false }(),
                "reachable_description": func() string { if _v, _ok := _m["reachable-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "last_received": func() int { if f, ok := _m["last-received"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["graceful-restart"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("graceful_restart", []interface{}{map[string]interface{}{
                "enabled": func() bool { if b, ok := _m["enabled"].(bool); ok { return b }; if s, ok := _m["enabled"].(string); ok { return s == "true" }; return false }(),
                "enabled_description": func() string { if _v, _ok := _m["enabled-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "state": func() string { if _v, _ok := _m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "restart_helper_expiry_time": func() int { if f, ok := _m["restart-helper-expiry-time"].(float64); ok { return int(f) }; return 0 }(),
                "stalepath_expiry_time": func() int { if f, ok := _m["stalepath-expiry-time"].(float64); ok { return int(f) }; return 0 }(),
                "stalepath_time": func() int { if f, ok := _m["stalepath-time"].(float64); ok { return int(f) }; return 0 }(),
                "restart_time_advertised": func() int { if f, ok := _m["restart-time-advertised"].(float64); ok { return int(f) }; return 0 }(),
                "restart_time_received": func() int { if f, ok := _m["restart-time-received"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["keepalives"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("keepalives", []interface{}{map[string]interface{}{
                "last_received": func() int { if f, ok := _m["last-received"].(float64); ok { return int(f) }; return 0 }(),
                "last_sent": func() int { if f, ok := _m["last-sent"].(float64); ok { return int(f) }; return 0 }(),
                "interval": func() int { if f, ok := _m["interval"].(float64); ok { return int(f) }; return 0 }(),
                "holdtime": func() int { if f, ok := _m["holdtime"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["received"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("received", []interface{}{map[string]interface{}{
                "routes_received": func() int { if f, ok := _m["routes-received"].(float64); ok { return int(f) }; return 0 }(),
                "routes_received_active": func() int { if f, ok := _m["routes-received-active"].(float64); ok { return int(f) }; return 0 }(),
                "ipv6_routes_received": func() int { if f, ok := _m["ipv6-routes-received"].(float64); ok { return int(f) }; return 0 }(),
                "ipv6_routes_received_active": func() int { if f, ok := _m["ipv6-routes-received-active"].(float64); ok { return int(f) }; return 0 }(),
                "notifications_list": func() []interface{} {
                    if _arr, _ok := _m["notifications-list"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "code": fmt.Sprintf("%v", _im["code"]),
                                    "sub_code": fmt.Sprintf("%v", _im["sub-code"]),
                                    "time": fmt.Sprintf("%v", _im["time"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["sent"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("sent", []interface{}{map[string]interface{}{
                "routes_sent": func() int { if f, ok := _m["routes-sent"].(float64); ok { return int(f) }; return 0 }(),
                "ipv6_routes_sent": func() int { if f, ok := _m["ipv6-routes-sent"].(float64); ok { return int(f) }; return 0 }(),
                "notifications_list": func() []interface{} {
                    if _arr, _ok := _m["notifications-list"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "code": fmt.Sprintf("%v", _im["code"]),
                                    "sub_code": fmt.Sprintf("%v", _im["sub-code"]),
                                    "time": fmt.Sprintf("%v", _im["time"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bgp-peer-" + acctest.RandString(10)))
    return nil
}

