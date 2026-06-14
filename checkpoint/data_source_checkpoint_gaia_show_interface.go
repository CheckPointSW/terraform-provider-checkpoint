package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowInterface() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv4_mask_length": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_autoconfig": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "comments": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_mask_length": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_local_link_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute show-interface - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-interface", payload)
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
            "interface",        // resource type
            "read",                       // operation
            "show-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-interface: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["ipv4-address"]; exists {
        d.Set("ipv4_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ipv6-address"]; exists {
        d.Set("ipv6_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ipv6-mask-length"]; exists {
        d.Set("ipv6_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["type"]; exists {
        d.Set("type", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-interface-" + acctest.RandString(10)))
    return nil
}

