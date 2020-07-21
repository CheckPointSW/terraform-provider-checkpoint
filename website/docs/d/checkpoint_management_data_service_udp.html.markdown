---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_udp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-udp"
description: |-
  Use this data source to get information on an existing Check Point Service Udp.
---

# checkpoint_management_data_service_udp

Use this data source to get information on an existing Check Point Service Udp.

## Example Usage


```hcl
resource "checkpoint_management_service_udp" "service_udp" {
    name = "service udp"
	port = "1123"
}

data "checkpoint_management_data_service_udp" "data_service_udp" {
    name = "${checkpoint_management_service_udp.service_udp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `accept_replies`
* `aggressive_aging` - Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
* `match_by_protocol_signature` - A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.
* `match_for_any` - Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.
* `override_default_settings` - Indicates whether this service is a Data Domain service which has been overridden.
* `port` - The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.
* `protocol` - Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).
* `session_timeout` - Time (in seconds) before the session times out.
* `source_port` - Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.
* `sync_connections_on_cluster` - Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.
* `use_default_session_timeout` - Use default virtual session timeout.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.

`aggressive_aging` supports the following:

* `default_timeout` - (Optional) Default aggressive aging timeout in seconds.
* `enable`
* `timeout` - (Optional) Aggressive aging timeout in seconds.
* `use_default_timeout`
