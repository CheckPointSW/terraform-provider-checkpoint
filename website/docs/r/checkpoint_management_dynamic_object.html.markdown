---
layout: "checkpoint"
page_title: "checkpoint_management_dynamic_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-dynamic-object"
description: |-
This resource allows you to execute Check Point Dynamic Object.
---

# checkpoint_management_dynamic_object

This resource allows you to execute Check Point Dynamic Object.

## Example Usage


```hcl
resource "checkpoint_management_dynamic_object" "example" {
  name = "Dynamic_Object_1"
  comments = "My Dynamic Object 1"
  color = "yellow"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
