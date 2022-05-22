---
layout: "checkpoint"
page_title: "checkpoint_management_md_permissions_profile"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-md-permissions-profile"
description: |-
This resource allows you to execute Check Point Md Permissions Profile.
---

# checkpoint_management_md_permissions_profile

This resource allows you to execute Check Point Md Permissions Profile.

## Example Usage


```hcl
resource "checkpoint_management_md_permissions_profile" "example" {
  name = "manager profile"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `mds_provisioning` - (Optional) Create and manage Multi-Domain Servers and Multi-Domain Log Servers.<br>Only a "Super User" permission-level profile can select this option. 
* `manage_admins` - (Optional) Create and manage Multi-Domain Security Management administrators with the same or lower permission level. For example, a Domain manager cannot create Superusers or global managers.<br>Only a 'Manager' permission-level profile can edit this permission. 
* `manage_sessions` - (Optional) Connect/disconnect Domain sessions, publish changes, and delete other administrator sessions.<br>Only a 'Manager' permission-level profile can edit this permission. 
* `management_api_login` - (Optional) Permission to log in to the Security Management Server and run API commands using these tools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole. 
* `cme_operations` - (Optional) Permission to read / edit the Cloud Management Extension (CME) configuration. 
* `global_vpn_management` - (Optional) Lets the administrator select Enable global use for a Security Gateway shown in the MDS Gateways & Servers view.<br>Only a 'Manager' permission-level profile can edit this permission. 
* `manage_global_assignments` - (Optional) Controls the ability to create, edit and delete global assignment and not the ability to reassign, which is set according to the specific Domain's permission profile. 
* `enable_default_profile_for_global_domains` - (Optional) Enable the option to specify a default profile for all global domains. 
* `default_profile_global_domains` - (Optional) Name or UID of the required default profile for all global domains. 
* `view_global_objects_in_domain` - (Optional) Lets an administrator with no global objects permissions view the global objects in the domain. This option is required for valid domain management. 
* `enable_default_profile_for_local_domains` - (Optional) Enable the option to specify a default profile for all local domains. 
* `default_profile_local_domains` - (Optional) Name or UID of the required default profile for all local domains. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `permission_level` - (Optional) The level of the Multi Domain Permissions Profile.<br>The level cannot be changed after creation. 
