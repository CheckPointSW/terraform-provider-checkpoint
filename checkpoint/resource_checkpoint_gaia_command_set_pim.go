package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetPim() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetPim,
        Read:   readGaiaSetPim,
        Delete: deleteGaiaSetPim,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "assert_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies the number of seconds that assert state should be maintained in the absence of a refreshing assert message.`,
            },
            "assert_rank": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Assert rank defines the cost of a routing protocolrelative to other protocols.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "protocol": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure the assert rank of a protocol.`,
                        },
                        "rank": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The cost metric.`,
                        },
                    },
                },
            },
            "bootstrap_candidate": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures candidate Bootstrap Router (candidate BSR) options.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "local_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configures the local address to use for this candidate Bootstrap Router (candidate BSR).`,
                        },
                        "priority": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configures the candidate Bootstrap Router (candidate BSR) priority.`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configures this router to be a candidate Bootstrap Router (candidate BSR).`,
                        },
                    },
                },
            },
            "candidate_rp": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures candidate Rendezvous Point (candidate RP) options.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "advertise_interval": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure the Advertisement Interval`,
                        },
                        "local_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure the Candidate RP Local Address.`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure a Candidate Rendezvous Point`,
                        },
                        "priority": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure Candidate Rendezvous Point Priority.`,
                        },
                        "multicast_group": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure a Candidate RP Multicast Group.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The multicast group prefix/mask, in CIDR notation.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "data_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure the Data Interval.`,
            },
            "hello_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure the Hello Interval.`,
            },
            "jp_delay_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure the Random Delay Join/Prune Interval`,
            },
            "jp_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure the Join/Prune Interval (PIM-SM/SSM only).`,
            },
            "mode": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure Dense Mode/Sparse Mode/SSM Mode`,
            },
            "register_suppress_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure Register-Suppression Interval`,
            },
            "enable_state_refresh": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure State Refresh`,
            },
            "state_refresh_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure State Refresh Interval`,
            },
            "state_refresh_ttl": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure State Refresh TTL`,
            },
            "static_rp": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure Static Rendezvous Point`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rp_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Adds the given static Rendezvous Point (static RP).`,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure a static Rendezvous Point (static RP).`,
                        },
                        "multicast_group": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure a Static RP Multicast Group.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The multicast group prefix/mask, in CIDR notation.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "spt_threshold": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure SPT Threshold`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "multicast_group": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The multicast group prefix/mask, in CIDR notation.`,
                        },
                        "threshold": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The threshold for the multicast group.`,
                        },
                    },
                },
            },
            "custom_ssm_prefix": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure Custom SSM Prefix`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The multicast group prefix/mask, in CIDR notation.`,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaSetPim(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("assert_interval"); ok {
        payload["assert-interval"] = v.(string)
    }

    if v := d.Get("assert_rank"); len(v.([]interface{})) > 0 {
        _ = v
        assertrankMap := make(map[string]interface{})
        if v, ok := d.GetOk("assert_rank.0.protocol"); ok {
            assertrankMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("assert_rank.0.rank"); ok {
            assertrankMap["rank"] = v.(string)
        }
        if len(assertrankMap) > 0 {
            payload["assert-rank"] = assertrankMap
        }
    }

    if v := d.Get("bootstrap_candidate"); len(v.([]interface{})) > 0 {
        _ = v
        bootstrapcandidateMap := make(map[string]interface{})
        if v, ok := d.GetOk("bootstrap_candidate.0.local_address"); ok {
            bootstrapcandidateMap["local-address"] = v.(string)
        }
        if v, ok := d.GetOk("bootstrap_candidate.0.priority"); ok {
            bootstrapcandidateMap["priority"] = v.(string)
        }
        if v, ok := d.GetOkExists("bootstrap_candidate.0.enable"); ok && v.(bool) {
            bootstrapcandidateMap["enable"] = v.(bool)
        }
        if len(bootstrapcandidateMap) > 0 {
            payload["bootstrap-candidate"] = bootstrapcandidateMap
        }
    }

    if v := d.Get("candidate_rp"); len(v.([]interface{})) > 0 {
        _ = v
        candidaterpMap := make(map[string]interface{})
        if v, ok := d.GetOk("candidate_rp.0.advertise_interval"); ok {
            candidaterpMap["advertise-interval"] = v.(int)
        }
        if v, ok := d.GetOk("candidate_rp.0.local_address"); ok {
            candidaterpMap["local-address"] = v.(string)
        }
        if v, ok := d.GetOkExists("candidate_rp.0.enable"); ok && v.(bool) {
            candidaterpMap["enable"] = v.(bool)
        }
        if v, ok := d.GetOk("candidate_rp.0.priority"); ok {
            candidaterpMap["priority"] = v.(int)
        }
        if v, ok := d.GetOk("candidate_rp.0.multicast_group"); ok {
            _ = v
            multicastgroupMap := make(map[string]interface{})
            if v, ok := d.GetOk("candidate_rp.0.multicast_group.0.address"); ok {
                multicastgroupMap["address"] = v.(string)
            }
            if len(multicastgroupMap) > 0 {
                candidaterpMap["multicast-group"] = multicastgroupMap
            }
        }
        if len(candidaterpMap) > 0 {
            payload["candidate-rp"] = candidaterpMap
        }
    }

    if v, ok := d.GetOk("data_interval"); ok {
        payload["data-interval"] = v.(string)
    }

    if v, ok := d.GetOk("hello_interval"); ok {
        payload["hello-interval"] = v.(string)
    }

    if v, ok := d.GetOk("jp_delay_interval"); ok {
        payload["jp-delay-interval"] = v.(string)
    }

    if v, ok := d.GetOk("jp_interval"); ok {
        payload["jp-interval"] = v.(string)
    }

    if v, ok := d.GetOk("mode"); ok {
        payload["mode"] = v.(string)
    }

    if v, ok := d.GetOk("register_suppress_interval"); ok {
        payload["register-suppress-interval"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_state_refresh"); ok {
        payload["enable-state-refresh"] = v.(bool)
    }

    if v, ok := d.GetOk("state_refresh_interval"); ok {
        payload["state-refresh-interval"] = v.(string)
    }

    if v, ok := d.GetOk("state_refresh_ttl"); ok {
        payload["state-refresh-ttl"] = v.(string)
    }

    if v := d.Get("static_rp"); len(v.([]interface{})) > 0 {
        _ = v
        staticrpMap := make(map[string]interface{})
        if v, ok := d.GetOk("static_rp.0.rp_address"); ok {
            staticrpMap["rp-address"] = v.(string)
        }
        if v, ok := d.GetOkExists("static_rp.0.enable"); ok && v.(bool) {
            staticrpMap["enable"] = v.(bool)
        }
        if v, ok := d.GetOk("static_rp.0.multicast_group"); ok {
            multicastgroupList := v.([]interface{})
            multicastgroupArray := make([]interface{}, 0, len(multicastgroupList))
            for i := range multicastgroupList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("static_rp.0.multicast_group.%d.address", i)); ok {
                    itemMap["address"] = v.(string)
                }
                if len(itemMap) > 0 {
                    multicastgroupArray = append(multicastgroupArray, itemMap)
                }
            }
            if len(multicastgroupArray) > 0 {
                staticrpMap["multicast-group"] = multicastgroupArray
            }
        }
        if len(staticrpMap) > 0 {
            payload["static-rp"] = staticrpMap
        }
    }

    if v := d.Get("spt_threshold"); len(v.([]interface{})) > 0 {
        sptthresholdList := v.([]interface{})
        sptthresholdArray := make([]interface{}, 0, len(sptthresholdList))
        for i := range sptthresholdList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("spt_threshold.%d.multicast_group", i)); ok {
                itemMap["multicast-group"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("spt_threshold.%d.threshold", i)); ok {
                itemMap["threshold"] = v.(string)
            }
            if len(itemMap) > 0 {
                sptthresholdArray = append(sptthresholdArray, itemMap)
            }
        }
        if len(sptthresholdArray) > 0 {
            payload["spt-threshold"] = sptthresholdArray
        }
    }

    if v := d.Get("custom_ssm_prefix"); len(v.([]interface{})) > 0 {
        customssmprefixList := v.([]interface{})
        customssmprefixArray := make([]interface{}, 0, len(customssmprefixList))
        for i := range customssmprefixList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("custom_ssm_prefix.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if len(itemMap) > 0 {
                customssmprefixArray = append(customssmprefixArray, itemMap)
            }
        }
        if len(customssmprefixArray) > 0 {
            payload["custom-ssm-prefix"] = customssmprefixArray
        }
    }

    log.Println("Execute set-pim - Payload = ", payload)

    GaiaSetPimRes, err := client.ApiCallSimple("set-pim", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetPimRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetPimRes.Success {
            errMsg = GaiaSetPimRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetPimRes.GetData()
        }

        debugLogOperation(
            "set-pim",        // resource type
            "command",                       // operation
            "set-pim",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-pim: %v", err)
    }
    if !GaiaSetPimRes.Success {
        if GaiaSetPimRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetPimRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }



    d.SetId(fmt.Sprintf("set-pim-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetPim(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetPim(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

