package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaStaticMroute() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaStaticMroute,
        Read:   readGaiaStaticMroute,
        Update: updateGaiaStaticMroute,
        Delete: deleteGaiaStaticMroute,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Address of the static-mroute to set configuration for.`,
            },
            "mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Mask length for the static-mroute.`,
            },
            "next_hop": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Static next-hop. Contains a list of next-hop gateways.<br><br>Each gateway is formatted in the following manner:<br>{\"gateway\": IP address, \"priority\": default or integer 1-8}`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "gateway": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IP address for the static next-hop gateway.`,
                        },
                        "priority": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Priority defines which gateway to select as the next-hop: the lower the priority, the higher the preference.<br>Possible values: default or integer 1-8`,
                        },
                    },
                },
            },
            "ping": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configures ping monitoring of the given IPv4 static-mroute.<br>Possible values: true, false<br><br>NOTE: Static-mroute ping is not supported in versions prior to R82.10`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaStaticMroute(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v := d.Get("next_hop"); len(v.([]interface{})) > 0 {
        nexthopList := v.([]interface{})
        nexthopArray := make([]interface{}, 0, len(nexthopList))
        for i := range nexthopList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("next_hop.%d.gateway", i)); ok {
                itemMap["gateway"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("next_hop.%d.priority", i)); ok {
                itemMap["priority"] = v.(string)
            }
            if len(itemMap) > 0 {
                nexthopArray = append(nexthopArray, itemMap)
            }
        }
        if len(nexthopArray) > 0 {
            payload["next-hop"] = nexthopArray
        }
    }

    if v, ok := d.GetOkExists("ping"); ok {
        payload["ping"] = v.(bool)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create StaticMroute - Map = ", payload)

    addStaticMrouteRes, err := client.ApiCallSimple("set-static-mroute", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addStaticMrouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addStaticMrouteRes.Success {
            errMsg = addStaticMrouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addStaticMrouteRes.GetData()
        }

        debugLogOperation(
            "static-mroute",        // resource type
            "create",                       // operation
            "set-static-mroute",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add static-mroute: %v", err)
    }
    if !addStaticMrouteRes.Success {
        if addStaticMrouteRes.ErrorMsg != "" {
            return fmt.Errorf(addStaticMrouteRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("static-mroute-" + acctest.RandString(10)))
    return readGaiaStaticMroute(d, m)
}

func readGaiaStaticMroute(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showStaticMrouteRes, err := client.ApiCallSimple("show-static-mroute", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showStaticMrouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showStaticMrouteRes.Success {
            errMsg = showStaticMrouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showStaticMrouteRes.GetData()
        }

        debugLogOperation(
            "static-mroute",        // resource type
            "read",                       // operation
            "show-static-mroute",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show static-mroute: %v", err)
    }
    if !showStaticMrouteRes.Success {
        if data := showStaticMrouteRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showStaticMrouteRes.ErrorMsg)
    }

    staticMroute := showStaticMrouteRes.GetData()

    log.Println("Read StaticMroute - Show JSON = ", staticMroute)

    if v, exists := staticMroute["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := staticMroute["mask-length"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("mask_length", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("mask_length", _n)
            }
        }
    }
    if v, exists := staticMroute["next-hop"]; exists {
        d.Set("next_hop", v.([]interface{}))
    }
    if v, exists := staticMroute["ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("ping", s == "true")
        }
    }
    if v, exists := staticMroute["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaStaticMroute(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v := d.Get("next_hop"); len(v.([]interface{})) > 0 {
        nexthopList := v.([]interface{})
        nexthopArray := make([]interface{}, 0, len(nexthopList))
        for i := range nexthopList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("next_hop.%d.gateway", i)); ok {
                itemMap["gateway"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("next_hop.%d.priority", i)); ok {
                itemMap["priority"] = v.(string)
            }
            if len(itemMap) > 0 {
                nexthopArray = append(nexthopArray, itemMap)
            }
        }
        if len(nexthopArray) > 0 {
            payload["next-hop"] = nexthopArray
        }
    }

    if v, ok := d.GetOkExists("ping"); ok {
        payload["ping"] = v.(bool)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setStaticMrouteRes, err := client.ApiCallSimple("set-static-mroute", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setStaticMrouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setStaticMrouteRes.Success {
            errMsg = setStaticMrouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setStaticMrouteRes.GetData()
        }

        debugLogOperation(
            "static-mroute",        // resource type
            "update",                       // operation
            "set-static-mroute",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set static-mroute: %v", err)
    }
    if !setStaticMrouteRes.Success {
        return fmt.Errorf(setStaticMrouteRes.ErrorMsg)
    }

    return readGaiaStaticMroute(d, m)
}

func deleteGaiaStaticMroute(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    deleteStaticMrouteRes, err := client.ApiCallSimple("delete-static-mroute", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteStaticMrouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteStaticMrouteRes.Success {
            errMsg = deleteStaticMrouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteStaticMrouteRes.GetData()
        }

        debugLogOperation(
            "static-mroute",        // resource type
            "delete",                       // operation
            "delete-static-mroute",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete static-mroute: %v", err)
    }
    if !deleteStaticMrouteRes.Success {
        return fmt.Errorf(deleteStaticMrouteRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

