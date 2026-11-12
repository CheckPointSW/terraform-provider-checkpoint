---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_weighted_keywords"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-weighted-keywords"
description: |-
This resource allows you to execute Check Point Data Type Weighted Keywords.
---

# checkpoint_management_data_type_weighted_keywords

This resource allows you to execute Check Point Data Type Weighted Keywords.

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
  weighted_keywords {
   keyword = "name2"
   weight = "5"
   max_weight = "8"
   regex = false
 }
  sum_of_weights_threshold = 10
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `weighted_keywords` - (Required) List of keywords or phrases.weighted_keywords blocks are documented below.
* `sum_of_weights_threshold` - (Optional) Define the number of appearances, by weight, of all the keywords that, beyond this threshold,
 the data containing this list of words or phrases will be recognized as data to be protected. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.  
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`weighted_keywords` supports the following:

* `keyword` - (Required) keyword or regular expression to be weighted. 
* `weight` - (Optional) Weight of the expression. 
* `max_weight` - (Optional) Max weight of the expression. 
* `regex` - (Optional) Determine whether to consider the expression as a regular expression. 
