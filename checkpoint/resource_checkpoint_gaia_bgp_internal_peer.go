package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBgpInternalPeer() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBgpInternalPeer,
        Read:   readGaiaBgpInternalPeer,
        Update: updateGaiaBgpInternalPeer,
        Delete: deleteGaiaBgpInternalPeer,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "peer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `IP address of the peer.`,
            },
            "accept_routes": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Whether or not to receive routes from the peer in the absence of an inbound route filter.`,
            },
            "authtype": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure authentication policy for this peer.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Authentication type for this peer.`,
                        },
                        "secret": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Secret key. Must be 1-80 characters.`,
                        },
                    },
                },
            },
            "capability": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configure the IP capabilities supported for this session. By default, IPv4 unicast is enabled and IPv6 unicast is disabled`,
            },
            "comment": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Set the comment for this peer.`,
            },
            "enable_graceful_restart": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configures Graceful Restart capability for the given BGP peer.`,
            },
            "graceful_restart_stalepath_time": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the time (seconds) that this router will wait for a restarting BGP peer to send the End-of-RIB notification.`,
            },
            "holdtime": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the holdtime (seconds) to use when negotiating the connection with this peer. The default value is 180s. The holdtime must always be three times the keepalive time. Setting holdtime will automatically set keepalive time to appropriate value`,
            },
            "enable_ignore_first_ashop": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Specifies that the router ignore the first AS number in the AS_PATH for routes learned from this peer.`,
            },
            "keepalive": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `This is an alternative way to specify the holdtime (seconds) when negotiating a peering session. The keepalive interval is one-third the holdtime; both values can be configured, as long as the ratio is maintained. The keepalive must be either 0, i.e., no keepalives are sent, or at least 2. The default value is 60s.`,
            },
            "local_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures the address to be used on the local end of the TCP connection.`,
            },
            "enable_log_state_transitions": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Directs the router to log a message whenever the peer enters or leaves ESTABLISHED state.`,
            },
            "enable_log_warnings": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Directs the router to log a message whenever a warning is encountered in the code path.`,
            },
            "enable_no_aggregator_id": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Directs this router to specify the Router ID in the aggregator attribute as zero, rather than the actual Router ID. This prevents different routers in an AS from creating aggregate routes with different AS paths.`,
            },
            "enable_passive_tcp": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Forces the router to wait for this peer to initiate the BGP session. By default, periodic messages are sent to all configured peers until a session is established. Modifying this option resets the peer connection.`,
            },
            "enable_ping": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enable or disable ping for this peer.`,
            },
            "enable_route_refresh": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enables or disables route refresh for this peer. Route Refresh is used to either re-learn routes from the peer, or to refresh the routing table of the peer without tearing down the BGP session. Both peers must support this capability.`,
            },
            "enable_send_keepalives": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Specifies that the router always send keepalives, even when an update would substitute.`,
            },
            "throttle_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `This option throttles the network traffic when there are many BGP peers by changing the number of updates sent at a time.`,
            },
            "trace": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure tracing for BGP. Initially, the default values for global trace options are used. The valid values that can be used are: keepalive, open, packets, update, all, general, normal, policy, route, task, timer, and cluster.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                    },
                },
            },
            "import_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure import policy for the given BGP peer.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Name of the routemap`,
                        },
                        "preference": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Preference for the routemap. Routemaps are evaluated in order of increasing preference value.`,
                        },
                        "family": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Describes which family of routes this routemap will be applied to.`,
                        },
                    },
                },
            },
            "ip_reachability": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Directs BGP to start BFD (Bidirectional Forwarding Detection) for this peer. Either \"single hop\" or \"multi hop\" BFD can be configured. Either \"single hop\" or \"multi hop\" BFD must be configured in order to use the \"check control plane\" feature.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure either \"single hop\" BFD, \"multi hop\" BFD, or none. The BFD protocol exists in \"single hop\" and \"multi hop\" variants (RFC 5881 and RFC 5883 respectively).  For multi hop BFD to work, the peer must also have multihop enabled, with this machine's local address as the remote peer address and vice versa. Multihop BFD cannot be configured with IPv6 link-local peers.`,
                        },
                        "local_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Configure the multi-hop local address if multi-hop BFD is enabled. The local address must be a local address of this host or VIP in the case of a cluster.`,
                        },
                        "check_control_plane_failure": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `This feature applies when the local node is helping the remote BGP peer undergo a graceful restart. Single hop or multi hop BFD must be enabled in order for this feature to be enabled.`,
                        },
                    },
                },
            },
            "outgoing_interface": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Directs the router to use the interface specified to reach the group peer(s). This is required for IPv6 peers that are identified with a link-local address (an address belonging to the fe80::/64 subnet).`,
            },
            "peer_type": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies if this is a route reflector client.`,
            },
            "enable_suppress_default_originate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Eliminates this peer from consideration when generating the BGP default route.`,
            },
            "weight": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the default weight associated with each route accepted from this peer. This can be overriden by the weight specified in the import policy.`,
            },
            "send_route_refresh": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Route Refresh is used to either re-learn routes from the peer, or to refresh the routing table of the peer without tearing down the BGP session. Both peers must support this capability. This field will not show up in the response if sent in the request.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Trigger either a route update or a request for a route update to be sent to the given peer.`,
                        },
                        "family": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The address family to send the route refresh for.`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaBgpInternalPeer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("accept_routes"); ok {
        payload["accept-routes"] = v.(string)
    }

    if v := d.Get("authtype"); len(v.([]interface{})) > 0 {
        _ = v
        authtypeMap := make(map[string]interface{})
        if v, ok := d.GetOk("authtype.0.type"); ok {
            authtypeMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("authtype.0.secret"); ok {
            authtypeMap["secret"] = v.(string)
        }
        if len(authtypeMap) > 0 {
            payload["authtype"] = authtypeMap
        }
    }

    if v, ok := d.GetOk("capability"); ok {
        payload["capability"] = v.(string)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_graceful_restart"); ok {
        payload["enable-graceful-restart"] = v.(bool)
    }

    if v, ok := d.GetOk("graceful_restart_stalepath_time"); ok {
        payload["graceful-restart-stalepath-time"] = v.(string)
    }

    if v, ok := d.GetOk("holdtime"); ok {
        payload["holdtime"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_ignore_first_ashop"); ok {
        payload["enable-ignore-first-ashop"] = v.(bool)
    }

    if v, ok := d.GetOk("keepalive"); ok {
        payload["keepalive"] = v.(string)
    }

    if v, ok := d.GetOk("local_address"); ok {
        payload["local-address"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_log_state_transitions"); ok {
        payload["enable-log-state-transitions"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_log_warnings"); ok {
        payload["enable-log-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_no_aggregator_id"); ok {
        payload["enable-no-aggregator-id"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_passive_tcp"); ok {
        payload["enable-passive-tcp"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_ping"); ok {
        payload["enable-ping"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_route_refresh"); ok {
        payload["enable-route-refresh"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_send_keepalives"); ok {
        payload["enable-send-keepalives"] = v.(bool)
    }

    if v, ok := d.GetOk("throttle_count"); ok {
        payload["throttle-count"] = v.(string)
    }

    if v := d.Get("trace"); len(v.([]interface{})) > 0 {
        traceList := v.([]interface{})
        traceArray := make([]interface{}, 0, len(traceList))
        for _ = range traceList {
            itemMap := make(map[string]interface{})
            if len(itemMap) > 0 {
                traceArray = append(traceArray, itemMap)
            }
        }
        if len(traceArray) > 0 {
            payload["trace"] = traceArray
        }
    }

    if v := d.Get("import_routemap_list"); len(v.([]interface{})) > 0 {
        importroutemaplistList := v.([]interface{})
        importroutemaplistArray := make([]interface{}, 0, len(importroutemaplistList))
        for i := range importroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                importroutemaplistArray = append(importroutemaplistArray, itemMap)
            }
        }
        if len(importroutemaplistArray) > 0 {
            payload["import-routemap-list"] = importroutemaplistArray
        }
    }

    if v := d.Get("ip_reachability"); len(v.([]interface{})) > 0 {
        _ = v
        ipreachabilityMap := make(map[string]interface{})
        if v, ok := d.GetOk("ip_reachability.0.type"); ok {
            ipreachabilityMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("ip_reachability.0.local_address"); ok {
            ipreachabilityMap["local-address"] = v.(string)
        }
        if v, ok := d.GetOkExists("ip_reachability.0.check_control_plane_failure"); ok && v.(bool) {
            ipreachabilityMap["check-control-plane-failure"] = v.(bool)
        }
        if len(ipreachabilityMap) > 0 {
            payload["ip-reachability"] = ipreachabilityMap
        }
    }

    if v, ok := d.GetOk("outgoing_interface"); ok {
        payload["outgoing-interface"] = v.(string)
    }

    if v, ok := d.GetOk("peer_type"); ok {
        payload["peer-type"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_suppress_default_originate"); ok {
        payload["enable-suppress-default-originate"] = v.(bool)
    }

    if v, ok := d.GetOk("weight"); ok {
        payload["weight"] = v.(string)
    }

    log.Println("Create BgpInternalPeer - Map = ", payload)

    addBgpInternalPeerRes, err := client.ApiCallSimple("add-bgp-internal-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBgpInternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBgpInternalPeerRes.Success {
            errMsg = addBgpInternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBgpInternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-internal-peer",        // resource type
            "create",                       // operation
            "add-bgp-internal-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add bgp-internal-peer: %v", err)
    }
    if !addBgpInternalPeerRes.Success {
        if addBgpInternalPeerRes.ErrorMsg != "" {
            return fmt.Errorf(addBgpInternalPeerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("bgp-internal-peer-" + acctest.RandString(10)))
    return readGaiaBgpInternalPeer(d, m)
}

func readGaiaBgpInternalPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    // show-configuration-bgp-internal-peer requires peer.
    payload["peer"] = d.Get("peer")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showBgpInternalPeerRes, err := client.ApiCallSimple("show-configuration-bgp-internal-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBgpInternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBgpInternalPeerRes.Success {
            errMsg = showBgpInternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBgpInternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-internal-peer",        // resource type
            "read",                       // operation
            "show-configuration-bgp-internal-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show bgp-internal-peer: %v", err)
    }
    if !showBgpInternalPeerRes.Success {
        if data := showBgpInternalPeerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBgpInternalPeerRes.ErrorMsg)
    }

    bgpInternalPeer := showBgpInternalPeerRes.GetData()

    log.Println("Read BgpInternalPeer - Show JSON = ", bgpInternalPeer)

    if v, exists := bgpInternalPeer["accept-routes"]; exists {
        d.Set("accept_routes", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["authtype"]; exists {
        d.Set("authtype", v)
    }
    if v, exists := bgpInternalPeer["capability"]; exists {
        d.Set("capability", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-graceful-restart"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_graceful_restart", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_graceful_restart", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["graceful-restart-stalepath-time"]; exists {
        d.Set("graceful_restart_stalepath_time", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["holdtime"]; exists {
        d.Set("holdtime", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-ignore-first-ashop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ignore_first_ashop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ignore_first_ashop", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["keepalive"]; exists {
        d.Set("keepalive", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-log-state-transitions"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_state_transitions", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_state_transitions", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["enable-log-warnings"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_warnings", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_warnings", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["enable-no-aggregator-id"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_no_aggregator_id", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_no_aggregator_id", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["outgoing-interface"]; exists {
        d.Set("outgoing_interface", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-passive-tcp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_passive_tcp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_passive_tcp", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["peer"]; exists {
        d.Set("peer", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ping", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["enable-route-refresh"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_route_refresh", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_route_refresh", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["enable-send-keepalives"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_send_keepalives", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_send_keepalives", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["throttle-count"]; exists {
        d.Set("throttle_count", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["trace"]; exists {
        d.Set("trace", v.([]interface{}))
    }
    if v, exists := bgpInternalPeer["import-routemap-list"]; exists {
        d.Set("import_routemap_list", v.([]interface{}))
    }
    if v, exists := bgpInternalPeer["ip-reachability"]; exists {
        d.Set("ip_reachability", v)
    }
    if v, exists := bgpInternalPeer["peer-type"]; exists {
        d.Set("peer_type", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["enable-suppress-default-originate"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_suppress_default_originate", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_suppress_default_originate", s == "true")
        }
    }
    if v, exists := bgpInternalPeer["weight"]; exists {
        d.Set("weight", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpInternalPeer["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBgpInternalPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("accept_routes"); ok {
        payload["accept-routes"] = v.(string)
    }

    if v := d.Get("authtype"); len(v.([]interface{})) > 0 {
        _ = v
        authtypeMap := make(map[string]interface{})
        if v, ok := d.GetOk("authtype.0.type"); ok {
            authtypeMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("authtype.0.secret"); ok {
            authtypeMap["secret"] = v.(string)
        }
        if len(authtypeMap) > 0 {
            payload["authtype"] = authtypeMap
        }
    }

    if v, ok := d.GetOk("capability"); ok {
        payload["capability"] = v.(string)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_graceful_restart"); ok {
        payload["enable-graceful-restart"] = v.(bool)
    }

    if v, ok := d.GetOk("graceful_restart_stalepath_time"); ok {
        payload["graceful-restart-stalepath-time"] = v.(string)
    }

    if v, ok := d.GetOk("holdtime"); ok {
        payload["holdtime"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_ignore_first_ashop"); ok {
        payload["enable-ignore-first-ashop"] = v.(bool)
    }

    if v, ok := d.GetOk("keepalive"); ok {
        payload["keepalive"] = v.(string)
    }

    if v, ok := d.GetOk("local_address"); ok {
        payload["local-address"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_log_state_transitions"); ok {
        payload["enable-log-state-transitions"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_log_warnings"); ok {
        payload["enable-log-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_no_aggregator_id"); ok {
        payload["enable-no-aggregator-id"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_passive_tcp"); ok {
        payload["enable-passive-tcp"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_ping"); ok {
        payload["enable-ping"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_route_refresh"); ok {
        payload["enable-route-refresh"] = v.(bool)
    }

    if v, ok := d.GetOkExists("enable_send_keepalives"); ok {
        payload["enable-send-keepalives"] = v.(bool)
    }

    if v, ok := d.GetOk("throttle_count"); ok {
        payload["throttle-count"] = v.(string)
    }

    if v := d.Get("trace"); len(v.([]interface{})) > 0 {
        traceList := v.([]interface{})
        traceArray := make([]interface{}, 0, len(traceList))
        for _ = range traceList {
            itemMap := make(map[string]interface{})
            if len(itemMap) > 0 {
                traceArray = append(traceArray, itemMap)
            }
        }
        if len(traceArray) > 0 {
            payload["trace"] = traceArray
        }
    }

    if v := d.Get("import_routemap_list"); len(v.([]interface{})) > 0 {
        importroutemaplistList := v.([]interface{})
        importroutemaplistArray := make([]interface{}, 0, len(importroutemaplistList))
        for i := range importroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                importroutemaplistArray = append(importroutemaplistArray, itemMap)
            }
        }
        if len(importroutemaplistArray) > 0 {
            payload["import-routemap-list"] = importroutemaplistArray
        }
    }

    if v := d.Get("ip_reachability"); len(v.([]interface{})) > 0 {
        _ = v
        ipreachabilityMap := make(map[string]interface{})
        if v, ok := d.GetOk("ip_reachability.0.type"); ok {
            ipreachabilityMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("ip_reachability.0.local_address"); ok {
            ipreachabilityMap["local-address"] = v.(string)
        }
        if v, ok := d.GetOkExists("ip_reachability.0.check_control_plane_failure"); ok && v.(bool) {
            ipreachabilityMap["check-control-plane-failure"] = v.(bool)
        }
        if len(ipreachabilityMap) > 0 {
            payload["ip-reachability"] = ipreachabilityMap
        }
    }

    if v, ok := d.GetOkExists("enable_suppress_default_originate"); ok {
        payload["enable-suppress-default-originate"] = v.(bool)
    }

    if v, ok := d.GetOk("weight"); ok {
        payload["weight"] = v.(string)
    }

    if v := d.Get("send_route_refresh"); len(v.([]interface{})) > 0 {
        _ = v
        sendrouterefreshMap := make(map[string]interface{})
        if v, ok := d.GetOk("send_route_refresh.0.type"); ok {
            sendrouterefreshMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("send_route_refresh.0.family"); ok {
            sendrouterefreshMap["family"] = v.(string)
        }
        if len(sendrouterefreshMap) > 0 {
            payload["send-route-refresh"] = sendrouterefreshMap
        }
    }

    setBgpInternalPeerRes, err := client.ApiCallSimple("set-bgp-internal-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBgpInternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBgpInternalPeerRes.Success {
            errMsg = setBgpInternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBgpInternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-internal-peer",        // resource type
            "update",                       // operation
            "set-bgp-internal-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set bgp-internal-peer: %v", err)
    }
    if !setBgpInternalPeerRes.Success {
        return fmt.Errorf(setBgpInternalPeerRes.ErrorMsg)
    }

    return readGaiaBgpInternalPeer(d, m)
}

func deleteGaiaBgpInternalPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    deleteBgpInternalPeerRes, err := client.ApiCallSimple("delete-bgp-internal-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteBgpInternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteBgpInternalPeerRes.Success {
            errMsg = deleteBgpInternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteBgpInternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-internal-peer",        // resource type
            "delete",                       // operation
            "delete-bgp-internal-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete bgp-internal-peer: %v", err)
    }
    if !deleteBgpInternalPeerRes.Success {
        return fmt.Errorf(deleteBgpInternalPeerRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

