package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaInboundRouteFilterBgpPolicy() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaInboundRouteFilterBgpPolicy,
        Read:   readGaiaInboundRouteFilterBgpPolicy,
        Update: updateGaiaInboundRouteFilterBgpPolicy,
        Delete: deleteGaiaInboundRouteFilterBgpPolicy,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "policy_id": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Specifies the BGP import policy identifier.<br><br>Note: In order to filter based on Autonomous System path, the policy identifier must be between 1-511.<br>In order to filter based on Autonomous System number, the policy identifier must be between 512-1024`,
            },
            "restrict_all_ipv4": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `When the specified value is set to true, the policy rule rejects all matching IPv4 routes, except when there exists a more specific rule, which is set to \"accept\".<br>When the specified value is set to false, the policy rule accepts all matching IPv4 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv4 routes`,
            },
            "restrict_all_ipv6": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `When the specified value is set to true, the policy rule rejects all matching IPv6 routes, except when there exists a more specific rule, which is set to \"accept\".<br>When the specified value is set to false, the policy rule accepts all matching IPv6 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv6 routes.<br><br>Note: The following value can only be specified when IPv6 state is enabled`,
            },
            "default_localpref": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Assigns a BGP local preference to all routes that match this filter. By default, no local preference value is assigned to the matched routes`,
            },
            "default_weight": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Assignes a BGP weight to all routes that match this filter. By default, no weight value is assigned to the matched routes`,
            },
            "community_match": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Matches routes containing a given BGP Community.<br><br>Note: A maximum of 25 Communites can be configured on each policy identifier`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Configures the Community ID value for BGP Communities`,
                        },
                        "as": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Configures the Autonomous System number for BGP Communities`,
                        },
                    },
                },
            },
            "extcommunity_match": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Matches routes containing a given BGP Extended Community<br><br>Note: A maximum of 25 Communities can be configured on each policy identifier`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configured Type for extended communities`,
                        },
                        "sub_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configured Sub-Type for extended communities. Valid sub type values are dependent on the type, the valid values are as follows:<br><br><table class=\"table\"><tr> <th>Type</th> <th>Sub Types</th> </tr><tr> <td>transitive-two-octet-as</td> <td>route-target, route-origin, ospf-domain-id, bgp-data-collect, source-as, l2vpn-id, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>link-bandwidth</td> </tr><tr> <td>transitive-four-octet-as</td> <td>route-target, route-origin, generic, ospf-domain-id, bgp-data-collect, source-as, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>generic</td> </tr><tr> <td>transitive-ipv4-address</td> <td>route-target, route-origin, ospf-domain-id, ospf-route-id, l2vpn-id, vrf-route-import, cisco-vpn-dist</td> </tr></table>`,
                        },
                        "value": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configured Value for extended communities. Valid values are dependent on the type, the valid values are as follows:<br><br><table class=\"table\"><tr> <th>Type</th> <th>Values</th> </tr><tr> <td>transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>transitive-ipv4-address</td> <td>IPv4:0 - 65,535</td> </tr></table>`,
                        },
                    },
                },
            },
            "based_on_as": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures a new policy for importing BGP routes from a particular Autonomous System.<br><br>Note: In order to configure filtering based on AS, the specified policy identifier must be in between 512-1024. Additionally the ASN cannot be configured in any other policy id`,
            },
            "based_on_aspath": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configures a new policy for importing BGP routes whose Autonomous Systems path matches the specified regular expression.<br><br>Note: In order to configure filtering based on AS path, the specified policy identifier must be in between 1-511. Additionally the AS path cannot be configured in any other policy id`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aspath_regex": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the regular expression, which is used to filter by Autonomous Systems paths. A valid AS path regular expression contains only digits and the following special characters:<br><br><table class=\"table\"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table>`,
                        },
                        "origin": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the completeness of the AS path information. The origin values are defined as follows:<br><br><table class=\"table\"><tr> <th>Origin</th> <th>Description</th> </tr><tr> <td>any</td> <td>Matches any route, regardless of origin</td> </tr><tr> <td>IGP</td> <td>Route was learned from an interior routing protocol, and the AS path is probably complete</td> </tr><tr> <td>EGP</td> <td>Route was learned from an exterior routing protocol that does not support AS paths, and the path is probably complete</td> </tr><tr> <td>incomplete</td> <td>Use when the AS path information is incomplete</td> </tr></table>`,
                        },
                    },
                },
            },
            "route": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configures filtering of imported routes for a given policy rule`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "subnet": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the address range with which to filter imported IPv4 and IPv6 routes.<br><br>Note: In order to configure subnets of type IPv6, the IPv6 state needs to be enabled`,
                        },
                        "restrict": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `When the specified value is true, all routes matching this rule will be rejected, unless a more specific filter accepts the imported routes.<br>When the specified value is false, all routes matching this rule will be accepted, unless a more specific filter accepts them. By default, the given route will be accepted`,
                        },
                        "match_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Routes can be matched with the following types: <br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>exact</td> <td>Matches only routes with prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>refines</td> <td>Matches only routes that are contained within the specified network (i.e., with greater mask length)</td> </tr><tr> <td>between</td> <td>Matches any route with prefix equal to the specified network whose mask length falls within a particular range</td> </tr></table><br><br>Note: When the given subnet is of type IPv6, the \"between\" value cannot be specified`,
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
                        "localpref": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Assigns a BGP local preference to all routes matching this filter, unless there exists a more specific rule with a different local preference value.<br><br>Note: The following value cannot be specified when the rule is restricted`,
                        },
                        "weight": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Assinges a BGP Weight to all routes matching this filter unless there exists a more specific rule with a different weight value.<br><br>Note: The following value cannot be specified when the rule is restricted`,
                        },
                    },
                },
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Resets Inbound Policy Filter Configuration to a default state for a given policy identifier`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaInboundRouteFilterBgpPolicy(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("policy_id"); ok {
        payload["policy-id"] = v.(int)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv4"); ok {
        payload["restrict-all-ipv4"] = v.(bool)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv6"); ok {
        payload["restrict-all-ipv6"] = v.(bool)
    }

    if v, ok := d.GetOk("default_localpref"); ok {
        payload["default-localpref"] = v.(string)
    }

    if v, ok := d.GetOk("default_weight"); ok {
        payload["default-weight"] = v.(string)
    }

    if v := d.Get("community_match"); len(v.([]interface{})) > 0 {
        communitymatchList := v.([]interface{})
        communitymatchArray := make([]interface{}, 0, len(communitymatchList))
        for i := range communitymatchList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("community_match.%d.resource_id", i)); ok {
                itemMap["id"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("community_match.%d.as", i)); ok {
                itemMap["as"] = v.(int)
            }
            if len(itemMap) > 0 {
                communitymatchArray = append(communitymatchArray, itemMap)
            }
        }
        if len(communitymatchArray) > 0 {
            payload["community-match"] = communitymatchArray
        }
    }

    if v := d.Get("extcommunity_match"); len(v.([]interface{})) > 0 {
        extcommunitymatchList := v.([]interface{})
        extcommunitymatchArray := make([]interface{}, 0, len(extcommunitymatchList))
        for i := range extcommunitymatchList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.sub_type", i)); ok {
                itemMap["sub-type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if len(itemMap) > 0 {
                extcommunitymatchArray = append(extcommunitymatchArray, itemMap)
            }
        }
        if len(extcommunitymatchArray) > 0 {
            payload["extcommunity-match"] = extcommunitymatchArray
        }
    }

    if v, ok := d.GetOk("based_on_as"); ok {
        payload["based-on-as"] = v.(string)
    }

    if v := d.Get("based_on_aspath"); len(v.([]interface{})) > 0 {
        _ = v
        basedonaspathMap := make(map[string]interface{})
        if v, ok := d.GetOk("based_on_aspath.0.aspath_regex"); ok {
            basedonaspathMap["aspath-regex"] = v.(string)
        }
        if v, ok := d.GetOk("based_on_aspath.0.origin"); ok {
            basedonaspathMap["origin"] = v.(string)
        }
        if len(basedonaspathMap) > 0 {
            payload["based-on-aspath"] = basedonaspathMap
        }
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
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.localpref", i)); ok {
                itemMap["localpref"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.weight", i)); ok {
                itemMap["weight"] = v.(string)
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

    log.Println("Create InboundRouteFilterBgpPolicy - Map = ", payload)

    // Check if inbound-route-filter-bgp-policy already exists on the device.
    _checkPayload := map[string]interface{}{"policy-id": d.Get("policy_id")}
    _existsRes, _existsErr := client.ApiCallSimple("show-inbound-route-filter-bgp-policy", _checkPayload)
    if _existsErr == nil && _existsRes.Success && func() bool {
            _d := _existsRes.GetData()
            if _bgp, ok := _d["bgp"]; ok {
                if _bgpMap, isMap := _bgp.(map[string]interface{}); isMap && len(_bgpMap) > 0 {
                    return true
                }
            }
            return false
        }() {
        delete(payload, "based-on-aspath")
        delete(payload, "based-on-as")
    }

    addInboundRouteFilterBgpPolicyRes, err := client.ApiCallSimple("set-inbound-route-filter-bgp-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addInboundRouteFilterBgpPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addInboundRouteFilterBgpPolicyRes.Success {
            errMsg = addInboundRouteFilterBgpPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addInboundRouteFilterBgpPolicyRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-bgp-policy",        // resource type
            "create",                       // operation
            "set-inbound-route-filter-bgp-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add inbound-route-filter-bgp-policy: %v", err)
    }
    if !addInboundRouteFilterBgpPolicyRes.Success {
        if addInboundRouteFilterBgpPolicyRes.ErrorMsg != "" {
            return fmt.Errorf(addInboundRouteFilterBgpPolicyRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("inbound-route-filter-bgp-policy-" + acctest.RandString(10)))
    return readGaiaInboundRouteFilterBgpPolicy(d, m)
}

func readGaiaInboundRouteFilterBgpPolicy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("policy_id"); ok {
        payload["policy-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showInboundRouteFilterBgpPolicyRes, err := client.ApiCallSimple("show-inbound-route-filter-bgp-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showInboundRouteFilterBgpPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showInboundRouteFilterBgpPolicyRes.Success {
            errMsg = showInboundRouteFilterBgpPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showInboundRouteFilterBgpPolicyRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-bgp-policy",        // resource type
            "read",                       // operation
            "show-inbound-route-filter-bgp-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show inbound-route-filter-bgp-policy: %v", err)
    }
    if !showInboundRouteFilterBgpPolicyRes.Success {
        if data := showInboundRouteFilterBgpPolicyRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showInboundRouteFilterBgpPolicyRes.ErrorMsg)
    }

    inboundRouteFilterBgpPolicy := showInboundRouteFilterBgpPolicyRes.GetData()

    log.Println("Read InboundRouteFilterBgpPolicy - Show JSON = ", inboundRouteFilterBgpPolicy)

    if v, exists := inboundRouteFilterBgpPolicy["bgp"]; exists {
        if items, ok := v.([]interface{}); ok && len(items) > 0 {
            item, _ := items[0].(map[string]interface{})
            if val, ok := item["policy-id"]; ok { if f, ok := val.(float64); ok { d.Set("policy_id", int(f)) } }
            if val, ok := item["restrict-all-ipv4"]; ok { if b, ok := val.(bool); ok { d.Set("restrict_all_ipv4", b) } }
            if val, ok := item["restrict-all-ipv6"]; ok { if b, ok := val.(bool); ok { d.Set("restrict_all_ipv6", b) } }
            if val, ok := item["default-localpref"]; ok { d.Set("default_localpref", fmt.Sprintf("%v", val)) }
            if val, ok := item["default-weight"]; ok { d.Set("default_weight", fmt.Sprintf("%v", val)) }
            if val, ok := item["based-on-as"]; ok { d.Set("based_on_as", fmt.Sprintf("%v", val)) }
            if val, ok := item["based-on-aspath"]; ok {
                if bm, ok := val.(map[string]interface{}); ok {
                    baEntry := map[string]interface{}{}
                    if sv, ok := bm["aspath-regex"]; ok { baEntry["aspath_regex"] = fmt.Sprintf("%v", sv) }
                    if sv, ok := bm["origin"]; ok { baEntry["origin"] = fmt.Sprintf("%v", sv) }
                    if len(baEntry) > 0 { d.Set("based_on_aspath", []interface{}{baEntry}) }
                }
            }
            if val, ok := item["community-match"]; ok {
                if cmList, ok := val.([]interface{}); ok {
                    cms := make([]interface{}, 0, len(cmList))
                    for _, c := range cmList {
                        if cm, ok := c.(map[string]interface{}); ok {
                            entry := map[string]interface{}{}
                            if sv, ok := cm["id"]; ok { if f, ok := sv.(float64); ok { entry["resource_id"] = int(f) } else { var n int; if _, err := fmt.Sscanf(fmt.Sprintf("%v", sv), "%d", &n); err == nil { entry["resource_id"] = n } } }
                            if sv, ok := cm["as"]; ok { if f, ok := sv.(float64); ok { entry["as"] = int(f) } else { var n int; if _, err := fmt.Sscanf(fmt.Sprintf("%v", sv), "%d", &n); err == nil { entry["as"] = n } } }
                            cms = append(cms, entry)
                        }
                    }
                    d.Set("community_match", cms)
                }
            }
            if val, ok := item["extcommunity-match"]; ok {
                if ecList, ok := val.([]interface{}); ok {
                    ecs := make([]interface{}, 0, len(ecList))
                    for _, c := range ecList {
                        if ec, ok := c.(map[string]interface{}); ok {
                            entry := map[string]interface{}{}
                            if sv, ok := ec["type"]; ok { entry["type"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := ec["sub-type"]; ok { entry["sub_type"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := ec["value"]; ok { entry["value"] = fmt.Sprintf("%v", sv) }
                            ecs = append(ecs, entry)
                        }
                    }
                    d.Set("extcommunity_match", ecs)
                }
            }
            if val, ok := item["route"]; ok {
                if routeList, ok := val.([]interface{}); ok {
                    routes := make([]interface{}, 0, len(routeList))
                    for _, r := range routeList {
                        if rm, ok := r.(map[string]interface{}); ok {
                            entry := map[string]interface{}{}
                            if sv, ok := rm["subnet"]; ok { entry["subnet"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := rm["restrict"]; ok { if b, ok := sv.(bool); ok { entry["restrict"] = b } }
                            if sv, ok := rm["match-type"]; ok { entry["match_type"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := rm["localpref"]; ok { entry["localpref"] = fmt.Sprintf("%v", sv) }
                            if sv, ok := rm["weight"]; ok { entry["weight"] = fmt.Sprintf("%v", sv) }
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
        }
    }
    if v, exists := inboundRouteFilterBgpPolicy["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaInboundRouteFilterBgpPolicy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("policy_id"); ok {
        payload["policy-id"] = v.(int)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv4"); ok {
        payload["restrict-all-ipv4"] = v.(bool)
    }

    if v, ok := d.GetOkExists("restrict_all_ipv6"); ok {
        payload["restrict-all-ipv6"] = v.(bool)
    }

    if v, ok := d.GetOk("default_localpref"); ok {
        payload["default-localpref"] = v.(string)
    }

    if v, ok := d.GetOk("default_weight"); ok {
        payload["default-weight"] = v.(string)
    }

    if v := d.Get("community_match"); len(v.([]interface{})) > 0 {
        communitymatchList := v.([]interface{})
        communitymatchArray := make([]interface{}, 0, len(communitymatchList))
        for i := range communitymatchList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("community_match.%d.resource_id", i)); ok {
                itemMap["id"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("community_match.%d.as", i)); ok {
                itemMap["as"] = v.(int)
            }
            if len(itemMap) > 0 {
                communitymatchArray = append(communitymatchArray, itemMap)
            }
        }
        if len(communitymatchArray) > 0 {
            payload["community-match"] = communitymatchArray
        }
    }

    if v := d.Get("extcommunity_match"); len(v.([]interface{})) > 0 {
        extcommunitymatchList := v.([]interface{})
        extcommunitymatchArray := make([]interface{}, 0, len(extcommunitymatchList))
        for i := range extcommunitymatchList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.sub_type", i)); ok {
                itemMap["sub-type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("extcommunity_match.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if len(itemMap) > 0 {
                extcommunitymatchArray = append(extcommunitymatchArray, itemMap)
            }
        }
        if len(extcommunitymatchArray) > 0 {
            payload["extcommunity-match"] = extcommunitymatchArray
        }
    }

    if v, ok := d.GetOk("based_on_as"); ok {
        payload["based-on-as"] = v.(string)
    }

    if v := d.Get("based_on_aspath"); len(v.([]interface{})) > 0 {
        _ = v
        basedonaspathMap := make(map[string]interface{})
        if v, ok := d.GetOk("based_on_aspath.0.aspath_regex"); ok {
            basedonaspathMap["aspath-regex"] = v.(string)
        }
        if v, ok := d.GetOk("based_on_aspath.0.origin"); ok {
            basedonaspathMap["origin"] = v.(string)
        }
        if len(basedonaspathMap) > 0 {
            payload["based-on-aspath"] = basedonaspathMap
        }
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
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.localpref", i)); ok {
                itemMap["localpref"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("route.%d.weight", i)); ok {
                itemMap["weight"] = v.(string)
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

    setInboundRouteFilterBgpPolicyRes, err := client.ApiCallSimple("set-inbound-route-filter-bgp-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setInboundRouteFilterBgpPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setInboundRouteFilterBgpPolicyRes.Success {
            errMsg = setInboundRouteFilterBgpPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setInboundRouteFilterBgpPolicyRes.GetData()
        }

        debugLogOperation(
            "inbound-route-filter-bgp-policy",        // resource type
            "update",                       // operation
            "set-inbound-route-filter-bgp-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set inbound-route-filter-bgp-policy: %v", err)
    }
    if !setInboundRouteFilterBgpPolicyRes.Success {
        return fmt.Errorf(setInboundRouteFilterBgpPolicyRes.ErrorMsg)
    }

    return readGaiaInboundRouteFilterBgpPolicy(d, m)
}

func deleteGaiaInboundRouteFilterBgpPolicy(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)
    payload := map[string]interface{}{"policy-id": d.Get("policy_id"), "reset": true}
    res, err := client.ApiCallSimple("set-inbound-route-filter-bgp-policy", payload)
    if err != nil {
        return fmt.Errorf("Failed to reset inbound-route-filter-bgp-policy: %v", err)
    }
    if !res.Success {
        return fmt.Errorf(res.ErrorMsg)
    }
    d.SetId("")
    return nil
}

