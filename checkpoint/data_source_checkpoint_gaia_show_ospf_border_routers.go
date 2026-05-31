package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfBorderRouters() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfBorderRouters,
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
                Description: `Sorts the results in either ascending or descending order.`,
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Existing OSPFv2 Instance`,
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
                        "area_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "cost": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "path_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "protocol_instance": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "router_id": {
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

func readGaiaShowOspfBorderRouters(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-border-routers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-border-routers", payload)
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
            "ospf-border-routers",        // resource type
            "read",                       // operation
            "show-ospf-border-routers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-border-routers: %v", err)
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
                        "area_id": func() string { if _v, _ok := m["area-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "cost": func() int { if f, ok := m["cost"].(float64); ok { return int(f) }; return 0 }(),
                        "path_type": func() string { if _v, _ok := m["path-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "protocol_instance": func() int { if f, ok := m["protocol-instance"].(float64); ok { return int(f) }; return 0 }(),
                        "router_id": func() string { if _v, _ok := m["router-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-ospf-border-routers-" + acctest.RandString(10)))
    return nil
}

