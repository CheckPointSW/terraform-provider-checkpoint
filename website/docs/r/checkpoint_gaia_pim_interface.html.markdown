---
layout: "checkpoint"
page_title: "checkpoint_gaia_pim_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-pim-interface"
description: |-
This resource allows you to execute Check Point Pim Interface.
---

# checkpoint_gaia_pim_interface

This resource allows you to execute Check Point Pim Interface.

## Example Usage


```hcl
# Step 1: configure PIM mode
resource "checkpoint_gaia_command_set_pim" "pim_setup" {
  mode = "sparse"
}

# Step 2: add the PIM interface
resource "checkpoint_gaia_pim_interface" "example" {
  name                   = "eth0"
  dr_priority            = "12"
  enable_virtual_address = false

  depends_on = [checkpoint_gaia_command_set_pim.pim_setup]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The interface name. 
* `dr_priority` - (Optional) Used to determine the relative preference when electing a Designated Router (DR). 
* `enable_virtual_address` - (Optional) Configures VRRP mode for the given interface. 
* `neighbor_filter` - (Optional) Configure Neighbor Filter neighbor_filter blocks are documented below.
* `ip_reachability_detection` - (Optional) Configure BFD IP-Reachability Detection 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`neighbor_filter` supports the following:

* `address` - (Optional) The multicast group prefix/mask, in CIDR notation. 


## Import

Use the existing resource's `name` as the import ID.

```shell
terraform import checkpoint_gaia_pim_interface.example 'myname'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

```shell
terraform import checkpoint_gaia_pim_interface.example '{"name":"myname","member_id":"1_1"}'
```

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
