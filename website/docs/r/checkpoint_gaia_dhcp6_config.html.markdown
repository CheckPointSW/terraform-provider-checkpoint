---
layout: "checkpoint"
page_title: "checkpoint_gaia_dhcp6_config"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-dhcp6-config"
description: |-
This resource allows you to execute Check Point Dhcp6 Config.
---

# checkpoint_gaia_dhcp6_config

This resource allows you to execute Check Point Dhcp6 Config.

## Example Usage


```hcl
resource "checkpoint_gaia_dhcp6_config" "example" {
  client_mode = "prefix-delegation"
}
```

## Argument Reference

The following arguments are supported:

* `prefix_delegation_options` - (Optional) General configuration for the prefix-delegation feature. prefix_delegation_options blocks are documented below.
* `client_mode` - (Optional) The working mode of the DHCPv6 client in this system. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`prefix_delegation_options` supports the following:

* `interface` - (Optional) The interface on which to send prefix-delegation request packets to the prefix-delegation DHCP server. 
* `method` - (Optional) The method of performing the delegation of the received subnets. Each method balances automation with granularity.                        <b>Manual</b> - Only configure client interfaces set to receive IPv6 via Prefix-Delegation.                        <b>Router Discovery</b> - In addition to IPv6, also automatically configure Router Discovery protocol                         on configured interfaces.                        <b>DHCPv6</b> - In addition to IPv6, also automatically configure the DHCPv6 Server feature on                         the new subnets and configure the Router Discory protocol with Managed Configuration flag on. 
* `suffix_pools` - (Optional) Pools of IPv6 suffixes to use with DHCPv6 delegation method.                        These will be used to automatically configure IPv6 pools for each subnet in the DHCPv6 server feature. suffix_pools blocks are documented below.


`suffix_pools` supports the following:

* `start` - (Optional) The first IPv6 address of the suffix range. 
* `end` - (Optional) The last IPv6 address of the suffix range. 
* `type` - (Optional) Specifies whether to include or exclude this range of IPv6 suffixes in the IP pools. 
