---
layout: "checkpoint"
page_title: "checkpoint_management_data_multicast_address_range"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-multicast-address-range"
description: |-
  Use this data source to get information on an existing Check Point Multicast Address Range.
---

# Data Source: checkpoint_management_data_multicast_address_range

Use this data source to get information on an existing Check Point Multicast Address Range.

## Example Usage


```hcl
resource "checkpoint_management_multicast_address_range" "multicast_address_range" {
    name = "multicast address range"
    ipv4_address_first = "224.0.0.1"
    ipv4_address_last = "224.0.0.4"
}

data "checkpoint_management_data_multicast_address_range" "data_multicast_address_range" {
    name = "${checkpoint_management_multicast_address_range.multicast_address_range.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 
* `ipv4_address` - IPv4 address. 
* `ipv6_address` - IPv6 address. 
* `ipv4_address_first` - First IPv4 address in the range. 
* `ipv6_address_first` - First IPv6 address in the range. 
* `ipv4_address_last` - Last IPv4 address in the range. 
* `ipv6_address_last` - Last IPv6 address in the range. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers. 