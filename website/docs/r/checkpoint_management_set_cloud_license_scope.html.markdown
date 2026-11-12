---
layout: "checkpoint"
page_title: "checkpoint_management_set_cloud_license_scope"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-cloud-license-scope"
description: |-
This resource allows you to execute Check Point Set Cloud License Scope.
---

# checkpoint_management_set_cloud_license_scope

This resource allows you to execute Check Point Set Cloud License Scope.

## Example Usage


```hcl
resource "checkpoint_management_set_cloud_license_scope" "example" {
  mode = "mds"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - (Required) Set cloud license scope mode.
* `task_id` - (Computed) Asynchronous task unique identifier.
