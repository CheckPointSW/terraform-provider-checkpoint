---
layout: "checkpoint"
page_title: "checkpoint_management_data_host"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-host"
description: |-
  Use this data source to get information on an existing Check Point Host.
---

# Data Source: checkpoint_management_data_host

Use this data source to get information on an existing Check Point Host.

## Example Usage


```hcl
resource "checkpoint_management_host" "host" {
    name = "My Host"
    ipv4_address = "1.2.3.4"
}

data "checkpoint_management_data_host" "data_host" {
    name = "${checkpoint_management_host.host.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `interfaces` - Host interfaces. interfaces blocks are documented below.
* `nat_settings` - NAT settings. nat_settings blocks are documented below.
* `host_servers` - Servers Configuration. host_servers blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.


`interfaces` supports the following:

* `name` - Object name. Should be unique in the domain.
* `subnet4` - IPv4 network address.
* `subnet6` - IPv6 network address.
* `mask_length4` - IPv4 network mask length.
* `mask_length6` - IPv6 network mask length.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `hide_behind` - Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - Which gateway should apply the NAT translation.
* `method` - NAT translation method.

`host_servers` supports the following:

* `dns_server` - Gets True if this server is a DNS Server.
* `mail_server` - Gets True if this server is a Mail Server.
* `web_server` - Gets True if this server is a Web Server.
* `web_server_config` - Web Server configuration. Web Server configuration blocks are documented below.

`web_server_config` supports the following:

* `additional_ports` - Server additional ports.
* `application_engines` - Application engines of this web server.
* `listen_standard_port` - "Whether server listens to standard port.
* `operating_system` - Operating System.
* `protected_by` - Network object which protects this server identified by the name or UID.
