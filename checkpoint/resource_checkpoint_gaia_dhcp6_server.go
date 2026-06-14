package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaDhcp6Server() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaDhcp6Server,
        Read:   readGaiaDhcp6Server,
        Update: updateGaiaDhcp6Server,
        Delete: deleteGaiaDhcp6Server,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `DHCPv6 server status`,
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
                        "prefix": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Subnet prefix length.`,
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
                                        Description: `The IPv6 address of the Primary DNS server for the DHCP clients. If not exist, empty string will be returned.`,
                                    },
                                    "secondary": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv6 address of the Secondary DNS server for the DHCP clients (to use if the primary DNS server does not respond). If not exist, empty string will be returned.`,
                                    },
                                    "tertiary": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IPv6 address of the Tertiary DNS server for the DHCP clients (to use if the primary and secondary DNS servers do not respond). If not exist, empty string will be returned.`,
                                    },
                                    "domain_name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The domain name of the IPv6 subnet. If not exist, empty string will be returned.`,
                                    },
                                },
                            },
                        },
                        "ip_pools": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Range of IPv6 addresses that the server assigns to hosts.`,
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
                                        Description: `The first IPv6 address of the range.`,
                                    },
                                    "end": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The last IPv6 address of the range.`,
                                    },
                                },
                            },
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

func createGaiaDhcp6Server(d *schema.ResourceData, m interface{}) error {
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
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.prefix", i)); ok {
                itemMap["prefix"] = v.(int)
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

    log.Println("Create Dhcp6Server - Map = ", payload)

    addDhcp6ServerRes, err := client.ApiCallSimple("set-dhcp6-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addDhcp6ServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addDhcp6ServerRes.Success {
            errMsg = addDhcp6ServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addDhcp6ServerRes.GetData()
        }

        debugLogOperation(
            "dhcp6-server",        // resource type
            "create",                       // operation
            "set-dhcp6-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add dhcp6-server: %v", err)
    }
    if !addDhcp6ServerRes.Success {
        if addDhcp6ServerRes.ErrorMsg != "" {
            return fmt.Errorf(addDhcp6ServerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("dhcp6-server-" + acctest.RandString(10)))
    return readGaiaDhcp6Server(d, m)
}

func readGaiaDhcp6Server(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showDhcp6ServerRes, err := client.ApiCallSimple("show-dhcp6-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showDhcp6ServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showDhcp6ServerRes.Success {
            errMsg = showDhcp6ServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showDhcp6ServerRes.GetData()
        }

        debugLogOperation(
            "dhcp6-server",        // resource type
            "read",                       // operation
            "show-dhcp6-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show dhcp6-server: %v", err)
    }
    if !showDhcp6ServerRes.Success {
        if data := showDhcp6ServerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showDhcp6ServerRes.ErrorMsg)
    }

    dhcp6Server := showDhcp6ServerRes.GetData()

    log.Println("Read Dhcp6Server - Show JSON = ", dhcp6Server)

    if v, exists := dhcp6Server["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := dhcp6Server["subnets"]; exists {
        if rawSubnets, ok := v.([]interface{}); ok {
            subnets := make([]interface{}, 0, len(rawSubnets))
            for _, s := range rawSubnets {
                raw, ok := s.(map[string]interface{})
                if !ok {
                    continue
                }
                subnet := map[string]interface{}{
                    "subnet":  fmt.Sprintf("%v", raw["subnet"]),
                    "enabled": raw["enabled"],
                    "prefix": func() int {
                        if f, ok := raw["prefix"].(float64); ok { return int(f) }
                        return 0
                    }(),
                    "max_lease": func() int {
                        if f, ok := raw["max-lease"].(float64); ok { return int(f) }
                        return 0
                    }(),
                    "default_lease": func() int {
                        if f, ok := raw["default-lease"].(float64); ok { return int(f) }
                        return 0
                    }(),
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
    d.SetId(d.Id())
    return nil
}

func updateGaiaDhcp6Server(d *schema.ResourceData, m interface{}) error {

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
            if v, ok := d.GetOk(fmt.Sprintf("subnets.%d.prefix", i)); ok {
                itemMap["prefix"] = v.(int)
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

    setDhcp6ServerRes, err := client.ApiCallSimple("set-dhcp6-server", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setDhcp6ServerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setDhcp6ServerRes.Success {
            errMsg = setDhcp6ServerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setDhcp6ServerRes.GetData()
        }

        debugLogOperation(
            "dhcp6-server",        // resource type
            "update",                       // operation
            "set-dhcp6-server",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set dhcp6-server: %v", err)
    }
    if !setDhcp6ServerRes.Success {
        return fmt.Errorf(setDhcp6ServerRes.ErrorMsg)
    }

    return readGaiaDhcp6Server(d, m)
}

func deleteGaiaDhcp6Server(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    