---
layout: "checkpoint"
page_title: "checkpoint_management_lsv_profile"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsv-profile"
description: |-
This resource allows you to add/update/delete Check Point Lsv Profile.
---

# Resource: checkpoint_management_lsv_profile

This resource allows you to add/update/delete Check Point Lsv Profile.

## Example Usage


```hcl
resource "checkpoint_management_lsv_profile" "example" {
  name = "Lsv profile"
  certificate_authority = "internal_ca"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `certificate_authority` - (Required) Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.
* `allowed_ip_addresses` - (Optional) Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.
* `restrict_allowed_addresses` - (Optional) Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.
* `tags` - (Optional) Collection of tag identifiers.
* `vpn_domain` - (Optional) peers' VPN Domain properties. vpn_domain blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`vpn_domain` supports the following: 

* `limit_peer_domain_size` - (Optional) Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.
* `max_allowed_addresses` - (Optional) Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.
