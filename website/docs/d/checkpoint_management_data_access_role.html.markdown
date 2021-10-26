---
layout: "checkpoint"
page_title: "checkpoint_management_data_access_role"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-role"
description: |- Use this data source to get information on an existing Check Point Access Role.
---

# Data Source: checkpoint_management_data_access_role

Use this data source to get information on an existing Check Point Access Role.

## Example Usage

```hcl
resource "checkpoint_management_access_role" "access_role" {
  name     = "My Access Role"
  comments = "my-Access-Role"
  machines {
    source    = "any"
    selection = ["any"]
  }
  users {
    selection = ["any"]
    source = "any"
  }
}

data "checkpoint_management_data_access_role" "data_access_role" {
  name = "${checkpoint_management_access_role.access_role.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required if uid is not given) Object name.
* `uid` - (Required name uid is not given) Object unique identifier.
