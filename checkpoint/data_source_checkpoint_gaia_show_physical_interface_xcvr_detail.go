package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPhysicalInterfaceXcvrDetail() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPhysicalInterfaceXcvrDetail,
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
            "temperature": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "product_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vendor_name": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vendor_pn": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vendor_rev": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vendor_sn": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "cp_part_number": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "cp_material_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "laser_wavelength": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_smf_km": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_smf": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_50um": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_62_5um": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_copper": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_om3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_om3_50um": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_om2_50um": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_om1_62_5um": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "link_length_copper_active": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "voltage": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "temperature_alarm_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "temperature_alarm_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "temperature_warning_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "temperature_warning_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "voltage_alarm_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "voltage_alarm_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "voltage_warning_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "voltage_warning_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_high_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_low_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_high_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_low_ch_1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_high_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_low_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_high_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_low_ch_2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_high_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_low_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_high_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_low_ch_3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_high_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_alarm_low_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_high_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "current_warning_low_ch_4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_high_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_low_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_high_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_low_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_high_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_low_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_high_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_low_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_high_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_low_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_high_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_low_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_high_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_alarm_low_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_high_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_power_warning_low_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_high": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_low": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_high_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_low_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_high_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_low_ch1": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_high_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_low_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_high_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_low_ch2": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_high_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_low_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_high_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_low_ch3": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_high_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_alarm_low_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_high_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_power_warning_low_ch4": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowPhysicalInterfaceXcvrDetail(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute show-physical-interface-xcvr-detail - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-physical-interface-xcvr-detail", payload)
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
            "physical-interface-xcvr-detail",        // resource type
            "read",                       // operation
            "show-physical-interface-xcvr-detail",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-physical-interface-xcvr-detail: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["temperature"]; exists {
        d.Set("temperature", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["product-type"]; exists {
        d.Set("product_type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["vendor-name"]; exists {
        d.Set("vendor_name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["vendor-pn"]; exists {
        d.Set("vendor_pn", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["vendor-rev"]; exists {
        d.Set("vendor_rev", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["vendor-sn"]; exists {
        d.Set("vendor_sn", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["CP-part-number"]; exists {
        d.Set("cp_part_number", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["CP-material-id"]; exists {
        d.Set("cp_material_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["laser-wavelength"]; exists {
        d.Set("laser_wavelength", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-smf-km"]; exists {
        d.Set("link_length_smf_km", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-smf"]; exists {
        d.Set("link_length_smf", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-50um"]; exists {
        d.Set("link_length_50um", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-62.5um"]; exists {
        d.Set("link_length_62_5um", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-copper"]; exists {
        d.Set("link_length_copper", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-om3"]; exists {
        d.Set("link_length_om3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-om3-50um"]; exists {
        d.Set("link_length_om3_50um", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-om2-50um"]; exists {
        d.Set("link_length_om2_50um", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-om1-62.5um"]; exists {
        d.Set("link_length_om1_62_5um", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["link-length-copper-active"]; exists {
        d.Set("link_length_copper_active", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage"]; exists {
        d.Set("voltage", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current"]; exists {
        d.Set("current", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-ch-1"]; exists {
        d.Set("current_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-ch-2"]; exists {
        d.Set("current_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-ch-3"]; exists {
        d.Set("current_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-ch-4"]; exists {
        d.Set("current_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power"]; exists {
        d.Set("tx_power", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-ch-1"]; exists {
        d.Set("tx_power_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-ch-2"]; exists {
        d.Set("tx_power_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-ch-3"]; exists {
        d.Set("tx_power_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-ch-4"]; exists {
        d.Set("tx_power_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power"]; exists {
        d.Set("rx_power", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-ch-1"]; exists {
        d.Set("rx_power_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-ch-2"]; exists {
        d.Set("rx_power_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-ch-3"]; exists {
        d.Set("rx_power_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-ch-4"]; exists {
        d.Set("rx_power_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["temperature-alarm-high"]; exists {
        d.Set("temperature_alarm_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["temperature-alarm-low"]; exists {
        d.Set("temperature_alarm_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["temperature-warning-high"]; exists {
        d.Set("temperature_warning_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["temperature-warning-low"]; exists {
        d.Set("temperature_warning_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage-alarm-high"]; exists {
        d.Set("voltage_alarm_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage-alarm-low"]; exists {
        d.Set("voltage_alarm_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage-warning-high"]; exists {
        d.Set("voltage_warning_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["voltage-warning-low"]; exists {
        d.Set("voltage_warning_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-high"]; exists {
        d.Set("current_alarm_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-low"]; exists {
        d.Set("current_alarm_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-high"]; exists {
        d.Set("current_warning_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-low"]; exists {
        d.Set("current_warning_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-high-ch-1"]; exists {
        d.Set("current_alarm_high_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-low-ch-1"]; exists {
        d.Set("current_alarm_low_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-high-ch-1"]; exists {
        d.Set("current_warning_high_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-low-ch-1"]; exists {
        d.Set("current_warning_low_ch_1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-high-ch-2"]; exists {
        d.Set("current_alarm_high_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-low-ch-2"]; exists {
        d.Set("current_alarm_low_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-high-ch-2"]; exists {
        d.Set("current_warning_high_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-low-ch-2"]; exists {
        d.Set("current_warning_low_ch_2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-high-ch-3"]; exists {
        d.Set("current_alarm_high_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-low-ch-3"]; exists {
        d.Set("current_alarm_low_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-high-ch-3"]; exists {
        d.Set("current_warning_high_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-low-ch-3"]; exists {
        d.Set("current_warning_low_ch_3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-high-ch-4"]; exists {
        d.Set("current_alarm_high_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-alarm-low-ch-4"]; exists {
        d.Set("current_alarm_low_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-high-ch-4"]; exists {
        d.Set("current_warning_high_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["current-warning-low-ch-4"]; exists {
        d.Set("current_warning_low_ch_4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-high"]; exists {
        d.Set("tx_power_alarm_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-low"]; exists {
        d.Set("tx_power_alarm_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-high"]; exists {
        d.Set("tx_power_warning_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-low"]; exists {
        d.Set("tx_power_warning_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-high-ch1"]; exists {
        d.Set("tx_power_alarm_high_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-low-ch1"]; exists {
        d.Set("tx_power_alarm_low_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-high-ch1"]; exists {
        d.Set("tx_power_warning_high_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-low-ch1"]; exists {
        d.Set("tx_power_warning_low_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-high-ch2"]; exists {
        d.Set("tx_power_alarm_high_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-low-ch2"]; exists {
        d.Set("tx_power_alarm_low_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-high-ch2"]; exists {
        d.Set("tx_power_warning_high_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-low-ch2"]; exists {
        d.Set("tx_power_warning_low_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-high-ch3"]; exists {
        d.Set("tx_power_alarm_high_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-low-ch3"]; exists {
        d.Set("tx_power_alarm_low_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-high-ch3"]; exists {
        d.Set("tx_power_warning_high_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-low-ch3"]; exists {
        d.Set("tx_power_warning_low_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-high-ch4"]; exists {
        d.Set("tx_power_alarm_high_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-alarm-low-ch4"]; exists {
        d.Set("tx_power_alarm_low_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-high-ch4"]; exists {
        d.Set("tx_power_warning_high_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["tx-power-warning-low-ch4"]; exists {
        d.Set("tx_power_warning_low_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-high"]; exists {
        d.Set("rx_power_alarm_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-low"]; exists {
        d.Set("rx_power_alarm_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-high"]; exists {
        d.Set("rx_power_warning_high", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-low"]; exists {
        d.Set("rx_power_warning_low", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-high-ch1"]; exists {
        d.Set("rx_power_alarm_high_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-low-ch1"]; exists {
        d.Set("rx_power_alarm_low_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-high-ch1"]; exists {
        d.Set("rx_power_warning_high_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-low-ch1"]; exists {
        d.Set("rx_power_warning_low_ch1", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-high-ch2"]; exists {
        d.Set("rx_power_alarm_high_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-low-ch2"]; exists {
        d.Set("rx_power_alarm_low_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-high-ch2"]; exists {
        d.Set("rx_power_warning_high_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-low-ch2"]; exists {
        d.Set("rx_power_warning_low_ch2", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-high-ch3"]; exists {
        d.Set("rx_power_alarm_high_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-low-ch3"]; exists {
        d.Set("rx_power_alarm_low_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-high-ch3"]; exists {
        d.Set("rx_power_warning_high_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-low-ch3"]; exists {
        d.Set("rx_power_warning_low_ch3", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-high-ch4"]; exists {
        d.Set("rx_power_alarm_high_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-alarm-low-ch4"]; exists {
        d.Set("rx_power_alarm_low_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-high-ch4"]; exists {
        d.Set("rx_power_warning_high_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["rx-power-warning-low-ch4"]; exists {
        d.Set("rx_power_warning_low_ch4", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-physical-interface-xcvr-detail-" + acctest.RandString(10)))
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    return nil
}

