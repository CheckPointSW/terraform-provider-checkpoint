---
layout: "checkpoint"
page_title: "checkpoint_management_user_template"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-user-template"
description: |-
This resource allows you to execute Check Point User Template.
---

# Data Source: checkpoint_management_user_template

This resource allows you to execute Check Point User Template.

## Example Usage


```hcl
resource "checkpoint_management_user_template" "user_template" {
    name = "my template"
    expiration_date = "2030-12-31"
    expiration_by_global_properties = false
}

data "checkpoint_management_user_template" "test_user_template" {
    name = "${checkpoint_management_user_template.user_template.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `expiration_by_global_properties` - Expiration date according to global properties. 
* `expiration_date` - Expiration date in format: yyyy-MM-dd. 
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