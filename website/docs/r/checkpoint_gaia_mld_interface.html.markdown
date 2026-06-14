---
layout: "checkpoint"
page_title: "checkpoint_gaia_mld_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-mld-interface"
description: |-
This resource allows you to execute Check Point Mld Interface.
---

# checkpoint_gaia_mld_interface

This resource allows you to execute Check Point Mld Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_mld_interface" "example" {
  name            = "eth0"
  loss_robustness = "7"
  query_interval  = "125"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the MLD interface 
* `last_listener_query_count` - (Optional) The number of queries to send when a listener leaves a group 
* `last_listener_query_interval` - (Optional) The number of seconds between queries when a listener leaves a group 
* `loss_robustness` - (Optional) The loss-robustness value 
* `query_interval` - (Optional) The number of seconds between MLD general queries 
* `query_response_interval` - (Optional) The maximum delay time in seconds for hosts to respond to an MLD membership query 
* `startup_query_count` - (Optional) The number of queries sent when MLD starts up 
* `startup_query_interval` - (Optional) The number of seconds between MLD startup queries 
* `mld_version` - (Optional) The MLD version running 
* `reset` - (Optional) Reset all attributes of this interface to default values 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
