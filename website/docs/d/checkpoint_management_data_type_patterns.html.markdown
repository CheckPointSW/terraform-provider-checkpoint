---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_patterns"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-patterns"
description: |-
Use this data source to get information on an existing Check Point Data Type Patterns.
---

# Data Source: checkpoint_management_data_type_patterns

Use this data source to get information on an existing Check Point Data Type Patterns.

## Example Usage


```hcl
resource "checkpoint_management_data_type_patterns" "example" {
  name = "pattern-obj"
  description = "data type pattern object"
  patterns = [ "g*" , "^k" ]
  number_of_occurrences = 4
}

data "checkpoint_management_data_type_patterns" "data" {
  name = "${checkpoint_management_data_type_patterns.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` - For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `patterns` -  Regular expressions to be evaluated. patterns blocks are documented below.
* `number_of_occurrences` - Define how many times the patterns must appear to be considered data to be protected. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

