---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_mld_summary"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-mld-summary"
description: |-
This resource allows you to execute Check Point Show Mld Summary.
---

# checkpoint_gaia_show_mld_summary

This resource allows you to execute Check Point Show Mld Summary.

## Example Usage


```hcl
data "checkpoint_gaia_show_mld_summary" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

