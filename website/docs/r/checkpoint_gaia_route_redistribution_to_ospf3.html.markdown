---
layout: "checkpoint"
page_title: "checkpoint_gaia_route_redistribution_to_ospf3"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-route-redistribution-to-ospf3"
description: |-
This resource allows you to execute Check Point Route Redistribution To Ospf3.
---

# checkpoint_gaia_route_redistribution_to_ospf3

This resource allows you to execute Check Point Route Redistribution To Ospf3.

## Example Usage


```hcl
resource "checkpoint_gaia_route_redistribution_to_ospf3" "example" {
  instance = "3"
  from {
    kernel {
      all_ipv6_routes {
        enable = true
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `from` - (Optional) Configure policy for exporting routes to IPv6 OSPF from blocks are documented below.
* `instance` - (Optional) Configures OSPF3 for specified instance 
* `reset` - (Optional) Removes OSPF3 Route Redistribution configuration 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`from` supports the following:

* `aggregate` - (Optional) Specifies the aggregate route to redistribute into IPv6 OSPF aggregate blocks are documented below.
* `kernel` - (Optional) Redistribution of kernel routes into IPv6 OSPF.  Note: It may be inadvisable in certain cases to redistribute kernel routes into another protocol. Kernel routes usually exist upon startup of routed, before the routing table has settled, when error conditions or bad routes may be present. Use caution when configuring route redistribution from the kernel. kernel blocks are documented below.
* `nat_pool` - (Optional) Redistribution of NAT pools into IPv6 OSPF nat_pool blocks are documented below.
* `ripng` - (Optional) Redistribution of RIPng routes into IPv6 OSPF ripng blocks are documented below.
* `static_route` - (Optional) Redistribution of static routes into IPv6 OSPF static_route blocks are documented below.
* `bgp_as_number` - (Optional) Configures Autonomous System numbers of the BGP group from which to export routes into IPv6 OSPF bgp_as_number blocks are documented below.
* `bgp_as_path` - (Optional) Configures the redistribution of BGP routes, whose AS path matches a given regular expression into IPv6 OSPF bgp_as_path blocks are documented below.
* `interface` - (Optional) Configures the redistribution of all directly connected routes from an interface into IPv6 OSPF interface blocks are documented below.
* `isis` - (Optional) Configures the redistribution of IS-IS routes into IPv6 OSPF isis blocks are documented below.
* `ospf3` - (Optional) Configures the redistribution of IPv6 OSPF routes into IPv6 OSPF ospf3 blocks are documented below.
* `ospf3ase` - (Optional) Configures the redistribution of IPv6 OSPF Autonomous System External routes into IPv6 OSPF ospf3ase blocks are documented below.


`aggregate` supports the following:

* `all_ipv6_routes` - (Optional) Matches all IPv6 aggregate routes all_ipv6_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv6 aggregate routes. The aggregate routes have to be already configured. network blocks are documented below.


`kernel` supports the following:

* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`nat_pool` supports the following:

* `all_ipv6_routes` - (Optional) Matches all IPv4 NAT pools all_ipv6_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv6 NAT pools. The NAT pool has to be already configured. network blocks are documented below.


`ripng` supports the following:

* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`static_route` supports the following:

* `all_ipv6_routes` - (Optional) Matches all IPv4 static route all_ipv6_routes blocks are documented below.
* `default6` - (Optional) Matches the default IPv4 static route default6 blocks are documented below.
* `network` - (Optional) Matches specific IPv6 static routes. The static route has to be already configured. network blocks are documented below.


`bgp_as_number` supports the following:

* `as_number` - (Optional) Configured Autonomous System Number. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.  The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'. 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`bgp_as_path` supports the following:

* `aspath_regex` - (Optional) Configures the redistribution of BGP routes, whose AS path matches the given regular expression.  Valid Values are regular expressions surrounded by double quotes ("). The regular expression can only have digits, a colon (:) and the following special characters:  <table class="table"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table> 
* `origin` - (Optional) Specifies the completeness of the AS path information. Only a single origin should be used with a regular expression.  Any - Matches any routes, regardless of origin. IGP - Route was learned from an interior routing protocol and the AS path is probably complete. EGP - Route was learned from an exterior routing protocol that does not support AS paths and the path is probably incomplete. incomplete - Use when the AS path information is incomplete. 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`interface` supports the following:

* `interface` - (Optional) Specifies the name of the interface 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule  The metric used by OSPF is a cost, representing the overhead required (i.e. due to bandwidth) to reach a destination. Routes with higher OSPF cost are more expensive 


`isis` supports the following:

* `level` - (Optional) Specifies which IS-IS level the route redistribution is applied to 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf3` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf3ase` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `metric` - (Optional) Specifies the metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `metric` - (Optional) Specifies the metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`default6` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `metric` - (Optional) Specifies the metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies IPv6 OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule 
