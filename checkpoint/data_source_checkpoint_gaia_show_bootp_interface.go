package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBootpInterface() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBootpInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Interface name`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "max_hopcount": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "wait_time": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "primary": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "gateway_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "relay_to": {
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
                        "status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "option_82_enabled": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "circuit_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "remote_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowBootpInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bootp-interface - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bootp-interface", payload)
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
            "bootp-interface",        // resource type
            "read",                       // operation
            "show-bootp-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bootp-interface: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["max-hopcount"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("max_hopcount", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["wait-time"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("wait_time", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["primary"]; exists {
        d.Set("primary", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["gateway-address"]; exists {
        d.Set("gateway_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["relay-to"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "address": func() string { if _v, _ok := m["address"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "status": func() string { if _v, _ok := m["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("relay_to", mapped)
        }
    } else {
        d.Set("relay_to", []interface{}{})
    }
    if v, exists := commandRes.GetData()["option-82-enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("option_82_enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("option_82_enabled", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["circuit-id"]; exists {
        d.Set("circuit_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["remote-id"]; exists {
        d.Set("remote_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["status"]; exists {
        d.Set("status", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bootp-interface-" + acctest.RandString(10)))
    return nil
}

