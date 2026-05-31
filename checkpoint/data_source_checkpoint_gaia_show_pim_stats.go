package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPimStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPimStats,
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
            "objects": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "packet_type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "received": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rx_errors": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "transmit": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tx_errors": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowPimStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-pim-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-pim-stats", payload)
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
            "pim-stats",        // resource type
            "read",                       // operation
            "show-pim-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-pim-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["objects"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "packet_type": func() string { if _v, _ok := m["packet-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "received": func() int { if f, ok := m["received"].(float64); ok { return int(f) }; return 0 }(),
                        "rx_errors": func() int { if f, ok := m["rx-errors"].(float64); ok { return int(f) }; return 0 }(),
                        "transmit": func() int { if f, ok := m["transmit"].(float64); ok { return int(f) }; return 0 }(),
                        "tx_errors": func() int { if f, ok := m["tx-errors"].(float64); ok { return int(f) }; return 0 }(),
                    }
                }
            }
            d.Set("objects", mapped)
        }
    } else {
        d.Set("objects", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-pim-stats-" + acctest.RandString(10)))
    return nil
}

