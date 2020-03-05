---
layout: "checkpoint"
page_title: "checkpoint_management_set_ips_update_schedule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-ips-update-schedule"
description: |-
This resource allows you to execute Check Point Set Ips Update Schedule.
---

# checkpoint_management_set_ips_update_schedule

This resource allows you to execute Check Point Set Ips Update Schedule.

## Example Usage


```hcl
resource "checkpoint_management_set_ips_update_schedule" "example" {
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Enable/Disable IPS Update Schedule. 
* `time` - (Optional) Time in format HH:mm. 
* `recurrence` - (Optional) Days recurrence.recurrence blocks are documented below.


`recurrence` supports the following:

* `days` - (Optional) Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].days blocks are documented below.
* `minutes` - (Optional) Valid on interval. The length of time in minutes between updates. 
* `pattern` - (Optional) Valid on "Interval", "Daily", "Weekly", "Monthly" base. 
* `weekdays` - (Optional) Valid on weekdays. Example: "Sun", "Mon"..."Sat".weekdays blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

