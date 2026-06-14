---
layout: "checkpoint"
page_title: "checkpoint_gaia_gaia_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-gaia-interface"
description: |-
This resource allows you to execute Check Point Management Interface.
---

# checkpoint_gaia_gaia_interface

This resource allows you to execute Check Point Management Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_gaia_interface" "example" {
  name = "eth1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Interface name 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
