---
layout: "checkpoint"
page_title: "checkpoint_gaia_radius"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-radius"
description: |-
This resource allows you to execute Check Point Radius.
---

# checkpoint_gaia_radius

This resource allows you to execute Check Point Radius.

## Example Usage


```hcl
resource "checkpoint_gaia_radius" "example" {
  enabled        = true
  default_shell  = "cli"
  super_user_uid = "96"
  servers {
    priority = 3
    address  = "1.2.1.2"
    port     = 1812
    timeout  = 3
    secret   = "mySecret"
  }
}
```

## Argument Reference

The following arguments are supported:

* `nas_ip` - (Optional) The NAS-IP for the RADIUS client 
* `default_shell` - (Optional) Default shell when login 
* `super_user_uid` - (Optional) The UID that will be given to a super user 
* `servers` - (Optional) RADIUS servers list servers blocks are documented below.
* `enabled` - (Optional) RADIUS authentication state 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`servers` supports the following:

* `priority` - (Optional) Server priority (lower values comes first) 
* `address` - (Optional) Server address 
* `port` - (Optional) UDP port to contact on the RADIUS server 
* `timeout` - (Optional)  
* `secret` - (Optional)  
