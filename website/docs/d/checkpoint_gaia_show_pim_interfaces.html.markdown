---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_pim_interfaces"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-pim-interfaces"
description: |-
This resource allows you to execute Check Point Show Pim Interfaces.
---

# checkpoint_gaia_show_pim_interfaces

This resource allows you to execute Check Point Show Pim Interfaces.

## Example Usage


```hcl
data "checkpoint_gaia_show_pim_interfaces" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results. 
* `offset` - (Optional) The number of results to initially skip. 
* `order` - (Optional) Sorts results in either ascending or descending order. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

