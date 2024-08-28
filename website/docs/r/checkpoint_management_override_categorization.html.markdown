---
layout: "checkpoint"
page_title: "checkpoint_management_override_categorization"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-override-categorization"
description: |-
This resource allows you to execute Check Point Override Categorization.
---

# checkpoint_management_override_categorization

This resource allows you to execute Check Point Override Categorization.

## Example Usage


```hcl
resource "checkpoint_management_override_categorization" "example" {
  url = "newOverride"
  new_primary_category = "Botnets"
  risk = "low"
}
```

## Argument Reference

The following arguments are supported:

* `url` - (Required) The URL for which we want to update the category and risk definitions, the URL and the object name are the same for Override Categorization. 
* `url_defined_as_regular_expression` - (Optional) States whether the URL is defined as a Regular Expression or not. 
* `new_primary_category` - (Optional) Uid or name of the primary category based on its most defining aspect.
* `tags` - (Optional) Collection of tag identifiers. 
* `risk` - (Optional) States the override categorization risk. 
* `additional_categories` - (Optional) Uid or name of the categories to override in the Application and URL Filtering or Threat Prevention.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
