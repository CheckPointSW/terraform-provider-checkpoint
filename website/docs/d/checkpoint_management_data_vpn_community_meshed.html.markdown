---
layout: "checkpoint"
page_title: "checkpoint_management_data_vpn_community_meshed"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-vpn-community-meshed"
description: |-
  Use this data source to get information on an existing Check Point Vpn Community Meshed.
---

# Data Source: checkpoint_management_data_vpn_community_meshed

Use this data source to get information on an existing Check Point Vpn Community Meshed.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_meshed" "vpn_community_meshed" {
    name = "vpn community meshed"
	encryption_method = "ikev1 for ipv4 and ikev2 for ipv6 only"
	encryption_suite = "custom"
}

data "checkpoint_management_data_vpn_community_meshed" "data_vpn_community_meshed" {
    name = "${checkpoint_management_vpn_community_meshed.vpn_community_meshed.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `encryption_method` - The encryption method to be used. 
* `encryption_suite` - The encryption suite to be used. 
* `gateways` - Collection of Gateway objects identified by the name or UID. gateways blocks are documented below.
* `ike_phase_1` - Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.
* `ike_phase_2` - Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.
* `override_vpn_domains` - The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `shared_secrets` - Shared secrets for external gateways. shared_secrets blocks are documented below.
* `tags` - Collection of tag identifiers.
* `use_shared_secret` - Indicates whether the shared secret should be used for all external gateways. 
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 


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