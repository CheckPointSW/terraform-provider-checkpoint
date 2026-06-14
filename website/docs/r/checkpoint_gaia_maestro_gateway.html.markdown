---
layout: "checkpoint"
page_title: "checkpoint_gaia_maestro_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-maestro-gateway"
description: |-
This resource allows you to execute Check Point Maestro Gateway.
---

# checkpoint_gaia_maestro_gateway

This resource allows you to execute Check Point Maestro Gateway.

## Example Usage


```hcl
resource "checkpoint_gaia_maestro_gateway" "example" {
  resource_id = ""
  security_group = 2
  description = "New Gateway Description"
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required) ID of Gateway to modify 
* `description` - (Optional) New Gateway description 
* `security_group` - (Optional) ID of a Security Group. If specified, the Gateway will be assigned to this Security Group,regardless of it's current assignment status. In case you want to unassign Gateway from Security Group, use 0 
* `include_pending_changes` - (Computed) If true, show pending topology. If false, show deployed topology 
