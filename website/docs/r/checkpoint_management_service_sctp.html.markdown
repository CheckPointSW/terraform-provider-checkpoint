---
layout: "checkpoint"
page_title: "checkpoint_management_service_sctp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-sctp"
description: |-
This resource allows you to execute Check Point Service Sctp.
---

# checkpoint_management_service_sctp

This resource allows you to execute Check Point Service Sctp.

## Example Usage


```hcl
resource "checkpoint_management_service_sctp" "example" {
  name = "New_SCTP_Service_1"
  port = "5669"
  keep_connections_open_after_policy_installation = false
  session_timeout = 100
  match_for_any = true
  sync_connections_on_cluster = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `aggressive_aging` - (Optional) Sets short (aggressive) timeouts for idle connections.aggressive_aging blocks are documented below.
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `match_for_any` - (Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol. 
* `port` - (Optional) Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. 
* `session_timeout` - (Optional) Time (in seconds) before the session times out. 
* `source_port` - (Optional) Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. 
* `sync_connections_on_cluster` - (Optional) Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `use_default_session_timeout` - (Optional) Use default virtual session timeout. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`aggressive_aging` supports the following:

* `default_timeout` - (Optional) Default aggressive aging timeout in seconds. 
* `enable` - (Optional) N/A 
* `timeout` - (Optional) Aggressive aging timeout in seconds. 
* `use_default_timeout` - (Optional) N/A 
