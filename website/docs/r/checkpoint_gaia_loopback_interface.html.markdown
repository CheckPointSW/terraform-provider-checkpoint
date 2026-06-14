---
layout: "checkpoint"
page_title: "checkpoint_gaia_loopback_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-loopback-interface"
description: |-
This resource allows you to execute Check Point Loopback Interface.
---

# checkpoint_gaia_loopback_interface

This resource allows you to execute Check Point Loopback Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_loopback_interface" "example" {
  name             = "loop00"
  ipv4_address     = "1.2.3.4"
  ipv4_mask_length = 32
  enabled          = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, Computed)  
* `ipv4_address` - (Optional) Either this, 'ipv6-address' or both must be specified 
* `ipv4_mask_length` - (Optional)  
* `enabled` - (Optional)  
* `ipv6_autoconfig` - (Optional)  
* `comments` - (Optional)  
* `ipv6_address` - (Optional) Either this, 'ipv4-address' or both must be specified 
* `ipv6_mask_length` - (Optional)  
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `link_state` - (Computed) Computed field, returned in the response. 
* `speed` - (Computed) Computed field, returned in the response. 
* `duplex` - (Computed) Computed field, returned in the response. 
* `tx_bytes` - (Computed) Computed field, returned in the response. 
* `tx_packets` - (Computed) Computed field, returned in the response. 
* `rx_bytes` - (Computed) Computed field, returned in the response. 
* `rx_packets` - (Computed) Computed field, returned in the response. 
* `mtu` - (Computed) Computed field, returned in the response. 
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
