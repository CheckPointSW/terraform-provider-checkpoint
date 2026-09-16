---
layout: "checkpoint"
page_title: "checkpoint_management_vpn_community_remote_access"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-vpn-community-remote-access"
description: |-
This resource allows you to execute Check Point VPN Community Remote Access.
---

# Resource: checkpoint_management_vpn_community_remote_access

This resource allows you to execute Check Point VPN Community Remote Access.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_remote_access" "example" {
    name = "RemoteAccess"
    user_groups = ["All Users"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `gateways` - (Optional) Collection of Gateway objects identified by the name or UID.
* `user_groups` - (Optional) Collection of User group objects identified by the name or UID.
* `override_vpn_domains` - (Optional) The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `encryption` - (Optional) Encryption settings.encryption blocks are documented below.

`override_vpn_domains` supports the following:

* `gateway` - Participant gateway in override VPN domain identified by the name or UID.
* `vpn_domain` - VPN domain network identified by the name or UID. 


`encryption` supports the following:

* `encryption_method` - (Optional) The encryption method to be used.
* `ike_phase_1` - (Optional) IKE Phase 1 settings.ike_phase_1 blocks are documented below.
* `ike_phase_2` - (Optional) IKE Phase 2 settings.ike_phase_2 blocks are documented below.


`ike_phase_1` supports the following:

* `configured_data_integrity_algorithms` - (Optional) The list of configured hash algorithms.
* `configured_diffie_hellman_groups` - (Optional) The list of configured Diffie-Hellman groups.
* `configured_encryption_algorithms` - (Optional) The list of configured encryption algorithms.
* `data_integrity` - (Optional) The hash algorithm to be used.
* `diffie_hellman_group` - (Optional) The Diffie-Hellman group to be used.
* `encryption_algorithm` - (Optional) The encryption algorithm to be used.
* `ike_p1_rekey_time` - (Optional) Indicates the time interval for IKE phase 1 renegotiation.
* `ike_p1_rekey_time_unit` - (Optional) Indicates the time unit for the 'ike-p1-rekey-time-unit' parameter, rounded up to minutes scale.
* `multiple_key_exchanges` - (Optional) Name of the Multiple Key Exchanges proposal object.
* `use_multiple_key_exchanges` - (Optional) Indicates whether to use a proposal with Multiple Key Exchanges.
* `use_standard_proposal` - (Optional) Indicates whether to use a proposal with a single Diffie-Hellman group.


`ike_phase_2` supports the following:

* `configured_data_integrity_algorithms` - (Optional) The list of configured hash algorithms.
* `configured_diffie_hellman_groups` - (Optional) The list of configured Diffie-Hellman groups.
* `configured_encryption_algorithms` - (Optional) The list of configured encryption algorithms.
* `data_integrity` - (Optional) The hash algorithm to be used.
* `encryption_algorithm` - (Optional) The encryption algorithm to be used.
* `enforce_encryption_alg_and_data_integrity_on_all_users` - (Optional) Enforce Encryption Algorithm and Data Integrity on all users.
* `ike_p2_pfs_dh_grp` - (Optional) The Diffie-Hellman group to be used.
* `ike_p2_rekey_time` - (Optional) Indicates the time interval for IKE phase 2 renegotiation.
* `ike_p2_rekey_time_unit` - (Optional) Indicates the time unit for the 'ike-p2-rekey-time-unit' parameter, rounded up to minutes scale.
* `ike_p2_use_pfs` - (Optional) Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.
* `multiple_key_exchanges` - (Optional) Name of the Multiple Key Exchanges proposal object to use when PFS is enabled and multiple key exchanges are configured.
* `use_multiple_key_exchanges` - (Optional) Indicates whether to use a proposal with Multiple Key Exchanges when PFS is enabled.
* `use_standard_proposal` - (Optional) Indicates whether to use a proposal with a single Diffie-Hellman group when PFS is enabled.
