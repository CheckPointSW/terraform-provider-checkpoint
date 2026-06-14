package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPhysicalInterfaceXcvr() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPhysicalInterfaceXcvr,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
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
    }
}

func readGaiaShowPhysicalInterfaceXcvr(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    log.Println("Execute show-physical-interface-xcvr - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-physical-interface-xcvr", payload)
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
            "physical-interface-xcvr",        // resource type
            "read",                       // operation
            "show-physical-interface-xcvr",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-physical-interface-xcvr: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["chkp-xcvr"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("chkp_xcvr", b)
        } else if s, ok := v.(string); ok {
            d.Set("chkp_xcvr", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["temp"]; exists {
        d.Set("temp", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage"]; exists {
        d.Set("voltage", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["laser-bias-current"]; exists {
        d.Set("laser_bias_current", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["transmit-power"]; exists {
        d.Set("transmit_power", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rec-power"]; exists {
        d.Set("rec_power", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["los"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("los", b)
        } else if s, ok := v.(string); ok {
            d.Set("los", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["transmitter-fault"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("transmitter_fault", b)
        } else if s, ok := v.(string); ok {
            d.Set("transmitter_fault", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["laser-bias-current-ch-1"]; exists {
        d.Set("laser_bias_current_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["laser-bias-current-ch-2"]; exists {
        d.Set("laser_bias_current_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["laser-bias-current-ch-3"]; exists {
        d.Set("laser_bias_current_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["laser-bias-current-ch-4"]; exists {
        d.Set("laser_bias_current_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["transmit-power-ch-1"]; exists {
        d.Set("transmit_power_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["transmit-power-ch-2"]; exists {
        d.Set("transmit_power_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["transmit-power-ch-3"]; exists {
        d.Set("transmit_power_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["transmit-power-ch-4"]; exists {
        d.Set("transmit_power_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rec-power-ch-1"]; exists {
        d.Set("rec_power_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rec-power-ch-2"]; exists {
        d.Set("rec_power_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rec-power-ch-3"]; exists {
        d.Set("rec_power_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rec-power-ch-4"]; exists {
        d.Set("rec_power_ch_4", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-physical-interface-xcvr-" + acctest.RandString(10)))
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    return nil
}

