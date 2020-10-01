---
layout: "checkpoint"
page_title: "checkpoint_management_user_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-user-group"
description: |-
This resource allows you to execute Check Point User Group.
---

# Resource: checkpoint_management_user_group

This resource allows you to execute Check Point User Group.

## Example Usage


```hcl
resource "checkpoint_management_user_group" "example" {
  name = "myusergroup"
  email = "myusergroup@email.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `email` - (Optional) Email Address. 
* `members` - (Optional) Collection of User Group objects identified by the name or UID.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.