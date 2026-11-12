---
layout: "checkpoint"
page_title: "checkpoint_management_command_abort_get_interfaces"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-abort-get-interfaces"
description: |-
This resource allows you to execute Check Point Abort Get Interfaces.
---

# Resource: checkpoint_management_command_abort_get_interfaces

This resource allows you to execute Check Point Abort Get Interfaces.

## Example Usage


```hcl
resource "checkpoint_management_command_get_interfaces" "get_interfaces" {
  target_uid = "2220d9ad-a251-5555-9a0a-4772a6511111"
}

resource "checkpoint_management_command_abort_get_interfaces" "abort_get_interfaces" {
  task_id = "${checkpoint_management_command_get_interfaces.get_interfaces.task_id}"
}
```

## Argument Reference

The following arguments are supported:

* `task_id` - (Required) get-interfaces task UID.
* `force_cleanup` - (Optional) Forcefully abort the "get-interfaces" task.
* `message` - Operation status.