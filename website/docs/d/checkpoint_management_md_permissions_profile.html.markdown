---
layout: "checkpoint"
page_title: "checkpoint_management_md_permissions_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-md-permissions-profile"
description: |-
Use this data source to get information on an existing Check Point Md Permissions Profile.
---

# Data Source: checkpoint_management_md_permissions_profile

Use this data source to get information on an existing Check Point Md Permissions Profile.

## Example Usage


```hcl
resource "checkpoint_management_md_permissions_profile" "example" {
  name = "manager profile"
}

data "checkpoint_management_md_permissions_profile" "data_md_permissions_profile" {
  name = "${checkpoint_management_md_permissions_profile.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
