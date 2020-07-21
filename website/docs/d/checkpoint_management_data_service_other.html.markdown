---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_other"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-other"
description: |-
  Use this data source to get information on an existing Check Point Service Other.
---

# checkpoint_management_data_service_other

Use this data source to get information on an existing Check Point Service Other.

## Example Usage


```hcl
resource "checkpoint_management_service_other" "service_other" {
    name = "service other"
    keep_connections_open_after_policy_installation = false
	session_timeout = 100
	match_for_any = true
	sync_connections_on_cluster = true
	ip_protocol = 51
	aggressive_aging = {
		use_default_timeout = true
		enable = true
		default_timeout = 600
		timeout = 600
	}
}

data "checkpoint_management_data_service_other" "data_service_other" {
    name = "${checkpoint_management_service_other.service_other.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `accept_replies` - Specifies whether Other Service replies are to be accepted. 
* `action` - Contains an INSPECT expression that defines the action to take if a rule containing this service is matched.
Example: set r_mhandler &open_ssl_handler sets a handler on the connection. 
* `aggressive_aging` - Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.
* `ip_protocol` - IP protocol number. 
* `keep_connections_open_after_policy_installation` - Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. 
* `match` - Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet.
Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections. 
* `match_for_any` - Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol. 
* `override_default_settings` - Indicates whether this service is a Data Domain service which has been overridden. 
* `session_timeout` - Time (in seconds) before the session times out. 
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