---
layout: "checkpoint"
page_title: "checkpoint_gaia_expert_password"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-expert-password"
description: |-
This resource allows you to execute Check Point Expert Password.
---

# checkpoint_gaia_expert_password

This resource allows you to execute Check Point Expert Password.

## Example Usage


```hcl
resource "checkpoint_gaia_expert_password" "example" {
  password = "Admin1234!"
}
```

## Argument Reference

The following arguments are supported:

* `password` - (Optional) expert new password 
* `password_hash` - (Optional) An encrypted representation of the password 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
