---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_isis_topology"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-isis-topology"
description: |-
This resource allows you to execute Check Point Show Isis Topology.
---

# checkpoint_gaia_show_isis_topology

This resource allows you to execute Check Point Show Isis Topology.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_isis" "isis_setup" {
  system_id = "0101.0101.0101"
}

data "checkpoint_gaia_show_isis_topology" "example" {
  protocol_instance = "default"
  topology          = "ipv4"

  depends_on = [checkpoint_gaia_command_set_isis.isis_setup]
}
```

## Argument Reference

The following arguments are supported:

* `protocol_instance` - (Optional) The instance to be queried 
* `topology` - (Optional) The topology to be queried 
* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the topology entries by their level first, and then their system-ids in either ascending or descending order. 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

