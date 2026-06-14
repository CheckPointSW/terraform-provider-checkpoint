package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowDynamicLayers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowDynamicLayers,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "layers": {
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
                        "meta_info": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "last_modifier": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "last_modify_time": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "iso": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "posix": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "comments": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "tags": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "custom_fields": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "field_1": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "field_2": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "field_3": {
                                                    Type:        schema.TypeString,
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
                },
            },
        },
    }
}

func readGaiaShowDynamicLayers(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-dynamic-layers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-dynamic-layers", payload)
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
            "dynamic-layers",        // resource type
            "read",                       // operation
            "show-dynamic-layers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-dynamic-layers: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["layers"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "meta_info": func() []interface{} {
                            if _obj, _ok := m["meta-info"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "last_modifier": func() string { if _v, _ok := _obj["last-modifier"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "last_modify_time": func() []interface{} {
                                        if _d2, _ok := _obj["last-modify-time"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "iso": func() string { if _v, _ok := _d2["iso"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "posix": func() int { if f, ok := _d2["posix"].(float64); ok { return int(f) }; return 0 }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "comments": func() string { if _v, _ok := _obj["comments"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "tags": func() []interface{} {
                                        if _sl, _ok := _obj["tags"].([]interface{}); _ok {
                                            return _sl
                                        }
                                        return nil
                                    }(),
                                    "custom_fields": func() []interface{} {
                                        if _d2, _ok := _obj["custom-fields"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "field_1": func() string { if _v, _ok := _d2["field-1"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "field_2": func() string { if _v, _ok := _d2["field-2"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                "field_3": func() string { if _v, _ok := _d2["field-3"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                    }
                }
            }
            d.Set("layers", mapped)
        }
    } else {
        d.Set("layers", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-dynamic-layers-" + acctest.RandString(10)))
    return nil
}

