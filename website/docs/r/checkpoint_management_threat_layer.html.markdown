---
layout: "checkpoint"
page_title: "checkpoint_management_threat_layer"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-threat-layer"
description: |-
This resource allows you to add/update/delete Check Point Threat Layer.
---

# Resource: checkpoint_management_threat_layer

This resource allows you to add/update/delete Check Point Threat Layer.

## Example Usage


```hcl
resource "checkpoint_management_threat_layer" "example" {
  name = "New Layer 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `add_default_rule` - (Optional) Indicates whether to include a default rule in the new layer. 
