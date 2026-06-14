---
layout: "checkpoint"
page_title: "checkpoint_gaia_grub_password"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-grub-password"
description: |-
This resource allows you to execute Check Point Grub Password.
---

# checkpoint_gaia_grub_password

This resource allows you to execute Check Point Grub Password.

## Example Usage


```hcl
resource "checkpoint_gaia_grub_password" "example" {
  password = "Admin1234!"
}
```

## Argument Reference

The following arguments are supported:

* `password` - (Optional) GRUB new password 
* `password_hash` - (Optional) An encrypted representation of the password 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
