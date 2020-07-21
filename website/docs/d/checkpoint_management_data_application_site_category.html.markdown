---
layout: "checkpoint"
page_title: "checkpoint_management_data_application_site_category"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-application-site-category"
description: |-
  Use this data source to get information on an existing Check Point Application Site Category.
---

# checkpoint_management_data_application_site_category

Use this data source to get information on an existing Check Point Application Site Category.

## Example Usage


```hcl
resource "checkpoint_management_application_site_category" "application_site_category" {
    name = "applicationcategory1"
}

data "checkpoint_management_data_application_site_category" "data_application_site_category" {
    name = "${checkpoint_management_application_site_category.application_site_category.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `description` - Description string
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.
