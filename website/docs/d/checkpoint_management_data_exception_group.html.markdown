---
layout: "checkpoint"
page_title: "checkpoint_management_data_exception_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-exception-group"
description: |-
  Use this data source to get information on an existing Check Point Exception Group.
---

# Data Source: checkpoint_management_data_exception_group

Use this data source to get information on an existing Check Point Exception Group.

## Example Usage


```hcl
resource "checkpoint_management_exception_group" "exception_group" {
    name = "exception group"
	apply_on = "manually-select-threat-rules"
}

data "checkpoint_management_data_exception_group" "data_exception_group" {
    name = "${checkpoint_management_exception_group.exception_group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `applied_profile` - The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile. 
* `apply_on` - An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 