package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaVlanInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaVlanInterface,
        Read:   readGaiaVlanInterface,
        Update: updateGaiaVlanInterface,
        Delete: deleteGaiaVlanInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "sd_wan": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `SD-WAN configuration.  Supported starting from R81.20 JHF 14`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: `Enable SD-WAN on this interface.`,
                        },
                        "next_hop": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure interface's next hop IPv4 address, obtain next hop IPv4 address automatically         or set as a layer 2-only link`,
                        },
                        "next_hop_ipv6": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure interface's next hop IPv6 address or obtain next hop IPv6 address automatically.              IPv6 configuration is supported starting from R82 latest Jumbo Hotfix`,
                        },
                        "nat": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Optional NAT configuration`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: `Enable NAT IP address on this interface`,
                                    },
                                    "ip": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Configure NAT IPv4 address on this interface or obtain NAT IPv4 address automatically.`,
                                    },
                                    "ipv6": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Configure NAT IPv6 address on this interface or obtain NAT IPv6 address automatically.              IPv6 configuration is supported starting from R82 latest Jumbo Hotfix`,
                                    },
                                },
                            },
                        },
                        "tag": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Optional tag configuration.             Must contain only alphanumeric characters, '-' or '_' (max length is 64)`,
                        },
                        "bandwidth": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Optional Bandwidth configuration.              Bandwidth configuration is supported starting from R81.20 JHF 79`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "upload_speed": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `In Mbps`,
                                    },
                                    "download_speed": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `In Mbps`,
                                    },
                                },
                            },
                        },
                        "circuit_id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Optional override interface circuit id value.              Circuit-ID configuration is supported starting from R81.20 JHF 79`,
                        },
                    },
                },
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
                Description: `N/A`,
            },
            "ipv6_mask_length": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `N/A`,
            },
            "parent": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `VLAN Trunk`,
            },
            "resource_id": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `VLAN Tag`,
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

func createGaiaVlanInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("sd_wan"); len(v.([]interface{})) > 0 {
        _ = v
        sdwanMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("sd_wan.0.enabled"); ok && v.(bool) {
            sdwanMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("sd_wan.0.next_hop"); ok {
            sdwanMap["next-hop"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.next_hop_ipv6"); ok {
            sdwanMap["next-hop-ipv6"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.nat"); ok {
            _ = v
            natMap := make(map[string]interface{})
            if v, ok := d.GetOkExists("sd_wan.0.nat.0.enabled"); ok && v.(bool) {
                natMap["enabled"] = v.(bool)
            }
            if v, ok := d.GetOk("sd_wan.0.nat.0.ip"); ok {
                natMap["ip"] = v.(string)
            }
            if v, ok := d.GetOk("sd_wan.0.nat.0.ipv6"); ok {
                natMap["ipv6"] = v.(string)
            }
            if len(natMap) > 0 {
                sdwanMap["nat"] = natMap
            }
        }
        if v, ok := d.GetOk("sd_wan.0.tag"); ok {
            sdwanMap["tag"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.bandwidth"); ok {
            _ = v
            bandwidthMap := make(map[string]interface{})
            if v, ok := d.GetOk("sd_wan.0.bandwidth.0.upload_speed"); ok {
                bandwidthMap["upload-speed"] = v.(string)
            }
            if v, ok := d.GetOk("sd_wan.0.bandwidth.0.download_speed"); ok {
                bandwidthMap["download-speed"] = v.(string)
            }
            if len(bandwidthMap) > 0 {
                sdwanMap["bandwidth"] = bandwidthMap
            }
        }
        if v, ok := d.GetOk("sd_wan.0.circuit_id"); ok {
            sdwanMap["circuit-id"] = v.(string)
        }
        if len(sdwanMap) > 0 {
            payload["sd-wan"] = sdwanMap
        }
    }

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

    if v, ok := d.GetOk("parent"); ok {
        payload["parent"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create VlanInterface - Map = ", payload)

    addVlanInterfaceRes, err := client.ApiCallSimple("add-vlan-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addVlanInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addVlanInterfaceRes.Success {
            errMsg = addVlanInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addVlanInterfaceRes.GetData()
        }

        debugLogOperation(
            "vlan-interface",        // resource type
            "create",                       // operation
            "add-vlan-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add vlan-interface: %v", err)
    }
    if !addVlanInterfaceRes.Success {
        if addVlanInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addVlanInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
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
        
        setVlanInterfaceRes, err := client.ApiCallSimple("set-vlan-interface", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for vlan-interface: %v", err)
        }
        if !setVlanInterfaceRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setVlanInterfaceRes.ErrorMsg)
        }
    }

    // Derive and persist the interface name from parent, resource_id so Read/Delete can use it.
    _nameField0, _ok0 := d.GetOk("parent")
    _nameField1, _ok1 := d.GetOk("resource_id")
    if _ok0 && _ok1 {
        d.Set("name", fmt.Sprintf("%s.%d", _nameField0.(string), _nameField1.(int)))
    }

    d.SetId(fmt.Sprintf("vlan-interface-" + acctest.RandString(10)))
    return readGaiaVlanInterface(d, m)
}

func readGaiaVlanInterface(d *schema.ResourceData, m interface{}) error {

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

    showVlanInterfaceRes, err := client.ApiCallSimple("show-vlan-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showVlanInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showVlanInterfaceRes.Success {
            errMsg = showVlanInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showVlanInterfaceRes.GetData()
        }

        debugLogOperation(
            "vlan-interface",        // resource type
            "read",                       // operation
            "show-vlan-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show vlan-interface: %v", err)
    }
    if !showVlanInterfaceRes.Success {
        if data := showVlanInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showVlanInterfaceRes.ErrorMsg)
    }

    vlanInterface := showVlanInterfaceRes.GetData()

    log.Println("Read VlanInterface - Show JSON = ", vlanInterface)

    if v, exists := vlanInterface["sd-wan"]; exists {
        d.Set("sd_wan", v)
    }
    if v, exists := vlanInterface["dhcp6"]; exists {
        d.Set("dhcp6", v)
    }
    if v, exists := vlanInterface["dhcp"]; exists {
        d.Set("dhcp", v)
    }
    if v, exists := vlanInterface["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := vlanInterface["speed"]; exists {
        d.Set("speed", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["duplex"]; exists {
        d.Set("duplex", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["tx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_bytes", _n)
            }
        }
    }
    if v, exists := vlanInterface["tx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_packets", _n)
            }
        }
    }
    if v, exists := vlanInterface["rx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_bytes", _n)
            }
        }
    }
    if v, exists := vlanInterface["rx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_packets", _n)
            }
        }
    }
    if v, exists := vlanInterface["parent"]; exists {
        d.Set("parent", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["mtu"]; exists {
        d.Set("mtu", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["ipv4-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv4_address", _val)
    }
    if v, exists := vlanInterface["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := vlanInterface["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := vlanInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["ipv6-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv6_address", _val)
    }
    if v, exists := vlanInterface["ipv6-mask-length"]; exists {
        d.Set("ipv6_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := vlanInterface["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaVlanInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("sd_wan"); len(v.([]interface{})) > 0 {
        _ = v
        sdwanMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("sd_wan.0.enabled"); ok && v.(bool) {
            sdwanMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("sd_wan.0.next_hop"); ok {
            sdwanMap["next-hop"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.next_hop_ipv6"); ok {
            sdwanMap["next-hop-ipv6"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.nat"); ok {
            _ = v
            natMap := make(map[string]interface{})
            if v, ok := d.GetOkExists("sd_wan.0.nat.0.enabled"); ok && v.(bool) {
                natMap["enabled"] = v.(bool)
            }
            if v, ok := d.GetOk("sd_wan.0.nat.0.ip"); ok {
                natMap["ip"] = v.(string)
            }
            if v, ok := d.GetOk("sd_wan.0.nat.0.ipv6"); ok {
                natMap["ipv6"] = v.(string)
            }
            if len(natMap) > 0 {
                sdwanMap["nat"] = natMap
            }
        }
        if v, ok := d.GetOk("sd_wan.0.tag"); ok {
            sdwanMap["tag"] = v.(string)
        }
        if v, ok := d.GetOk("sd_wan.0.bandwidth"); ok {
            _ = v
            bandwidthMap := make(map[string]interface{})
            if v, ok := d.GetOk("sd_wan.0.bandwidth.0.upload_speed"); ok {
                bandwidthMap["upload-speed"] = v.(string)
            }
            if v, ok := d.GetOk("sd_wan.0.bandwidth.0.download_speed"); ok {
                bandwidthMap["download-speed"] = v.(string)
            }
            if len(bandwidthMap) > 0 {
                sdwanMap["bandwidth"] = bandwidthMap
            }
        }
        if v, ok := d.GetOk("sd_wan.0.circuit_id"); ok {
            sdwanMap["circuit-id"] = v.(string)
        }
        if len(sdwanMap) > 0 {
            payload["sd-wan"] = sdwanMap
        }
    }

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

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    setVlanInterfaceRes, err := client.ApiCallSimple("set-vlan-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setVlanInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setVlanInterfaceRes.Success {
            errMsg = setVlanInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setVlanInterfaceRes.GetData()
        }

        debugLogOperation(
            "vlan-interface",        // resource type
            "update",                       // operation
            "set-vlan-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set vlan-interface: %v", err)
    }
    if !setVlanInterfaceRes.Success {
        return fmt.Errorf(setVlanInterfaceRes.ErrorMsg)
    }

    return readGaiaVlanInterface(d, m)
}

func deleteGaiaVlanInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    deleteVlanInterfaceRes, err := client.ApiCallSimple("delete-vlan-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteVlanInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteVlanInterfaceRes.Success {
            errMsg = deleteVlanInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteVlanInterfaceRes.GetData()
        }

        debugLogOperation(
            "vlan-interface",        // resource type
            "delete",                       // operation
            "delete-vlan-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete vlan-interface: %v", err)
    }
    if !deleteVlanInterfaceRes.Success {
        return fmt.Errorf(deleteVlanInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

