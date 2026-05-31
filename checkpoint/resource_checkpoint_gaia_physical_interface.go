package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaPhysicalInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPhysicalInterface,
        Read:   readGaiaPhysicalInterface,
        Update: updateGaiaPhysicalInterface,
        Delete: deleteGaiaPhysicalInterface,
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
                Computed:    true,
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
                Computed:    true,
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
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "auto_negotiation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Activating Auto-Negotiation will skip the speed and duplex configuration`,
            },
            "speed": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Speed is not relevant when 'auto-negotiation' is enabled`,
            },
            "duplex": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Duplex is not relevant when 'auto-negotiation' is enabled`,
            },
            "monitor_mode": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "mac_addr": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_ringsize": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_ringsize": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaPhysicalInterface(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("auto_negotiation"); ok {
        payload["auto-negotiation"] = v.(bool)
    }

    if v, ok := d.GetOk("speed"); ok {
        payload["speed"] = v.(string)
    }

    if v, ok := d.GetOk("duplex"); ok {
        payload["duplex"] = v.(string)
    }

    if v, ok := d.GetOkExists("monitor_mode"); ok {
        payload["monitor-mode"] = v.(bool)
    }

    if v, ok := d.GetOk("mac_addr"); ok {
        payload["mac-addr"] = v.(string)
    }

    if v, ok := d.GetOk("rx_ringsize"); ok {
        payload["rx-ringsize"] = v.(int)
    }

    if v, ok := d.GetOk("tx_ringsize"); ok {
        payload["tx-ringsize"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create PhysicalInterface - Map = ", payload)

    addPhysicalInterfaceRes, err := client.ApiCallSimple("set-physical-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPhysicalInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPhysicalInterfaceRes.Success {
            errMsg = addPhysicalInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPhysicalInterfaceRes.GetData()
        }

        debugLogOperation(
            "physical-interface",        // resource type
            "create",                       // operation
            "set-physical-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add physical-interface: %v", err)
    }
    if !addPhysicalInterfaceRes.Success {
        if addPhysicalInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addPhysicalInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("physical-interface-" + acctest.RandString(10)))
    return readGaiaPhysicalInterface(d, m)
}

func readGaiaPhysicalInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPhysicalInterfaceRes, err := client.ApiCallSimple("show-physical-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPhysicalInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPhysicalInterfaceRes.Success {
            errMsg = showPhysicalInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPhysicalInterfaceRes.GetData()
        }

        debugLogOperation(
            "physical-interface",        // resource type
            "read",                       // operation
            "show-physical-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show physical-interface: %v", err)
    }
    if !showPhysicalInterfaceRes.Success {
        if data := showPhysicalInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPhysicalInterfaceRes.ErrorMsg)
    }

    physicalInterface := showPhysicalInterfaceRes.GetData()

    log.Println("Read PhysicalInterface - Show JSON = ", physicalInterface)

    if v, exists := physicalInterface["sd-wan"]; exists {
        d.Set("sd_wan", v)
    }
    if v, exists := physicalInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["ip-conflicts"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ip_conflicts", b)
        } else if s, ok := v.(string); ok {
            d.Set("ip_conflicts", s == "true")
        }
    }
    if v, exists := physicalInterface["dhcp6"]; exists {
        d.Set("dhcp6", v)
    }
    if v, exists := physicalInterface["dhcp"]; exists {
        d.Set("dhcp", v)
    }
    if v, exists := physicalInterface["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := physicalInterface["speed"]; exists {
        d.Set("speed", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["duplex"]; exists {
        d.Set("duplex", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["tx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_bytes", _n)
            }
        }
    }
    if v, exists := physicalInterface["tx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_packets", _n)
            }
        }
    }
    if v, exists := physicalInterface["rx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_bytes", _n)
            }
        }
    }
    if v, exists := physicalInterface["rx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_packets", _n)
            }
        }
    }
    if v, exists := physicalInterface["mtu"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("mtu", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("mtu", _n)
            }
        }
    }
    if v, exists := physicalInterface["ipv4-address"]; exists {
        d.Set("ipv4_address", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["ipv4-mask-length"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("ipv4_mask_length", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("ipv4_mask_length", _n)
            }
        }
    }
    if v, exists := physicalInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := physicalInterface["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := physicalInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["ipv6-address"]; exists {
        _val := fmt.Sprintf("%v", v)
        if strings.HasPrefix(_val, "Not") {
            _val = ""
        }
        d.Set("ipv6_address", _val)
    }
    if v, exists := physicalInterface["ipv6-mask-length"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("ipv6_mask_length", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("ipv6_mask_length", _n)
            }
        }
    }
    if v, exists := physicalInterface["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["status"]; exists {
        d.Set("status", v)
    }
    if v, exists := physicalInterface["auto-negotiation"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("auto_negotiation", b)
        } else if s, ok := v.(string); ok {
            d.Set("auto_negotiation", s == "true")
        }
    }
    if v, exists := physicalInterface["monitor-mode"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("monitor_mode", b)
        } else if s, ok := v.(string); ok {
            d.Set("monitor_mode", s == "true")
        }
    }
    if v, exists := physicalInterface["mac-addr"]; exists {
        d.Set("mac_addr", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["rx-ringsize"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_ringsize", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_ringsize", _n)
            }
        }
    }
    if v, exists := physicalInterface["tx-ringsize"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_ringsize", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_ringsize", _n)
            }
        }
    }
    if v, exists := physicalInterface["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := physicalInterface["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPhysicalInterface(d *schema.ResourceData, m interface{}) error {

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

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("auto_negotiation"); ok {
        payload["auto-negotiation"] = v.(bool)
    }

    if v, ok := d.GetOk("speed"); ok {
        payload["speed"] = v.(string)
    }

    if v, ok := d.GetOk("duplex"); ok {
        payload["duplex"] = v.(string)
    }

    if v, ok := d.GetOkExists("monitor_mode"); ok {
        payload["monitor-mode"] = v.(bool)
    }

    if v, ok := d.GetOk("mac_addr"); ok {
        payload["mac-addr"] = v.(string)
    }

    if v, ok := d.GetOk("rx_ringsize"); ok {
        payload["rx-ringsize"] = v.(int)
    }

    if v, ok := d.GetOk("tx_ringsize"); ok {
        payload["tx-ringsize"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setPhysicalInterfaceRes, err := client.ApiCallSimple("set-physical-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPhysicalInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPhysicalInterfaceRes.Success {
            errMsg = setPhysicalInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPhysicalInterfaceRes.GetData()
        }

        debugLogOperation(
            "physical-interface",        // resource type
            "update",                       // operation
            "set-physical-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set physical-interface: %v", err)
    }
    if !setPhysicalInterfaceRes.Success {
        return fmt.Errorf(setPhysicalInterfaceRes.ErrorMsg)
    }

    return readGaiaPhysicalInterface(d, m)
}

func deleteGaiaPhysicalInterface(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    