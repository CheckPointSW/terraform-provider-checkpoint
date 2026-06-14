---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_pbr_table_static_next_hop_priority"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-pbr-table-static-next-hop-priority"
description: |-
This resource allows you to execute Check Point Set Pbr Table Static Next Hop Priority.
---

# checkpoint_gaia_command_set_pbr_table_static_next_hop_priority

This resource allows you to execute Check Point Set Pbr Table Static Next Hop Priority.

## Example Usage


```hcl
# Step 1: create the PBR table with the target static route and next-hop gateway
resource "checkpoint_gaia_pbr_table" "pbr_setup" {
  table = "mytable"
  static_routes {
    address     = "1.2.3.0"
    mask_length = 24
    type        = "gateway"
    next_hop {
      gateway  = "1.1.1.1"
      priority = "1"
    }
  }
}

# Step 2: update the next-hop priority
resource "checkpoint_gaia_command_set_pbr_table_static_next_hop_priority" "example" {
  table              = "mytable"
  static_address     = "1.2.3.0"
  static_mask_length = 24
  next_hop_gateway   = "1.1.1.1"
  priority           = "3"

  depends_on = [checkpoint_gaia_pbr_table.pbr_setup]
}
```

## Argument Reference

The following arguments are supported:

* `table` - (Required) Name of PBR Table 
* `static_address` - (Required) IP address of PBR Table static route 
* `static_mask_length` - (Required) Mask length of PBR Table static route 
* `next_hop_gateway` - (Required) Nexthop gateway of PBR Table static route, can be IP address or interface name 
* `priority` - (Required) This value will replace the current priority of the specified nexthop gateway. Priority defines which gateway to select as the next hop, the lower the priority, the higher the preference. Can be default or integer from 1 to 8 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

