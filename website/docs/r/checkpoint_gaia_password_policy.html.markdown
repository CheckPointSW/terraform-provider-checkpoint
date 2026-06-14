---
layout: "checkpoint"
page_title: "checkpoint_gaia_password_policy"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-password-policy"
description: |-
This resource allows you to execute Check Point Password Policy.
---

# checkpoint_gaia_password_policy

This resource allows you to execute Check Point Password Policy.

## Example Usage


```hcl
resource "checkpoint_gaia_password_policy" "policy" {
  lock_settings {
    failed_attempts_settings {
      failed_attempts_allowed       = 10
      failed_lock_duration_seconds  = 1200
      failed_lock_enabled           = false
      failed_lock_enforced_on_admin = false
    }
    inactivity_settings {
      inactivity_threshold_days    = 365
      lock_unused_accounts_enabled = false
    }
    must_one_time_password_enabled               = false
    password_expiration_days                     = "never"
    password_expiration_maximum_days_before_lock = "never"
    password_expiration_warning_days             = 7
  }
  password_history {
    check_history_enabled   = true
    repeated_history_length = 10
  }
  password_strength {
    complexity               = 2
    minimum_length           = 6
    palindrome_check_enabled = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `lock_settings` - (Optional) password change configuration lock_settings blocks are documented below.
* `password_history` - (Optional) password history configuration password_history blocks are documented below.
* `password_strength` - (Optional) password strength configuration password_strength blocks are documented below.
* `all_users_require_two_factor_authentication` - (Optional) Force Two-Factor Authentication for all users. Upon their next login, if Two-Factor Authentication is not already set up, the users will be required to generate the authentication keys. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`lock_settings` supports the following:

* `inactivity_settings` - (Optional) inactivity configuration inactivity_settings blocks are documented below.
* `failed_attempts_settings` - (Optional) failed attempts configuration failed_attempts_settings blocks are documented below.
* `password_expiration_days` - (Optional) Password expiration lifetime, default value is 'never' 
* `password_expiration_warning_days` - (Optional) Number of days before a password expires that the user gets warned, default value is 7 days 
* `password_expiration_maximum_days_before_lock` - (Optional) Password expiration lockout in days, default value is 'never' 
* `must_one_time_password_enabled` - (Optional) Forces a user to change their password after                it has been set via "User Management" (but not via "Self Password Change" or forced change at login).Use this command to set the value. Default value is false 


`password_history` supports the following:

* `check_history_enabled` - (Optional) Password history check, default value is false 
* `repeated_history_length` - (Optional) Password history length, default value is 10 entries 


`password_strength` supports the following:

* `minimum_length` - (Optional) default length is 6 
* `complexity` - (Optional) default value is 2 
* `palindrome_check_enabled` - (Optional) Password palindrome check, default value is true 


`inactivity_settings` supports the following:

* `lock_unused_accounts_enabled` - (Optional) Password lock unused accounts, default: false 
* `inactivity_threshold_days` - (Optional) Inactivity days to password expiration lockout, default value is 365 days 


`failed_attempts_settings` supports the following:

* `failed_lock_duration_seconds` - (Optional) Password failed logging lockout duration, default value is 1200 
* `failed_lock_enforced_on_admin` - (Optional) Enforce failed lockout on admin user, default value is false 
* `failed_lock_enabled` - (Optional) Lock user after exceeded maximum allowed login attempts, default value is false 
* `failed_attempts_allowed` - (Optional) Amount of login attempts allowed before lockout, default value is 10 attempts 
