---
layout: "checkpoint"
page_title: "checkpoint_gaia_igmp_interface_static_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-igmp-interface-static-group"
description: |-
This resource allows you to execute Check Point Igmp Interface Static Group.
---

# checkpoint_gaia_igmp_interface_static_group

This resource allows you to execute Check Point Igmp Interface Static Group.

## Example Usage


```hcl
resource "checkpoint_gaia_igmp_interface_static_group" "example" {
  interface   = "eth0"
  static_group = "224.5.5.5"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the IGMP interface 
* `static_group` - (Required) The statically configured group address that this IGMP interface receives multicast data for 
* `group_count` - (Optional) The number of adjacent static groups 
* `group_increment` - (Optional) The increment between IGMP static groups (default: 0.0.0.1) 
* `sources` - (Optional) The list of IPv4 sources from which to receive traffic for this static group sources blocks are documented below.
* `source_all_off` - (Optional) Remove all sources of a static group 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`sources` supports the following:

* `source` - (Optional) The IPv4 source from which to receive traffic for this static group 
* `source_count` - (Optional) The number of adjacent static group sources 
* `source_increment` - (Optional) The increment between IGMP static group sources (default: 0.0.0.1) 
