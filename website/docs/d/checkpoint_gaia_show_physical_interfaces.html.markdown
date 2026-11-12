---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_physical_interfaces"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-physical-interfaces"
description: |-
This resource allows you to execute Check Point Show Physical Interfaces.
---

# checkpoint_gaia_show_physical_interfaces

This resource allows you to execute Check Point Show Physical Interfaces.

## Example Usage


```hcl
data "checkpoint_gaia_show_physical_interfaces" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

