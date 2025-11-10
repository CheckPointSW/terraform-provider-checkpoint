---
layout: "checkpoint"
page_title: "checkpoint_management_logical_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-logical-server"
description: |-
 This resource allows you to execute Check Point Logical Server.
---

# checkpoint_management_logical_server

This resource allows you to execute Check Point Logical Server.

## Example Usage


```hcl
resource "checkpoint_management_logical_server" "example" {
  name = "logicalServer1"
  server_group = "testGroup"
  server_type = "other"
  persistence_mode = true
  persistency_type = "by_server"
  balance_method = "domain"
  ipv4_address = "1.1.1.1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `server_group` - (Optional) Server group associated with the logical server. 
Identified by name or UID. 
* `server_type` - (Optional) Type of server for the logical server. 
* `persistence_mode` - (Optional) Indicates if persistence mode is enabled for the logical server. 
* `persistency_type` - (Optional) Persistency type for the logical server. 
* `balance_method` - (Optional) Load balancing method for the logical server. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
