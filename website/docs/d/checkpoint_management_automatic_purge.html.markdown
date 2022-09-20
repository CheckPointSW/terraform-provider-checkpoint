---
layout: "checkpoint"
page_title: "checkpoint_management_automatic_purge"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-automatic-purge"
description: |-
Use this data source to get information on an existing Check Point Automatic Purge.
---

# Data Source: checkpoint_management_automatic_purge

Use this data source to get information on an existing Check Point Automatic Purge.


## Example Usage


```hcl
data "checkpoint_management_automatic_purge" "data_automatic_purge" {
   
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - Turn on/off the automatic-purge feature.
* `keep_sessions_by_count` - Whether or not to keep the latest N sessions.
* `number_of_sessions_to_keep` - The number of newest sessions to preserve, by the sessions's publish date.
* `keep_sessions_by_days` - Whether or not to keep the sessions for D days.
* `number_of_days_to_keep` - When "keep-sessions-by-days = true" this sets the number of days to keep the sessions.
* `scheduling` - When to purge sessions that do not meet the "keep" criteria. scheduling is type List. scheduling blocks are documented below.

`scheduling` supports the following:

* `start_date` - The first time to check whether or not there are sessions to purge.
* `time_units` - The time units.
* `check_interval` - Number of time-units between two purge checks.
* `last_check` - Last time purge check was executed.
* `next_check` - Next time purge check will be executed.