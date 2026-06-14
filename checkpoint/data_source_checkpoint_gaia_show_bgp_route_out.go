package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBgpRouteOut() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBgpRouteOut,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Filter the results for a specific route address.`,
            },
            "mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Filter the results for a specific route mask-length.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "object": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "route": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "as_path": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "communities": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "local_preference": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "med": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "nexthop": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "link_local_nexthop": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "origin": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "originator_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "extended_communities": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "value": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "sub_type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "sub_type_description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "peer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "peer_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "peer_as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowBgpRouteOut(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bgp-route-out - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bgp-route-out", payload)
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
            "bgp-route-out",        // resource type
            "read",                       // operation
            "show-bgp-route-out",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bgp-route-out: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["object"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "route": func() string { if _v, _ok := m["route"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "as_path": func() string { if _v, _ok := m["as-path"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "communities": func() string { if _v, _ok := m["communities"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "local_preference": func() string { if _v, _ok := m["local-preference"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "med": func() string { if _v, _ok := m["med"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "nexthop": func() string { if _v, _ok := m["nexthop"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "link_local_nexthop": func() string { if _v, _ok := m["link-local-nexthop"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "origin": func() string { if _v, _ok := m["origin"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "originator_id": func() string { if _v, _ok := m["originator-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "extended_communities": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["extended-communities"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "value": func() string { if _v, _ok := _sgm["value"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "type": func() string { if _v, _ok := _sgm["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "sub_type": func() string { if _v, _ok := _sgm["sub-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "sub_type_description": func() string { if _v, _ok := _sgm["sub-type-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "peer": func() string { if _v, _ok := m["peer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "peer_type": func() string { if _v, _ok := m["peer-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "peer_as": func() string { if _v, _ok := m["peer-as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("object", mapped)
        }
    } else {
        d.Set("object", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bgp-route-out-" + acctest.RandString(10)))
    return nil
}

