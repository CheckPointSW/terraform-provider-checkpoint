---
layout: "checkpoint"
page_title: "checkpoint_management_data_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-group"
description: |-
  Use this data source to get information on an existing Check Point Group.
---

# checkpoint_management_data_group

Use this data source to get information on an existing Check Point Group.

## Example Usage


```hcl
resource "checkpoint_management_group" "group" {
    name = "My Group"
}

data "checkpoint_management_data_group" "data_group" {
    name = "${checkpoint_management_group.group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `members` - Collection of Network objects identified by the name or UID.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.

