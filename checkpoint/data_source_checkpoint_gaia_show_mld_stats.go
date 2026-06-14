package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowMldStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowMldStats,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The type of MLD statistics`,
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
                        "listener_query": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "listener_report": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "num_v2_report": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "listener_done": {
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
                        "listener_query": {
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
                        "tx_truncated": {
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
                        "no_mld": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_group": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_router_alert": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "non_linklocal_sender": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "version_mismatch": {
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

func readGaiaShowMldStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-mld-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-mld-stats", payload)
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
            "mld-stats",        // resource type
            "read",                       // operation
            "show-mld-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-mld-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["receive"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("receive", []interface{}{map[string]interface{}{
                "total": func() int { if f, ok := _m["total"].(float64); ok { return int(f) }; return 0 }(),
                "listener_query": func() int { if f, ok := _m["listener-query"].(float64); ok { return int(f) }; return 0 }(),
                "listener_report": func() int { if f, ok := _m["listener-report"].(float64); ok { return int(f) }; return 0 }(),
                "num_v2_report": func() int { if f, ok := _m["num-v2-report"].(float64); ok { return int(f) }; return 0 }(),
                "listener_done": func() int { if f, ok := _m["listener-done"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["transmit"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("transmit", []interface{}{map[string]interface{}{
                "total": func() int { if f, ok := _m["total"].(float64); ok { return int(f) }; return 0 }(),
                "listener_query": func() int { if f, ok := _m["listener-query"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["error"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("error", []interface{}{map[string]interface{}{
                "tx_truncated": func() int { if f, ok := _m["tx-truncated"].(float64); ok { return int(f) }; return 0 }(),
                "truncated": func() int { if f, ok := _m["truncated"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ttl": func() int { if f, ok := _m["bad-ttl"].(float64); ok { return int(f) }; return 0 }(),
                "no_mld": func() int { if f, ok := _m["no-mld"].(float64); ok { return int(f) }; return 0 }(),
                "bad_group": func() int { if f, ok := _m["bad-group"].(float64); ok { return int(f) }; return 0 }(),
                "bad_router_alert": func() int { if f, ok := _m["bad-router-alert"].(float64); ok { return int(f) }; return 0 }(),
                "non_linklocal_sender": func() int { if f, ok := _m["non-linklocal-sender"].(float64); ok { return int(f) }; return 0 }(),
                "version_mismatch": func() int { if f, ok := _m["version-mismatch"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-mld-stats-" + acctest.RandString(10)))
    return nil
}

