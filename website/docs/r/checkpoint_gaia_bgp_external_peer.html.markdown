---
layout: "checkpoint"
page_title: "checkpoint_gaia_bgp_external_peer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-bgp-external-peer"
description: |-
This resource allows you to execute Check Point Bgp External Peer.
---

# checkpoint_gaia_bgp_external_peer

This resource allows you to execute Check Point Bgp External Peer.

## Example Usage


```hcl
# Step 1: enable the external peer group for the target AS
resource "checkpoint_gaia_command_set_bgp_external" "ext_group" {
  remote_as = "1111.2222"
  enabled   = true
}

# Step 2: add the external peer
resource "checkpoint_gaia_bgp_external_peer" "example" {
  peer      = "10.1.1.22"
  remote_as = "1111.2222"

  depends_on = [checkpoint_gaia_command_set_bgp_external.ext_group]
}
```

## Argument Reference

The following arguments are supported:

* `peer` - (Required) IP address of the peer. 
* `remote_as` - (Required) The Autonomous System number of the peer group to configure.The value can be one of the following: An integer from 1-4294967295 A float from 0.1-65535.65535 
* `accept_routes` - (Optional) Whether or not to receive routes from the peer in the absence of an inbound route filter. 
* `authtype` - (Optional) Configure authentication policy for this peer. authtype blocks are documented below.
* `capability` - (Optional) Configure the IP capabilities supported for this session. By default, IPv4 unicast is enabled and IPv6 unicast is disabled 
* `comment` - (Optional) Set the comment for this peer. 
* `enable_graceful_restart` - (Optional) Configures Graceful Restart capability for the given BGP peer. 
* `graceful_restart_stalepath_time` - (Optional) Specifies the time (seconds) that this router will wait for a restarting BGP peer to send the End-of-RIB notification. 
* `holdtime` - (Optional) Specifies the holdtime (seconds) to use when negotiating the connection with this peer. The default value is 180s. The holdtime must always be three times the keepalive time. Setting holdtime will automatically set keepalive time to appropriate value 
* `enable_ignore_first_ashop` - (Optional) Specifies that the router ignore the first AS number in the AS_PATH for routes learned from this peer. 
* `keepalive` - (Optional) This is an alternative way to specify the holdtime (seconds) when negotiating a peering session. The keepalive interval is one-third the holdtime; both values can be configured, as long as the ratio is maintained. The keepalive must be either 0, i.e., no keepalives are sent, or at least 2. The default value is 60s. 
* `local_address` - (Optional) Configures the address to be used on the local end of the TCP connection. 
* `enable_log_state_transitions` - (Optional) Directs the router to log a message whenever the peer enters or leaves ESTABLISHED state. 
* `enable_log_warnings` - (Optional) Directs the router to log a message whenever a warning is encountered in the code path. 
* `enable_no_aggregator_id` - (Optional) Directs this router to specify the Router ID in the aggregator attribute as zero, rather than the actual Router ID. This prevents different routers in an AS from creating aggregate routes with different AS paths. 
* `enable_passive_tcp` - (Optional) Forces the router to wait for this peer to initiate the BGP session. By default, periodic messages are sent to all configured peers until a session is established. Modifying this option resets the peer connection. 
* `enable_ping` - (Optional) Enable or disable ping for this peer. 
* `enable_route_refresh` - (Optional) Enables or disables route refresh for this peer. Route Refresh is used to either re-learn routes from the peer, or to refresh the routing table of the peer without tearing down the BGP session. Both peers must support this capability. 
* `enable_send_keepalives` - (Optional) Specifies that the router always send keepalives, even when an update would substitute. 
* `throttle_count` - (Optional) This option throttles the network traffic when there are many BGP peers by changing the number of updates sent at a time. 
* `trace` - (Optional) Configure tracing for BGP. Initially, the default values for global trace options are used. The valid values that can be used are: keepalive, open, packets, update, all, general, normal, policy, route, task, timer, and cluster. trace blocks are documented below.
* `enable_accept_med` - (Optional) Specifies whether to accept the MED attribute received from this external peer. MEDs are always accepted from internal and confederation peers. If this parameter is set to 'off', the MED is stripped before the update is added to the routing table. If this parameter is reconfigured, the affected peering sessions are automatically restarted. 
* `allowas_in_count` - (Optional) Specifies the number of times the Local AS can occur in an AS path received from this peer. A value of 0 means that the Local AS cannot be in the received AS path. The default value is 0.  If the Peer Local AS feature is enabled, then this value represents the total cumulative occurances of the Local AS and Peer Local AS that can occur in an AS path. 
* `enable_as_override` - (Optional) Directs the router to overwrite this peer's AS number with this router's AS in the AS path.  If the Peer Local AS feature is enabled, this router will use the configured Peer Local AS to override the remote peer's AS number. 
* `aspath_prepend_count` - (Optional) Specifies the number of times this router adds its AS to the route's AS path for eBGP (external) or CBGP (Confederation) sessions. The default value is 1.  If the Peer Local AS feature is enabled, the configured Peer Local AS will be the AS number prepended. 
* `export_routemap_list` - (Optional) Configure export policy for the given BGP peer group or peer. export_routemap_list blocks are documented below.
* `inject_routemap_list` - (Optional) Configure conditional route injection for a routemap inject_routemap_list blocks are documented below.
* `import_routemap_list` - (Optional) Configure import policy for the given BGP peer. import_routemap_list blocks are documented below.
* `ip_reachability` - (Optional) Directs BGP to start BFD (Bidirectional Forwarding Detection) for this peer. Either "single hop" or "multi hop" BFD can be configured. Either "single hop" or "multi hop" BFD must be configured in order to use the "check control plane" feature. ip_reachability blocks are documented below.
* `med_out` - (Optional) Specifies the Multi-Exit Discriminator (MED) metric used on BGP routes sent to this peer. This BGP attribute is optional, and if none is specified, then no metric will be propagated to the peer. This metric is overridden by any metric specified in export policy. 
* `enable_multihop` - (Optional) Multihop is used to establish peering with External BGP (eBGP) peers that are not directly connected. The router then uses the IGP route table to reach the peer. The feature can be used to perform eBGP load balancing.  Cannot be configured with IPv6 link-local peers. 
* `outgoing_interface` - (Optional) Directs the router to use the interface specified to reach the group peer(s). This is required for IPv6 peers that are identified with a link-local address (an address belonging to the fe80::/64 subnet). 
* `peer_local_as` - (Optional) Configures a peer-specific Local AS number different to the systemwide Local AS number. The Peer Local AS will replace the Local AS in the BGP session. peer_local_as blocks are documented below.
* `enable_remove_private_as` - (Optional) Specifies that private AS number be removed from updates to this peer. The following conditions apply:  If the AS path includes both public and private AS numbers, no private AS numbers are removed.  If the AS path contains the AS number of the destination peer, no private AS numbers are removed.  If the AS path contains only confederations and private AS numbers, private AS numbers are removed. 
* `enable_suppress_default_originate` - (Optional) Eliminates this peer from consideration when generating the BGP default route. 
* `ttl` - (Optional) Limits the number of hops over which the eBGP multihop session is established. This feature is used only with multihop. The default value is 64. 
* `send_route_refresh` - (Optional) Route Refresh is used to either re-learn routes from the peer, or to refresh the routing table of the peer without tearing down the BGP session. Both peers must support this capability. This field will not show up in the response if sent in the request. send_route_refresh blocks are documented below.
* `member_id` - (Computed) No description available. 
* `peer_type` - (Computed) Computed field, returned in the response. 


