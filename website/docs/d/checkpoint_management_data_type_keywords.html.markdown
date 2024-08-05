---
layout: "checkpoint"
page_title: "checkpoint_management_data_type_keywords"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-type-keywords"
description: |-
Use this data source to get information on an existing Check Point Data Type Keywords.
---

# Data Source: checkpoint_management_data_type_keywords

Use this data source to get information on an existing Check Point Data Type Keywords.

## Example Usage


```hcl
resource "checkpoint_management_data_type_keywords" "example" {
  name = "keywords_obj"
  description = "keywords object"
  keywords = ["word1","word2"]
  data_match_threshold = "all-keywords"
}

data "checkpoint_management_data_type_keywords" "data" {
  name = "${checkpoint_management_data_type_keywords.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `description` -  For built-in data types, the description explains the purpose of this type of data representation.
For custom-made data types, you can use this field to provide more details. 
* `keywords` -  Specify keywords or phrases to search for.keywords blocks are documented below.
* `data_match_threshold` -  If set to all-keywords - the data will be matched to the rule only if all the words in the list appear in the data contents.
When set to min-keywords any number of the words may appear according to configuration. 
* `min_number_of_keywords` -  Define how many of the words in the list must appear in the contents of the data to match the rule. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

