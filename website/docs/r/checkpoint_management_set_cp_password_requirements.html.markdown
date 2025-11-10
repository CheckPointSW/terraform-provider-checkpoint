---
layout: "checkpoint"
page_title: "checkpoint_management_set_cp_password_requirements"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-cp-password-requirements"
description: |-
 This resource allows you to execute Check Point Set Cp Password Requirements.
---

# checkpoint_management_set_cp_password_requirements

This resource allows you to execute Check Point Set Cp Password Requirements.

## Example Usage


```hcl
resource "checkpoint_management_set_cp_password_requirements" "example" {
  min_password_length = 7
}
```

## Argument Reference

The following arguments are supported:

* `min_password_length` - (Optional) Minimum Check Point password length. 
* `uid` - Object unique identifier.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

