---
layout: "checkpoint"
page_title: "checkpoint_gaia_router_id"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-router-id"
description: |-
This resource allows you to execute Check Point Router Id.
---

# checkpoint_gaia_router_id

This resource allows you to execute Check Point Router Id.

## Example Usage


```hcl
resource "checkpoint_gaia_router_id" "example" {
  router_id = "1.2.3.4"
}
```

## Argument Reference

The following arguments are supported:

* `router_id` - (Optional) Configures the Router ID used by BGP and OSPF. It is usually the IPv4 address of one of the local interfaces, and should uniquely identify the router within the local Autonomous System. It is generally recommended that a non-127.0.0.1 loopback address be used. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
