---
layout: "checkpoint"
page_title: "checkpoint_gaia_dhcp6_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-dhcp6-server"
description: |-
This resource allows you to execute Check Point Dhcp6 Server.
---

# checkpoint_gaia_dhcp6_server

This resource allows you to execute Check Point Dhcp6 Server.

## Example Usage


```hcl
resource "checkpoint_gaia_dhcp6_server" "example" {
  enabled = false
  subnets {
    subnet = "2222:0:1::"
    prefix = 48
    max_lease = 86400
    default_lease = 43200
    enabled = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) DHCPv6 server status 
* `subnets` - (Optional) Subnets subnets blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`subnets` supports the following:

* `subnet` - (Optional) Subnet name. 
* `prefix` - (Optional) Subnet prefix length. 
* `enabled` - (Optional) Enable DHCP on this subnet. 
* `max_lease` - (Optional) The longest lease that the server can allocate, in seconds. 
* `default_lease` - (Optional) The default lease that the server allocates, in seconds. 
* `dns` - (Optional) DNS configuration. dns blocks are documented below.
* `ip_pools` - (Optional) Range of IPv6 addresses that the server assigns to hosts. ip_pools blocks are documented below.


`dns` supports the following:

* `primary` - (Optional) The IPv6 address of the Primary DNS server for the DHCP clients. If not exist, empty string will be returned. 
* `secondary` - (Optional) The IPv6 address of the Secondary DNS server for the DHCP clients (to use if the primary DNS server does not respond). If not exist, empty string will be returned. 
* `tertiary` - (Optional) The IPv6 address of the Tertiary DNS server for the DHCP clients (to use if the primary and secondary DNS servers do not respond). If not exist, empty string will be returned. 
* `domain_name` - (Optional) The domain name of the IPv6 subnet. If not exist, empty string will be returned. 


`ip_pools` supports the following:

* `enabled` - (Optional) Enables or disables the DHCP Server for this subnet IP pool. 
* `include` - (Optional) Specifies whether to include or exclude this range of IPv4 addresses in the IP pool. 
* `start` - (Optional) The first IPv6 address of the range. 
* `end` - (Optional) The last IPv6 address of the range. 
