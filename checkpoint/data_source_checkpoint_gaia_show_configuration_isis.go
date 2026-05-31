package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationIsis() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationIsis,
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
            "adjacency_check": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "area_list": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "object": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "add": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "remove": {
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
                    },
                },
            },
            "default_metric": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "metric": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "is_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "max_areas": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "overload_bit": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "system_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ignore_attached_bit": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "overload_bit": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_attached_bit": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "multi_topology": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "prc_interval": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "second_offset": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "spf": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "second_offset": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "default_metric": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
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
            "dynamic_hostname": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "hello": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interface_point_to_point": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface_broadcast": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "metric_type": {
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
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "lsp": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "lifetime": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mtu": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "refresh_interval": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "gen_interval": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "second_offset": {
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
            "spf": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "max_interval": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "initial_offset": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "second_offset": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "prc_interval": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "max_interval": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "initial_offset": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "second_offset": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "authentication_ignore": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ignore_all": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_csnp": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_hello": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_lsp": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_psnp": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ignore_none": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "authentication": {
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
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "encrypted_secret": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Sensitive:   true,
                            Description: `N/A`,
                        },
                        "secret": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Sensitive:   true,
                            Description: `N/A`,
                        },
                        "keys": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "encrypted_secret": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                    "secret": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                    "resource_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "algorithm": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "active_key": {
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

