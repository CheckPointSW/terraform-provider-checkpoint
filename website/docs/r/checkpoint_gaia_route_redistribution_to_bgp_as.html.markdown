---
layout: "checkpoint"
page_title: "checkpoint_gaia_route_redistribution_to_bgp_as"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-route-redistribution-to-bgp-as"
description: |-
This resource allows you to execute Check Point Route Redistribution To Bgp As.
---

# checkpoint_gaia_route_redistribution_to_bgp_as

This resource allows you to execute Check Point Route Redistribution To Bgp As.

## Example Usage


```hcl
# Step 1: clear any leftover BGP confederation state
resource "checkpoint_gaia_command_set_bgp" "clear_conf" {
  confederation {
    identifier = "off"
  }
  routing_domain {
    identifier = "off"
  }
}

# Step 2: configure a local BGP AS number
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  as = "65001"

  depends_on = [checkpoint_gaia_command_set_bgp.clear_conf]
}

# Step 3: configure the external peer group for AS 65002
resource "checkpoint_gaia_command_set_bgp_external" "bgp_ext_setup" {
  remote_as = "65002"
  enabled   = true

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}

# Step 4: configure route redistribution to BGP AS 65002
resource "checkpoint_gaia_route_redistribution_to_bgp_as" "example" {
  as_number = "65002"
  localpref = "300"
  med = "2"
  extcommunity_match {
    type = "transitive-two-octet-as"
    sub_type = "route-target"
    value = "1:1"
  }
  extcommunity_match {
    type = "transitive-ipv4-address"
    sub_type = "ospf-route-id"
    value = "1.2.3.4:5"
  }

  depends_on = [checkpoint_gaia_command_set_bgp_external.bgp_ext_setup]
}
```

## Argument Reference

The following arguments are supported:

* `as_number` - (Required) Specifies the Autonomous System Number for the BGP Route Redistribution. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.  The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'. 
* `from` - (Optional) Configure policy for exporting routes to a BGP peer AS from blocks are documented below.
* `localpref` - (Optional) Configures a Local Preference value 
* `med` - (Optional) Configures a Multi-Exit Descriminator for export routes 
* `community_append` - (Optional) Appends BGP Community to exported routes community_append blocks are documented below.
* `community_match` - (Optional) Configures match value for BGP Community community_match blocks are documented below.
* `extcommunity_append` - (Optional) Appends BGP Extended Community to exported routes extcommunity_append blocks are documented below.
* `extcommunity_match` - (Optional) Configures match value for BGP Extended Community extcommunity_match blocks are documented below.
* `reset` - (Optional) Removes Route Redistribution configuration for the specified configured ASN 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`from` supports the following:

* `aggregate` - (Optional) Specifies the aggregate route to redistribute into BGP aggregate blocks are documented below.
* `default_origin` - (Optional) Default rule for redistributing all IPv4 route to the given BGP AS default_origin blocks are documented below.
* `kernel` - (Optional) Redistribution of kernel routes to the given BGP AS.  Note: It may be inadvisable in certain cases to redistribute kernel routes into another protocol. Kernel routes usually exist upon startup of routed, before the routing table has settled, when error conditions or bad routes may be present. Use caution when configuring route redistribution from the kernel. kernel blocks are documented below.
* `nat_pool` - (Optional) Redistribution of NAT pools to the given BGP AS nat_pool blocks are documented below.
* `rip` - (Optional) Redistribution of RIP routes to the given BGP AS rip blocks are documented below.
* `ripng` - (Optional) Redistribution of RIPng routes to the given BGP AS.  Note: IPv6 state needs to be enabled. ripng blocks are documented below.
* `static_route` - (Optional) Redistribution of static routes to the given BGP AS static_route blocks are documented below.
* `bgp_as_number` - (Optional) Configures Autonomous System numbers of the BGP group from which to export routes to the given BGP AS bgp_as_number blocks are documented below.
* `bgp_as_path` - (Optional) Configures the redistribution of BGP routes, whose AS path matches a given regular expression into BGP bgp_as_path blocks are documented below.
* `interface` - (Optional) Configures the redistribution of all directly connected routes from an interface to the give BGP AS interface blocks are documented below.
* `isis` - (Optional) Configures the redistribution of IS-IS routes into BGP-AS isis blocks are documented below.
* `ospf2` - (Optional) Configures the redistribution of IPv4 OSPF routes to the given BGP AS ospf2 blocks are documented below.
* `ospf2ase` - (Optional) Configures the redistribution of OSPF Autonomous System External routes to the given BGP AS ospf2ase blocks are documented below.
* `ospf3` - (Optional) Configures the redistribution of IPv6 OSPF routes to the given BGP AS.  Note: IPv6 state needs to be enabled. ospf3 blocks are documented below.
* `ospf3ase` - (Optional) Configures the redistribution of IPv6 OSPF Autonomous System External routes to the given BGP AS.  Note: IPv6 state needs to be enabled. ospf3ase blocks are documented below.


