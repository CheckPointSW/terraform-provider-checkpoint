package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaStaticRoute() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaStaticRoute,
        Read:   readGaiaStaticRoute,
        Update: updateGaiaStaticRoute,
        Delete: deleteGaiaStaticRoute,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `N/A`,
            },
            "type": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Type of next hop. Possible values: blackhole, gateway, reject`,
            },
            "next_hop": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Static next-hop. Contains a list of next-hop gateways. Each gateway is formatted in the following manner: {\"gateway\": IP address or logical name, \"priority\": default or integer 1-8}`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "gateway": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IP address or logical name for the static next-hop gateway`,
                        },
                        "priority": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Priority defines which gateway to select as the next-hop. The lower the priority, the higher the preference. Possible values: default or integer 1-8`,
                        },
                    },
                },
            },
            "ping": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `Configures ping monitoring of the given IPv4 static route. Possible values: true, false`,
            },
            "rank": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Selects a route when there are many routes to a destination that use different routing protocols. The route with the lowest rank value is selected. Possible values: default or integer 0-255`,
            },
            "scope_local": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `Configure the local-interface scope option, When the this option is enabled, the route treated as directly connected to local machine. Possible values: true, false`,
            },
            "comment": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
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

func createGaiaStaticRoute(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
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

    if v, ok := d.GetOk("rank"); ok {
        payload["rank"] = v.(int)
    }

    if v, ok := d.GetOkExists("scope_local"); ok {
        payload["scope-local"] = v.(bool)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create StaticRoute - Map = ", payload)

    addStaticRouteRes, err := client.ApiCallSimple("set-static-route", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addStaticRouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addStaticRouteRes.Success {
            errMsg = addStaticRouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addStaticRouteRes.GetData()
        }

        debugLogOperation(
            "static-route",        // resource type
            "create",                       // operation
            "set-static-route",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add static-route: %v", err)
    }
    if !addStaticRouteRes.Success {
        if addStaticRouteRes.ErrorMsg != "" {
            return fmt.Errorf(addStaticRouteRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("static-route-" + acctest.RandString(10)))
    return readGaiaStaticRoute(d, m)
}

func readGaiaStaticRoute(d *schema.ResourceData, m interface{}) error {

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

    showStaticRouteRes, err := client.ApiCallSimple("show-static-route", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showStaticRouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showStaticRouteRes.Success {
            errMsg = showStaticRouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showStaticRouteRes.GetData()
        }

        debugLogOperation(
            "static-route",        // resource type
            "read",                       // operation
            "show-static-route",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show static-route: %v", err)
    }
    if !showStaticRouteRes.Success {
        if data := showStaticRouteRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showStaticRouteRes.ErrorMsg)
    }

    staticRoute := showStaticRouteRes.GetData()

    log.Println("Read StaticRoute - Show JSON = ", staticRoute)

    if v, exists := staticRoute["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := staticRoute["mask-length"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("mask_length", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("mask_length", _n)
            }
        }
    }
    if v, exists := staticRoute["type"]; exists {
        d.Set("type", fmt.Sprintf("%v", v))
    }
    if v, exists := staticRoute["next-hop"]; exists {
        d.Set("next_hop", v.([]interface{}))
    }
    if v, exists := staticRoute["ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("ping", s == "true")
        }
    }
    if v, exists := staticRoute["rank"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rank", int(f))
        }
    }
    if v, exists := staticRoute["scope-local"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("scope_local", b)
        } else if s, ok := v.(string); ok {
            d.Set("scope_local", s == "true")
        }
    }
    if v, exists := staticRoute["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    if v, exists := staticRoute["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := staticRoute["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaStaticRoute(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("mask_length"); ok {
        payload["mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("type"); ok {
        payload["type"] = v.(string)
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

    if v, ok := d.GetOk("rank"); ok {
        payload["rank"] = v.(int)
    }

    if v, ok := d.GetOkExists("scope_local"); ok {
        payload["scope-local"] = v.(bool)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setStaticRouteRes, err := client.ApiCallSimple("set-static-route", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setStaticRouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setStaticRouteRes.Success {
            errMsg = setStaticRouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setStaticRouteRes.GetData()
        }

        debugLogOperation(
            "static-route",        // resource type
            "update",                       // operation
            "set-static-route",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set static-route: %v", err)
    }
    if !setStaticRouteRes.Success {
        return fmt.Errorf(setStaticRouteRes.ErrorMsg)
    }

    return readGaiaStaticRoute(d, m)
}

func deleteGaiaStaticRoute(d *schema.ResourceData, m interface{}) error {

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

    deleteStaticRouteRes, err := client.ApiCallSimple("delete-static-route", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteStaticRouteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteStaticRouteRes.Success {
            errMsg = deleteStaticRouteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteStaticRouteRes.GetData()
        }

        debugLogOperation(
            "static-route",        // resource type
            "delete",                       // operation
            "delete-static-route",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete static-route: %v", err)
    }
    if !deleteStaticRouteRes.Success {
        return fmt.Errorf(deleteStaticRouteRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

