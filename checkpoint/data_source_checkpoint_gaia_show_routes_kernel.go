package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowRoutesKernel() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowRoutesKernel,
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
                Description: `Sorts the routes in either ascending or descending order.`,
            },
            "address_family": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Address family of routes returned. IPv6 route monitoring, or specifying \"inet6\" for this field, is only supported on GAIA versions R81.10 and up.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "objects": {
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
                        "mask_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "protocol": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "next_hop": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "gateways": {
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
                                                "interface": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "interface": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "age": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "active_age": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "cost": {
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

func readGaiaShowRoutesKernel(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOk("address_family"); ok {
        payload["address-family"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-routes-kernel - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-routes-kernel", payload)
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
            "routes-kernel",        // resource type
            "read",                       // operation
            "show-routes-kernel",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-routes-kernel: %v", err)
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
                        "address": func() string { if _v, _ok := m["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "mask_length": func() int { if f, ok := m["mask-length"].(float64); ok { return int(f) }; return 0 }(),
                        "protocol": func() string { if _v, _ok := m["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "next_hop": func() []interface{} {
                            if _obj, _ok := m["next-hop"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "gateways": func() []interface{} {
                                        if _gws, _ok := _obj["gateways"].([]interface{}); _ok {
                                            _gwOut := make([]interface{}, len(_gws))
                                            for _gi, _gw := range _gws {
                                                if _gwm, _ok := _gw.(map[string]interface{}); _ok {
                                                    _gwOut[_gi] = map[string]interface{}{
                                                        "address": func() string { if _v, _ok := _gwm["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "interface": func() string { if _v, _ok := _gwm["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }
                                                }
                                            }
                                            return _gwOut
                                        }
                                        return nil
                                    }(),
                                    "interface": func() string { if _v, _ok := _obj["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "age": func() int { if f, ok := m["age"].(float64); ok { return int(f) }; return 0 }(),
                        "active_age": func() int { if f, ok := m["active-age"].(float64); ok { return int(f) }; return 0 }(),
                        "cost": func() string { if _v, _ok := m["cost"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-routes-kernel-" + acctest.RandString(10)))
    return nil
}

