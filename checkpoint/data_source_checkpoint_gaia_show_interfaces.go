package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowInterfaces() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowInterfaces,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "names": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv4_mask_length": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_autoconfig": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "comments": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_mask_length": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_local_link_address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "virtual_system_id": {
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

func readGaiaShowInterfaces(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("names"); len(v.(*schema.Set).List()) > 0 {
        payload["names"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-interfaces - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-interfaces", payload)
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
            "interfaces",        // resource type
            "read",                       // operation
            "show-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-interfaces: %v", err)
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
                        "ipv4_address": func() string { if _v, _ok := m["ipv4-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv4_mask_length": func() string { if _v, _ok := m["ipv4-mask-length"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enabled": func() bool { if b, ok := m["enabled"].(bool); ok { return b }; if s, ok := m["enabled"].(string); ok { return s == "true" }; return false }(),
                        "ipv6_autoconfig": func() bool { if b, ok := m["ipv6-autoconfig"].(bool); ok { return b }; if s, ok := m["ipv6-autoconfig"].(string); ok { return s == "true" }; return false }(),
                        "comments": func() string { if _v, _ok := m["comments"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_address": func() string { if _v, _ok := m["ipv6-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_mask_length": func() string { if _v, _ok := m["ipv6-mask-length"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "ipv6_local_link_address": func() string { if _v, _ok := m["ipv6-local-link-address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "member_id": func() string { if _v, _ok := m["member-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "virtual_system_id": func() string { if _v, _ok := m["virtual-system-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-interfaces-" + acctest.RandString(10)))
    return nil
}

