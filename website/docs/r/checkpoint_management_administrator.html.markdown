---
layout: "checkpoint"
page_title: "checkpoint_management_administrator"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-administrator"
description: |-
This resource allows you to add/update/delete Check Point Administrator.
---

# Resource: checkpoint_management_host

This resource allows you to add/update/delete Check Point Administrator.

## Example Usage: MDS


```hcl
resource "checkpoint_management_administrator" "admin" {
  name = "example"
  permissions_profile {
    domain = "domain1"
    profile = "Read Only All"
  }

  multi_domain_profile = "Domain Level Only"
  password = "1233"

}

```

## Example Usage: SMC


```hcl
resource "checkpoint_management_administrator" "admin" {
  name = "example"
  permissions_profile {
    domain = "SMC User"
    profile = "Read Only All"
  }
  password = "1233"

}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `password` - (Required) Administrator password.
* `authentication_method` - (Required) Authentication method.
* `permission_profile` - (Required) Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to "Multi-Domain Super User" or "Domain Super User". In SMC, permissions_profile with single object, domain must be "SMC User".
* `email` - (Optional) Administrator email.
* `expiration_date` - (Optional) Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.
* `multi_domain_profile` - (Optional) Administrator multi-domain profile. Only in MDS.
* `must_change_password` - (Optional) True if administrator must change password on the next login.
* `password_hash` (Optional) Administrator password hash.
* `phone_number` - (Optional) Administrator phone number.
* `radius_server` - (Optional) RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".
* `tacacs_server` - (Optional) TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `sic_name` - Name of the Secure Internal Connection Trust.