func readGaiaShowConfigurationIsis(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-isis - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-isis", payload)
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
            "configuration-isis",        // resource type
            "read",                       // operation
            "show-configuration-isis",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-isis: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["adjacency-check"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("adjacency_check", b)
        } else if s, ok := v.(string); ok {
            d.Set("adjacency_check", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["default-metric"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "metric": func() string { if _v, _ok := m["metric"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("default_metric", mapped)
        }
    } else {
        d.Set("default_metric", []interface{}{})
    }
    if v, exists := commandRes.GetData()["is-type"]; exists {
        d.Set("is_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["max-areas"]; exists {
        d.Set("max_areas", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["overload-bit"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("overload_bit", b)
        } else if s, ok := v.(string); ok {
            d.Set("overload_bit", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["system-id"]; exists {
        d.Set("system_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ignore-attached-bit"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ignore_attached_bit", b)
        } else if s, ok := v.(string); ok {
            d.Set("ignore_attached_bit", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["ipv6"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("ipv6", []interface{}{map[string]interface{}{
                "overload_bit": func() bool { if b, ok := _m["overload-bit"].(bool); ok { return b }; if s, ok := _m["overload-bit"].(string); ok { return s == "true" }; return false }(),
                "ignore_attached_bit": func() bool { if b, ok := _m["ignore-attached-bit"].(bool); ok { return b }; if s, ok := _m["ignore-attached-bit"].(string); ok { return s == "true" }; return false }(),
                "multi_topology": func() bool { if b, ok := _m["multi-topology"].(bool); ok { return b }; if s, ok := _m["multi-topology"].(string); ok { return s == "true" }; return false }(),
                "prc_interval": func() []interface{} {
                    if _arr, _ok := _m["prc-interval"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "level": fmt.Sprintf("%v", _im["level"]),
                                    "max_interval": fmt.Sprintf("%v", _im["max-interval"]),
                                    "initial_offset": fmt.Sprintf("%v", _im["initial-offset"]),
                                    "second_offset": fmt.Sprintf("%v", _im["second-offset"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "spf": func() []interface{} {
                    if _arr, _ok := _m["spf"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "level": fmt.Sprintf("%v", _im["level"]),
                                    "max_interval": fmt.Sprintf("%v", _im["max-interval"]),
                                    "initial_offset": fmt.Sprintf("%v", _im["initial-offset"]),
                                    "second_offset": fmt.Sprintf("%v", _im["second-offset"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "default_metric": func() []interface{} {
                    if _arr, _ok := _m["default-metric"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "level": fmt.Sprintf("%v", _im["level"]),
                                    "metric": fmt.Sprintf("%v", _im["metric"]),
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
    if v, exists := commandRes.GetData()["dynamic-hostname"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("dynamic_hostname", b)
        } else if s, ok := v.(string); ok {
            d.Set("dynamic_hostname", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["hello"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("hello", []interface{}{map[string]interface{}{
                "interface_point_to_point": func() string { if _v, _ok := _m["interface-point-to-point"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "interface_broadcast": func() string { if _v, _ok := _m["interface-broadcast"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["metric-type"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("metric_type", mapped)
        }
    } else {
        d.Set("metric_type", []interface{}{})
    }
    if v, exists := commandRes.GetData()["lsp"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("lsp", []interface{}{map[string]interface{}{
                "lifetime": func() string { if _v, _ok := _m["lifetime"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "mtu": func() string { if _v, _ok := _m["mtu"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "refresh_interval": func() string { if _v, _ok := _m["refresh-interval"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "gen_interval": func() []interface{} {
                    if _arr, _ok := _m["gen-interval"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "level": fmt.Sprintf("%v", _im["level"]),
                                    "max_interval": fmt.Sprintf("%v", _im["max-interval"]),
                                    "initial_offset": fmt.Sprintf("%v", _im["initial-offset"]),
                                    "second_offset": fmt.Sprintf("%v", _im["second-offset"]),
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
    if v, exists := commandRes.GetData()["spf"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "max_interval": func() string { if _v, _ok := m["max-interval"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "initial_offset": func() string { if _v, _ok := m["initial-offset"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "second_offset": func() string { if _v, _ok := m["second-offset"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("spf", mapped)
        }
    } else {
        d.Set("spf", []interface{}{})
    }
    if v, exists := commandRes.GetData()["prc-interval"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "max_interval": func() string { if _v, _ok := m["max-interval"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "initial_offset": func() string { if _v, _ok := m["initial-offset"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "second_offset": func() string { if _v, _ok := m["second-offset"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("prc_interval", mapped)
        }
    } else {
        d.Set("prc_interval", []interface{}{})
    }
    if v, exists := commandRes.GetData()["authentication-ignore"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "ignore_all": func() bool { if b, ok := m["ignore-all"].(bool); ok { return b }; if s, ok := m["ignore-all"].(string); ok { return s == "true" }; return false }(),
                        "ignore_csnp": func() bool { if b, ok := m["ignore-csnp"].(bool); ok { return b }; if s, ok := m["ignore-csnp"].(string); ok { return s == "true" }; return false }(),
                        "ignore_hello": func() bool { if b, ok := m["ignore-hello"].(bool); ok { return b }; if s, ok := m["ignore-hello"].(string); ok { return s == "true" }; return false }(),
                        "ignore_lsp": func() bool { if b, ok := m["ignore-lsp"].(bool); ok { return b }; if s, ok := m["ignore-lsp"].(string); ok { return s == "true" }; return false }(),
                        "ignore_psnp": func() bool { if b, ok := m["ignore-psnp"].(bool); ok { return b }; if s, ok := m["ignore-psnp"].(string); ok { return s == "true" }; return false }(),
                        "ignore_none": func() bool { if b, ok := m["ignore-none"].(bool); ok { return b }; if s, ok := m["ignore-none"].(string); ok { return s == "true" }; return false }(),
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("authentication_ignore", mapped)
        }
    } else {
        d.Set("authentication_ignore", []interface{}{})
    }
    if v, exists := commandRes.GetData()["authentication"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "encrypted_secret": func() string { if _v, _ok := m["encrypted-secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "secret": func() string { if _v, _ok := m["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "keys": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["keys"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "encrypted_secret": func() string { if _v, _ok := _sgm["encrypted-secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "secret": func() string { if _v, _ok := _sgm["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "resource_id": func() int { if f, ok := _sgm["id"].(float64); ok { return int(f) }; return 0 }(),
                                            "algorithm": func() string { if _v, _ok := _sgm["algorithm"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "active_key": func() string { if _v, _ok := m["active-key"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("authentication", mapped)
        }
    } else {
        d.Set("authentication", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-configuration-isis-" + acctest.RandString(10)))
    return nil
}

