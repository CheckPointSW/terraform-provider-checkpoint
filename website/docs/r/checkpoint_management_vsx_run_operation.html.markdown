---
layout: "checkpoint"
page_title: "checkpoint_management_command_vsx_run_operation"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-vsx-run-operation"
description: |-
This resource allows you to execute Check Point Vsx Run Operation.
---

# Resource: checkpoint_management_command_vsx_run_operation

This resource allows you to execute Check Point Vsx Run Operation.

## Example Usage


```hcl
resource "checkpoint_management_command_vsx_run_operation" "example" {
  operation = "add-member"
}
```

## Argument Reference

The following arguments are supported:

* `operation` - (Required) The name of the operation to run. Each operation has its specific parameters.<br>The available operations are:<ul><li><i>upgrade</i> - Upgrades the VSX Gateway or VSX Cluster object to a higher version</li><li><i>downgrade</i> - Downgrades the VSX Gateway or VSX Cluster object to a lower version</li><li><i>add-member</i> - Adds a new VSX Cluster member object</li><li><i>remove-member</i> - Removes a VSX Cluster member object</li><li><i>reconf-gw</i> - Reconfigures a VSX Gateway after a clean install</li><li><i>reconf-member</i> - Reconfigures a VSX Cluster member after a clean install</li></ul>. 
* `add_member_params` - (Optional) Parameters for the operation to add a VSX Cluster member.add_member_params blocks are documented below.
* `downgrade_params` - (Optional) Parameters for the operation to downgrade a VSX Gateway or VSX Cluster object to a lower version.<br>In case the current version is already the target version, or is lower than the target version, no change is done.downgrade_params blocks are documented below.
* `reconf_gw_params` - (Optional) Parameters for the operation to reconfigure a VSX Gateway after a clean install.reconf_gw_params blocks are documented below.
* `reconf_member_params` - (Optional) Parameters for the operation to reconfigure a VSX Cluster member after a clean install.reconf_member_params blocks are documented below.
* `remove_member_params` - (Optional) Parameters for the operation to remove a VSX Cluster member object.remove_member_params blocks are documented below.
* `upgrade_params` - (Optional) Parameters for the operation to upgrade a VSX Gateway or VSX Cluster object to a higher version.<br>In case the current version is already the target version, or is higher than the target version, no change is done.upgrade_params blocks are documented below.


`add_member_params` supports the following:

* `ipv4_address` - (Optional) The IPv4 address of the management interface of the VSX Cluster member. 
* `ipv4_sync_address` - (Optional) The IPv4 address of the sync interface of the VSX Cluster member. 
* `member_name` - (Optional) Name of the new VSX Cluster member object. 
* `vsx_name` - (Optional) Name of the VSX Cluster object. 
* `vsx_uid` - (Optional) UID of the VSX Cluster object. 


`downgrade_params` supports the following:

* `target_version` - (Optional) The target version. 
* `vsx_name` - (Optional) Name of the VSX Gateway or VSX Cluster object. 
* `vsx_uid` - (Optional) UID of the VSX Gateway or VSX Cluster object. 


`reconf_gw_params` supports the following:

* `ipv4_corexl_number` - (Optional) Number of IPv4 CoreXL Firewall instances on the target VSX Gateway.<br>Valid values:<br><ul><li>To configure CoreXL Firewall instances, enter an integer greater or equal to 2.</li><li>To disable CoreXL, enter 1.</li></ul>. 
* `one_time_password` - (Optional) A password required for establishing a Secure Internal Communication (SIC). Enter the same password you used during the First Time Configuration Wizard on the target VSX Gateway. 
* `vsx_name` - (Optional) Name of the VSX Gateway object. 
* `vsx_uid` - (Optional) UID of the VSX Gateway object. 


`reconf_member_params` supports the following:

* `ipv4_corexl_number` - (Optional) Number of IPv4 CoreXL Firewall instances on the target VSX Cluster member.<br>Valid values:<br><ul><li>To configure CoreXL Firewall instances, enter an integer greater or equal to 2.</li><li>To disable CoreXL, enter 1.</li></ul>Important - The CoreXL configuration must be the same on all the cluster members. 
* `member_uid` - (Optional) UID of the VSX Cluster member object. 
* `member_name` - (Optional) Name of the VSX Cluster member object. 
* `one_time_password` - (Optional) A password required for establishing a Secure Internal Communication (SIC). Enter the same password you used during the First Time Configuration Wizard on the target VSX Cluster member. 


`remove_member_params` supports the following:

* `member_uid` - (Optional) UID of the VSX Cluster member object. 
* `member_name` - (Optional) Name of the VSX Cluster member object. 


`upgrade_params` supports the following:

* `target_version` - (Optional) The target version. 
* `vsx_name` - (Optional) Name of the VSX Gateway or VSX Cluster object. 
* `vsx_uid` - (Optional) UID of the VSX Gateway or VSX Cluster object. 



Response:
* `task_id` - Operation task UID. Use the show-task command to check the progress of the task.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

