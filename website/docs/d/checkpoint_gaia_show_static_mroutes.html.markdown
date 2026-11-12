---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_static_mroutes"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-static-mroutes"
description: |-
This resource allows you to execute Check Point Show Static Mroutes.
---

# checkpoint_gaia_show_static_mroutes

This resource allows you to execute Check Point Show Static Mroutes.

## Example Usage


```hcl
data "checkpoint_gaia_show_static_mroutes" "example" {
  limit = 2
  offset = 0
  order = "DESC"
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the routes by priority in either ascending or descending order 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

