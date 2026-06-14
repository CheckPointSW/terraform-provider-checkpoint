package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIgmpInterfaceStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIgmpInterfaceStats,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of IGMP interface`,
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
    }
}

func readGaiaShowIgmpInterfaceStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-igmp-interface-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-igmp-interface-stats", payload)
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
            "igmp-interface-stats",        // resource type
            "read",                       // operation
            "show-igmp-interface-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-igmp-interface-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["receive"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("receive", []interface{}{map[string]interface{}{
                "membership_query": func() int { if f, ok := _m["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                "membership_report": func() int { if f, ok := _m["membership-report"].(float64); ok { return int(f) }; return 0 }(),
                "num_v3_report": func() int { if f, ok := _m["num-v3-report"].(float64); ok { return int(f) }; return 0 }(),
                "group_leave": func() int { if f, ok := _m["group-leave"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_request": func() int { if f, ok := _m["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_response": func() int { if f, ok := _m["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                "dvmrp": func() int { if f, ok := _m["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                "pim_v1": func() int { if f, ok := _m["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
                "version_mismatch": func() int { if f, ok := _m["version-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "igmp_v1_host": func() int { if f, ok := _m["igmp-v1-host"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["transmit"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("transmit", []interface{}{map[string]interface{}{
                "membership_query": func() int { if f, ok := _m["membership-query"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_request": func() int { if f, ok := _m["mtrace-request"].(float64); ok { return int(f) }; return 0 }(),
                "mtrace_response": func() int { if f, ok := _m["mtrace-response"].(float64); ok { return int(f) }; return 0 }(),
                "dvmrp": func() int { if f, ok := _m["dvmrp"].(float64); ok { return int(f) }; return 0 }(),
                "pim_v1": func() int { if f, ok := _m["pim-v1"].(float64); ok { return int(f) }; return 0 }(),
                "truncated": func() int { if f, ok := _m["truncated"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    d.SetId(fmt.Sprintf("show-igmp-interface-stats-" + acctest.RandString(10)))
    return nil
}

