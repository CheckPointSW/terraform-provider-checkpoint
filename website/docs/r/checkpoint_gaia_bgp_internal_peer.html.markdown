---
layout: "checkpoint"
page_title: "checkpoint_gaia_bgp_internal_peer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-bgp-internal-peer"
description: |-
This resource allows you to execute Check Point Bgp Internal Peer.
---

# checkpoint_gaia_bgp_internal_peer

This resource allows you to execute Check Point Bgp Internal Peer.

## Example Usage


```hcl
# Step 1: configure a BGP AS number
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  as = "65001"
}

# Step 2: enable the internal peer group
resource "checkpoint_gaia_command_set_bgp_internal" "int_group" {
  enabled = true

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}

# Step 3: add the internal peer
resource "checkpoint_gaia_bgp_internal_peer" "example" {
  peer = "10.1.1.12"

  depends_on = [checkpoint_gaia_command_set_bgp_internal.int_group]
}
```

## Argument Reference

The following arguments are supported:

* `peer` - (Required) IP address of the peer. 
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
* `import_routemap_list` - (Optional) Configure import policy for the given BGP peer. import_routemap_list blocks are documented below.
* `ip_reachability` - (Optional) Directs BGP to start BFD (Bidirectional Forwarding Detection) for this peer. Either "single hop" or "multi hop" BFD can be configured. Either "single hop" or "multi hop" BFD must be configured in order to use the "check control plane" feature. ip_reachability blocks are documented below.
* `outgoing_interface` - (Optional) Directs the router to use the interface specified to reach the group peer(s). This is required for IPv6 peers that are identified with a link-local address (an address belonging to the fe80::/64 subnet). 
* `peer_type` - (Optional) Specifies if this is a route reflector client. 
* `enable_suppress_default_originate` - (Optional) Eliminates this peer from consideration when generating the BGP default route. 
* `weight` - (Optional) Specifies the default weight associated with each route accepted from this peer. This can be overriden by the weight specified in the import policy. 
* `send_route_refresh` - (Optional) Route Refresh is used to either re-learn routes from the peer, or to refresh the routing table of the peer without tearing down the BGP session. Both peers must support this capability. This field will not show up in the response if sent in the request. send_route_refresh blocks are documented below.
* `member_id` - (Computed) No description available. 


`authtype` supports the following:

* `type` - (Optional) Authentication type for this peer. 
* `secret` - (Optional) Secret key. Must be 1-80 characters. 


`trace` supports the following:



`import_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`ip_reachability` supports the following:

* `type` - (Optional) Configure either "single hop" BFD, "multi hop" BFD, or none. The BFD protocol exists in "single hop" and "multi hop" variants (RFC 5881 and RFC 5883 respectively).  For multi hop BFD to work, the peer must also have multihop enabled, with this machine's local address as the remote peer address and vice versa. Multihop BFD cannot be configured with IPv6 link-local peers. 
* `local_address` - (Optional) Configure the multi-hop local address if multi-hop BFD is enabled. The local address must be a local address of this host or VIP in the case of a cluster. 
* `check_control_plane_failure` - (Optional) This feature applies when the local node is helping the remote BGP peer undergo a graceful restart. Single hop or multi hop BFD must be enabled in order for this feature to be enabled. 


`send_route_refresh` supports the following:

* `type` - (Optional) Trigger either a route update or a request for a route update to be sent to the given peer. 
* `family` - (Optional) The address family to send the route refresh for. 
