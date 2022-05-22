---
layout: "checkpoint"
page_title: "checkpoint_management_domain_permissions_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-domain-permissions-profile"
description: |-
Use this data source to get information on an existing Check Point Domain Permissions Profile.
---

# Data Source: checkpoint_management_domain_permissions_profile

Use this data source to get information on an existing Check Point Domain Permissions Profile.

## Example Usage


```hcl
resource "checkpoint_management_domain_permissions_profile" "example" {
  name = "customize profile"
}

data "checkpoint_management_domain_permissions_profile" "data_domain_permissions_profile" {
  name = "${checkpoint_management_domain_permissions_profile.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
