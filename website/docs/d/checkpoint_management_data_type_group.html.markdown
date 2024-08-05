---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-group"
description: |-
Use this data source to get information on an existing Check Point Data Type Group.
---

# Data Source: checkpoint_management_data_type_group

Use this data source to get information on an existing Check Point Data Type Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_group" "example" {
  name = "data-group-obj"
  description = "add data type group object"
  file_type = ["Archive File"]
  file_content = ["CSV File"]
}

data "checkpoint_management_data_type_group" "data" {
  name = "${checkpoint_management_data_type_group.test.name}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` -  For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `file_type` -  List of data-types-file-attributes objects.
Identified by name or UID.file_type blocks are documented below.
* `file_content` -  List of Data Types. At least one must be matched.
Identified by name or UID.file_content blocks are documented below.
* `tags` -  Collection of tag identifiers. tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string.  
