---
layout: "checkpoint"
page_title: "checkpoint_gaia_inbound_route_filter_ospf3"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-inbound-route-filter-ospf3"
description: |-
This resource allows you to execute Check Point Inbound Route Filter Ospf3.
---

# checkpoint_gaia_inbound_route_filter_ospf3

This resource allows you to execute Check Point Inbound Route Filter Ospf3.

## Example Usage


```hcl
resource "checkpoint_gaia_inbound_route_filter_ospf3" "example" {
  instance = "default"
  rank     = "255"
}
```

## Argument Reference

The following arguments are supported:

* `instance` - (Optional) Configures filtering of IPv6 routes for a specific OSPFv3 instance.  Note: The specified OSPFv3 instance must be configured 
* `restrict_all_ipv6` - (Optional) When the specified value is set to true, the policy rule rejects all matching IPv6 routes, except when there exists a more specific rule, which is set to "accept".  When the specified value is set to false, the policy rule accepts all matching IPv6 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv6 routes 
* `rank` - (Optional) Assigns a rank to all incoming routes matching the filter. Rank is used by the routing system when there are routes from different protocols to the same destination. The route from the protocol with the lowest rank will be used.  Note: This value cannot be specified when rule is set to restrict 
* `route` - (Optional) Configures filtering of imported IPv6 routes for a given policy rule route blocks are documented below.
* `reset` - (Optional) Resets Inbound Route Filter configuration to a default state for the given IPv6 OSPF Instance 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`route` supports the following:

* `subnet` - (Optional) Specifies the address range with which to filter imported IPv6 routes 
* `restrict` - (Optional) When the specified value is true, all routes matching this rule will be rejected, unless a more specific filter accepts the imported routes. When the specified value is false, all routes matching this rule will be accepted, unless a more specific filter accepts them. By default, the given route will be accepted 
* `match_type` - (Optional) Routes can be matched with the following types:   <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>exact</td> <td>Matches only routes with prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>refines</td> <td>Matches only routes that are contained within the specified network (i.e., with greater mask length)</td> </tr></table> 
* `rank` - (Optional) Assigns a rank to all incoming routes matching this filter, except those matching a more specific rule with a different rank configured.  Rank is used by the routing system when there are routes from different protocols to the same destination. The route with the lowest rank from the protocol will be used 
