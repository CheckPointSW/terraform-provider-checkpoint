---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_static_mroute_next_hop_priority"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-static-mroute-next-hop-priority"
description: |-
This resource allows you to execute Check Point Set Static Mroute Next Hop Priority.
---

# checkpoint_gaia_command_set_static_mroute_next_hop_priority

This resource allows you to execute Check Point Set Static Mroute Next Hop Priority.

## Example Usage


```hcl
# Step 1: create the static mroute with the target next-hop gateway
resource "checkpoint_gaia_static_mroute" "mroute_setup" {
  address     = "40.40.40.0"
  mask_length = 24

  next_hop {
    gateway  = "10.10.10.85"
    priority = "default"
  }
}

# Step 2: update the next-hop priority
resource "checkpoint_gaia_command_set_static_mroute_next_hop_priority" "example" {
  address          = "40.40.40.0"
  mask_length      = 24
  next_hop_gateway = "10.10.10.85"
  priority         = 3

  depends_on = [checkpoint_gaia_static_mroute.mroute_setup]
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) Address of the static-mroute to set configuration for. 
* `mask_length` - (Required) Mask length of the static-mroute. 
* `next_hop_gateway` - (Required) Next-hop gateway, must be an IP address. 
* `priority` - (Required) Priority defines which gateway to select as the next hop: the lower the priority, the higher the preference. Can be default or integer from 1 to 8 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

