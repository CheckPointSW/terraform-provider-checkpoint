---
layout: "checkpoint"
page_title: "checkpoint_management_services-tcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-services-tcp"
description: |- Use this data source to get information on all Services TCP.
---


# checkpoint_management_identity_provider

Use this data source to get information on all Services TCP.

## Example Usage


```hcl
data "checkpoint_management_services_tcp" "my_query" {
  limit = 15
}

# Fetch all results
data "checkpoint_management_services_tcp" "my_query_fetch_all" {
  fetch_all = true
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Search expression to filter objects by.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional) Number of the results to initially skip.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.
* `fetch_all` - (Optional) If true, fetches all results.
* `from` - From which element number the query was done.
* `to` - To which element number the query was done.
* `total` - Total number of elements returned by the query.
* `objects` - Objects list. objects blocks are documented below.

`order` supports the following:
* `asc` - Sorts results by the given field in ascending order.
* `desc` - Sorts results by the given field in descending order.

`objects` supports the following:
* `name` - Object name. Should be unique in the domain.
* `uid` - Object unique identifier.
* `type` - Object type.
* `port` - The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.
* `aggressive_aging` - Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
* `match_by_protocol_signature` - A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.
* `match_for_any` - Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.
* `override_default_settings` - Indicates whether this service is a Data Domain service which has been overridden.
* `protocol` - Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).
* `session_timeout` - Time (in seconds) before the session times out.
* `source_port` - Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.
* `sync_connections_on_cluster` - Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.
* `use_default_session_timeout` - Use default virtual session timeout.
* `use_delayed_sync` - Enable this option to delay notifying the Security Gateway about a connection, so that the connection will only be synchronized if it still exists x seconds after the connection is initiated.
* `delayed_sync_value` - Specify the delay (in seconds) in which a synchronization will start after connection initiation.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.
* `domain` - Information about the domain that holds the Object. domain blocks are documented below.
* `icon` - Object icon.

`aggressive_aging` supports the following:

* `default_timeout` - Default aggressive aging timeout in seconds.
* `enable`
* `timeout` - Aggressive aging timeout in seconds.
* `use_default_timeout`

`domain` supports the following:

* `name` - Object name.
* `uid` - Object unique identifier.
* `domain_type` - Domain type.