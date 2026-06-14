---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_pim_joins"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-pim-joins"
description: |-
This resource allows you to execute Check Point Show Pim Joins.
---

# checkpoint_gaia_show_pim_joins

This resource allows you to execute Check Point Show Pim Joins.

## Example Usage


```hcl
data "checkpoint_gaia_show_pim_joins" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results. 
* `offset` - (Optional) The number of results to initially skip. 
* `order` - (Optional) Sorts results in either ascending or descending order. 
* `detailed` - (Optional) Show sparse-mode detailed join state. 
* `group` - (Optional) Show sparse-mode join state by group. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

