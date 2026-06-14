---
layout: "checkpoint"
page_title: "checkpoint_gaia_role"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-role"
description: |-
This resource allows you to execute Check Point Role.
---

# checkpoint_gaia_role

This resource allows you to execute Check Point Role.

## Example Usage


```hcl
resource "checkpoint_gaia_role" "example" {
  name = "myrole"
  extended_commands = ["ifconfig",]
  features {
    name = "bgp"
    permission = "read-write"
  }
  features {
    name = "dhcp"
    permission = "read-write"
  }
  features {
    name = "igmp"
    permission = "read-write"
  }
  features {
    name = "ntp"
    permission = "read-write"
  }
  features {
    name = "syslog"
    permission = "read-write"
  }
  features {
    name = "backup"
    permission = "read-only"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Role name 
* `features` - (Optional) Specifies which features will be assigned to the role. features blocks are documented below.
* `extended_commands` - (Optional) Specifies which extended commands will be assigned to the role.Valid values: extended commands as shown in show-extended-commands API output or 'all' to specify all extended-commands.  extended_commands blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `users` - (Computed) Computed field, returned in the response. users blocks are documented below.


`features` supports the following:

* `name` - (Optional) Feature name. Valid values: feature name as shown in show-features API output or 'all' to specify all features.  
* `permission` - (Optional) Feature permission. Valid values: read-write ,read-only.  
