---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_maestro_security_groups"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-maestro-security-groups"
description: |-
This resource allows you to execute Check Point Show Maestro Security Groups.
---

# checkpoint_gaia_show_maestro_security_groups

This resource allows you to execute Check Point Show Maestro Security Groups.

## Example Usage


```hcl
data "checkpoint_gaia_show_maestro_security_groups" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `include_pending_changes` - (Optional) If true, show pending Security Groups changes. If false, show deployed topology 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

