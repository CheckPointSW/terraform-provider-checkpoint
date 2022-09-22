---
layout: "checkpoint"
page_title: "checkpoint_management_threat_layer"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-layer"
description: |-
Use this data source to get information on an existing Check Point Threat Layer.
---

# Data Source: checkpoint_management_threat_layer

Use this data source to get information on an existing Check Point Threat Layer.

## Example Usage


```hcl
resource "checkpoint_management_threat_layer" "example" {
  name = "New Layer 1"
}

data "checkpoint_management_threat_layer" "data_threat_layer" {
    name = "${checkpoint_management_threat_layer.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string. 