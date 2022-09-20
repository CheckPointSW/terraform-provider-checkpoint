---
layout: "checkpoint"
page_title: "checkpoint_management_command_import_management"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-import-management"
description: |-
This resource allows you to execute Check Point Import Management.
---

# Resource: checkpoint_management_command_import_management

This resource allows you to execute Check Point Import Management.

## Example Usage


```hcl
resource "checkpoint_management_command_import_management" "example" {
  file_path = "/var/log/domain1_exported.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path to the exported database file to be imported. 
* `domain_name` - (Required) Domain name to be imported. Must be unique in the Multi-Domain Server.<br><font color="red">Required only for</font> importing the Security Management Server into the Multi-Domain Server. 
* `domain_ip_address` - (Required) IPv4 address for the imported Domain.<br><font color="red">Required only for</font> importing the Security Management Server into the Multi-Domain Server. 
* `domain_server_name` - (Required) Multi-Domain Server name for the imported Domain.<br><font color="red">Required only for</font> importing the Security Management Server into the Multi-Domain Server. 
* `include_logs` - (Optional) Import logs without log indexes. 
* `include_logs_indexes` - (Optional) Import logs with log indexes. 
* `include_endpoint_configuration` - (Optional) Include import of the Endpoint Security Management configuration files. 
* `include_endpoint_database` - (Optional) Include import of the Endpoint Security Management database. 
* `verify_domain_restore` - (Optional) If true, verify that the restore operation is valid for this input file and this environment. <br>Note: Restore operation will not be executed. 
* `pre_import_verification_only` - (Optional) If true, only runs the pre-import verifications instead of the full import. 
* `ignore_warnings` - (Optional) Ignoring the verification warnings. By Setting this parameter to 'true' import will not be blocked by warnings. 
* `task_id` - Asynchronous task unique identifier. Use show-task command to check the progress of the task.
* `login_required` - If set to "True", session is expired and login is required.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

