---
layout: "checkpoint"
page_title: "checkpoint_management_passcode_profile"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-passcode-profile"
description: |-
Use this data source to get information on an existing Check Point Passcode Profile.
---

# checkpoint_management_passcode_profile

Use this data source to get information on an existing Check Point Passcode Profile.

## Example Usage


```hcl
resource "checkpoint_management_passcode_profile" "example" {
  name = "New App Passcode Policy"
  allow_simple_passcode =false
  min_passcode_length = 10
  require_alphanumeric_passcode = true
  min_passcode_complex_characters = 3
  force_passcode_expiration = false
  passcode_expiration_period = 190
  enable_inactivity_time_lock = true
  max_inactivity_time_lock =  10
  enable_passcode_failed_attempts = false
}

data "checkpoint_management_passcode_profile" "data" {
  uid = "${checkpoint_management_passcode_profile.test.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `allow_simple_passcode` - The passcode length is 4 and only numeric values allowed. 
* `min_passcode_length` - Minimum passcode length - relevant if "allow-simple-passcode" is disable. 
* `require_alphanumeric_passcode` - Require alphanumeric characters in the passcode - relevant if "allow-simple-passcode" is disable. 
* `min_passcode_complex_characters` -Minimum number of complex characters (if "require-alphanumeric-passcode" is enabled). The number of the complex characters cannot be greater than number of the passcode length. 
* `force_passcode_expiration` - Enable/disable expiration date to the passcode. 
* `passcode_expiration_period` - The period in days after which the passcode will expire. 
* `enable_inactivity_time_lock` - Lock the device if app is inactive. 
* `max_inactivity_time_lock` - Time without user input before passcode must be re-entered (in minutes). 
* `enable_passcode_failed_attempts` - Exit after few failures in passcode verification. 
* `max_passcode_failed_attempts` - Number of failed attempts allowed. 
* `enable_passcode_history` - Check passcode history for reparations. 
* `passcode_history` -  Number of passcodes that will be kept in history. 
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
