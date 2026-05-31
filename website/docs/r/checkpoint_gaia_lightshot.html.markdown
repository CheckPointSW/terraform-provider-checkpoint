---
layout: "checkpoint"
page_title: "checkpoint_gaia_lightshot"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-lightshot"
description: |-
This resource allows you to execute Check Point Lightshot.
---

# checkpoint_gaia_lightshot

This resource allows you to execute Check Point Lightshot.

## Example Usage


```hcl
resource "checkpoint_gaia_lightshot" "example" {
  name = "lightshot_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of lightshot to add to lightshots list 
* `member_id` - (Computed) No description available. 
* `lightshot` - (Computed) Computed field, returned in the response. lightshot blocks are documented below.


`lightshot` supports the following:

* `name` - (Computed) Computed field, returned in the response. 
* `description` - (Computed) Computed field, returned in the response. 
* `size` - (Computed) Computed field, returned in the response. 
* `date` - (Computed) Computed field, returned in the response. 
