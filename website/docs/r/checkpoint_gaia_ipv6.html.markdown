---
layout: "checkpoint"
page_title: "checkpoint_gaia_ipv6"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-ipv6"
description: |-
This resource allows you to execute Check Point Ipv6.
---

# checkpoint_gaia_ipv6

This resource allows you to execute Check Point Ipv6.

## Example Usage


```hcl
resource "checkpoint_gaia_ipv6" "example" {
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Required)  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
