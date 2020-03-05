---
layout: "checkpoint"
page_title: "checkpoint_management_where_used"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-where-used"
description: |-
This resource allows you to execute Check Point Where Used.
---

# checkpoint_management_where_used

This resource allows you to execute Check Point Where Used.

## Example Usage


```hcl
resource "checkpoint_management_where_used" "example" {
  name = "Host 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `dereference_group_members` - (Optional) Indicates whether to dereference "members" field by details level for every object in reply. 
* `show_membership` - (Optional) Indicates whether to calculate and show "groups" field for every object in reply. 
* `indirect` - (Optional) Search for indirect usage. 
* `indirect_max_depth` - (Optional) Maximum nesting level during indirect usage search. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

