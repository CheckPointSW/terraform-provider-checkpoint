---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_file_attributes"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-file-attributes"
description: |-
Use this data source to get information on an existing Check Point Data Type File Attributes.
---

# Data Source: checkpoint_management_data_type_file_attributes

Use this data source to get information on an existing  Check Point Data Type File Attributes.

## Example Usage


```hcl
resource "checkpoint_management_data_type_file_type_attributes" "example" {
  name = "file-attr-obj"
  match_by_file_type = true
  file_groups_list = ["Viewer"]
  match_by_file_name = true
  file_name_contains = "expression"
  match_by_file_size = true
  file_size = 14
}

data "checkpoint_management_data_type_file_type_attributes" "data" {
  name = "${checkpoint_management_data_type_file_type_attributes.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` -  For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `match_by_file_type` -  Determine whether to consider file type. 
* `file_groups_list` -  The file must be one of the types specified in the list.
Identified by name or UID.file_groups_list blocks are documented below.
* `match_by_file_name` -  Determine whether to consider file name. 
* `file_name_contains` -  File name should contain the expression. 
* `match_by_file_size` -  Determine whether to consider file size. 
* `file_size` -  Min File size in KB. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 

