---
layout: "checkpoint"
page_title: "checkpoint_gaia_igmp_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-igmp-interface"
description: |-
This resource allows you to execute Check Point Igmp Interface.
---

# checkpoint_gaia_igmp_interface

This resource allows you to execute Check Point Igmp Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_igmp_interface" "example" {
  name            = "eth0"
  loss_robustness = "7"
  query_interval  = "125"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the IGMP interface 
* `last_member_query_interval` - (Optional) The number of seconds between queries that an IGMP router sends after it receives a "Leave Group" message from a host 
* `loss_robustness` - (Optional) The loss-robustness value 
* `query_interval` - (Optional) The number of seconds between IGMP general queries 
* `query_response_interval` - (Optional) The maximum delay time in seconds for hosts to respond to an IGMP membership query 
* `router_alert` - (Optional) Configure the router-alert option for this IGMP interface 
* `igmp_version` - (Optional) The IGMP version running 
* `reset` - (Optional) Reset all attributes of this interface to default values 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
