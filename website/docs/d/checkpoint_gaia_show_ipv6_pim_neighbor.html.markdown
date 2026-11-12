---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_ipv6_pim_neighbor"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-ipv6-pim-neighbor"
description: |-
This resource allows you to execute Check Point Show Ipv6 Pim Neighbor.
---

# checkpoint_gaia_show_ipv6_pim_neighbor

This resource allows you to execute Check Point Show Ipv6 Pim Neighbor.

## Example Usage


```hcl
data "checkpoint_gaia_show_ipv6_pim_neighbor" "example" {
  address = "c22::231"
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) The neighbor address. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

