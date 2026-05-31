---
layout: "checkpoint"
page_title: "checkpoint_gaia_fips"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-fips"
description: |-
This resource allows you to execute Check Point Fips.
---

# checkpoint_gaia_fips

This resource allows you to execute Check Point Fips.

## Example Usage


```hcl
resource "checkpoint_gaia_fips" "example" {
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Required) FIPS mode enabled status 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
