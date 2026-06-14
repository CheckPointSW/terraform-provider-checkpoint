---
layout: "checkpoint"
page_title: "checkpoint_gaia_ipv6_pim_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-ipv6-pim-interface"
description: |-
This resource allows you to execute Check Point Ipv6 Pim Interface.
---

# checkpoint_gaia_ipv6_pim_interface

This resource allows you to execute Check Point Ipv6 Pim Interface.

## Example Usage


```hcl
# Step 1: configure IPv6 PIM mode
resource "checkpoint_gaia_command_set_ipv6_pim" "ipv6_pim_setup" {
  mode = "sparse"
}

# Step 2: add the IPv6 PIM interface
resource "checkpoint_gaia_ipv6_pim_interface" "example" {
  name                   = "eth0"
  dr_priority            = "12"
  enable_virtual_address = false

  depends_on = [checkpoint_gaia_command_set_ipv6_pim.ipv6_pim_setup]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The interface name. 
* `dr_priority` - (Optional) Used to determine the relative preference when electing a Designated Router (DR). 
* `enable_virtual_address` - (Optional) Configures VRRP mode for the given interface. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `ip_reachability_detection` - (Computed) Computed field, returned in the response. 
* `neighbor_filter` - (Computed) Computed field, returned in the response. neighbor_filter blocks are documented below.


`neighbor_filter` supports the following:

* `address` - (Computed) Computed field, returned in the response. 
