---
layout: "checkpoint"
page_title: "checkpoint_management_set_app_control_update_schedule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-app-control-update-schedule"
description: |-
This resource allows you to execute Check Point Set App Control Update Schedule.
---

# checkpoint_management_set_app_control_update_schedule

This resource allows you to execute Check Point Set App Control Update Schedule.

## Example Usage


```hcl
resource "checkpoint_management_set_app_control_update_schedule" "example" {
  schedule_gateway_update {
    schedule {
      recurrence {
        pattern = "interval"
        interval_hours = 4
        interval_minutes = 30
        interval_seconds = 10
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `schedule_management_update` - (Optional) Application Control & URL Filtering Update Schedule on Management Server.schedule_management_update blocks are documented below.
* `schedule_gateway_update` - (Optional) Application Control & URL Filtering Update Schedule on Gateway.schedule_gateway_update blocks are documented below.


`schedule_management_update` supports the following:

* `enabled` - (Optional) Enable/Disable Application Control & URL Filtering Update Schedule on Management Server. 
* `schedule` - (Optional) Schedule Configuration.schedule blocks are documented below.


`schedule_gateway_update` supports the following:

* `enabled` - (Optional) Enable/Disable Application Control & URL Filtering Update Schedule on Gateway. 
* `schedule` - (Optional) Schedule Configuration.schedule blocks are documented below.


`schedule` supports the following:

* `time` - (Optional) Time in format HH:mm. 
* `recurrence` - (Optional) Days recurrence.recurrence blocks are documented below.


`schedule` supports the following:

* `time` - (Optional) Time in format HH:mm. 
* `recurrence` - (Optional) Days recurrence.recurrence blocks are documented below.


`recurrence` supports the following:

* `pattern` - (Optional) Days recurrence pattern. 
* `weekdays` - (Optional) Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color="red">Required only when</font> pattern is set to 'Weekly'.
* `days` - (Optional) Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31. <font color="red">Required only when</font> pattern is set to 'Monthly'.


`recurrence` supports the following:

* `pattern` - (Optional) Days recurrence pattern. 
* `interval_hours` - (Optional) The amount of hours between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `interval_minutes` - (Optional) The amount of minutes between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `interval_seconds` - (Optional) The amount of seconds between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `weekdays` - (Optional) Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color="red">Required only when</font> pattern is set to 'Weekly'.
* `days` - (Optional) Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31. <font color="red">Required only when</font> pattern is set to 'Monthly'.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

