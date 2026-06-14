package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgp() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "as": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "cluster_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_communities": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "confederation": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aspath_loops_permitted": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "identifier": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "dampening": {
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
                        "keep_history": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "max_flap": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "reachable_decay": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "reuse_below": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "suppress_above": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unreachable_decay": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "default_med": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "default_route_gateway": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_ecmp": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "graceful_restart": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restart_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "selection_deferral_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "ping": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "count": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interval": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "routing_domain": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aspath_loops_permitted": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "identifier": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "enable_synchronization": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowConfigurationBgp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp", payload)
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
            "configuration-bgp",        // resource type
            "read",                       // operation
            "show-configuration-bgp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["as"]; exists {
        d.Set("as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["cluster-id"]; exists {
        d.Set("cluster_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-communities"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_communities", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_communities", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["confederation"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("confederation", []interface{}{map[string]interface{}{
                "aspath_loops_permitted": func() string { if _v, _ok := _m["aspath-loops-permitted"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "identifier": func() string { if _v, _ok := _m["identifier"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["dampening"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("dampening", []interface{}{map[string]interface{}{
                "enabled": func() bool { if b, ok := _m["enabled"].(bool); ok { return b }; if s, ok := _m["enabled"].(string); ok { return s == "true" }; return false }(),
                "keep_history": func() string { if _v, _ok := _m["keep-history"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "max_flap": func() string { if _v, _ok := _m["max-flap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "reachable_decay": func() string { if _v, _ok := _m["reachable-decay"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "reuse_below": func() string { if _v, _ok := _m["reuse-below"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "suppress_above": func() string { if _v, _ok := _m["suppress-above"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "unreachable_decay": func() string { if _v, _ok := _m["unreachable-decay"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["default-med"]; exists {
        d.Set("default_med", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["default-route-gateway"]; exists {
        d.Set("default_route_gateway", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-ecmp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ecmp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ecmp", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["graceful-restart"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("graceful_restart", []interface{}{map[string]interface{}{
                "restart_time": func() string { if _v, _ok := _m["restart-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "selection_deferral_time": func() string { if _v, _ok := _m["selection-deferral-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["ping"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("ping", []interface{}{map[string]interface{}{
                "count": func() string { if _v, _ok := _m["count"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "interval": func() string { if _v, _ok := _m["interval"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["routing-domain"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("routing_domain", []interface{}{map[string]interface{}{
                "aspath_loops_permitted": func() string { if _v, _ok := _m["aspath-loops-permitted"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "identifier": func() string { if _v, _ok := _m["identifier"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["enable-synchronization"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_synchronization", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_synchronization", s == "true")
        }
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-" + acctest.RandString(10)))
    return nil
}

