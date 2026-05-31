---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_static_route_next_hop_priority"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-static-route-next-hop-priority"
description: |-
This resource allows you to execute Check Point Set Static Route Next Hop Priority.
---

# checkpoint_gaia_command_set_static_route_next_hop_priority

This resource allows you to execute Check Point Set Static Route Next Hop Priority.

## Example Usage


```hcl
# Step 1: create the static route with the target next-hop gateway
resource "checkpoint_gaia_static_route" "route_setup" {
  address     = "1.2.3.0"
  mask_length = 24
  type        = "gateway"

  next_hop {
    gateway  = "1.1.1.1"
    priority = "default"
  }
}

# Step 2: update the next-hop priority
resource "checkpoint_gaia_command_set_static_route_next_hop_priority" "example" {
  address          = "1.2.3.0"
  mask_length      = 24
  next_hop_gateway = "1.1.1.1"
  priority         = 3

  depends_on = [checkpoint_gaia_static_route.route_setup]
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required)  
* `mask_length` - (Required)  
* `next_hop_gateway` - (Required) nexthop gateway, can be IP address or interface name 
* `priority` - (Required) Priority defines which gateway to select as the next hop, the lower the priority, the higher the preference. can be default or integer from 1 to 8 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

