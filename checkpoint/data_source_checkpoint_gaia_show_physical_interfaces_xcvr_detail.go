package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPhysicalInterfacesXcvrDetail() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPhysicalInterfacesXcvrDetail,
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
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "chkp_xcvr": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "temp": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "voltage": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "laser_bias_current": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit_power": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rec_power": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "los": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmitter_fault": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "laser_bias_current_ch_1": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "laser_bias_current_ch_2": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "laser_bias_current_ch_3": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "laser_bias_current_ch_4": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit_power_ch_1": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit_power_ch_2": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit_power_ch_3": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit_power_ch_4": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rec_power_ch_1": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rec_power_ch_2": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rec_power_ch_3": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rec_power_ch_4": {
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

func readGaiaShowPhysicalInterfacesXcvrDetail(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-physical-interfaces-xcvr-detail - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-physical-interfaces-xcvr-detail", payload)
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
            "physical-interfaces-xcvr-detail",        // resource type
            "read",                       // operation
            "show-physical-interfaces-xcvr-detail",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-physical-interfaces-xcvr-detail: %v", err)
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
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "chkp_xcvr": func() bool { if b, ok := m["chkp-xcvr"].(bool); ok { return b }; if s, ok := m["chkp-xcvr"].(string); ok { return s == "true" }; return false }(),
                        "temp": func() string { if _v, _ok := m["temp"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "voltage": func() string { if _v, _ok := m["voltage"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "laser_bias_current": func() string { if _v, _ok := m["laser-bias-current"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "transmit_power": func() string { if _v, _ok := m["transmit-power"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rec_power": func() string { if _v, _ok := m["rec-power"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "los": func() bool { if b, ok := m["los"].(bool); ok { return b }; if s, ok := m["los"].(string); ok { return s == "true" }; return false }(),
                        "transmitter_fault": func() bool { if b, ok := m["transmitter-fault"].(bool); ok { return b }; if s, ok := m["transmitter-fault"].(string); ok { return s == "true" }; return false }(),
                        "laser_bias_current_ch_1": func() string { if _v, _ok := m["laser-bias-current-ch-1"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "laser_bias_current_ch_2": func() string { if _v, _ok := m["laser-bias-current-ch-2"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "laser_bias_current_ch_3": func() string { if _v, _ok := m["laser-bias-current-ch-3"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "laser_bias_current_ch_4": func() string { if _v, _ok := m["laser-bias-current-ch-4"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "transmit_power_ch_1": func() string { if _v, _ok := m["transmit-power-ch-1"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "transmit_power_ch_2": func() string { if _v, _ok := m["transmit-power-ch-2"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "transmit_power_ch_3": func() string { if _v, _ok := m["transmit-power-ch-3"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "transmit_power_ch_4": func() string { if _v, _ok := m["transmit-power-ch-4"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rec_power_ch_1": func() string { if _v, _ok := m["rec-power-ch-1"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rec_power_ch_2": func() string { if _v, _ok := m["rec-power-ch-2"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rec_power_ch_3": func() string { if _v, _ok := m["rec-power-ch-3"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "rec_power_ch_4": func() string { if _v, _ok := m["rec-power-ch-4"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
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
    d.SetId(fmt.Sprintf("show-physical-interfaces-xcvr-detail-" + acctest.RandString(10)))
    return nil
}

