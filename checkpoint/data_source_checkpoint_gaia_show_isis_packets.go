package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisPackets() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisPackets,
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
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "sent": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "unknown": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "hello": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "csnp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "psnp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "dropped": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "unknown": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "hello": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "csnp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "psnp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "received": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "unknown": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "hello": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "csnp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "psnp": {
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

func readGaiaShowIsisPackets(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-isis-packets - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-packets", payload)
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
            "isis-packets",        // resource type
            "read",                       // operation
            "show-isis-packets",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-packets: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["sent"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("sent", []interface{}{map[string]interface{}{
                "unknown": func() int { if f, ok := _m["unknown"].(float64); ok { return int(f) }; return 0 }(),
                "lsp": func() int { if f, ok := _m["lsp"].(float64); ok { return int(f) }; return 0 }(),
                "hello": func() int { if f, ok := _m["hello"].(float64); ok { return int(f) }; return 0 }(),
                "csnp": func() int { if f, ok := _m["csnp"].(float64); ok { return int(f) }; return 0 }(),
                "psnp": func() int { if f, ok := _m["psnp"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["dropped"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("dropped", []interface{}{map[string]interface{}{
                "unknown": func() int { if f, ok := _m["unknown"].(float64); ok { return int(f) }; return 0 }(),
                "lsp": func() int { if f, ok := _m["lsp"].(float64); ok { return int(f) }; return 0 }(),
                "hello": func() int { if f, ok := _m["hello"].(float64); ok { return int(f) }; return 0 }(),
                "csnp": func() int { if f, ok := _m["csnp"].(float64); ok { return int(f) }; return 0 }(),
                "psnp": func() int { if f, ok := _m["psnp"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["received"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("received", []interface{}{map[string]interface{}{
                "unknown": func() int { if f, ok := _m["unknown"].(float64); ok { return int(f) }; return 0 }(),
                "lsp": func() int { if f, ok := _m["lsp"].(float64); ok { return int(f) }; return 0 }(),
                "hello": func() int { if f, ok := _m["hello"].(float64); ok { return int(f) }; return 0 }(),
                "csnp": func() int { if f, ok := _m["csnp"].(float64); ok { return int(f) }; return 0 }(),
                "psnp": func() int { if f, ok := _m["psnp"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-isis-packets-" + acctest.RandString(10)))
    return nil
}

