---
layout: "checkpoint"
page_title: "checkpoint_management_repository_package"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-repository-package"
description: |-
Use this data source to get information on an existing Check Point Repository Package.
---

# Data Source: checkpoint_management_repository_package

Use this data source to get information on an existing Check Point Repository Package.

## Example Usage


```hcl
resource "checkpoint_management_add_repository_package" "example" {
  name = "my_rep_package"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the repository package. 

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

