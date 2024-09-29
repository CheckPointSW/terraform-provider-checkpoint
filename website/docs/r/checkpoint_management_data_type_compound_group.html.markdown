---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_compound_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-compound-group"
description: |-
This resource allows you to execute Check Point Data Type Compound Group.
---

# checkpoint_management_data_type_compound_group

This resource allows you to execute Check Point Data Type Compound Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_compound_group" "example" {
  name = "compound-group-obj"
  description = "Compound group object"
  matched_groups =  ["Source Code"]
  unmatched_groups = ["Large File"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `matched_groups` - (Optional) Each one of these data types must be matched - Select existing data types to add. Traffic must match all the data types of this group to match a rule.
Identified by name or UID.matched_groups blocks are documented below.
* `unmatched_groups` - (Optional) Each one of these data types must not be matched - Select existing data types to add to the definition. Traffic that does not contain any data matching the types in this list will match this compound data type.
Identified by name or UID.unmatched_groups blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 