---
layout: "checkpoint"
page_title: "checkpoint_gaia_system_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-system-group"
description: |-
This resource allows you to execute Check Point System Group.
---

# checkpoint_gaia_system_group

This resource allows you to execute Check Point System Group.

## Example Usage


```hcl
resource "checkpoint_gaia_system_group" "example" {
  name = "financeGroup"
  gid = 2000
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required)  
* `gid` - (Required) Numeric ID which is used in identifying a group; it must be unique 
* `users` - (Optional) New users to be added to a group. Users, as well as the group, must exist. users blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use the existing resource's `name` as the import ID.

```shell
terraform import checkpoint_gaia_system_group.example 'myname'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

```shell
terraform import checkpoint_gaia_system_group.example '{"name":"myname","member_id":"1_1"}'
```

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
