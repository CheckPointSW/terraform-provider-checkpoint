---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_system_groups"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-system-groups"
description: |-
This resource allows you to execute Check Point Show System Groups.
---

# checkpoint_gaia_show_system_groups

This resource allows you to execute Check Point Show System Groups.

## Example Usage


```hcl
data "checkpoint_gaia_show_system_groups" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

