package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaDhcp6Config() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaDhcp6Config,
        Read:   readGaiaDhcp6Config,
        Update: updateGaiaDhcp6Config,
        Delete: deleteGaiaDhcp6Config,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "prefix_delegation_options": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `General configuration for the prefix-delegation feature.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interface": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The interface on which to send prefix-delegation request packets to the prefix-delegation DHCP server.`,
                        },
                        "method": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The method of performing the delegation of the received subnets. Each method balances automation with granularity.                       <br><b>Manual</b> - Only configure client interfaces set to receive IPv6 via Prefix-Delegation.                       <br><b>Router Discovery</b> - In addition to IPv6, also automatically configure Router Discovery protocol                         on configured interfaces.                       <br><b>DHCPv6</b> - In addition to IPv6, also automatically configure the DHCPv6 Server feature on                         the new subnets and configure the Router Discory protocol with Managed Configuration flag on.`,
                        },
                        "suffix_pools": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Pools of IPv6 suffixes to use with DHCPv6 delegation method.                        These will be used to automatically configure IPv6 pools for each subnet in the DHCPv6 server feature.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "start": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The first IPv6 address of the suffix range.`,
                                    },
                                    "end": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The last IPv6 address of the suffix range.`,
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies whether to include or exclude this range of IPv6 suffixes in the IP pools.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "client_mode": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The working mode of the DHCPv6 client in this system.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaDhcp6Config(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("prefix_delegation_options"); len(v.([]interface{})) > 0 {
        _ = v
        prefixdelegationoptionsMap := make(map[string]interface{})
        if v, ok := d.GetOk("prefix_delegation_options.0.interface"); ok {
            prefixdelegationoptionsMap["interface"] = v.(string)
        }
        if v, ok := d.GetOk("prefix_delegation_options.0.method"); ok {
            prefixdelegationoptionsMap["method"] = v.(string)
        }
        if v, ok := d.GetOk("prefix_delegation_options.0.suffix_pools"); ok {
            suffixpoolsList := v.([]interface{})
            suffixpoolsArray := make([]interface{}, 0, len(suffixpoolsList))
            for i := range suffixpoolsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.start", i)); ok {
                    itemMap["start"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.end", i)); ok {
                    itemMap["end"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.type", i)); ok {
                    itemMap["type"] = v.(string)
                }
                if len(itemMap) > 0 {
                    suffixpoolsArray = append(suffixpoolsArray, itemMap)
                }
            }
            if len(suffixpoolsArray) > 0 {
                prefixdelegationoptionsMap["suffix-pools"] = suffixpoolsArray
            }
        }
        if len(prefixdelegationoptionsMap) > 0 {
            payload["prefix-delegation-options"] = prefixdelegationoptionsMap
        }
    }

    if v, ok := d.GetOk("client_mode"); ok {
        payload["client-mode"] = v.(string)
    }

    log.Println("Create Dhcp6Config - Map = ", payload)

    addDhcp6ConfigRes, err := client.ApiCallSimple("set-dhcp6-config", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addDhcp6ConfigRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addDhcp6ConfigRes.Success {
            errMsg = addDhcp6ConfigRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addDhcp6ConfigRes.GetData()
        }

        debugLogOperation(
            "dhcp6-config",        // resource type
            "create",                       // operation
            "set-dhcp6-config",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add dhcp6-config: %v", err)
    }
    if !addDhcp6ConfigRes.Success {
        if addDhcp6ConfigRes.ErrorMsg != "" {
            return fmt.Errorf(addDhcp6ConfigRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("dhcp6-config-" + acctest.RandString(10)))
    return readGaiaDhcp6Config(d, m)
}

func readGaiaDhcp6Config(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showDhcp6ConfigRes, err := client.ApiCallSimple("show-dhcp6-config", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showDhcp6ConfigRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showDhcp6ConfigRes.Success {
            errMsg = showDhcp6ConfigRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showDhcp6ConfigRes.GetData()
        }

        debugLogOperation(
            "dhcp6-config",        // resource type
            "read",                       // operation
            "show-dhcp6-config",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show dhcp6-config: %v", err)
    }
    if !showDhcp6ConfigRes.Success {
        if data := showDhcp6ConfigRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showDhcp6ConfigRes.ErrorMsg)
    }

    dhcp6Config := showDhcp6ConfigRes.GetData()

    log.Println("Read Dhcp6Config - Show JSON = ", dhcp6Config)

    if v, exists := dhcp6Config["client-mode"]; exists {
        d.Set("client_mode", fmt.Sprintf("%v", v))
    }
    if v, exists := dhcp6Config["prefix-delegation-options"]; exists {
        d.Set("prefix_delegation_options", v)
    }
    if v, exists := dhcp6Config["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaDhcp6Config(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("prefix_delegation_options"); len(v.([]interface{})) > 0 {
        _ = v
        prefixdelegationoptionsMap := make(map[string]interface{})
        if v, ok := d.GetOk("prefix_delegation_options.0.interface"); ok {
            prefixdelegationoptionsMap["interface"] = v.(string)
        }
        if v, ok := d.GetOk("prefix_delegation_options.0.method"); ok {
            prefixdelegationoptionsMap["method"] = v.(string)
        }
        if v, ok := d.GetOk("prefix_delegation_options.0.suffix_pools"); ok {
            suffixpoolsList := v.([]interface{})
            suffixpoolsArray := make([]interface{}, 0, len(suffixpoolsList))
            for i := range suffixpoolsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.start", i)); ok {
                    itemMap["start"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.end", i)); ok {
                    itemMap["end"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("prefix_delegation_options.0.suffix_pools.%d.type", i)); ok {
                    itemMap["type"] = v.(string)
                }
                if len(itemMap) > 0 {
                    suffixpoolsArray = append(suffixpoolsArray, itemMap)
                }
            }
            if len(suffixpoolsArray) > 0 {
                prefixdelegationoptionsMap["suffix-pools"] = suffixpoolsArray
            }
        }
        if len(prefixdelegationoptionsMap) > 0 {
            payload["prefix-delegation-options"] = prefixdelegationoptionsMap
        }
    }

    if v, ok := d.GetOk("client_mode"); ok {
        payload["client-mode"] = v.(string)
    }

    setDhcp6ConfigRes, err := client.ApiCallSimple("set-dhcp6-config", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setDhcp6ConfigRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setDhcp6ConfigRes.Success {
            errMsg = setDhcp6ConfigRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setDhcp6ConfigRes.GetData()
        }

        debugLogOperation(
            "dhcp6-config",        // resource type
            "update",                       // operation
            "set-dhcp6-config",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set dhcp6-config: %v", err)
    }
    if !setDhcp6ConfigRes.Success {
        return fmt.Errorf(setDhcp6ConfigRes.ErrorMsg)
    }

    return readGaiaDhcp6Config(d, m)
}

func deleteGaiaDhcp6Config(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    