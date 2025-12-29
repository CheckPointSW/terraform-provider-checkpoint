---
layout: "checkpoint"
page_title: "checkpoint_management_app_control_update_schedule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-app-control-update-schedule"
description: |-
Use this data source to get information on an existing Check Point App Control update schedule.
---

# checkpoint_management_app_control_update_schedule

Use this data source to get information on an existing Check Point App Control update schedule.

## Example Usage


```hcl
data "checkpoint_management_app_control_update_schedule" "data" {

}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object Identifier.
* `schedule_management_update` - Application Control & URL Filtering Update Schedule on Management Server.schedule_management_update blocks are documented below.
* `schedule_gateway_update` - Application Control & URL Filtering Update Schedule on Gateway.schedule_gateway_update blocks are documented below.


`schedule_management_update` supports the following:

* `enabled` - Enable/Disable Application Control & URL Filtering Update Schedule on Management Server.
* `schedule` - Schedule Configuration.schedule blocks are documented below.


`schedule_gateway_update` supports the following:

* `enabled` - Enable/Disable Application Control & URL Filtering Update Schedule on Gateway.
* `schedule` - Schedule Configuration.schedule blocks are documented below.


`schedule` of `schedule_management_update` supports the following:

* `time` - Time in format HH:mm.
* `recurrence` - Days recurrence. recurrence blocks are documented below.


`schedule` of `schedule_gateway_update` supports the following:

* `time` - Time in format HH:mm.
* `recurrence` - Days recurrence. recurrence blocks are documented below.


`recurrence` of `schedule_management_update` supports the following:

* `pattern` - Days recurrence pattern.
* `weekdays` - Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}.
* `days` - Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31.


`recurrence` of `schedule_gateway_update` supports the following:

* `pattern` - Days recurrence pattern.
* `interval_hours` - The amount of hours between updates.
* `interval_minutes` - The amount of minutes between updates.
* `interval_seconds` - The amount of seconds between updates.
* `weekdays` - Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}.
* `days` - Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

