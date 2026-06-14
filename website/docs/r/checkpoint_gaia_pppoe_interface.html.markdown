---
layout: "checkpoint"
page_title: "checkpoint_gaia_pppoe_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-pppoe-interface"
description: |-
This resource allows you to execute Check Point Pppoe Interface.
---

# checkpoint_gaia_pppoe_interface

This resource allows you to execute Check Point Pppoe Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_pppoe_interface" "example" {
  client_id                   = 1
  interface                   = "eth0"
  username                    = "admin"
  password                    = "password"
  use_peer_as_default_gateway = true
  use_peer_dns                = true
  comments                    = "example pppoe interface"
  enabled                     = true
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) The username needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP 
* `client_id` - (Required) The PPPoE client Id. This ID must be unique for every PPPoE interface. 
* `interface` - (Required) The name of the applicable physical interface. Gaia uses this interface to forward PPPoE frames. 
* `name` - (Computed) The PPPoE interface name. 
* `sd_wan` - (Optional) SD-WAN configuration. 
Supported starting from R81.20 JHF 14 sd_wan blocks are documented below.
* `password` - (Optional) The password needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP 
* `password_hash` - (Optional) The hash of the password needed to connect to the PPPoE server at the Internet Service Provider (ISP). Get it from the ISP. 
* `use_peer_as_default_gateway` - (Optional) Enable to make the ISP server the Default Gateway for the Gaia. 
* `use_peer_dns` - (Optional) Enable to allow the ISP to define the IPv4 DNS server for the Gaia. 
* `fake_peer_settings` - (Optional) Fake peer settings fake_peer_settings blocks are documented below.
* `enabled` - (Optional) Enable to turn on the interface. 
* `comments` - (Optional) User comments. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `status` - (Computed) Computed field, returned in the response. 


`sd_wan` supports the following:

* `enabled` - (Optional) Enable SD-WAN on this interface. 
* `next_hop` - (Optional) Configure interface's next hop IPv4 address, obtain next hop IPv4 address automatically         or set as a layer 2-only link 
* `next_hop_ipv6` - (Optional) Configure interface's next hop IPv6 address or obtain next hop IPv6 address automatically.              IPv6 configuration is supported starting from R82 latest Jumbo Hotfix 
* `nat` - (Optional) Optional NAT configuration nat blocks are documented below.
* `tag` - (Optional) Optional tag configuration.             Must contain only alphanumeric characters, '-' or '_' (max length is 64) 
* `bandwidth` - (Optional) Optional Bandwidth configuration.              Bandwidth configuration is supported starting from R81.20 JHF 79 bandwidth blocks are documented below.
* `circuit_id` - (Optional) Optional override interface circuit id value.              Circuit-ID configuration is supported starting from R81.20 JHF 79 


`fake_peer_settings` supports the following:

* `address` - (Optional) The fake unicast peer IPv4 address (the default value is 0.0.0.0). 
* `enabled` - (Optional) Enable to use the configured fake peer IPv4 address. 


`nat` supports the following:

* `enabled` - (Optional) Enable NAT IP address on this interface 
* `ip` - (Optional) Configure NAT IPv4 address on this interface or obtain NAT IPv4 address automatically. 
* `ipv6` - (Optional) Configure NAT IPv6 address on this interface or obtain NAT IPv6 address automatically.              IPv6 configuration is supported starting from R82 latest Jumbo Hotfix 


`bandwidth` supports the following:

* `upload_speed` - (Optional) In Mbps 
* `download_speed` - (Optional) In Mbps 
