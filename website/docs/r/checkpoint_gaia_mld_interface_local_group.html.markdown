---
layout: "checkpoint"
page_title: "checkpoint_gaia_mld_interface_local_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-mld-interface-local-group"
description: |-
This resource allows you to execute Check Point Mld Interface Local Group.
---

# checkpoint_gaia_mld_interface_local_group

This resource allows you to execute Check Point Mld Interface Local Group.

## Example Usage


```hcl
resource "checkpoint_gaia_mld_interface_local_group" "example" {
  interface   = "eth0"
  local_group = "ff02::beef"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the MLD interface 
* `local_group` - (Required) The locally configured group address that this MLD interface receives multicast data for 
