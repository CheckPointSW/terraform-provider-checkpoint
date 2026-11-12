---
layout: "checkpoint"
page_title: "checkpoint_management_add_repository_package"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-add-repository-package"
description: |-
This resource allows you to execute Check Point Add Repository Package.
---

# checkpoint_management_add_repository_package

This resource allows you to execute Check Point Add Repository Package.

## Example Usage


```hcl
resource "checkpoint_management_add_repository_package" "example" {
  name = "my_rep_package"
  path = "/home/admin/"
  source = "local"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the repository package. 
* `path` - (Required) The path of the repository package.<br><font color="red">Required only for</font> adding package from local. 
* `source` - (Required) The source of the repository package. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

