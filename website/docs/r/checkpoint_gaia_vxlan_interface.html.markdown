---
layout: "checkpoint"
page_title: "checkpoint_gaia_vxlan_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-vxlan-interface"
description: |-
This resource allows you to execute Check Point Vxlan Interface.
---

# checkpoint_gaia_vxlan_interface

This resource allows you to execute Check Point Vxlan Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_vxlan_interface" "example" {
  vxlan_id          = 100
  member_of         = "eth0"
  remote_ip_address = "10.0.0.1"
  comments          = "example vxlan tunnel"
}
```

## Argument Reference

The following arguments are supported:

* `vxlan_id` - (Required) The Vxlan Network Identifier (aka VNI) 
* `member_of` - (Required) The physical device used for tunnel endpoint communication 
* `remote_ip_address` - (Required) Vxlan remote address 
* `name` - (Optional, Computed) Existing Vxlan Interface name 
* `ipv4_address` - (Optional)  
* `ipv4_mask_length` - (Optional)  
* `enabled` - (Optional)  
* `ipv6_autoconfig` - (Optional)  
* `comments` - (Optional)  
* `ipv6_address` - (Optional)  
* `ipv6_mask_length` - (Optional)  
* `destination_port` - (Optional) The UDP destination port to communicate to the remote Vxlan tunnel endpoint 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `mtu` - (Optional)  
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


`status` supports the following:

* `link_state` - (Computed) Computed field, returned in the response. 
* `speed` - (Computed) Computed field, returned in the response. 
* `duplex` - (Computed) Computed field, returned in the response. 
* `tx_bytes` - (Computed) Computed field, returned in the response. 
* `tx_packets` - (Computed) Computed field, returned in the response. 
* `rx_bytes` - (Computed) Computed field, returned in the response. 
* `rx_packets` - (Computed) Computed field, returned in the response. 
