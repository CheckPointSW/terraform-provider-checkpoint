---
layout: "checkpoint"
page_title: "checkpoint_management_logical_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-logical-server"
description: |-
Use this data source to get information on an existing Check Point Logical Server.
---

# Data Source: checkpoint_management_logical_server

Use this data source to get information on an existing Check Point Logical Server.

## Example Usage
```hcl
data "checkpoint_management_logical_server" "data_test" {
    name = "logical_server1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `server_group` - Server group associated with the logical server.
  Identified by name or UID.
* `server_type` - Type of server for the logical server.
* `persistence_mode` - Indicates if persistence mode is enabled for the logical server.
* `persistency_type` - Persistency type for the logical server.
* `balance_method` - Load balancing method for the logical server.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
