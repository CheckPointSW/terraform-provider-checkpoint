---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_configuration_bgp_confederation_peer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-configuration-bgp-confederation-peer"
description: |-
This resource allows you to execute Check Point Show Configuration Bgp Confederation Peer.
---

# checkpoint_gaia_show_configuration_bgp_confederation_peer

This resource allows you to execute Check Point Show Configuration Bgp Confederation Peer.

## Example Usage


```hcl
data "checkpoint_gaia_show_configuration_bgp_confederation_peer" "example" {
  member_as = "1234.4321"
  peer = "fe80::34"
}
```

## Argument Reference

The following arguments are supported:

* `peer` - (Required) IP address of the peer. 
* `member_as` - (Required) Specify the Routing Domain identifier of the Confederation peer.  If the peer group specified is the local Routing Domain, it will run IBGP in a full mesh (just as an internal peer group normally would in non-Confederation mode). Otherwise, if an external Routing Domain within the Confederation is specified, the peer group will run a modified version of eBGP, which preserves route metrics and other BGP attributes.  The value can be one of the following: An integer from 1-4294967295 A float from 0.1-65535.65535 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

