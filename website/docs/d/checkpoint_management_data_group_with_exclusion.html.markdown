---
layout: "checkpoint"
page_title: "checkpoint_management_data_group_with_exclusion"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-group-with-exclusion"
description: |-
  Use this data source to get information on an existing Check Point Group With Exclusion.
---

# Data Source: checkpoint_management_data_group_with_exclusion

Use this data source to get information on an existing Check Point Group With Exclusion.

## Example Usage


```hcl
resource "checkpoint_management_group" "group1" {
    name = "group1"
}

resource "checkpoint_management_group" "group2" {
    name = "group2"
}

resource "checkpoint_management_group_with_exclusion" "group_with_exclusion" {
    name = "Group with exclusion"
    include = "${checkpoint_management_group.group1.name}"
    except = "${checkpoint_management_group.group2.name}"
}

data "checkpoint_management_data_group_with_exclusion" "data_group_with_exclusion" {
    name = "${checkpoint_management_group_with_exclusion.group_with_exclusion.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `except` - Name or UID of an object which the group excludes. 
* `include` - Name or UID of an object which the group includes. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.
