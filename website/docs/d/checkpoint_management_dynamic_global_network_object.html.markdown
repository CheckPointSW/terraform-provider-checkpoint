---
layout: "checkpoint"
page_title: "checkpoint_management_dynamic_global_network_object"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-dynamic-global-network-object"
description: |-
Use this data source to get information on an existing Check Point Dynamic Global Network Object.
---

# Data Source: checkpoint_management_dynamic_global_network_object

Use this data source to get information on an existing Check Point Dynamic Global Network Object.

## Example Usage


```hcl
resource "checkpoint_management_dynamic_global_network_object" "example" {
  name = "obj_global"
}

data "checkpoint_management_dynamic_global_network_object" "data_example" {
  name = "${checkpoint_management_dynamic_global_network_object.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.