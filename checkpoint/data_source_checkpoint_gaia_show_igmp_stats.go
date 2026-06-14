package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIgmpStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIgmpStats,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The type of IGMP statistics`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "receive": {
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
                    },
                },
            },
            "transmit": {
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
                    },
                },
            },
            "error": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rx_truncated": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tx_truncated": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "checksum": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "truncated": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ttl": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "non_local_neighbor": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unsupported_version": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unsupported_type": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_dispatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_igmp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_group": {
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
        },
    }
}

func readGaiaShowIgmpStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-igmp-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-igmp-stats", payload)
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
            "igmp-stats",        // resource type
            "read",                       // operation
            "show-igmp-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-igmp-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["receive"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("receive", []interface{}{map[string]interface{}{
                "total": func() int { if f, ok := _m["total"].(float64); ok { return int(f) }; return 0 }(),
                "membership_query": func() int { if f, ok := _m["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                "membership_report": func() int { if f, ok := _m["membership-report"].(float64); ok { return int(f) }; return 0 }(),
                "num_v3_report": func() int { if f, ok := _m["num-v3-report"].(float64); ok { return int(f) }; return 0 }(),
                "group_leave": func() int { if f, ok := _m["group-leave"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_request": func() int { if f, ok := _m["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_response": func() int { if f, ok := _m["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                "dvmrp": func() int { if f, ok := _m["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                "pim_v1": func() int { if f, ok := _m["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["transmit"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("transmit", []interface{}{map[string]interface{}{
                "total": func() int { if f, ok := _m["total"].(float64); ok { return int(f) }; return 0 }(),
                "membership_query": func() int { if f, ok := _m["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_request": func() int { if f, ok := _m["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_response": func() int { if f, ok := _m["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                "dvmrp": func() int { if f, ok := _m["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                "pim_v1": func() int { if f, ok := _m["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["error"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("error", []interface{}{map[string]interface{}{
                "rx_truncated": func() int { if f, ok := _m["rx-truncated"].(float64); ok { return int(f) }; return 0 }(),
                "tx_truncated": func() int { if f, ok := _m["tx-truncated"].(float64); ok { return int(f) }; return 0 }(),
                "checksum": func() int { if f, ok := _m["checksum"].(float64); ok { return int(f) }; return 0 }(),
                "truncated": func() int { if f, ok := _m["truncated"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ttl": func() int { if f, ok := _m["bad-ttl"].(float64); ok { return int(f) }; return 0 }(),
                "non_local_neighbor": func() int { if f, ok := _m["non-local-neighbor"].(float64); ok { return int(f) }; return 0 }(),
                "unsupported_version": func() int { if f, ok := _m["unsupported-version"].(float64); ok { return int(f) }; return 0 }(),
                "unsupported_type": func() int { if f, ok := _m["unsupported-type"].(float64); ok { return int(f) }; return 0 }(),
                "no_dispatch": func() int { if f, ok := _m["no-dispatch"].(float64); ok { return int(f) }; return 0 }(),
                "no_igmp": func() int { if f, ok := _m["no-igmp"].(float64); ok { return int(f) }; return 0 }(),
                "bad_group": func() int { if f, ok := _m["bad-group"].(float64); ok { return int(f) }; return 0 }(),
                "version_mismatch": func() int { if f, ok := _m["version-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "igmp_v1_host": func() int { if f, ok := _m["igmp-v1-host"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-igmp-stats-" + acctest.RandString(10)))
    return nil
}

