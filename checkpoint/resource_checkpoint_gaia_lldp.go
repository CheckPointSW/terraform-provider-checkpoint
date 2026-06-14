package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaLldp() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLldp,
        Read:   readGaiaLldp,
        Update: updateGaiaLldp,
        Delete: deleteGaiaLldp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `LLDP State`,
            },
            "timers": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `LLDP Timers`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hold_time_multiplier": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Define LLDP hold time multiplier interval to cache learned information before discarding. Range: 2-10, Default: 4 (giving 120 seconds LLDP cache lifetime with other defaults).`,
                        },
                        "transmit_interval": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Define LLDP packet transmitting interval (Seconds). Range: 8-32768 seconds, Default: 30 seconds.`,
                        },
                    },
                },
            },
            "tlv": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `LLDP Tlv`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "management_address": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Define Gaia to send the Management Address information in the LLDP packets.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Description: `Define Gaia to send the Management Address information in the LLDP packets.`,
                                    },
                                    "ip_from": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `configured-interface - Send Configured interface IP within the LLDP packets, mgmt-interface - Send Management interface IP within the LLDP packets, Default is configured-interface. (supported from version R81.20)`,
                                    },
                                },
                            },
                        },
                        "port_description": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Define Gaia to send the Port Description information in the LLDP packets.`,
                        },
                        "system_capabilities": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Define Gaia to send the System Capabilities information in the LLDP packets.`,
                        },
                        "system_description": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Define Gaia to send the System Description information in the LLDP packets.`,
                        },
                        "system_name": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Define Gaia to send the System Name information in the LLDP packets.`,
                        },
                    },
                },
            },
            "interfaces": {
                Type:     schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interface_name": {
                            Type:     schema.TypeString,
                            Required: true,
                        },
                        "mode": {
                            Type:     schema.TypeString,
                            Optional: true,
                            Description: `transmit-and-receive, transmit, receive`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaLldp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("timers"); len(v.([]interface{})) > 0 {
        _ = v
        timersMap := make(map[string]interface{})
        if v, ok := d.GetOk("timers.0.hold_time_multiplier"); ok {
            timersMap["hold-time-multiplier"] = v.(int)
        }
        if v, ok := d.GetOk("timers.0.transmit_interval"); ok {
            timersMap["transmit-interval"] = v.(int)
        }
        if len(timersMap) > 0 {
            payload["timers"] = timersMap
        }
    }

    if v := d.Get("tlv"); len(v.([]interface{})) > 0 {
        _ = v
        tlvMap := make(map[string]interface{})
        if v, ok := d.GetOk("tlv.0.management_address"); ok {
            _ = v
            managementaddressMap := make(map[string]interface{})
            if v, ok := d.GetOkExists("tlv.0.management_address.0.enabled"); ok && v.(bool) {
                managementaddressMap["enabled"] = v.(bool)
            }
            if v, ok := d.GetOk("tlv.0.management_address.0.ip_from"); ok {
                managementaddressMap["ip-from"] = v.(string)
            }
            if len(managementaddressMap) > 0 {
                tlvMap["management-address"] = managementaddressMap
            }
        }
        if v, ok := d.GetOkExists("tlv.0.port_description"); ok && v.(bool) {
            tlvMap["port-description"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_capabilities"); ok && v.(bool) {
            tlvMap["system-capabilities"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_description"); ok && v.(bool) {
            tlvMap["system-description"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_name"); ok && v.(bool) {
            tlvMap["system-name"] = v.(bool)
        }
        if len(tlvMap) > 0 {
            payload["tlv"] = tlvMap
        }
    }

    if v := d.Get("interfaces"); len(v.([]interface{})) > 0 {
        ifaceList := v.([]interface{})
        ifaceArray := make([]interface{}, 0, len(ifaceList))
        for _, item := range ifaceList {
            m := item.(map[string]interface{})
            entry := map[string]interface{}{}
            if n, ok := m["interface_name"].(string); ok && n != "" {
                entry["interface-name"] = n
            }
            if mo, ok := m["mode"].(string); ok && mo != "" {
                entry["mode"] = mo
            }
            ifaceArray = append(ifaceArray, entry)
        }
        payload["interfaces"] = ifaceArray
    }

    log.Println("Create Lldp - Map = ", payload)

    addLldpRes, err := client.ApiCallSimple("set-lldp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addLldpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addLldpRes.Success {
            errMsg = addLldpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addLldpRes.GetData()
        }

        debugLogOperation(
            "lldp",        // resource type
            "create",                       // operation
            "set-lldp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add lldp: %v", err)
    }
    if !addLldpRes.Success {
        if addLldpRes.ErrorMsg != "" {
            return fmt.Errorf(addLldpRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("lldp-" + acctest.RandString(10)))
    return readGaiaLldp(d, m)
}

func readGaiaLldp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showLldpRes, err := client.ApiCallSimple("show-lldp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showLldpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showLldpRes.Success {
            errMsg = showLldpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showLldpRes.GetData()
        }

        debugLogOperation(
            "lldp",        // resource type
            "read",                       // operation
            "show-lldp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show lldp: %v", err)
    }
    if !showLldpRes.Success {
        if data := showLldpRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showLldpRes.ErrorMsg)
    }

    lldp := showLldpRes.GetData()

    log.Println("Read Lldp - Show JSON = ", lldp)

    if v, exists := lldp["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := lldp["interfaces"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "interface_name": fmt.Sprintf("%v", m["interface-name"]),
                        "mode":           fmt.Sprintf("%v", m["mode"]),
                    })
                }
            }
            d.Set("interfaces", out)
        }
    }
    if v, exists := lldp["timers"]; exists {
        d.Set("timers", v)
    }
    if v, exists := lldp["tlv"]; exists {
        d.Set("tlv", v)
    }
    if v, exists := lldp["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaLldp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("timers"); len(v.([]interface{})) > 0 {
        _ = v
        timersMap := make(map[string]interface{})
        if v, ok := d.GetOk("timers.0.hold_time_multiplier"); ok {
            timersMap["hold-time-multiplier"] = v.(int)
        }
        if v, ok := d.GetOk("timers.0.transmit_interval"); ok {
            timersMap["transmit-interval"] = v.(int)
        }
        if len(timersMap) > 0 {
            payload["timers"] = timersMap
        }
    }

    if v := d.Get("tlv"); len(v.([]interface{})) > 0 {
        _ = v
        tlvMap := make(map[string]interface{})
        if v, ok := d.GetOk("tlv.0.management_address"); ok {
            _ = v
            managementaddressMap := make(map[string]interface{})
            if v, ok := d.GetOkExists("tlv.0.management_address.0.enabled"); ok && v.(bool) {
                managementaddressMap["enabled"] = v.(bool)
            }
            if v, ok := d.GetOk("tlv.0.management_address.0.ip_from"); ok {
                managementaddressMap["ip-from"] = v.(string)
            }
            if len(managementaddressMap) > 0 {
                tlvMap["management-address"] = managementaddressMap
            }
        }
        if v, ok := d.GetOkExists("tlv.0.port_description"); ok && v.(bool) {
            tlvMap["port-description"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_capabilities"); ok && v.(bool) {
            tlvMap["system-capabilities"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_description"); ok && v.(bool) {
            tlvMap["system-description"] = v.(bool)
        }
        if v, ok := d.GetOkExists("tlv.0.system_name"); ok && v.(bool) {
            tlvMap["system-name"] = v.(bool)
        }
        if len(tlvMap) > 0 {
            payload["tlv"] = tlvMap
        }
    }

    if v := d.Get("interfaces"); len(v.([]interface{})) > 0 {
        ifaceList := v.([]interface{})
        ifaceArray := make([]interface{}, 0, len(ifaceList))
        for _, item := range ifaceList {
            m := item.(map[string]interface{})
            entry := map[string]interface{}{}
            if n, ok := m["interface_name"].(string); ok && n != "" {
                entry["interface-name"] = n
            }
            if mo, ok := m["mode"].(string); ok && mo != "" {
                entry["mode"] = mo
            }
            ifaceArray = append(ifaceArray, entry)
        }
        payload["interfaces"] = ifaceArray
    }

    setLldpRes, err := client.ApiCallSimple("set-lldp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setLldpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setLldpRes.Success {
            errMsg = setLldpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setLldpRes.GetData()
        }

        debugLogOperation(
            "lldp",        // resource type
            "update",                       // operation
            "set-lldp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set lldp: %v", err)
    }
    if !setLldpRes.Success {
        return fmt.Errorf(setLldpRes.ErrorMsg)
    }

    return readGaiaLldp(d, m)
}

func deleteGaiaLldp(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    