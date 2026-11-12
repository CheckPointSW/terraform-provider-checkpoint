---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_maestro_gateways"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-maestro-gateways"
description: |-
This resource allows you to execute Check Point Show Maestro Gateways.
---

# checkpoint_gaia_show_maestro_gateways

This resource allows you to execute Check Point Show Maestro Gateways.

## Example Usage


```hcl
data "checkpoint_gaia_show_maestro_gateways" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `include_pending_changes` - (Optional) If true, show pending topology. If false, show deployed topology 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

