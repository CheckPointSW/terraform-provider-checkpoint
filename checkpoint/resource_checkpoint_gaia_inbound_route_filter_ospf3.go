package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaInboundRouteFilterOspf3() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaInboundRouteFilterOspf3,
        Read:   readGaiaInboundRouteFilterOspf3,
        Update: updateGaiaInboundRouteFilterOspf3,
        Delete: deleteGaiaInboundRouteFilterOspf3,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures filtering of IPv6 routes for a specific OSPFv3 instance.<br><br>Note: The specified OSPFv3 instance must be configured`,
            },
            "restrict_all_ipv6": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `When the specified value is set to true, the policy rule rejects all matching IPv6 routes, except when there exists a more specific rule, which is set to \"accept\".<br><br>When the specified value is set to false, the policy rule accepts all matching IPv6 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv6 routes`,
            },
            "rank": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Assigns a rank to all incoming routes matching the filter. Rank is used by the routing system when there are routes from different protocols to the same destination. The route from the protocol with the lowest rank will be used.<br><br>Note: This value cannot be specified when rule is set to restrict`,
            },
            "route": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configures filtering of imported IPv6 routes for a given policy rule`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "subnet": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the address range with which to filter imported IPv6 routes`,
                        },
                        "restrict": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `When the specified value is true, all routes matching this rule will be rejected, unless a more specific filter accepts the imported routes.<br>When the specified value is false, all routes matching this rule will be accepted, unless a more specific filter accepts them. By default, the given route will be accepted`,
                        },
                        "match_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Routes can be matched with the following types: <br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>exact</td> <td>Matches only routes with prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>refines</td> <td>Matches only routes that are contained within the specified network (i.e., with greater mask length)</td> </tr></table>`,
                        },
                        "rank": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Assigns a rank to all incoming routes matching this filter, except those matching a more specific rule with a different rank configured.<br><br>Rank is used by the routing system when there are routes from different protocols to the same destination. The route with the lowest rank from the protocol will be used`,
                        },
                    },
                },
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Resets Inbound Route Filter configuration to a default state for the given IPv6 OSPF Instance`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaInboundRouteFilterOspf3(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv6"); ok {
        payload["restrict-all-ipv6"] = v.(bool)
    }

    if v, ok := d.GetOk("rank"); ok {
        payload["rank"] = v.(string)
    }

    if v := d.Get("route"); len(v.([]interface{})) > 0 {
        routeList := v.([]interface{})
        routeArray := make([]interface{}, 0, len(routeList))
        for i := range routeList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v := d.Get(fmt.Sprintf("route.%d.restrict", i)).(bool); v {
                itemMap["restrict"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.match_type", i)); ok {
                itemMap["match-type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.rank", i)); ok {
                itemMap["rank"] = v.(string)
            }
            if len(itemMap) > 0 {
                routeArray = append(routeArray, itemMap)
            }
        }
        if len(routeArray) > 0 {
            payload["route"] = routeArray
        }
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    log.Println("Create InboundRouteFilterOspf3 - Map = ", payload)

    addInboundRouteFilterOspf3Res, err := client.ApiCallSimple("set-inbound-route-filter-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addInboundRouteFilterOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addInboundRouteFilterOspf3Res.Success {
            errMsg = addInboundRouteFilterOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addInboundRouteFilterOspf3Res.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-ospf3",        // resource type
            "create",                       // operation
            "set-inbound-route-filter-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add inbound-route-filter-ospf3: %v", err)
    }
    if !addInboundRouteFilterOspf3Res.Success {
        if addInboundRouteFilterOspf3Res.ErrorMsg != "" {
            return fmt.Errorf(addInboundRouteFilterOspf3Res.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("inbound-route-filter-ospf3-" + acctest.RandString(10)))
    return readGaiaInboundRouteFilterOspf3(d, m)
}

func readGaiaInboundRouteFilterOspf3(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showInboundRouteFilterOspf3Res, err := client.ApiCallSimple("show-inbound-route-filter-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showInboundRouteFilterOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showInboundRouteFilterOspf3Res.Success {
            errMsg = showInboundRouteFilterOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showInboundRouteFilterOspf3Res.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-ospf3",        // resource type
            "read",                       // operation
            "show-inbound-route-filter-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show inbound-route-filter-ospf3: %v", err)
    }
    if !showInboundRouteFilterOspf3Res.Success {
        if data := showInboundRouteFilterOspf3Res.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showInboundRouteFilterOspf3Res.ErrorMsg)
    }

    inboundRouteFilterOspf3 := showInboundRouteFilterOspf3Res.GetData()

    log.Println("Read InboundRouteFilterOspf3 - Show JSON = ", inboundRouteFilterOspf3)

    if v, exists := inboundRouteFilterOspf3["ospf3"]; exists {
        if items, ok := v.([]interface{}); ok && len(items) > 0 {
            item, _ := items[0].(map[string]interface{})
            if val, ok := item["instance"]; ok { d.Set("instance", fmt.Sprintf("%v", val)) }
            if val, ok := item["rank"]; ok { d.Set("rank", fmt.Sprintf("%v", val)) }
            if val, ok := item["restrict-all-ipv6"]; ok { if b, ok := val.(bool); ok { d.Set("restrict_all_ipv6", b) } }
            if val, ok := item["route"]; ok {
                if routeList, ok := val.([]interface{}); ok {
                    routes := make([]interface{}, 0, len(routeList))
                    for _, r := range routeList {
                        if rm, ok := r.(map[string]interface{}); ok {
                            entry := map[string]interface{}{}
                            if sv, ok := rm["subnet"]; ok { entry["subnet"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := rm["restrict"]; ok { if b, ok := sv.(bool); ok { entry["restrict"] = b } }
                            if sv, ok := rm["match-type"]; ok { entry["match_type"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := rm["rank"]; ok { entry["rank"] = fmt.Sprintf("%v", sv) }
                            routes = append(routes, entry)
                        }
                    }
                    d.Set("route", routes)
                }
            }
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaInboundRouteFilterOspf3(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv6"); ok {
        payload["restrict-all-ipv6"] = v.(bool)
    }

    if v, ok := d.GetOk("rank"); ok {
        payload["rank"] = v.(string)
    }

    if v := d.Get("route"); len(v.([]interface{})) > 0 {
        routeList := v.([]interface{})
        routeArray := make([]interface{}, 0, len(routeList))
        for i := range routeList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v := d.Get(fmt.Sprintf("route.%d.restrict", i)).(bool); v {
                itemMap["restrict"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.match_type", i)); ok {
                itemMap["match-type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.rank", i)); ok {
                itemMap["rank"] = v.(string)
            }
            if len(itemMap) > 0 {
                routeArray = append(routeArray, itemMap)
            }
        }
        if len(routeArray) > 0 {
            payload["route"] = routeArray
        }
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    setInboundRouteFilterOspf3Res, err := client.ApiCallSimple("set-inbound-route-filter-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setInboundRouteFilterOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setInboundRouteFilterOspf3Res.Success {
            errMsg = setInboundRouteFilterOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setInboundRouteFilterOspf3Res.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-ospf3",        // resource type
            "update",                       // operation
            "set-inbound-route-filter-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set inbound-route-filter-ospf3: %v", err)
    }
    if !setInboundRouteFilterOspf3Res.Success {
        return fmt.Errorf(setInboundRouteFilterOspf3Res.ErrorMsg)
    }

    return readGaiaInboundRouteFilterOspf3(d, m)
}

func deleteGaiaInboundRouteFilterOspf3(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    