package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowVxlanInterfaces() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowVxlanInterfaces,
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
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "objects": {
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
                        "member_of": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "vxlan_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "remote_ip_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "destination_port": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
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
                        "name": {
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

func readGaiaShowVxlanInterfaces(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-vxlan-interfaces - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-vxlan-interfaces", payload)
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
            "vxlan-interfaces",        // resource type
            "read",                       // operation
            "show-vxlan-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-vxlan-interfaces: %v", err)
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
                        "link_state": func() bool { if b, ok := m["link-state"].(bool); ok { return b }; if s, ok := m["link-state"].(string); ok { return s == "true" }; return false }(),
                        "speed": func() string { if _v, _ok := m["speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "duplex": func() string { if _v, _ok := m["duplex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "tx_bytes": func() int { if f, ok := m["tx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                        "tx_packets": func() int { if f, ok := m["tx-packets"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_bytes": func() int { if f, ok := m["rx-bytes"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_packets": func() int { if f, ok := m["rx-packets"].(float64); ok { return int(f) }; return 0 }(),
                        "member_of": func() string { if _v, _ok := m["member-of"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "vxlan_id": func() int { if f, ok := m["vxlan-id"].(float64); ok { return int(f) }; return 0 }(),
                        "remote_ip_address": func() string { if _v, _ok := m["remote-ip-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "destination_port": func() int { if f, ok := m["destination-port"].(float64); ok { return int(f) }; return 0 }(),
                        "mtu": func() string { if _v, _ok := m["mtu"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_address": func() string { if _v, _ok := m["ipv4-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_mask_length": func() string { if _v, _ok := m["ipv4-mask-length"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-vxlan-interfaces-" + acctest.RandString(10)))
    return nil
}

