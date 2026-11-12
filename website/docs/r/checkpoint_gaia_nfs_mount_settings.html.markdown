---
layout: "checkpoint"
page_title: "checkpoint_gaia_nfs_mount_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-nfs-mount-settings"
description: |-
This resource allows you to execute Check Point Nfs Mount Settings.
---

# checkpoint_gaia_nfs_mount_settings

This resource allows you to execute Check Point Nfs Mount Settings.

## Example Usage


```hcl
resource "checkpoint_gaia_nfs_mount_settings" "example" {
  timeout = 5
}
```

## Argument Reference

The following arguments are supported:

* `timeout` - (Optional) Nfs timeout in seconds. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_nfs_mount_settings.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
