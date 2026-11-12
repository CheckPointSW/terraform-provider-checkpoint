---
layout: "checkpoint"
page_title: "checkpoint_management_ha_full_sync"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-ha-full-sync"
description: |-
This resource allows you to execute Check Point HA Full Sync.
---

# Resource: checkpoint_management_ha_full_sync

This command resource allows you to execute Check Point HA Full Sync.

## Example Usage


```hcl
resource "checkpoint_management_ha_full_sync" "example" {
  name = "standby peer"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Peer name (Multi Domain Server, Domain Server or Security Management Server). 
* `uid` - (Optional) Peer unique identifier (Multi Domain Server, Domain Server or Security Management Server).
* `task_id` - (Computed) Asynchronous task unique identifier. 


## How To Use
Make sure this command will be executed in the right execution order.