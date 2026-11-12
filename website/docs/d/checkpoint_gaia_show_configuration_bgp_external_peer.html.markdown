---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_configuration_bgp_external_peer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-configuration-bgp-external-peer"
description: |-
This resource allows you to execute Check Point Show Configuration Bgp External Peer.
---

# checkpoint_gaia_show_configuration_bgp_external_peer

This resource allows you to execute Check Point Show Configuration Bgp External Peer.

## Example Usage


```hcl
data "checkpoint_gaia_show_configuration_bgp_external_peer" "example" {
  peer      = "192.168.1.254"
  remote_as = "65002"
}
```

## Argument Reference

The following arguments are supported:

* `peer` - (Required) IP address of the peer. 
* `remote_as` - (Required) The Autonomous System number of the peerThe value can be one of the following: An integer from 1-4294967295 A float from 0.1-65535.65535 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

