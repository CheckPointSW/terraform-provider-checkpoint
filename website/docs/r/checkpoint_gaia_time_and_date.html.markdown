---
layout: "checkpoint"
page_title: "checkpoint_gaia_time_and_date"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-time-and-date"
description: |-
This resource allows you to execute Check Point Time And Date.
---

# checkpoint_gaia_time_and_date

This resource allows you to execute Check Point Time And Date.

## Example Usage


```hcl
resource "checkpoint_gaia_time_and_date" "example" {
  timezone = "Asia / Jerusalem"
}
```

## Argument Reference

The following arguments are supported:

* `time` - (Optional) Time to set, in HH:MM[:SS] format 
* `timezone` - (Optional) Timezone in Area / Region format. See timezones list via 'show-timezones' 
* `date` - (Optional) Date to set, in YYYY-MM-DD format 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_time_and_date.example 'settings'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
