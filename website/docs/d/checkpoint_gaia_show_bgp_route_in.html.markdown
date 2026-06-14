---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_bgp_route_in"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-bgp-route-in"
description: |-
This resource allows you to execute Check Point Show Bgp Route In.
---

# checkpoint_gaia_show_bgp_route_in

This resource allows you to execute Check Point Show Bgp Route In.

## Example Usage


```hcl
data "checkpoint_gaia_show_bgp_route_in" "example" {
  address = "1.1.1.1"
  mask_length = 32
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) Filter the results for a specific route address. 
* `mask_length` - (Required) Filter the results for a specific route mask-length. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

