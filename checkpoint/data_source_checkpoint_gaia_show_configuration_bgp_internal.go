package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpInternal() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpInternal,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "description": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
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
                                    "condition": {
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
            "interface_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "local_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "med": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_nexthop_self": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "outdelay": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "protocol_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
        },
    }
}

func readGaiaShowConfigurationBgpInternal(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-internal - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-internal", payload)
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
            "configuration-bgp-internal",        // resource type
            "read",                       // operation
            "show-configuration-bgp-internal",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-internal: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["description"]; exists {
        d.Set("description", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["export-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "conditional_routemap": func() []interface{} {
                            if _obj, _ok := m["conditional-routemap"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "name": func() string { if _v, _ok := _obj["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "condition": func() string { if _v, _ok := _obj["condition"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                    }
                }
            }
            d.Set("export_routemap_list", mapped)
        }
    } else {
        d.Set("export_routemap_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["import-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("import_routemap_list", mapped)
        }
    } else {
        d.Set("import_routemap_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["interface-list"]; exists {
        switch val := v.(type) {
        case []interface{}:
            d.Set("interface_list", val)
        case string:
            d.Set("interface_list", []interface{}{val})
        default:
            d.Set("interface_list", []interface{}{})
        }
    } else {
        d.Set("interface_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["med"]; exists {
        d.Set("med", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-nexthop-self"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_nexthop_self", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_nexthop_self", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["outdelay"]; exists {
        d.Set("outdelay", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["protocol-list"]; exists {
        switch val := v.(type) {
        case []interface{}:
            d.Set("protocol_list", val)
        case string:
            d.Set("protocol_list", []interface{}{val})
        default:
            d.Set("protocol_list", []interface{}{})
        }
    } else {
        d.Set("protocol_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-internal-" + acctest.RandString(10)))
    return nil
}

