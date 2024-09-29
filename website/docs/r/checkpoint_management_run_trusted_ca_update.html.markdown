---
layout: "checkpoint"
page_title: "checkpoint_management_run_trusted_ca_update"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-run-trusted-ca-update"
description: |-
This resource allows you to execute Check Point Run Trusted Ca Update.
---

# checkpoint_management_run_trusted_ca_update

This resource allows you to execute Check Point Run Trusted Ca Update.

## Example Usage
```hcl
resource "checkpoint_management_run_trusted_ca_update" "test" {
  
}
```

## Argument Reference

The following arguments are supported:

* `package_path` - (Optional) Path on the management server for offline Trusted CAs package update. 
* `task_id` - (Computed) Asynchronous task unique identifier.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

