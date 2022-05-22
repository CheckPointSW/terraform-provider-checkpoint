---
layout: "checkpoint"
page_title: "checkpoint_management_set_idp_default_assignment"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-idp-default-assignment"
description: |-
This resource allows you to execute Check Point Set Idp Default Assignment.
---

# checkpoint_management_set_idp_default_assignment

This resource allows you to execute Check Point Set Idp Default Assignment.

## Example Usage


```hcl
resource "checkpoint_management_set_idp_default_assignment" "example" {
  identity_provider = "azure"
}
```

## Argument Reference

The following arguments are supported:

* `identity_provider` - (Required) Represents the Identity Provider to be used for Login by this assignment identified by the name or UID, to cancel existing assignment should set to 'none'. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

