---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_dynamic_layer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-dynamic-layer"
description: |-
This resource allows you to execute Check Point Show Dynamic Layer.
---

# checkpoint_gaia_show_dynamic_layer

This resource allows you to execute Check Point Show Dynamic Layer.

## Example Usage


```hcl
data "checkpoint_gaia_show_dynamic_layer" "example" {
  name = "Network"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Must be unique in the domain. 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

