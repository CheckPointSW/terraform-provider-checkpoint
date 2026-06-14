---
layout: "checkpoint"
page_title: "checkpoint_gaia_virtual_switch"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-virtual-switch"
description: |-
This resource allows you to execute Check Point Virtual Switch.
---

# checkpoint_gaia_virtual_switch

This resource allows you to execute Check Point Virtual Switch.

## Example Usage


```hcl
resource "checkpoint_gaia_virtual_switch" "example" {
  resource_id = 1
  name        = "DMZ-Switch"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Virtual switch name 
* `resource_id` - (Optional) Virtual switch identifier 
* `interface` - (Optional) Network interface to be added 
* `set_if_exist` - (Optional) If another virtual switch with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original virtual switch's fields will be overwritten by the fields provided in the request payload! 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `action` - (Computed) Computed field, returned in the response. 
* `status` - (Computed) Computed field, returned in the response. 
* `message` - (Computed) Computed field, returned in the response. 
* `vsxd_task_id` - (Computed) Computed field, returned in the response. 
* `vs_id` - (Computed) Computed field, returned in the response. 
