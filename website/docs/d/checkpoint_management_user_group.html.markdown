---
layout: "checkpoint"
page_title: "checkpoint_management_user_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-user-group"
description: |-
This resource allows you to execute Check Point User Group.
---

# Data Source: checkpoint_management_user_group

This resource allows you to execute Check Point User Group.

## Example Usage


```hcl
resource "checkpoint_management_user_group" "user_group" {
    name = "user_group"
    email = "user@email.com"
}

data "checkpoint_management_user_group" "test_user_group" {
    name = "${checkpoint_management_user_group.user_group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `email` - Email Address. 
* `members` - Collection of User Group objects identified by the name or UID.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. 
* `comments` - Comments string.