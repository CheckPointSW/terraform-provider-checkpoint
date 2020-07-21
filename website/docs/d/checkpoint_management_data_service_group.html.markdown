---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-group"
description: |-
  Use this data source to get information on an existing Check Point Service Group.
---

# checkpoint_management_data_service_group

Use this data source to get information on an existing Check Point Service Group.

## Example Usage


```hcl
resource "checkpoint_management_service_group" "service_group" {
    name = "service group"
}

data "checkpoint_management_data_service_group" "data_service_group" {
    name = "${checkpoint_management_service_group.service_group.name}"
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