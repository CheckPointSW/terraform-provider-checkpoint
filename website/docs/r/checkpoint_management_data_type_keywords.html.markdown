---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_keywords"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-keywords"
description: |-
This resource allows you to execute Check Point Data Type Keywords.
---

# checkpoint_management_data_type_keywords

This resource allows you to execute Check Point Data Type Keywords.

## Example Usage


```hcl
resource "checkpoint_management_data_type_keywords" "example" {
  name = "keywords_obj"
  description = "keywords object"
  keywords = ["word1","word2"]
  data_match_threshold = "all-keywords"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `description` - (Optional) For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `keywords` - (Optional) Specify keywords or phrases to search for.keywords blocks are documented below.
* `data_match_threshold` - (Optional) If set to all-keywords - the data will be matched to the rule only if all the words in the list appear in the data contents.
When set to min-keywords any number of the words may appear according to configuration. 
* `min_number_of_keywords` - (Optional) Define how many of the words in the list must appear in the contents of the data to match the rule. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
