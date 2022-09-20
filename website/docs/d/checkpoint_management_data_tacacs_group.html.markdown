---
layout: "checkpoint"
page_title: "checkpoint_management_data_tacacs_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-tacscs-group"
description: |-
Use this data source to get information on an existing Check Point Tacacs Group.
---

# Data Source: checkpoint_management_tacacs_group

Use this data source to get information on an exsisting Check Point Tacacs Group.

## Example Usage


```hcl
resource "checkpoint_management_tacacs_group" "tacacsGroup" {
    name = "My Tacacs Group"
}

data "checkpoint_management_tacacs_group" "data_tacacs_group" {
    name = "${checkpoint_management_tacacs_group.tacacsGroup.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `members` - Collection of tacacs servers identified by the name or UID.members blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
