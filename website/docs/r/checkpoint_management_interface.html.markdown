---
layout: "checkpoint"
page_title: "checkpoint_management_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-interface"
description: |-
This resource allows you to execute Check Point Interface.
---

# checkpoint_management_interface

This resource allows you to execute Check Point Interface.

## Example Usage


```hcl
resource "checkpoint_management_interface" "example" {
  name = "eth0"
  ipv4_address = "1.1.1.111"
  ipv4_mask_length = 24
  gateway_uid = "20ec49e8-8cd8-4ad4-b204-0de8ae4e0e17"
  topology = "internal"
  cluster_network_type = "cluster"
  anti_spoofing = true
  ignore_warnings = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Network interface name. 
* `gateway_uid` - (Required) Gateway or cluster object uid that the interface belongs to. <font color="red">Required only if</font> name was specified. 
* `anti_spoofing` - (Optional) Enable anti-spoofing. 
* `anti_spoofing_settings` - (Optional) Anti Spoofing Settings.anti_spoofing_settings blocks are documented below.
* `cluster_members` - (Optional) Network interface settings for cluster members.cluster_members blocks are documented below.
* `cluster_network_type` - (Optional) Cluster interface type. 
* `dynamic_ip` - (Optional) Enable dynamic interface. 
* `ipv4_address` - (Optional) IPv4 network address. 
* `ipv4_mask_length` - (Optional) IPv4 mask length. 
* `ipv4_network_mask` - (Optional) IPv4 network mask. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv6_mask_length` - (Optional) IPv6 mask length. 
* `ipv6_network_mask` - (Optional) IPv6 network mask. 
* `monitored_by_cluster` - (Optional) When Private is selected as the Cluster interface type, cluster can monitor or not monitor the interface. 
* `network_interface_type` - (Optional) Network Interface Type. 
* `security_zone_settings` - (Optional) Security Zone Settings.security_zone_settings blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) Topology configuration. 
* `topology_automatic` - Topology configuration automatically calculated by get-interfaces command.
* `topology_manual` - Topology configuration manually defined.
* `topology_settings` - (Optional) Topology Settings.topology_settings blocks are documented below.
* `topology_settings_automatic` -  Topology settings automatically calculated by get-interfaces command.
* `topology_settings_manual` -  Topology settings manually defined.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`anti_spoofing_settings` supports the following:

* `action` - (Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). 
* `exclude_packets` - (Optional) Don't check packets from excluded network. 
* `excluded_network_name` - (Optional) Excluded network name. 
* `excluded_network_uid` - (Optional) Excluded network UID. 
* `spoof_tracking` - (Optional) Spoof tracking. 


`cluster_members` supports the following:

* `name` - (Optional) Cluster member network interface name. 
* `member_uid` - (Optional) Cluster member object uid. 
* `member_name` - (Optional) Cluster member object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 
* `network_mask` - (Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` - (Optional) IPv4 network address. 
* `ipv6_network_mask` - (Optional) IPv6 network address. 
* `ipv4_mask_length` - (Optional) IPv4 network mask length. 
* `ipv6_mask_length` - (Optional) IPv6 network mask length. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`security_zone_settings` supports the following:

* `auto_calculated` - (Optional) Security Zone is calculated according to where the interface leads to. 
* `specific_zone` - (Optional) Security Zone specified manually. 
* `auto_calculated_zone` - (Optional) N/A 
* `auto_calculated_zone_uid` - (Optional) N/A 
* `specific_security_zone_enabled` - (Optional) N/A 


`topology_settings` supports the following:

* `interface_leads_to_dmz` - (Optional) Whether this interface leads to demilitarized zone (perimeter network). 
* `ip_address_behind_this_interface` - (Optional) Network settings behind this interface.
* `specific_network` - (Optional) Network behind this interface. 
* `specific_network_uid` - (Optional) N/A 

`topology_settings_automatic` supports the following:

* `interface_leads_to_dmz` -  Whether this interface leads to demilitarized zone (perimeter network).
* `ip_address_behind_this_interface` - Network settings behind this interface.
* `specific_network` -  Network behind this interface.
* `specific_network_uid` - N/A


`topology_settings_manual` supports the following:

* `interface_leads_to_dmz` -  Whether this interface leads to demilitarized zone (perimeter network).
* `ip_address_behind_this_interface` - Network settings behind this interface.
* `specific_network` -  Network behind this interface.
* `specific_network_uid` - N/A 