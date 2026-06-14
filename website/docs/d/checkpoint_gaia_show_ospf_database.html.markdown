---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_ospf_database"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-ospf-database"
description: |-
This resource allows you to execute Check Point Show Ospf Database.
---

# checkpoint_gaia_show_ospf_database

This resource allows you to execute Check Point Show Ospf Database.

## Example Usage


```hcl
data "checkpoint_gaia_show_ospf_database" "example" {
  protocol_instance = "default"
  ospf2_area        = "all"
  lsa_type          = "all"
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the results in either ascending or descending order. 
* `protocol_instance` - (Optional) Existing OSPFv2 Instance 
* `ospf2_area` - (Optional) Existing OSPFv2 Area 
* `lsa_type` - (Optional) LSA Type 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

