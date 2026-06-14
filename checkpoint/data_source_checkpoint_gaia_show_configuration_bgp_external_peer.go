package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpExternalPeer() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpExternalPeer,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "peer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `IP address of the peer.`,
            },
            "remote_as": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The Autonomous System number of the peerThe value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
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
    }
}

func readGaiaShowConfigurationBgpExternalPeer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-external-peer - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-external-peer", payload)
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
            "configuration-bgp-external-peer",        // resource type
            "read",                       // operation
            "show-configuration-bgp-external-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-external-peer: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["accept-routes"]; exists {
        d.Set("accept_routes", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["authtype"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("authtype", []interface{}{map[string]interface{}{
                "type": func() string { if _v, _ok := _m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "secret": func() string { if _v, _ok := _m["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["capability"]; exists {
        d.Set("capability", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-graceful-restart"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_graceful_restart", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_graceful_restart", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["graceful-restart-stalepath-time"]; exists {
        d.Set("graceful_restart_stalepath_time", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["holdtime"]; exists {
        d.Set("holdtime", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-ignore-first-ashop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ignore_first_ashop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ignore_first_ashop", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["keepalive"]; exists {
        d.Set("keepalive", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-log-state-transitions"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_state_transitions", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_state_transitions", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["enable-log-warnings"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_warnings", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_warnings", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["enable-no-aggregator-id"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_no_aggregator_id", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_no_aggregator_id", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["outgoing-interface"]; exists {
        d.Set("outgoing_interface", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-passive-tcp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_passive_tcp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_passive_tcp", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["peer"]; exists {
        d.Set("peer", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ping", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["enable-route-refresh"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_route_refresh", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_route_refresh", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["enable-send-keepalives"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_send_keepalives", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_send_keepalives", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["throttle-count"]; exists {
        d.Set("throttle_count", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["trace"]; exists {
        d.Set("trace", v.([]interface{}))
    } else {
        d.Set("trace", []interface{}{})
    }
    if v, exists := commandRes.GetData()["peer-type"]; exists {
        d.Set("peer_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-accept-med"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_accept_med", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_accept_med", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["allowas-in-count"]; exists {
        d.Set("allowas_in_count", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-as-override"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_as_override", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_as_override", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["aspath-prepend-count"]; exists {
        d.Set("aspath_prepend_count", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["export-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "conditional_routemap": func() []interface{} {
                            if _obj, _ok := m["conditional-routemap"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "name": func() string { if _v, _ok := _obj["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "condition": func() string { if _v, _ok := _obj["condition"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                    }
                }
            }
            d.Set("export_routemap_list", mapped)
        }
    } else {
        d.Set("export_routemap_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["inject-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "any_pass_routemap": func() string { if _v, _ok := m["any-pass-routemap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("inject_routemap_list", mapped)
        }
    } else {
        d.Set("inject_routemap_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["import-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("import_routemap_list", mapped)
        }
    } else {
        d.Set("import_routemap_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["ip-reachability"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("ip_reachability", []interface{}{map[string]interface{}{
                "type": func() string { if _v, _ok := _m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "local_address": func() string { if _v, _ok := _m["local-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "check_control_plane_failure": func() bool { if b, ok := _m["check-control-plane-failure"].(bool); ok { return b }; if s, ok := _m["check-control-plane-failure"].(string); ok { return s == "true" }; return false }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["med-out"]; exists {
        d.Set("med_out", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-multihop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_multihop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_multihop", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["peer-local-as"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("peer_local_as", []interface{}{map[string]interface{}{
                "as": func() string { if _v, _ok := _m["as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "enable_dual_peering": func() bool { if b, ok := _m["enable-dual-peering"].(bool); ok { return b }; if s, ok := _m["enable-dual-peering"].(string); ok { return s == "true" }; return false }(),
                "enable_inbound_peer_local": func() bool { if b, ok := _m["enable-inbound-peer-local"].(bool); ok { return b }; if s, ok := _m["enable-inbound-peer-local"].(string); ok { return s == "true" }; return false }(),
                "enable_outbound_local": func() bool { if b, ok := _m["enable-outbound-local"].(bool); ok { return b }; if s, ok := _m["enable-outbound-local"].(string); ok { return s == "true" }; return false }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["enable-remove-private-as"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_remove_private_as", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_remove_private_as", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["remote-as"]; exists {
        d.Set("remote_as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["send-route-refresh"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("send_route_refresh", []interface{}{map[string]interface{}{
                "type": func() string { if _v, _ok := _m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "family": func() string { if _v, _ok := _m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["enable-suppress-default-originate"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_suppress_default_originate", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_suppress_default_originate", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["ttl"]; exists {
        d.Set("ttl", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-external-peer-" + acctest.RandString(10)))
    return nil
}