`community_append` supports the following:

* `resource_id` - (Optional) Configures the Community ID value for BGP Communities. 
* `as` - (Optional) Configures the Autonomous System number for BGP Communities. 


`community_match` supports the following:

* `resource_id` - (Optional) Configures the Community ID value for BGP Communities. 
* `as` - (Optional) Configures the Autonomous System number for BGP Communities. 


`extcommunity_append` supports the following:

* `type` - (Optional) Configured Type for extended communities. 
* `sub_type` - (Optional) Configured Sub-Type for extended communities. Valid sub type values are dependent on the type, the valid values are as follows:   <table class="table"><tr> <th>Type</th> <th>Sub Types</th> </tr><tr> <td>transitive-two-octet-as</td> <td>route-target, route-origin, ospf-domain-id, bgp-data-collect, source-as, l2vpn-id, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>link-bandwidth</td> </tr><tr> <td>transitive-four-octet-as</td> <td>route-target, route-origin, generic, ospf-domain-id, bgp-data-collect, source-as, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>generic</td> </tr><tr> <td>transitive-ipv4-address</td> <td>route-target, route-origin, ospf-domain-id, ospf-route-id, l2vpn-id, vrf-route-import, cisco-vpn-dist</td> </tr></table> 
* `value` - (Optional) Configured Value for extended communities. Valid values are dependent on the type, the valid values are as follows:   <table class="table"><tr> <th>Type</th> <th>Values</th> </tr><tr> <td>transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>transitive-ipv4-address</td> <td>IPv4:0 - 65,535</td> </tr></table> 


`extcommunity_match` supports the following:

* `type` - (Optional) Configured Type for extended communities. 
* `sub_type` - (Optional) Configured Sub-Type for extended communities. Valid sub type values are dependent on the type, the valid values are as follows:   <table class="table"><tr> <th>Type</th> <th>Sub Types</th> </tr><tr> <td>transitive-two-octet-as</td> <td>route-target, route-origin, ospf-domain-id, bgp-data-collect, source-as, l2vpn-id, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>link-bandwidth</td> </tr><tr> <td>transitive-four-octet-as</td> <td>route-target, route-origin, generic, ospf-domain-id, bgp-data-collect, source-as, cisco-vpn-dist</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>generic</td> </tr><tr> <td>transitive-ipv4-address</td> <td>route-target, route-origin, ospf-domain-id, ospf-route-id, l2vpn-id, vrf-route-import, cisco-vpn-dist</td> </tr></table> 
* `value` - (Optional) Configured Value for extended communities. Valid values are dependent on the type, the valid values are as follows:   <table class="table"><tr> <th>Type</th> <th>Values</th> </tr><tr> <td>transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>non-transitive-two-octet-as</td> <td>1 - 65,535:0 - 4,294,967,295</td> </tr><tr> <td>transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>non-transitive-four-octet-as</td> <td>65,536 - 4,294,967,295:0 - 65,535</td> </tr><tr> <td>transitive-ipv4-address</td> <td>IPv4:0 - 65,535</td> </tr></table> 


