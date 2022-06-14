---
layout: "checkpoint"
page_title: "checkpoint_management_vpn_community_meshed"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-vpn-community-meshed"
description: |-
  This resource allows you to execute Check Point Vpn Community Meshed.
---

# Resource: checkpoint_management_vpn_community_meshed

This resource allows you to execute Check Point Vpn Community Meshed.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_meshed" "example" {
  name = "New_VPN_Community_Meshed_1"
  encryption_method = "prefer ikev2 but support ikev1"
  encryption_suite = "custom"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `encryption_method` - (Optional) The encryption method to be used. 
* `encryption_suite` - (Optional) The encryption suite to be used. 
* `gateways` - (Optional) Collection of Gateway objects identified by the name or UID.gateways blocks are documented below.
* `ike_phase_1` - (Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.
* `ike_phase_2` - (Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.
* `override_vpn_domains` - (Optional) The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `shared_secrets` - (Optional) Shared secrets for external gateways. shared_secrets blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `use_shared_secret` - (Optional) Indicates whether the shared secret should be used for all external gateways.
* `tunnel_granularity` - (Optional) VPN tunnel sharing option to be used.
* `granular_encryptions` - (Optional) VPN granular encryption settings. granular_encryptions blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`override_vpn_domains` supports the following:

* `gateway` - (Optional) Participant gateway in override VPN domain identified by the name or UID. 
* `vpn_domain` - (Optional) VPN domain network identified by the name or UID. 


`shared_secrets` supports the following:

* `external_gateway` - (Optional) External gateway identified by the name or UID. 
* `shared_secret` - (Optional) Shared secret. 

`granular_encryptions` supports the following:

* `internal-gateway` - (Required) Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.
* `external_gateway` - (Required) Externally managed or 3rd party gateway identified by name or UID.
* `encryption_method` - (Optional) The encryption method to be used: prefer ikev2 but support ikev1, ikev2 only, ikev1 for ipv4 and ikev2 for ipv6 only.
* `encryption_suite` - (Optional) The encryption suite to be used: suite-b-gcm-256, custom, vpn b, vpn a, suite-b-gcm-128.
* `ike_phase_1` - (Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.
* `ike_phase_2` - (Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.

`ike_phase_1` supports the following:

* `data_integrity` - (Optional) The hash algorithm to be used.
* `diffie_hellman_group` - (Optional) The Diffie-Hellman group to be used.
* `encryption_algorithm` - (Optional) The encryption algorithm to be used.
* `ike_p1_rekey_time` - (Optional) Indicates the time interval for IKE phase 1 renegotiation.


`ike_phase_2` supports the following:

* `data_integrity` - (Optional) The hash algorithm to be used.
* `encryption_algorithm` - (Optional) The encryption algorithm to be used.
* `ike_p2_use_pfs` - (Optional) Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.
* `ike_p2_pfs_dh_grp` - (Optional) The Diffie-Hellman group to be used: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24.
* `ike_p2_rekey_time` - (Optional) Indicates the time interval for IKE phase 2 renegotiation.