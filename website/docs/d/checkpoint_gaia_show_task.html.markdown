---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_task"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-task"
description: |-
This resource allows you to execute Check Point Show Task.
---

# checkpoint_gaia_show_task

This resource allows you to execute Check Point Show Task.

## Example Usage


```hcl
data "checkpoint_gaia_show_task" "example" {
  task_id = ["00000000-0000-0000-0000-000000000000"]
}
```

## Argument Reference

The following arguments are supported:

* `task_id` - (Required) task id to show. expiration default time for task id is 1 day, after this time the task id will not be available task_id blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

