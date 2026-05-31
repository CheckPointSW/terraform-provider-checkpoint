package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBootpStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBootpStats,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "summary": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the bootp stats returned`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "receive_summary": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "total": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "truncated": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "request": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "reply": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "request_summary": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dropped": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "no_bootp_on_interface": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "wait_time_policy": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "max_hops_policy": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "relayed": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "requests": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "packets": {
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
            "reply_summary": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dropped": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "non_local_client": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_bootp_on_interface": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_client_ip": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_client_hardware_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_client_hardware_address": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_client_arp_entry": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "relayed": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "unicast": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "broadcast": {
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

func readGaiaShowBootpStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("summary"); ok {
        payload["summary"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bootp-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bootp-stats", payload)
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
            "bootp-stats",        // resource type
            "read",                       // operation
            "show-bootp-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bootp-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["receive-summary"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("receive_summary", []interface{}{map[string]interface{}{
                "total": func() int { if f, ok := _m["total"].(float64); ok { return int(f) }; return 0 }(),
                "truncated": func() int { if f, ok := _m["truncated"].(float64); ok { return int(f) }; return 0 }(),
                "request": func() int { if f, ok := _m["request"].(float64); ok { return int(f) }; return 0 }(),
                "reply": func() int { if f, ok := _m["reply"].(float64); ok { return int(f) }; return 0 }(),
                "unknown": func() int { if f, ok := _m["unknown"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["request-summary"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("request_summary", []interface{}{map[string]interface{}{
                "dropped": func() []interface{} {
                    if _nd, _ok := _m["dropped"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "no_bootp_on_interface": func() int { if f, ok := _nd["no-bootp-on-interface"].(float64); ok { return int(f) }; return 0 }(),
                            "wait_time_policy": func() int { if f, ok := _nd["wait-time-policy"].(float64); ok { return int(f) }; return 0 }(),
                            "max_hops_policy": func() int { if f, ok := _nd["max-hops-policy"].(float64); ok { return int(f) }; return 0 }(),
                        }}
                    }
                    return []interface{}{}
                }(),
                "relayed": func() []interface{} {
                    if _nd, _ok := _m["relayed"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "requests": func() int { if f, ok := _nd["requests"].(float64); ok { return int(f) }; return 0 }(),
                            "packets": func() int { if f, ok := _nd["packets"].(float64); ok { return int(f) }; return 0 }(),
                        }}
                    }
                    return []interface{}{}
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["reply-summary"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("reply_summary", []interface{}{map[string]interface{}{
                "dropped": func() []interface{} {
                    if _nd, _ok := _m["dropped"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "non_local_client": func() int { if f, ok := _nd["non-local-client"].(float64); ok { return int(f) }; return 0 }(),
                            "no_bootp_on_interface": func() int { if f, ok := _nd["no-bootp-on-interface"].(float64); ok { return int(f) }; return 0 }(),
                            "bad_client_ip": func() int { if f, ok := _nd["bad-client-ip"].(float64); ok { return int(f) }; return 0 }(),
                            "bad_client_hardware_type": func() int { if f, ok := _nd["bad-client-hardware-type"].(float64); ok { return int(f) }; return 0 }(),
                            "bad_client_hardware_address": func() int { if f, ok := _nd["bad-client-hardware-address"].(float64); ok { return int(f) }; return 0 }(),
                            "bad_client_arp_entry": func() int { if f, ok := _nd["bad-client-arp-entry"].(float64); ok { return int(f) }; return 0 }(),
                        }}
                    }
                    return []interface{}{}
                }(),
                "relayed": func() []interface{} {
                    if _nd, _ok := _m["relayed"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "unicast": func() int { if f, ok := _nd["unicast"].(float64); ok { return int(f) }; return 0 }(),
                            "broadcast": func() int { if f, ok := _nd["broadcast"].(float64); ok { return int(f) }; return 0 }(),
                        }}
                    }
                    return []interface{}{}
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bootp-stats-" + acctest.RandString(10)))
    return nil
}

