---
layout: "checkpoint"
page_title: "checkpoint_gaia_hostname_on_login_page"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-hostname-on-login-page"
description: |-
This resource allows you to execute Check Point Hostname On Login Page.
---

# checkpoint_gaia_hostname_on_login_page

This resource allows you to execute Check Point Hostname On Login Page.

## Example Usage


```hcl
resource "checkpoint_gaia_hostname_on_login_page" "example" {
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Hostname on Gaia Portal login page enabled (true/false) 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
