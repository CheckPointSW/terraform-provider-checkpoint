---
layout: "checkpoint"
page_title: "checkpoint_gaia_scheduled_snapshot"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-scheduled-snapshot"
description: |-
This resource allows you to execute Check Point Scheduled Snapshot.
---

# checkpoint_gaia_scheduled_snapshot

This resource allows you to execute Check Point Scheduled Snapshot.

## Example Usage


```hcl
resource "checkpoint_gaia_scheduled_snapshot" "example" {
  enabled     = true
  name_prefix = "weeklySnap"
  description = "weekly"

  host {
    target = "lvm"
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) State of the snapshot scheduler 
* `name_prefix` - (Optional) Prefix for the snapshots name created by the scheduler 
* `description` - (Optional) Description of the scheduled snapshot 
* `host` - (Optional) Target host for the snapshots creation host blocks are documented below.
* `recurrence` - (Optional) Recurrence of the scheduled snapshot recurrence blocks are documented below.
* `retention_policy` - (Optional) Retention-policy for the snapshot scheduler retention_policy blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`host` supports the following:

* `target` - (Optional) Host target type 
* `ip_address` - (Optional) IP-Address of the target 
* `upload_path` - (Optional) Upload path for scp/ftp targets 
* `username` - (Optional) Username for scp/ftp targets 
* `password` - (Optional) Password for scp/ftp targets 


`recurrence` supports the following:

* `pattern` - (Optional) Recurrence pattern 
* `days` - (Optional) Recurrence days days blocks are documented below.
* `months` - (Optional) Recurrence months months blocks are documented below.
* `weekdays` - (Optional) Recurrence weekdays weekdays blocks are documented below.
* `time` - (Optional) Recurrence time time blocks are documented below.


`retention_policy` supports the following:

* `keep_disk_space_above_in_gb` - (Optional) Minimum diskspace to keep on the local machine (GB) 
* `min_snapshots_to_keep` - (Optional) Minimum snapshots to keep 
* `max_snapshots_to_keep` - (Optional) Maximum snapshots to keep 


`time` supports the following:

* `hour` - (Optional) Time hour 
* `minute` - (Optional) Time minute 
