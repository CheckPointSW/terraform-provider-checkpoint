package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfPackets() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfPackets,
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
            "dd_rx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "dd_tx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "hello_rx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "hello_tx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsack_rx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsack_tx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsr_rx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsr_tx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsu_rx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "lsu_tx": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowOspfPackets(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-packets - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-packets", payload)
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
            "ospf-packets",        // resource type
            "read",                       // operation
            "show-ospf-packets",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-packets: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["dd-rx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("dd_rx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["dd-tx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("dd_tx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["hello-rx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("hello_rx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["hello-tx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("hello_tx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsack-rx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsack_rx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsack-tx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsack_tx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsr-rx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsr_rx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsr-tx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsr_tx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsu-rx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsu_rx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["lsu-tx"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("lsu_tx", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ospf-packets-" + acctest.RandString(10)))
    return nil
}

