---
layout: "checkpoint"
page_title: "checkpoint_management_checkpoint_host"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-checkpoint-host"
description: |-
This resource allows you to execute Check Point Checkpoint Host.
---

# Resource: checkpoint_management_checkpoint_host

This resource allows you to execute Check Point Checkpoint Host.

## Example Usage


```hcl
resource "checkpoint_management_checkpoint_host" "example" {
  name = "checkpoint host"
  ipv4_address = "5.5.5.5"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `interfaces` - (Optional) Checkpoint host interfaces. interfaces blocks are documented below.
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `nat_settings` - (Optional) NAT settings. nat_settings blocks are documented below.
* `one_time_password` - (Optional) Secure internal connection one time password. 
* `hardware` - (Optional) Hardware name. 
* `os` - (Optional) Operating system name. 
* `version` - (Optional) Checkpoint host platform version. 
* `management_blades` - (Optional) Management blades. management_blades blocks are documented below.
* `logs_settings` - (Optional) Logs settings. logs_settings blocks are documented below.
* `save_logs_locally` - (Optional) Enable save logs locally. 
* `send_alerts_to_server` - (Optional) Collection of Server(s) to send alerts to identified by the name or UID.
* `send_logs_to_backup_server` - (Optional) Collection of Backup server(s) to send logs to identified by the name or UID.
* `send_logs_to_server` - (Optional) Collection of Server(s) to send logs to identified by the name or UID.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `sic_name` - (Computed) Name of the Secure Internal Connection Trust.
* `sic_state` - (Computed) State the Secure Internal Connection Trust.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`interfaces` supports the following:

* `name` - (Optional) Interface name. 
* `subnet4` - (Optional) IPv4 network address. 
* `subnet6` - (Optional) IPv6 network address. 
* `mask_length4` - (Optional) IPv4 network mask length. 
* `mask_length6` - (Optional) IPv6 network mask length. 
* `subnet_mask` - (Optional) IPv4 network mask. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `hide_behind` - (Optional) Hide behind method. This parameter is not required in case "method" parameter is "static". 
* `install_on` - (Optional) Which gateway should apply the NAT translation. 
* `method` - (Optional) NAT translation method. 


`management_blades` supports the following:

* `network_policy_management` - (Optional) Enable Network Policy Management. 
* `logging_and_status` - (Optional) Enable Logging & Status. 
* `smart_event_server` - (Optional) Enable SmartEvent server. When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164. 
* `smart_event_correlation` - (Optional) Enable SmartEvent Correlation Unit. 
* `endpoint_policy` - (Optional) Enable Endpoint Policy. To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. Field is not supported on Multi Domain Server environment. 
* `compliance` - (Optional) Compliance blade. Can be set when 'network-policy-management' was selected to be True. 
* `user_directory` - (Optional) Enable User Directory. Can be set when 'network-policy-management' was selected to be True. 
* `secondary` - (Computed) Secondary Management enabled.
* `identity_logging` - (Computed) Identity Logging enabled.

`logs_settings` supports the following:

* `free_disk_space_metrics` - (Optional) Free disk space metrics. 
* `accept_syslog_messages` - (Optional) Enable accept syslog messages. 
* `alert_when_free_disk_space_below` - (Optional) Enable alert when free disk space is below threshold. 
* `alert_when_free_disk_space_below_threshold` - (Optional) Alert when free disk space below threshold. 
* `alert_when_free_disk_space_below_type` - (Optional) Alert when free disk space below type. 
* `before_delete_keep_logs_from_the_last_days` - (Optional) Enable before delete keep logs from the last days. 
* `before_delete_keep_logs_from_the_last_days_threshold` - (Optional) Before delete keep logs from the last days threshold. 
* `before_delete_run_script` - (Optional) Enable Before delete run script. 
* `before_delete_run_script_command` - (Optional) Before delete run script command. 
* `delete_index_files_older_than_days` - (Optional) Enable delete index files older than days. 
* `delete_index_files_older_than_days_threshold` - (Optional) Delete index files older than days threshold. 
* `delete_when_free_disk_space_below` - (Optional) Enable delete when free disk space below. 
* `delete_when_free_disk_space_below_threshold` - (Optional) Delete when free disk space below threshold. 
* `detect_new_citrix_ica_application_names` - (Optional) Enable detect new citrix ica application names. 
* `enable_log_indexing` - (Optional) Enable log indexing. 
* `forward_logs_to_log_server` - (Optional) Enable forward logs to log server. 
* `forward_logs_to_log_server_name` - (Optional) Forward logs to log server name. 
* `forward_logs_to_log_server_schedule_name` - (Optional) Forward logs to log server schedule name. 
* `rotate_log_by_file_size` - (Optional) Enable rotate log by file size. 
* `rotate_log_file_size_threshold` - (Optional) Log file size threshold. 
* `rotate_log_on_schedule` - (Optional) Enable rotate log on schedule. 
* `rotate_log_schedule_name` - (Optional) Rotate log schedule name. 
* `smart_event_intro_correletion_unit` - (Optional) Enable SmartEvent intro correletion unit. 
* `stop_logging_when_free_disk_space_below` - (Optional) Enable stop logging when free disk space below. 
* `stop_logging_when_free_disk_space_below_threshold` - (Optional) Stop logging when free disk space below threshold. 
* `turn_on_qos_logging` - (Optional) Enable turn on qos logging. 
* `update_account_log_every` - (Optional) Update account log in every amount of seconds.