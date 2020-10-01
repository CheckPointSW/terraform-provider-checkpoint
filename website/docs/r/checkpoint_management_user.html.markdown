---
layout: "checkpoint"
page_title: "checkpoint_management_user"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-user"
description: |-
This resource allows you to execute Check Point User.
---

# Resource: checkpoint_management_user

This resource allows you to execute Check Point User.

## Example Usage


```hcl
resource "checkpoint_management_user" "example" {
  name = "myuser"
  email = "myuser@email.com"
  expiration_date = "2030-05-30"
  phone_number = "0501112233"
  authentication_method = "securid"
  connect_daily = true
  from_hour = "08:00"
  to_hour = "17:00"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `email` - (Optional) User email. 
* `expiration_date` - (Optional) Expiration date in format: yyyy-MM-dd. 
* `phone_number` - (Optional) User phone number. 
* `authentication_method` - (Optional) Authentication method. 
* `password` - (Optional) Checkpoint password authentication method identified by the name or UID. Must be set when "authentication-method" was selected to be "Check Point Password". 
* `radius_server` - (Optional) RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS". 
* `tacacs_server` - (Optional) TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS". 
* `connect_on_days` - (Optional) Days users allow to connect.
* `connect_daily` - (Optional) Connect every day.
* `from_hour` - (Optional) Allow users connect from hour.
* `to_hour` - (Optional) Allow users connect until hour.
* `allowed_locations` - (Optional) User allowed locations. allowed_locations blocks are documented below.
* `encryption` - (Optional) User encryption. encryption blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `template` - (Optional) User template name or UID. 


`allowed_locations` supports the following:

* `destinations` - (Optional) Collection of allowed destination locations name or uid.
* `sources` - (Optional) Collection of allowed source locations name or uid.

`encryption` supports the following:

* `enable_ike` - (Optional) Enable IKE encryption for users. 
* `enable_public_key` - (Optional) Enable IKE public key. 
* `enable_shared_secret` - (Optional) Enable IKE shared secret. 
* `shared_secret` - (Optional) IKE shared secret.