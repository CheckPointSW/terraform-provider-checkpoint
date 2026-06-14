package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPppoeInterfaces() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPppoeInterfaces,
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
                        "client_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "username": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "password_hash": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Sensitive:   true,
                            Description: `N/A`,
                        },
                        "use_peer_as_default_gateway": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "use_peer_dns": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "fake_peer_settings": {
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
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "comments": {
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

func readGaiaShowPppoeInterfaces(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-pppoe-interfaces - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-pppoe-interfaces", payload)
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
            "pppoe-interfaces",        // resource type
            "read",                       // operation
            "show-pppoe-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-pppoe-interfaces: %v", err)
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
                        "client_id": func() int { if f, ok := m["client-id"].(float64); ok { return int(f) }; return 0 }(),
                        "interface": func() string { if _v, _ok := m["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "username": func() string { if _v, _ok := m["username"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "password_hash": func() string { if _v, _ok := m["password-hash"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "use_peer_as_default_gateway": func() bool { if b, ok := m["use-peer-as-default-gateway"].(bool); ok { return b }; if s, ok := m["use-peer-as-default-gateway"].(string); ok { return s == "true" }; return false }(),
                        "use_peer_dns": func() bool { if b, ok := m["use-peer-dns"].(bool); ok { return b }; if s, ok := m["use-peer-dns"].(string); ok { return s == "true" }; return false }(),
                        "fake_peer_settings": func() []interface{} {
                            if _obj, _ok := m["fake-peer-settings"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "address": func() string { if _v, _ok := _obj["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                }}
                            }
                            return nil
                        }(),
                        "enabled": func() bool { if b, ok := m["enabled"].(bool); ok { return b }; if s, ok := m["enabled"].(string); ok { return s == "true" }; return false }(),
                        "status": func() string { if _v, _ok := m["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "comments": func() string { if _v, _ok := m["comments"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-pppoe-interfaces-" + acctest.RandString(10)))
    return nil
}

