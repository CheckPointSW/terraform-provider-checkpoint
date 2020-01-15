---
layout: "checkpoint"
page_title: "checkpoint_management_address_range"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-address-range"
description: |-
  This resource allows you to add/update/delete Check Point Address Range.
---

# checkpoint_management_address_range

This resource allows you to add/update/delete Check Point Address Range.

## Example Usage


```hcl
resource "checkpoint_management_address_range" "example" {
  name = "New Address Range 1"
  ipv4_address_first = "192.0.2.1"
  ipv4_address_last = "192.0.2.10"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `ipv4_address_first` - (Optional) First IPv4 address in the range.
* `ipv6_address_first` - (Optional) First IPv6 address in the range.
* `ipv4_address_last` - (Optional) Last IPv4 address in the range.
* `ipv6_address_last` - (Optional) Last IPv6 address in the range.
* `nat_settings` - (Optional) NAT settings. NAT settings blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `groups` - (Optional) Collection of group identifiers.
* `tags` - (Optional) Collection of tag identifiers.

`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `hide_behind` - (Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - (Optional) Which gateway should apply the NAT translation.
* `method` - (Optional) NAT translation method.
