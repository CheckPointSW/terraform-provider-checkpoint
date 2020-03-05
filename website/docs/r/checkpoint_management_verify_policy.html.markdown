---
layout: "checkpoint"
page_title: "checkpoint_management_verify_policy"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-verify-policy"
description: |-
This resource allows you to execute Check Point Verify Policy.
---

# checkpoint_management_verify_policy

This resource allows you to execute Check Point Verify Policy.

## Example Usage


```hcl
resource "checkpoint_management_verify_policy" "example" {
  policy_package = "standard"
}
```

## Argument Reference

The following arguments are supported:

* `policy_package` - (Required) Policy package identified by the name or UID. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

