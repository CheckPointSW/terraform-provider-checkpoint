package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaArp() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaArp,
        Read:   readGaiaArp,
        Update: updateGaiaArp,
        Delete: deleteGaiaArp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure ARP settings`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restriction_level": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Define different restriction levels for announcing the local source IP address          from IP packets in ARP requests sent on interface:          0 - Any local address          1 - Use address from the same subnet as the target address          2 - Prefer primary address`,
                        },
                        "cache_size": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Specify the maximum number of entries in the arp cache`,
                        },
                        "validity_timeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Specify time, in seconds, to keep resolved dynamic ARP entries.         If the entry is not referred to and is not used by traffic before the time elapses, it is deleted.        Otherwise, a request will be sent to verify the MAC address.`,
                        },
                        "auto_cache_size": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Update cache size to be automatically changed depending on the current ARP table entries in the system.  ARP table default size: 4096, Range: 1024-131072, Supported starting from R82.00. Supported starting from Gaia version R82`,
                        },
                    },
                },
            },
            "proxy": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Add a specified Proxy ARP entries`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Define the IP address of a new proxy ARP entry`,
                        },
                        "interface": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Define the interface used when forwarding packets to the given IP address`,
                        },
                        "mac_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Define the hardware address used when forwarding packets to the given IP address`,
                        },
                        "real_ipv4_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Define the real IP address used when forwarding packets to the given IP address`,
                        },
                    },
                },
            },
            "static": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Add a specified Static ARP entries`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Define the IP address of a new static ARP entry`,
                        },
                        "mac_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specify the hardware address used when forwarding packets to the given IP address`,
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

