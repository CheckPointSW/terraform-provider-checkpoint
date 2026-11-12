---
layout: "checkpoint"
page_title: "checkpoint_management_identity_tag"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-identity-tag"
description: |-
This resource allows you to execute Check Point Identity Tag.
---

# Data Source: checkpoint_management_identity_tag

This resource allows you to execute Check Point Identity Tag.

## Example Usage


```hcl
resource "checkpoint_management_identity_tag" "test" {
    name = "my identity tag"
    external_identifier = "cisco ise security group"
}

data "checkpoint_management_identity_tag" "data_test" {
    name = "${checkpoint_management_identity_tag.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `external_identifier` - External identifier. For example: Cisco ISE security group tag. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string.