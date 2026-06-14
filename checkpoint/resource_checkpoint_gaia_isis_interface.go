package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaIsisInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIsisInterface,
        Read:   readGaiaIsisInterface,
        Update: updateGaiaIsisInterface,
        Delete: deleteGaiaIsisInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the interface`,
            },
            "address_family": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Address family that the interface will run on`,
            },
            "advertise": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Advertise this interfaces IP address`,
            },
            "circuit_type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Set level for the interface to run on`,
            },
            "csnp_interval": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure IS-IS Interface Csnp Interval`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the csnp interval configuration`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the level for the csnp configuration`,
                        },
                    },
                },
            },
            "hello": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure ISIS interface hello`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "padding": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set hello padding for interface`,
                        },
                        "timers": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Set level 1 configuration`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "holdtime": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Set holdtime`,
                                    },
                                    "interval": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Set interval`,
                                    },
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The IS-IS level that this entry belongs to`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "ip_reachability": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configure bidirectional forwarding detection (BFD) for interface`,
            },
            "ipv6": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure IS-IS ipv6 options. Note that ipv6 multi topology must be enabled`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "advertise": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Advertise this interfaces IP address`,
                        },
                        "ip_reachability": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Configure bidirectional forwarding detection (BFD) for interface`,
                        },
                        "metric": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Set the metric (cost) of this interface`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "metric": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Set the interface metric interval configuration`,
                                    },
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Set the level for this metric configuration`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "lsp_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configure delay between sending LSPs`,
            },
            "mesh_group": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configure this interface as a member of a mesh group`,
            },
            "metric": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Set the metric (cost) of this interface`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "metric": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the interface metric interval configuration`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the level for this metric configuration`,
                        },
                    },
                },
            },
            "passive_mode": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable or disable passive operation`,
            },
            "point_to_point": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure point to point options`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "toggle": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Configure toggle`,
                        },
                        "retransmit_interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure retransmit interval`,
                        },
                        "retransmit_throttle_interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure retransmit Throttle interval`,
                        },
                    },
                },
            },
            "priority": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Set DIS priority`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "value": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the level 1 interface priority interval configuration`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set the level 2 csnp configuration`,
                        },
                    },
                },
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The instance to be queried`,
            },
            "name": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `The interface name of the interface to be queried`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaIsisInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("address_family"); ok {
        payload["address-family"] = v.(string)
    }

    if v, ok := d.GetOkExists("advertise"); ok {
        payload["advertise"] = v.(bool)
    }

    if v, ok := d.GetOk("circuit_type"); ok {
        payload["circuit-type"] = v.(string)
    }

    if v := d.Get("csnp_interval"); len(v.([]interface{})) > 0 {
        csnpintervalList := v.([]interface{})
        csnpintervalArray := make([]interface{}, 0, len(csnpintervalList))
        for i := range csnpintervalList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("csnp_interval.%d.interval", i)); ok {
                itemMap["interval"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("csnp_interval.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                csnpintervalArray = append(csnpintervalArray, itemMap)
            }
        }
        if len(csnpintervalArray) > 0 {
            payload["csnp-interval"] = csnpintervalArray
        }
    }

    if v := d.Get("hello"); len(v.([]interface{})) > 0 {
        _ = v
        helloMap := make(map[string]interface{})
        if v, ok := d.GetOk("hello.0.padding"); ok {
            helloMap["padding"] = v.(string)
        }
        if v, ok := d.GetOk("hello.0.timers"); ok {
            timersList := v.([]interface{})
            timersArray := make([]interface{}, 0, len(timersList))
            for i := range timersList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.holdtime", i)); ok {
                    itemMap["holdtime"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.interval", i)); ok {
                    itemMap["interval"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if len(itemMap) > 0 {
                    timersArray = append(timersArray, itemMap)
                }
            }
            if len(timersArray) > 0 {
                helloMap["timers"] = timersArray
            }
        }
        if len(helloMap) > 0 {
            payload["hello"] = helloMap
        }
    }

    if v, ok := d.GetOkExists("ip_reachability"); ok {
        payload["ip-reachability"] = v.(bool)
    }

    if v := d.Get("ipv6"); len(v.([]interface{})) > 0 {
        _ = v
        ipv6Map := make(map[string]interface{})
        if v, ok := d.GetOkExists("ipv6.0.advertise"); ok && v.(bool) {
            ipv6Map["advertise"] = v.(bool)
        }
        if v, ok := d.GetOkExists("ipv6.0.ip_reachability"); ok && v.(bool) {
            ipv6Map["ip-reachability"] = v.(bool)
        }
        if v, ok := d.GetOk("ipv6.0.metric"); ok {
            metricList := v.([]interface{})
            metricArray := make([]interface{}, 0, len(metricList))
            for i := range metricList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.metric.%d.metric", i)); ok {
                    itemMap["metric"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.metric.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if len(itemMap) > 0 {
                    metricArray = append(metricArray, itemMap)
                }
            }
            if len(metricArray) > 0 {
                ipv6Map["metric"] = metricArray
            }
        }
        if len(ipv6Map) > 0 {
            payload["ipv6"] = ipv6Map
        }
    }

    if v, ok := d.GetOk("lsp_interval"); ok {
        payload["lsp-interval"] = v.(string)
    }

    if v, ok := d.GetOk("mesh_group"); ok {
        payload["mesh-group"] = v.(string)
    }

    if v := d.Get("metric"); len(v.([]interface{})) > 0 {
        metricList := v.([]interface{})
        metricArray := make([]interface{}, 0, len(metricList))
        for i := range metricList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("metric.%d.metric", i)); ok {
                itemMap["metric"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("metric.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                metricArray = append(metricArray, itemMap)
            }
        }
        if len(metricArray) > 0 {
            payload["metric"] = metricArray
        }
    }

    if v, ok := d.GetOkExists("passive_mode"); ok {
        payload["passive-mode"] = v.(bool)
    }

    if v := d.Get("point_to_point"); len(v.([]interface{})) > 0 {
        _ = v
        pointtopointMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("point_to_point.0.toggle"); ok && v.(bool) {
            pointtopointMap["toggle"] = v.(bool)
        }
        if v, ok := d.GetOk("point_to_point.0.retransmit_interval"); ok {
            pointtopointMap["retransmit-interval"] = v.(string)
        }
        if v, ok := d.GetOk("point_to_point.0.retransmit_throttle_interval"); ok {
            pointtopointMap["retransmit-throttle-interval"] = v.(string)
        }
        if len(pointtopointMap) > 0 {
            payload["point-to-point"] = pointtopointMap
        }
    }

    if v := d.Get("priority"); len(v.([]interface{})) > 0 {
        priorityList := v.([]interface{})
        priorityArray := make([]interface{}, 0, len(priorityList))
        for i := range priorityList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("priority.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("priority.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                priorityArray = append(priorityArray, itemMap)
            }
        }
        if len(priorityArray) > 0 {
            payload["priority"] = priorityArray
        }
    }

    log.Println("Create IsisInterface - Map = ", payload)

    addIsisInterfaceRes, err := client.ApiCallSimple("add-isis-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIsisInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIsisInterfaceRes.Success {
            errMsg = addIsisInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIsisInterfaceRes.GetData()
        }

        debugLogOperation(
            "isis-interface",        // resource type
            "create",                       // operation
            "add-isis-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add isis-interface: %v", err)
    }
    if !addIsisInterfaceRes.Success {
        if addIsisInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addIsisInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("isis-interface-" + acctest.RandString(10)))
    return readGaiaIsisInterface(d, m)
}

func readGaiaIsisInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

   payload["name"] = d.Get("name")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    // show-isis-interfaces (plural) does not accept a "name" parameter.
    // build_payload may have added it from the field_set; remove it unconditionally.
    delete(payload, "name")

    showIsisInterfaceRes, err := client.ApiCallSimple("show-isis-interfaces", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showIsisInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showIsisInterfaceRes.Success {
            errMsg = showIsisInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showIsisInterfaceRes.GetData()
        }

        debugLogOperation(
            "isis-interface",        // resource type
            "read",                       // operation
            "show-isis-interfaces",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show isis-interface: %v", err)
    }
    if !showIsisInterfaceRes.Success {
        if data := showIsisInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showIsisInterfaceRes.ErrorMsg)
    }

    isisInterface := showIsisInterfaceRes.GetData()

    log.Println("Read IsisInterface - Show JSON = ", isisInterface)

    if v, exists := isisInterface["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    // show-isis-interfaces returns objects list; find entry matching "interface" attr.
    if objects, ok := isisInterface["objects"].([]interface{}); ok {
        ifaceVal := ""
        if v, ok := d.GetOk("interface"); ok {
            ifaceVal = v.(string)
        }
        for _, obj := range objects {
            if item, ok := obj.(map[string]interface{}); ok {
                if fmt.Sprintf("%v", item["name"]) == ifaceVal {
                    if v, ok := item["address-family"]; ok {
                        d.Set("address_family", fmt.Sprintf("%v", v))
                    }
                    if v, ok := item["advertise"]; ok {
                        if b, ok := v.(bool); ok { d.Set("advertise", b) }
                    }
                    if v, ok := item["circuit-type"]; ok {
                        d.Set("circuit_type", fmt.Sprintf("%v", v))
                    }
                    if cv, ok := item["csnp-interval"].([]interface{}); ok {
                        csnpMapped := make([]interface{}, 0, len(cv))
                        for _, e := range cv {
                            if em, ok := e.(map[string]interface{}); ok {
                                csnpMapped = append(csnpMapped, map[string]interface{}{
                                    "interval": fmt.Sprintf("%v", em["interval"]),
                                    "level":    fmt.Sprintf("%v", em["level"]),
                                })
                            }
                        }
                        d.Set("csnp_interval", csnpMapped)
                    }
                    if hv, ok := item["hello"].(map[string]interface{}); ok {
                        hEntry := map[string]interface{}{
                            "padding": fmt.Sprintf("%v", hv["padding"]),
                        }
                        if timers, ok := hv["timers"].([]interface{}); ok {
                            tMapped := make([]interface{}, 0, len(timers))
                            for _, t := range timers {
                                if tm, ok := t.(map[string]interface{}); ok {
                                    tMapped = append(tMapped, map[string]interface{}{
                                        "holdtime": fmt.Sprintf("%v", tm["holdtime"]),
                                        "interval": fmt.Sprintf("%v", tm["interval"]),
                                        "level":    fmt.Sprintf("%v", tm["level"]),
                                    })
                                }
                            }
                            hEntry["timers"] = tMapped
                        }
                        d.Set("hello", []interface{}{hEntry})
                    }
                    if v, ok := item["ip-reachability"]; ok {
                        if b, ok := v.(bool); ok { d.Set("ip_reachability", b) }
                    }
                    if iv, ok := item["ipv6"].(map[string]interface{}); ok {
                        ipv6Entry := map[string]interface{}{}
                        if b, ok := iv["advertise"].(bool); ok { ipv6Entry["advertise"] = b }
                        if b, ok := iv["ip-reachability"].(bool); ok { ipv6Entry["ip_reachability"] = b }
                        if metrics, ok := iv["metric"].([]interface{}); ok {
                            mMapped := make([]interface{}, 0, len(metrics))
                            for _, me := range metrics {
                                if mm, ok := me.(map[string]interface{}); ok {
                                    mMapped = append(mMapped, map[string]interface{}{
                                        "metric": fmt.Sprintf("%v", mm["metric"]),
                                        "level":  fmt.Sprintf("%v", mm["level"]),
                                    })
                                }
                            }
                            ipv6Entry["metric"] = mMapped
                        }
                        d.Set("ipv6", []interface{}{ipv6Entry})
                    }
                    if v, ok := item["lsp-interval"]; ok {
                        d.Set("lsp_interval", fmt.Sprintf("%v", v))
                    }
                    if v, ok := item["mesh-group"]; ok {
                        d.Set("mesh_group", fmt.Sprintf("%v", v))
                    }
                    if mv, ok := item["metric"].([]interface{}); ok {
                        metricMapped := make([]interface{}, 0, len(mv))
                        for _, e := range mv {
                            if em, ok := e.(map[string]interface{}); ok {
                                metricMapped = append(metricMapped, map[string]interface{}{
                                    "metric": fmt.Sprintf("%v", em["metric"]),
                                    "level":  fmt.Sprintf("%v", em["level"]),
                                })
                            }
                        }
                        d.Set("metric", metricMapped)
                    }
                    if v, ok := item["passive-mode"]; ok {
                        if b, ok := v.(bool); ok { d.Set("passive_mode", b) }
                    }
                    if pv, ok := item["point-to-point"].(map[string]interface{}); ok {
                        d.Set("point_to_point", []interface{}{map[string]interface{}{
                            "toggle":                       func() bool { b, _ := pv["toggle"].(bool); return b }(),
                            "retransmit_interval":          fmt.Sprintf("%v", pv["retransmit-interval"]),
                            "retransmit_throttle_interval": fmt.Sprintf("%v", pv["retransmit-throttle-interval"]),
                        }})
                    }
                    if pv, ok := item["priority"].([]interface{}); ok {
                        priorityMapped := make([]interface{}, 0, len(pv))
                        for _, e := range pv {
                            if em, ok := e.(map[string]interface{}); ok {
                                priorityMapped = append(priorityMapped, map[string]interface{}{
                                    "value": fmt.Sprintf("%v", em["value"]),
                                    "level": fmt.Sprintf("%v", em["level"]),
                                })
                            }
                        }
                        d.Set("priority", priorityMapped)
                    }
                    d.Set("interface", ifaceVal)
                    break
                }
            }
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaIsisInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("address_family"); ok {
        payload["address-family"] = v.(string)
    }

    if v, ok := d.GetOkExists("advertise"); ok {
        payload["advertise"] = v.(bool)
    }

    if v, ok := d.GetOk("circuit_type"); ok {
        payload["circuit-type"] = v.(string)
    }

    if v := d.Get("csnp_interval"); len(v.([]interface{})) > 0 {
        csnpintervalList := v.([]interface{})
        csnpintervalArray := make([]interface{}, 0, len(csnpintervalList))
        for i := range csnpintervalList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("csnp_interval.%d.interval", i)); ok {
                itemMap["interval"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("csnp_interval.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                csnpintervalArray = append(csnpintervalArray, itemMap)
            }
        }
        if len(csnpintervalArray) > 0 {
            payload["csnp-interval"] = csnpintervalArray
        }
    }

    if v := d.Get("hello"); len(v.([]interface{})) > 0 {
        _ = v
        helloMap := make(map[string]interface{})
        if v, ok := d.GetOk("hello.0.padding"); ok {
            helloMap["padding"] = v.(string)
        }
        if v, ok := d.GetOk("hello.0.timers"); ok {
            timersList := v.([]interface{})
            timersArray := make([]interface{}, 0, len(timersList))
            for i := range timersList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.holdtime", i)); ok {
                    itemMap["holdtime"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.interval", i)); ok {
                    itemMap["interval"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("hello.0.timers.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if len(itemMap) > 0 {
                    timersArray = append(timersArray, itemMap)
                }
            }
            if len(timersArray) > 0 {
                helloMap["timers"] = timersArray
            }
        }
        if len(helloMap) > 0 {
            payload["hello"] = helloMap
        }
    }

    if v, ok := d.GetOkExists("ip_reachability"); ok {
        payload["ip-reachability"] = v.(bool)
    }

    if v := d.Get("ipv6"); len(v.([]interface{})) > 0 {
        _ = v
        ipv6Map := make(map[string]interface{})
        if v, ok := d.GetOkExists("ipv6.0.advertise"); ok && v.(bool) {
            ipv6Map["advertise"] = v.(bool)
        }
        if v, ok := d.GetOkExists("ipv6.0.ip_reachability"); ok && v.(bool) {
            ipv6Map["ip-reachability"] = v.(bool)
        }
        if v, ok := d.GetOk("ipv6.0.metric"); ok {
            metricList := v.([]interface{})
            metricArray := make([]interface{}, 0, len(metricList))
            for i := range metricList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.metric.%d.metric", i)); ok {
                    itemMap["metric"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.metric.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if len(itemMap) > 0 {
                    metricArray = append(metricArray, itemMap)
                }
            }
            if len(metricArray) > 0 {
                ipv6Map["metric"] = metricArray
            }
        }
        if len(ipv6Map) > 0 {
            payload["ipv6"] = ipv6Map
        }
    }

    if v, ok := d.GetOk("lsp_interval"); ok {
        payload["lsp-interval"] = v.(string)
    }

    if v, ok := d.GetOk("mesh_group"); ok {
        payload["mesh-group"] = v.(string)
    }

    if v := d.Get("metric"); len(v.([]interface{})) > 0 {
        metricList := v.([]interface{})
        metricArray := make([]interface{}, 0, len(metricList))
        for i := range metricList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("metric.%d.metric", i)); ok {
                itemMap["metric"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("metric.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                metricArray = append(metricArray, itemMap)
            }
        }
        if len(metricArray) > 0 {
            payload["metric"] = metricArray
        }
    }

    if v, ok := d.GetOkExists("passive_mode"); ok {
        payload["passive-mode"] = v.(bool)
    }

    if v := d.Get("point_to_point"); len(v.([]interface{})) > 0 {
        _ = v
        pointtopointMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("point_to_point.0.toggle"); ok && v.(bool) {
            pointtopointMap["toggle"] = v.(bool)
        }
        if v, ok := d.GetOk("point_to_point.0.retransmit_interval"); ok {
            pointtopointMap["retransmit-interval"] = v.(string)
        }
        if v, ok := d.GetOk("point_to_point.0.retransmit_throttle_interval"); ok {
            pointtopointMap["retransmit-throttle-interval"] = v.(string)
        }
        if len(pointtopointMap) > 0 {
            payload["point-to-point"] = pointtopointMap
        }
    }

    if v := d.Get("priority"); len(v.([]interface{})) > 0 {
        priorityList := v.([]interface{})
        priorityArray := make([]interface{}, 0, len(priorityList))
        for i := range priorityList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("priority.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("priority.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                priorityArray = append(priorityArray, itemMap)
            }
        }
        if len(priorityArray) > 0 {
            payload["priority"] = priorityArray
        }
    }

    setIsisInterfaceRes, err := client.ApiCallSimple("set-isis-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setIsisInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setIsisInterfaceRes.Success {
            errMsg = setIsisInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setIsisInterfaceRes.GetData()
        }

        debugLogOperation(
            "isis-interface",        // resource type
            "update",                       // operation
            "set-isis-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set isis-interface: %v", err)
    }
    if !setIsisInterfaceRes.Success {
        return fmt.Errorf(setIsisInterfaceRes.ErrorMsg)
    }

    return readGaiaIsisInterface(d, m)
}

func deleteGaiaIsisInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    deleteIsisInterfaceRes, err := client.ApiCallSimple("delete-isis-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteIsisInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteIsisInterfaceRes.Success {
            errMsg = deleteIsisInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteIsisInterfaceRes.GetData()
        }

        debugLogOperation(
            "isis-interface",        // resource type
            "delete",                       // operation
            "delete-isis-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete isis-interface: %v", err)
    }
    if !deleteIsisInterfaceRes.Success {
        return fmt.Errorf(deleteIsisInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

