---
layout: "checkpoint"
page_title: "checkpoint_management_delete_infinity_idp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-infinity-idp"
description: |-
Use this data source to get information on an existing Check Point Delete Infinity Idp.
---

# Data Source: checkpoint_management_delete_infinity_idp

Use this data source to get information on an existing Check Point Delete Infinity Idp.

## Example Usage


```hcl
data "checkpoint_management_infinity_idp" "data" {
  name = "object-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `idp_domains` - List of domains configured in the Infinity Identity Provider object in Infinity Portal.
* `idp_id` - Identity Provider unique identifier in Infinity Portal.
* `idp_name` - Identity Provider name in Infinity Portal.
* `idp_type` - Identity Provider type in Infinity Portal.
* `tags` - Collection of tag identifiers.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

