---
layout: "checkpoint"
page_title: "checkpoint_gaia_scheduled_job"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-scheduled-job"
description: |-
This resource allows you to execute Check Point Scheduled Job.
---

# checkpoint_gaia_scheduled_job

This resource allows you to execute Check Point Scheduled Job.

## Example Usage


```hcl
resource "checkpoint_gaia_scheduled_job" "example" {
  name    = "startup_job"
  command = "/home/admin/new_job.sh"

  recurrence {
    type = "system-startup"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Scheduled job name 
* `command` - (Required) Scheduled command (expert CLI style) 
* `recurrence` - (Required) Recurrence schedule recurrence blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`recurrence` supports the following:

* `type` - (Optional) Job recurrence type 
* `interval` - (Optional) Time interval in minutes. Relevant for "interval" recurrence type 
* `time_of_day` - (Optional) Time of day in 24 hour format. Relevant for "daily", "weekly" and "monthly" recurrence types time_of_day blocks are documented below.
* `hourly` - (Optional) Hours of day in 24 hour format. Can choose multiple hours. Relevant for "hourly" recurrence type hourly blocks are documented below.
* `weekdays` - (Optional) Days of the week. Relevant for "weekly" recurrence type weekdays blocks are documented below.
* `days` - (Optional) Days of the month. Relevant for "monthly" recurrence type days blocks are documented below.
* `months` - (Optional) Month numbers. Relevant for "monthly" recurrence type months blocks are documented below.


`time_of_day` supports the following:

* `hour` - (Optional) Time hour 
* `minute` - (Optional) Time minute 


`hourly` supports the following:

* `hours_of_day` - (Optional) Hours of day in 24 hour format hours_of_day blocks are documented below.
* `minute` - (Optional) Time minute 
