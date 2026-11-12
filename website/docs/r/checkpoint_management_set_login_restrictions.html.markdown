---
layout: "checkpoint"
page_title: "checkpoint_management_set_login_restrictions"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-login-restrictions"
description: |-
 This resource allows you to execute Check Point Set Login Restrictions.
---

# checkpoint_management_set_login_restrictions

This resource allows you to execute Check Point Set Login Restrictions.

## Example Usage


```hcl
resource "checkpoint_management_set_login_restrictions" "example" {
  lockout_admin_account = true
  failed_authentication_attempts = 10
  unlock_admin_account = false
  lockout_duration = 30
  display_access_denied_message = false
}
```

## Argument Reference

The following arguments are supported:

* `lockout_admin_account` - (Optional) Indicates whether to lockout administrator's account after specified number of failed authentication attempts. 
* `failed_authentication_attempts` - (Optional) Number of failed authentication attempts before lockout administrator account. <font color="red">Required only when</font> lockout-admin-account is set to true. 
* `unlock_admin_account` - (Optional) Indicates whether to unlock administrator account after specified number of minutes. <font color="red">Required only when</font> lockout-admin-account is set to true. 
* `lockout_duration` - (Optional) Number of minutes of administrator account lockout. <font color="red">Required only when</font> lockout-admin-account is set to true. 
* `display_access_denied_message` - (Optional) Indicates whether to display informative message upon denying access. <font color="red">Required only when</font> lockout-admin-account is set to true. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

