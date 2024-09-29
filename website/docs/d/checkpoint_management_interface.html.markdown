---
layout: "checkpoint"
page_title: "checkpoint_management_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-interface"
description: |-
Use this data source to get information on an existing Check Point Interface.
---

# checkpoint_management_interface

Use this data source to get information on an existing Check Point Interface.

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

data "checkpoint_management_interface" "data" {
  name = "${checkpoint_management_interface.example.name}"
  gateway_uid = "${checkpoint_management_interface.example.gateway_uid}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `gateway_uid` - (Optional) Gateway or cluster object uid that the interface belongs to. <font color="red">Required only if</font> name was specified. 
* `anti_spoofing` - Enable anti-spoofing. 
* `anti_spoofing_settings` -  Anti Spoofing Settings.anti_spoofing_settings blocks are documented below.
* `cluster_members` -  Network interface settings for cluster members.cluster_members blocks are documented below.
* `cluster_network_type` -Cluster interface type. 
* `dynamic_ip` -  Enable dynamic interface. 
* `ipv4_address` -IPv4 network address. 
* `ipv4_mask_length` -  IPv4 mask length. 
* `ipv4_network_mask` -  IPv4 network mask. 
* `ipv6_address` - IPv6 address. 
* `ipv6_mask_length` -  IPv6 mask length. 
* `ipv6_network_mask` - IPv6 network mask. 
* `monitored_by_cluster` - When Private is selected as the Cluster interface type, cluster can monitor or not monitor the interface. 
* `network_interface_type` -  Network Interface Type. 
* `security_zone_settings` -Security Zone Settings.security_zone_settings blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `topology` - Topology configuration. 
* `topology_automatic` - Topology configuration automatically calculated by get-interfaces command.
* `topology_manual` - Topology configuration manually defined.
* `topology_settings` -  Topology Settings.topology_settings blocks are documented below.
* `topology_settings_automatic` -  Topology settings automatically calculated by get-interfaces command.
* `topology_settings_manual` -  Topology settings manually defined.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 


`anti_spoofing_settings` supports the following:

* `action` -  If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). 
* `exclude_packets` -  Don't check packets from excluded network. 
* `excluded_network_name` -  Excluded network name. 
* `excluded_network_uid` -  Excluded network UID. 
* `spoof_tracking` -  Spoof tracking. 


`cluster_members` supports the following:

* `name` -  Cluster member network interface name. 
* `member_uid` -  Cluster member object uid. 
* `member_name` -  Cluster member object name. 
* `ipv4_address` -  IPv4 address. 
* `ipv6_address` -  IPv6 address. 
* `network_mask` -  IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. 
* `ipv4_network_mask` -  IPv4 network address. 
* `ipv6_network_mask` -  IPv6 network address. 
* `ipv4_mask_length` - IPv4 network mask length. 
* `ipv6_mask_length` -  IPv6 network mask length. 



`security_zone_settings` supports the following:

* `auto_calculated` -  Security Zone is calculated according to where the interface leads to. 
* `specific_zone` -  Security Zone specified manually. 
* `auto_calculated_zone` -  N/A 
* `auto_calculated_zone_uid` -  N/A 
* `specific_security_zone_enabled` -  N/A 


`topology_settings` supports the following:

* `interface_leads_to_dmz` -  Whether this interface leads to demilitarized zone (perimeter network). 
* `ip_address_behind_this_interface` - Network settings behind this interface.
* `specific_network` -  Network behind this interface. 
* `specific_network_uid` - N/A 

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
