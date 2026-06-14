package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfErrors() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfErrors,
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
            "error_type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Error Type`,
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
                        "ip_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "bad_destination": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_source": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_such_index": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "own_packet": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "protocol": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "size": {
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
                                    "no_ospf": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "instance_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dd_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ase_in_stub": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_ls_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dd_duplicate_router_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dd_too_low": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "master_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mtu": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "not_duplicate": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "options_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "runt": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "slave_seq": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type7_in_non_nssa": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "hello_protocol_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "bad_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "dead_interval_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "external_option_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hello_duplicate_router_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "hello_timer_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "network_mask_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "nssa_option_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "runt": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "lsack_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "bad_ls_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsack_duplicate_router_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsack_too_low": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "question_ack": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "lsr_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "bad_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_state": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "empty_request": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsr_duplicate_router_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "lsu_errors": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ase_in_stub": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_ase_lsa_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_checksum": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_ls_req": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_ls_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_network_lsa_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_router_lsa_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_summary_lsa_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_type7_lsa_size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "invalid_seq_num": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsu_duplicate_router_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsu_too_low": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "lsu_too_new": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "runt": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "seq_num_wrap": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "summary_in_total_stub": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type7_in_non_nssa": {
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
                                    "area_id_mismatch": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "auth_crypto_seq": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "auth_key_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "auth_key_time": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "auth_key_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_area_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bad_destination": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "checksum": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "if_down": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_neighbor": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_virtual_neighbor": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "non_local": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "packet_type": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "passive_interface": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "size": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "tx": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "version": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "virtual_link": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "zero_rid": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowOspfErrors(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("error_type"); ok {
        payload["error-type"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-errors - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-errors", payload)
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
            "ospf-errors",        // resource type
            "read",                       // operation
            "show-ospf-errors",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-errors: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["global-errors"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "ip_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ip-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "bad_destination": func() int { if f, ok := _sgm["bad-destination"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_source": func() int { if f, ok := _sgm["bad-source"].(float64); ok { return int(f) }; return 0 }(),
                                            "no_such_index": func() int { if f, ok := _sgm["no-such-index"].(float64); ok { return int(f) }; return 0 }(),
                                            "own_packet": func() int { if f, ok := _sgm["own-packet"].(float64); ok { return int(f) }; return 0 }(),
                                            "protocol": func() int { if f, ok := _sgm["protocol"].(float64); ok { return int(f) }; return 0 }(),
                                            "size": func() int { if f, ok := _sgm["size"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "protocol_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["protocol-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "no_ospf": func() int { if f, ok := _sgm["no-ospf"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                    }
                }
            }
            d.Set("global_errors", mapped)
        }
    } else {
        d.Set("global_errors", []interface{}{})
    }
    if v, exists := commandRes.GetData()["instance-errors"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "dd_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["dd-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ase_in_stub": func() int { if f, ok := _sgm["ase-in-stub"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_ls_type": func() int { if f, ok := _sgm["bad-ls-type"].(float64); ok { return int(f) }; return 0 }(),
                                            "dd_duplicate_router_id": func() int { if f, ok := _sgm["dd-duplicate-router-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "dd_too_low": func() int { if f, ok := _sgm["dd-too-low"].(float64); ok { return int(f) }; return 0 }(),
                                            "master_mismatch": func() int { if f, ok := _sgm["master-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "mtu": func() int { if f, ok := _sgm["mtu"].(float64); ok { return int(f) }; return 0 }(),
                                            "not_duplicate": func() int { if f, ok := _sgm["not-duplicate"].(float64); ok { return int(f) }; return 0 }(),
                                            "options_mismatch": func() int { if f, ok := _sgm["options-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "runt": func() int { if f, ok := _sgm["runt"].(float64); ok { return int(f) }; return 0 }(),
                                            "slave_seq": func() int { if f, ok := _sgm["slave-seq"].(float64); ok { return int(f) }; return 0 }(),
                                            "type7_in_non_nssa": func() int { if f, ok := _sgm["type7-in-non-nssa"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "hello_protocol_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["hello-protocol-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "bad_size": func() int { if f, ok := _sgm["bad-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "dead_interval_mismatch": func() int { if f, ok := _sgm["dead-interval-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "external_option_mismatch": func() int { if f, ok := _sgm["external-option-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "hello_duplicate_router_id": func() int { if f, ok := _sgm["hello-duplicate-router-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "hello_timer_mismatch": func() int { if f, ok := _sgm["hello-timer-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "network_mask_mismatch": func() int { if f, ok := _sgm["network-mask-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "nssa_option_mismatch": func() int { if f, ok := _sgm["nssa-option-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "runt": func() int { if f, ok := _sgm["runt"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "lsack_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["lsack-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "bad_ls_type": func() int { if f, ok := _sgm["bad-ls-type"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_size": func() int { if f, ok := _sgm["bad-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsack_duplicate_router_id": func() int { if f, ok := _sgm["lsack-duplicate-router-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsack_too_low": func() int { if f, ok := _sgm["lsack-too-low"].(float64); ok { return int(f) }; return 0 }(),
                                            "question_ack": func() int { if f, ok := _sgm["question-ack"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "lsr_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["lsr-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "bad_size": func() int { if f, ok := _sgm["bad-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_state": func() int { if f, ok := _sgm["bad-state"].(float64); ok { return int(f) }; return 0 }(),
                                            "empty_request": func() int { if f, ok := _sgm["empty-request"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsr_duplicate_router_id": func() int { if f, ok := _sgm["lsr-duplicate-router-id"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "lsu_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["lsu-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "ase_in_stub": func() int { if f, ok := _sgm["ase-in-stub"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_ase_lsa_size": func() int { if f, ok := _sgm["bad-ase-lsa-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_checksum": func() int { if f, ok := _sgm["bad-checksum"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_ls_req": func() int { if f, ok := _sgm["bad-ls-req"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_ls_type": func() int { if f, ok := _sgm["bad-ls-type"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_network_lsa_size": func() int { if f, ok := _sgm["bad-network-lsa-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_router_lsa_size": func() int { if f, ok := _sgm["bad-router-lsa-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_summary_lsa_size": func() int { if f, ok := _sgm["bad-summary-lsa-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_type7_lsa_size": func() int { if f, ok := _sgm["bad-type7-lsa-size"].(float64); ok { return int(f) }; return 0 }(),
                                            "invalid_seq_num": func() int { if f, ok := _sgm["invalid-seq-num"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsu_duplicate_router_id": func() int { if f, ok := _sgm["lsu-duplicate-router-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsu_too_low": func() int { if f, ok := _sgm["lsu-too-low"].(float64); ok { return int(f) }; return 0 }(),
                                            "lsu_too_new": func() int { if f, ok := _sgm["lsu-too-new"].(float64); ok { return int(f) }; return 0 }(),
                                            "runt": func() int { if f, ok := _sgm["runt"].(float64); ok { return int(f) }; return 0 }(),
                                            "seq_num_wrap": func() int { if f, ok := _sgm["seq-num-wrap"].(float64); ok { return int(f) }; return 0 }(),
                                            "summary_in_total_stub": func() int { if f, ok := _sgm["summary-in-total-stub"].(float64); ok { return int(f) }; return 0 }(),
                                            "type7_in_non_nssa": func() int { if f, ok := _sgm["type7-in-non-nssa"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "protocol_errors": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["protocol-errors"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "area_id_mismatch": func() int { if f, ok := _sgm["area-id-mismatch"].(float64); ok { return int(f) }; return 0 }(),
                                            "auth_crypto_seq": func() int { if f, ok := _sgm["auth-crypto-seq"].(float64); ok { return int(f) }; return 0 }(),
                                            "auth_key_id": func() int { if f, ok := _sgm["auth-key-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "auth_key_time": func() int { if f, ok := _sgm["auth-key-time"].(float64); ok { return int(f) }; return 0 }(),
                                            "auth_key_type": func() int { if f, ok := _sgm["auth-key-type"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_area_id": func() int { if f, ok := _sgm["bad-area-id"].(float64); ok { return int(f) }; return 0 }(),
                                            "bad_destination": func() int { if f, ok := _sgm["bad-destination"].(float64); ok { return int(f) }; return 0 }(),
                                            "checksum": func() int { if f, ok := _sgm["checksum"].(float64); ok { return int(f) }; return 0 }(),
                                            "if_down": func() int { if f, ok := _sgm["if-down"].(float64); ok { return int(f) }; return 0 }(),
                                            "no_neighbor": func() int { if f, ok := _sgm["no-neighbor"].(float64); ok { return int(f) }; return 0 }(),
                                            "no_virtual_neighbor": func() int { if f, ok := _sgm["no-virtual-neighbor"].(float64); ok { return int(f) }; return 0 }(),
                                            "non_local": func() int { if f, ok := _sgm["non-local"].(float64); ok { return int(f) }; return 0 }(),
                                            "packet_type": func() int { if f, ok := _sgm["packet-type"].(float64); ok { return int(f) }; return 0 }(),
                                            "passive_interface": func() int { if f, ok := _sgm["passive-interface"].(float64); ok { return int(f) }; return 0 }(),
                                            "size": func() int { if f, ok := _sgm["size"].(float64); ok { return int(f) }; return 0 }(),
                                            "tx": func() int { if f, ok := _sgm["tx"].(float64); ok { return int(f) }; return 0 }(),
                                            "version": func() int { if f, ok := _sgm["version"].(float64); ok { return int(f) }; return 0 }(),
                                            "virtual_link": func() int { if f, ok := _sgm["virtual-link"].(float64); ok { return int(f) }; return 0 }(),
                                            "zero_rid": func() int { if f, ok := _sgm["zero-rid"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                    }
                }
            }
            d.Set("instance_errors", mapped)
        }
    } else {
        d.Set("instance_errors", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ospf-errors-" + acctest.RandString(10)))
    return nil
}

