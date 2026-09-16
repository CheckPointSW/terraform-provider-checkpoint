---
layout: "checkpoint"
page_title: "checkpoint_management_requirement"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-requirement"
description: |-
This resource allows you to execute Check Point Requirement.
---

# checkpoint_management_requirement

This resource allows you to execute Check Point Requirement.

## Example Usage


```hcl
resource "checkpoint_management_requirement" "example" {
  name = "MyReq"
  regulation = "MyReg"
  comments = "My New Requirement"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Requirement name.
* `regulation` - (Required) The relevant regulation. Identified by name or UID.
* `best_practices` - (Optional) UIDs or IDs of the relevant best practices for the requirement.best_practices blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) The requirement comments.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `new_regulation_uid` - (Optional) The new regulation UID for this requirement.
* `regulation_uid` - (Optional) The relevant regulation UID for this requirement.
