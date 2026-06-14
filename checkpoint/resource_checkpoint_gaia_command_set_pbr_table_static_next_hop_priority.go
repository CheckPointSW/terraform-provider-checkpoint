package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"
)
func dataGaiaSetPbrTableStaticNextHopPriority() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetPbrTableStaticNextHopPriority,
        Read:   readGaiaSetPbrTableStaticNextHopPriority,
        Delete: deleteGaiaSetPbrTableStaticNextHopPriority,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "table": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Name of PBR Table`,
            },
            "static_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `IP address of PBR Table static route`,
            },
            "static_mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `Mask length of PBR Table static route`,
            },
            "next_hop_gateway": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Nexthop gateway of PBR Table static route, can be IP address or interface name`,
            },
            "priority": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `This value will replace the current priority of the specified nexthop gateway. Priority defines which gateway to select as the next hop, the lower the priority, the higher the preference. Can be default or integer from 1 to 8`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "message": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSetPbrTableStaticNextHopPriority(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }

    if v, ok := d.GetOk("static_address"); ok {
        payload["static-address"] = v.(string)
    }

    if v, ok := d.GetOk("static_mask_length"); ok {
        payload["static-mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("next_hop_gateway"); ok {
        payload["next-hop-gateway"] = v.(string)
    }

    if v, ok := d.GetOk("priority"); ok {
        payload["priority"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute set-pbr-table-static-next-hop-priority - Payload = ", payload)

    GaiaSetPbrTableStaticNextHopPriorityRes, err := client.ApiCallSimple("set-pbr-table-static-next-hop-priority", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetPbrTableStaticNextHopPriorityRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetPbrTableStaticNextHopPriorityRes.Success {
            errMsg = GaiaSetPbrTableStaticNextHopPriorityRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetPbrTableStaticNextHopPriorityRes.GetData()
        }

        debugLogOperation(
            "set-pbr-table-static-next-hop-priority",        // resource type
            "command",                       // operation
            "set-pbr-table-static-next-hop-priority",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-pbr-table-static-next-hop-priority: %v", err)
    }
    if !GaiaSetPbrTableStaticNextHopPriorityRes.Success {
        if GaiaSetPbrTableStaticNextHopPriorityRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetPbrTableStaticNextHopPriorityRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaSetPbrTableStaticNextHopPriorityRes.GetData()
    if v, exists := _respData["message"]; exists {
        d.Set("message", toString(v))
    }


    d.SetId(fmt.Sprintf("set-pbr-table-static-next-hop-priority-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetPbrTableStaticNextHopPriority(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }
    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Read set-pbr-table-static-next-hop-priority - show-pbr-table payload = ", payload)

    showRes, err := client.ApiCallSimple("show-pbr-table", payload)
    if err != nil {
        return fmt.Errorf("Failed to show pbr-table: %v", err)
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
        return fmt.Errorf("Unknown error from show-pbr-table")
    }

    tableData := showRes.GetData()

    wantAddr := ""
    if v, ok := d.GetOk("static_address"); ok {
        wantAddr = v.(string)
    }
    wantMask := 0
    if v, ok := d.GetOk("static_mask_length"); ok {
        wantMask = v.(int)
    }
    wantGW := ""
    if v, ok := d.GetOk("next_hop_gateway"); ok {
        wantGW = v.(string)
    }
    wantPri := ""
    if v, ok := d.GetOk("priority"); ok {
        wantPri = v.(string)
    }

    found := false
    if routes, ok := tableData["static-routes"].([]interface{}); ok {
        for _, r := range routes {
            route, ok := r.(map[string]interface{})
            if !ok {
                continue
            }
            if fmt.Sprintf("%v", route["address"]) != wantAddr {
                continue
            }
            if fmt.Sprintf("%v", route["mask-length"]) != fmt.Sprintf("%d", wantMask) {
                continue
            }
            if nexthops, ok := route["next-hop"].([]interface{}); ok {
                for _, nh := range nexthops {
                    if nhMap, ok := nh.(map[string]interface{}); ok {
                        if fmt.Sprintf("%v", nhMap["gateway"]) == wantGW && fmt.Sprintf("%v", nhMap["priority"]) == wantPri {
                            found = true
                            break
                        }
                    }
                }
            }
            if found {
                break
            }
        }
    }

    if !found {
        d.SetId("")
        return nil
    }

    return nil
}

func deleteGaiaSetPbrTableStaticNextHopPriority(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

