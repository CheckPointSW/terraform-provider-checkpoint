package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowRoutemaps() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowRoutemaps,
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
                Description: `Sorts the routes by priority in either ascending or descending order.`,
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
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ids": {
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
                                    "state": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "match": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "aspath_regex": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "community_regex": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "extcommunity_regex": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "protocol": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "as": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "community": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "extended_community": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "ifaddress": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "interface": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "level": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "metric_type": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "neighbor": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "network": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "nexthop": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "ospf_instance": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "prefix_list": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "prefix_tree": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "route_type": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "tag": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeInt,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    "action": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "aspath_prepend_count": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "localpref": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "metric_type": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "ospfautomatictag": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "ospfmanualtag": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "precedence": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "preference": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "prefix_list": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "riptag": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "route_type": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "community": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "extended_community": {
                                                    Type:        schema.TypeSet,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "nexthop": {
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

func readGaiaShowRoutemaps(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-routemaps - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-routemaps", payload)
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
            "routemaps",        // resource type
            "read",                       // operation
            "show-routemaps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-routemaps: %v", err)
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
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ids": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ids"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "resource_id": func() int { if f, ok := _sgm["id"].(float64); ok { return int(f) }; return 0 }(),
                                            "state": func() string { if _v, _ok := _sgm["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "match": func() []interface{} {
                                                if _dobj, _ok := _sgm["match"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "aspath_regex": func() string { if _v, _ok := _dobj["aspath-regex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "community_regex": func() string { if _v, _ok := _dobj["community-regex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "extcommunity_regex": func() string { if _v, _ok := _dobj["extcommunity-regex"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "metric": func() int { if f, ok := _dobj["metric"].(float64); ok { return int(f) }; return 0 }(),
                                                        "protocol": func() string { if _v, _ok := _dobj["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "as": func() string { if _v, _ok := _dobj["as"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "community": func() string { if _v, _ok := _dobj["community"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "extended_community": func() string { if _v, _ok := _dobj["extended-community"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "ifaddress": func() string { if _v, _ok := _dobj["ifaddress"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "interface": func() string { if _v, _ok := _dobj["interface"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "level": func() string { if _v, _ok := _dobj["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "metric_type": func() string { if _v, _ok := _dobj["metric-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "neighbor": func() string { if _v, _ok := _dobj["neighbor"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "network": func() string { if _v, _ok := _dobj["network"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "nexthop": func() string { if _v, _ok := _dobj["nexthop"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "ospf_instance": func() string { if _v, _ok := _dobj["ospf-instance"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "prefix_list": func() string { if _v, _ok := _dobj["prefix-list"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "prefix_tree": func() string { if _v, _ok := _dobj["prefix-tree"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "route_type": func() string { if _v, _ok := _dobj["route-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "tag": func() string { if _v, _ok := _dobj["tag"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                            "action": func() []interface{} {
                                                if _dobj, _ok := _sgm["action"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "aspath_prepend_count": func() int { if f, ok := _dobj["aspath-prepend-count"].(float64); ok { return int(f) }; return 0 }(),
                                                        "localpref": func() int { if f, ok := _dobj["localpref"].(float64); ok { return int(f) }; return 0 }(),
                                                        "metric": func() string { if _v, _ok := _dobj["metric"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "metric_type": func() string { if _v, _ok := _dobj["metric-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "ospfautomatictag": func() int { if f, ok := _dobj["ospfautomatictag"].(float64); ok { return int(f) }; return 0 }(),
                                                        "ospfmanualtag": func() int { if f, ok := _dobj["ospfmanualtag"].(float64); ok { return int(f) }; return 0 }(),
                                                        "precedence": func() int { if f, ok := _dobj["precedence"].(float64); ok { return int(f) }; return 0 }(),
                                                        "preference": func() int { if f, ok := _dobj["preference"].(float64); ok { return int(f) }; return 0 }(),
                                                        "prefix_list": func() string { if _v, _ok := _dobj["prefix-list"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "riptag": func() int { if f, ok := _dobj["riptag"].(float64); ok { return int(f) }; return 0 }(),
                                                        "route_type": func() string { if _v, _ok := _dobj["route-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "community": func() string { if _v, _ok := _dobj["community"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "extended_community": func() string { if _v, _ok := _dobj["extended-community"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                        "nexthop": func() string { if _v, _ok := _dobj["nexthop"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-routemaps-" + acctest.RandString(10)))
    return nil
}

