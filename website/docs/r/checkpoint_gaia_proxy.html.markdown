---
layout: "checkpoint"
page_title: "checkpoint_gaia_proxy"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-proxy"
description: |-
This resource allows you to execute Check Point Proxy.
---

# checkpoint_gaia_proxy

This resource allows you to execute Check Point Proxy.

## Example Usage


```hcl
resource "checkpoint_gaia_proxy" "example" {
  address = "1.1.1.1"
  port = 89
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Optional)  
* `port` - (Optional)  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
