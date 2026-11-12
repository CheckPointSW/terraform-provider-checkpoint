---
layout: "checkpoint"
page_title: "checkpoint_management_delete_infinity_idp_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-infinity-idp-object"
description: |-
Use this data source to get information on an Check Point Delete Infinity Idp Object.
---

# Data Source: checkpoint_management_delete_infinity_idp_object

Use this data source to get information on an Check Point Delete Infinity Idp Object.

## Example Usage


```hcl
resource "checkpoint_management_infinity_idp_object" "example" {
  name = "object-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `description` - Description string.
* `display_name` - Entity name in the Management Server.
* `ext_id` - Entity unique identifier in the Identity Provider.
* `idp_display_name` - Identity Provider name in Management Server.
* `idp_id` - Identity Provider unique identifier in Infinity Portal.
* `idp_name` - Identity Provider name in Infinity Portal.
* `object_type` - Entity type - can be user/group/machine.
* `tags` - Collection of tag identifiers.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

