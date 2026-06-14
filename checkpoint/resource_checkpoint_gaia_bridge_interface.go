package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBridgeInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBridgeInterface,
        Read:   readGaiaBridgeInterface,
        Update: updateGaiaBridgeInterface,
        Delete: deleteGaiaBridgeInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "dhcp6": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `DHCPv6 configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: `Enable DHCP on this interface.`,
                        },
                        "server_timeout": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the amount of time, in seconds, that must pass between the time that the interface begins to try to determine its address and the time that it decides that it's not going to be able to contact a server.`,
                        },
                        "retry": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the time, in seconds, that must pass after the interface has determined that there is no DHCP server present before it tries again to contact a DHCP server.`,
                        },
                        "leasetime": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the lease time, in seconds, when requesting for an IP address. Default value is \"default\" - according to the server.`,
                        },
                        "reacquire_timeout": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `When trying to reacquire the last IP address, the reacquire-timeout statement sets the time, in seconds, that must elapse after the first try to reacquire the old address before it gives up and tries to discover a new address.`,
                        },
                        "using": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Choose the DHCPv6 client working mode of this interface.          Interface will receive IPv6 only if the chosen mode and the system's configured mode match`,
                        },
                    },
                },
            },
            "dhcp": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `DHCP configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: `Enable DHCP on this interface.`,
                        },
                        "server_timeout": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the amount of time, in seconds, that must pass between the time that the interface begins to try to determine its address and the time that it decides that it's not going to be able to contact a server.`,
                        },
                        "retry": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the time, in seconds, that must pass after the interface has determined that there is no DHCP server present before it tries again to contact a DHCP server.`,
                        },
                        "leasetime": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the lease time, in seconds, when requesting for an IP address. Default value is \"default\" - according to the server.`,
                        },
                        "reacquire_timeout": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `When trying to reacquire the last IP address, the reacquire-timeout statement sets the time, in seconds, that must elapse after the first try to reacquire the old address before it gives up and tries to discover a new address.`,
                        },
                    },
                },
            },
            "mtu": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `N/A`,
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv4_mask_length": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `N/A`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_autoconfig": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
            },
            "ipv6_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_mask_length": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `N/A`,
            },
            "members": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "resource_id": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "link_state": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "speed": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "duplex": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_bytes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_packets": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_bytes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_packets": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_local_link_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaBridgeInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("dhcp6"); len(v.([]interface{})) > 0 {
        _ = v
        dhcp6Map := make(map[string]interface{})
        if v, ok := d.GetOkExists("dhcp6.0.enabled"); ok && v.(bool) {
            dhcp6Map["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("dhcp6.0.server_timeout"); ok {
            dhcp6Map["server-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.retry"); ok {
            dhcp6Map["retry"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.leasetime"); ok {
            dhcp6Map["leasetime"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.reacquire_timeout"); ok {
            dhcp6Map["reacquire-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.using"); ok {
            dhcp6Map["using"] = v.(string)
        }
        if len(dhcp6Map) > 0 {
            payload["dhcp6"] = dhcp6Map
        }
    }

    if v := d.Get("dhcp"); len(v.([]interface{})) > 0 {
        _ = v
        dhcpMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("dhcp.0.enabled"); ok && v.(bool) {
            dhcpMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("dhcp.0.server_timeout"); ok {
            dhcpMap["server-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.retry"); ok {
            dhcpMap["retry"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.leasetime"); ok {
            dhcpMap["leasetime"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.reacquire_timeout"); ok {
            dhcpMap["reacquire-timeout"] = v.(string)
        }
        if len(dhcpMap) > 0 {
            payload["dhcp"] = dhcpMap
        }
    }

    if v, ok := d.GetOk("mtu"); ok {
        payload["mtu"] = v.(int)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        payload["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_mask_length"); ok {
        payload["ipv4-mask-length"] = v.(int)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ipv6_autoconfig"); ok {
        payload["ipv6-autoconfig"] = v.(bool)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address"); ok {
        payload["ipv6-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_mask_length"); ok {
        payload["ipv6-mask-length"] = v.(int)
    }

    if v := d.Get("members"); len(v.(*schema.Set).List()) > 0 {
        payload["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create BridgeInterface - Map = ", payload)

    addBridgeInterfaceRes, err := client.ApiCallSimple("add-bridge-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBridgeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBridgeInterfaceRes.Success {
            errMsg = addBridgeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBridgeInterfaceRes.GetData()
        }

        debugLogOperation(
            "bridge-interface",        // resource type
            "create",                       // operation
            "add-bridge-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add bridge-interface: %v", err)
    }
    if !addBridgeInterfaceRes.Success {
        if addBridgeInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addBridgeInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addBridgeInterfaceRes.GetData(); data != nil {
        if v, exists := data["name"]; exists {
            d.Set("name", v)
        }
    }

    // Two-phase creation: Apply update-only fields if present
    hasUpdateOnlyFields := false
    updatePayload := map[string]interface{}{
    }

    if v, ok := d.GetOk("name"); ok {
        updatePayload["name"] = v.(string)
        hasUpdateOnlyFields = true
    }

    if hasUpdateOnlyFields {
        log.Println("Two-phase creation: Applying update-only fields - Map = ", updatePayload)
        
        setBridgeInterfaceRes, err := client.ApiCallSimple("set-bridge-interface", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for bridge-interface: %v", err)
        }
        if !setBridgeInterfaceRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setBridgeInterfaceRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("bridge-interface-" + acctest.RandString(10)))
    return readGaiaBridgeInterface(d, m)
}

func readGaiaBridgeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showBridgeInterfaceRes, err := client.ApiCallSimple("show-bridge-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBridgeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBridgeInterfaceRes.Success {
            errMsg = showBridgeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBridgeInterfaceRes.GetData()
        }

        debugLogOperation(
            "bridge-interface",        // resource type
            "read",                       // operation
            "show-bridge-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show bridge-interface: %v", err)
    }
    if !showBridgeInterfaceRes.Success {
        if data := showBridgeInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBridgeInterfaceRes.ErrorMsg)
    }

    bridgeInterface := showBridgeInterfaceRes.GetData()

    log.Println("Read BridgeInterface - Show JSON = ", bridgeInterface)

    if v, exists := bridgeInterface["dhcp6"]; exists {
        d.Set("dhcp6", v)
    }
    if v, exists := bridgeInterface["dhcp"]; exists {
        d.Set("dhcp", v)
    }
    if v, exists := bridgeInterface["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := bridgeInterface["speed"]; exists {
        d.Set("speed", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["duplex"]; exists {
        d.Set("duplex", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["tx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_bytes", _n)
            }
        }
    }
    if v, exists := bridgeInterface["tx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_packets", _n)
            }
        }
    }
    if v, exists := bridgeInterface["rx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_bytes", _n)
            }
        }
    }
    if v, exists := bridgeInterface["rx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_packets", _n)
            }
        }
    }
    if v, exists := bridgeInterface["members"]; exists {
        d.Set("members", v.([]interface{}))
    }
    if v, exists := bridgeInterface["mtu"]; exists {
        d.Set("mtu", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["ipv4-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv4_address", _val)
    }
    if v, exists := bridgeInterface["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := bridgeInterface["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := bridgeInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["ipv6-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv6_address", _val)
    }
    if v, exists := bridgeInterface["ipv6-mask-length"]; exists {
        d.Set("ipv6_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := bridgeInterface["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBridgeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("dhcp6"); len(v.([]interface{})) > 0 {
        _ = v
        dhcp6Map := make(map[string]interface{})
        if v, ok := d.GetOkExists("dhcp6.0.enabled"); ok && v.(bool) {
            dhcp6Map["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("dhcp6.0.server_timeout"); ok {
            dhcp6Map["server-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.retry"); ok {
            dhcp6Map["retry"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.leasetime"); ok {
            dhcp6Map["leasetime"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.reacquire_timeout"); ok {
            dhcp6Map["reacquire-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp6.0.using"); ok {
            dhcp6Map["using"] = v.(string)
        }
        if len(dhcp6Map) > 0 {
            payload["dhcp6"] = dhcp6Map
        }
    }

    if v := d.Get("dhcp"); len(v.([]interface{})) > 0 {
        _ = v
        dhcpMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("dhcp.0.enabled"); ok && v.(bool) {
            dhcpMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("dhcp.0.server_timeout"); ok {
            dhcpMap["server-timeout"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.retry"); ok {
            dhcpMap["retry"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.leasetime"); ok {
            dhcpMap["leasetime"] = v.(string)
        }
        if v, ok := d.GetOk("dhcp.0.reacquire_timeout"); ok {
            dhcpMap["reacquire-timeout"] = v.(string)
        }
        if len(dhcpMap) > 0 {
            payload["dhcp"] = dhcpMap
        }
    }

    if v, ok := d.GetOk("mtu"); ok {
        payload["mtu"] = v.(int)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        payload["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_mask_length"); ok {
        payload["ipv4-mask-length"] = v.(int)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ipv6_autoconfig"); ok {
        payload["ipv6-autoconfig"] = v.(bool)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address"); ok {
        payload["ipv6-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_mask_length"); ok {
        payload["ipv6-mask-length"] = v.(int)
    }

    if v := d.Get("members"); len(v.(*schema.Set).List()) > 0 {
        payload["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    setBridgeInterfaceRes, err := client.ApiCallSimple("set-bridge-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBridgeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBridgeInterfaceRes.Success {
            errMsg = setBridgeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBridgeInterfaceRes.GetData()
        }

        debugLogOperation(
            "bridge-interface",        // resource type
            "update",                       // operation
            "set-bridge-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set bridge-interface: %v", err)
    }
    if !setBridgeInterfaceRes.Success {
        return fmt.Errorf(setBridgeInterfaceRes.ErrorMsg)
    }

    return readGaiaBridgeInterface(d, m)
}

func deleteGaiaBridgeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    deleteBridgeInterfaceRes, err := client.ApiCallSimple("delete-bridge-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteBridgeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteBridgeInterfaceRes.Success {
            errMsg = deleteBridgeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteBridgeInterfaceRes.GetData()
        }

        debugLogOperation(
            "bridge-interface",        // resource type
            "delete",                       // operation
            "delete-bridge-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete bridge-interface: %v", err)
    }
    if !deleteBridgeInterfaceRes.Success {
        return fmt.Errorf(deleteBridgeInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

