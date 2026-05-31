package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationIpv6Pim() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationIpv6Pim,
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
            "assert_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "assert_rank": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "protocol": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rank": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "bootstrap_candidate": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "local_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "priority": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "candidate_rp": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "advertise_interval": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "local_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "priority": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "multicast_group": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
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
            "data_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "hello_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "jp_delay_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "jp_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "mode": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "register_suppress_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_state_refresh": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "state_refresh_interval": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "state_refresh_ttl": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "static_rp": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rp_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "multicast_group": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
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
            "custom_ssm_prefix": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "spt_threshold": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "multicast_group": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "threshold": {
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

func readGaiaShowConfigurationIpv6Pim(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-ipv6-pim - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-ipv6-pim", payload)
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
            "configuration-ipv6-pim",        // resource type
            "read",                       // operation
            "show-configuration-ipv6-pim",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-ipv6-pim: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["assert-interval"]; exists {
        d.Set("assert_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["assert-rank"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "protocol": func() string { if _v, _ok := m["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rank": func() string { if _v, _ok := m["rank"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("assert_rank", mapped)
        }
    } else {
        d.Set("assert_rank", []interface{}{})
    }
    if v, exists := commandRes.GetData()["bootstrap-candidate"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("bootstrap_candidate", []interface{}{map[string]interface{}{
                "local_address": func() string { if _v, _ok := _m["local-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "priority": func() string { if _v, _ok := _m["priority"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "enable": func() bool { if b, ok := _m["enable"].(bool); ok { return b }; if s, ok := _m["enable"].(string); ok { return s == "true" }; return false }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["candidate-rp"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("candidate_rp", []interface{}{map[string]interface{}{
                "advertise_interval": func() int { if f, ok := _m["advertise-interval"].(float64); ok { return int(f) }; return 0 }(),
                "local_address": func() string { if _v, _ok := _m["local-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "enable": func() bool { if b, ok := _m["enable"].(bool); ok { return b }; if s, ok := _m["enable"].(string); ok { return s == "true" }; return false }(),
                "priority": func() int { if f, ok := _m["priority"].(float64); ok { return int(f) }; return 0 }(),
                "multicast_group": func() []interface{} {
                    if _nd, _ok := _m["multicast-group"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "address": func() string { if _v, _ok := _nd["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        }}
                    }
                    return []interface{}{}
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["data-interval"]; exists {
        d.Set("data_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["hello-interval"]; exists {
        d.Set("hello_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["jp-delay-interval"]; exists {
        d.Set("jp_delay_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["jp-interval"]; exists {
        d.Set("jp_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["register-suppress-interval"]; exists {
        d.Set("register_suppress_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-state-refresh"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_state_refresh", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_state_refresh", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["state-refresh-interval"]; exists {
        d.Set("state_refresh_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["state-refresh-ttl"]; exists {
        d.Set("state_refresh_ttl", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["static-rp"]; exists {
        buildStaticRP := func(m map[string]interface{}) map[string]interface{} {
            return map[string]interface{}{
                "rp_address": func() string { if _v, _ok := m["rp-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "enable": func() bool { if b, ok := m["enable"].(bool); ok { return b }; if s, ok := m["enable"].(string); ok { return s == "true" }; return false }(),
                "multicast_group": func() []interface{} {
                    if _nd, _ok := m["multicast-group"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "address": func() string { if _v, _ok := _nd["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        }}
                    }
                    return []interface{}{}
                }(),
            }
        }
        switch val := v.(type) {
        case map[string]interface{}:
            d.Set("static_rp", []interface{}{buildStaticRP(val)})
        case []interface{}:
            _mapped := make([]interface{}, 0, len(val))
            for _, item := range val {
                if m, ok := item.(map[string]interface{}); ok {
                    _mapped = append(_mapped, buildStaticRP(m))
                }
            }
            d.Set("static_rp", _mapped)
        }
    }
    if v, exists := commandRes.GetData()["custom-ssm-prefix"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "address": func() string { if _v, _ok := m["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("custom_ssm_prefix", mapped)
        }
    } else {
        d.Set("custom_ssm_prefix", []interface{}{})
    }
    if v, exists := commandRes.GetData()["spt-threshold"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "multicast_group": func() string { if _v, _ok := m["multicast-group"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "threshold": func() string { if _v, _ok := m["threshold"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("spt_threshold", mapped)
        }
    } else {
        d.Set("spt_threshold", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-configuration-ipv6-pim-" + acctest.RandString(10)))
    return nil
}

