package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfSummary() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfSummary,
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
            "spf_schedule_delay": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "spf_hold_time": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "num_areas": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "num_normal_areas": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "num_stub_areas": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "num_nssa_areas": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "rfc1583compatibility": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "abr": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "asbr": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "graceful_restart_capable": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "grace_period": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "graceful_restart_status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "last_graceful_restart_exit_status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "virtual_links": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "interface_up": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "interface_down": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ase_cost": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ase_type": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "areas": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "area": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interfaces": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "default_route_cost": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "default_metric_type": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "num_abrs": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "num_asbrs": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "spf_runs": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "ranges": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "range": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "status": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "stubnets": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "stubnet": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "cost": {
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

func readGaiaShowOspfSummary(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-summary - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-summary", payload)
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
            "ospf-summary",        // resource type
            "read",                       // operation
            "show-ospf-summary",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-summary: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["spf-schedule-delay"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("spf_schedule_delay", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["spf-hold-time"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("spf_hold_time", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["num-areas"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("num_areas", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["num-normal-areas"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("num_normal_areas", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["num-stub-areas"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("num_stub_areas", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["num-nssa-areas"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("num_nssa_areas", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["rfc1583compatibility"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("rfc1583compatibility", b)
        } else if s, ok := v.(string); ok {
            d.Set("rfc1583compatibility", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["abr"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("abr", b)
        } else if s, ok := v.(string); ok {
            d.Set("abr", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["asbr"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("asbr", b)
        } else if s, ok := v.(string); ok {
            d.Set("asbr", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["graceful-restart-capable"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("graceful_restart_capable", b)
        } else if s, ok := v.(string); ok {
            d.Set("graceful_restart_capable", s == "true")
        }
    }
    if v, exists := commandRes.GetData()["grace-period"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("grace_period", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["graceful-restart-status"]; exists {
        d.Set("graceful_restart_status", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["last-graceful-restart-exit-status"]; exists {
        d.Set("last_graceful_restart_exit_status", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-links"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("virtual_links", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["interface-up"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("interface_up", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["interface-down"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("interface_down", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["ase-cost"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("ase_cost", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["ase-type"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("ase_type", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["areas"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "area": func() string { if _v, _ok := m["area"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "interfaces": func() int { if f, ok := m["interfaces"].(float64); ok { return int(f) }; return 0 }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "default_route_cost": func() int { if f, ok := m["default-route-cost"].(float64); ok { return int(f) }; return 0 }(),
                        "default_metric_type": func() int { if f, ok := m["default-metric-type"].(float64); ok { return int(f) }; return 0 }(),
                        "num_abrs": func() int { if f, ok := m["num-abrs"].(float64); ok { return int(f) }; return 0 }(),
                        "num_asbrs": func() int { if f, ok := m["num-asbrs"].(float64); ok { return int(f) }; return 0 }(),
                        "spf_runs": func() int { if f, ok := m["spf-runs"].(float64); ok { return int(f) }; return 0 }(),
                        "ranges": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["ranges"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "range": func() string { if _v, _ok := _sgm["range"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "status": func() string { if _v, _ok := _sgm["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "type": func() string { if _v, _ok := _sgm["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "stubnets": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["stubnets"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "stubnet": func() string { if _v, _ok := _sgm["stubnet"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "cost": func() int { if f, ok := _sgm["cost"].(float64); ok { return int(f) }; return 0 }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                    }
                }
            }
            d.Set("areas", mapped)
        }
    } else {
        d.Set("areas", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ospf-summary-" + acctest.RandString(10)))
    return nil
}

