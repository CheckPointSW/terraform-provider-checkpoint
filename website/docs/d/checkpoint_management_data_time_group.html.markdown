---
layout: "checkpoint"
page_title: "checkpoint_management_data_time_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-time-group"
description: |-
  Use this data source to get information on an existing Check Point Time Group.
---

# Data Source: checkpoint_management_data_time_group

Use this data source to get information on an existing Check Point Time Group.

## Example Usage


```hcl
resource "checkpoint_management_time_group" "time_group" {
    name = "time group"
}

data "checkpoint_management_data_time_group" "data_time_group" {
    name = "${checkpoint_management_time_group.time_group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `members` - Collection of Time Group objects identified by the name or UID.members blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.