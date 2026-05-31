package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBgpExternalPeer() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBgpExternalPeer,
        Read:   readGaiaBgpExternalPeer,
        Update: updateGaiaBgpExternalPeer,
        Delete: deleteGaiaBgpExternalPeer,
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
            "enable_accept_med": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Specifies whether to accept the MED attribute received from this external peer. MEDs are always accepted from internal and confederation peers. If this parameter is set to 'off', the MED is stripped before the update is added to the routing table. If this parameter is reconfigured, the affected peering sessions are automatically restarted.`,
            },
            "allowas_in_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the number of times the Local AS can occur in an AS path received from this peer. A value of 0 means that the Local AS cannot be in the received AS path. The default value is 0.<br><br>If the Peer Local AS feature is enabled, then this value represents the total cumulative occurances of the Local AS and Peer Local AS that can occur in an AS path.`,
            },
            "enable_as_override": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Directs the router to overwrite this peer's AS number with this router's AS in the AS path.<br><br>If the Peer Local AS feature is enabled, this router will use the configured Peer Local AS to override the remote peer's AS number.`,
            },
            "aspath_prepend_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the number of times this router adds its AS to the route's AS path for eBGP (external) or CBGP (Confederation) sessions. The default value is 1.<br><br>If the Peer Local AS feature is enabled, the configured Peer Local AS will be the AS number prepended.`,
            },
            "export_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure export policy for the given BGP peer group or peer.`,
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
                        "conditional_routemap": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Condition to apply to the routemap`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The name of the routemap condition`,
                                    },
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `The condition can be any-pass or no-pass`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "inject_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure conditional route injection for a routemap`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The name of the inject routemap`,
                        },
                        "preference": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Preference for the routemap. Routemaps are evaluated in order of increasing preference value.`,
                        },
                        "any_pass_routemap": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The name of the any-pass-routemap that will be the condition for injection`,
                        },
                        "family": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Describes which family of routes this routemap will be applied to.`,
                        },
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
            "med_out": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specifies the Multi-Exit Discriminator (MED) metric used on BGP routes sent to this peer. This BGP attribute is optional, and if none is specified, then no metric will be propagated to the peer. This metric is overridden by any metric specified in export policy.`,
            },
            "enable_multihop": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Multihop is used to establish peering with External BGP (eBGP) peers that are not directly connected. The router then uses the IGP route table to reach the peer. The feature can be used to perform eBGP load balancing.<br><br>Cannot be configured with IPv6 link-local peers.`,
            },
            "outgoing_interface": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Directs the router to use the interface specified to reach the group peer(s). This is required for IPv6 peers that are identified with a link-local address (an address belonging to the fe80::/64 subnet).`,
            },
            "peer_local_as": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configures a peer-specific Local AS number different to the systemwide Local AS number. The Peer Local AS will replace the Local AS in the BGP session.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "as": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the Peer Local AS number to use when peering with this peer.The value can be one of the following:<br>'off'<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
                        },
                        "enable_dual_peering": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Enabling this option allows the peer to connect to either the Local AS or the Peer Local AS number. When not enabled, only connections to the Peer Local AS number are accepted.<br><br>Only one connection can exist between this system and the peer.<br><br>If peering is established with the Local AS number, the BGP session will behave as if the Peer Local AS feature is not configured.<br><br>This feature should not be used with another system that already has Peer Local AS with Dual-Peering enabled as it is possible for the two systems to alternate sending AS numbers in OPEN messages in a manner that never converges. Cisco and Juniper have similar features named 'dual-as' and 'alias' respectively.`,
                        },
                        "enable_inbound_peer_local": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Specifies that the Peer Local AS number be prepended to the AS path of prefix updates received from the peer.`,
                        },
                        "enable_outbound_local": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Specifies that the Local AS number be prepended to the AS path of prefix updates advertised to the peer. The Local AS number is prepended before the Peer Local AS number.`,
                        },
                    },
                },
            },
            "enable_remove_private_as": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Specifies that private AS number be removed from updates to this peer. The following conditions apply:<br><br>If the AS path includes both public and private AS numbers, no private AS numbers are removed.<br><br>If the AS path contains the AS number of the destination peer, no private AS numbers are removed.<br><br>If the AS path contains only confederations and private AS numbers, private AS numbers are removed.`,
            },
            "remote_as": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The Autonomous System number of the peer group to configure.The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
            },
            "enable_suppress_default_originate": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Eliminates this peer from consideration when generating the BGP default route.`,
            },
            "ttl": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Limits the number of hops over which the eBGP multihop session is established. This feature is used only with multihop. The default value is 64.`,
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
            "peer_type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaBgpExternalPeer(d *schema.ResourceData, m interface{}) error {
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

    if v, ok := d.GetOkExists("enable_accept_med"); ok {
        payload["enable-accept-med"] = v.(bool)
    }

    if v, ok := d.GetOk("allowas_in_count"); ok {
        payload["allowas-in-count"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_as_override"); ok {
        payload["enable-as-override"] = v.(bool)
    }

    if v, ok := d.GetOk("aspath_prepend_count"); ok {
        payload["aspath-prepend-count"] = v.(string)
    }

    if v := d.Get("export_routemap_list"); len(v.([]interface{})) > 0 {
        exportroutemaplistList := v.([]interface{})
        exportroutemaplistArray := make([]interface{}, 0, len(exportroutemaplistList))
        for i := range exportroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if sv, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.conditional_routemap", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    conditional_routemapMap := make(map[string]interface{})
                    if sv, ok := rawDict["name"]; ok && sv.(string) != "" {
                        conditional_routemapMap["name"] = sv.(string)
                    }
                    if sv, ok := rawDict["condition"]; ok && sv.(string) != "" {
                        conditional_routemapMap["condition"] = sv.(string)
                    }
                    if len(conditional_routemapMap) > 0 {
                        itemMap["conditional-routemap"] = conditional_routemapMap
                    }
                }
            }
            if len(itemMap) > 0 {
                exportroutemaplistArray = append(exportroutemaplistArray, itemMap)
            }
        }
        if len(exportroutemaplistArray) > 0 {
            payload["export-routemap-list"] = exportroutemaplistArray
        }
    }

    if v := d.Get("inject_routemap_list"); len(v.([]interface{})) > 0 {
        injectroutemaplistList := v.([]interface{})
        injectroutemaplistArray := make([]interface{}, 0, len(injectroutemaplistList))
        for i := range injectroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.any_pass_routemap", i)); ok {
                itemMap["any-pass-routemap"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                injectroutemaplistArray = append(injectroutemaplistArray, itemMap)
            }
        }
        if len(injectroutemaplistArray) > 0 {
            payload["inject-routemap-list"] = injectroutemaplistArray
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

    if v, ok := d.GetOk("med_out"); ok {
        payload["med-out"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_multihop"); ok {
        payload["enable-multihop"] = v.(bool)
    }

    if v, ok := d.GetOk("outgoing_interface"); ok {
        payload["outgoing-interface"] = v.(string)
    }

    if v := d.Get("peer_local_as"); len(v.([]interface{})) > 0 {
        _ = v
        peerlocalasMap := make(map[string]interface{})
        if v, ok := d.GetOk("peer_local_as.0.as"); ok {
            peerlocalasMap["as"] = v.(string)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_dual_peering"); ok && v.(bool) {
            peerlocalasMap["enable-dual-peering"] = v.(bool)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_inbound_peer_local"); ok && v.(bool) {
            peerlocalasMap["enable-inbound-peer-local"] = v.(bool)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_outbound_local"); ok && v.(bool) {
            peerlocalasMap["enable-outbound-local"] = v.(bool)
        }
        if len(peerlocalasMap) > 0 {
            payload["peer-local-as"] = peerlocalasMap
        }
    }

    if v, ok := d.GetOkExists("enable_remove_private_as"); ok {
        payload["enable-remove-private-as"] = v.(bool)
    }

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_suppress_default_originate"); ok {
        payload["enable-suppress-default-originate"] = v.(bool)
    }

    if v, ok := d.GetOk("ttl"); ok {
        payload["ttl"] = v.(string)
    }

    log.Println("Create BgpExternalPeer - Map = ", payload)

    addBgpExternalPeerRes, err := client.ApiCallSimple("add-bgp-external-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBgpExternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBgpExternalPeerRes.Success {
            errMsg = addBgpExternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBgpExternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-external-peer",        // resource type
            "create",                       // operation
            "add-bgp-external-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add bgp-external-peer: %v", err)
    }
    if !addBgpExternalPeerRes.Success {
        if addBgpExternalPeerRes.ErrorMsg != "" {
            return fmt.Errorf(addBgpExternalPeerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("bgp-external-peer-" + acctest.RandString(10)))
    return readGaiaBgpExternalPeer(d, m)
}

func readGaiaBgpExternalPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    // show-configuration-bgp-external-peer requires peer and remote-as.
    payload["peer"] = d.Get("peer")
    payload["remote-as"] = d.Get("remote_as")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showBgpExternalPeerRes, err := client.ApiCallSimple("show-configuration-bgp-external-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBgpExternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBgpExternalPeerRes.Success {
            errMsg = showBgpExternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBgpExternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-external-peer",        // resource type
            "read",                       // operation
            "show-configuration-bgp-external-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show bgp-external-peer: %v", err)
    }
    if !showBgpExternalPeerRes.Success {
        if data := showBgpExternalPeerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBgpExternalPeerRes.ErrorMsg)
    }

    bgpExternalPeer := showBgpExternalPeerRes.GetData()

    log.Println("Read BgpExternalPeer - Show JSON = ", bgpExternalPeer)

    if v, exists := bgpExternalPeer["accept-routes"]; exists {
        d.Set("accept_routes", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["authtype"]; exists {
        d.Set("authtype", v)
    }
    if v, exists := bgpExternalPeer["capability"]; exists {
        d.Set("capability", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-graceful-restart"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_graceful_restart", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_graceful_restart", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["graceful-restart-stalepath-time"]; exists {
        d.Set("graceful_restart_stalepath_time", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["holdtime"]; exists {
        d.Set("holdtime", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-ignore-first-ashop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ignore_first_ashop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ignore_first_ashop", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["keepalive"]; exists {
        d.Set("keepalive", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["local-address"]; exists {
        d.Set("local_address", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-log-state-transitions"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_state_transitions", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_state_transitions", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["enable-log-warnings"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_log_warnings", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_log_warnings", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["enable-no-aggregator-id"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_no_aggregator_id", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_no_aggregator_id", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["outgoing-interface"]; exists {
        d.Set("outgoing_interface", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-passive-tcp"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_passive_tcp", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_passive_tcp", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["peer"]; exists {
        d.Set("peer", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-ping"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_ping", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_ping", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["enable-route-refresh"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_route_refresh", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_route_refresh", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["enable-send-keepalives"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_send_keepalives", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_send_keepalives", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["throttle-count"]; exists {
        d.Set("throttle_count", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["trace"]; exists {
        d.Set("trace", v.([]interface{}))
    }
    if v, exists := bgpExternalPeer["peer-type"]; exists {
        d.Set("peer_type", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-accept-med"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_accept_med", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_accept_med", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["allowas-in-count"]; exists {
        d.Set("allowas_in_count", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-as-override"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_as_override", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_as_override", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["aspath-prepend-count"]; exists {
        d.Set("aspath_prepend_count", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["export-routemap-list"]; exists {
        d.Set("export_routemap_list", v.([]interface{}))
    }
    if v, exists := bgpExternalPeer["inject-routemap-list"]; exists {
        d.Set("inject_routemap_list", v.([]interface{}))
    }
    if v, exists := bgpExternalPeer["import-routemap-list"]; exists {
        d.Set("import_routemap_list", v.([]interface{}))
    }
    if v, exists := bgpExternalPeer["ip-reachability"]; exists {
        d.Set("ip_reachability", v)
    }
    if v, exists := bgpExternalPeer["med-out"]; exists {
        d.Set("med_out", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["enable-multihop"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_multihop", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_multihop", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["peer-local-as"]; exists {
        d.Set("peer_local_as", v)
    }
    if v, exists := bgpExternalPeer["enable-remove-private-as"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_remove_private_as", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_remove_private_as", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["remote-as"]; exists {
        d.Set("remote_as", fmt.Sprintf("%v", v))
    }
    if v, exists := bgpExternalPeer["send-route-refresh"]; exists {
        d.Set("send_route_refresh", v)
    }
    if v, exists := bgpExternalPeer["enable-suppress-default-originate"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_suppress_default_originate", b)
        } else if s, ok := v.(string); ok {
            d.Set("enable_suppress_default_originate", s == "true")
        }
    }
    if v, exists := bgpExternalPeer["ttl"]; exists {
        d.Set("ttl", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBgpExternalPeer(d *schema.ResourceData, m interface{}) error {

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

    if v, ok := d.GetOkExists("enable_accept_med"); ok {
        payload["enable-accept-med"] = v.(bool)
    }

    if v, ok := d.GetOk("allowas_in_count"); ok {
        payload["allowas-in-count"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_as_override"); ok {
        payload["enable-as-override"] = v.(bool)
    }

    if v, ok := d.GetOk("aspath_prepend_count"); ok {
        payload["aspath-prepend-count"] = v.(string)
    }

    if v := d.Get("export_routemap_list"); len(v.([]interface{})) > 0 {
        exportroutemaplistList := v.([]interface{})
        exportroutemaplistArray := make([]interface{}, 0, len(exportroutemaplistList))
        for i := range exportroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if sv, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.conditional_routemap", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    conditional_routemapMap := make(map[string]interface{})
                    if sv, ok := rawDict["name"]; ok && sv.(string) != "" {
                        conditional_routemapMap["name"] = sv.(string)
                    }
                    if sv, ok := rawDict["condition"]; ok && sv.(string) != "" {
                        conditional_routemapMap["condition"] = sv.(string)
                    }
                    if len(conditional_routemapMap) > 0 {
                        itemMap["conditional-routemap"] = conditional_routemapMap
                    }
                }
            }
            if len(itemMap) > 0 {
                exportroutemaplistArray = append(exportroutemaplistArray, itemMap)
            }
        }
        if len(exportroutemaplistArray) > 0 {
            payload["export-routemap-list"] = exportroutemaplistArray
        }
    }

    if v := d.Get("inject_routemap_list"); len(v.([]interface{})) > 0 {
        injectroutemaplistList := v.([]interface{})
        injectroutemaplistArray := make([]interface{}, 0, len(injectroutemaplistList))
        for i := range injectroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.any_pass_routemap", i)); ok {
                itemMap["any-pass-routemap"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                injectroutemaplistArray = append(injectroutemaplistArray, itemMap)
            }
        }
        if len(injectroutemaplistArray) > 0 {
            payload["inject-routemap-list"] = injectroutemaplistArray
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

    if v, ok := d.GetOk("med_out"); ok {
        payload["med-out"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_multihop"); ok {
        payload["enable-multihop"] = v.(bool)
    }

    if v := d.Get("peer_local_as"); len(v.([]interface{})) > 0 {
        _ = v
        peerlocalasMap := make(map[string]interface{})
        if v, ok := d.GetOk("peer_local_as.0.as"); ok {
            peerlocalasMap["as"] = v.(string)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_dual_peering"); ok && v.(bool) {
            peerlocalasMap["enable-dual-peering"] = v.(bool)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_inbound_peer_local"); ok && v.(bool) {
            peerlocalasMap["enable-inbound-peer-local"] = v.(bool)
        }
        if v, ok := d.GetOkExists("peer_local_as.0.enable_outbound_local"); ok && v.(bool) {
            peerlocalasMap["enable-outbound-local"] = v.(bool)
        }
        if len(peerlocalasMap) > 0 {
            payload["peer-local-as"] = peerlocalasMap
        }
    }

    if v, ok := d.GetOkExists("enable_remove_private_as"); ok {
        payload["enable-remove-private-as"] = v.(bool)
    }

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_suppress_default_originate"); ok {
        payload["enable-suppress-default-originate"] = v.(bool)
    }

    if v, ok := d.GetOk("ttl"); ok {
        payload["ttl"] = v.(string)
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

    setBgpExternalPeerRes, err := client.ApiCallSimple("set-bgp-external-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBgpExternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBgpExternalPeerRes.Success {
            errMsg = setBgpExternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBgpExternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-external-peer",        // resource type
            "update",                       // operation
            "set-bgp-external-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set bgp-external-peer: %v", err)
    }
    if !setBgpExternalPeerRes.Success {
        return fmt.Errorf(setBgpExternalPeerRes.ErrorMsg)
    }

    return readGaiaBgpExternalPeer(d, m)
}

func deleteGaiaBgpExternalPeer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("peer"); ok {
        payload["peer"] = v.(string)
    }

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    deleteBgpExternalPeerRes, err := client.ApiCallSimple("delete-bgp-external-peer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteBgpExternalPeerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteBgpExternalPeerRes.Success {
            errMsg = deleteBgpExternalPeerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteBgpExternalPeerRes.GetData()
        }

        debugLogOperation(
            "bgp-external-peer",        // resource type
            "delete",                       // operation
            "delete-bgp-external-peer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete bgp-external-peer: %v", err)
    }
    if !deleteBgpExternalPeerRes.Success {
        return fmt.Errorf(deleteBgpExternalPeerRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

