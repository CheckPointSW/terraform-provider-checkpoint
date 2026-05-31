package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBondInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBondInterface,
        Read:   readGaiaBondInterface,
        Update: updateGaiaBondInterface,
        Delete: deleteGaiaBondInterface,
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
            "ip_conflicts": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable ip-conflicts on this interface to monitor the Address Resolution Protocol traffic on the connected network.`,
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
            "xmit_hash_policy": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `layer2, layer3+4`,
            },
            "down_delay": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Unit: Milliseconds`,
            },
            "up_delay": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Unit: Milliseconds`,
            },
            "primary": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Primary member of the bond interface`,
            },
            "lacp_rate": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "mode": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
            },
            "mii_interval": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Media monitoring interval`,
            },
            "min_links": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Valid only while mode 8023AD is configured`,
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

func createGaiaBondInterface(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOkExists("ip_conflicts"); ok {
        payload["ip-conflicts"] = v.(bool)
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

    if v := d.Get("members"); len(v.(*schema.Set).List()) > 0 {
        payload["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("xmit_hash_policy"); ok {
        payload["xmit-hash-policy"] = v.(string)
    }

    if v, ok := d.GetOk("down_delay"); ok {
        payload["down-delay"] = v.(int)
    }

    if v, ok := d.GetOk("up_delay"); ok {
        payload["up-delay"] = v.(int)
    }

    if v, ok := d.GetOk("primary"); ok {
        payload["primary"] = v.(string)
    }

    if v, ok := d.GetOk("lacp_rate"); ok {
        payload["lacp-rate"] = v.(string)
    }

    if v, ok := d.GetOk("mode"); ok {
        payload["mode"] = v.(string)
    }

    if v, ok := d.GetOk("mii_interval"); ok {
        payload["mii-interval"] = v.(int)
    }

    if v, ok := d.GetOk("min_links"); ok {
        payload["min-links"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create BondInterface - Map = ", payload)

    addBondInterfaceRes, err := client.ApiCallSimple("add-bond-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBondInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBondInterfaceRes.Success {
            errMsg = addBondInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBondInterfaceRes.GetData()
        }

        debugLogOperation(
            "bond-interface",        // resource type
            "create",                       // operation
            "add-bond-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add bond-interface: %v", err)
    }
    if !addBondInterfaceRes.Success {
        if addBondInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addBondInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addBondInterfaceRes.GetData(); data != nil {
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
        
        setBondInterfaceRes, err := client.ApiCallSimple("set-bond-interface", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for bond-interface: %v", err)
        }
        if !setBondInterfaceRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setBondInterfaceRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("bond-interface-" + acctest.RandString(10)))
    return readGaiaBondInterface(d, m)
}

func readGaiaBondInterface(d *schema.ResourceData, m interface{}) error {

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

    showBondInterfaceRes, err := client.ApiCallSimple("show-bond-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBondInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBondInterfaceRes.Success {
            errMsg = showBondInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBondInterfaceRes.GetData()
        }

        debugLogOperation(
            "bond-interface",        // resource type
            "read",                       // operation
            "show-bond-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show bond-interface: %v", err)
    }
    if !showBondInterfaceRes.Success {
        if data := showBondInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBondInterfaceRes.ErrorMsg)
    }

    bondInterface := showBondInterfaceRes.GetData()

    log.Println("Read BondInterface - Show JSON = ", bondInterface)

    if v, exists := bondInterface["sd-wan"]; exists {
        d.Set("sd_wan", v)
    }
    if v, exists := bondInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["ip-conflicts"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ip_conflicts", b)
        } else if s, ok := v.(string); ok {
            d.Set("ip_conflicts", s == "true")
        }
    }
    if v, exists := bondInterface["dhcp6"]; exists {
        d.Set("dhcp6", v)
    }
    if v, exists := bondInterface["dhcp"]; exists {
        d.Set("dhcp", v)
    }
    if v, exists := bondInterface["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := bondInterface["speed"]; exists {
        d.Set("speed", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["duplex"]; exists {
        d.Set("duplex", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["tx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_bytes", _n)
            }
        }
    }
    if v, exists := bondInterface["tx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_packets", _n)
            }
        }
    }
    if v, exists := bondInterface["rx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_bytes", _n)
            }
        }
    }
    if v, exists := bondInterface["rx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_packets", _n)
            }
        }
    }
    if v, exists := bondInterface["members"]; exists {
        d.Set("members", v.([]interface{}))
    }
    if v, exists := bondInterface["mtu"]; exists {
        d.Set("mtu", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["ipv4-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv4_address", _val)
    }
    if v, exists := bondInterface["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := bondInterface["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := bondInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["ipv6-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv6_address", _val)
    }
    if v, exists := bondInterface["ipv6-mask-length"]; exists {
        d.Set("ipv6_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["xmit-hash-policy"]; exists {
        d.Set("xmit_hash_policy", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["down-delay"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("down_delay", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("down_delay", _n)
            }
        }
    }
    if v, exists := bondInterface["up-delay"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("up_delay", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("up_delay", _n)
            }
        }
    }
    if v, exists := bondInterface["primary"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("primary", _val)
    }
    if v, exists := bondInterface["lacp-rate"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("lacp_rate", _val)
    }
    if v, exists := bondInterface["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := bondInterface["mii-interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("mii_interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("mii_interval", _n)
            }
        }
    }
    if v, exists := bondInterface["min-links"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("min_links", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("min_links", _n)
            }
        }
    }
    if v, exists := bondInterface["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBondInterface(d *schema.ResourceData, m interface{}) error {

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

    if v, ok := d.GetOkExists("ip_conflicts"); ok {
        payload["ip-conflicts"] = v.(bool)
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

    if v := d.Get("members"); len(v.(*schema.Set).List()) > 0 {
        payload["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("xmit_hash_policy"); ok {
        payload["xmit-hash-policy"] = v.(string)
    }

    if v, ok := d.GetOk("down_delay"); ok {
        payload["down-delay"] = v.(int)
    }

    if v, ok := d.GetOk("up_delay"); ok {
        payload["up-delay"] = v.(int)
    }

    if v, ok := d.GetOk("primary"); ok {
        payload["primary"] = v.(string)
    }

    if v, ok := d.GetOk("lacp_rate"); ok {
        payload["lacp-rate"] = v.(string)
    }

    if v, ok := d.GetOk("mode"); ok {
        payload["mode"] = v.(string)
    }

    if v, ok := d.GetOk("mii_interval"); ok {
        payload["mii-interval"] = v.(int)
    }

    if v, ok := d.GetOk("min_links"); ok {
        payload["min-links"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    setBondInterfaceRes, err := client.ApiCallSimple("set-bond-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBondInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBondInterfaceRes.Success {
            errMsg = setBondInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBondInterfaceRes.GetData()
        }

        debugLogOperation(
            "bond-interface",        // resource type
            "update",                       // operation
            "set-bond-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set bond-interface: %v", err)
    }
    if !setBondInterfaceRes.Success {
        return fmt.Errorf(setBondInterfaceRes.ErrorMsg)
    }

    return readGaiaBondInterface(d, m)
}

func deleteGaiaBondInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    deleteBondInterfaceRes, err := client.ApiCallSimple("delete-bond-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteBondInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteBondInterfaceRes.Success {
            errMsg = deleteBondInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteBondInterfaceRes.GetData()
        }

        debugLogOperation(
            "bond-interface",        // resource type
            "delete",                       // operation
            "delete-bond-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete bond-interface: %v", err)
    }
    if !deleteBondInterfaceRes.Success {
        return fmt.Errorf(deleteBondInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