`aggregate` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 aggregate routes all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Matches all IPv6 aggregate routes  Note: IPv6 state must be enabled all_ipv6_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv4 or IPv6 aggregate routes. The aggregate routes have to be already configured.  Note: IPv6 state must be enabled for IPv6 aggregate routes. network blocks are documented below.


`default_origin` supports the following:

* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.


`kernel` supports the following:

* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies  Note: IPv6 state must be enabled all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`nat_pool` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 NAT pools all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Matches all IPv6 NAT pools all_ipv6_routes blocks are documented below.
* `network` - (Optional) Matches specific IPv4 or IPv6 NAT pools. The NAT pool has to be already configured.  Note: IPv6 state must be enabled for IPv6 NAT pools. network blocks are documented below.


`rip` supports the following:

* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ripng` supports the following:

* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`static_route` supports the following:

* `all_ipv4_routes` - (Optional) Matches all IPv4 static route all_ipv4_routes blocks are documented below.
* `default` - (Optional) Matches the default IPv4 static route default blocks are documented below.
* `all_ipv6_routes` - (Optional) Matches all IPv6 static route all_ipv6_routes blocks are documented below.
* `default6` - (Optional) Matches the default IPv6 static route default6 blocks are documented below.
* `network` - (Optional) Matches specific IPv4 or IPv6 static routes. The static route has to be already configured.  Note: IPv6 state must be enabled for IPv6 static routes. network blocks are documented below.


`bgp_as_number` supports the following:

* `as_number` - (Optional) Configured Autonomous System Number. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.  The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'. 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies  Note: IPv6 state must be enabled all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`bgp_as_path` supports the following:

* `aspath_regex` - (Optional) Configures the redistribution of BGP routes, whose AS path matches the given regular expression.  Valid Values are regular expressions surrounded by double quotes ("). The regular expression can only have digits, a colon (:) and the following special characters:  <table class="table"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table> 
* `origin` - (Optional) Specifies the completeness of the AS path information. Only a single origin should be used with a regular expression.  Any - Matches any routes, regardless of origin. IGP - Route was learned from an interior routing protocol and the AS path is probably complete. EGP - Route was learned from an exterior routing protocol that does not support AS paths and the path is probably incomplete. incomplete - Use when the AS path information is incomplete. 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies  Note: IPv6 state must be enabled all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`interface` supports the following:

* `interface` - (Optional) Specifies the name of the interface 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule  The metric in BGP is the Multi-Exit Discriminator (MED), used to break ties between routes with equal preference from the same neighboring Autonomous System. Lower MED values are preferred, and routes with no MED tie with a MED value of 0 for most preferred. 


`isis` supports the following:

* `level` - (Optional) Specifies which IS-IS level the route redistribution is applied to 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies  Note: IPv6 state must be enabled all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf2` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf2ase` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv4_routes` - (Optional) Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies all_ipv4_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf3` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`ospf3ase` supports the following:

* `instance` - (Optional) Redistribute routes from a specific OSPF instance 
* `all_ipv6_routes` - (Optional) Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies all_ipv6_routes blocks are documented below.
* `network` - (Optional) Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.  Note: When network objects are specified, previous objects will be overwritten network blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" and network must be IPv4 range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`default` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`default6` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" and network must be IPv4 range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" and network must be IPv4 range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 or IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" and network must be IPv4 range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv4_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv4 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 
* `range` - (Optional) Specifies the mask length range  Note: The match-type needs to be of type "range" range blocks are documented below.


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


`all_ipv6_routes` supports the following:

* `metric` - (Optional) Specifies BGP metric value to routes matching this rule 
* `enable` - (Optional) Enables or disables the metric value 


`network` supports the following:

* `address` - (Optional) Specifies IPv6 network 
* `restrict` - (Optional) Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted 
* `match_type` - (Optional) Defines how routes are matched to the network. The match types are as follows:  <table class="table"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table> 
* `metric` - (Optional) Specifies the BGP metric to be added to routes redistributed via this rule 


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
