---
layout: "checkpoint"
page_title: "checkpoint_management_scheduled_event"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-scheduled-event"
description: |-
Use this data source to get information on an existing Check Point Scheduled Event.
---

# Data Source: checkpoint_management_scheduled_event

Use this data source to get information on an existing Check Point Scheduled Event.

## Example Usage

```hcl
data "checkpoint_management_scheduled_event" "data_test" {
    name = "scheduled-event1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `schedule` - (Computed) Schedule Configuration. schedule blocks are documented below.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `comments` - (Computed) Comments string.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`schedule` supports the following:

* `time` - (Computed) Time in format HH:mm.
* `recurrence` - (Computed) Days recurrence. recurrence blocks are documented below.


`recurrence` supports the following:

* `pattern` - (Computed) Days recurrence pattern.
* `interval_hours` - (Computed) The amount of hours between updates. <font color="red">Required only when</font> pattern is set to 'Interval'.
* `interval_minutes` - (Computed) The amount of minutes between updates. <font color="red">Required only when</font> pattern is set to 'Interval'.
* `interval_seconds` - (Computed) The amount of seconds between updates. <font color="red">Required only when</font> pattern is set to 'Interval'.
* `weekdays` - (Computed) Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color="red">Required only when</font> pattern is set to 'Weekly'.
* `days` - (Computed) Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31. <font color="red">Required only when</font> pattern is set to 'Monthly'.
