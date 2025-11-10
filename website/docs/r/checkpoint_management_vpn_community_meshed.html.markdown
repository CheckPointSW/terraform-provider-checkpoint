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
* `disable_nat` - (Optional) Indicates whether to disable NAT inside the VPN Community.
* `encrypted_traffic` - (Optional) Encrypted traffic settings. encrypted_traffic blocks are documented below.
* `encryption_method` - (Optional) The encryption method to be used. 
* `encryption_suite` - (Optional) The encryption suite to be used. 
* `excluded_services` - (Optional) Collection of services that are excluded from the community identified by the name or UID.<br> Connections with these services will not be encrypted and will not match rules specifying the community in the VPN community.excluded_services blocks are documented below.
* `gateways` - (Optional) Collection of Gateway objects identified by the name or UID.gateways blocks are documented below.
* `ike_phase_1` - (Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.
* `ike_phase_2` - (Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.
* `override_vpn_domains` - (Optional) The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `shared_secrets` - (Optional) Shared secrets for external gateways. shared_secrets blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `use_shared_secret` - (Optional) Indicates whether the shared secret should be used for all external gateways.
* `tunnel_granularity` - (Optional) VPN tunnel sharing option to be used.
* `granular_encryptions` - (Optional) VPN granular encryption settings. granular_encryptions blocks are documented below.
* `permanent_tunnels` - (Optional) Permanent tunnels properties.permanent_tunnels blocks are documented below.
* `wire_mode` - (Optional) VPN Community Wire mode properties.wire_mode blocks are documented below.
* `routing_mode` - (Optional) VPN Community Routing Mode.
* `advanced_properties` - (Optional) Advanced properties.advanced_properties blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

`encrypted_traffic` supports the following:

* `enabled` - (Optional) Indicates whether to accept all encrypted traffic.

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

`permanent_tunnels` supports the following:

* `set_permanent_tunnels` - (Optional) Indicates which tunnels to set as permanent.
* `gateways` - (Optional) List of gateways to set all their tunnels to permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to all-tunnels-of-specific-gateways.gateways blocks are documented below.
* `tunnels` - (Optional) List of tunnels to set as permanent with specified track options. Will take effect only if set-permanent-tunnels-on is set to specific-tunnels-in-the-community.tunnels blocks are documented below.
* `rim` - (Optional) Route Injection Mechanism settings.rim blocks are documented below.
* `tunnel_down_track` - (Optional) VPN community permanent tunnels down track option.
* `tunnel_up_track` - (Optional) Permanent tunnels up track option. 

`wire_mode` supports the following:

* `allow_uninspected_encrypted_traffic` - (Optional) Allow uninspected encrypted traffic between Wire mode interfaces of this Community members.
* `allow_uninspected_encrypted_routing` - (Optional) Allow members to route uninspected encrypted traffic in VPN routing configurations.


`advanced_properties` supports the following:

* `support_ip_compression` - (Optional) Indicates whether to support IP compression.
* `use_aggressive_mode` - (Optional) Indicates whether to use aggressive mode.

`gateways` supports the following:

* `gateway` - (Optional) Gateway to set all is tunnels to permanent with specified track options.<br>
  Identified by name or UID.
* `track_options` - (Optional) Indicates whether to use the community track options or to override track options for the permanent tunnels.
* `override_tunnel_down_track` - (Optional) Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.
* `override_tunnel_up_track` - (Optional) Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.


`tunnels` supports the following:

* `first_tunnel_endpoint` - (Optional) First tunnel endpoint (center gateway).
  Identified by name or UID.
* `second_tunnel_endpoint` - (Optional) Second tunnel endpoint (center gateway for meshed VPN community and satellitegateway for star VPN community).
  Identified by name or UID.
* `track_options` - (Optional) Indicates whether to use the community track options or to override track options for the permanent tunnels.
* `override_tunnel_down_track` - (Optional) Gateway tunnel down track option. Relevant only if the track-options is set to 'override track options'.
* `override_tunnel_up_track` - (Optional) Gateway tunnel up track option. Relevant only if the track-options is set to 'override track options'.


`rim` supports the following:

* `enabled` - (Optional) Indicates whether Route Injection Mechanism is enabled.
* `enable_on_gateways` - (Optional) Indicates whether to enable automatic Route Injection Mechanism for gateways.
* `route_injection_track` - (Optional) Route injection track method. 
