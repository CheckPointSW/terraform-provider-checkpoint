---
layout: "checkpoint"
page_title: "checkpoint_management_smart_task_trigger"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-smart_task_trigger"
description: |-
Use this data source to get information on an existing Check Point Smart Task Trigger.
---

# Data Source: checkpoint_management_smart_task_trigger

Use this data source to get information on an existing Check Point Smart Task Trigger.

## Example Usage


```hcl
data "checkpoint_management_data_host" "data_host" {
    name = "Example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `before_operation` - Whether or not this trigger is fired before an operation.