package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBondInterfaces() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBondInterfaces,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
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
                        "sd_wan": {
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
                                    "next_hop": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "next_hop_ipv6": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "nat": {
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
                                                "ip": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "ipv6": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "tag": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bandwidth": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "upload_speed": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "download_speed": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "circuit_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "link_type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ip_conflicts": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "dhcp6": {
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
                                    "server_timeout": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "retry": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "leasetime": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "reacquire_timeout": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "using": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "dhcp": {
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
                                    "server_timeout": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "retry": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "leasetime": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "reacquire_timeout": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "link_state": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "speed": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "duplex": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tx_bytes": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tx_packets": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rx_bytes": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rx_packets": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "members": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "mtu": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv4_mask_length": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_autoconfig": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "comments": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_mask_length": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_local_link_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "status": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "link_state": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "speed": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "duplex": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "tx_bytes": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "tx_packets": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "rx_bytes": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "rx_packets": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "xmit_hash_policy": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "down_delay": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "up_delay": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "primary": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lacp_rate": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mode": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mii_interval": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "min_links": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "virtual_system_id": {
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

func readGaiaShowBondInterfaces(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bond-interfaces - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bond-interfaces", payload)
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
            "bond-interfaces",        // resource type
            "read",                       // operation
            "show-bond-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bond-interfaces: %v", err)
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
                        "sd_wan": func() []interface{} {
                            if _obj, _ok := m["sd-wan"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                    "next_hop": func() string { if _v, _ok := _obj["next-hop"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "next_hop_ipv6": func() string { if _v, _ok := _obj["next-hop-ipv6"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "nat": func() []interface{} {
                                        if _d2, _ok := _obj["nat"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "enabled": func() bool { if b, ok := _d2["enabled"].(bool); ok { return b }; if s, ok := _d2["enabled"].(string); ok { return s == "true" }; return false }(),
                                                "ip": func() string { if _v, _ok := _d2["ip"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "ipv6": func() string { if _v, _ok := _d2["ipv6"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "tag": func() string { if _v, _ok := _obj["tag"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "bandwidth": func() []interface{} {
                                        if _d2, _ok := _obj["bandwidth"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "upload_speed": func() string { if _v, _ok := _d2["upload-speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "download_speed": func() string { if _v, _ok := _d2["download-speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "circuit_id": func() string { if _v, _ok := _obj["circuit-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "link_type": func() string { if _v, _ok := _obj["link-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ip_conflicts": func() bool { if b, ok := m["ip-conflicts"].(bool); ok { return b }; if s, ok := m["ip-conflicts"].(string); ok { return s == "true" }; return false }(),
                        "dhcp6": func() []interface{} {
                            if _obj, _ok := m["dhcp6"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                    "server_timeout": func() string { if _v, _ok := _obj["server-timeout"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "retry": func() string { if _v, _ok := _obj["retry"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "leasetime": func() string { if _v, _ok := _obj["leasetime"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "reacquire_timeout": func() string { if _v, _ok := _obj["reacquire-timeout"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "using": func() string { if _v, _ok := _obj["using"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "dhcp": func() []interface{} {
                            if _obj, _ok := m["dhcp"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                    "server_timeout": func() string { if _v, _ok := _obj["server-timeout"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "retry": func() string { if _v, _ok := _obj["retry"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "leasetime": func() string { if _v, _ok := _obj["leasetime"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "reacquire_timeout": func() string { if _v, _ok := _obj["reacquire-timeout"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "link_state": func() bool { if b, ok := m["link-state"].(bool); ok { return b }; if s, ok := m["link-state"].(string); ok { return s == "true" }; return false }(),
                        "speed": func() string { if _v, _ok := m["speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "duplex": func() string { if _v, _ok := m["duplex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "tx_bytes": func() int { if f, ok := m["tx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                        "tx_packets": func() int { if f, ok := m["tx-packets"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_bytes": func() int { if f, ok := m["rx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_packets": func() int { if f, ok := m["rx-packets"].(float64); ok { return int(f) }; return 0 }(),
                        "members": func() []interface{} {
                            switch _ev := m["members"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "mtu": func() string { if _v, _ok := m["mtu"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_address": func() string { if _v, _ok := m["ipv4-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_mask_length": func() string { if _v, _ok := m["ipv4-mask-length"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enabled": func() bool { if b, ok := m["enabled"].(bool); ok { return b }; if s, ok := m["enabled"].(string); ok { return s == "true" }; return false }(),
                        "ipv6_autoconfig": func() bool { if b, ok := m["ipv6-autoconfig"].(bool); ok { return b }; if s, ok := m["ipv6-autoconfig"].(string); ok { return s == "true" }; return false }(),
                        "comments": func() string { if _v, _ok := m["comments"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_address": func() string { if _v, _ok := m["ipv6-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_mask_length": func() string { if _v, _ok := m["ipv6-mask-length"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_local_link_address": func() string { if _v, _ok := m["ipv6-local-link-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "status": func() []interface{} {
                            if _obj, _ok := m["status"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "link_state": func() bool { if b, ok := _obj["link-state"].(bool); ok { return b }; if s, ok := _obj["link-state"].(string); ok { return s == "true" }; return false }(),
                                    "speed": func() string { if _v, _ok := _obj["speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "duplex": func() string { if _v, _ok := _obj["duplex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "tx_bytes": func() int { if f, ok := _obj["tx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                                    "tx_packets": func() int { if f, ok := _obj["tx-packets"].(float64); ok { return int(f) }; return 0 }(),
                                    "rx_bytes": func() int { if f, ok := _obj["rx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                                    "rx_packets": func() int { if f, ok := _obj["rx-packets"].(float64); ok { return int(f) }; return 0 }(),
                                }}
                            }
                            return nil
                        }(),
                        "xmit_hash_policy": func() string { if _v, _ok := m["xmit-hash-policy"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "down_delay": func() int { if f, ok := m["down-delay"].(float64); ok { return int(f) }; return 0 }(),
                        "up_delay": func() int { if f, ok := m["up-delay"].(float64); ok { return int(f) }; return 0 }(),
                        "primary": func() string { if _v, _ok := m["primary"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "lacp_rate": func() string { if _v, _ok := m["lacp-rate"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "mode": func() string { if _v, _ok := m["mode"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "mii_interval": func() int { if f, ok := m["mii-interval"].(float64); ok { return int(f) }; return 0 }(),
                        "min_links": func() int { if f, ok := m["min-links"].(float64); ok { return int(f) }; return 0 }(),
                        "virtual_system_id": func() string { if _v, _ok := m["virtual-system-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bond-interfaces-" + acctest.RandString(10)))
    return nil
}