`authtype` supports the following:

* `type` - (Optional) Authentication type for this peer. 
* `secret` - (Optional) Secret key. Must be 1-80 characters. 


`trace` supports the following:



`export_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 
* `conditional_routemap` - (Optional) Condition to apply to the routemap conditional_routemap blocks are documented below.


`inject_routemap_list` supports the following:

* `name` - (Optional) The name of the inject routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `any_pass_routemap` - (Optional) The name of the any-pass-routemap that will be the condition for injection 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`import_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`ip_reachability` supports the following:

* `type` - (Optional) Configure either "single hop" BFD, "multi hop" BFD, or none. The BFD protocol exists in "single hop" and "multi hop" variants (RFC 5881 and RFC 5883 respectively).  For multi hop BFD to work, the peer must also have multihop enabled, with this machine's local address as the remote peer address and vice versa. Multihop BFD cannot be configured with IPv6 link-local peers. 
* `local_address` - (Optional) Configure the multi-hop local address if multi-hop BFD is enabled. The local address must be a local address of this host or VIP in the case of a cluster. 
* `check_control_plane_failure` - (Optional) This feature applies when the local node is helping the remote BGP peer undergo a graceful restart. Single hop or multi hop BFD must be enabled in order for this feature to be enabled. 


`peer_local_as` supports the following:

* `as` - (Optional) Specifies the Peer Local AS number to use when peering with this peer.The value can be one of the following: 'off' An integer from 1-4294967295 A float from 0.1-65535.65535 
* `enable_dual_peering` - (Optional) Enabling this option allows the peer to connect to either the Local AS or the Peer Local AS number. When not enabled, only connections to the Peer Local AS number are accepted.  Only one connection can exist between this system and the peer.  If peering is established with the Local AS number, the BGP session will behave as if the Peer Local AS feature is not configured.  This feature should not be used with another system that already has Peer Local AS with Dual-Peering enabled as it is possible for the two systems to alternate sending AS numbers in OPEN messages in a manner that never converges. Cisco and Juniper have similar features named 'dual-as' and 'alias' respectively. 
* `enable_inbound_peer_local` - (Optional) Specifies that the Peer Local AS number be prepended to the AS path of prefix updates received from the peer. 
* `enable_outbound_local` - (Optional) Specifies that the Local AS number be prepended to the AS path of prefix updates advertised to the peer. The Local AS number is prepended before the Peer Local AS number. 


`send_route_refresh` supports the following:

* `type` - (Optional) Trigger either a route update or a request for a route update to be sent to the given peer. 
* `family` - (Optional) The address family to send the route refresh for. 


`conditional_routemap` supports the following:

* `name` - (Optional) The name of the routemap condition 
* `condition` - (Optional) The condition can be any-pass or no-pass 
