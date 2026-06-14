package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowPimSparseModeStats() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowPimSparseModeStats,
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
            "pim_not_enabled_on_link_to_src": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_grp_to_rp_mapping": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "border_router_mismatch": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_route_to_dr": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_route_to_src": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "src_net_mismatch": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "not_dr_on_link_to_src": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "register_encaps_rp_not_local": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_receiver_for_sg": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_matching_active_entry": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_route_to_rp": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "sg_iif_not_active": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "msg_from_non_rpf_nbr": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "msg_from_less_preferred_bsr": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_route_to_sender": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "crp_msg_sent_to_non_elected_bsr": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "join_denied_af_mismatch": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "join_denied_not_dr": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "no_pim_on_link_to_receiver": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "join_denied_no_rp": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowPimSparseModeStats(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-pim-sparse-mode-stats - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-pim-sparse-mode-stats", payload)
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
            "pim-sparse-mode-stats",        // resource type
            "read",                       // operation
            "show-pim-sparse-mode-stats",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-pim-sparse-mode-stats: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["pim-not-enabled-on-link-to-src"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("pim_not_enabled_on_link_to_src", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-grp-to-rp-mapping"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_grp_to_rp_mapping", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["border-router-mismatch"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("border_router_mismatch", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-route-to-dr"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_route_to_dr", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-route-to-src"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_route_to_src", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["src-net-mismatch"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("src_net_mismatch", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["not-dr-on-link-to-src"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("not_dr_on_link_to_src", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["register-encaps-rp-not-local"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("register_encaps_rp_not_local", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-receiver-for-sg"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_receiver_for_sg", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-matching-active-entry"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_matching_active_entry", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-route-to-rp"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_route_to_rp", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["sg-iif-not-active"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("sg_iif_not_active", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["msg-from-non-rpf-nbr"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("msg_from_non_rpf_nbr", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["msg-from-less-preferred-bsr"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("msg_from_less_preferred_bsr", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-route-to-sender"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_route_to_sender", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["crp-msg-sent-to-non-elected-bsr"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("crp_msg_sent_to_non_elected_bsr", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["join-denied-af-mismatch"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("join_denied_af_mismatch", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["join-denied-not-dr"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("join_denied_not_dr", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["no-pim-on-link-to-receiver"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("no_pim_on_link_to_receiver", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["join-denied-no-rp"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("join_denied_no_rp", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-pim-sparse-mode-stats-" + acctest.RandString(10)))
    return nil
}

