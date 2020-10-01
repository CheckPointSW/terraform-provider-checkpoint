---
layout: "checkpoint"
page_title: "checkpoint_management_checkpoint_host"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-checkpoint-host"
description: |-
This resource allows you to execute Check Point Checkpoint Host.
---

# Data Source: checkpoint_management_checkpoint_host

This resource allows you to execute Check Point Checkpoint Host.

## Example Usage


```hcl
resource "checkpoint_management_checkpoint_host" "checkpoint_host" {
	name = "checkpoint host"
	ipv4_address = "1.2.3.4"
}

data "checkpoint_management_checkpoint_host" "data_checkpoint_host" {
	name = "${checkpoint_management_checkpoint_host.checkpoint_host.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `interfaces` - Checkpoint host interfaces. interfaces blocks are documented below.
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address. 
* `nat_settings` - NAT settings. nat_settings blocks are documented below.
* `hardware` - Hardware name. 
* `os` - Operating system name. 
* `version` - Checkpoint host platform version. 
* `management_blades` - Management blades. management_blades blocks are documented below.
* `logs_settings` - Logs settings. logs_settings blocks are documented below.
* `save_logs_locally` - Enable save logs locally. 
* `send_alerts_to_server` - Collection of Server(s) to send alerts to identified by the name or UID.
* `send_logs_to_backup_server` - Collection of Backup server(s) to send logs to identified by the name or UID.
* `send_logs_to_server` - Collection of Server(s) to send logs to identified by the name or UID.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 
* `sic_name` - Name of the Secure Internal Connection Trust.
* `sic_state` - State the Secure Internal Connection Trust.


`interfaces` supports the following:

* `name` - Interface name. 
* `subnet4` - IPv4 network address. 
* `subnet6` - IPv6 network address. 
* `mask_length4` - IPv4 network mask length. 
* `mask_length6` - IPv6 network mask length. 
* `subnet_mask` - IPv4 network mask. 
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 


`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules. 
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address. 
* `hide_behind` - Hide behind method. This parameter is not required in case "method" parameter is "static". 
* `install_on` - Which gateway should apply the NAT translation. 
* `method` - NAT translation method. 


`management_blades` supports the following:

* `network_policy_management` - Enable Network Policy Management. 
* `logging_and_status` - Enable Logging & Status. 
* `smart_event_server` - Enable SmartEvent server. When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164. 
* `smart_event_correlation` - Enable SmartEvent Correlation Unit. 
* `endpoint_policy` - Enable Endpoint Policy. To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. Field is not supported on Multi Domain Server environment. 
* `compliance` - Compliance blade. Can be set when 'network-policy-management' was selected to be True. 
* `user_directory` - Enable User Directory. Can be set when 'network-policy-management' was selected to be True. 
* `secondary` - Secondary Management enabled.
* `identity_logging` - Identity Logging enabled.

`logs_settings` supports the following:

* `free_disk_space_metrics` - Free disk space metrics. 
* `accept_syslog_messages` - Enable accept syslog messages. 
* `alert_when_free_disk_space_below` - Enable alert when free disk space is below threshold. 
* `alert_when_free_disk_space_below_threshold` - Alert when free disk space below threshold. 
* `alert_when_free_disk_space_below_type` - Alert when free disk space below type. 
* `before_delete_keep_logs_from_the_last_days` - Enable before delete keep logs from the last days. 
* `before_delete_keep_logs_from_the_last_days_threshold` - Before delete keep logs from the last days threshold. 
* `before_delete_run_script` - Enable Before delete run script. 
* `before_delete_run_script_command` - Before delete run script command. 
* `delete_index_files_older_than_days` - Enable delete index files older than days. 
* `delete_index_files_older_than_days_threshold` - Delete index files older than days threshold. 
* `delete_when_free_disk_space_below` - Enable delete when free disk space below. 
* `delete_when_free_disk_space_below_threshold` - Delete when free disk space below threshold. 
* `detect_new_citrix_ica_application_names` - Enable detect new citrix ica application names. 
* `enable_log_indexing` - Enable log indexing. 
* `forward_logs_to_log_server` - Enable forward logs to log server. 
* `forward_logs_to_log_server_name` - Forward logs to log server name. 
* `forward_logs_to_log_server_schedule_name` - Forward logs to log server schedule name. 
* `rotate_log_by_file_size` - Enable rotate log by file size. 
* `rotate_log_file_size_threshold` - Log file size threshold. 
* `rotate_log_on_schedule` - Enable rotate log on schedule. 
* `rotate_log_schedule_name` - Rotate log schedule name. 
* `smart_event_intro_correletion_unit` - Enable SmartEvent intro correletion unit. 
* `stop_logging_when_free_disk_space_below` - Enable stop logging when free disk space below. 
* `stop_logging_when_free_disk_space_below_threshold` - Stop logging when free disk space below threshold. 
* `turn_on_qos_logging` - Enable turn on qos logging. 
* `update_account_log_every` - Update account log in every amount of seconds.