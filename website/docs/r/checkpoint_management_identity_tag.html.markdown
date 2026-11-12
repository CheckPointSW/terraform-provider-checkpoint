---
layout: "checkpoint"
page_title: "checkpoint_management_identity_tag"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-identity-tag"
description: |-
This resource allows you to execute Check Point Identity Tag.
---

# Resource: checkpoint_management_identity_tag

This resource allows you to execute Check Point Identity Tag.

## Example Usage


```hcl
resource "checkpoint_management_identity_tag" "example" {
  name = "mytag"
  external_identifier = "some external identifier"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `external_identifier` - (Optional) External identifier. For example: Cisco ISE security group tag. 
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.