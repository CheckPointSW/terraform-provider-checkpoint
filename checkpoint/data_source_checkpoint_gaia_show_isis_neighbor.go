package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisNeighbor() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisNeighbor,
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
            "neighbor": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The system-id of the neighbor to be queried`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "system_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "interface": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "hold_time": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "last_state_change": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "level": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "state": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "restart_capable": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "restarting_now": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "last_restart": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "circuit_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "is_dis": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "priority": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "mac_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "area_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "mtid_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "bfd_support_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "bfd_sessions_list": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "topology": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
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

func readGaiaShowIsisNeighbor(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("neighbor"); ok {
        payload["neighbor"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-isis-neighbor - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-neighbor", payload)
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
            "isis-neighbor",        // resource type
            "read",                       // operation
            "show-isis-neighbor",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-neighbor: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["system-id"]; exists {
        d.Set("system_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["hold-time"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("hold_time", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["last-state-change"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("last_state_change", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["level"]; exists {
        d.Set("level", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["restart-capable"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("restart_capable", b)
        } else if s, ok := v.(string); ok {
            d.Set("restart_capable", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["restarting-now"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("restarting_now", b)
        } else if s, ok := v.(string); ok {
            d.Set("restarting_now", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["last-restart"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("last_restart", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["circuit-id"]; exists {
        d.Set("circuit_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["is-dis"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("is_dis", b)
        } else if s, ok := v.(string); ok {
            d.Set("is_dis", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["priority"]; exists {
        d.Set("priority", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ipv4-address"]; exists {
        d.Set("ipv4_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["mac-address"]; exists {
        d.Set("mac_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["area-list"]; exists {
        d.Set("area_list", v.([]interface{}))
    } else {
        d.Set("area_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["mtid-list"]; exists {
        d.Set("mtid_list", v.([]interface{}))
    } else {
        d.Set("mtid_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["bfd-support-list"]; exists {
        d.Set("bfd_support_list", v.([]interface{}))
    } else {
        d.Set("bfd_support_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["bfd-sessions-list"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "topology": func() string { if _v, _ok := m["topology"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("bfd_sessions_list", mapped)
        }
    } else {
        d.Set("bfd_sessions_list", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-isis-neighbor-" + acctest.RandString(10)))
    return nil
}

