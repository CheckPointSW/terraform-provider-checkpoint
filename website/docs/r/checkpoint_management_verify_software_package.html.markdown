---
layout: "checkpoint"
page_title: "checkpoint_management_verify_software_package"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-verify-software-package"
description: |-
  This resource allows you to execute Check Point Verify Software Package.
---

# Resource: checkpoint_management_verify_software_package

This command resource allows you to execute Check Point Verify Software Package.

## Example Usage


```hcl
resource "checkpoint_management_verify_software_package" "example" {
  name = "Check_Point_R80_40_JHF_MCD_DEMO_019_MAIN_Bundle_T1_VISIBLE_FULL.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the software package. 
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.
* `concurrency_limit` - (Optional) The number of targets, on which the same package is installed at the same time. 
* `task_id` - (Computed) Asynchronous task unique identifier. 
* `cluster_installation_settings` - (Optional) Installation settings for cluster.cluster_installation_settings blocks are documented below.
* `download_package` - (Optional) Should the package be downloaded before verification.
* `download_package_from` - (Optional) Where is the package located.
* `operation_context` - (Optional) The operation can be: 'install' (default) or 'uninstall'.

`cluster_installation_settings` supports the following:

* `cluster_delay` - (Optional) The delay between end of installation on one cluster members and start of installation on the next cluster member.
* `cluster_strategy` - (Optional) The cluster installation strategy. all-members - Install the package on all members in the cluster non-active-members-and-failover - In the High Avail...


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  
