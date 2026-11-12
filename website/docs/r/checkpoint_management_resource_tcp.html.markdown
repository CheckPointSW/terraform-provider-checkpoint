---
layout: "checkpoint"
page_title: "checkpoint_management_resource_tcp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-tcp"
description: |-
This resource allows you to execute Check Point Resource Tcp.
---

# checkpoint_management_resource_tcp

This resource allows you to execute Check Point Resource Tcp.

## Example Usage


```hcl
resource "checkpoint_management_resource_tcp" "example" {
  name = "newTcpResource"
  ufp-settings {
    server = "ufpServer",
    caching-control = "security_gateway_one_request",
    ignore-ufp-server-after-failure = true,
    number-of-failures-before-ignore = 3
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `resource_type` - (Optional) The type of the TCP resource. 
* `exception_track` - (Optional) Configures how to track connections that match this rule but fail the content security checks. An example of an exception is a connection with an unsupported scheme or method. 
* `ufp_settings` - (Optional) UFP settings. ufp_settings blocks are documented below.
* `cvp_settings` - (Optional) CVP settings. cvp_settings blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`ufp_settings` supports the following:

* `server` - (Optional) UFP server identified by name or UID.
The UFP server must already be defined as an OPSEC Application. 
* `caching_control` - (Optional) Specifies if and how caching is to be enabled. 
* `ignore_ufp_server_after_failure` - (Optional) The UFP server will be ignored after numerous UFP server connections were unsuccessful. 
* `number_of_failures_before_ignore` - (Optional) Signifies at what point the UFP server should be ignored, Applicable only if 'ignore after fail' is enabled. 
* `timeout_before_reconnecting` - (Optional) The amount of time, in seconds, that must pass before a UFP server connection should be attempted, Applicable only if 'ignore after fail' is enabled. 


`cvp_settings` supports the following:

* `server` - (Optional) CVP server identified by name or UID.
The CVP server must already be defined as an OPSEC Application. 
* `allowed_to_modify_content` - (Optional) Configures the CVP server to inspect but not modify content. 
* `reply_order` - (Optional) Designates when the CVP server returns data to the Security Gateway security server. 
