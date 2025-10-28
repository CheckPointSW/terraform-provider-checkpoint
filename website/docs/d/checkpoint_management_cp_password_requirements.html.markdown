---
layout: "checkpoint"
page_title: "checkpoint_management_cp_password_requirements"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cp-password-requirements"
description: |-
  Use this data source to get information on an existing Check Point Cp Password Requirements.
---

# Data Source: checkpoint_management_cp_password_requirements

Use this data source to get information on an existing Check Point Cp Password Requirements.

## Example Usage


```hcl
data "checkpoint_management_cp_password_requirements" "data_test" {

}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object unique identifier.
* `min_password_length` - Minimum Check Point password length. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

