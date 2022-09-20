---
layout: "checkpoint"
page_title: "checkpoint_management_command_import_smart_task"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-import-smart-task"
description: |-
This resource allows you to execute Check Point Import Smart Task.
---

# Resource: checkpoint_management_command_import_smart_task

This resource allows you to execute Check Point Import Smart Task.

## Example Usage


```hcl
resource "checkpoint_management_command_import_smart_task" "example" {
  file_path = "/home/admin/smart-task.txt"
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path to the SmartTask file to be imported. Should be the full file path (example, "/home/admin/exported-smart-task.txt"). 
* `message` - Operation status.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

