---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_bgp_routes_in"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-bgp-routes-in"
description: |-
This resource allows you to execute Check Point Show Bgp Routes In.
---

# checkpoint_gaia_show_bgp_routes_in

This resource allows you to execute Check Point Show Bgp Routes In.

## Example Usage


```hcl
data "checkpoint_gaia_show_bgp_routes_in" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the peers first by their AS, then by their IDs in either ascending or descending order. 
* `peer` - (Optional) Filter the results for a specific peer. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

