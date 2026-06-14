---
layout: "checkpoint"
page_title: "checkpoint_gaia_user"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-user"
description: |-
This resource allows you to execute Check Point User.
---

# checkpoint_gaia_user

This resource allows you to execute Check Point User.

## Example Usage


```hcl
resource "checkpoint_gaia_user" "example" {
  name     = "myusername"
  uid      = 5555
  password = "Mypass123!"
  shell    = "no-login"
  allow_access_using       = ["CLI"]
  primary_system_group_id  = 60
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required)  
* `uid` - (Required) Specifies a numeric user ID used to identify permissions of a user, duplicate UIDs are not allowed 
* `homedir` - (Optional) Specifies the user's home directory as the full UNIX path name where the user is placed on login. If the directory doesn't exist, it is created. Range: Must be under '/home' and must not contain colon (:).  
* `primary_system_group_id` - (Optional) GID. Numeric ID which is used in identifying the primary group to which this user belongs.  
* `secondary_system_groups` - (Optional) This operation assigns groups to the user. Valid values: must be names of known groups.  secondary_system_groups blocks are documented below.
* `password` - (Optional) Specifies new password on command line. Check Point recommends that a password be at least eight characters long. A password must contain at least six characters. Enforcement level can be modified via 'password control' feature.  
* `password_hash` - (Optional) An encrypted representation of the password.  
* `real_name` - (Optional) Specifies a string describing a user; conventionally it's the user's full name.  
* `shell` - (Optional) Specifies the user's default command-line interpreter during login.  
* `allow_access_using` - (Optional) Modify the access-mechanisms available for a user. Valid values: CLI, Web-UI, Gaia-API (supported from R81.10).  allow_access_using blocks are documented below.
* `must_change_password` - (Optional) Forcing password change is relevant only when a password is set. When set to 'True': Force the user to change their password the next time they log in. If they don't log in within the time limit configured in 'set password-controls expiration-lockout-days' they may not be able to log in at all. When set to 'False': If the user was being forced to change their password, cancel that. If the user was locked out due to failure to change their password within the time limit configured in 'set password-controls expiration-lockout-days' they will no longer be locked out.  
* `roles` - (Optional)  roles blocks are documented below.
* `requires_two_factor_authentication` - (Optional) Force Two-Factor Authentication for this user. Upon their next login, if Two-Factor Authentication is not already set up, the user will be required to generate the authentication keys. 
* `unlock` - (Optional) If the user has been locked out, cancel that. True: cancel lock-out. False: do nothing.  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `locked` - (Computed) Computed field, returned in the response. 
