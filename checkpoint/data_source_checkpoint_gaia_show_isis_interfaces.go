package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisInterfaces() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisInterfaces,
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
                Description: `Sorts the interfaces by their names in either ascending or descending order`,
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
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "local_circuit": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mac_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "circuit_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_link_local": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_global": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "network_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "next_p2p_hello": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "extended_circuit": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "p2p_neighbor_state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "level_1": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hello_interval": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hold_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "next_lan_hello": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "circuit_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "priority": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dis": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "next_csnp": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "neighbor_count": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "neighbor_up_count": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "authentication": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "level_2": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_metric": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hello_interval": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hold_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "next_lan_hello": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "circuit_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "priority": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dis": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "next_csnp": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "neighbor_count": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "neighbor_up_count": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "authentication": {
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

func readGaiaShowIsisInterfaces(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
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

    log.Println("Execute show-isis-interfaces - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-interfaces", payload)
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
            "isis-interfaces",        // resource type
            "read",                       // operation
            "show-isis-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-interfaces: %v", err)
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
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "local_circuit": func() string { if _v, _ok := m["local-circuit"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "mac_address": func() string { if _v, _ok := m["mac-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "circuit_type": func() string { if _v, _ok := m["circuit-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_address": func() string { if _v, _ok := m["ipv4-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_link_local": func() string { if _v, _ok := m["ipv6-link-local"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_global": func() string { if _v, _ok := m["ipv6-global"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "network_type": func() string { if _v, _ok := m["network-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "next_p2p_hello": func() int { if f, ok := m["next-p2p-hello"].(float64); ok { return int(f) }; return 0 }(),
                        "extended_circuit": func() string { if _v, _ok := m["extended-circuit"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "p2p_neighbor_state": func() string { if _v, _ok := m["p2p-neighbor-state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "level_1": func() []interface{} {
                            if _obj, _ok := m["level-1"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "metric": func() int { if f, ok := _obj["metric"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_metric": func() int { if f, ok := _obj["ipv6-metric"].(float64); ok { return int(f) }; return 0 }(),
                                    "hello_interval": func() int { if f, ok := _obj["hello-interval"].(float64); ok { return int(f) }; return 0 }(),
                                    "hold_time": func() int { if f, ok := _obj["hold-time"].(float64); ok { return int(f) }; return 0 }(),
                                    "next_lan_hello": func() int { if f, ok := _obj["next-lan-hello"].(float64); ok { return int(f) }; return 0 }(),
                                    "circuit_id": func() string { if _v, _ok := _obj["circuit-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "priority": func() int { if f, ok := _obj["priority"].(float64); ok { return int(f) }; return 0 }(),
                                    "dis": func() string { if _v, _ok := _obj["dis"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "next_csnp": func() int { if f, ok := _obj["next-csnp"].(float64); ok { return int(f) }; return 0 }(),
                                    "neighbor_count": func() int { if f, ok := _obj["neighbor-count"].(float64); ok { return int(f) }; return 0 }(),
                                    "neighbor_up_count": func() int { if f, ok := _obj["neighbor-up-count"].(float64); ok { return int(f) }; return 0 }(),
                                    "authentication": func() string { if _v, _ok := _obj["authentication"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "level_2": func() []interface{} {
                            if _obj, _ok := m["level-2"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "metric": func() int { if f, ok := _obj["metric"].(float64); ok { return int(f) }; return 0 }(),
                                    "ipv6_metric": func() int { if f, ok := _obj["ipv6-metric"].(float64); ok { return int(f) }; return 0 }(),
                                    "hello_interval": func() int { if f, ok := _obj["hello-interval"].(float64); ok { return int(f) }; return 0 }(),
                                    "hold_time": func() int { if f, ok := _obj["hold-time"].(float64); ok { return int(f) }; return 0 }(),
                                    "next_lan_hello": func() int { if f, ok := _obj["next-lan-hello"].(float64); ok { return int(f) }; return 0 }(),
                                    "circuit_id": func() string { if _v, _ok := _obj["circuit-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "priority": func() int { if f, ok := _obj["priority"].(float64); ok { return int(f) }; return 0 }(),
                                    "dis": func() string { if _v, _ok := _obj["dis"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "next_csnp": func() int { if f, ok := _obj["next-csnp"].(float64); ok { return int(f) }; return 0 }(),
                                    "neighbor_count": func() int { if f, ok := _obj["neighbor-count"].(float64); ok { return int(f) }; return 0 }(),
                                    "neighbor_up_count": func() int { if f, ok := _obj["neighbor-up-count"].(float64); ok { return int(f) }; return 0 }(),
                                    "authentication": func() string { if _v, _ok := _obj["authentication"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
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
    d.SetId(fmt.Sprintf("show-isis-interfaces-" + acctest.RandString(10)))
    return nil
}

