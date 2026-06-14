package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBgpConfederationPeer() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBgpConfederationPeer,
        Read:   readGaiaBgpConfederationPeer,
        Update: updateGaiaBgpConfederationPeer,
        Delete: deleteGaiaBgpConfederationPeer,
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
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Specifies the time (seconds) that this router will wait for a restarting BGP peer to send the End-of-RIB notification.`,
            },
            "holdtime": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Specifies the holdtime (seconds) to use when negotiating the connection with this peer. The default value is 180s. The holdtime must always be three times the keepalive time. Setting holdtime will automatically set keepalive time to appropriate value`,
            },
            "enable_ignore_first_ashop": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Specifies that the router ignore the first AS number in the AS_PATH for routes learned from this peer.`,
            },
            "keepalive": {
                Type:        schema.TypeInt,
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
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `This option throttles the network traffic when there are many BGP peers by changing the number of updates sent at a time.`,
            },
            "trace": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure tracing for BGP. Initially, the default values for global trace options are used. The valid values that can be used are: keepalive, open, packets, update, all, general, normal, policy, route, task, timer, and cluster.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "member_as": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Specify the Routing Domain identifier of the Confederation peer group to configure.<br><br>If the peer group specified is the local Routing Domain, it will run IBGP in a full mesh (just as an internal peer group normally would in non-Confederation mode). Otherwise, if an external Routing Domain within the Confederation is specified, the peer group will run a modified version of eBGP, which preserves route metrics and other BGP attributes.<br><br>The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
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
            "weight": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Specifies the default weight associated with each route accepted from this peer. This can be overriden by the weight specified in the import policy.`,
            },
        },
    }
}

