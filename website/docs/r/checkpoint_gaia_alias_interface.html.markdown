---
layout: "checkpoint"
page_title: "checkpoint_gaia_alias_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-alias-interface"
description: |-
This resource allows you to execute Check Point Alias Interface.
---

# checkpoint_gaia_alias_interface

This resource allows you to execute Check Point Alias Interface.

## Example Usage


```hcl
resource "checkpoint_gaia_alias_interface" "example" {
  parent           = "eth0"
  ipv4_address     = "192.168.1.10"
  ipv4_mask_length = 24
}
```

## Argument Reference

The following arguments are supported:

* `parent` - (Required)  
* `ipv4_address` - (Required)  
* `ipv4_mask_length` - (Required)  
* `name` - (Computed)  
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
