---
layout: "checkpoint"
page_title: "checkpoint_management_host"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-host"
description: |-
  This resource allows you to add/update/delete Check Point Host.
---

# checkpoint_management_host

This resource allows you to add/update/delete Check Point Host.

## Example Usage


```hcl
resource "checkpoint_management_host" "example" {
  name = "New Host 1"
  ipv4_address = "192.0.2.1"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `interfaces` - (Optional) Host interfaces. Host interfaces blocks are documented below.
* `nat_settings` - (Optional) NAT settings. NAT settings blocks are documented below.
* `host_servers` - (Optional) Servers Configuration. Servers Configuration blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `groups` - (Optional) Collection of group identifiers.
* `tags` - (Optional) Collection of tag identifiers.


`interfaces` supports the following:

* `name` - (Required) Object name. Should be unique in the domain.
* `subnet4` - (Optional) IPv4 network address.
* `subnet6` - (Optional) IPv6 network address.
* `mask_length4` - (Optional) IPv4 network mask length.
* `mask_length6` - (Optional) IPv6 network mask length.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.

`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `hide_behind` - (Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - (Optional) Which gateway should apply the NAT translation.
* `method` - (Optional) NAT translation method.

`host_servers` supports the following:

* `dns_server` - (Optional) Gets True if this server is a DNS Server.
* `mail_server` - (Optional) Gets True if this server is a Mail Server.
* `web_server` - (Optional) Gets True if this server is a Web Server.
* `web_server_config` - (Optional) Web Server configuration. Web Server configuration blocks are documented below.

`web_server_config` supports the following:

* `additional_ports` - (Optional) Server additional ports.
* `application_engines` - (Optional) Application engines of this web server.
* `listen_standard_port` - (Optional) "Whether server listens to standard port.
* `operating_system` - (Optional) Operating System.
* `protected_by` - (Optional) Network object which protects this server identified by the name or UID.
