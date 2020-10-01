---
layout: "checkpoint"
page_title: "checkpoint_management_set_automatic_purge"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-automatic-purge"
description: |-
This resource allows you to execute Check Point Set Automatic Purge.
---

# Resource: checkpoint_management_set_automatic_purge

This command resource allows you to execute Check Point Set Automatic Purge.

## Example Usage


```hcl
resource "checkpoint_management_set_automatic_purge" "example" {
  enabled = true
  keep_sessions_by_days = false
  number_of_sessions_to_keep = 10
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Required) Turn on/off the automatic-purge feature. 
* `keep_sessions_by_count` - (Optional) Whether or not to keep the latest N sessions.
Note: when the automatic purge feature is enabled, this field and/or the "keep-sessions-by-date" field must be set to 'true'. 
* `number_of_sessions_to_keep` - (Optional) When "keep-sessions-by-count = true" this sets the number of newest sessions to preserve, by the sessions's publish date. 
* `keep_sessions_by_days` - (Optional) Whether or not to keep the sessions for D days.
Note: when the automatic purge feature is enabled, this field and/or the "keep-sessions-by-count" field must be set to 'true'. 
* `number_of_days_to_keep` - (Optional) When "keep-sessions-by-days = true" this sets the number of days to keep the sessions. 
* `scheduling` - (Optional) When to purge sessions that do not meet the "keep" criteria. Note: when the automatic purge feature is enabled, this field must be set.scheduling blocks are documented below.


`scheduling` supports the following:

* `start_date` - (Optional) The first time to check whether or not there are sessions to purge. ISO 8601. If timezone isn't specified in the input, the Management server's timezone is used. Instead - If you want to start immediately, type: "now". Note: when the automatic purge feature is enabled, this field must be set. 
* `time_units` - (Optional) Note: when the automatic purge feature is enabled, this field must be set. 
* `check_interval` - (Optional) Number of time-units between two purge checks.  Note: when the automatic purge feature is enabled, this field must be set. 


## How To Use
Make sure this command will be executed in the right execution order.