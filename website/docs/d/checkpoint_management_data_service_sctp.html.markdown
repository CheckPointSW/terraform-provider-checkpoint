---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_sctp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-sctp"
description: |-
  Use this data source to get information on an existing Check Point Service Sctp.
---

# checkpoint_management_data_service_sctp

Use this data source to get information on an existing Check Point Service Sctp.

## Example Usage


```hcl
resource "checkpoint_management_service_sctp" "service_sctp" {
    name = "service sctp"
    port = "1234"
    session_timeout = "3600"
    sync_connections_on_cluster = true
}

data "checkpoint_management_data_service_sctp" "data_service_sctp" {
    name = "${checkpoint_management_service_sctp.service_sctp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `aggressive_aging` - Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `match_for_any` - Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol. 
* `port` - Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. 
* `session_timeout` - Time (in seconds) before the session times out. 
* `source_port` - Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. 
* `sync_connections_on_cluster` - Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster. 
* `tags` - Collection of tag identifiers.
* `use_default_session_timeout` - Use default virtual session timeout. 
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.


`aggressive_aging` supports the following:

* `default_timeout` - Default aggressive aging timeout in seconds. 
* `enable`
* `timeout` - Aggressive aging timeout in seconds. 
* `use_default_timeout`