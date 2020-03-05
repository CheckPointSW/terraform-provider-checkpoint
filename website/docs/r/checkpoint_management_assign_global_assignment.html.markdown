---
layout: "checkpoint"
page_title: "checkpoint_management_assign_global_assignment"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-assign-global-assignment"
description: |-
This resource allows you to execute Check Point Assign Global Assignment.
---

# checkpoint_management_assign_global_assignment

This resource allows you to execute Check Point Assign Global Assignment.

## Example Usage


```hcl
resource "checkpoint_management_assign_global_assignment" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `dependent_domains` - (Optional) N/Adependent_domains blocks are documented below.
* `global_domains` - (Optional) N/Aglobal_domains blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

