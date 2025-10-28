---
layout: "checkpoint"
page_title: "checkpoint_management_set_smart_console_idle_timeout"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-smart-console-idle-timeout"
description: |-
 This resource allows you to execute Check Point Set Smart Console Idle Timeout.
---

# checkpoint_management_set_smart_console_idle_timeout

This resource allows you to execute Check Point Set Smart Console Idle Timeout.

## Example Usage


```hcl
resource "checkpoint_management_set_smart_console_idle_timeout" "example" {
  enabled = true
  timeout_duration = 30
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Indicates whether to perform logout after being idle. 
* `timeout_duration` - (Optional) Number of minutes that the SmartConsole will automatically logout after being idle.<br>Updating the interval will take effect only on the next login. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

