---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_traditional_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-traditional-group"
description: |-
This resource allows you to execute Check Point Data Type Traditional Group.
---

# checkpoint_management_data_type_traditional_group

This resource allows you to execute Check Point Data Type Traditional Group.

## Example Usage


```hcl
resource "checkpoint_management_data_type_traditional_group" "example" {
  name = "trad-group-obj"
  description = "traditional group object"
  data_types = [ "SSH Private Key" , "CSV File"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `data_types` - (Optional) List of data-types.
If data matches any of the data types in the group, the data type group is matched.
Identified by name or UID.data_types blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 