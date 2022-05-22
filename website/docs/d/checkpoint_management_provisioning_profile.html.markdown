---
layout: "checkpoint"
page_title: "checkpoint_management_provisioning_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-provisioning-profile"
description: |-
Use this data source to get information on an existing Check Point Provisioning Profile.
---

# Data Source: checkpoint_management_provisioning_profile

Use this data source to get information on an existing Check Point Provisioning Profile.

## Example Usage


```hcl
data "checkpoint_management_provisioning_profile" "example" {
  name = "prv_gaia_profile"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

