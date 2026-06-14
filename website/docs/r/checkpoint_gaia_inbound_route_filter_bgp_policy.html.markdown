---
layout: "checkpoint"
page_title: "checkpoint_gaia_inbound_route_filter_bgp_policy"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-inbound-route-filter-bgp-policy"
description: |-
This resource allows you to execute Check Point Inbound Route Filter Bgp Policy.
---

# checkpoint_gaia_inbound_route_filter_bgp_policy

This resource allows you to execute Check Point Inbound Route Filter Bgp Policy.

## Example Usage


```hcl
resource "checkpoint_gaia_inbound_route_filter_bgp_policy" "example" {
  policy_id         = 512
  based_on_as       = "65002"
  restrict_all_ipv4 = false
  default_localpref = "1"
  default_weight    = "10"
  route {
    subnet     = "1.2.3.0/24"
    restrict   = false
    match_type = "normal"
    localpref  = "456"
    weight     = "65535"
  }
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Required) Specifies the BGP import policy identifier.  Note: In order to filter based on Autonomous System path, the policy identifier must be between 1-511. In order to filter based on Autonomous System number, the policy identifier must be between 512-1024 
* `restrict_all_ipv4` - (Optional) When the specified value is set to true, the policy rule rejects all matching IPv4 routes, except when there exists a more specific rule, which is set to "accept". When the specified value is set to false, the policy rule accepts all matching IPv4 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv4 routes 
* `restrict_all_ipv6` - (Optional) When the specified value is set to true, the policy rule rejects all matching IPv6 routes, except when there exists a more specific rule, which is set to "accept". When the specified value is set to false, the policy rule accepts all matching IPv6 routes, except when there exists a more specific rule, which rejects the routes. By default, the rule accepts all IPv6 routes.  Note: The following value can only be specified when IPv6 state is enabled 
* `default_localpref` - (Optional) Assigns a BGP local preference to all routes that match this filter. By default, no local preference value is assigned to the matched routes 
* `default_weight` - (Optional) Assignes a BGP weight to all routes that match this filter. By default, no weight value is assigned to the matched routes 
* `community_match` - (Optional) Matches routes containing a given BGP Community.  Note: A maximum of 25 Communites can be configured on each policy identifier community_match blocks are documented below.
* `extcommunity_match` - (Optional) Matches routes containing a given BGP Extended Community  Note: A maximum of 25 Communities can be configured on each policy identifier extcommunity_match blocks are documented below.
* `based_on_as` - (Optional) Configures a new policy for importing BGP routes from a particular Autonomous System.  Note: In order to configure filtering based on AS, the specified policy identifier must be in between 512-1024. Additionally the ASN cannot be configured in any other policy id 
* `based_on_aspath` - (Optional) Configures a new policy for importing BGP routes whose Autonomous Systems path matches the specified regular expression.  Note: In order to configure filtering based on AS path, the specified policy identifier must be in between 1-511. Additionally the AS path cannot be configured in any other policy id based_on_aspath blocks are documented below.
* `route` - (Optional) Configures filtering of imported routes for a given policy rule route blocks are documented below.
* `reset` - (Optional) Resets Inbound Policy Filter Configuration to a default state for a given policy identifier 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`community_match` supports the following:

* `resource_id` - (Optional) Configures the Community ID value for BGP Communities 
* `as` - (Optional) Configures the Autonomous System number for BGP Communities 


`extcommunity_match` supports the following:

* `type` - (Optional) Configured Type for extended communities 
* `sub_type` - (Optional) Configured Sub-Type for extended communities. Valid sub type values are dependent on the type, the valid values are as follows:  <table class="table"><tr> <th>Type</th> <th>Sub Types</th> </tr><tr> <td>transitive-two-octet-as</td> <td>route-target, route-origin, ospf-domain-id, bgp-data-collect, source-as, l2vpn-id, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>link-bandwidth</td> </tr><tr> <td>transitive-four-octet-as</td> <td>route-target, route-origin, generic, ospf-domain-id, bgp-data-collect, source-as, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>generic</td> </tr><tr> <td>transitive-ipv4-address</td> <td>route-target, route-origin, ospf-domain-id, ospf-route-id, l2vpn-id, vrf-route-import, cisco-vpn-dist</td> </tr></table> 
* `value` - (Optional) Configured Value for extended communities. Valid values are dependent on the type, the valid values are as follows:  <table class="table"><tr> <th>Type</th> <th>Values</th> </tr><tr> <td>transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>transitive-ipv4-address</td> <td>IPv4:0 - 65,535</td> </tr></table> 


`based_on_aspath` supports the following:

* `aspath_regex` - (Optional) Specifies the regular expression, which is used to filter by Autonomous Systems paths. A valid AS path regular expression contains only digits and the following special characters:  <table class="table"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table> 
* `origin` - (Optional) Specifies the completeness of the AS path information. The origin values are defined as follows:  <table class="table"><tr> <th>Origin</th> <th>Description</th> </tr><tr> <td>any</td> <td>Matches any route, regardless of origin</td> </tr><tr> <td>IGP</td> <td>Route was learned from an interior routing protocol, and the AS path is probably complete</td> </tr><tr> <td>EGP</td> <td>Route was learned from an exterior routing protocol that does not support AS paths, and the path is probably complete</td> </tr><tr> <td>incomplete</td> <td>Use when the AS path information is incomplete</td> </tr></table> 


`route` supports the following:

* `subnet` - (Optional) Specifies the address range with which to filter imported IPv4 and IPv6 routes.  Note: In order to configure subnets of type IPv6, the IPv6 state needs to be enabled 
* `restrict` - (Optional) When the specified value is true, all routes matching this rule will be rejected, unless a more specific filter accepts the imported routes. When the specified value is false, all routes matching this rule will be accepted, unless a more specific filter accepts them. By default, the given route will be accepted 
* `match_type` - (Optional) Routes can be matched with the following types:   <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>exact</td> <td>Matches only routes with prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>refines</td> <td>Matches only routes that are contained within the specified network (i.e., with greater mask length)</td> </tr><tr> <td>between</td> <td>Matches any route with prefix equal to the specified network whose mask length falls within a particular range</td> </tr></table>  Note: When the given subnet is of type IPv6, the "between" value cannot be specified 
* `range` - (Optional) Specifies the range with which to match the routes.  This attribute can only be specified when the match type is "between" range blocks are documented below.
* `localpref` - (Optional) Assigns a BGP local preference to all routes matching this filter, unless there exists a more specific rule with a different local preference value.  Note: The following value cannot be specified when the rule is restricted 
* `weight` - (Optional) Assinges a BGP Weight to all routes matching this filter unless there exists a more specific rule with a different weight value.  Note: The following value cannot be specified when the rule is restricted 


`range` supports the following:

* `from` - (Optional) Specifies the lower limit of the range of mask lengths 
* `to` - (Optional) Specifies the upper limit of the range of mask lengths 
