---
layout: "checkpoint"
page_title: "checkpoint_management_resource_tcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-resource-tcp"
description: |- Use this data source to get information on an existing TCP resource.
---


# checkpoint_management_resource_tcp

Use this data source to get information on an existing TCP resource.

## Example Usage


```hcl
data "checkpoint_management_resource_tcp" "data_tcp" {
  name = "tcp_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `resource_type` - The type of the TCP resource.
* `exception_track` - Configures how to track connections that match this rule but fail the content security checks. An example of an exception is a connection with an unsupported scheme or method.
* `ufp_settings` - UFP settings. ufp_settings blocks are documented below.
* `cvp_settings` - CVP settings. cvp_settings blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.


`ufp_settings` supports the following:

* `server` - UFP server identified by name or UID.
  The UFP server must already be defined as an OPSEC Application.
* `caching_control` - Specifies if and how caching is to be enabled.
* `ignore_ufp_server_after_failure` - The UFP server will be ignored after numerous UFP server connections were unsuccessful.
* `number_of_failures_before_ignore` - Signifies at what point the UFP server should be ignored, Applicable only if 'ignore after fail' is enabled.
* `timeout_before_reconnecting` - The amount of time, in seconds, that must pass before a UFP server connection should be attempted, Applicable only if 'ignore after fail' is enabled.


`cvp_settings` supports the following:

* `server` - CVP server identified by name or UID.
  The CVP server must already be defined as an OPSEC Application.
* `allowed_to_modify_content` - Configures the CVP server to inspect but not modify content.
* `reply_order` - Designates when the CVP server returns data to the Security Gateway security server. 
