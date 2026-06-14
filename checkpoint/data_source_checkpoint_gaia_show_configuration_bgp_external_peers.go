package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpExternalPeers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpExternalPeers,
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
            "remote_as": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The Autonomous System number of the peerThe value can be one of the following:<br>'all'<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
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
                        "accept_routes": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authtype": {
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
                                    "secret": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "capability": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "comment": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_graceful_restart": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "graceful_restart_stalepath_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "holdtime": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_ignore_first_ashop": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "keepalive": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "local_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_log_state_transitions": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_log_warnings": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_no_aggregator_id": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "outgoing_interface": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_passive_tcp": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "peer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_ping": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_route_refresh": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_send_keepalives": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "throttle_count": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "trace": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "peer_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_accept_med": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "allowas_in_count": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_as_override": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "aspath_prepend_count": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "export_routemap_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "preference": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "family": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "conditional_routemap": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "condition": {
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
                        "inject_routemap_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "preference": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "any_pass_routemap": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "family": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "import_routemap_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "preference": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "family": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ip_reachability": {
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
                                    "local_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "check_control_plane_failure": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "med_out": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable_multihop": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "peer_local_as": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "as": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "enable_dual_peering": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "enable_inbound_peer_local": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "enable_outbound_local": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "enable_remove_private_as": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "remote_as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "send_route_refresh": {
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
                                    "family": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "enable_suppress_default_originate": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ttl": {
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

func readGaiaShowConfigurationBgpExternalPeers(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-external-peers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-external-peers", payload)
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
            "configuration-bgp-external-peers",        // resource type
            "read",                       // operation
            "show-configuration-bgp-external-peers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-external-peers: %v", err)
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
                        "accept_routes": func() string { if _v, _ok := m["accept-routes"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "authtype": func() []interface{} {
                            if _obj, _ok := m["authtype"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "secret": func() string { if _v, _ok := _obj["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "capability": func() string { if _v, _ok := m["capability"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "comment": func() string { if _v, _ok := m["comment"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_graceful_restart": func() bool { if b, ok := m["enable-graceful-restart"].(bool); ok { return b }; if s, ok := m["enable-graceful-restart"].(string); ok { return s == "true" }; return false }(),
                        "graceful_restart_stalepath_time": func() string { if _v, _ok := m["graceful-restart-stalepath-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "holdtime": func() string { if _v, _ok := m["holdtime"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_ignore_first_ashop": func() bool { if b, ok := m["enable-ignore-first-ashop"].(bool); ok { return b }; if s, ok := m["enable-ignore-first-ashop"].(string); ok { return s == "true" }; return false }(),
                        "keepalive": func() string { if _v, _ok := m["keepalive"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "local_address": func() string { if _v, _ok := m["local-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_log_state_transitions": func() bool { if b, ok := m["enable-log-state-transitions"].(bool); ok { return b }; if s, ok := m["enable-log-state-transitions"].(string); ok { return s == "true" }; return false }(),
                        "enable_log_warnings": func() bool { if b, ok := m["enable-log-warnings"].(bool); ok { return b }; if s, ok := m["enable-log-warnings"].(string); ok { return s == "true" }; return false }(),
                        "enable_no_aggregator_id": func() bool { if b, ok := m["enable-no-aggregator-id"].(bool); ok { return b }; if s, ok := m["enable-no-aggregator-id"].(string); ok { return s == "true" }; return false }(),
                        "outgoing_interface": func() string { if _v, _ok := m["outgoing-interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_passive_tcp": func() bool { if b, ok := m["enable-passive-tcp"].(bool); ok { return b }; if s, ok := m["enable-passive-tcp"].(string); ok { return s == "true" }; return false }(),
                        "peer": func() string { if _v, _ok := m["peer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_ping": func() bool { if b, ok := m["enable-ping"].(bool); ok { return b }; if s, ok := m["enable-ping"].(string); ok { return s == "true" }; return false }(),
                        "enable_route_refresh": func() bool { if b, ok := m["enable-route-refresh"].(bool); ok { return b }; if s, ok := m["enable-route-refresh"].(string); ok { return s == "true" }; return false }(),
                        "enable_send_keepalives": func() bool { if b, ok := m["enable-send-keepalives"].(bool); ok { return b }; if s, ok := m["enable-send-keepalives"].(string); ok { return s == "true" }; return false }(),
                        "throttle_count": func() string { if _v, _ok := m["throttle-count"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "trace": func() []interface{} {
                            switch _ev := m["trace"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "peer_type": func() string { if _v, _ok := m["peer-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_accept_med": func() bool { if b, ok := m["enable-accept-med"].(bool); ok { return b }; if s, ok := m["enable-accept-med"].(string); ok { return s == "true" }; return false }(),
                        "allowas_in_count": func() string { if _v, _ok := m["allowas-in-count"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_as_override": func() bool { if b, ok := m["enable-as-override"].(bool); ok { return b }; if s, ok := m["enable-as-override"].(string); ok { return s == "true" }; return false }(),
                        "aspath_prepend_count": func() string { if _v, _ok := m["aspath-prepend-count"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "export_routemap_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["export-routemap-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "preference": func() int { if f, ok := _sgm["preference"].(float64); ok { return int(f) }; return 0 }(),
                                            "family": func() string { if _v, _ok := _sgm["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "conditional_routemap": func() []interface{} {
                                                if _dobj, _ok := _sgm["conditional-routemap"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "name": func() string { if _v, _ok := _dobj["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "condition": func() string { if _v, _ok := _dobj["condition"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "inject_routemap_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["inject-routemap-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "preference": func() int { if f, ok := _sgm["preference"].(float64); ok { return int(f) }; return 0 }(),
                                            "any_pass_routemap": func() string { if _v, _ok := _sgm["any-pass-routemap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "family": func() string { if _v, _ok := _sgm["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "import_routemap_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["import-routemap-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "preference": func() int { if f, ok := _sgm["preference"].(float64); ok { return int(f) }; return 0 }(),
                                            "family": func() string { if _v, _ok := _sgm["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ip_reachability": func() []interface{} {
                            if _obj, _ok := m["ip-reachability"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "local_address": func() string { if _v, _ok := _obj["local-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "check_control_plane_failure": func() bool { if b, ok := _obj["check-control-plane-failure"].(bool); ok { return b }; if s, ok := _obj["check-control-plane-failure"].(string); ok { return s == "true" }; return false }(),
                                }}
                            }
                            return nil
                        }(),
                        "med_out": func() string { if _v, _ok := m["med-out"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_multihop": func() bool { if b, ok := m["enable-multihop"].(bool); ok { return b }; if s, ok := m["enable-multihop"].(string); ok { return s == "true" }; return false }(),
                        "peer_local_as": func() []interface{} {
                            if _obj, _ok := m["peer-local-as"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "as": func() string { if _v, _ok := _obj["as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "enable_dual_peering": func() bool { if b, ok := _obj["enable-dual-peering"].(bool); ok { return b }; if s, ok := _obj["enable-dual-peering"].(string); ok { return s == "true" }; return false }(),
                                    "enable_inbound_peer_local": func() bool { if b, ok := _obj["enable-inbound-peer-local"].(bool); ok { return b }; if s, ok := _obj["enable-inbound-peer-local"].(string); ok { return s == "true" }; return false }(),
                                    "enable_outbound_local": func() bool { if b, ok := _obj["enable-outbound-local"].(bool); ok { return b }; if s, ok := _obj["enable-outbound-local"].(string); ok { return s == "true" }; return false }(),
                                }}
                            }
                            return nil
                        }(),
                        "enable_remove_private_as": func() bool { if b, ok := m["enable-remove-private-as"].(bool); ok { return b }; if s, ok := m["enable-remove-private-as"].(string); ok { return s == "true" }; return false }(),
                        "remote_as": func() string { if _v, _ok := m["remote-as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "send_route_refresh": func() []interface{} {
                            if _obj, _ok := m["send-route-refresh"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "family": func() string { if _v, _ok := _obj["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "enable_suppress_default_originate": func() bool { if b, ok := m["enable-suppress-default-originate"].(bool); ok { return b }; if s, ok := m["enable-suppress-default-originate"].(string); ok { return s == "true" }; return false }(),
                        "ttl": func() string { if _v, _ok := m["ttl"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-configuration-bgp-external-peers-" + acctest.RandString(10)))
    return nil
}

