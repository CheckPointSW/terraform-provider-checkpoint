---
layout: "checkpoint"
page_title: "checkpoint_management_set_sync_with_user_center"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-sync-with-user-center"
description: |-
This resource allows you to execute Check Point Set Sync With User Center.
---

# checkpoint_management_set_sync_with_user_center

This resource allows you to execute Check Point Set Sync With User Center.

## Example Usage


```hcl
resource "checkpoint_management_set_sync_with_user_center" "example" {
  enabled = false
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Synchronize information once a day.<br>Warning: Synchronizing with the Check Point UserCenter requires a valid licence. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

