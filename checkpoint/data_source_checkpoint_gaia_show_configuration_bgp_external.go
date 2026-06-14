package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpExternal() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpExternal,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "remote_as": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The Autonomous System number of the peer group.The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
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
            "inject_routemap_list": {
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
                        "any_pass_routemap": {
                            Type:        schema.TypeString,
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
            "local_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "outdelay": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowConfigurationBgpExternal(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-external - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-external", payload)
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
            "configuration-bgp-external",        // resource type
            "read",                       // operation
            "show-configuration-bgp-external",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-external: %v", err)
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
    if v, exists := commandRes.GetData()["inject-routemap-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "preference": func() int { if f, ok := m["preference"].(float64); ok { return int(f) }; return 0 }(),
                        "any_pass_routemap": func() string { if _v, _ok := m["any-pass-routemap"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "family": func() string { if _v, _ok := m["family"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("inject_routemap_list", mapped)
        }
    } else {
        d.Set("inject_routemap_list", []interface{}{})
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
    if v, exists := commandRes.GetData()["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["outdelay"]; exists {
        d.Set("outdelay", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["remote-as"]; exists {
        d.Set("remote_as", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-external-" + acctest.RandString(10)))
    return nil
}

