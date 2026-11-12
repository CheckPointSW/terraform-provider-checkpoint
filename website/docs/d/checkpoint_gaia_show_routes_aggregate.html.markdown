---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_routes_aggregate"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-routes-aggregate"
description: |-
This resource allows you to execute Check Point Show Routes Aggregate.
---

# checkpoint_gaia_show_routes_aggregate

This resource allows you to execute Check Point Show Routes Aggregate.

## Example Usage


```hcl
data "checkpoint_gaia_show_routes_aggregate" "example" {
  limit = 1
  offset = 0
  order = "DESC"
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the routes in either ascending or descending order. 
* `address_family` - (Optional) Address family of routes returned. IPv6 route monitoring, or specifying "inet6" for this field, is only supported on GAIA versions R81.10 and up. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

