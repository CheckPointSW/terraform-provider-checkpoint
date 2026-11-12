---
layout: "checkpoint"
page_title: "checkpoint_management_administrator"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-administrator"
description: |-
Use this data source to get information on an existing Check Point Administrator.
---

# Data Source: checkpoint_management_administrator

Use this data source to get information on an existing Check Point Administrator.

## Example Usage


```hcl
resource "checkpoint_management_administrator" "admin" {
  name = "example"
  permissions_profile {
    domain = "domain1"
    profile = "Read Only All"
  }

  multi_domain_profile = "domain level only"
  password = "1233"

}

data "checkpoint_management_administrator" "data_admin" {
    name = "${checkpoint_management_administrator.admin.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `authentication_method` - Authentication method.
* `email` - Administrator email.
* `expiration_date`
* `multi_domain_profile` - Administrator multi-domain profile. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `must_change_password` - True if administrator must change password on the next login.
* `permissions_profile` - Administrator permissions profile. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. permissions_profile blocks are documented below.
* `phone_number` - Administrator phone number.
* `radius_server` - RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `sic_name` - Name of the Secure Internal Connection Trust.
* `tacacs_server` - TACACS server object identified by the name or UID . Must be set when "authentication-method" was selected to be "TACACS". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`permissions_profile` supports the following:

* `domain` - The domain's profile.
* `profile` - Permission profile.