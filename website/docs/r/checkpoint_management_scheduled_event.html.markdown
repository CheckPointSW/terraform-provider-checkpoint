---
layout: "checkpoint"
page_title: "checkpoint_management_scheduled_event"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-scheduled-event"
description: |-
This resource allows you to execute Check Point Scheduled Event.
---

# checkpoint_management_scheduled_event

This resource allows you to execute Check Point Scheduled Event.

## Example Usage


```hcl
resource "checkpoint_management_scheduled_event" "example" {
  name = "Daily Event"
  comments = "Scheduled event for daily backup operations"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `schedule` - (Optional) Schedule Configuration.schedule blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`schedule` supports the following:

* `time` - (Optional) Time in format HH:mm. 
* `recurrence` - (Optional) Days recurrence.recurrence blocks are documented below.


`recurrence` supports the following:

* `pattern` - (Optional) Days recurrence pattern. 
* `interval_hours` - (Optional) The amount of hours between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `interval_minutes` - (Optional) The amount of minutes between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `interval_seconds` - (Optional) The amount of seconds between updates. <font color="red">Required only when</font> pattern is set to 'Interval'. 
* `weekdays` - (Optional) Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color="red">Required only when</font> pattern is set to 'Weekly'.weekdays blocks are documented below.
* `days` - (Optional) Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31. <font color="red">Required only when</font> pattern is set to 'Monthly'.days blocks are documented below.
