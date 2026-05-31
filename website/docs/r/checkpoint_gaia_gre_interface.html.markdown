---
layout: "checkpoint"
page_title: "checkpoint_gaia_gre_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-gre-interface"
description: |-
This resource allows you to execute Check Point Gre Interface.
---

# checkpoint_gaia_gre_interface

This resource allows you to execute Check Point Gre Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_gre_interface" "example" {
  gre_id            = 1
  local_ip_address  = "172.23.22.31"
  remote_ip_address = "10.0.0.1"
  ttl               = 64
  ipv4_address      = "192.168.100.1"
  ipv4_mask_length  = 30
  peer_address      = "10.0.0.1"
  enabled           = true
}
```

## Argument Reference

The following arguments are supported:

* `gre_id` - (Required) ID number represents the tunnel ID. 
* `local_ip_address` - (Required) IP address of the underlying local interface on this gateway. 
* `remote_ip_address` - (Required) IP address of the underlying remote interface on the router on the other end of the tunnel. 
* `ttl` - (Required)  
* `ipv4_address` - (Required) Assigned IP of the GRE. 
* `ipv4_mask_length` - (Required)  
* `peer_address` - (Required) IP address of the remote peer. 
* `name` - (Optional, Computed)  
* `comments` - (Optional)  
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `mtu` - (Optional)  
* `enabled` - (Optional)  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `link_state` - (Computed) Computed field, returned in the response. 
* `speed` - (Computed) Computed field, returned in the response. 
* `duplex` - (Computed) Computed field, returned in the response. 
* `tx_bytes` - (Computed) Computed field, returned in the response. 
* `tx_packets` - (Computed) Computed field, returned in the response. 
* `rx_bytes` - (Computed) Computed field, returned in the response. 
* `rx_packets` - (Computed) Computed field, returned in the response. 
* `ipv6_autoconfig` - (Computed) Computed field, returned in the response. 
* `ipv6_address` - (Computed) Computed field, returned in the response. 
* `ipv6_mask_length` - (Computed) Computed field, returned in the response. 
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
