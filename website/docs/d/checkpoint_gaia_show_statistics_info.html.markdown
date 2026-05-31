---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_statistics_info"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-statistics-info"
description: |-
This resource allows you to execute Check Point Show Statistics Info.
---

# checkpoint_gaia_show_statistics_info

This resource allows you to execute Check Point Show Statistics Info.

## Example Usage


```hcl
data "checkpoint_gaia_show_statistics_info" "example" {
  filter = ["UM_STAT.UM_CPU.num_of_cores", "UM_STAT.UM_CPU.UM_CPU_TABLE",]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filter the results by a list of labels and stat IDs filter blocks are documented below.
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

