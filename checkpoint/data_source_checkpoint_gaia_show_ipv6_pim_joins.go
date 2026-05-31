package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIpv6PimJoins() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIpv6PimJoins,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "limit": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The maximum number of returned results.`,
            },
            "offset": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The number of results to initially skip.`,
            },
            "order": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Sorts results in either ascending or descending order.`,
            },
            "detailed": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Show sparse-mode detailed join state.`,
            },
            "group": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Show sparse-mode join state by group.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "objects": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "creation_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "expire_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "flags": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "group": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "outgoing_interfaces": {
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
                                    "jp_expiry_time": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "creation_time": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "jp_state": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "rp_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "src": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "register_sent": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "register_stop_received": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "register_received": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "register_stop_sent": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "spt_join_sent": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "register_state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "could_register": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "total": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
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
        },
    }
}

func readGaiaShowIpv6PimJoins(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("limit"); ok {
        payload["limit"] = v.(int)
    }

    if v, ok := d.GetOk("offset"); ok {
        payload["offset"] = v.(int)
    }

    if v, ok := d.GetOk("order"); ok {
        payload["order"] = v.(string)
    }

    if v, ok := d.GetOk("detailed"); ok {
        payload["detailed"] = v.(string)
    }

    if v, ok := d.GetOk("group"); ok {
        payload["group"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ipv6-pim-joins - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ipv6-pim-joins", payload)
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
            "ipv6-pim-joins",        // resource type
            "read",                       // operation
            "show-ipv6-pim-joins",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ipv6-pim-joins: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["objects"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "creation_time": func() string { if _v, _ok := m["creation-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "expire_time": func() string { if _v, _ok := m["expire-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "flags": func() string { if _v, _ok := m["flags"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "group": func() string { if _v, _ok := m["group"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "interface": func() string { if _v, _ok := m["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "interface_address": func() string { if _v, _ok := m["interface-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "outgoing_interfaces": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["outgoing-interfaces"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "jp_expiry_time": func() string { if _v, _ok := _sgm["jp-expiry-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "creation_time": func() string { if _v, _ok := _sgm["creation-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "jp_state": func() string { if _v, _ok := _sgm["jp-state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "rp_address": func() string { if _v, _ok := m["rp-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "src": func() string { if _v, _ok := m["src"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "register_sent": func() string { if _v, _ok := m["register-sent"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "register_stop_received": func() string { if _v, _ok := m["register-stop-received"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "register_received": func() string { if _v, _ok := m["register-received"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "register_stop_sent": func() string { if _v, _ok := m["register-stop-sent"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "spt_join_sent": func() string { if _v, _ok := m["spt-join-sent"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "register_state": func() string { if _v, _ok := m["register-state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "could_register": func() string { if _v, _ok := m["could-register"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("objects", mapped)
        }
    } else {
        d.Set("objects", []interface{}{})
    }
    if v, exists := commandRes.GetData()["total"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("total", int(_f))
        }
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
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ipv6-pim-joins-" + acctest.RandString(10)))
    return nil
}

