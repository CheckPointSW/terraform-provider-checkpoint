package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIpv6PimRps() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIpv6PimRps,
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
                        "expires": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "groups": {
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
                                },
                            },
                        },
                        "holdtime": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "group_amount": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "priority": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rp_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "type": {
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

func readGaiaShowIpv6PimRps(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-ipv6-pim-rps - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ipv6-pim-rps", payload)
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
            "ipv6-pim-rps",        // resource type
            "read",                       // operation
            "show-ipv6-pim-rps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ipv6-pim-rps: %v", err)
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
                        "expires": func() string { if _v, _ok := m["expires"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "groups": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["groups"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "address": func() string { if _v, _ok := _sgm["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "holdtime": func() int { if f, ok := m["holdtime"].(float64); ok { return int(f) }; return 0 }(),
                        "group_amount": func() int { if f, ok := m["group-amount"].(float64); ok { return int(f) }; return 0 }(),
                        "priority": func() int { if f, ok := m["priority"].(float64); ok { return int(f) }; return 0 }(),
                        "rp_address": func() string { if _v, _ok := m["rp-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-ipv6-pim-rps-" + acctest.RandString(10)))
    return nil
}

