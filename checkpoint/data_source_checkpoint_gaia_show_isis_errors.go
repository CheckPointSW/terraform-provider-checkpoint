package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowIsisErrors() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowIsisErrors,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The instance to be queried`,
            },
            "summary": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Filter the results`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "global_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "no_interface_config": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interface_down": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "received_own_packet": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "protocol_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "packet_size_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "misformatted_header": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "system_id_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "max_areas": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_packet_type": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "wrong_multicast": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "p2p_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "level_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authentication_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "malformed_authentication": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "hello_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pdu_length_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tlv_too_long": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_tlv": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "area_too_long": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_area_match": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "adj_state_tlv_in_lan": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_adj_state_tlv_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "adj_state_sysid_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "adj_state_circuit_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "is_neighbors_tlv_in_p2p": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_is_neighbors_tvl_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_nlpid": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "nlpid_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ipv4_interfaces_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "duplicate_ipv4_address": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_ipv4_subnet_match": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv4_support_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ipv6_interfaces_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "duplicate_ipv6_address": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_ipv6_subnet_match": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ipv6_support_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "neighbor_mac_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_adj_state_tlv_in_p2p": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_level_match_in_p2p": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "p2p_neighbor_sysid_change": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "p2p_neighbor_level_change": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authentication_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "malformed_authentication": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mtid_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "lsp_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pdu_length_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_adjacency": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_checksum": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "checksum_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "received_own_lsp": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_areas_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "areas_in_nonzero_fragment": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_nlpid": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ipv4_interfaces_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ipv6_interfaces_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_is_reach_tlv_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ext_is_reach_tlv_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ip_reach_tlv_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ext_ip_reach_tlv_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_ipv6_reach_tlv_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_tlv": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_sub_tlv": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authentication_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "malformed_authentication": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_mtid": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mt_is_reach_with_mtid_zero": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_mt_ip_reach_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_mt_is_reach_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_mt_ipv6_reach_len": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "mtid_in_nonzero_fragment": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "csnp_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pdu_length_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_adjacency": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "sender_is_not_dis": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_tlv": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_tlv_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp_below_minimum": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "lsp_above_maximum": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authentication_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "malformed_authentication": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "psnp_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pdu_length_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "no_adjacency": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "received_while_not_dis": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "unknown_tlv": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "bad_tlv_length": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "authentication_mismatch": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "malformed_authentication": {
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

func readGaiaShowIsisErrors(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("summary"); ok {
        payload["summary"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-isis-errors - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-isis-errors", payload)
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
            "isis-errors",        // resource type
            "read",                       // operation
            "show-isis-errors",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-isis-errors: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["global-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("global_errors", []interface{}{map[string]interface{}{
                "no_interface_config": func() int { if f, ok := _m["no-interface-config"].(float64); ok { return int(f) }; return 0 }(),
                "interface_down": func() int { if f, ok := _m["interface-down"].(float64); ok { return int(f) }; return 0 }(),
                "received_own_packet": func() int { if f, ok := _m["received-own-packet"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["protocol-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("protocol_errors", []interface{}{map[string]interface{}{
                "packet_size_mismatch": func() int { if f, ok := _m["packet-size-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "misformatted_header": func() int { if f, ok := _m["misformatted-header"].(float64); ok { return int(f) }; return 0 }(),
                "system_id_length": func() int { if f, ok := _m["system-id-length"].(float64); ok { return int(f) }; return 0 }(),
                "max_areas": func() int { if f, ok := _m["max-areas"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_packet_type": func() int { if f, ok := _m["unknown-packet-type"].(float64); ok { return int(f) }; return 0 }(),
                "wrong_multicast": func() int { if f, ok := _m["wrong-multicast"].(float64); ok { return int(f) }; return 0 }(),
                "p2p_mismatch": func() int { if f, ok := _m["p2p-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "level_mismatch": func() int { if f, ok := _m["level-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "authentication_mismatch": func() int { if f, ok := _m["authentication-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "malformed_authentication": func() int { if f, ok := _m["malformed-authentication"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["hello-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("hello_errors", []interface{}{map[string]interface{}{
                "pdu_length_mismatch": func() int { if f, ok := _m["pdu-length-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "tlv_too_long": func() int { if f, ok := _m["tlv-too-long"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_tlv": func() int { if f, ok := _m["unknown-tlv"].(float64); ok { return int(f) }; return 0 }(),
                "area_too_long": func() int { if f, ok := _m["area-too-long"].(float64); ok { return int(f) }; return 0 }(),
                "no_area_match": func() int { if f, ok := _m["no-area-match"].(float64); ok { return int(f) }; return 0 }(),
                "adj_state_tlv_in_lan": func() int { if f, ok := _m["adj-state-tlv-in-lan"].(float64); ok { return int(f) }; return 0 }(),
                "bad_adj_state_tlv_length": func() int { if f, ok := _m["bad-adj-state-tlv-length"].(float64); ok { return int(f) }; return 0 }(),
                "adj_state_sysid_mismatch": func() int { if f, ok := _m["adj-state-sysid-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "adj_state_circuit_mismatch": func() int { if f, ok := _m["adj-state-circuit-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "is_neighbors_tlv_in_p2p": func() int { if f, ok := _m["is-neighbors-tlv-in-p2p"].(float64); ok { return int(f) }; return 0 }(),
                "bad_is_neighbors_tvl_length": func() int { if f, ok := _m["bad-is-neighbors-tvl-length"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_nlpid": func() int { if f, ok := _m["unknown-nlpid"].(float64); ok { return int(f) }; return 0 }(),
                "nlpid_mismatch": func() int { if f, ok := _m["nlpid-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ipv4_interfaces_length": func() int { if f, ok := _m["bad-ipv4-interfaces-length"].(float64); ok { return int(f) }; return 0 }(),
                "duplicate_ipv4_address": func() int { if f, ok := _m["duplicate-ipv4-address"].(float64); ok { return int(f) }; return 0 }(),
                "no_ipv4_subnet_match": func() int { if f, ok := _m["no-ipv4-subnet-match"].(float64); ok { return int(f) }; return 0 }(),
                "ipv4_support_mismatch": func() int { if f, ok := _m["ipv4-support-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ipv6_interfaces_length": func() int { if f, ok := _m["bad-ipv6-interfaces-length"].(float64); ok { return int(f) }; return 0 }(),
                "duplicate_ipv6_address": func() int { if f, ok := _m["duplicate-ipv6-address"].(float64); ok { return int(f) }; return 0 }(),
                "no_ipv6_subnet_match": func() int { if f, ok := _m["no-ipv6-subnet-match"].(float64); ok { return int(f) }; return 0 }(),
                "ipv6_support_mismatch": func() int { if f, ok := _m["ipv6-support-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "neighbor_mac_mismatch": func() int { if f, ok := _m["neighbor-mac-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "no_adj_state_tlv_in_p2p": func() int { if f, ok := _m["no-adj-state-tlv-in-p2p"].(float64); ok { return int(f) }; return 0 }(),
                "no_level_match_in_p2p": func() int { if f, ok := _m["no-level-match-in-p2p"].(float64); ok { return int(f) }; return 0 }(),
                "p2p_neighbor_sysid_change": func() int { if f, ok := _m["p2p-neighbor-sysid-change"].(float64); ok { return int(f) }; return 0 }(),
                "p2p_neighbor_level_change": func() int { if f, ok := _m["p2p-neighbor-level-change"].(float64); ok { return int(f) }; return 0 }(),
                "authentication_mismatch": func() int { if f, ok := _m["authentication-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "malformed_authentication": func() int { if f, ok := _m["malformed-authentication"].(float64); ok { return int(f) }; return 0 }(),
                "mtid_mismatch": func() int { if f, ok := _m["mtid-mismatch"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["lsp-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("lsp_errors", []interface{}{map[string]interface{}{
                "pdu_length_mismatch": func() int { if f, ok := _m["pdu-length-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "no_adjacency": func() int { if f, ok := _m["no-adjacency"].(float64); ok { return int(f) }; return 0 }(),
                "bad_checksum": func() int { if f, ok := _m["bad-checksum"].(float64); ok { return int(f) }; return 0 }(),
                "checksum_mismatch": func() int { if f, ok := _m["checksum-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "received_own_lsp": func() int { if f, ok := _m["received-own-lsp"].(float64); ok { return int(f) }; return 0 }(),
                "bad_areas_len": func() int { if f, ok := _m["bad-areas-len"].(float64); ok { return int(f) }; return 0 }(),
                "areas_in_nonzero_fragment": func() int { if f, ok := _m["areas-in-nonzero-fragment"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_nlpid": func() int { if f, ok := _m["unknown-nlpid"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ipv4_interfaces_len": func() int { if f, ok := _m["bad-ipv4-interfaces-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ipv6_interfaces_len": func() int { if f, ok := _m["bad-ipv6-interfaces-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_is_reach_tlv_len": func() int { if f, ok := _m["bad-is-reach-tlv-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ext_is_reach_tlv_len": func() int { if f, ok := _m["bad-ext-is-reach-tlv-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ip_reach_tlv_len": func() int { if f, ok := _m["bad-ip-reach-tlv-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ext_ip_reach_tlv_len": func() int { if f, ok := _m["bad-ext-ip-reach-tlv-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_ipv6_reach_tlv_len": func() int { if f, ok := _m["bad-ipv6-reach-tlv-len"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_tlv": func() int { if f, ok := _m["unknown-tlv"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_sub_tlv": func() int { if f, ok := _m["unknown-sub-tlv"].(float64); ok { return int(f) }; return 0 }(),
                "authentication_mismatch": func() int { if f, ok := _m["authentication-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "malformed_authentication": func() int { if f, ok := _m["malformed-authentication"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_mtid": func() int { if f, ok := _m["unknown-mtid"].(float64); ok { return int(f) }; return 0 }(),
                "mt_is_reach_with_mtid_zero": func() int { if f, ok := _m["mt-is-reach-with-mtid-zero"].(float64); ok { return int(f) }; return 0 }(),
                "bad_mt_ip_reach_len": func() int { if f, ok := _m["bad-mt-ip-reach-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_mt_is_reach_len": func() int { if f, ok := _m["bad-mt-is-reach-len"].(float64); ok { return int(f) }; return 0 }(),
                "bad_mt_ipv6_reach_len": func() int { if f, ok := _m["bad-mt-ipv6-reach-len"].(float64); ok { return int(f) }; return 0 }(),
                "mtid_in_nonzero_fragment": func() int { if f, ok := _m["mtid-in-nonzero-fragment"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["csnp-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("csnp_errors", []interface{}{map[string]interface{}{
                "pdu_length_mismatch": func() int { if f, ok := _m["pdu-length-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "no_adjacency": func() int { if f, ok := _m["no-adjacency"].(float64); ok { return int(f) }; return 0 }(),
                "sender_is_not_dis": func() int { if f, ok := _m["sender-is-not-dis"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_tlv": func() int { if f, ok := _m["unknown-tlv"].(float64); ok { return int(f) }; return 0 }(),
                "bad_tlv_length": func() int { if f, ok := _m["bad-tlv-length"].(float64); ok { return int(f) }; return 0 }(),
                "lsp_below_minimum": func() int { if f, ok := _m["lsp-below-minimum"].(float64); ok { return int(f) }; return 0 }(),
                "lsp_above_maximum": func() int { if f, ok := _m["lsp-above-maximum"].(float64); ok { return int(f) }; return 0 }(),
                "authentication_mismatch": func() int { if f, ok := _m["authentication-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "malformed_authentication": func() int { if f, ok := _m["malformed-authentication"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["psnp-errors"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("psnp_errors", []interface{}{map[string]interface{}{
                "pdu_length_mismatch": func() int { if f, ok := _m["pdu-length-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "no_adjacency": func() int { if f, ok := _m["no-adjacency"].(float64); ok { return int(f) }; return 0 }(),
                "received_while_not_dis": func() int { if f, ok := _m["received-while-not-dis"].(float64); ok { return int(f) }; return 0 }(),
                "unknown_tlv": func() int { if f, ok := _m["unknown-tlv"].(float64); ok { return int(f) }; return 0 }(),
                "bad_tlv_length": func() int { if f, ok := _m["bad-tlv-length"].(float64); ok { return int(f) }; return 0 }(),
                "authentication_mismatch": func() int { if f, ok := _m["authentication-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                "malformed_authentication": func() int { if f, ok := _m["malformed-authentication"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-isis-errors-" + acctest.RandString(10)))
    return nil
}

