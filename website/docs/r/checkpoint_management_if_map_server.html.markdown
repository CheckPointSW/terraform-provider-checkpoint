---
layout: "checkpoint"
page_title: "checkpoint_management_if_map_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-if-map-server"
description: |-
This resource allows you to execute Check Point If Map Server.
---

# checkpoint_management_if_map_server

This resource allows you to execute Check Point If Map Server.

## Example Usage


```hcl
resource "checkpoint_management_if_map_server" "example" {
  name = "TestIfMapServer"
  host = "TestHost"
  monitored_ips {
    first_ip = "0.0.0.0"
    last_ip = "0.0.0.0"
  }
  version = "2"
  port = 1
  path = "path"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `host` - (Required) Host that is IF-MAP server. Identified by name or UID.
* `monitored_ips` - (Required) IP ranges to be monitored by the IF-MAP client. monitored_ips blocks are documented below.
* `port` - (Optional) IF-MAP server port number. 
* `version` - (Optional) IF-MAP version.
* `path` - (Optional) N/A 
* `query_whole_ranges` - (Optional) Indicate whether to query whole ranges instead of single IP. 
* `authentication` - (Optional) Authentication configuration for the IF-MAP server. authentication blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`monitored_ips` supports the following:

* `first_ip` - (Optional) First IPv4 address in the range to be monitored. 
* `last_ip` - (Optional) Last IPv4 address in the range to be monitored. 


`authentication` supports the following:

* `authentication_method` - (Optional) Authentication method for the IF-MAP server. 
* `username` - (Optional) Username for the IF-MAP server authentication. <font color="red">Required only when</font> 'authentication-method' is set to 'basic'. 
* `password` - (Optional) Username for the IF-MAP server authentication. <font color="red">Required only when</font> 'authentication-method' is set to 'basic'. 
