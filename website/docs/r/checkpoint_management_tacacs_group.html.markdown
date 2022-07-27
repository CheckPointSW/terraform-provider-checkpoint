---
layout: "checkpoint"
page_title: "checkpoint_management_tacacs_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-tacacs-group"
description: |-
This resource allows you to execute Check Point Tacacs Group.
---

# Resource: checkpoint_management_tacacs_group

This resource allows you to  add/update/delete Check Point Tacacs Group.

## Example Usage


```hcl
resource "checkpoint_management_tacacs_group" "example" {
  name = "New Tacacs Group 1"
  members = ["t1", "t3", "group1",]
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
