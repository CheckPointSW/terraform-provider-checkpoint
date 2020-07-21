---
layout: "checkpoint"
page_title: "checkpoint_management_data_network "
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-network"
description: |-
  Use this data source to get information on an existing Check Point Network Object.
---

# checkpoint_management_data_network

Use this data source to get information on an existing Check Point Network Object.

## Example Usage


```hcl
resource "checkpoint_management_network" "network" {
    name = "My Network"
	subnet4 = "10.0.0.0"
	mask_length4 = "24"
}

data "checkpoint_management_data_network" "data_network" {
    name = "${checkpoint_management_network.network.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `subnet4` - IPv4 network address.
* `subnet6` - IPv6 network address.
* `mask_length4` - IPv4 network mask length.
* `mask_length6` - IPv6 network mask length.
* `nat_settings` - NAT settings. nat_settings blocks are documented below.
* `tags` - Collection of tag identifiers.
* `groups` - Collection of group identifiers.
* `broadcast` - Allow broadcast address inclusion.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`nat_settings` supports the following:

* `auto_rule` - Whether to add automatic address translation rules.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `hide_behind` - Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".
* `install_on` - Which gateway should apply the NAT translation.
* `method` - NAT translation method.