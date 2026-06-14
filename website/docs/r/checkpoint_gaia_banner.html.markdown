---
layout: "checkpoint"
page_title: "checkpoint_gaia_banner"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-banner"
description: |-
This resource allows you to execute Check Point Banner.
---

# checkpoint_gaia_banner

This resource allows you to execute Check Point Banner.

## Example Usage


```hcl
resource "checkpoint_gaia_banner" "example" {
  enabled = true
  message = "This is a new banner message"
}
```

## Argument Reference

The following arguments are supported:

* `message` - (Optional) Banner message for the web, ssh and serial login. Empty string returns to default 
* `enabled` - (Optional) Banner message enabled (true/false) 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
