---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_configuration_ipv6_pim"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-configuration-ipv6-pim"
description: |-
This resource allows you to execute Check Point Show Configuration Ipv6 Pim.
---

# checkpoint_gaia_show_configuration_ipv6_pim

This resource allows you to execute Check Point Show Configuration Ipv6 Pim.

## Example Usage


```hcl
data "checkpoint_gaia_show_configuration_ipv6_pim" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

