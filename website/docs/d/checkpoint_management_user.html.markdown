---
layout: "checkpoint"
page_title: "checkpoint_management_user"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-user"
description: |-
This resource allows you to execute Check Point User.
---

# Data Source: checkpoint_management_user

This resource allows you to execute Check Point User.

## Example Usage


```hcl
resource "checkpoint_management_user" "user" {
    name = "my user"
    email = "email@email.com"
    expiration_date = "2030-12-31"
    phone_number = "12345678"
    authentication_method = "securid"
    connect_daily = true
    from_hour = "08:00"
    to_hour = "17:00"
}

data "checkpoint_management_user" "test_user" {
    name = "${checkpoint_management_user.user.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.   
* `email` - User email. 
* `expiration_date` - Expiration date in format: yyyy-MM-dd. 
* `phone_number` - User phone number. 
* `authentication_method` - Authentication method. 
* `radius_server` - RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS". 
* `tacacs_server` - TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS". 
* `connect_on_days` - Days users allow to connect.
* `connect_daily` - Connect every day.
* `from_hour` - Allow users connect from hour.
* `to_hour` - Allow users connect until hour.
* `allowed_locations` - User allowed locations. allowed_locations blocks are documented below.
* `encryption` - User encryption. encryption blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object.
* `comments` - Comments string. 


`allowed_locations` supports the following:

* `destinations` - Collection of allowed destination locations name or uid.
* `sources` - Collection of allowed source locations name or uid.

`encryption` supports the following:

* `enable_ike` - Enable IKE encryption for users.
* `enable_public_key` - Enable IKE public key.
* `enable_shared_secret` - Enable IKE shared secret.