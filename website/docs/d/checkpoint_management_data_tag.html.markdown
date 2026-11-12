---
layout: "checkpoint"
page_title: "checkpoint_management_data_tag"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-tag"
description: |-
Use this data source to get information on an existing Check Point Tag.
---

# Data Source: checkpoint_management_tag

Use this data source to get information on an exsisting Check Point Tag.

## Example Usage


```hcl
resource "checkpoint_management_tag" "myTag" {
    name = "My Tag"
}

data "checkpoint_management_tag" "data_tag" {
    name = "${checkpoint_management_tag.myTag.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string. 