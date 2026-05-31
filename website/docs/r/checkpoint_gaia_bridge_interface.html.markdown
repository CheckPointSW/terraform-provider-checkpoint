---
layout: "checkpoint"
page_title: "checkpoint_gaia_bridge_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-bridge-interface"
description: |-
This resource allows you to execute Check Point Bridge Interface.
---

# checkpoint_gaia_bridge_interface

This resource allows you to execute Check Point Bridge Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_bridge_interface" "example" {
  resource_id = 9
  comments    = "Bridge9"
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required)  
* `name` - (Optional, Computed)  
* `dhcp6` - (Optional) DHCPv6 configuration dhcp6 blocks are documented below.
* `dhcp` - (Optional) DHCP configuration dhcp blocks are documented below.
* `mtu` - (Optional)  
* `ipv4_address` - (Optional)  
* `ipv4_mask_length` - (Optional)  
* `enabled` - (Optional)  
* `ipv6_autoconfig` - (Optional)  
* `comments` - (Optional)  
* `ipv6_address` - (Optional)  
* `ipv6_mask_length` - (Optional)  
* `members` - (Optional)  members blocks are documented below.
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `link_state` - (Computed) Computed field, returned in the response. 
* `speed` - (Computed) Computed field, returned in the response. 
* `duplex` - (Computed) Computed field, returned in the response. 
* `tx_bytes` - (Computed) Computed field, returned in the response. 
* `tx_packets` - (Computed) Computed field, returned in the response. 
* `rx_bytes` - (Computed) Computed field, returned in the response. 
* `rx_packets` - (Computed) Computed field, returned in the response. 
* `ipv6_local_link_address` - (Computed) Computed field, returned in the response. 
* `status` - (Computed) Computed field, returned in the response. status blocks are documented below.


`dhcp6` supports the following:

* `enabled` - (Optional) Enable DHCP on this interface. 
* `server_timeout` - (Optional) Specifies the amount of time, in seconds, that must pass between the time that the interface begins to try to determine its address and the time that it decides that it's not going to be able to contact a server. 
* `retry` - (Optional) Specifies the time, in seconds, that must pass after the interface has determined that there is no DHCP server present before it tries again to contact a DHCP server. 
* `leasetime` - (Optional) Specifies the lease time, in seconds, when requesting for an IP address. Default value is "default" - according to the server. 
* `reacquire_timeout` - (Optional) When trying to reacquire the last IP address, the reacquire-timeout statement sets the time, in seconds, that must elapse after the first try to reacquire the old address before it gives up and tries to discover a new address. 
* `using` - (Optional) Choose the DHCPv6 client working mode of this interface.          Interface will receive IPv6 only if the chosen mode and the system's configured mode match 


`dhcp` supports the following:

* `enabled` - (Optional) Enable DHCP on this interface. 
* `server_timeout` - (Optional) Specifies the amount of time, in seconds, that must pass between the time that the interface begins to try to determine its address and the time that it decides that it's not going to be able to contact a server. 
* `retry` - (Optional) Specifies the time, in seconds, that must pass after the interface has determined that there is no DHCP server present before it tries again to contact a DHCP server. 
* `leasetime` - (Optional) Specifies the lease time, in seconds, when requesting for an IP address. Default value is "default" - according to the server. 
* `reacquire_timeout` - (Optional) When trying to reacquire the last IP address, the reacquire-timeout statement sets the time, in seconds, that must elapse after the first try to reacquire the old address before it gives up and tries to discover a new address. 


`status` supports the following:

* `link_state` - (Computed) Computed field, returned in the response. 
* `speed` - (Computed) Computed field, returned in the response. 
* `duplex` - (Computed) Computed field, returned in the response. 
* `tx_bytes` - (Computed) Computed field, returned in the response. 
* `tx_packets` - (Computed) Computed field, returned in the response. 
* `rx_bytes` - (Computed) Computed field, returned in the response. 
* `rx_packets` - (Computed) Computed field, returned in the response. 
