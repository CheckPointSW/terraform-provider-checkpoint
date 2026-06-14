package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfEvents() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfEvents,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Existing OSPFv2 Instance`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "1583_mode_changes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "abr_changes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "asbr_changes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "dr_elections_run": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "interface_down": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "interface_up": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "router_id_changes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "self_lsa_promotions": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "virtual_interface_down": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "virtual_interface_up": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowOspfEvents(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-events - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-events", payload)
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
            "ospf-events",        // resource type
            "read",                       // operation
            "show-ospf-events",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-events: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["1583-mode-changes"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("1583_mode_changes", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["abr-changes"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("abr_changes", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["asbr-changes"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("asbr_changes", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["dr-elections-run"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("dr_elections_run", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["interface-down"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("interface_down", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["interface-up"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("interface_up", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["router-id-changes"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("router_id_changes", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["self-lsa-promotions"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("self_lsa_promotions", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["virtual-interface-down"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("virtual_interface_down", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["virtual-interface-up"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("virtual_interface_up", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ospf-events-" + acctest.RandString(10)))
    return nil
}

