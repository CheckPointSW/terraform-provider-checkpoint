---
layout: "checkpoint"
page_title: "checkpoint_management_multicast_address_range"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-multicast-address-range"
description: |-
This resource allows you to execute Check Point Multicast Address Range.
---

# checkpoint_management_multicast_address_range

This resource allows you to execute Check Point Multicast Address Range.

## Example Usage


```hcl
resource "checkpoint_management_multicast_address_range" "example" {
  name = "New Multicast Address Range"
  ipv4_address_first = "224.0.0.1"
  ipv4_address_last = "224.0.0.4"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv4_address_first` - (Optional) First IPv4 address in the range. 
* `ipv6_address_first` - (Optional) First IPv6 address in the range. 
* `ipv4_address_last` - (Optional) Last IPv4 address in the range. 
* `ipv6_address_last` - (Optional) Last IPv6 address in the range. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