func createGaiaBgpConfederationPeer(d *schema.ResourceData, m interface{}) error {
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
        payload["graceful-restart-stalepath-time"] = v.(int)
    }

    if v, ok := d.GetOk("holdtime"); ok {
        payload["holdtime"] = v.(int)
    }

    if v, ok := d.GetOkExists("enable_ignore_first_ashop"); ok {
        payload["enable-ignore-first-ashop"] = v.(bool)
    }

    if v, ok := d.GetOk("keepalive"); ok {
        payload["keepalive"] = v.(int)
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
        payload["throttle-count"] = v.(int)
    }

    if v := d.Get("trace"); len(v.([]interface{})) > 0 {
        payload["trace"] = v.([]interface{})
    }

    if v, ok := d.GetOk("member_as"); ok {
        payload["member-as"] = v.(string)
    }

    if v, ok := d.GetOk("outgoing_interface"); ok {
        payload["outgoing-interface"] = v.(string)
    }

    if v, ok := d.GetOk("peer_type"); ok {
        payload["peer-type"] = v.(string)
    }

    if v, ok := d.GetOk("weight"); ok {
        payload["weight"] = v.(int)
    }

    log.Println("Create BgpConfederationPeer - Map = ", payload)

    addBgpConfederationPeerRes, err := client.ApiCallSimple("add-bgp-confederation-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBgpConfederationPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBgpConfederationPeerRes.Success {
            errMsg = addBgpConfederationPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBgpConfederationPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-confederation-peer",        // resource type
            "create",                       // operation
            "add-bgp-confederation-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add bgp-confederation-peer: %v", err)
    }
    if !addBgpConfederationPeerRes.Success {
        if addBgpConfederationPeerRes.ErrorMsg != "" {
            return fmt.Errorf(addBgpConfederationPeerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("bgp-confederation-peer-" + acctest.RandString(10)))
    return readGaiaBgpConfederationPeer(d, m)
}

func readGaiaBgpConfederationPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    // show-configuration-bgp-confederation-peer requires peer and member-as.
    payload["peer"] = d.Get("peer")
    payload["member-as"] = d.Get("member_as")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showBgpConfederationPeerRes, err := client.ApiCallSimple("show-configuration-bgp-confederation-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBgpConfederationPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBgpConfederationPeerRes.Success {
            errMsg = showBgpConfederationPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBgpConfederationPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-confederation-peer",        // resource type
            "read",                       // operation
            "show-configuration-bgp-confederation-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show bgp-confederation-peer: %v", err)
    }
    if !showBgpConfederationPeerRes.Success {
        if data := showBgpConfederationPeerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBgpConfederationPeerRes.ErrorMsg)
    }

    bgpConfederationPeer := showBgpConfederationPeerRes.GetData()

    log.Println("Read BgpConfederationPeer - Show JSON = ", bgpConfederationPeer)

    if v, exists := bgpConfederationPeer["accept-routes"]; exists {
        d.Set("accept_routes", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["authtype"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("authtype", []interface{}{map[string]interface{}{
                "type":   func() string { if _v, _ok := _m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "secret": func() string { if _v, _ok := _m["secret"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := bgpConfederationPeer["capability"]; exists {
        d.Set("capability", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["enable-graceful-restart"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_graceful_restart", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_graceful_restart", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["graceful-restart-stalepath-time"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("graceful_restart_stalepath_time", int(f))
        }
    }
    if v, exists := bgpConfederationPeer["holdtime"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("holdtime", int(f))
        }
    }
    if v, exists := bgpConfederationPeer["enable-ignore-first-ashop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ignore_first_ashop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ignore_first_ashop", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["keepalive"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("keepalive", int(f))
        }
    }
    if v, exists := bgpConfederationPeer["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["enable-log-state-transitions"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_state_transitions", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_state_transitions", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["enable-log-warnings"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_warnings", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_warnings", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["enable-no-aggregator-id"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_no_aggregator_id", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_no_aggregator_id", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["outgoing-interface"]; exists {
        d.Set("outgoing_interface", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["enable-passive-tcp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_passive_tcp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_passive_tcp", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["peer"]; exists {
        d.Set("peer", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["enable-ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ping", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["enable-route-refresh"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_route_refresh", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_route_refresh", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["enable-send-keepalives"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_send_keepalives", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_send_keepalives", s == "true")
        }
    }
    if v, exists := bgpConfederationPeer["throttle-count"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("throttle_count", int(f))
        }
    }
    if v, exists := bgpConfederationPeer["trace"]; exists {
        d.Set("trace", v.([]interface{}))
    }
    if v, exists := bgpConfederationPeer["member-as"]; exists {
        d.Set("member_as", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["peer-type"]; exists {
        d.Set("peer_type", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpConfederationPeer["weight"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("weight", int(f))
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBgpConfederationPeer(d *schema.ResourceData, m interface{}) error {

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
        payload["graceful-restart-stalepath-time"] = v.(int)
    }

    if v, ok := d.GetOk("holdtime"); ok {
        payload["holdtime"] = v.(int)
    }

    if v, ok := d.GetOkExists("enable_ignore_first_ashop"); ok {
        payload["enable-ignore-first-ashop"] = v.(bool)
    }

    if v, ok := d.GetOk("keepalive"); ok {
        payload["keepalive"] = v.(int)
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
        payload["throttle-count"] = v.(int)
    }

    if v := d.Get("trace"); len(v.([]interface{})) > 0 {
        payload["trace"] = v.([]interface{})
    }

    if v, ok := d.GetOk("member_as"); ok {
        payload["member-as"] = v.(string)
    }

    if v, ok := d.GetOk("weight"); ok {
        payload["weight"] = v.(int)
    }

    setBgpConfederationPeerRes, err := client.ApiCallSimple("set-bgp-confederation-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBgpConfederationPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBgpConfederationPeerRes.Success {
            errMsg = setBgpConfederationPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBgpConfederationPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-confederation-peer",        // resource type
            "update",                       // operation
            "set-bgp-confederation-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set bgp-confederation-peer: %v", err)
    }
    if !setBgpConfederationPeerRes.Success {
        return fmt.Errorf(setBgpConfederationPeerRes.ErrorMsg)
    }

    return readGaiaBgpConfederationPeer(d, m)
}

func deleteGaiaBgpConfederationPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("member_as"); ok {
        payload["member-as"] = v.(string)
    }

    deleteBgpConfederationPeerRes, err := client.ApiCallSimple("delete-bgp-confederation-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteBgpConfederationPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteBgpConfederationPeerRes.Success {
            errMsg = deleteBgpConfederationPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteBgpConfederationPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-confederation-peer",        // resource type
            "delete",                       // operation
            "delete-bgp-confederation-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete bgp-confederation-peer: %v", err)
    }
    if !deleteBgpConfederationPeerRes.Success {
        return fmt.Errorf(deleteBgpConfederationPeerRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

