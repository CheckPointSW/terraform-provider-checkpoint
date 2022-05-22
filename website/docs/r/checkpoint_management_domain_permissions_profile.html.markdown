---
layout: "checkpoint"
page_title: "checkpoint_management_domain_permissions_profile"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-domain-permissions-profile"
description: |-
This resource allows you to execute Check Point Domain Permissions Profile.
---

# checkpoint_management_domain_permissions_profile

This resource allows you to execute Check Point Domain Permissions Profile.

## Example Usage


```hcl
resource "checkpoint_management_domain_permissions_profile" "example" {
  name = "customize profile"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `permission_type` - (Optional) The type of the Permissions Profile. 
* `edit_common_objects` - (Optional) Define and manage objects in the Check Point database: Network Objects, Services, Custom Application Site, VPN Community, Users, Servers, Resources, Time, UserCheck, and Limit.<br>Only a 'Customized' permission-type profile can edit this permission. 
* `access_control` - (Optional) Access Control permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.access_control blocks are documented below.
* `endpoint` - (Optional) Endpoint permissions. Not supported for Multi-Domain Servers.<br>Only a 'Customized' permission-type profile can edit these permissions.endpoint blocks are documented below.
* `events_and_reports` - (Optional) Events and Reports permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.events_and_reports blocks are documented below.
* `gateways` - (Optional) Gateways permissions. <br>Only a 'Customized' permission-type profile can edit these permissions.gateways blocks are documented below.
* `management` - (Optional) Management permissions.management blocks are documented below.
* `monitoring_and_logging` - (Optional) Monitoring and Logging permissions.<br>'Customized' permission-type profile can edit all these permissions. "Read Write All" permission-type can edit only dlp-logs-including-confidential-fields and manage-dlp-messages permissions.monitoring_and_logging blocks are documented below.
* `threat_prevention` - (Optional) Threat Prevention permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.threat_prevention blocks are documented below.
* `others` - (Optional) Additional permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.others blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`access_control` supports the following:

* `show_policy` - (Optional) Select to let administrators work with Access Control rules and NAT rules. If not selected, administrators cannot see these rules. 
* `policy_layers` - (Optional) Layer editing permissions.<br>Available only if show-policy is set to true.policy_layers blocks are documented below.
* `dlp_policy` - (Optional) Configure DLP rules and Policies. 
* `geo_control_policy` - (Optional) Work with Access Control rules that control traffic to and from specified countries. 
* `nat_policy` - (Optional) Work with NAT in Access Control rules. 
* `qos_policy` - (Optional) Work with QoS Policies and rules. 
* `access_control_objects_and_settings` - (Optional) Allow editing of the following objet types: VPN Community, Access Role, Custom application group,Custom application, Custom category, Limit, Application - Match Settings, Application Category - Match Settings,Override Categorization, Application and URL filtering blade - Advanced Settings, Content Awareness blade - Advanced Settings. 
* `app_control_and_url_filtering_update` - (Optional) Install Application and URL Filtering updates. 
* `install_policy` - (Optional) Install Access Control Policies. 


`endpoint` supports the following:

* `manage_policies_and_software_deployment` - (Optional) The administrator can work with policies, rules and actions. 
* `edit_endpoint_policies` - (Optional) Available only if manage-policies-and-software-deployment is set to true. 
* `policies_installation` - (Optional) The administrator can install policies on endpoint computers. 
* `edit_software_deployment` - (Optional) The administrator can define deployment rules, create packages for export, and configure advanced package settings.<br>Available only if manage-policies-and-software-deployment is set to true. 
* `software_deployment_installation` - (Optional) The administrator can deploy packages and install endpoint clients. 
* `allow_executing_push_operations` - (Optional) The administrator can start operations that the Security Management Server pushes directly to client computers with no policy installation required. 
* `authorize_preboot_users` - (Optional) The administrator can add and remove the users who are permitted to log on to Endpoint Security client computers with Full Disk Encryption. 
* `recovery_media` - (Optional) The administrator can create recovery media on endpoint computers and devices. 
* `remote_help` - (Optional) The administrator can use the Remote Help feature to reset user passwords and give access to locked out users. 
* `reset_computer_data` - (Optional) The administrator can reset a computer, which deletes all information about the computer from the Security Management Server. 


`events_and_reports` supports the following:

* `smart_event` - (Optional) 'Custom' - Configure SmartEvent permissions. 
* `events` - (Optional) Work with event queries on the Events tab. Create custom event queries.<br>Available only if smart-event is set to 'Custom'. 
* `policy` - (Optional) Configure SmartEvent Policy rules and install SmartEvent Policies.<br>Available only if smart-event is set to 'Custom'. 
* `reports` - (Optional) Create and run SmartEvent reports.<br>Available only if smart-event is set to 'Custom'. 


`gateways` supports the following:

* `smart_update` - (Optional) Install, update and delete Check Point licenses. This includes permissions to use SmartUpdate to manage licenses. 
* `lsm_gw_db` - (Optional) Access to objects defined in LSM gateway tables. These objects are managed in the SmartProvisioning GUI or LSMcli command-line.<br>Note: 'Write' permission on lsm-gw-db allows administrator to run a script on SmartLSM gateway in Expert mode. 
* `manage_provisioning_profiles` - (Optional) Administrator can add, edit, delete, and assign provisioning profiles to gateways (both LSM and non-LSM).<br>Available for edit only if lsm-gw-db is set with 'Write' permission.<br>Note: 'Read' permission on lsm-gw-db enables 'Read' permission for manage-provisioning-profiles. 
* `vsx_provisioning` - (Optional) Create and configure Virtual Systems and other VSX virtual objects. 
* `system_backup` - (Optional) Backup Security Gateways. 
* `system_restore` - (Optional) Restore Security Gateways from saved backups. 
* `open_shell` - (Optional) Use the SmartConsole CLI to run commands. 
* `run_one_time_script` - (Optional) Run user scripts from the command line. 
* `run_repository_script` - (Optional) Run scripts from the repository. 
* `manage_repository_scripts` - (Optional) Add, change and remove scripts in the repository. 


`management` supports the following:

* `cme_operations` - (Optional) Permission to read / edit the Cloud Management Extension (CME) configuration.<br>Not supported for Multi-Domain Servers. 
* `manage_admins` - (Optional) Controls the ability to manage Administrators, Permission Profiles, Trusted clients,API settings and Policy settings.<br>Only a "Read Write All" permission-type profile can edit this permission.<br>Not supported for Multi-Domain Servers. 
* `management_api_login` - (Optional) Permission to log in to the Security Management Server and run API commands using thesetools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.<br>Not supported for Multi-Domain Servers. 
* `manage_sessions` - (Optional) Lets you disconnect, discard, publish, or take over other administrator sessions.<br>Only a "Read Write All" permission-type profile can edit this permission. 
* `high_availability_operations` - (Optional) Configure and work with Domain High Availability.<br>Only a 'Customized' permission-type profile can edit this permission. 
* `approve_or_reject_sessions` - (Optional) Approve / reject other sessions. 
* `publish_sessions` - (Optional) Allow session publishing without an approval. 
* `manage_integration_with_cloud_services` - (Optional) Manage integration with Cloud Services. 


`monitoring_and_logging` supports the following:

* `monitoring` - (Optional) See monitoring views and reports. 
* `management_logs` - (Optional) See Multi-Domain Server audit logs. 
* `track_logs` - (Optional) Use the log tracking features in SmartConsole. 
* `app_and_url_filtering_logs` - (Optional) Work with Application and URL Filtering logs. 
* `https_inspection_logs` - (Optional) See logs generated by HTTPS Inspection. 
* `packet_capture_and_forensics` - (Optional) See logs generated by the IPS and Forensics features. 
* `show_packet_capture_by_default` - (Optional) Enable packet capture by default. 
* `identities` - (Optional) Show user and computer identity information in logs. 
* `show_identities_by_default` - (Optional) Show user and computer identity information in logs by default. 
* `dlp_logs_including_confidential_fields` - (Optional) Show DLP logs including confidential fields. 
* `manage_dlp_messages` - (Optional) View/Release/Discard DLP messages.<br>Available only if dlp-logs-including-confidential-fields is set to true. 


`threat_prevention` supports the following:

* `policy_layers` - (Optional) Configure Threat Prevention Policy rules.<br>Note: To have policy-layers permissions you must set policy-exceptionsand profiles permissions. To have 'Write' permissions for policy-layers, policy-exceptions must be set with 'Write' permission as well. 
* `edit_layers` - (Optional) 'ALL' -  Gives permission to edit all layers.<br>"By Selected Profile In A Layer Editor" -  Administrators can only edit the layer if the Threat Prevention layer editor gives editing permission to their profiles.<br>Available only if policy-layers is set to 'Write'. 
* `edit_settings` - (Optional) Work with general Threat Prevention settings. 
* `policy_exceptions` - (Optional) Configure exceptions to Threat Prevention rules.<br>Note: To have policy-exceptions you must set the protections permission. 
* `profiles` - (Optional) Configure Threat Prevention profiles. 
* `protections` - (Optional) Work with malware protections. 
* `install_policy` - (Optional) Install Policies. 
* `ips_update` - (Optional) Update IPS protections.<br>Note: You do not have to log into the User Center to receive IPS updates. 


`others` supports the following:

* `client_certificates` - (Optional) Create and manage client certificates for Mobile Access. 
* `edit_cp_users_db` - (Optional) Work with user accounts and groups. 
* `https_inspection` - (Optional) Enable and configure HTTPS Inspection rules. 
* `ldap_users_db` - (Optional) Work with the LDAP database and user accounts, groups and OUs. 
* `user_authority_access` - (Optional) Work with Check Point User Authority authentication. 
* `user_device_mgmt_conf` - (Optional) Gives access to the UDM (User & Device Management) web-based application that handles security challenges in a "bring your own device" (BYOD) workspace. 


`policy_layers` supports the following:

* `edit_layers` - (Optional) "By Software Blades" - Edit Access Control layers that contain the blades enabled in the Permissions Profile.<br>"By Selected Profile In A Layer Editor" - Administrators can only edit the layer if the Access Control layer editor gives editing permission to their profiles. 
* `app_control_and_url_filtering` - (Optional) Use Application and URL Filtering in Access Control rules.<br>Available only if edit-layers is set to "By Software Blades". 
* `content_awareness` - (Optional) Use specified data types in Access Control rules.<br>Available only if edit-layers is set to "By Software Blades". 
* `firewall` - (Optional) Work with Access Control and other Software Blades that do not have their own Policies.<br>Available only if edit-layers is set to "By Software Blades". 
* `mobile_access` - (Optional) Work with Mobile Access rules.<br>Available only if edit-layers is set to "By Software Blades". 
