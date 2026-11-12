---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_bgp_peer"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-bgp-peer"
description: |-
This resource allows you to execute Check Point Show Bgp Peer.
---

# checkpoint_gaia_show_bgp_peer

This resource allows you to execute Check Point Show Bgp Peer.

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

# Step 2: configure BGP AS
resource "checkpoint_gaia_command_set_bgp" "bgp_setup" {
  as = "65001"

  depends_on = [checkpoint_gaia_command_set_bgp.clear_conf]
}

# Step 3: configure external peer group for AS 65002
resource "checkpoint_gaia_command_set_bgp_external" "ext_group" {
  remote_as = "65002"
  enabled   = true

  depends_on = [checkpoint_gaia_command_set_bgp.bgp_setup]
}

# Step 4: add the external peer
resource "checkpoint_gaia_bgp_external_peer" "peer_setup" {
  peer      = "192.168.1.254"
  remote_as = "65002"

  depends_on = [checkpoint_gaia_command_set_bgp_external.ext_group]
}

# Step 5: query the peer stats
data "checkpoint_gaia_show_bgp_peer" "example" {
  peer = "192.168.1.254"

  depends_on = [checkpoint_gaia_bgp_external_peer.peer_setup]
}
```

## Argument Reference

The following arguments are supported:

* `peer` - (Required) The peer to be queried 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

