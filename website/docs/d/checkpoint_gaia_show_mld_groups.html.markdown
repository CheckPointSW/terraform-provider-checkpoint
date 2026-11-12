---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_mld_groups"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-mld-groups"
description: |-
This resource allows you to execute Check Point Show Mld Groups.
---

# checkpoint_gaia_show_mld_groups

This resource allows you to execute Check Point Show Mld Groups.

## Example Usage


```hcl
data "checkpoint_gaia_show_mld_groups" "example" {
  type = "local"
  interface = "all"
}
```

## Argument Reference

The following arguments are supported:

* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the interface MLD groups entries by interface name in either ascending or descending order 
* `type` - (Optional) The type of the MLD group 
* `interface` - (Optional) The name of the interface associated with the MLD groups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

