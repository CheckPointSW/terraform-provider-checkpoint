---
layout: "checkpoint"
page_title: "checkpoint_management_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-group"
description: |-
  This resource allows you to add/update/delete Check Point Group.
---

# checkpoint_management_group

This resource allows you to add/update/delete Check Point Group.

## Example Usage


```hcl
resource "checkpoint_management_group" "example" {
  name = "New Group 4"
  members = [ "New Host 1", "My Test Host 3" ]
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `members` - (Optional) Collection of Network objects identified by the name or UID.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `groups` - (Optional) Collection of group identifiers.
* `tags` - (Optional) Collection of tag identifiers.

