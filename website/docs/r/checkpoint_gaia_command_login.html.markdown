---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_login"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-login"
description: |-
This resource allows you to execute Check Point Login.
---

# checkpoint_gaia_command_login

This resource allows you to execute Check Point Login.

## Example Usage


```hcl
resource "checkpoint_gaia_command_login" "example" {
  user     = "admin"
  password = "zubur1"
}
```

## Argument Reference

The following arguments are supported:

* `user` - (Required) Administrator user name 
* `password` - (Required) Administrator password 
* `session_timeout` - (Optional) Session expiration timeout in seconds 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `verification_code` - (Optional) Verification code, if Two-Factor Authentication is enabled for this user. This field must be a string comprised solely of digits. 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

