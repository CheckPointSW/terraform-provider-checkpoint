---
layout: "checkpoint"
page_title: "checkpoint_management_data_tacacs_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-tacscs-group"
description: |-
Use this data source to get information on an existing Check Point Tacacs Group.
---

# Data Source: checkpoint_management_data_tacacs_group

Use this data source to get information on an exsisting Check Point Tacacs Group.

## Example Usage


```hcl
resource "checkpoint_management_tacacs_group" "tacacsGroup" {
    name = "My Tacacs Group"
}

data "checkpoint_management_data_tacacs_group" "data_tacacs_group" {
    name = "${checkpoint_management_tacacs_group.tacacsGroup.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `members` - (Optional) Collection of tacacs servers identified by the name or UID.members blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 