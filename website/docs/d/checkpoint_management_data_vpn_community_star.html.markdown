---
layout: "checkpoint"
page_title: "checkpoint_management_data_vpn_community_star"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-vpn-community-star"
description: |-
  Use this data source to get information on an existing Check Point Vpn Community Star.
---

# checkpoint_management_data_vpn_community_star

Use this data source to get information on an existing Check Point Vpn Community Star.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_star" "vpn_community_star" {
    name = "%s"
	encryption_method = "ikev1 for ipv4 and ikev2 for ipv6 only"
	encryption_suite = "custom"
}

data "checkpoint_management_data_vpn_community_star" "data_vpn_community_star" {
    name = "${checkpoint_management_vpn_community_star.vpn_community_star.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `center_gateways` - Collection of Gateway objects representing center gateways identified by the name or UID. center_gateways blocks are documented below.
* `encryption_method` - The encryption method to be used. 
* `encryption_suite` - The encryption suite to be used. 
* `ike_phase_1` - Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.
* `ike_phase_2` - Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.
* `mesh_center_gateways` - Indicates whether the meshed community is in center. 
* `override_vpn_domains` - The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `satellite_gateways` - Collection of Gateway objects representing satellite gateways identified by the name or UID. satellite_gateways blocks are documented below.
* `shared_secrets` - Shared secrets for external gateways. shared_secrets blocks are documented below.
* `tags` - Collection of tag identifiers.
* `use_shared_secret` - Indicates whether the shared secret should be used for all external gateways. 
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `ignore_warnings` - Apply changes ignoring warnings. 
* `ignore_errors` - Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`ike_phase_1` supports the following:

* `data_integrity` - The hash algorithm to be used. 
* `diffie_hellman_group` - The Diffie-Hellman group to be used. 
* `encryption_algorithm` - The encryption algorithm to be used. 


`ike_phase_2` supports the following:

* `data_integrity` - The hash algorithm to be used. 
* `encryption_algorithm` - The encryption algorithm to be used. 


`override_vpn_domains` supports the following:

* `gateway` - Participant gateway in override VPN domain identified by the name or UID. 
* `vpn_domain` - VPN domain network identified by the name or UID. 


`shared_secrets` supports the following:

* `external_gateway` - External gateway identified by the name or UID. 
* `shared_secret` - Shared secret. 