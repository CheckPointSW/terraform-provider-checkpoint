---
layout: "checkpoint"
page_title: "checkpoint_gaia_tacacs"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-tacacs"
description: |-
This resource allows you to execute Check Point Tacacs.
---

# checkpoint_gaia_tacacs

This resource allows you to execute Check Point Tacacs.

## Example Usage


```hcl
resource "checkpoint_gaia_tacacs" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) TACACS+ authentication state 
* `super_user_uid` - (Optional) The UID that will be given to a TACACS+ user 
* `servers` - (Optional) TACACS+ servers list servers blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`servers` supports the following:

* `priority` - (Optional) Server priority (lower values comes first) 
* `address` - (Optional) The server address 
* `timeout` - (Optional)  
* `secret` - (Optional)  


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_tacacs.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
