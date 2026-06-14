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
