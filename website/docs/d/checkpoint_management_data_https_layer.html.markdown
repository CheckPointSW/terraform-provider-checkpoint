---
layout: "checkpoint"
page_title: "checkpoint_management_data_https_layer"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-https-layer"
description: |-
  Use this data source to get information on an existing Check Point Https Layer.
---

# Data Source: checkpoint_management_data_https_layer

Use this data source to get information on an existing Check Point Https Layer.

## Example Usage


```hcl
resource "checkpoint_management_https_layer" "https_layer" {
    name = "%s"
}

data "checkpoint_management_data_https_layer" "data_https_layer" {
    name = "${checkpoint_management_https_layer.https_layer.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `shared` - Define the Layer as Shared (TRUE/FALSE). 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
