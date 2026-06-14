---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_mld_interface_stats"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-mld-interface-stats"
description: |-
This resource allows you to execute Check Point Show Mld Interface Stats.
---

# checkpoint_gaia_show_mld_interface_stats

This resource allows you to execute Check Point Show Mld Interface Stats.

## Example Usage


```hcl
data "checkpoint_gaia_show_mld_interface_stats" "example" {
  interface = "eth0"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the MLD interface 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

