---
layout: "checkpoint"
page_title: "checkpoint_management_data_address_range"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-address-range"
description: |-
  Use this data source to get information on an existing Check Point Address Range.
---

# checkpoint_management_data_address_range

Use this data source to get information on an existing Check Point Address Range.

## Example Usage


```hcl
resource "checkpoint_management_address_range" "address_range" {
    name = "My Address Range"
    ipv4_address_first = "1.1.1.1"
    ipv4_address_last = "2.2.2.2"
}

data "checkpoint_management_data_address_range" "data_address_range" {
    name = "${checkpoint_management_address_range.address_range.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `ipv4_address_first` - First IPv4 address in the range.
* `ipv6_address_first` - First IPv6 address in the range.
* `ipv4_address_last` - Last IPv4 address in the range.
* `ipv6_address_last` - Last IPv6 address in the range.
* `nat_settings` - NAT settings. NAT settings blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.

`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `hide_behind` - (Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - (Optional) Which gateway should apply the NAT translation.
* `method` - (Optional) NAT translation method.