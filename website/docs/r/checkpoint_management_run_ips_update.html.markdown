---
layout: "checkpoint"
page_title: "checkpoint_management_run_ips_update"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-run-ips-update"
description: |-
  Runs IPS database update. If "package-path" is not provided server will try to get the latest package from the User Center.
---

# Resource: checkpoint_management_run_ips_update

This command resource allows you to Runs IPS database update. If "package-path" is not provided server will try to get the latest package from the User Center.

## Example Usage

```hcl
resource "checkpoint_management_run_ips_update" "example" {}
```

## Argument Reference

The following arguments are supported:

* `package_path` - (Optional) Offline update package path.
* `task_id` - (Computed) Asynchronous task unique identifier. 

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.    



