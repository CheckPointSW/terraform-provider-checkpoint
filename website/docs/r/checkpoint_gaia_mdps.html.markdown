---
layout: "checkpoint"
page_title: "checkpoint_gaia_mdps"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-mdps"
description: |-
This resource allows you to execute Check Point Mdps.
---

# checkpoint_gaia_mdps

This resource allows you to execute Check Point Mdps.

## Example Usage


```hcl
resource "checkpoint_gaia_mdps" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `separation_interfaces` - (Optional) Management and/or Sync interface for the separation options separation_interfaces blocks are documented below.
* `routing_separation` - (Optional) Routing Separation routing_separation blocks are documented below.
* `resource_separation` - (Optional) Resource Separation resource_separation blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`separation_interfaces` supports the following:

* `management` - (Optional) Management interface 
* `sync` - (Optional) Sync interface (used in ClusterXL) 


`routing_separation` supports the following:

* `enabled` - (Optional) Routing separation state 


`resource_separation` supports the following:

* `enabled` - (Optional) Resource separation state 
* `allocated_cpus` - (Optional) Number of CPU's for resource separation 
