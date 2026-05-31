package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowVirtualGateways() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowVirtualGateways,
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
                Description: `Sorts the virtual gateways by either ascending or descending order.`,
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
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interfaces": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "virtual_switches": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "resource_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "connected_through": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "resources": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "firewall_ipv4_instances": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "firewall_ipv6_instances": {
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

func readGaiaShowVirtualGateways(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-virtual-gateways - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-virtual-gateways", payload)
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
            "virtual-gateways",        // resource type
            "read",                       // operation
            "show-virtual-gateways",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-virtual-gateways: %v", err)
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
                        "resource_id": func() int { if f, ok := m["id"].(float64); ok { return int(f) }; return 0 }(),
                        "interfaces": func() []interface{} {
                            switch _ev := m["interfaces"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "virtual_switches": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["virtual-switches"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "resource_id": func() int { if f, ok := _sgm["id"].(float64); ok { return int(f) }; return 0 }(),
                                            "connected_through": func() string { if _v, _ok := _sgm["connected-through"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "resources": func() []interface{} {
                            if _obj, _ok := m["resources"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "firewall_ipv4_instances": func() int { if f, ok := _obj["firewall-ipv4-instances"].(float64); ok { return int(f) }; return 0 }(),
                                    "firewall_ipv6_instances": func() int { if f, ok := _obj["firewall-ipv6-instances"].(float64); ok { return int(f) }; return 0 }(),
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
    d.SetId(fmt.Sprintf("show-virtual-gateways-" + acctest.RandString(10)))
    return nil
}

