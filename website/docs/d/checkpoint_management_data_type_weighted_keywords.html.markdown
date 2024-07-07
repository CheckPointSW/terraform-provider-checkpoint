---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_weighted_keywords"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-weighted-keywords"
description: |-
Use this data source to get information on an existing Check Point Data Type Weighted Keywords.
---

# Data Source: checkpoint_management_data_type_weighted_keywords

Use this data source to get information on an existing Check Point Data Type Weighted Keywords.

## Example Usage


```hcl
resource "checkpoint_management_data_type_weighted_keywords" "example" {
  name = "weighted-words-obj"
  weighted_keywords {
     keyword = "name"
     weight = "4"
     max_weight = "4"
    regex = true
  }
  sum_of_weights_threshold = 10
}

data "checkpoint_management_data_type_weighted_keywords" "data" {
 name = "${checkpoint_management_data_type_weighted_keywords.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `description` -  For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `weighted_keywords` -  List of keywords or phrases.weighted_keywords blocks are documented below.
* `sum_of_weights_threshold` -  Define the number of appearances, by weight, of all the keywords that, beyond this threshold,
 the data containing this list of words or phrases will be recognized as data to be protected. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string.  


`weighted_keywords` supports the following:

* `keyword` -  keyword or regular expression to be weighted. 
* `weight` -  Weight of the expression. 
* `max_weight` -  Max weight of the expression. 
* `regex` -  Determine whether to consider the expression as a regular expression. 
