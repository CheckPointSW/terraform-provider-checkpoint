package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBgpRoutemaps() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBgpRoutemaps,
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
                Description: `Sorts the output by the group type in either ascending or descending order. By default, the group types will be sorted in the order: confederation, external, internal. Within each group type, the items will be sorted according to the AS number in either ascending or descending order.`,
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
                        "as_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "as": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "default": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "peer": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "import_routemap_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "preference": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "family": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "export_routemap_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "preference": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "family": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "conditional_routemap": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "peer_list": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "peer": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "import_routemap_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "preference": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "family": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "export_routemap_list": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "preference": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "family": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "conditional_routemap": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
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

func readGaiaShowBgpRoutemaps(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-bgp-routemaps - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bgp-routemaps", payload)
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
            "bgp-routemaps",        // resource type
            "read",                       // operation
            "show-bgp-routemaps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bgp-routemaps: %v", err)
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
                        "as_type": func() string { if _v, _ok := m["as-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "as": func() string { if _v, _ok := m["as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "default": func() []interface{} {
                            if _obj, _ok := m["default"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "peer": func() string { if _v, _ok := _obj["peer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "import_routemap_list": func() []interface{} {
                                        if _d2, _ok := _obj["import-routemap-list"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "name": func() string { if _v, _ok := _d2["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "preference": func() int { if f, ok := _d2["preference"].(float64); ok { return int(f) }; return 0 }(),
                                                "family": func() string { if _v, _ok := _d2["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "export_routemap_list": func() []interface{} {
                                        if _d2, _ok := _obj["export-routemap-list"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "name": func() string { if _v, _ok := _d2["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "preference": func() int { if f, ok := _d2["preference"].(float64); ok { return int(f) }; return 0 }(),
                                                "family": func() string { if _v, _ok := _d2["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "conditional_routemap": func() string { if _v, _ok := _d2["conditional-routemap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                        "peer_list": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["peer-list"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "peer": func() string { if _v, _ok := _sgm["peer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "import_routemap_list": func() []interface{} {
                                                if _dobj, _ok := _sgm["import-routemap-list"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "name": func() string { if _v, _ok := _dobj["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "preference": func() int { if f, ok := _dobj["preference"].(float64); ok { return int(f) }; return 0 }(),
                                                        "family": func() string { if _v, _ok := _dobj["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                            "export_routemap_list": func() []interface{} {
                                                if _dobj, _ok := _sgm["export-routemap-list"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "name": func() string { if _v, _ok := _dobj["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "preference": func() int { if f, ok := _dobj["preference"].(float64); ok { return int(f) }; return 0 }(),
                                                        "family": func() string { if _v, _ok := _dobj["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "conditional_routemap": func() string { if _v, _ok := _dobj["conditional-routemap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
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
    d.SetId(fmt.Sprintf("show-bgp-routemaps-" + acctest.RandString(10)))
    return nil
}

