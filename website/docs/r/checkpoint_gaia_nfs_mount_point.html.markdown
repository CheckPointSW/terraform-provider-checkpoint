---
layout: "checkpoint"
page_title: "checkpoint_gaia_nfs_mount_point"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-nfs-mount-point"
description: |-
This resource allows you to execute Check Point Nfs Mount Point.
---

# checkpoint_gaia_nfs_mount_point

This resource allows you to execute Check Point Nfs Mount Point.

## Example Usage


```hcl
resource "checkpoint_gaia_nfs_mount_point" "example" {
  device_path = "192.168.100.1:/nfs/export"
  mount_point = "/mnt/nfstest"
}
```

## Argument Reference

The following arguments are supported:

* `mount_point` - (Required) The directory on your root file system from which it will be possible to access the content of the device. Mount points should not have spaces in the names. 
* `device_path` - (Required) The device that contains a file system. 
* `options` - (Optional) Mount options of access to the device. For the list of the supported options, see the Gaia Administration Guide, or the built-in help in the corresponding Gaia Clish command. For explanations about these options, see the Linux man pages 'mount(8)' and 'nfs(5)'. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
