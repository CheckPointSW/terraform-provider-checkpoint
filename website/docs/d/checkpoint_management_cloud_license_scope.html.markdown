---
layout: "checkpoint"
page_title: "checkpoint_management_cloud_license_scope"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cloud-license-scope"
description: |-
Use this data source to get information on an existing Check Point Cloud License Scope.
---

# checkpoint_management_cloud_license_scope

Use this data source to get information on an existing Check Point Cloud License Scope.

## Example Usage


```hcl
data "checkpoint_management_cloud_license_scope" "scope" {
}
```

## Argument Reference

The following arguments are supported:

* `task_id` - (Computed) update-cloud-license task UID. Use show-task command to check the progress of the task.
