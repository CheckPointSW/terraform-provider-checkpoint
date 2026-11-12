---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_connections"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-connections"
description: |-
This resource allows you to execute Check Point Show Connections.
---

# checkpoint_gaia_show_connections

This resource allows you to execute Check Point Show Connections.

## Example Usage


```hcl
data "checkpoint_gaia_show_connections" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `max_results` - (Optional) Max number of connection to display, or "all" (10 by default) 
* `use_preset` - (Optional) A preset represents a pre-configured request with default values. But any value can be overridden by being explicitly provided in the request. 
* `preset` - (Optional) The names of the presets to be used (API call "show-connections-presets" returns the list of available presets) preset blocks are documented below.
* `filter` - (Optional) Connection filtering options filter blocks are documented below.
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`filter` supports the following:

* `instance_id` - (Optional) Query a specific instance / set of instances (all by default) instance_id blocks are documented below.
* `source` - (Optional) Return connections with the given source IP address 
* `destination` - (Optional) Return connections with the given destination IP address 
* `ip_protocol` - (Optional) Return connections with the given IP protocol number 
* `source_port` - (Optional) Return connections with the given source port 
* `destination_port` - (Optional) Return connections with the given destination port 
* `ip_version` - (Optional) Return connections of the given IP version (v4 / v6 / any (default)) 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

