---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_remote_syslogs"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-remote-syslogs"
description: |-
This resource allows you to execute Check Point Show Remote Syslogs.
---

# checkpoint_gaia_show_remote_syslogs

This resource allows you to execute Check Point Show Remote Syslogs.

## Example Usage


```hcl
data "checkpoint_gaia_show_remote_syslogs" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

