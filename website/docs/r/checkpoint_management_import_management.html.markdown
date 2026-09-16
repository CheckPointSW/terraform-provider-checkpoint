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
* `cancel_prepare_import` - (Optional) If 'true', cancels the import in background process. If you do not run this command within the number of days defined in 'show-background-upgrade-sett...
* `change_ips` - (Optional) New IP addresses (IPv4, IPv6, or both) of the servers.<br><font color='red'>Required only if</font> one or more of the servers in the Security Managem...change_ips blocks are documented below.
* `complete_background_import` - (Optional) If 'true', import the changes you made during the 'Prepare' phase, and the 'Complete' phase will be achieved. You can't make any changes during the 'C...
* `days_of_logs` - (Optional) Export <N> last days of logs.
* `domain_ipv6_address` - (Optional) IPv6 address for the imported Domain.
* `keep_cloud_sharing` - (Optional) Preserve the connection of the Management Server to Check Point's Infinity Portal.<br>Use this flag after ensuring that the original Management Server...
* `prepare_background_import` - (Optional) If 'true', the import will run in the background and 'Prepare' phase will be achieved. You can continue making changes on the Management Server during...


`change_ips` supports the following:

* `new_ipv4_address` - (Optional) The new IPv4 address of the server that migrates to a new IP address.
* `server_name` - (Optional) The object name of the server that migrates to a new IP address.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  
