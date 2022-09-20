---
layout: "checkpoint"
page_title: "checkpoint_management_command_export_management"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-export-management"
description: |-
This resource allows you to execute Check Point Export Management.
---

# Resource: checkpoint_management_command_export_management

This resource allows you to execute Check Point Export Management.

## Example Usage


```hcl
resource "checkpoint_management_command_export_management" "example" {
  domain_name = "domain1"
  file_path = "/var/log/domain1_backup.tgz"
  is_domain_backup = true
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path in which the exported database file is saved.<br><font color="red">Required only</font> when not using pre-export-verification-only flag. 
* `domain_name` - (Required) Domain name to be exported.<br><font color="red">Required only for</font> exporting a Domain from the Multi-Domain Server or backing up Domain. 
* `version` - (Optional) Target version. 
* `include_logs` - (Optional) Export logs without log indexes. 
* `include_logs_indexes` - (Optional) Export logs with log indexes. 
* `include_endpoint_configuration` - (Optional) Include export of the Endpoint Security Management configuration files. 
* `include_endpoint_database` - (Optional) Include export of the Endpoint Security Management database. 
* `is_domain_backup` - (Optional) If true, the exported Domain will be suitable for import on the same Multi-Domain Server only. 
* `is_smc_to_mds` - (Optional) If true, the exported Security Management Server will be suitable for import on the Multi-Domain Server only. 
* `pre_export_verification_only` - (Optional) If true, only runs the pre-export verifications instead of the full export. 
* `ignore_warnings` - (Optional) Ignoring the verification warnings. By Setting this parameter to 'true' export will not be blocked by warnings. 
* `task_id` - Asynchronous task unique identifier. Use show-task command to check the progress of the task.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

