---
layout: "checkpoint"
page_title: "checkpoint_management_data_dynamic_object"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-dynamic-object"
description: |-
  Use this data source to get information on an existing Check Point Dynamic Object.
---

# Data Source: checkpoint_management_data_dynamic_object

Use this data source to get information on an existing Check Point Dynamic Object.

## Example Usage


```hcl
resource "checkpoint_management_dynamic_object" "dynamic_object" {
    name = "Dynamic Object"
}

data "checkpoint_management_data_dynamic_object" "data_dynamic_object" {
    name = "${checkpoint_management_dynamic_object.dynamic_object.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
