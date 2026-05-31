---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_bgp"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-bgp"
description: |-
This resource allows you to execute Check Point Set Bgp.
---

# checkpoint_gaia_command_set_bgp

This resource allows you to execute Check Point Set Bgp.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_bgp" "example" {
  cluster_id = "10.1.1.121"
  default_med = "123"
  default_route_gateway = "1.2.3.4"
  enable_communities = false
  enable_ecmp = true
}
```

## Argument Reference

The following arguments are supported:

* `as` - (Optional) Autonomous System Number The value can be one of the following: 'off' An integer from 1-4294967295 A float from 0.1-65535.65535  WARNING: Removing the AS number will result in all BGP configurations and any associated route-redistribution or inbound-route-filter configurations being removed. 
* `cluster_id` - (Optional) Specifies the cluster-id used for route reflection. 
* `enable_communities` - (Optional) This option controls whether or not community information is included in BGP advertisements. This option must be enabled in order to configure the routing policy to filter incoming or outgoing advertisements based on community information. 
* `confederation` - (Optional) Configure BGP Confederation parameters.  A BGP Confederation is a single large AS that is divided into sub-AS's called Routing Domains. The Routing Domains are only visible within the Confederation; to the outside world, the entire Confederation appears as one AS. Confederations improve BGP performance for large AS's by reducing IBGP mesh size. IBGP is used within each Routing Domain, but between them, a modified form of eBGP is used which preserves route metrics and other BGP attributes.  Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents. confederation blocks are documented below.
* `dampening` - (Optional) Weighted-route dampening minimizes the propagation of flapping routes across an internetwork. dampening blocks are documented below.
* `default_med` - (Optional) Defines the Multi-Exit Discriminator metric (MED) used when advertising routes through BGP. If no value is specified, no metric is propagated. Any metrics configured in peer, Route Map, or route redistribution configurations will override the value configured here. 
* `default_route_gateway` - (Optional) A default route is generated if any BGP peer is up. This route has a higher rank than the default configured through static route configuration. If a specific BGP peer should not be considered for generating the default route, it should be explicitly suppressed via the peer-specific 'suppress-default-originate' configuration. 
* `enable_ecmp` - (Optional) Enables or disables ECMP (Equal-Cost Multi-Path) routing for IPv4 BGP routes. 
* `graceful_restart` - (Optional) Configures global settings for BGP Graceful Restart. graceful_restart blocks are documented below.
* `ping` - (Optional) Configures global settings for BGP ping. ping blocks are documented below.
* `routing_domain` - (Optional) Configure Routing Domain parameters.  In Confederation mode, the Routing Domain is the Confederation sub-AS. It acts as an independent AS within the Confederation, but is not visible outside the Confederation. The Routing Domain identifier (RDI) is the equivalent of its AS number.   Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents. routing_domain blocks are documented below.
* `enable_synchronization` - (Optional) Enabling this option directs Internal BGP (IBGP) peers to check for a matching route from IGP protocols before installing a route. 


`confederation` supports the following:

* `aspath_loops_permitted` - (Optional) Specifies the number of times the Local AS can appear in an AS path for routes learned via BGP. Routes with numbers higher than the configured value are rejected. The default value is 1. 
* `identifier` - (Optional) Specifies the identifier for the entire Confederation.  The Confederation identifier is used as the AS number in external BGP sessions, so it must be a globally unique, normally assigned AS number.  Both the Confederation identifer and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.  The value can be one of the following: 'off' An integer from 1-4294967295 A float from 0.1-65535.65535 


`dampening` supports the following:

* `enabled` - (Optional) Enable/disable weighted-route dampening. 
* `keep_history` - (Optional) Specifies the period over which route flapping history for a given route is maintained. The default value is 1800. 
* `max_flap` - (Optional) Specifies the upper limit of the instability accepted. The value must be higher than one plus the suppress-above value. The default value is 16. 
* `reachable_decay` - (Optional) Specifies the length of time (seconds) it takes for the instability metric to reach one half of its current value when the route is reachable. The default value is 300. 
* `reuse_below` - (Optional) Specifies the value of the instability metric at which a suppressed but reachable route becomes unsuppressed. The value must be less than the suppress-above value. The default value is 2. 
* `suppress_above` - (Optional) Specifies the value of the instability metric at which a route is suppressed. While suppressed, the route is neither installed, nor advertised as reachable. The default value is 3. 
* `unreachable_decay` - (Optional) Specifies the rate at which the instability metric is decayed when a route is unreachable. This value must be equal to or greater than the reachable decay. The default value is 900. 


`graceful_restart` supports the following:

* `restart_time` - (Optional) Specifies the time (seconds) that BGP peers of this router should keep the routes advertised to them while this router restarts. The default value is 360. 
* `selection_deferral_time` - (Optional) Specifies the time (seconds) that this router will wait for the End-of-RIB notification from each of its BGP peers after a restart. The default value is 360. 


`ping` supports the following:

* `count` - (Optional) Set the number of failed pings to an individual BGP peer with ping enabled before BGP will drop that peer. This value is common across all peers. The default value is 3. 
* `interval` - (Optional) Set the interval between pings sent to all BGP peers with ping enabled. The default value is 2. 


`routing_domain` supports the following:

* `aspath_loops_permitted` - (Optional) Specifies the number of times the Local AS can appear in an AS path for routes learned via BGP. Routes with numbers higher than the configured value are rejected. The default value is 1. 
* `identifier` - (Optional) In Confederation mode, the Routing Domain identifier (RDI) identifies the Routing Domain, or Confederation sub-AS, to which this router belongs. The RDI is used as the AS number for peers within the Confederation, while the Confederation identifier is used outside the Confederation. The RDI does not have to be globally unique, since it is never used outside the Confederation.  Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.  The value can be one of the following: 'off' An integer from 1-4294967295 A float from 0.1-65535.65535 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

