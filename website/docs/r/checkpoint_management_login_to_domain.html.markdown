---
layout: "checkpoint"
page_title: "checkpoint_management_login_to_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-login-to-domain"
description: |-
This resource allows you to execute Check Point Login To Domain.
---

# checkpoint_management_login_to_domain

This resource allows you to execute Check Point Login To Domain.

## Example Usage


```hcl
resource "checkpoint_management_login_to_domain" "example" {
  domain = "Global"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) Domain identified by the name or UID. 
* `continue_last_session` - (Optional) When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.
* `read_only` - (Optional) Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.
* `sid` - Session unique identifier. Enter this session unique identifier in the 'X-chkp-sid' header of each request.
* `api_server_version` - API Server version.
* `disk_space_message` - Information about the available disk space on the management server.
* `last_login_was_at` - Timestamp when administrator last accessed the management server. last_login_was_at blocks are documented below.
* `login_message` - Login message. login_message blocks are documented below.
* `read_only` - True if this session is read only.
* `session_timeout` - Session expiration timeout in seconds.
* `standby` - True if this management server is in the standby mode.
* `uid` - Session object unique identifier. This identifier may be used in the discard API to discard changes that were made in this session, when administrator is working from another session, or in the 'switch-session' API.
* `url` - URL that was used to reach the API server.

`last_login_was_at` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`login_message` supports the following:

* `header` - Message header.
* `message` - Message content.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

