---
layout: "checkpoint"
page_title: "checkpoint_gaia_scheduled_backup"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-scheduled-backup"
description: |-
This resource allows you to execute Check Point Scheduled Backup.
---

# checkpoint_gaia_scheduled_backup

This resource allows you to execute Check Point Scheduled Backup.

## Example Usage


```hcl
resource "checkpoint_gaia_scheduled_backup" "example" {
  name = "my_backup"
  host {
    target      = "local"
  }
  recurrence {
    pattern  = "monthly"
    days     = ["1"]
    months   = ["1"]
  }
  time {
    hour   = 1
    minute = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) backup schedule name 
* `host` - (Required) scheduled backup host host blocks are documented below.
* `recurrence` - (Required) scheduled backup recurrence recurrence blocks are documented below.
* `time` - (Required) scheduled backup time time blocks are documented below.
* `retention_policy` - (Optional) Retention-policy for the backup scheduler, supported from R81.10 and above retention_policy blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`host` supports the following:

* `target` - (Optional) backup host type 
* `ip_address` - (Optional) backup host IPv4 address 
* `upload_path` - (Optional) backup host upload path 
* `username` - (Optional) backup host username 
* `password` - (Optional) backup host password 


`recurrence` supports the following:

* `pattern` - (Optional) backup recurrence pattern 
* `days` - (Optional) backup recurrence days days blocks are documented below.
* `months` - (Optional) backup recurrence months months blocks are documented below.
* `weekdays` - (Optional) backup recurrence weekdays weekdays blocks are documented below.


`time` supports the following:

* `hour` - (Optional) backup time hour 
* `minute` - (Optional) backup time minute 


`retention_policy` supports the following:

* `max_disk_space` - (Optional) Maximum diskspace to keep on the local machine (MB) 
* `min_num_of_backups` - (Optional) Minimum backups to keep 
* `max_num_of_backups` - (Optional) Maximum backups to keep 
