package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowBgpSummary() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowBgpSummary,
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
            "enabled": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "confederation_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "routing_domain_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "as": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "default_weight": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv4_route_rank": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_route_rank": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "default_med": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_ecmp": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "enable_synchronization": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowBgpSummary(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-bgp-summary - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-bgp-summary", payload)
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
            "bgp-summary",        // resource type
            "read",                       // operation
            "show-bgp-summary",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-bgp-summary: %v", err)
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
    if v, exists := commandRes.GetData()["confederation-id"]; exists {
        d.Set("confederation_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["routing-domain-id"]; exists {
        d.Set("routing_domain_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["as"]; exists {
        d.Set("as", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["default-weight"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("default_weight", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["ipv4-route-rank"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("ipv4_route_rank", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["ipv6-route-rank"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("ipv6_route_rank", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["default-med"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("default_med", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["enable-ecmp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ecmp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ecmp", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["enable-synchronization"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_synchronization", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_synchronization", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-bgp-summary-" + acctest.RandString(10)))
    return nil
}

