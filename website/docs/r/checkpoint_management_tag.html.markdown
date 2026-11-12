---
layout: "checkpoint"
page_title: "checkpoint_management_tag"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-tag"
description: |-
This resource allows you to execute Check Point Tag.
---

# Resource: checkpoint_management_tag

This resource allows you to add/update/delete Check Point Tag.

## Example Usage


```hcl
resource "checkpoint_management_tag" "example" {
  name = "My Tag"
  tags = ["tag1", "tag2",]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name Must be unique in the domain. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
