package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"
)
func dataGaiaSetStaticMrouteNextHopPriority() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetStaticMrouteNextHopPriority,
        Read:   readGaiaSetStaticMrouteNextHopPriority,
        Delete: deleteGaiaSetStaticMrouteNextHopPriority,
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
                Description: `Address of the static-mroute to set configuration for.`,
            },
            "mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `Mask length of the static-mroute.`,
            },
            "next_hop_gateway": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Next-hop gateway, must be an IP address.`,
            },
            "priority": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `Priority defines which gateway to select as the next hop: the lower the priority, the higher the preference.<br>Can be default or integer from 1 to 8`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
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
        },
    }
}

func createGaiaSetStaticMrouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Execute set-static-mroute-next-hop-priority - Payload = ", payload)

    GaiaSetStaticMrouteNextHopPriorityRes, err := client.ApiCallSimple("set-static-mroute-next-hop-priority", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetStaticMrouteNextHopPriorityRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetStaticMrouteNextHopPriorityRes.Success {
            errMsg = GaiaSetStaticMrouteNextHopPriorityRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetStaticMrouteNextHopPriorityRes.GetData()
        }

        debugLogOperation(
            "set-static-mroute-next-hop-priority",        // resource type
            "command",                       // operation
            "set-static-mroute-next-hop-priority",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-static-mroute-next-hop-priority: %v", err)
    }
    if !GaiaSetStaticMrouteNextHopPriorityRes.Success {
        if GaiaSetStaticMrouteNextHopPriorityRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetStaticMrouteNextHopPriorityRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaSetStaticMrouteNextHopPriorityRes.GetData()
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


    d.SetId(fmt.Sprintf("set-static-mroute-next-hop-priority-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetStaticMrouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Read set-static-mroute-next-hop-priority - show-static-mroute payload = ", payload)

    showRes, err := client.ApiCallSimple("show-static-mroute", payload)
    if err != nil {
        return fmt.Errorf("Failed to show static-mroute: %v", err)
    }
    if !showRes.Success {
        data := showRes.GetData()
        if code, ok := data["code"]; ok {
            if strings.Contains(strings.ToLower(fmt.Sprintf("%v", code)), "not_found") {
                d.SetId("")
                return nil
            }
        }
        if showRes.ErrorMsg != "" {
            return fmt.Errorf(showRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error from show-static-mroute")
    }

    routeData := showRes.GetData()

    // Find the next-hop entry matching next_hop_gateway + priority.
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

    if ping, ok := routeData["ping"]; ok {
        if bv, bOk := ping.(bool); bOk {
            d.Set("ping", bv)
        }
    }

    if !found {
        d.SetId("")
        return nil
    }

    return nil
}

func deleteGaiaSetStaticMrouteNextHopPriority(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

