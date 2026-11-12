---
layout: "checkpoint"
page_title: "checkpoint_management_vpn_community_remote_access"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-vpn-community-remote-access"
description: |-
This resource allows you to execute Check Point VPN Community Remote Access.
---

# Data Source: checkpoint_management_vpn_community_remote_access

This resource allows you to execute Check Point VPN Community Remote Access.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_remote_access" "vpn_community_remote_access" {
    name = "RemoteAccess"
	user_groups = ["All Users"]
}

data "checkpoint_management_vpn_community_remote_access" "data_vpn_community_remote_access" {
    name = "${checkpoint_management_vpn_community_remote_access.vpn_community_remote_access.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `gateways` - Collection of Gateway objects identified by the name or UID.
* `user_groups` - Collection of User group objects identified by the name or UID.
* `override_vpn_domains` - The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. 
* `comments` - Comments string.
* `encryption` - Encryption settings.encryption blocks are documented below.

`override_vpn_domains` supports the following:

* `gateway` - Participant gateway in override VPN domain identified by the name or UID.
* `vpn_domain` - VPN domain network identified by the name or UID. 


`encryption` supports the following:

* `encryption_method` - The encryption method to be used.
* `ike_phase_1` - IKE Phase 1 settings.ike_phase_1 blocks are documented below.
* `ike_phase_2` - IKE Phase 2 settings.ike_phase_2 blocks are documented below.


`ike_phase_1` supports the following:

* `configured_data_integrity_algorithms` - The list of configured hash algorithms.
* `configured_diffie_hellman_groups` - The list of configured Diffie-Hellman groups.
* `configured_encryption_algorithms` - The list of configured encryption algorithms.
* `data_integrity` - The hash algorithm to be used.
* `diffie_hellman_group` - The Diffie-Hellman group to be used.
* `encryption_algorithm` - The encryption algorithm to be used.
* `ike_p1_rekey_time` - Indicates the time interval for IKE phase 1 renegotiation.
* `multiple_key_exchanges` - Multiple Key Exchanges proposal object.multiple_key_exchanges blocks are documented below.
* `use_multiple_key_exchanges` - Indicates whether to use a proposal with Multiple Key Exchanges.
* `use_standard_proposal` - Indicates whether to use a proposal with a single Diffie-Hellman group.


`multiple_key_exchanges` supports the following:

* `additional_key_exchange_1_methods` - Additional Key-Exchange 1 methods to use.
* `additional_key_exchange_2_methods` - Additional Key-Exchange 2 methods to use.
* `additional_key_exchange_3_methods` - Additional Key-Exchange 3 methods to use.
* `additional_key_exchange_4_methods` - Additional Key-Exchange 4 methods to use.
* `additional_key_exchange_5_methods` - Additional Key-Exchange 5 methods to use.
* `additional_key_exchange_6_methods` - Additional Key-Exchange 6 methods to use.
* `additional_key_exchange_7_methods` - Additional Key-Exchange 7 methods to use.
* `key_exchange_methods` - Key-Exchange methods to use. Can contain only Diffie-Hellman groups.


`ike_phase_2` supports the following:

* `configured_data_integrity_algorithms` - The list of configured hash algorithms.
* `configured_diffie_hellman_groups` - The list of configured Diffie-Hellman groups.
* `configured_encryption_algorithms` - The list of configured encryption algorithms.
* `data_integrity` - The hash algorithm to be used.
* `encryption_algorithm` - The encryption algorithm to be used.
* `enforce_encryption_alg_and_data_integrity_on_all_users` - Enforce Encryption Algorithm and Data Integrity on all users.
* `ike_p2_pfs_dh_grp` - The Diffie-Hellman group to be used.
* `ike_p2_rekey_time` - Indicates the time interval for IKE phase 2 renegotiation.
* `ike_p2_use_pfs` - Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.
* `multiple_key_exchanges` - Multiple Key Exchanges proposal object to use when PFS is enabled and multiple key exchanges are configured.multiple_key_exchanges blocks are documented below.
* `use_multiple_key_exchanges` - Indicates whether to use a proposal with Multiple Key Exchanges when PFS is enabled.
* `use_standard_proposal` - Indicates whether to use a proposal with a single Diffie-Hellman group when PFS is enabled.


`multiple_key_exchanges` supports the following:

* `additional_key_exchange_1_methods` - Additional Key-Exchange 1 methods to use.
* `additional_key_exchange_2_methods` - Additional Key-Exchange 2 methods to use.
* `additional_key_exchange_3_methods` - Additional Key-Exchange 3 methods to use.
* `additional_key_exchange_4_methods` - Additional Key-Exchange 4 methods to use.
* `additional_key_exchange_5_methods` - Additional Key-Exchange 5 methods to use.
* `additional_key_exchange_6_methods` - Additional Key-Exchange 6 methods to use.
* `additional_key_exchange_7_methods` - Additional Key-Exchange 7 methods to use.
* `key_exchange_methods` - Key-Exchange methods to use. Can contain only Diffie-Hellman groups.
