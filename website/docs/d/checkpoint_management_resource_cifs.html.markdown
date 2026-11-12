---
layout: "checkpoint"
page_title: "checkpoint_management_resource_cifs"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-cifs"
description: |-
Use this data source to get information on an existing Check Point Resource Cifs.
---

# Data Source: checkpoint_management_resource_cifs

Use this data source to get information on an existing Check Point Resource Cifs.

## Example Usage


```hcl
resource "checkpoint_management_resource_cifs" "cifs" {
  name = "newCifsResource"
  allowed_disk_and_print_shares {
    server_name = "server1"
    share_name = "share1"
  }
  allowed_disk_and_print_shares {
    server_name = "server2"
    share_name = "share2"
  }
  log_mapped_shares = true
  log_access_violation = true
  block_remote_registry_access = false
}

data "checkpoint_management_resource_cifs" "data" {
  uid = "${checkpoint_management_resource_cifs.test.id}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.
* `allowed_disk_and_print_shares` -  The list of Allowed Disk and Print Shares. Must be added in pairs.allowed_disk_and_print_shares blocks are documented below.
* `log_mapped_shares` -  Logs each share map attempt. 
* `log_access_violation` - Logs any attempt to violate the restrictions imposed by the Resource. 
* `block_remote_registry_access` -  Blocks the ability to remotely manipulate a the window's registry. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 


`allowed_disk_and_print_shares` supports the following:

* `server_name` -  Blocks the ability to remotely manipulate a the window's registry. 
* `share_name` -  Disk shares. 
