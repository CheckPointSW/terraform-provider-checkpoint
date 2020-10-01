---
layout: "checkpoint"
page_title: "checkpoint_management_user_template"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-user-template"
description: |-
This resource allows you to execute Check Point User Template.
---

# Resource: checkpoint_management_user_template

This resource allows you to execute Check Point User Template.

## Example Usage


```hcl
resource "checkpoint_management_user_template" "example" {
  name = "myusertemplate"
  expiration_date = "2030-05-30"
  expiration_by_global_properties = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `expiration_by_global_properties` - (Optional) Expiration date according to global properties. 
* `expiration_date` - (Optional) Expiration date in format: yyyy-MM-dd. 
* `authentication_method` - (Optional) Authentication method. 
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


`allowed_locations` supports the following:

* `destinations` - (Optional) Collection of allowed destination locations name or uid.destinations blocks are documented below.
* `sources` - (Optional) Collection of allowed source locations name or uid.sources blocks are documented below.


`encryption` supports the following:

* `enable_ike` - (Optional) Enable IKE encryption for users. 
* `enable_public_key` - (Optional) Enable IKE public key. 
* `enable_shared_secret` - (Optional) Enable IKE shared secret.