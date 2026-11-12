---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_ospf_neighbor"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-ospf-neighbor"
description: |-
This resource allows you to execute Check Point Show Ospf Neighbor.
---

# checkpoint_gaia_show_ospf_neighbor

This resource allows you to execute Check Point Show Ospf Neighbor.

## Example Usage


```hcl
data "checkpoint_gaia_show_ospf_neighbor" "example" {
  protocol_instance = "default"
  neighbor = "10.10.10.80"
}
```

## Argument Reference

The following arguments are supported:

* `neighbor` - (Required) Existing OSPFv2 Neighbor 
* `protocol_instance` - (Optional) Existing OSPFv2 Instance 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

