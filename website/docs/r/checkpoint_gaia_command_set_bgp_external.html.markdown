---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_bgp_external"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-bgp-external"
description: |-
This resource allows you to execute Check Point Set Bgp External.
---

# checkpoint_gaia_command_set_bgp_external

This resource allows you to execute Check Point Set Bgp External.

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

# Step 2: configure a BGP AS number (required before configuring external peers)
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  as = "65001"

  depends_on = [checkpoint_gaia_command_set_bgp.clear_conf]
}

# Step 3: configure the external peer group
resource "checkpoint_gaia_command_set_bgp_external" "example" {
  remote_as   = "65002"
  enabled     = true
  description = "example desc"
  outdelay    = "70"

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}
```

## Argument Reference

The following arguments are supported:

* `remote_as` - (Required) The Autonomous System number of the peer group to configure.The value can be one of the following: An integer from 1-4294967295 A float from 0.1-65535.65535 
* `enabled` - (Optional) Enable/disable the peer group for the specified AS. 
* `description` - (Optional) Adds a brief description of the peer group. 
* `export_routemap_list` - (Optional) Configure export policy for the given BGP peer group or peer. export_routemap_list blocks are documented below.
* `inject_routemap_list` - (Optional) Configure conditional route injection for a routemap inject_routemap_list blocks are documented below.
* `import_routemap_list` - (Optional) Configure import policy for the given BGP peer group. import_routemap_list blocks are documented below.
* `local_address` - (Optional) Configures the address to be used on the local end of the TCP connection. 
* `outdelay` - (Optional) Specifies the length of time (seconds) that a route must be present in the routing database before it is redistributed to BGP. 


`export_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 
* `conditional_routemap` - (Optional) Condition to apply to the routemap conditional_routemap blocks are documented below.


`inject_routemap_list` supports the following:

* `name` - (Optional) The name of the inject routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `any_pass_routemap` - (Optional) The name of the any-pass-routemap that will be the condition for injection 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`import_routemap_list` supports the following:

* `name` - (Optional) Name of the routemap 
* `preference` - (Optional) Preference for the routemap. Routemaps are evaluated in order of increasing preference value. 
* `family` - (Optional) Describes which family of routes this routemap will be applied to. 


`conditional_routemap` supports the following:

* `name` - (Optional) The name of the routemap condition 
* `condition` - (Optional) The condition can be any-pass or no-pass 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

