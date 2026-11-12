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


## Import

Use the existing resource's `name` as the import ID.

```shell
terraform import checkpoint_gaia_role.example 'myname'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

```shell
terraform import checkpoint_gaia_role.example '{"name":"myname","member_id":"1_1"}'
```

IDs beginning with `{` must use the JSON-object form; JSON syntax is reserved for structured import IDs.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
