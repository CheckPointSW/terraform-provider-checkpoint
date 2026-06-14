---
layout: "checkpoint"
page_title: "checkpoint_gaia_route_redistribution_to_ospf2"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-route-redistribution-to-ospf2"
description: |-
This resource allows you to execute Check Point Route Redistribution To Ospf2.
---

# checkpoint_gaia_route_redistribution_to_ospf2

This resource allows you to execute Check Point Route Redistribution To Ospf2.

## Example Usage


```hcl
resource "checkpoint_gaia_route_redistribution_to_ospf2" "example" {
  instance = "default"
  from {
    static_route {
      default {
        enable = true
        metric = "10"
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `from` - (Optional) Configure policy for exporting routes to OSPF from blocks are documented below.
* `instance` - (Optional) Configures OSPF2 for the specified instance instance 
* `reset` - (Optional) Removes OSPF2 Route Redistribution configuration 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`from` supports the following:

* `aggregate` - (Optional) Specifies the aggregate route to redistribute into OSPF aggregate blocks are documented below.
* `kernel` - (Optional) Redistribution of kernel routes into OSPF.  Note: It may be inadvisable in certain cases to redistribute kernel routes into another protocol. Kernel routes usually exist upon startup of routed, before the routing table has settled, when error conditions or bad routes may be present. Use caution when configuring route redistribution from the kernel. kernel blocks are documented below.
* `nat_pool` - (Optional) Redistribution of NAT pools into OSPF nat_pool blocks are documented below.
* `rip` - (Optional) Redistribution of RIP routes into OSPF rip blocks are documented below.
* `static_route` - (Optional) Redistribution of static routes into OSPF static_route blocks are documented below.
* `bgp_as_number` - (Optional) Configures Autonomous System numbers of the BGP group from which to export routes into OSPF bgp_as_number blocks are documented below.
* `bgp_as_path` - (Optional) Configures the redistribution of BGP routes, whose AS path matches a given regular expression into OSPF bgp_as_path blocks are documented below.
* `interface` - (Optional) Configures the redistribution of all directly connected routes from an interface into OSPF interface blocks are documented below.
* `isis` - (Optional) Configures the redistribution of IS-IS routes into OSPF2 isis blocks are documented below.
* `ospf2` - (Optional) Configures the redistribution of IPv4 OSPF routes into OSPF ospf2 blocks are documented below.
* `ospf2ase` - (Optional) Configures the redistribution of OSPF Autonomous System External routes into OSPF ospf2ase blocks are documented below.


`aggregate` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 aggregate routes all_ipv4_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv4 aggregate routes. The aggregate routes have to be already configured. network blocks are documented below.


`kernel` supports the following:

* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`nat_pool` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 NAT pools all_ipv4_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv4 NAT pools. The NAT pool has to be already configured. network blocks are documented below.


`rip` supports the following:

* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`static_route` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 static route all_ipv4_routes blocks are documented below.
* `default` - (Optional) Matches the default IPv4 static route default blocks are documented below.
* `network` - (Optional) Matches specific IPv4 static routes. The static route has to be already configured. network blocks are documented below.


`bgp_as_number` supports the following:

* `as_number` - (Optional) Configured Autonomous System Number. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.  The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'. 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.
* `ospf_automatic_tag` - (Optional) Enables or disables the use of an automatically generated OSPF route tag, based on the BGP AS. Tag is attached to external OSPF routes upon export 
* `ospf_automatic_tag_value` - (Optional) This feature allows the user to input an integer to modify the OSPF route tag, automatically generated based on the BGP AS. This route tag is attached to external OSPF routes upon export. OSPF Automatic Tag value has to be be enabled. 
* `ospf_manual_tag_value` - (Optional) Specifies the value to place in the external OSPF route tag field. This configuration overrides any automatic tag configuration 


`bgp_as_path` supports the following:

* `aspath_regex` - (Optional) Configures the redistribution of BGP routes, whose AS path matches the given regular expression.  Valid Values are regular expressions surrounded by double quotes ("). The regular expression can only have digits, a colon (:) and the following special characters:  <table class="table"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table> 
* `origin` - (Optional) Specifies the completeness of the AS path information. Only a single origin should be used with a regular expression.  Any - Matches any routes, regardless of origin. IGP - Route was learned from an interior routing protocol and the AS path is probably complete. EGP - Route was learned from an exterior routing protocol that does not support AS paths and the path is probably incomplete. incomplete - Use when the AS path information is incomplete. 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.
* `ospf_automatic_tag` - (Optional) Enables or disables the use of an automatically generated OSPF route tag, based on the BGP AS. Tag is attached to external OSPF routes upon export 
* `ospf_automatic_tag_value` - (Optional) This feature allows the user to input an integer to modify the OSPF route tag, automatically generated based on the BGP AS. This route tag is attached to external OSPF routes upon export. OSPF Automatic Tag value has to be be enabled. 
* `ospf_manual_tag_value` - (Optional) Specifies the value to place in the external OSPF route tag field. This configuration overrides any automatic tag configuration 


`interface` supports the following:

* `interface` - (Optional) Specifies the name of the interface 
* `metric` - (Optional) Specifies the metric to be added to routes redistributed via this rule  The metric used by OSPF is a cost, representing the overhead required (i.e. due to bandwidth) to reach a destination. Routes with higher OSPF cost are more expensive. 


`isis` supports the following:

* `level` - (Optional) Specifies which IS-IS level the route redistribution is applied to 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf2` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf2ase` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `metric` - (Optional) Specifies the  metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `metric` - (Optional) Specifies the  metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`default` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `metric` - (Optional) Specifies the  metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies OSPF metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the OSPF metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 
