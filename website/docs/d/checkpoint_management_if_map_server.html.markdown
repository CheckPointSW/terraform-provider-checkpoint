---
layout: "checkpoint"
page_title: "checkpoint_management_if_map_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-if-map-server"
description: |- Use this data source to get information on an existing IF-MAP Server.
---


# checkpoint_management_if_map_server

Use this data source to get information on an existing IF-MAP Server.

## Example Usage


```hcl
data "checkpoint_management_if_map_server" "data_if_map_server" {
name = "example_if_map_server"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `host` - Host that is IF-MAP server.
  Identified by name or UID.
* `monitored_ips` - IP ranges to be monitored by the IF-MAP client. monitored_ips blocks are documented below.
* `port` - IF-MAP server port number.
* `version` - IF-MAP version.
* `path` - N/A
* `query_whole_ranges` - Indicate whether to query whole ranges instead of single IP.
* `authentication` - Authentication configuration for the IF-MAP server. authentication blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.


`monitored_ips` supports the following:

* `first_ip` - First IPv4 address in the range to be monitored.
* `last_ip` - Last IPv4 address in the range to be monitored.


`authentication` supports the following:

* `authentication_method` - Authentication method for the IF-MAP server.
* `username` - Username for the IF-MAP server authentication.
* `password` - Username for the IF-MAP server authentication.
