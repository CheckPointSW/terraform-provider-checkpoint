---
layout: "checkpoint"
page_title: "checkpoint_management_hosts"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-hosts"
description: |- Use this data source to get information on hosts.
---


# checkpoint_management_hosts

Use this data source to get information on hosts.

## Example Usage


```hcl
data "checkpoint_management_hosts" "my_query" {
  limit = 15
}

# Fetch all results
data "checkpoint_management_hosts" "my_query_fetch_all" {
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
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `interfaces` - Host interfaces. Host interfaces blocks are documented below.
* `nat_settings` - NAT settings. NAT settings blocks are documented below.
* `host_servers` - Servers Configuration. Servers Configuration blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.
* `domain` - Information about the domain that holds the Object. Domain information blocks are documented below.
* `icon` - Object icon.


`interfaces` supports the following:

* `name` - Object name. Should be unique in the domain.
* `subnet4` - IPv4 network address.
* `subnet6` - IPv6 network address.
* `mask_length4` - IPv4 network mask length.
* `mask_length6` - IPv6 network mask length.
* `ignore_warnings` - Apply changes ignoring warnings.
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
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

`domain` supports the following:
* `name` - Object name.
* `uid` - Object unique identifier.
* `domain_type` - Domain type.
