package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIgmpInterfacesStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIgmpInterfacesStats,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
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
                Description: `Sorts the interface statistics entries by interface name in either ascending or descending order`,
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
                        "interface": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "receive": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "membership_query": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "membership_report": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "num_v3_report": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "group_leave": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mtrace_request": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mtrace_response": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dvmrp": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "pim_v1": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "version_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "igmp_v1_host": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "transmit": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "membership_query": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mtrace_request": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mtrace_response": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dvmrp": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "pim_v1": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "truncated": {
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

func readGaiaShowIgmpInterfacesStats(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-igmp-interfaces-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-igmp-interfaces-stats", payload)
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
            "igmp-interfaces-stats",        // resource type
            "read",                       // operation
            "show-igmp-interfaces-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-igmp-interfaces-stats: %v", err)
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
                        "interface": func() string { if _v, _ok := m["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "receive": func() []interface{} {
                            if _obj, _ok := m["receive"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "membership_query": func() int { if f, ok := _obj["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                                    "membership_report": func() int { if f, ok := _obj["membership-report"].(float64); ok { return int(f) }; return 0 }(),
                                    "num_v3_report": func() int { if f, ok := _obj["num-v3-report"].(float64); ok { return int(f) }; return 0 }(),
                                    "group_leave": func() int { if f, ok := _obj["group-leave"].(float64); ok { return int(f) }; return 0 }(),
                                    "mtrace_request": func() int { if f, ok := _obj["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                                    "mtrace_response": func() int { if f, ok := _obj["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                                    "dvmrp": func() int { if f, ok := _obj["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                                    "pim_v1": func() int { if f, ok := _obj["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
                                    "version_mismatch": func() int { if f, ok := _obj["version-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                    "igmp_v1_host": func() int { if f, ok := _obj["igmp-v1-host"].(float64); ok { return int(f) }; return 0 }(),
                                }}
                            }
                            return nil
                        }(),
                        "transmit": func() []interface{} {
                            if _obj, _ok := m["transmit"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "membership_query": func() int { if f, ok := _obj["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                                    "mtrace_request": func() int { if f, ok := _obj["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                                    "mtrace_response": func() int { if f, ok := _obj["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                                    "dvmrp": func() int { if f, ok := _obj["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                                    "pim_v1": func() int { if f, ok := _obj["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
                                    "truncated": func() int { if f, ok := _obj["truncated"].(float64); ok { return int(f) }; return 0 }(),
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
    d.SetId(fmt.Sprintf("show-igmp-interfaces-stats-" + acctest.RandString(10)))
    return nil
}

