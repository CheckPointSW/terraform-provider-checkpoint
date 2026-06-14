package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPimNeighbor() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPimNeighbor,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The neighbor address.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "dr_priority": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "expires": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "gen_id": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "holdtime": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "interface": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowPimNeighbor(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-pim-neighbor - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-pim-neighbor", payload)
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
            "pim-neighbor",        // resource type
            "read",                       // operation
            "show-pim-neighbor",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-pim-neighbor: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["dr-priority"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("dr_priority", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["expires"]; exists {
        d.Set("expires", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["gen-id"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("gen_id", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["holdtime"]; exists {
        d.Set("holdtime", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-pim-neighbor-" + acctest.RandString(10)))
    return nil
}

