---
layout: "checkpoint"
page_title: "checkpoint_management_login "
sidebar_current: "docs-checkpoint-resource-checkpoint-management-login"
description: |-
  Log in to the server with username and password.
---

# checkpoint_management_login

Log in to the server with username and password.

## Example Usage

```hcl
resource "checkpoint_management_login" "example" {
  user = "aa"
  password = "aaaa"
}
```

## Argument Reference

The following arguments are supported:

* `user` - (Required) Session unique identifier. Specify it to publish a different session than the one you currently use.
* `password` - (Required) Administrator password.
* `continue_last_session` - (Optional) When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.
* `domain` - (Optional) Use domain to login to specific domain. Domain can be identified by name or UID.
* `enter_last_published_session` - (Optional) Login to the last published session. Such login is done with the Read Only permissions.
* `read_only` - (Optional) Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.
* `session_comments` - (Optional) Session comments.
* `session_description` - (Optional) Session description.
* `session_name` - (Optional) Session unique name.
* `session_timeout` - (Optional) Session expiration timeout in seconds. Default 600 seconds.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  



