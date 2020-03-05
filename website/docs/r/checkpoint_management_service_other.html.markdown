---
layout: "checkpoint"
page_title: "checkpoint_management_service_other"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-other"
description: |-
This resource allows you to execute Check Point Service Other.
---

# checkpoint_management_service_other

This resource allows you to execute Check Point Service Other.

## Example Usage


```hcl
resource "checkpoint_management_service_other" "example" {
  name = "New_Service_1"
  keep_connections_open_after_policy_installation = false
  session_timeout = 100
  match_for_any = true
  sync_connections_on_cluster = true
  ip_protocol = 51
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `accept_replies` - (Optional) Specifies whether Other Service replies are to be accepted. 
* `action` - (Optional) Contains an INSPECT expression that defines the action to take if a rule containing this service is matched.
Example: set r_mhandler &open_ssl_handler sets a handler on the connection. 
* `aggressive_aging` - (Optional) Sets short (aggressive) timeouts for idle connections.aggressive_aging blocks are documented below.
* `ip_protocol` - (Optional) IP protocol number. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `match` - (Optional) Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet.
Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections. 
* `match_for_any` - (Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol. 
* `override_default_settings` - (Optional) Indicates whether this service is a Data Domain service which has been overridden. 
* `session_timeout` - (Optional) Time (in seconds) before the session times out. 
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
