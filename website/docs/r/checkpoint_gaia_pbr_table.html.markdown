---
layout: "checkpoint"
page_title: "checkpoint_gaia_pbr_table"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-pbr-table"
description: |-
This resource allows you to execute Check Point Pbr Table.
---

# checkpoint_gaia_pbr_table

This resource allows you to execute Check Point Pbr Table.

## Example Usage


```hcl
resource "checkpoint_gaia_pbr_table" "example" {
  table = "mytable"
  static_routes {
    address     = "10.0.0.0"
    mask_length = 8
    type        = "blackhole"
  }
}
```

## Argument Reference

The following arguments are supported:

* `table` - (Required) Name of PBR Table 
* `static_routes` - (Required) List of static routes configured on PBR Table static_routes blocks are documented below.
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `static_route_limit` - (Computed) The maximum number of configured static-routes to show in response 
* `static_route_offset` - (Computed) The number of configured static-routes to initially skip 
* `static_route_order` - (Computed) Sorts the static-routes by address in either ascending or descending order. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`static_routes` supports the following:

* `address` - (Optional) IPv4 address of route 
* `mask_length` - (Optional) Mask length of route 
* `type` - (Optional) Type of next-hop. Possible values: blackhole, gateway, reject 
* `next_hop` - (Optional) Static next-hop. Contains a list of next-hop gateways. Each gateway is formatted in the following manner:{"gateway": IP address or logical name, "priority": default or integer 1-8} next_hop blocks are documented below.
* `ping` - (Optional) Configures ping monitoring of the given IPv4 static route. Possible values: true, false 


`next_hop` supports the following:

* `gateway` - (Optional) IP address or logical name for the static next-hop gateway 
* `priority` - (Optional) Priority defines which gateway to select as the next-hop. The lower the priority, the higher the preference. Possible values: default or integer 1-8 
