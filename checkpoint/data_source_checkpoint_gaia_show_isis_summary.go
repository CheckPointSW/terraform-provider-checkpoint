package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisSummary() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisSummary,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The instance to be queried`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "instance": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "system_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "is_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "mt_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "state": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "shutdown_timer": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "last_restart": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "manual_area_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "configured_interfaces_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "authentication": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "auth_mode": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "auth_ignore": {
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
    }
}

func readGaiaShowIsisSummary(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-isis-summary - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-summary", payload)
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
            "isis-summary",        // resource type
            "read",                       // operation
            "show-isis-summary",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-summary: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["instance"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("instance", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["system-id"]; exists {
        d.Set("system_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["is-type"]; exists {
        d.Set("is_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["mt-type"]; exists {
        d.Set("mt_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["shutdown-timer"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("shutdown_timer", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["last-restart"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("last_restart", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["manual-area-list"]; exists {
        d.Set("manual_area_list", v.([]interface{}))
    } else {
        d.Set("manual_area_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["configured-interfaces-list"]; exists {
        d.Set("configured_interfaces_list", v.([]interface{}))
    } else {
        d.Set("configured_interfaces_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["authentication"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "auth_mode": func() string { if _v, _ok := m["auth-mode"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "auth_ignore": func() []interface{} {
                            switch _ev := m["auth-ignore"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                    }
                }
            }
            d.Set("authentication", mapped)
        }
    } else {
        d.Set("authentication", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-isis-summary-" + acctest.RandString(10)))
    return nil
}

