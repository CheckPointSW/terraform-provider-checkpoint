package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisDatabase() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisDatabase,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The instance to be queried`,
            },
            "level": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the results by IS-IS level`,
            },
            "lsp_type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the results by lsp-type`,
            },
            "system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the results by system-id`,
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
                Description: `Sorts the database entries by their level first, then their LSP IDs in either ascending or descending order.`,
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
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "sequence": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "checksum": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lifetime": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "flags": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "ipv6_flags": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "size": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
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
                                    "secret": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "area_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "mtid_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "nlpid_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "hostname": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "is_reach_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "neighbor_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "is_reach_ext_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "neighbor_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "is_reach_mt_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "neighbor_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ip_address_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "ipv6_address_list": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "ip_internal_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ip_prefix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ip_external_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ip_prefix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ip_extended_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ip_prefix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ipv6_reach_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ip_prefix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ipv6_reach_mt_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ip_prefix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeInt,
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

func readGaiaShowIsisDatabase(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOk("lsp_type"); ok {
        payload["lsp-type"] = v.(string)
    }

    if v, ok := d.GetOk("system_id"); ok {
        payload["system-id"] = v.(string)
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

    log.Println("Execute show-isis-database - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-database", payload)
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
            "isis-database",        // resource type
            "read",                       // operation
            "show-isis-database",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-database: %v", err)
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
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "lsp_id": func() string { if _v, _ok := m["lsp-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "sequence": func() string { if _v, _ok := m["sequence"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "checksum": func() string { if _v, _ok := m["checksum"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "lifetime": func() int { if f, ok := m["lifetime"].(float64); ok { return int(f) }; return 0 }(),
                        "flags": func() []interface{} {
                            switch _ev := m["flags"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "ipv6_flags": func() []interface{} {
                            switch _ev := m["ipv6-flags"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "size": func() int { if f, ok := m["size"].(float64); ok { return int(f) }; return 0 }(),
                        "authentication": func() []interface{} {
                            if _obj, _ok := m["authentication"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "secret": func() string { if _v, _ok := _obj["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "area_list": func() []interface{} {
                            switch _ev := m["area-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "mtid_list": func() []interface{} {
                            switch _ev := m["mtid-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "nlpid_list": func() []interface{} {
                            switch _ev := m["nlpid-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "hostname": func() string { if _v, _ok := m["hostname"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "is_reach_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["is-reach-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "neighbor_id": func() string { if _v, _ok := _sgm["neighbor-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "is_reach_ext_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["is-reach-ext-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "neighbor_id": func() string { if _v, _ok := _sgm["neighbor-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "is_reach_mt_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["is-reach-mt-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "neighbor_id": func() string { if _v, _ok := _sgm["neighbor-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ip_address_list": func() []interface{} {
                            switch _ev := m["ip-address-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "ipv6_address_list": func() []interface{} {
                            switch _ev := m["ipv6-address-list"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "ip_internal_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ip-internal-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ip_prefix": func() string { if _v, _ok := _sgm["ip-prefix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ip_external_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ip-external-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ip_prefix": func() string { if _v, _ok := _sgm["ip-prefix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ip_extended_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ip-extended-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ip_prefix": func() string { if _v, _ok := _sgm["ip-prefix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ipv6_reach_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ipv6-reach-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ip_prefix": func() string { if _v, _ok := _sgm["ip-prefix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "ipv6_reach_mt_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ipv6-reach-mt-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ip_prefix": func() string { if _v, _ok := _sgm["ip-prefix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "metric": func() int { if f, ok := _sgm["metric"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
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
    d.SetId(fmt.Sprintf("show-isis-database-" + acctest.RandString(10)))
    return nil
}

