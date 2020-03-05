---
layout: "checkpoint"
page_title: "checkpoint_management_group_with_exclusion"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-group-with-exclusion"
description: |-
This resource allows you to execute Check Point Group With Exclusion.
---

# checkpoint_management_group_with_exclusion

This resource allows you to execute Check Point Group With Exclusion.

## Example Usage


```hcl
resource "checkpoint_management_group_with_exclusion" "example" {
  name = "Group with exclusion"
  include = "New Group 1"
  except = "New Group 2"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `except` - (Optional) Name or UID of an object which the group excludes. 
* `include` - (Optional) Name or UID of an object which the group includes. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
