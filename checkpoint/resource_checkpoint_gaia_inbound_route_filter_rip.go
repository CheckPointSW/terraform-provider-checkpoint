package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaInboundRouteFilterRip() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaInboundRouteFilterRip,
        Read:   readGaiaInboundRouteFilterRip,
        Update: updateGaiaInboundRouteFilterRip,
        Delete: deleteGaiaInboundRouteFilterRip,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "restrict_all_ipv4": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `When the specified value is set to true, the policy rule rejects all matching IPv4 routes, except when there exists a more specific filter, which is set to \"accept\".<br><br>When the specified value is set to false, the policy rule accepts all matching IPv4 routes, except when there exists a more specific filter, which rejects the routes. By default, the rule accepts all IPv4 routes`,
            },
            "rank": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Assigns a rank to all incoming routes matching the filter. Rank is used by the routing system when there are routes from different protocols to the same destination. The route from the protocol with the lowest rank will be used.<br><br>Note: This value cannot be specified when rule is set to restrict`,
            },
            "route": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configures filtering of imported IPv4 routes for a given policy rule`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "subnet": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the address range with which to filter imported IPv4 routes`,
                        },
                        "restrict": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `When the specified value is true, all routes matching this rule will be rejected, unless a more specific filter accepts the imported routes.<br>When the specified value is false, all routes matching this rule will be accepted, unless a more specific filter accepts them. By default, the given route will be accepted`,
                        },
                        "match_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Routes can be matched with the following types: <br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>exact</td> <td>Matches only routes with prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>refines</td> <td>Matches only routes that are contained within the specified network (i.e., with greater mask length)</td> </tr><tr> <td>between</td> <td>Matches any route with prefix equal to the specified network whose mask length falls within a particular range</td> </tr></table>`,
                        },
                        "range": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Specifies the range with which to match the routes.<br><br>This attribute can only be specified when the match type is \"between\"`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "from": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Specifies the lower limit of the range of mask lengths`,
                                    },
                                    "to": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Specifies the upper limit of the range of mask lengths`,
                                    },
                                },
                            },
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
                Description: `Reset Inbound Route Filter configuration to a default state`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaInboundRouteFilterRip(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("restrict_all_ipv4"); ok {
        payload["restrict-all-ipv4"] = v.(bool)
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
            if sv, ok := d.GetOk(fmt.Sprintf("route.%d.range", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    rangeMap := make(map[string]interface{})
                    if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                        rangeMap["from"] = sv.(int)
                    }
                    if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                        rangeMap["to"] = sv.(int)
                    }
                    if len(rangeMap) > 0 {
                        itemMap["range"] = rangeMap
                    }
                }
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

    log.Println("Create InboundRouteFilterRip - Map = ", payload)

    addInboundRouteFilterRipRes, err := client.ApiCallSimple("set-inbound-route-filter-rip", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addInboundRouteFilterRipRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addInboundRouteFilterRipRes.Success {
            errMsg = addInboundRouteFilterRipRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addInboundRouteFilterRipRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-rip",        // resource type
            "create",                       // operation
            "set-inbound-route-filter-rip",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add inbound-route-filter-rip: %v", err)
    }
    if !addInboundRouteFilterRipRes.Success {
        if addInboundRouteFilterRipRes.ErrorMsg != "" {
            return fmt.Errorf(addInboundRouteFilterRipRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("inbound-route-filter-rip-" + acctest.RandString(10)))
    return readGaiaInboundRouteFilterRip(d, m)
}

func readGaiaInboundRouteFilterRip(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showInboundRouteFilterRipRes, err := client.ApiCallSimple("show-inbound-route-filter-rip", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showInboundRouteFilterRipRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showInboundRouteFilterRipRes.Success {
            errMsg = showInboundRouteFilterRipRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showInboundRouteFilterRipRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-rip",        // resource type
            "read",                       // operation
            "show-inbound-route-filter-rip",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show inbound-route-filter-rip: %v", err)
    }
    if !showInboundRouteFilterRipRes.Success {
        if data := showInboundRouteFilterRipRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showInboundRouteFilterRipRes.ErrorMsg)
    }

    inboundRouteFilterRip := showInboundRouteFilterRipRes.GetData()

    log.Println("Read InboundRouteFilterRip - Show JSON = ", inboundRouteFilterRip)

    if val, ok := inboundRouteFilterRip["rank"]; ok {
        d.Set("rank", fmt.Sprintf("%v", val))
    }
    if val, ok := inboundRouteFilterRip["restrict-all-ipv4"]; ok {
        if b, ok := val.(bool); ok {
            d.Set("restrict_all_ipv4", b)
        }
    }
    if val, ok := inboundRouteFilterRip["route"]; ok {
        if routeList, ok := val.([]interface{}); ok {
            routes := make([]interface{}, 0, len(routeList))
            for _, r := range routeList {
                if rm, ok := r.(map[string]interface{}); ok {
                    entry := map[string]interface{}{}
                    if sv, ok := rm["subnet"]; ok { entry["subnet"] = fmt.Sprintf("%v", sv) }
                    if sv, ok := rm["restrict"]; ok { if b, ok := sv.(bool); ok { entry["restrict"] = b } }
                    if sv, ok := rm["match-type"]; ok { entry["match_type"] = fmt.Sprintf("%v", sv) }
                    if sv, ok := rm["rank"]; ok { entry["rank"] = fmt.Sprintf("%v", sv) }
                    if sv, ok := rm["range"]; ok {
                        if rangeMap, ok := sv.(map[string]interface{}); ok {
                            re := map[string]interface{}{}
                            if fv, ok := rangeMap["from"]; ok { var n int; if _, err := fmt.Sscanf(fmt.Sprintf("%v", fv), "%d", &n); err == nil { re["from"] = n } }
                            if tv, ok := rangeMap["to"]; ok { var n int; if _, err := fmt.Sscanf(fmt.Sprintf("%v", tv), "%d", &n); err == nil { re["to"] = n } }
                            if len(re) > 0 { entry["range"] = []interface{}{re} }
                        }
                    }
                    routes = append(routes, entry)
                }
            }
            d.Set("route", routes)
        }
    }
    if v, exists := inboundRouteFilterRip["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaInboundRouteFilterRip(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("restrict_all_ipv4"); ok {
        payload["restrict-all-ipv4"] = v.(bool)
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
            if sv, ok := d.GetOk(fmt.Sprintf("route.%d.range", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    rangeMap := make(map[string]interface{})
                    if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                        rangeMap["from"] = sv.(int)
                    }
                    if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                        rangeMap["to"] = sv.(int)
                    }
                    if len(rangeMap) > 0 {
                        itemMap["range"] = rangeMap
                    }
                }
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

    setInboundRouteFilterRipRes, err := client.ApiCallSimple("set-inbound-route-filter-rip", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setInboundRouteFilterRipRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setInboundRouteFilterRipRes.Success {
            errMsg = setInboundRouteFilterRipRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setInboundRouteFilterRipRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-rip",        // resource type
            "update",                       // operation
            "set-inbound-route-filter-rip",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set inbound-route-filter-rip: %v", err)
    }
    if !setInboundRouteFilterRipRes.Success {
        return fmt.Errorf(setInboundRouteFilterRipRes.ErrorMsg)
    }

    return readGaiaInboundRouteFilterRip(d, m)
}

func deleteGaiaInboundRouteFilterRip(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    