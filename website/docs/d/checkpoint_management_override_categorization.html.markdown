---
layout: "checkpoint"
page_title: "checkpoint_management_override_categorization"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-override-categorization"
description: |-
Use this data source to get information on an existing Check Point Override Categorization.
---

# Data Source: checkpoint_management_override_categorization

Use this data source to get information on an existing Check Point Override Categorization.

## Example Usage

```hcl
resource "checkpoint_management_override_categorization" "example" {
  url = "newOverride"
  new_primary_category = "Botnets"
  risk = "low"
}
data "checkpoint_management_override_categorization" "data" {
  url = "${checkpoint_management_override_categorization.example.url}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `url` - (Optional) The URL for which we want to update the category and risk definitions, the URL and the object name are the same for Override Categorization. 
* `url_defined_as_regular_expression` -  States whether the URL is defined as a Regular Expression or not. 
* `new_primary_category` -  Uid or name of the primary category based on its most defining aspect. 
* `tags` -  Collection of tag identifiers. 
* `risk` -  States the override categorization risk. 
* `additional_categories` -  Uid or name of the categories to override in the Application and URL Filtering or Threat Prevention.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

