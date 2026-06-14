---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_ipv6_pim_group_rp_mapping"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-ipv6-pim-group-rp-mapping"
description: |-
This resource allows you to execute Check Point Show Ipv6 Pim Group Rp Mapping.
---

# checkpoint_gaia_show_ipv6_pim_group_rp_mapping

This resource allows you to execute Check Point Show Ipv6 Pim Group Rp Mapping.

## Example Usage


```hcl
data "checkpoint_gaia_show_ipv6_pim_group_rp_mapping" "example" {
  group = "ff12::"
}
```

## Argument Reference

The following arguments are supported:

* `group` - (Optional) Specifies the relevant multicast group 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

