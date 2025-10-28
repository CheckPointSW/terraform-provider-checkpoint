---
layout: "checkpoint"
page_title: "checkpoint_management_smart_console_idle_timeout"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-smart-console-idle-timeout"
description: |-
  Use this data source to get information on an existing Check Point Smart Console Idle Timeout.
---

# Data Source: checkpoint_management_smart_console_idle_timeout

Use this data source to get information on an existing Check Point Smart Console Idle Timeout.

## Example Usage
```hcl
data "checkpoint_management_smart_console_idle_timeout" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - Indicates whether to perform logout after being idle. 
* `timeout_duration` - Number of minutes that the SmartConsole will automatically logout after being idle.<br>Updating the interval will take effect only on the next login. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.

