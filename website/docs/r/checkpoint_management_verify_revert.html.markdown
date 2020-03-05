---
layout: "checkpoint"
page_title: "checkpoint_management_verify_revert"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-verify-revert"
description: |-
This resource allows you to execute Check Point Verify Revert.
---

# checkpoint_management_verify_revert

This resource allows you to execute Check Point Verify Revert.

## Example Usage


```hcl
resource "checkpoint_management_verify_revert" "example" {
  to_session = "d49ed10c-649a-476a-8e80-8282eda00e15"
}
```

## Argument Reference

The following arguments are supported:

* `to_session` - (Required) Session unique identifier. Specify the session you would like to verify a revert operation to. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

