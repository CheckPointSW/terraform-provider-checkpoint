---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_configuration_bgp_external_peers"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-configuration-bgp-external-peers"
description: |-
This resource allows you to execute Check Point Show Configuration Bgp External Peers.
---

# checkpoint_gaia_show_configuration_bgp_external_peers

This resource allows you to execute Check Point Show Configuration Bgp External Peers.

## Example Usage


```hcl
data "checkpoint_gaia_show_configuration_bgp_external_peers" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filter the results 
* `limit` - (Optional) The maximum number of returned results 
* `offset` - (Optional) The number of results to initially skip 
* `order` - (Optional) Sorts the peers first by their AS, then by their IDs in either ascending or descending order. 
* `remote_as` - (Optional) The Autonomous System number of the peerThe value can be one of the following: 'all' An integer from 1-4294967295 A float from 0.1-65535.65535 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

