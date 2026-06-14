---
layout: "checkpoint"
page_title: "checkpoint_gaia_igmp_interface_local_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-igmp-interface-local-group"
description: |-
This resource allows you to execute Check Point Igmp Interface Local Group.
---

# checkpoint_gaia_igmp_interface_local_group

This resource allows you to execute Check Point Igmp Interface Local Group.

## Example Usage


```hcl
resource "checkpoint_gaia_igmp_interface_local_group" "example" {
  interface   = "eth0"
  local_group = "224.6.6.6"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the IGMP interface 
* `local_group` - (Required) The locally configured group address that this IGMP interface receives multicast data for 
