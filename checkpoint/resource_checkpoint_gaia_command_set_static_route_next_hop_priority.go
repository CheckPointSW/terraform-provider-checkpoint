package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetStaticRouteNextHopPriority() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetStaticRouteNextHopPriority,
        Read:   readGaiaSetStaticRouteNextHopPriority,
        Delete: deleteGaiaSetStaticRouteNextHopPriority,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "next_hop_gateway": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `nexthop gateway, can be IP address or interface name`,
            },
            "priority": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `Priority defines which gateway to select as the next hop, the lower the priority, the higher the preference. can be default or integer from 1 to 8`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "next_hop": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "gateway": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "priority": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "ping": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "rank": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "scope_local": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "comment": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSetStaticRouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("next_hop_gateway"); ok {
        payload["next-hop-gateway"] = v.(string)
    }

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute set-static-route-next-hop-priority - Payload = ", payload)

    GaiaSetStaticRouteNextHopPriorityRes, err := client.ApiCallSimple("set-static-route-next-hop-priority", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetStaticRouteNextHopPriorityRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetStaticRouteNextHopPriorityRes.Success {
            errMsg = GaiaSetStaticRouteNextHopPriorityRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetStaticRouteNextHopPriorityRes.GetData()
        }

        debugLogOperation(
            "set-static-route-next-hop-priority",        // resource type
            "command",                       // operation
            "set-static-route-next-hop-priority",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-static-route-next-hop-priority: %v", err)
    }
    if !GaiaSetStaticRouteNextHopPriorityRes.Success {
        if GaiaSetStaticRouteNextHopPriorityRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetStaticRouteNextHopPriorityRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaSetStaticRouteNextHopPriorityRes.GetData()
    if v, exists := _respData["type"]; exists {
        d.Set("type", toString(v))
    }
    if v, exists := _respData["next-hop"]; exists {
        if hops, ok := v.([]interface{}); ok {
            var hopList []interface{}
            for _, h := range hops {
                if hm, ok := h.(map[string]interface{}); ok {
                    hopList = append(hopList, map[string]interface{}{
                        "gateway":  fmt.Sprintf("%v", hm["gateway"]),
                        "priority": fmt.Sprintf("%v", hm["priority"]),
                    })
                }
            }
            d.Set("next_hop", hopList)
        }
    }
    if v, exists := _respData["ping"]; exists {
        if bv, ok := v.(bool); ok {
            d.Set("ping", bv)
        }
    }
    if v, exists := _respData["rank"]; exists {
        d.Set("rank", toString(v))
    }
    if v, exists := _respData["scope-local"]; exists {
        if bv, ok := v.(bool); ok {
            d.Set("scope_local", bv)
        }
    }
    if v, exists := _respData["comment"]; exists {
        d.Set("comment", toString(v))
    }
    if v, exists := _respData["member-id"]; exists {
        d.Set("member_id", toString(v))
    }


    d.SetId(fmt.Sprintf("set-static-route-next-hop-priority-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetStaticRouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }
    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }
    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Read set-static-route-next-hop-priority - show-static-route payload = ", payload)

    showRes, err := client.ApiCallSimple("show-static-route", payload)
    if err != nil {
        return fmt.Errorf("Failed to show static-route: %v", err)
    }
    if !showRes.Success {
        d.SetId("")
        return nil
    }

    routeData := showRes.GetData()

    wantGW := ""
    if v, ok := d.GetOk("next_hop_gateway"); ok {
        wantGW = v.(string)
    }
    wantPri := 0
    if v, ok := d.GetOk("priority"); ok {
        wantPri = v.(int)
    }

    found := false
    if nhRaw, ok := routeData["next-hop"].([]interface{}); ok {
        for _, item := range nhRaw {
            if nh, ok := item.(map[string]interface{}); ok {
                gw := fmt.Sprintf("%v", nh["gateway"])
                pri := fmt.Sprintf("%v", nh["priority"])
                if gw == wantGW && pri == fmt.Sprintf("%d", wantPri) {
                    found = true
                    break
                }
            }
        }
        if found {
            d.Set("next_hop", func() []interface{} {
                out := make([]interface{}, len(nhRaw))
                for i, item := range nhRaw {
                    if nh, ok := item.(map[string]interface{}); ok {
                        out[i] = map[string]interface{}{
                            "gateway":  fmt.Sprintf("%v", nh["gateway"]),
                            "priority": fmt.Sprintf("%v", nh["priority"]),
                        }
                    }
                }
                return out
            }())
        }
    }

    if v, ok := routeData["type"]; ok {
        d.Set("type", toString(v))
    }
    if v, ok := routeData["ping"]; ok {
        if bv, bOk := v.(bool); bOk {
            d.Set("ping", bv)
        }
    }
    if v, ok := routeData["scope-local"]; ok {
        if bv, bOk := v.(bool); bOk {
            d.Set("scope_local", bv)
        }
    }
    if v, ok := routeData["rank"]; ok {
        d.Set("rank", toString(v))
    }
    if v, ok := routeData["comment"]; ok {
        d.Set("comment", toString(v))
    }
    if v, ok := routeData["member-id"]; ok {
        d.Set("member_id", toString(v))
    }

    if !found {
        d.SetId("")
        return nil
    }

    return nil
}

func deleteGaiaSetStaticRouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

