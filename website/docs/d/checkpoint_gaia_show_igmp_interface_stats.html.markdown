---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_igmp_interface_stats"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-igmp-interface-stats"
description: |-
This resource allows you to execute Check Point Show Igmp Interface Stats.
---

# checkpoint_gaia_show_igmp_interface_stats

This resource allows you to execute Check Point Show Igmp Interface Stats.

## Example Usage


```hcl
# Step 1: add a local multicast group to enable IGMP on the interface
resource "checkpoint_gaia_igmp_interface_local_group" "igmp_local_group" {
  interface   = "eth0"
  local_group = "224.6.6.6"
}

# Step 2: retrieve the IGMP interface statistics
data "checkpoint_gaia_show_igmp_interface_stats" "example" {
  interface = "eth0"

  depends_on = [checkpoint_gaia_igmp_interface_local_group.igmp_local_group]
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of IGMP interface 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

