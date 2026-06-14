---
layout: "checkpoint"
page_title: "checkpoint_gaia_nfs_mount_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-nfs-mount-settings"
description: |-
This resource allows you to execute Check Point Nfs Mount Settings.
---

# checkpoint_gaia_nfs_mount_settings

This resource allows you to execute Check Point Nfs Mount Settings.

## Example Usage


```hcl
resource "checkpoint_gaia_nfs_mount_settings" "example" {
  timeout = 5
}
```

## Argument Reference

The following arguments are supported:

* `timeout` - (Optional) Nfs timeout in seconds. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
