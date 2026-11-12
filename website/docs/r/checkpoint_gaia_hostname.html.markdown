---
layout: "checkpoint"
page_title: "checkpoint_gaia_hostname"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-hostname"
description: |-
This resource allows you to execute Check Point Hostname.
---

# checkpoint_gaia_hostname

This resource allows you to execute Check Point Hostname.

## Example Usage


```hcl
resource "checkpoint_gaia_hostname" "example" {
  name = "new-hostname"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Hostname can be a combination of letters and numbers, it cannot be in IP format or start/end with characters such as '.' And '-'  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_hostname.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
