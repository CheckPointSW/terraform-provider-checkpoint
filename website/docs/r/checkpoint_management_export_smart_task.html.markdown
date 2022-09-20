---
layout: "checkpoint"
page_title: "checkpoint_management_command_export_smart_task"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-export-smart-task"
description: |-
This resource allows you to execute Check Point Export Smart Task.
---

# Resource: checkpoint_management_command_export_smart_task

This resource allows you to execute Check Point Export Smart Task.

## Example Usage


```hcl
resource "checkpoint_management_command_export_smart_task" "example" {
  name = "After Session Approve"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of task to be exported. 
* `file_path` - (Optional) Path to the SmartTask file to be exported. Should be the full file path (example, "/home/admin/exported-smart-task.txt)". If no path was inserted the default will be: "/var/log/<task_name>.txt". 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

