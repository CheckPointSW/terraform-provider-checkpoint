---
layout: "checkpoint"
page_title: "checkpoint_management_time"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-time"
description: |-
  This resource allows you to execute Check Point Time.
---

# Resource: checkpoint_management_application_site

This resource allows you to execute Check Point Time.

## Example Usage


```hcl
resource "checkpoint_management_application_site" "example" {
  name = "time1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Time object name. Cannot be more than 11 characters. Should be unique in the domain.
* `end` - (Optional) End time. Note: Each gateway may interpret this time differently according to its time zone.
* `end_never` - (Optional) End never.
* `hours_ranges` - (Optional) Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone.
* `start` - (Optional) Starting time. Note: Each gateway may interpret this time differently according to its time zone.
* `start_now` - (Optional) Start immediately.
* `tags` - (Optional) Collection of tag identifiers.
* `recurrence` - (Optional) Days recurrence.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

`end` supports the following:

* `date` - (Optional) Date in format dd-MMM-yyyy.
* `iso_8601` - (Optional) Date and time represented in international ISO 8601 format. Time zone information is ignored.
* `posix` - (Optional) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
* `time` - (Optional) Time in format HH:mm.

`start` supports the following:

* `date` - (Optional) Date in format dd-MMM-yyyy.
* `iso_8601` - (Optional) Date and time represented in international ISO 8601 format. Time zone information is ignored.
* `posix` - (Optional) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
* `time` - (Optional) Time in format HH:mm.

`recurrence` supports the following:

* `days` - (Optional) Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].
* `month` - (Optional) Valid on month. Example: "1", "2","12","Any".
* `pattern` - (Optional) Valid on "Daily", "Weekly", "Monthly" base.
* `weekday` - (Optional) Valid on weekdays. Example: "Sun", "Mon"..."Sat".

`hours_ranges` supports the following:

* `enabled` - (Optional) Is hour range enabled.
* `from` - (Optional) Time in format HH:MM.
* `index` - (Optional) Hour range index.
* `to` - (Optional) Time in format HH:MM.