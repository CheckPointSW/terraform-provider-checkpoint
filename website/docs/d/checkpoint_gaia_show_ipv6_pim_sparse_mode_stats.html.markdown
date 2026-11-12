---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_ipv6_pim_sparse_mode_stats"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-ipv6-pim-sparse-mode-stats"
description: |-
This resource allows you to execute Check Point Show Ipv6 Pim Sparse Mode Stats.
---

# checkpoint_gaia_show_ipv6_pim_sparse_mode_stats

This resource allows you to execute Check Point Show Ipv6 Pim Sparse Mode Stats.

## Example Usage


```hcl
data "checkpoint_gaia_show_ipv6_pim_sparse_mode_stats" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

