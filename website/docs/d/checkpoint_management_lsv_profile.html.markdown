---
layout: "checkpoint"
page_title: "checkpoint_management_lsv_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-lsv-profile"
description: |-
Use this data source to get information on an existing Check Point Lsv Profile.
---

# Data Source: checkpoint_management_lsv_profile

Use this data source to get information on an existing Check Point Lsv Profile.

## Example Usage


```hcl
resource "checkpoint_management_lsv_profile" "lsv_profile" {
  name = "Lsv profile"
  certificate_authority = "internal_ca"
}


data "checkpoint_management_lsv_profile" "data_lsv_profile" {
    name = "${checkpoint_management_lsv_profile.lsv_profile.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `allowed_ip_addresses` - Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.
* `certificate_authority` - Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.
* `restrict_allowed_addresses` - Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `vpn_domain` - peers' VPN Domain properties. vpn_domain blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.


`vpn_domain` supports the following:

* `limit_peer_domain_size` - Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.
* `max_allowed_addresses` - Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.
