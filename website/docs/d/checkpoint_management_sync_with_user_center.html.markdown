---
layout: "checkpoint"
page_title: "checkpoint_management_sync_with_user_center"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-sync-with-user-center"
description: |-
Use this data source to get information on an existing Check Point Set Sync With User Center.
---

# checkpoint_management_set_sync_with_user_center

Use this data source to get information on an existing Check Point Set Sync With User Center.

## Example Usage


```hcl
data "checkpoint_management_sync_with_user_center" "data" {
}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object Identifier.
* `enabled` - This indicates whether the information is being synchronized with the user center once a day.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

