---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_statistics_view_info"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-statistics-view-info"
description: |-
This resource allows you to execute Check Point Show Statistics View Info.
---

# checkpoint_gaia_show_statistics_view_info

This resource allows you to execute Check Point Show Statistics View Info.

## Example Usage


```hcl
data "checkpoint_gaia_show_statistics_view_info" "example" {
  filter = ["CPVIEW.Hardware-Health",]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filter the results by a list of view-IDs filter blocks are documented below.
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

