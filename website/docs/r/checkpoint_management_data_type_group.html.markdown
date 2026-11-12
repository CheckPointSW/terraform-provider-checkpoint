---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-group"
description: |-
This resource allows you to execute Check Point Data Type Group.
---

# checkpoint_management_data_type_group

This resource allows you to execute Check Point Data Type Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_group" "example" {
  name = "data-group-obj"
  description = "add data type group object"
  file_type = ["Archive File"]
  file_content = ["CSV File"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `file_type` - (Optional) List of data-types-file-attributes objects.
Identified by name or UID.file_type blocks are documented below.
* `file_content` - (Optional) List of Data Types. At least one must be matched.
Identified by name or UID.file_content blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.  
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 