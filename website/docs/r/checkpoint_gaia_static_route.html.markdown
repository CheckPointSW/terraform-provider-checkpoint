---
layout: "checkpoint"
page_title: "checkpoint_gaia_static_route"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-static-route"
description: |-
This resource allows you to execute Check Point Static Route.
---

# checkpoint_gaia_static_route

This resource allows you to execute Check Point Static Route.

## Example Usage


```hcl
resource "checkpoint_gaia_static_route" "example" {
  address     = "1.2.3.0"
  mask_length = 24
  type        = "gateway"
  rank        = 25
  comment     = "hello"

  next_hop {
    gateway  = "172.23.22.1"
    priority = "default"
  }
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required)  
* `mask_length` - (Required)  
* `type` - (Required) Type of next hop. Possible values: blackhole, gateway, reject 
* `next_hop` - (Optional) Static next-hop. Contains a list of next-hop gateways. Each gateway is formatted in the following manner: {"gateway": IP address or logical name, "priority": default or integer 1-8} next_hop blocks are documented below.
* `ping` - (Optional) Configures ping monitoring of the given IPv4 static route. Possible values: true, false 
* `rank` - (Optional) Selects a route when there are many routes to a destination that use different routing protocols. The route with the lowest rank value is selected. Possible values: default or integer 0-255 
* `scope_local` - (Optional) Configure the local-interface scope option, When the this option is enabled, the route treated as directly connected to local machine. Possible values: true, false 
* `comment` - (Optional)  
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`next_hop` supports the following:

* `gateway` - (Optional) IP address or logical name for the static next-hop gateway 
* `priority` - (Optional) Priority defines which gateway to select as the next-hop. The lower the priority, the higher the preference. Possible values: default or integer 1-8 
