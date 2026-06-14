---
layout: "checkpoint"
page_title: "checkpoint_gaia_allowed_clients"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-allowed-clients"
description: |-
This resource allows you to execute Check Point Allowed Clients.
---

# checkpoint_gaia_allowed_clients

This resource allows you to execute Check Point Allowed Clients.

## Example Usage


```hcl
resource "checkpoint_gaia_allowed_clients" "example" {
  allowed_any_host = true
}
```

## Argument Reference

The following arguments are supported:

* `allowed_networks` - (Optional)  allowed_networks blocks are documented below.
* `allowed_hosts` - (Optional)  allowed_hosts blocks are documented below.
* `allowed_any_host` - (Optional)  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`allowed_networks` supports the following:

* `subnet` - (Optional) The network subnet 
* `mask_length` - (Optional) The network mask length 


`allowed_hosts` supports the following:

