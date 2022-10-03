---
layout: "checkpoint"
page_title: "checkpoint_management_tacacs_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-tacacs-server"
description: |-
Use this data source to get information on an existing Check Point Tacacs Server.
---

# Data Source: checkpoint_management_tacacs_server

Use this data source to get information on an existing Check Point Tacacs Server.

## Example Usage


```hcl
resource "checkpoint_management_host" "host" {
  name = "My Host"
  ipv4_address = "1.2.3.4"
}

resource "checkpoint_management_tacacs_server" "tacacs_server" {
    name = "My Tacacs Server"
    server = "1.2.3.4"
}

data "checkpoint_management_tacacs_server" "data_tacacs_server" {
    name = "${checkpoint_management_tacacs_server.tacacs_server.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `encryption` - Is there a secret key defined on the server. Must be set true when "server-type" was selected to be "TACACS+".
* `groups` - Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `priority` - The priority of the TACACS Server in case it is a member of a TACACS Group.
* `server` - The UID or Name of the host that is the TACACS Server. server blocks are documented below.
* `server_type` - Server type, TACACS or TACACS+.
* `service` - Server service, only relevant when "server-type" is TACACS. service blocks are documented below.

`server` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.


`service` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.
* `aggressive_aging` - Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.
* `groups` - Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
* `match_for_any` - Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.
* `port` - The number of the port used to provide this service.
* `session_timeout` - Time (in seconds) before the session times out.
* `source_port` - Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.
* `sync_connections_on_cluster` - Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.
* `use_default_session_timeout` - Use default virtual session timeout.


`aggressive_aging` supports the following:

* `default_timeout` - Default aggressive aging timeout in seconds.
* `enabled`
* `timeout` - Aggressive aging timeout in seconds.
* `use_default_timeout`