func createGaiaArp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("settings"); len(v.([]interface{})) > 0 {
        _ = v
        settingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("settings.0.restriction_level"); ok {
            settingsMap["restriction-level"] = v.(int)
        }
        if v, ok := d.GetOk("settings.0.cache_size"); ok {
            settingsMap["cache-size"] = v.(int)
        }
        if v, ok := d.GetOk("settings.0.validity_timeout"); ok {
            settingsMap["validity-timeout"] = v.(int)
        }
        if v, ok := d.GetOkExists("settings.0.auto_cache_size"); ok && v.(bool) {
            settingsMap["auto-cache-size"] = v.(bool)
        }
        if len(settingsMap) > 0 {
            payload["settings"] = settingsMap
        }
    }

    if v := d.Get("proxy"); len(v.([]interface{})) > 0 {
        proxyList := v.([]interface{})
        proxyArray := make([]interface{}, 0, len(proxyList))
        for i := range proxyList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.ipv4_address", i)); ok {
                itemMap["ipv4-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.interface", i)); ok {
                itemMap["interface"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.mac_address", i)); ok {
                itemMap["mac-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.real_ipv4_address", i)); ok {
                itemMap["real-ipv4-address"] = v.(string)
            }
            if len(itemMap) > 0 {
                proxyArray = append(proxyArray, itemMap)
            }
        }
        if len(proxyArray) > 0 {
            payload["proxy"] = proxyArray
        }
    }

    if v := d.Get("static"); len(v.([]interface{})) > 0 {
        staticList := v.([]interface{})
        staticArray := make([]interface{}, 0, len(staticList))
        for i := range staticList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("static.%d.ipv4_address", i)); ok {
                itemMap["ipv4-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static.%d.mac_address", i)); ok {
                itemMap["mac-address"] = v.(string)
            }
            if len(itemMap) > 0 {
                staticArray = append(staticArray, itemMap)
            }
        }
        if len(staticArray) > 0 {
            payload["static"] = staticArray
        }
    }

    log.Println("Create Arp - Map = ", payload)

    addArpRes, err := client.ApiCallSimple("set-arp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addArpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addArpRes.Success {
            errMsg = addArpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addArpRes.GetData()
        }

        debugLogOperation(
            "arp",        // resource type
            "create",                       // operation
            "set-arp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add arp: %v", err)
    }
    if !addArpRes.Success {
        if addArpRes.ErrorMsg != "" {
            return fmt.Errorf(addArpRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("arp-" + acctest.RandString(10)))
    return readGaiaArp(d, m)
}

func readGaiaArp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showArpRes, err := client.ApiCallSimple("show-arp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showArpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showArpRes.Success {
            errMsg = showArpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showArpRes.GetData()
        }

        debugLogOperation(
            "arp",        // resource type
            "read",                       // operation
            "show-arp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show arp: %v", err)
    }
    if !showArpRes.Success {
        if data := showArpRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showArpRes.ErrorMsg)
    }

    arp := showArpRes.GetData()

    log.Println("Read Arp - Show JSON = ", arp)

    if v, exists := arp["settings"]; exists {
        if sm, ok := v.(map[string]interface{}); ok {
            if existing := d.Get("settings"); len(existing.([]interface{})) > 0 {
                settings := map[string]interface{}{}
                if val, ok := sm["restriction-level"]; ok {
                    if f, ok := val.(float64); ok { settings["restriction_level"] = int(f) }
                }
                if val, ok := sm["cache-size"]; ok {
                    if f, ok := val.(float64); ok { settings["cache_size"] = int(f) }
                }
                if val, ok := sm["validity-timeout"]; ok {
                    if f, ok := val.(float64); ok { settings["validity_timeout"] = int(f) }
                }
                if val, ok := sm["auto-cache-size"]; ok {
                    if b, ok := val.(bool); ok { settings["auto_cache_size"] = b }
                }
                d.Set("settings", []interface{}{settings})
            }
        }
    }
    if v, exists := arp["proxy"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, ri := range items {
                if im, ok := ri.(map[string]interface{}); ok {
                    item := map[string]interface{}{}
                    if val, ok := im["ipv4-address"]; ok {
                        item["ipv4_address"] = fmt.Sprintf("%v", val)
                    }
                    if val, ok := im["interface"]; ok && fmt.Sprintf("%v", val) != "" {
                        item["interface"] = fmt.Sprintf("%v", val)
                    }
                    if val, ok := im["mac-address"]; ok && fmt.Sprintf("%v", val) != "" {
                        item["mac_address"] = strings.ToLower(fmt.Sprintf("%v", val))
                    }
                    if val, ok := im["real-ipv4-address"]; ok && fmt.Sprintf("%v", val) != "" {
                        item["real_ipv4_address"] = fmt.Sprintf("%v", val)
                    }
                    if len(item) > 0 {
                        out = append(out, item)
                    }
                }
            }
            d.Set("proxy", out)
        }
    }
    if v, exists := arp["static"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, ri := range items {
                if im, ok := ri.(map[string]interface{}); ok {
                    item := map[string]interface{}{}
                    if val, ok := im["ipv4-address"]; ok {
                        item["ipv4_address"] = fmt.Sprintf("%v", val)
                    }
                    if val, ok := im["mac-address"]; ok {
                        item["mac_address"] = strings.ToLower(fmt.Sprintf("%v", val))
                    }
                    if len(item) > 0 {
                        out = append(out, item)
                    }
                }
            }
            d.Set("static", out)
        }
    }
    if v, exists := arp["dynamic"]; exists {
        d.Set("dynamic", v.([]interface{}))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaArp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("settings"); len(v.([]interface{})) > 0 {
        _ = v
        settingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("settings.0.restriction_level"); ok {
            settingsMap["restriction-level"] = v.(int)
        }
        if v, ok := d.GetOk("settings.0.cache_size"); ok {
            settingsMap["cache-size"] = v.(int)
        }
        if v, ok := d.GetOk("settings.0.validity_timeout"); ok {
            settingsMap["validity-timeout"] = v.(int)
        }
        if v, ok := d.GetOkExists("settings.0.auto_cache_size"); ok && v.(bool) {
            settingsMap["auto-cache-size"] = v.(bool)
        }
        if len(settingsMap) > 0 {
            payload["settings"] = settingsMap
        }
    }

    if v := d.Get("proxy"); len(v.([]interface{})) > 0 {
        proxyList := v.([]interface{})
        proxyArray := make([]interface{}, 0, len(proxyList))
        for i := range proxyList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.ipv4_address", i)); ok {
                itemMap["ipv4-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.interface", i)); ok {
                itemMap["interface"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.mac_address", i)); ok {
                itemMap["mac-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("proxy.%d.real_ipv4_address", i)); ok {
                itemMap["real-ipv4-address"] = v.(string)
            }
            if len(itemMap) > 0 {
                proxyArray = append(proxyArray, itemMap)
            }
        }
        if len(proxyArray) > 0 {
            payload["proxy"] = proxyArray
        }
    }

    if v := d.Get("static"); len(v.([]interface{})) > 0 {
        staticList := v.([]interface{})
        staticArray := make([]interface{}, 0, len(staticList))
        for i := range staticList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("static.%d.ipv4_address", i)); ok {
                itemMap["ipv4-address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static.%d.mac_address", i)); ok {
                itemMap["mac-address"] = v.(string)
            }
            if len(itemMap) > 0 {
                staticArray = append(staticArray, itemMap)
            }
        }
        if len(staticArray) > 0 {
            payload["static"] = staticArray
        }
    }

    setArpRes, err := client.ApiCallSimple("set-arp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setArpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setArpRes.Success {
            errMsg = setArpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setArpRes.GetData()
        }

        debugLogOperation(
            "arp",        // resource type
            "update",                       // operation
            "set-arp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set arp: %v", err)
    }
    if !setArpRes.Success {
        return fmt.Errorf(setArpRes.ErrorMsg)
    }

    return readGaiaArp(d, m)
}

func deleteGaiaArp(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    