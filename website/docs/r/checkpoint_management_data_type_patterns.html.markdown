---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_patterns"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-patterns"
description: |-
This resource allows you to execute Check Point Data Type Patterns.
---

# checkpoint_management_data_type_patterns

This resource allows you to execute Check Point Data Type Patterns.

## Example Usage


```hcl
resource "checkpoint_management_data_type_patterns" "example" {
  name = "pattern-obj"
  description = "data type pattern object"
  patterns = [ "g*" , "^k" ]
  number_of_occurrences = 4
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `patterns` - (Optional) Regular expressions to be evaluated.patterns blocks are documented below.
* `number_of_occurrences` - (Optional) Define how many times the patterns must appear to be considered data to be protected. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 