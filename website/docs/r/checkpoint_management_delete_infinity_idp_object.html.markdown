---
layout: "checkpoint"
page_title: "checkpoint_management_delete_infinity_idp_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-infinity-idp-object"
description: |-
This resource allows you to execute Check Point Delete Infinity Idp Object.
---

# checkpoint_management_delete_infinity_idp_object

This resource allows you to execute Check Point Delete Infinity Idp Object.

## Example Usage


```hcl
resource "checkpoint_management_delete_infinity_idp_object" "example" {
  name = "object-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object UID. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

