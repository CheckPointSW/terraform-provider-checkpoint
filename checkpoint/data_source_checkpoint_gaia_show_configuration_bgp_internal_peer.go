package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpInternalPeer() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpInternalPeer,
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
            "peer_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_suppress_default_originate": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "weight": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowConfigurationBgpInternalPeer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-internal-peer - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-internal-peer", payload)
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
            "configuration-bgp-internal-peer",        // resource type
            "read",                       // operation
            "show-configuration-bgp-internal-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-internal-peer: %v", err)
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
    if v, exists := commandRes.GetData()["peer-type"]; exists {
        d.Set("peer_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-suppress-default-originate"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_suppress_default_originate", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_suppress_default_originate", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["weight"]; exists {
        d.Set("weight", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-internal-peer-" + acctest.RandString(10)))
    return nil
}

