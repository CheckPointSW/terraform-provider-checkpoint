---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_file_attributes"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-file-attributes"
description: |-
This resource allows you to execute Check Point Data Type File Attributes.
---

# checkpoint_management_data_type_file_attributes

This resource allows you to execute Check Point Data Type File Attributes.

## Example Usage


```hcl
resource "checkpoint_management_data_type_file_attributes" "example" {
  name = "file-attr-obj"
  match_by_file_type = true
  file_groups_list = ["Viewer"]
  match_by_file_name = true
  file_name_contains = "expression"
  match_by_file_size = true
  file_size = 14
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `match_by_file_type` - (Optional) Determine whether to consider file type. 
* `file_groups_list` - (Optional) The file must be one of the types specified in the list.
Identified by name or UID.file_groups_list blocks are documented below.
* `match_by_file_name` - (Optional) Determine whether to consider file name. 
* `file_name_contains` - (Optional) File name should contain the expression. 
* `match_by_file_size` - (Optional) Determine whether to consider file size. 
* `file_size` - (Optional) Min File size in KB. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.  
