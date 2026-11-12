---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_compound_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-compound-group"
description: |-
TUse this data source to get information on an existing Check Point Data Type Compound Group.
---

# Data Source: checkpoint_management_data_type_compound_group

Use this data source to get information on an existing Check Point Data Type Compound Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_compound_group" "example" {
  name = "compound-group-obj"
  description = "Compound group object"
  matched_groups =  ["Source Code"]
  unmatched_groups = ["Large File"]
}

data "checkpoint_management_data_type_compound_group" "data" {
  name = "${checkpoint_management_data_type_compound_group.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` -  For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `matched_groups` - Each one of these data types must be matched - Select existing data types to add. Traffic must match all the data types of this group to match a rule.
Identified by name or UID.matched_groups blocks are documented below.
* `unmatched_groups` -  Each one of these data types must not be matched - Select existing data types to add to the definition. Traffic that does not contain any data matching the types in this list will match this compound data type.
Identified by name or UID.unmatched_groups blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
