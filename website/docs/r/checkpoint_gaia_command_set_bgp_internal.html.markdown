---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_bgp_internal"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-bgp-internal"
description: |-
This resource allows you to execute Check Point Set Bgp Internal.
---

# checkpoint_gaia_command_set_bgp_internal

This resource allows you to execute Check Point Set Bgp Internal.

## Example Usage


```hcl
# Step 1: clear any leftover BGP confederation state
resource "checkpoint_gaia_command_set_bgp" "clear_conf" {
  confederation {
    identifier = "off"
  }
  routing_domain {
    identifier = "off"
  }
}

# Step 2: configure a BGP AS number (required before configuring internal BGP)
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  as = "65001"

  depends_on = [checkpoint_gaia_command_set_bgp.clear_conf]
}

# Step 3: configure internal BGP
resource "checkpoint_gaia_command_set_bgp_internal" "example" {
  enabled             = true
  description         = "test desc"
  enable_nexthop_self = true
  med                 = "40"
  outdelay            = "40"

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Enable Internal BGP (IBGP) on this router. 
* `description` - (Optional) Adds a brief description of the peer group. 
* `export_routemap_list` - (Optional) Configure export policy for the given BGP peer group or peer. export_routemap_list blocks are documented below.
* `import_routemap_list` - (Optional) Configure import policy for the given BGP peer group. import_routemap_list blocks are documented below.
* `interface_list` - (Optional) Specifies the interfaces for which third-party next hops may be used. By default, all interfaces are enabled. interface_list blocks are documented below.
* `local_address` - (Optional) Configures the address to be used on the local end of the TCP connection. 
* `med` - (Optional) Defines the Multi-Exit Discriminator (MED) metric used when advertising routes to all peers in this group. 
* `enable_nexthop_self` - (Optional) When this option is enabled, the router sends its own IP address as the BGP next hop. 
* `outdelay` - (Optional) Specifies the length of time (seconds) that a route must be present in the routing database before it is redistributed to BGP. 
* `protocol_list` - (Optional) Enables specific routing protocols to use as an Interior Gateway Protocol. The possible values that can be used are: all, bgp, direct, rip, static, ospf, ospfase, ospf3, ospf3ase and ripng. By default, all protocols are enabled. protocol_list blocks are documented below.


`export_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 
* `conditional_routemap` - (Optional) Condition to apply to the routemap conditional_routemap blocks are documented below.


`import_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`interface_list` supports the following:



`protocol_list` supports the following:



`conditional_routemap` supports the following:

* `name` - (Optional) The name of the routemap condition 
* `condition` - (Optional) The condition can be any-pass or no-pass 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

