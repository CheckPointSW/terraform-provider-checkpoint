---
layout: "checkpoint"
page_title: "checkpoint_gaia_authentication_order"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-authentication-order"
description: |-
This resource allows you to execute Check Point Authentication Order.
---

# checkpoint_gaia_authentication_order

This resource allows you to execute Check Point Authentication Order.

## Example Usage


```hcl
resource "checkpoint_gaia_authentication_order" "example" {
  local {
    priority = 1
  }
  radius {
    priority = 2
    enabled  = true
  }
  tacacs {
    priority = 3
    enabled  = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `radius` - (Optional) Server type radius blocks are documented below.
* `tacacs` - (Optional) Server type tacacs blocks are documented below.
* `local` - (Optional) Server type local blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`radius` supports the following:

* `priority` - (Optional) Authentication priority 
* `enabled` - (Optional) Server state 


`tacacs` supports the following:

* `priority` - (Optional) Authentication priority 
* `enabled` - (Optional) Server state 


`local` supports the following:

* `priority` - (Optional) Authentication priority 
