package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpInternalPeers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpInternalPeers,
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

func readGaiaShowConfigurationBgpInternalPeers(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-configuration-bgp-internal-peers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-internal-peers", payload)
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
            "configuration-bgp-internal-peers",        // resource type
            "read",                       // operation
            "show-configuration-bgp-internal-peers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-internal-peers: %v", err)
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
                        "peer_type": func() string { if _v, _ok := m["peer-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enable_suppress_default_originate": func() bool { if b, ok := m["enable-suppress-default-originate"].(bool); ok { return b }; if s, ok := m["enable-suppress-default-originate"].(string); ok { return s == "true" }; return false }(),
                        "weight": func() string { if _v, _ok := m["weight"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-configuration-bgp-internal-peers-" + acctest.RandString(10)))
    return nil
}

