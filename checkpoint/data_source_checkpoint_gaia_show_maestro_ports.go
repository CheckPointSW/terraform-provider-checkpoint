package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowMaestroPorts() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowMaestroPorts,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "ports": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface_name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "qsfp_mode": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "link_state": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "auto_negotiation": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transceiver_state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "operating_speed": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mtu": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rx_frames": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tx_frames": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "orchestrator_id": {
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
        },
    }
}

func readGaiaShowMaestroPorts(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    log.Println("Execute show-maestro-ports - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-maestro-ports", payload)
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
            "maestro-ports",        // resource type
            "read",                       // operation
            "show-maestro-ports",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-maestro-ports: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["ports"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "resource_id": func() string { if _v, _ok := m["id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site": func() int { if f, ok := m["site"].(float64); ok { return int(f) }; return 0 }(),
                        "interface_name": func() string { if _v, _ok := m["interface-name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "qsfp_mode": func() string { if _v, _ok := m["qsfp-mode"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enabled": func() bool { if b, ok := m["enabled"].(bool); ok { return b }; if s, ok := m["enabled"].(string); ok { return s == "true" }; return false }(),
                        "link_state": func() bool { if b, ok := m["link-state"].(bool); ok { return b }; if s, ok := m["link-state"].(string); ok { return s == "true" }; return false }(),
                        "auto_negotiation": func() bool { if b, ok := m["auto-negotiation"].(bool); ok { return b }; if s, ok := m["auto-negotiation"].(string); ok { return s == "true" }; return false }(),
                        "transceiver_state": func() string { if _v, _ok := m["transceiver-state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "operating_speed": func() string { if _v, _ok := m["operating-speed"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "mtu": func() int { if f, ok := m["mtu"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_frames": func() int { if f, ok := m["rx-frames"].(float64); ok { return int(f) }; return 0 }(),
                        "tx_frames": func() int { if f, ok := m["tx-frames"].(float64); ok { return int(f) }; return 0 }(),
                        "orchestrator_id": func() []interface{} {
                            if _arr, _ok := m["orchestrator-id"].([]interface{}); _ok {
                                out := make([]interface{}, len(_arr))
                                for _i, _v := range _arr {
                                    if _f, _ok := _v.(float64); _ok { out[_i] = int(_f) }
                                }
                                return out
                            }
                            return []interface{}{}
                        }(),
                    }
                }
            }
            d.Set("ports", mapped)
        }
    } else {
        d.Set("ports", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-maestro-ports-" + acctest.RandString(10)))
    return nil
}

