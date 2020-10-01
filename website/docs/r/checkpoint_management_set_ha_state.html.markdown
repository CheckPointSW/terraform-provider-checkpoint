---
layout: "checkpoint"
page_title: "checkpoint_management_set_ha_state"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-ha-state"
description: |-
This resource allows you to execute Check Point Set HA State.
---

# Resource: checkpoint_management_set_ha_state

This command resource allows you to execute Check Point Set HA State.

## Example Usage


```hcl
resource "checkpoint_management_set_ha_state" "example" {
  new_state = "active"
}
```

## Argument Reference

The following arguments are supported:

* `new_state` - (Required) Domain server new state. 
* `task_id` - (Computed) Asynchronous task unique identifier. 

## How To Use
Make sure this command will be executed in the right execution order.