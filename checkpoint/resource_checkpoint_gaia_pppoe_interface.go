package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaPppoeInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPppoeInterface,
        Read:   readGaiaPppoeInterface,
        Update: updateGaiaPppoeInterface,
        Delete: deleteGaiaPppoeInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "username": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The username needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP`,
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
            "client_id": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `The PPPoE client Id. This ID must be unique for every PPPoE interface.`,
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the applicable physical interface. Gaia uses this interface to forward PPPoE frames.`,
            },
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `The password needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP`,
            },
            "password_hash": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Sensitive:   true,
                Description: `The hash of the password needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP.`,
            },
            "use_peer_as_default_gateway": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable to make the ISP server the Default Gateway for the Gaia.`,
            },
            "use_peer_dns": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable to allow the ISP to define the IPv4 DNS server for the Gaia.`,
            },
            "fake_peer_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Fake peer settings`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The fake unicast peer IPv4 address (the default value is 0.0.0.0).`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Enable to use the configured fake peer IPv4 address.`,
                        },
                    },
                },
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable to turn on the interface.`,
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `User comments.`,
            },
            "name": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `The PPPoE interface name.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaPppoeInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("username"); ok {
        payload["username"] = v.(string)
    }

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

    if v, ok := d.GetOk("client_id"); ok {
        payload["client-id"] = v.(int)
    }

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    if v, ok := d.GetOkExists("use_peer_as_default_gateway"); ok {
        payload["use-peer-as-default-gateway"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_peer_dns"); ok {
        payload["use-peer-dns"] = v.(bool)
    }

    if v := d.Get("fake_peer_settings"); len(v.([]interface{})) > 0 {
        _ = v
        fakepeersettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("fake_peer_settings.0.address"); ok {
            fakepeersettingsMap["address"] = v.(string)
        }
        if v, ok := d.GetOkExists("fake_peer_settings.0.enabled"); ok && v.(bool) {
            fakepeersettingsMap["enabled"] = v.(bool)
        }
        if len(fakepeersettingsMap) > 0 {
            payload["fake-peer-settings"] = fakepeersettingsMap
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    log.Println("Create PppoeInterface - Map = ", payload)

    addPppoeInterfaceRes, err := client.ApiCallSimple("add-pppoe-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPppoeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPppoeInterfaceRes.Success {
            errMsg = addPppoeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPppoeInterfaceRes.GetData()
        }

        debugLogOperation(
            "pppoe-interface",        // resource type
            "create",                       // operation
            "add-pppoe-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add pppoe-interface: %v", err)
    }
    if !addPppoeInterfaceRes.Success {
        if addPppoeInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addPppoeInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addPppoeInterfaceRes.GetData(); data != nil {
        if v, exists := data["name"]; exists {
            d.Set("name", v)
        }
    }

    // Two-phase creation: Apply update-only fields if present
    hasUpdateOnlyFields := false
    updatePayload := map[string]interface{}{
    "username": payload["username"],
    "interface": payload["interface"],
    }

    if v, ok := d.GetOk("name"); ok {
        updatePayload["name"] = v.(string)
        hasUpdateOnlyFields = true
    }

    if hasUpdateOnlyFields {
        log.Println("Two-phase creation: Applying update-only fields - Map = ", updatePayload)
        
        setPppoeInterfaceRes, err := client.ApiCallSimple("set-pppoe-interface", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for pppoe-interface: %v", err)
        }
        if !setPppoeInterfaceRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setPppoeInterfaceRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("pppoe-interface-" + acctest.RandString(10)))
    return readGaiaPppoeInterface(d, m)
}

func readGaiaPppoeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

   payload["name"] = d.Get("name")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPppoeInterfaceRes, err := client.ApiCallSimple("show-pppoe-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPppoeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPppoeInterfaceRes.Success {
            errMsg = showPppoeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPppoeInterfaceRes.GetData()
        }

        debugLogOperation(
            "pppoe-interface",        // resource type
            "read",                       // operation
            "show-pppoe-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show pppoe-interface: %v", err)
    }
    if !showPppoeInterfaceRes.Success {
        if data := showPppoeInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPppoeInterfaceRes.ErrorMsg)
    }

    pppoeInterface := showPppoeInterfaceRes.GetData()

    log.Println("Read PppoeInterface - Show JSON = ", pppoeInterface)

    if v, exists := pppoeInterface["sd-wan"]; exists {
        if sm, ok := v.(map[string]interface{}); ok {
            entry := map[string]interface{}{
                "enabled":       sm["enabled"],
                "next_hop":      fmt.Sprintf("%v", sm["next-hop"]),
                "next_hop_ipv6": fmt.Sprintf("%v", sm["next-hop-ipv6"]),
                "tag":           fmt.Sprintf("%v", sm["tag"]),
                "circuit_id":    fmt.Sprintf("%v", sm["circuit-id"]),
            }
            if nat, ok := sm["nat"].(map[string]interface{}); ok {
                entry["nat"] = []interface{}{map[string]interface{}{
                    "enabled": nat["enabled"],
                    "ip":      fmt.Sprintf("%v", nat["ip"]),
                    "ipv6":    fmt.Sprintf("%v", nat["ipv6"]),
                }}
            }
            if bw, ok := sm["bandwidth"].(map[string]interface{}); ok {
                entry["bandwidth"] = []interface{}{map[string]interface{}{
                    "upload_speed":   fmt.Sprintf("%v", bw["upload-speed"]),
                    "download_speed": fmt.Sprintf("%v", bw["download-speed"]),
                }}
            }
            d.Set("sd_wan", []interface{}{entry})
        }
    }
    if v, exists := pppoeInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := pppoeInterface["client-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("client_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("client_id", _n)
            }
        }
    }
    if v, exists := pppoeInterface["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := pppoeInterface["username"]; exists {
        d.Set("username", fmt.Sprintf("%v", v))
    }
    if v, exists := pppoeInterface["password-hash"]; exists {
        d.Set("password_hash", fmt.Sprintf("%v", v))
    }
    if v, exists := pppoeInterface["use-peer-as-default-gateway"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("use_peer_as_default_gateway", b)
        } else if s, ok := v.(string); ok {
            d.Set("use_peer_as_default_gateway", s == "true")
        }
    }
    if v, exists := pppoeInterface["use-peer-dns"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("use_peer_dns", b)
        } else if s, ok := v.(string); ok {
            d.Set("use_peer_dns", s == "true")
        }
    }
    if v, exists := pppoeInterface["fake-peer-settings"]; exists {
        if fm, ok := v.(map[string]interface{}); ok {
            d.Set("fake_peer_settings", []interface{}{map[string]interface{}{
                "address": fmt.Sprintf("%v", fm["address"]),
                "enabled": fm["enabled"],
            }})
        }
    }
    if v, exists := pppoeInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := pppoeInterface["status"]; exists {
        d.Set("status", fmt.Sprintf("%v", v))
    }
    if v, exists := pppoeInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPppoeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("username"); ok {
        payload["username"] = v.(string)
    }

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

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOkExists("use_peer_as_default_gateway"); ok {
        payload["use-peer-as-default-gateway"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_peer_dns"); ok {
        payload["use-peer-dns"] = v.(bool)
    }

    if v := d.Get("fake_peer_settings"); len(v.([]interface{})) > 0 {
        _ = v
        fakepeersettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("fake_peer_settings.0.address"); ok {
            fakepeersettingsMap["address"] = v.(string)
        }
        if v, ok := d.GetOkExists("fake_peer_settings.0.enabled"); ok && v.(bool) {
            fakepeersettingsMap["enabled"] = v.(bool)
        }
        if len(fakepeersettingsMap) > 0 {
            payload["fake-peer-settings"] = fakepeersettingsMap
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

   payload["name"] = d.Get("name")
    setPppoeInterfaceRes, err := client.ApiCallSimple("set-pppoe-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPppoeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPppoeInterfaceRes.Success {
            errMsg = setPppoeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPppoeInterfaceRes.GetData()
        }

        debugLogOperation(
            "pppoe-interface",        // resource type
            "update",                       // operation
            "set-pppoe-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set pppoe-interface: %v", err)
    }
    if !setPppoeInterfaceRes.Success {
        return fmt.Errorf(setPppoeInterfaceRes.ErrorMsg)
    }

    return readGaiaPppoeInterface(d, m)
}

func deleteGaiaPppoeInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

   payload["name"] = d.Get("name")
    deletePppoeInterfaceRes, err := client.ApiCallSimple("delete-pppoe-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deletePppoeInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deletePppoeInterfaceRes.Success {
            errMsg = deletePppoeInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deletePppoeInterfaceRes.GetData()
        }

        debugLogOperation(
            "pppoe-interface",        // resource type
            "delete",                       // operation
            "delete-pppoe-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete pppoe-interface: %v", err)
    }
    if !deletePppoeInterfaceRes.Success {
        return fmt.Errorf(deletePppoeInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

