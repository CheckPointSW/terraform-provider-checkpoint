package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIpv6PimSummary() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIpv6PimSummary,
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
            "instance": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "mode": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "join_prune_amount": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowIpv6PimSummary(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ipv6-pim-summary - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ipv6-pim-summary", payload)
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
            "ipv6-pim-summary",        // resource type
            "read",                       // operation
            "show-ipv6-pim-summary",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ipv6-pim-summary: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["instance"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("instance", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["join-prune-amount"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("join_prune_amount", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ipv6-pim-summary-" + acctest.RandString(10)))
    return nil
}

