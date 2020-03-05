---
layout: "checkpoint"
page_title: "checkpoint_management_uninstall_software_package"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-uninstall-software-package"
description: |-
This resource allows you to execute Check Point Uninstall Software Package.
---

# checkpoint_management_uninstall_software_package

This resource allows you to execute Check Point Uninstall Software Package.

## Example Usage


```hcl
resource "checkpoint_management_uninstall_software_package" "example" {
  name = "Check_Point_R80_40_JHF_MCD_DEMO_019_MAIN_Bundle_T1_VISIBLE_FULL.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the software package. 
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.
* `cluster_installation_settings` - (Optional) Installation settings for cluster.cluster_installation_settings blocks are documented below.
* `concurrency_limit` - (Optional) The number of targets, on which the same package is installed at the same time. 


`cluster_installation_settings` supports the following:

* `cluster_delay` - (Optional) The delay between end of installation on one cluster members and start of installation on the next cluster member. 
* `cluster_strategy` - (Optional) The cluster installation strategy. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

