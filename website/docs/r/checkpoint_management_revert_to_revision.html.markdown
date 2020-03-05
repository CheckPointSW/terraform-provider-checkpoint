---
layout: "checkpoint"
page_title: "checkpoint_management_revert_to_revision"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-revert-to-revision"
description: |-
This resource allows you to execute Check Point Revert To Revision.
---

# checkpoint_management_revert_to_revision

This resource allows you to execute Check Point Revert To Revision.

## Example Usage


```hcl
resource "checkpoint_management_revert_to_revision" "example" {
  to_session = "d49ed10c-649a-476a-8e80-8282eda00e15"
}
```

## Argument Reference

The following arguments are supported:

* `to_session` - (Optional) Session unique identifier. Specify the session  id you would like to revert your database to. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

