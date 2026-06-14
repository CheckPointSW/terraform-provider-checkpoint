---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_bgp_confederation"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-bgp-confederation"
description: |-
This resource allows you to execute Check Point Set Bgp Confederation.
---

# checkpoint_gaia_command_set_bgp_confederation

This resource allows you to execute Check Point Set Bgp Confederation.

## Example Usage


```hcl
# Step 1: configure BGP confederation and routing-domain identifiers
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  confederation {
    identifier = "65000"
  }
  routing_domain {
    identifier = "1234.4321"
  }
}

# Step 2: configure the confederation peer group settings
resource "checkpoint_gaia_command_set_bgp_confederation" "example" {
  member_as           = "1234.4321"
  enabled             = true
  description         = "test desc"
  enable_nexthop_self = true
  med                 = "1000"
  outdelay            = "123"

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}
```

## Argument Reference

The following arguments are supported:

* `member_as` - (Required) Specify the Routing Domain identifier of the Confederation peer group to configure.  If the peer group specified is the local Routing Domain, it will run IBGP in a full mesh (just as an internal peer group normally would in non-Confederation mode). Otherwise, if an external Routing Domain within the Confederation is specified, the peer group will run a modified version of eBGP, which preserves route metrics and other BGP attributes.  The value can be one of the following: An integer from 1-4294967295 A float from 0.1-65535.65535 
* `enabled` - (Optional) Enable/disable the peer group for the specified AS. 
* `description` - (Optional) Adds a brief description of the peer group. 
* `interface_list` - (Optional) Specifies the interfaces for which third-party next hops may be used. By default, all interfaces are enabled. interface_list blocks are documented below.
* `local_address` - (Optional) Configures the address to be used on the local end of the TCP connection. 
* `med` - (Optional) Defines the Multi-Exit Discriminator (MED) metric used when advertising routes to all peers in this group. 
* `enable_nexthop_self` - (Optional) When this option is enabled, the router sends its own IP address as the BGP next hop. 
* `outdelay` - (Optional) Specifies the length of time (seconds) that a route must be present in the routing database before it is redistributed to BGP. 
* `protocol_list` - (Optional) Enables specific routing protocols to use as an Interior Gateway Protocol. The possible values that can be used are: all, bgp, direct, rip, static, ospf, ospfase, ospf3, ospf3ase and ripng. By default, all protocols are enabled. protocol_list blocks are documented below.


`interface_list` supports the following:



`protocol_list` supports the following:



## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

