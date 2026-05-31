package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConfigurationBgpConfederation() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConfigurationBgpConfederation,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_as": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Specify the Routing Domain identifier of the Confederation peer group.<br><br>If the peer group specified is the local Routing Domain, it will run IBGP in a full mesh (just as an internal peer group normally would in non-Confederation mode). Otherwise, if an external Routing Domain within the Confederation is specified, the peer group will run a modified version of eBGP, which preserves route metrics and other BGP attributes.<br><br>The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "description": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "interface_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "local_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "med": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_nexthop_self": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "outdelay": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "protocol_list": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
        },
    }
}

func readGaiaShowConfigurationBgpConfederation(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_as"); ok {
        payload["member-as"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-configuration-bgp-confederation - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-configuration-bgp-confederation", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "configuration-bgp-confederation",        // resource type
            "read",                       // operation
            "show-configuration-bgp-confederation",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-configuration-bgp-confederation: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["description"]; exists {
        d.Set("description", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["interface-list"]; exists {
        d.Set("interface_list", v.([]interface{}))
    } else {
        d.Set("interface_list", []interface{}{})
    }
    if v, exists := commandRes.GetData()["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["med"]; exists {
        d.Set("med", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-as"]; exists {
        d.Set("member_as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enable-nexthop-self"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_nexthop_self", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_nexthop_self", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["outdelay"]; exists {
        d.Set("outdelay", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["protocol-list"]; exists {
        switch val := v.(type) {
        case []interface{}:
            d.Set("protocol_list", val)
        case string:
            d.Set("protocol_list", []interface{}{val})
        default:
            d.Set("protocol_list", []interface{}{})
        }
    } else {
        d.Set("protocol_list", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-configuration-bgp-confederation-" + acctest.RandString(10)))
    return nil
}

