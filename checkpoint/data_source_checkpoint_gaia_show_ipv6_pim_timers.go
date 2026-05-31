package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIpv6PimTimers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIpv6PimTimers,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "hello": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "mrt": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "join_prune": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "active_rpset": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "candidate_rp": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "bootstrap": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowIpv6PimTimers(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ipv6-pim-timers - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ipv6-pim-timers", payload)
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
            "ipv6-pim-timers",        // resource type
            "read",                       // operation
            "show-ipv6-pim-timers",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ipv6-pim-timers: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["hello"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("hello", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["mrt"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("mrt", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["join-prune"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("join_prune", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["active-rpset"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("active_rpset", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["candidate-rp"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("candidate_rp", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["bootstrap"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("bootstrap", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ipv6-pim-timers-" + acctest.RandString(10)))
    return nil
}

