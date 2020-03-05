---
layout: "checkpoint"
page_title: "checkpoint_management_add_updatable_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-add-updatable-object"
description: |-
This resource allows you to execute Check Point Add Updatable Object.
---

# checkpoint_management_add_updatable_object

This resource allows you to execute Check Point Add Updatable Object.

## Example Usage


```hcl
resource "checkpoint_management_add_updatable_object" "example" {
  uri = "{{uri}}"
}
```

## Argument Reference

The following arguments are supported:

* `uri` - (Required) URI of the updatable object in the Updatable Objects Repository. 
* `uid_in_updatable_objects_repository` - (Required) Unique identifier of the updatable object in the Updatable Objects Repository. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

