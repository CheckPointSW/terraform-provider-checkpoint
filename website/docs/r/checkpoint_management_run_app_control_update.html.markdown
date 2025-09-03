---
layout: "checkpoint"
page_title: "checkpoint_management_run_app_control_update"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-run-app-control-update"
description: |-
This resource allows you to execute Check Point Run App Control Update.
---

# checkpoint_management_run_app_control_update

This resource allows you to execute Check Point Run App Control Update.

## Example Usage

```hcl
resource "checkpoint_management_run_app_control_update" "example" {

}
```

## Argument Reference

The following arguments are supported:
* `task_id` - Asynchronous task unique identifier. Use show-task command to check the progress of the task.



## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

