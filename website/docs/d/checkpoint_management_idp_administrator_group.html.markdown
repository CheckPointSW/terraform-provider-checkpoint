---
layout: "checkpoint"
page_title: "checkpoint_management_idp_administrator_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-idp-administrator-group"
description: |-
Use this data source to get information on an existing Check Point Idp Administrator Group.
---

# Data Source: checkpoint_management_idp_administrator_group

Use this data source to get information on an existing Check Point Idp Administrator Group.

## Example Usage


```hcl
resource "checkpoint_management_idp_administrator_group" "example" {
  name = "my super group"
  group_id = "it-team"
  multi_domain_profile = "domain super user"
}

data "checkpoint_management_idp_administrator_group" "data_idp_administrator_group" {
  name = "${checkpoint_management_idp_administrator_group.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
