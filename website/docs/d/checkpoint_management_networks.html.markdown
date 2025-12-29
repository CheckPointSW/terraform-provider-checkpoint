---
layout: "checkpoint"
page_title: "checkpoint_management_networks"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-networks"
description: |- Use this data source to get information on networks.
---


# checkpoint_management_networks

Use this data source to get information on networks.

## Example Usage


```hcl
data "checkpoint_management_networks" "my_query" {
  limit = 15
}

# Fetch all results
data "checkpoint_management_networks" "my_query_fetch_all" {
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
* `uid` - Object unique identifier.
* `subnet4` - IPv4 network address.
* `subnet6` - IPv6 network address.
* `mask_length4` - IPv4 network mask length.
* `mask_length6` - IPv6 network mask length.
* `subnet_mask` - IPv4 network mask.
* `nat_settings` - NAT settings. nat_settings blocks are documented below.
* `tags` - Collection of tag identifiers.
* `groups` - Collection of group identifiers.
* `broadcast` - Allow broadcast address inclusion.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `domain` - Information about the domain that holds the Object. domain blocks are documented below.
* `icon` - Object icon.

`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `hide_behind` - Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - Which gateway should apply the NAT translation.
* `method` - NAT translation method.

`domain` supports the following:
* `name` - Object name.
* `uid` - Object unique identifier.
* `domain_type` - Domain type.