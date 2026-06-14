package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMaestroPort() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMaestroPort,
        Read:   readGaiaMaestroPort,
        Update: updateGaiaMaestroPort,
        Delete: deleteGaiaMaestroPort,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "resource_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Port ID (e.g. '1/13/1')`,
            },
            "interface_name": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Interface name in case this port is an Uplink or MGMT interface (e.g. 'eth1-25')`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Setting this to false will disable this port, setting to true will enable it. AKA 'admin state'`,
            },
            "mtu": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `MTU of this port`,
            },
            "auto_negotiation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If true, Auto Negotiation will be turned on, and vice versa`,
            },
            "qsfp_mode": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Port QSFP mode. Valid values are: '4x10G', '4x25G', '25G', '40G', '100G'`,
            },
            "type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Port type. Valid values are: 'downlink', 'uplink', 'site_sync', 'ssm_sync', 'mgmt'`,
            },
            "site": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "link_state": {
                Type:     schema.TypeBool,
                Computed: true,
            },
            "transceiver_state": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "operating_speed": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "rx_frames": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "tx_frames": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "orchestrator_id": {
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeInt},
            },
        },
    }
}

func createGaiaMaestroPort(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("interface_name"); ok {
        payload["interface-name"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("mtu"); ok {
        payload["mtu"] = v.(int)
    }

    if v, ok := d.GetOkExists("auto_negotiation"); ok {
        payload["auto-negotiation"] = v.(bool)
    }

    if v, ok := d.GetOk("qsfp_mode"); ok {
        payload["qsfp-mode"] = v.(string)
    }

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
    }

    log.Println("Create MaestroPort - Map = ", payload)

    addMaestroPortRes, err := client.ApiCallSimple("set-maestro-port", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMaestroPortRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMaestroPortRes.Success {
            errMsg = addMaestroPortRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMaestroPortRes.GetData()
        }

        debugLogOperation(
            "maestro-port",        // resource type
            "create",                       // operation
            "set-maestro-port",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add maestro-port: %v", err)
    }
    if !addMaestroPortRes.Success {
        if addMaestroPortRes.ErrorMsg != "" {
            return fmt.Errorf(addMaestroPortRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("maestro-port-" + acctest.RandString(10)))
    return readGaiaMaestroPort(d, m)
}

func readGaiaMaestroPort(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("interface_name"); ok {
        payload["interface-name"] = v.(string)
    }

    showMaestroPortRes, err := client.ApiCallSimple("show-maestro-port", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMaestroPortRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMaestroPortRes.Success {
            errMsg = showMaestroPortRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMaestroPortRes.GetData()
        }

        debugLogOperation(
            "maestro-port",        // resource type
            "read",                       // operation
            "show-maestro-port",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show maestro-port: %v", err)
    }
    if !showMaestroPortRes.Success {
        if data := showMaestroPortRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMaestroPortRes.ErrorMsg)
    }

    maestroPort := showMaestroPortRes.GetData()

    log.Println("Read MaestroPort - Show JSON = ", maestroPort)

    if v, exists := maestroPort["id"]; exists {
        d.Set("resource_id", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["site"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("site", int(f))
        }
    }
    if v, exists := maestroPort["interface-name"]; exists {
        d.Set("interface_name", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["type"]; exists {
        d.Set("type", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["qsfp-mode"]; exists {
        d.Set("qsfp_mode", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := maestroPort["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := maestroPort["auto-negotiation"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("auto_negotiation", b)
        } else if s, ok := v.(string); ok {
            d.Set("auto_negotiation", s == "true")
        }
    }
    if v, exists := maestroPort["transceiver-state"]; exists {
        d.Set("transceiver_state", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["operating-speed"]; exists {
        d.Set("operating_speed", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroPort["mtu"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("mtu", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("mtu", _n)
            }
        }
    }
    if v, exists := maestroPort["rx-frames"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_frames", int(f))
        }
    }
    if v, exists := maestroPort["tx-frames"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_frames", int(f))
        }
    }
    if v, exists := maestroPort["orchestrator-id"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if f, ok := item.(float64); ok {
                    out = append(out, int(f))
                } else {
                    out = append(out, item)
                }
            }
            d.Set("orchestrator_id", out)
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMaestroPort(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("interface_name"); ok {
        payload["interface-name"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("mtu"); ok {
        payload["mtu"] = v.(int)
    }

    if v, ok := d.GetOkExists("auto_negotiation"); ok {
        payload["auto-negotiation"] = v.(bool)
    }

    if v, ok := d.GetOk("qsfp_mode"); ok {
        payload["qsfp-mode"] = v.(string)
    }

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
    }

    setMaestroPortRes, err := client.ApiCallSimple("set-maestro-port", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMaestroPortRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMaestroPortRes.Success {
            errMsg = setMaestroPortRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMaestroPortRes.GetData()
        }

        debugLogOperation(
            "maestro-port",        // resource type
            "update",                       // operation
            "set-maestro-port",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set maestro-port: %v", err)
    }
    if !setMaestroPortRes.Success {
        return fmt.Errorf(setMaestroPortRes.ErrorMsg)
    }

    return readGaiaMaestroPort(d, m)
}

func deleteGaiaMaestroPort(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    