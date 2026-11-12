---
layout: "checkpoint"
page_title: "checkpoint_management_set_login_restrictions"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-login-restrictions"
description: |-
  Use this data source to get information on an existing Check Point Login Restrictions.
---

# Data Source: checkpoint_management_login_restrictions

Use this data source to get information on an existing Check Point Login Restrictions.

## Example Usage


```hcl
data "checkpoint_management_login_restrictions" "data_test" {
}
```

## Argument Reference

The following arguments are supported:
* `uid` - Object unique identifier.
* `lockout_admin_account` -  Indicates whether to lockout administrator's account after specified number of failed authentication attempts. 
* `failed_authentication_attempts` -  Number of failed authentication attempts before lockout administrator account.
* `unlock_admin_account` -  Indicates whether to unlock administrator account after specified number of minutes.
* `lockout_duration` -  Number of minutes of administrator account lockout.
* `display_access_denied_message` -  Indicates whether to display informative message upon denying access.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

