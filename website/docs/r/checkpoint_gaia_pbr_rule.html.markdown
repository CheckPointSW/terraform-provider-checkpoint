---
layout: "checkpoint"
page_title: "checkpoint_gaia_pbr_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-pbr-rule"
description: |-
This resource allows you to execute Check Point Pbr Rule.
---

# checkpoint_gaia_pbr_rule

This resource allows you to execute Check Point Pbr Rule.

## Example Usage


```hcl
resource "checkpoint_gaia_pbr_rule" "example" {
  priority = 1

  match {
    interface = "eth0"
    protocol  = "tcp"
    port      = 80
    source {
      address     = "10.0.0.0"
      mask_length = 24
    }
    destination {
      address     = "10.1.0.0"
      mask_length = 24
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `priority` - (Required) PBR Rule Priority 
* `match` - (Required) PBR Rule match conditions. These determine what traffic will match the PBR Rule. match blocks are documented below.
* `action` - (Optional) PBR Rule actions. These specify the action to take if traffic matches the PBR Rule. action blocks are documented below.
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`match` supports the following:

* `interface` - (Optional) Match traffic on inbound interface 
* `port` - (Optional) Match traffic by service port 
* `protocol` - (Optional) Match traffic by protocol 
* `destination` - (Optional) Match traffic with destination network destination blocks are documented below.
* `source` - (Optional) Match traffic with source network source blocks are documented below.


`action` supports the following:

* `table` - (Optional) Name of PBR Table used to route matched traffic 
* `main_table` - (Optional) Use the main routing table to route matched traffic 
* `prohibit` - (Optional) Mark matched traffic as prohibited 
* `unreachable` - (Optional) Report matched traffic as having an unreachable destination 


`destination` supports the following:

* `address` - (Optional) IPv4 address of network 
* `mask_length` - (Optional) Mask length of network 


`source` supports the following:

* `address` - (Optional) IPv4 address of network 
* `mask_length` - (Optional) Mask length of network 
