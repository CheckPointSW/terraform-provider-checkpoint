---
layout: "checkpoint"
page_title: "checkpoint_management_delete_repository_package"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-repository-package"
description: |-
This resource allows you to execute Check Point Delete Repository Package.
---

# checkpoint_management_delete_repository_package

This resource allows you to execute Check Point Delete Repository Package.

## Example Usage


```hcl
resource "checkpoint_management_delete_repository_package" "example" {
  name = "my_rep_package"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the software package. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

