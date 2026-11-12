---
layout: "checkpoint"
page_title: "checkpoint_management_task"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-task"
description: |-
Use this data source to get information on an existing Check Point Task.
---

# Data Source: checkpoint_management_task

Use this data source to get information on an existing Check Point Task.

## Example Usage


```hcl
data "checkpoint_management_task" "example" {
  task_id = ["6682b963-fe1a-4a75-a86c-91cb13e91d83"]
}
```

## Argument Reference

The following arguments are supported:

* `task_id` - (Required) Unique identifier of one or more tasks.
* `tasks` - The tasks. tasks blocks are documented below.
* `response` - Response message in JSON format.

`tasks` supports the following:

* `task_id` - Asynchronous task unique identifier. Use show-task command to check the progress of the task.
* `task_name` - The task name.
* `status` - Task status.
* `progress_percentage` - The progress percentage of the task.
* `suppressed` - Is the task suppressed.
* `comments` - Comments string.
* `last_update_time` - N/A.last_update_time blocks are documented below.
* `previous_task_id` - ID of the previous task in the execution sequence, on which this task depends.
* `progress_description` - N/A.
* `revert_status` - N/A.
* `start_time` - N/A.start_time blocks are documented below.


`last_update_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`start_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
