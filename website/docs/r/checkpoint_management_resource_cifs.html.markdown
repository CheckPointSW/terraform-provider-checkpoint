---
layout: "checkpoint"
page_title: "checkpoint_management_resource_cifs"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-cifs"
description: |-
This resource allows you to execute Check Point Resource Cifs.
---

# checkpoint_management_resource_cifs

This resource allows you to execute Check Point Resource Cifs.

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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `allowed_disk_and_print_shares` - (Required) The list of Allowed Disk and Print Shares. Must be added in pairs.allowed_disk_and_print_shares blocks are documented below.
* `log_mapped_shares` - (Optional) Logs each share map attempt. 
* `log_access_violation` - (Optional) Logs any attempt to violate the restrictions imposed by the Resource. 
* `block_remote_registry_access` - (Optional) Blocks the ability to remotely manipulate a the window's registry. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`allowed_disk_and_print_shares` supports the following:

* `server_name` - (Required) Blocks the ability to remotely manipulate a the window's registry. 
* `share_name` - (Required) Disk shares. 
