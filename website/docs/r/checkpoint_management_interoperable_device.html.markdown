---
layout: "checkpoint"
page_title: "checkpoint_management_interoperable_device"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-interoperable-device"
description: |-
This resource allows you to execute Check Point Interoperable Device.
---

# checkpoint_management_interoperable_device

This resource allows you to execute Check Point Interoperable Device.

## Example Usage


```hcl
resource "checkpoint_management_interoperable_device" "example" {
  name = "NewInteroperableDevice"
  ipv4_address = "192.168.1.6"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address of the Interoperable Device. 
* `ipv6_address` - (Optional) IPv6 address of the Interoperable Device. 
* `interfaces` - (Optional) Network interfaces.interfaces blocks are documented below.
* `vpn_settings` - (Optional) VPN domain properties for the Interoperable Device.vpn_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`interfaces` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `network_mask` - (Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` - (Optional) IPv4 network address. 
* `ipv6_network_mask` - (Optional) IPv6 network address. 
* `ipv4_mask_length` - (Optional) IPv4 network mask length. 
* `ipv6_mask_length` - (Optional) IPv6 network mask length. 
* `anti_spoofing` - (Optional) Is anti spoofing enabled on the interface. 
* `anti_spoofing_settings` - (Optional) Anti spoofing settings.anti_spoofing_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) Topology configuration. 
* `topology_settings` - (Optional) Internal topology settings.topology_settings blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`vpn_settings` supports the following:

* `vpn_domain` - (Optional) Network group representing the customized encryption domain. Must be set when vpn-domain-type is set to 'manual' option. 
* `vpn_domain_exclude_external_ip_addresses` - (Optional) Exclude the external IP addresses from the VPN domain of this Interoperable device. 
* `vpn_domain_type` - (Optional) Indicates the encryption domain. 


`anti_spoofing_settings` supports the following:

* `action` - (Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). 
* `exclude_packets` - (Optional) Don't check packets from excluded network. 
* `excluded_network_name` - (Optional) Excluded network name. 
* `excluded_network_uid` - (Optional) Excluded network UID. 
* `spoof_tracking` - (Optional) Spoof tracking. 


`topology_settings` supports the following:

* `interface_leads_to_dmz` - (Optional) Whether this interface leads to demilitarized zone (perimeter network). 
* `specific_network` - (Optional) Network behind this interface. 
