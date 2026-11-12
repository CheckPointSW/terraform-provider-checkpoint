---
layout: "checkpoint"
page_title: "checkpoint_management_ips_update_schedule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-ips-update-schedule"
description: |-
Use this data source to get information on an existing Check Point Ips Update Schedule.
---

# Data Source: checkpoint_management_ips_update_schedule

Use this data source to get information on an existing Check Point Ips Update Schedule.

## Example Usage


```hcl
data "checkpoint_management_ips_update_schedule" "data_ips_update_schedule" {

}
```

## Argument Reference

The following arguments are supported:

* `enabled` - IPS Update Schedule status.
* `time` - Time in format HH:mm.
* `recurrence` - Days recurrence. recurrence blocks are documented below.

`recurrence` supports the following:

* `days` - Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].
* `minutes` - Valid on interval. The length of time in minutes between updates.
* `pattern` - Valid on "Interval", "Daily", "Weekly", "Monthly" base.
* `weekdays` - Valid on weekdays. Example: "Sun", "Mon"..."Sat".