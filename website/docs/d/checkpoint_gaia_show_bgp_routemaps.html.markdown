---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_bgp_routemaps"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-bgp-routemaps"
description: |-
This resource allows you to execute Check Point Show Bgp Routemaps.
---

# checkpoint_gaia_show_bgp_routemaps

This resource allows you to execute Check Point Show Bgp Routemaps.

## Example Usage


```hcl
data "checkpoint_gaia_show_bgp_routemaps" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the output by the group type in either ascending or descending order. By default, the group types will be sorted in the order: confederation, external, internal. Within each group type, the items will be sorted according to the AS number in either ascending or descending order. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

