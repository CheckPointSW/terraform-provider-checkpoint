---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_traditional_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-traditional-group"
description: |-
Use this data source to get information on an existing Check Point Data Type Traditional Group.
---

# Data Source: checkpoint_management_data_type_traditional_group

Use this data source to get information on an existing Check Point Data Type Traditional Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_traditional_group" "example" {
  name = "trad-group-obj"
  description = "traditional group object"
  data_types = [ "SSH Private Key" , "CSV File"]
}

data "checkpoint_management_data_type_traditional_group" "data" {
  name = "${checkpoint_management_data_type_traditional_group.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` - For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `data_types` -  List of data-types. If data matches any of the data types in the group, the data type group is matched.
Identified by name or UID.data_types blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string.
