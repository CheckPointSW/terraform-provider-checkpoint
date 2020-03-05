---
layout: "checkpoint"
page_title: "checkpoint_management_unlock_administrator"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-unlock-administrator"
description: |-
This resource allows you to execute Check Point Unlock Administrator.
---

# checkpoint_management_unlock_administrator

This resource allows you to execute Check Point Unlock Administrator.

## Example Usage


```hcl
resource "checkpoint_management_unlock_administrator" "example" {
  name = "aa"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

