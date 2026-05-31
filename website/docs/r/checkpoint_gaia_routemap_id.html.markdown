---
layout: "checkpoint"
page_title: "checkpoint_gaia_routemap_id"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-routemap-id"
description: |-
This resource allows you to execute Check Point Routemap Id.
---

# checkpoint_gaia_routemap_id

This resource allows you to execute Check Point Routemap Id.

## Example Usage


```hcl
resource "checkpoint_gaia_routemap_id" "example" {
  resource_id = 1
  name = "bgp_export"
  state = "restrict"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Configures Routemap name. If the Routemap with the name doesn't exist, it will be created.  The Routemap name will be used to configure export and import policy for various dynamic routing protocols. 
* `resource_id` - (Required) Configures Routemap ID for Routemap name.  The Routemap ID has match condition and actions for the import and export policies. 
* `state` - (Optional) Configures state for Routemap ID. Any of the following values are permissible: allow, restrict, inactive.  Allow - If route is matched, it will be accepted by import or export policy. Restrict - If route is matched, it will be rejected by the import or export policy. Inactive - If route is matched, it will not be taken into consideration by the import and export policy. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `ids` - (Computed) Computed field, returned in the response. ids blocks are documented below.


`ids` supports the following:

* `id` - (Computed) Computed field, returned in the response. 
* `state` - (Computed) Computed field, returned in the response. 
* `match` - (Computed) Computed field, returned in the response. match blocks are documented below.
* `action` - (Computed) Computed field, returned in the response. action blocks are documented below.


`match` supports the following:

* `aspath_regex` - (Computed) Computed field, returned in the response. aspath_regex blocks are documented below.
* `community_regex` - (Computed) Computed field, returned in the response. 
* `extcommunity_regex` - (Computed) Computed field, returned in the response. 
* `metric` - (Computed) Computed field, returned in the response. 
* `protocol` - (Computed) Computed field, returned in the response. 
* `as` - (Computed) Computed field, returned in the response. as blocks are documented below.
* `community` - (Computed) Computed field, returned in the response. community blocks are documented below.
* `extended_community` - (Computed) Computed field, returned in the response. extended_community blocks are documented below.
* `ifaddress` - (Computed) Computed field, returned in the response. ifaddress blocks are documented below.
* `interface` - (Computed) Computed field, returned in the response. interface blocks are documented below.
* `level` - (Computed) Computed field, returned in the response. level blocks are documented below.
* `metric_type` - (Computed) Computed field, returned in the response. metric_type blocks are documented below.
* `neighbor` - (Computed) Computed field, returned in the response. neighbor blocks are documented below.
* `network` - (Computed) Computed field, returned in the response. network blocks are documented below.
* `nexthop` - (Computed) Computed field, returned in the response. nexthop blocks are documented below.
* `ospf_instance` - (Computed) Computed field, returned in the response. ospf_instance blocks are documented below.
* `prefix_list` - (Computed) Computed field, returned in the response. prefix_list blocks are documented below.
* `prefix_tree` - (Computed) Computed field, returned in the response. prefix_tree blocks are documented below.
* `route_type` - (Computed) Computed field, returned in the response. route_type blocks are documented below.
* `tag` - (Computed) Computed field, returned in the response. tag blocks are documented below.


`action` supports the following:

* `aspath_prepend_count` - (Computed) Computed field, returned in the response. 
* `localpref` - (Computed) Computed field, returned in the response. 
* `metric` - (Computed) Computed field, returned in the response. metric blocks are documented below.
* `metric_type` - (Computed) Computed field, returned in the response. 
* `ospfautomatictag` - (Computed) Computed field, returned in the response. 
* `ospfmanualtag` - (Computed) Computed field, returned in the response. 
* `precedence` - (Computed) Computed field, returned in the response. 
* `preference` - (Computed) Computed field, returned in the response. 
* `prefix_list` - (Computed) Computed field, returned in the response. 
* `riptag` - (Computed) Computed field, returned in the response. 
* `route_type` - (Computed) Computed field, returned in the response. 
* `community` - (Computed) Computed field, returned in the response. community blocks are documented below.
* `extended_community` - (Computed) Computed field, returned in the response. extended_community blocks are documented below.
* `nexthop` - (Computed) Computed field, returned in the response. nexthop blocks are documented below.


`aspath_regex` supports the following:

* `regex` - (Computed) Computed field, returned in the response. 
* `origin` - (Computed) Computed field, returned in the response. 


`community` supports the following:

* `match_exact` - (Computed) Computed field, returned in the response. 
* `objects` - (Computed) Computed field, returned in the response. objects blocks are documented below.


`extended_community` supports the following:

* `match_type` - (Computed) Computed field, returned in the response. 
* `objects` - (Computed) Computed field, returned in the response. objects blocks are documented below.


`network` supports the following:

* `subnet` - (Computed) Computed field, returned in the response. 
* `match_type` - (Computed) Computed field, returned in the response. 
* `range` - (Computed) Computed field, returned in the response. range blocks are documented below.
* `restrict` - (Computed) Computed field, returned in the response. 


`prefix_list` supports the following:

* `prefix` - (Computed) Computed field, returned in the response. 
* `invert` - (Computed) Computed field, returned in the response. 
* `preference` - (Computed) Computed field, returned in the response. 


`prefix_tree` supports the following:

* `prefix` - (Computed) Computed field, returned in the response. 
* `invert` - (Computed) Computed field, returned in the response. 
* `preference` - (Computed) Computed field, returned in the response. 


`metric` supports the following:

* `value` - (Computed) Computed field, returned in the response. 
* `action_type` - (Computed) Computed field, returned in the response. 


`community` supports the following:

* `action_type` - (Computed) Computed field, returned in the response. 
* `objects` - (Computed) Computed field, returned in the response. objects blocks are documented below.


`extended_community` supports the following:

* `action_type` - (Computed) Computed field, returned in the response. 
* `objects` - (Computed) Computed field, returned in the response. objects blocks are documented below.


`nexthop` supports the following:

* `ipv4` - (Computed) Computed field, returned in the response. 
* `ipv6` - (Computed) Computed field, returned in the response. 


`objects` supports the following:

* `id` - (Computed) Computed field, returned in the response. 
* `as` - (Computed) Computed field, returned in the response. 


`objects` supports the following:

* `type` - (Computed) Computed field, returned in the response. 
* `sub_type` - (Computed) Computed field, returned in the response. 
* `value` - (Computed) Computed field, returned in the response. 


`range` supports the following:

* `from` - (Computed) Computed field, returned in the response. 
* `to` - (Computed) Computed field, returned in the response. 


`objects` supports the following:

* `id` - (Computed) Computed field, returned in the response. 
* `as` - (Computed) Computed field, returned in the response. 


`objects` supports the following:

* `type` - (Computed) Computed field, returned in the response. 
* `sub_type` - (Computed) Computed field, returned in the response. 
* `value` - (Computed) Computed field, returned in the response. 
