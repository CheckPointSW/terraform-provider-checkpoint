---
layout: "checkpoint"
page_title: "checkpoint_management_delete_updatable_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-updatable-object"
description: |-
This resource allows you to execute Check Point Delete Updatable Object.
---

# checkpoint_management_delete_updatable_object

This resource allows you to execute Check Point Delete Updatable Object.

## Example Usage


```hcl
resource "checkpoint_management_delete_updatable_object" "example" {
  name = "CodeBuild US East 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

