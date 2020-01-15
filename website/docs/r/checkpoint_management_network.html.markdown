---
layout: "checkpoint"
page_title: "checkpoint_management_network "
sidebar_current: "docs-checkpoint-resource-checkpoint-management-network"
description: |-
  This resource allows you to add/update/delete Check Point Network Object.
---

# checkpoint_management_network

This resource allows you to add/update/delete Check Point Network Object.

## Example Usage


```hcl
resource "checkpoint_management_network" "example" {
  name = "New Network 1"
  subnet4 = "192.0.2.0"
  mask_length4 = 32
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `subnet4` - (Optional) IPv4 network address.
* `subnet6` - (Optional) IPv6 network address..
* `mask_length4` - (Optional) IPv4 network mask length.
* `mask_length6` - (Optional) IPv6 network mask length.
* `nat_settings` - (Optional) NAT settings. NAT settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.
* `groups` - (Optional) Collection of group identifiers.
* `broadcast` - (Optional) "Allow broadcast address inclusion.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.

`nat_settings` supports the following:

* `auto_rule` - (Optional) Whether to add automatic address translation rules.
* `ipv4_address` - (Optional) IPv4 address.
* `ipv6_address` - (Optional) IPv6 address.
* `hide_behind` - (Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - (Optional) Which gateway should apply the NAT translation.
* `method` - (Optional) NAT translation method.
