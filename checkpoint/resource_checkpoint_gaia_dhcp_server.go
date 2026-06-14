package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaDhcpServer() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaDhcpServer,
        Read:   readGaiaDhcpServer,
        Update: updateGaiaDhcpServer,
        Delete: deleteGaiaDhcpServer,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `DHCP server status`,
            },
            "subnets": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Subnets`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Enable DHCP on this subnet.`,
                        },
                        "max_lease": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `The longest lease that the server can allocate, in seconds.`,
                        },
                        "default_lease": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `The default lease that the server allocates, in seconds.`,
                        },
                        "subnet": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Subnet name.`,
                        },
                        "netmask": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Subnet mask.`,
                        },
                        "default_gateway": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The IPv4 address of the default gateway for the DHCP clients. If not exist, empty string will be returned.`,
                        },
                        "dns": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `DNS configuration.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "primary": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv4 address of the Primary DNS server for the DHCP clients. If not exist, empty string will be returned.`,
                                    },
                                    "secondary": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv4 address of the Secondary DNS server for the DHCP clients (to use if the primary DNS server does not respond). If not exist, empty string will be returned.`,
                                    },
                                    "tertiary": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv4 address of the Tertiary DNS server for the DHCP clients (to use if the primary and secondary DNS servers do not respond). If not exist, empty string will be returned.`,
                                    },
                                    "domain_name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv4 address of the Tertiary DNS server for the DHCP clients (to use if the primary and secondary DNS servers do not respond). If not exist, empty string will be returned.`,
                                    },
                                },
                            },
                        },
                        "ip_pools": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Range of IPv4 addresses that the server assigns to hosts.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Description: `Enables or disables the DHCP Server for this subnet IP pool.`,
                                    },
                                    "include": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies whether to include or exclude this range of IPv4 addresses in the IP pool.`,
                                    },
                                    "start": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The first IPv4 address of the range.`,
                                    },
                                    "end": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The last IPv4 address of the range.`,
                                    },
                                },
                            },
                        },
                    },
                },
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

func createGaiaDhcpServer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("subnets"); len(v.([]interface{})) > 0 {
        subnetsList := v.([]interface{})
        subnetsArray := make([]interface{}, 0, len(subnetsList))
        for i := range subnetsList {
            itemMap := make(map[string]interface{})
            if v := d.Get(fmt.Sprintf("subnets.%d.enabled", i)).(bool); v {
                itemMap["enabled"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.max_lease", i)); ok {
                itemMap["max-lease"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.default_lease", i)); ok {
                itemMap["default-lease"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.netmask", i)); ok {
                itemMap["netmask"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.default_gateway", i)); ok {
                itemMap["default-gateway"] = v.(string)
            }
            if sv, ok := d.GetOk(fmt.Sprintf("subnets.%d.dns", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    dnsMap := make(map[string]interface{})
                    if sv, ok := rawDict["primary"]; ok && sv.(string) != "" {
                        dnsMap["primary"] = sv.(string)
                    }
                    if sv, ok := rawDict["secondary"]; ok && sv.(string) != "" {
                        dnsMap["secondary"] = sv.(string)
                    }
                    if sv, ok := rawDict["tertiary"]; ok && sv.(string) != "" {
                        dnsMap["tertiary"] = sv.(string)
                    }
                    if sv, ok := rawDict["domain_name"]; ok && sv.(string) != "" {
                        dnsMap["domain-name"] = sv.(string)
                    }
                    if len(dnsMap) > 0 {
                        itemMap["dns"] = dnsMap
                    }
                }
            }
            if sv := d.Get(fmt.Sprintf("subnets.%d.ip_pools", i)); len(sv.([]interface{})) > 0 {
                ip_poolsList := sv.([]interface{})
                ip_poolsArr := make([]interface{}, 0, len(ip_poolsList))
                for j := range ip_poolsList {
                    innerMap := make(map[string]interface{})
                    if v := d.Get(fmt.Sprintf("subnets.%d.ip_pools.%d.enabled", i, j)).(bool); v {
                        innerMap["enabled"] = v
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.include", i, j)); ok {
                        innerMap["include"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.start", i, j)); ok {
                        innerMap["start"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.end", i, j)); ok {
                        innerMap["end"] = iv.(string)
                    }
                    if len(innerMap) > 0 {
                        ip_poolsArr = append(ip_poolsArr, innerMap)
                    }
                }
                if len(ip_poolsArr) > 0 {
                    itemMap["ip-pools"] = ip_poolsArr
                }
            }
            if len(itemMap) > 0 {
                subnetsArray = append(subnetsArray, itemMap)
            }
        }
        if len(subnetsArray) > 0 {
            payload["subnets"] = subnetsArray
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create DhcpServer - Map = ", payload)

    addDhcpServerRes, err := client.ApiCallSimple("set-dhcp-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addDhcpServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addDhcpServerRes.Success {
            errMsg = addDhcpServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addDhcpServerRes.GetData()
        }

        debugLogOperation(
            "dhcp-server",        // resource type
            "create",                       // operation
            "set-dhcp-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add dhcp-server: %v", err)
    }
    if !addDhcpServerRes.Success {
        if addDhcpServerRes.ErrorMsg != "" {
            return fmt.Errorf(addDhcpServerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("dhcp-server-" + acctest.RandString(10)))
    return readGaiaDhcpServer(d, m)
}

func readGaiaDhcpServer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showDhcpServerRes, err := client.ApiCallSimple("show-dhcp-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showDhcpServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showDhcpServerRes.Success {
            errMsg = showDhcpServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showDhcpServerRes.GetData()
        }

        debugLogOperation(
            "dhcp-server",        // resource type
            "read",                       // operation
            "show-dhcp-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show dhcp-server: %v", err)
    }
    if !showDhcpServerRes.Success {
        if data := showDhcpServerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showDhcpServerRes.ErrorMsg)
    }

    dhcpServer := showDhcpServerRes.GetData()

    log.Println("Read DhcpServer - Show JSON = ", dhcpServer)

    if v, exists := dhcpServer["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := dhcpServer["subnets"]; exists {
        if rawSubnets, ok := v.([]interface{}); ok {
            subnets := make([]interface{}, 0, len(rawSubnets))
            for _, s := range rawSubnets {
                raw, ok := s.(map[string]interface{})
                if !ok {
                    continue
                }
                subnet := map[string]interface{}{
                    "subnet":          raw["subnet"],
                    "netmask":         raw["netmask"],
                    "enabled":         raw["enabled"],
                    "max_lease":       raw["max-lease"],
                    "default_lease":   raw["default-lease"],
                    "default_gateway": raw["default-gateway"],
                }
                if dns, ok := raw["dns"].(map[string]interface{}); ok {
                    subnet["dns"] = []interface{}{map[string]interface{}{
                        "primary":     fmt.Sprintf("%v", dns["primary"]),
                        "secondary":   fmt.Sprintf("%v", dns["secondary"]),
                        "tertiary":    fmt.Sprintf("%v", dns["tertiary"]),
                        "domain_name": fmt.Sprintf("%v", dns["domain-name"]),
                    }}
                }
                if pools, ok := raw["ip-pools"].([]interface{}); ok {
                    mapped := make([]interface{}, 0, len(pools))
                    for _, p := range pools {
                        if pm, ok := p.(map[string]interface{}); ok {
                            mapped = append(mapped, map[string]interface{}{
                                "start":   fmt.Sprintf("%v", pm["start"]),
                                "end":     fmt.Sprintf("%v", pm["end"]),
                                "enabled": pm["enabled"],
                                "include": fmt.Sprintf("%v", pm["include"]),
                            })
                        }
                    }
                    subnet["ip_pools"] = mapped
                }
                subnets = append(subnets, subnet)
            }
            d.Set("subnets", subnets)
        }
    }
    if v, exists := dhcpServer["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaDhcpServer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("subnets"); len(v.([]interface{})) > 0 {
        subnetsList := v.([]interface{})
        subnetsArray := make([]interface{}, 0, len(subnetsList))
        for i := range subnetsList {
            itemMap := make(map[string]interface{})
            if v := d.Get(fmt.Sprintf("subnets.%d.enabled", i)).(bool); v {
                itemMap["enabled"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.max_lease", i)); ok {
                itemMap["max-lease"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.default_lease", i)); ok {
                itemMap["default-lease"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.netmask", i)); ok {
                itemMap["netmask"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.default_gateway", i)); ok {
                itemMap["default-gateway"] = v.(string)
            }
            if sv, ok := d.GetOk(fmt.Sprintf("subnets.%d.dns", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    dnsMap := make(map[string]interface{})
                    if sv, ok := rawDict["primary"]; ok && sv.(string) != "" {
                        dnsMap["primary"] = sv.(string)
                    }
                    if sv, ok := rawDict["secondary"]; ok && sv.(string) != "" {
                        dnsMap["secondary"] = sv.(string)
                    }
                    if sv, ok := rawDict["tertiary"]; ok && sv.(string) != "" {
                        dnsMap["tertiary"] = sv.(string)
                    }
                    if sv, ok := rawDict["domain_name"]; ok && sv.(string) != "" {
                        dnsMap["domain-name"] = sv.(string)
                    }
                    if len(dnsMap) > 0 {
                        itemMap["dns"] = dnsMap
                    }
                }
            }
            if sv := d.Get(fmt.Sprintf("subnets.%d.ip_pools", i)); len(sv.([]interface{})) > 0 {
                ip_poolsList := sv.([]interface{})
                ip_poolsArr := make([]interface{}, 0, len(ip_poolsList))
                for j := range ip_poolsList {
                    innerMap := make(map[string]interface{})
                    if v := d.Get(fmt.Sprintf("subnets.%d.ip_pools.%d.enabled", i, j)).(bool); v {
                        innerMap["enabled"] = v
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.include", i, j)); ok {
                        innerMap["include"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.start", i, j)); ok {
                        innerMap["start"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("subnets.%d.ip_pools.%d.end", i, j)); ok {
                        innerMap["end"] = iv.(string)
                    }
                    if len(innerMap) > 0 {
                        ip_poolsArr = append(ip_poolsArr, innerMap)
                    }
                }
                if len(ip_poolsArr) > 0 {
                    itemMap["ip-pools"] = ip_poolsArr
                }
            }
            if len(itemMap) > 0 {
                subnetsArray = append(subnetsArray, itemMap)
            }
        }
        if len(subnetsArray) > 0 {
            payload["subnets"] = subnetsArray
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setDhcpServerRes, err := client.ApiCallSimple("set-dhcp-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setDhcpServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setDhcpServerRes.Success {
            errMsg = setDhcpServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setDhcpServerRes.GetData()
        }

        debugLogOperation(
            "dhcp-server",        // resource type
            "update",                       // operation
            "set-dhcp-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set dhcp-server: %v", err)
    }
    if !setDhcpServerRes.Success {
        return fmt.Errorf(setDhcpServerRes.ErrorMsg)
    }

    return readGaiaDhcpServer(d, m)
}

func deleteGaiaDhcpServer(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    