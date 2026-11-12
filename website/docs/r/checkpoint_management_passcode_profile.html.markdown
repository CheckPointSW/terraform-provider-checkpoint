---
layout: "checkpoint"
page_title: "checkpoint_management_passcode_profile"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-passcode-profile"
description: |-
This resource allows you to execute Check Point Passcode Profile.
---

# checkpoint_management_passcode_profile

This resource allows you to execute Check Point Passcode Profile.

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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `allow_simple_passcode` - (Optional) The passcode length is 4 and only numeric values allowed. 
* `min_passcode_length` - (Optional) Minimum passcode length - relevant if "allow-simple-passcode" is disable. 
* `require_alphanumeric_passcode` - (Optional) Require alphanumeric characters in the passcode - relevant if "allow-simple-passcode" is disable. 
* `min_passcode_complex_characters` - (Optional) Minimum number of complex characters (if "require-alphanumeric-passcode" is enabled). The number of the complex characters cannot be greater than number of the passcode length. 
* `force_passcode_expiration` - (Optional) Enable/disable expiration date to the passcode. 
* `passcode_expiration_period` - (Optional) The period in days after which the passcode will expire. 
* `enable_inactivity_time_lock` - (Optional) Lock the device if app is inactive. 
* `max_inactivity_time_lock` - (Optional) Time without user input before passcode must be re-entered (in minutes). 
* `enable_passcode_failed_attempts` - (Optional) Exit after few failures in passcode verification. 
* `max_passcode_failed_attempts` - (Optional) Number of failed attempts allowed. 
* `enable_passcode_history` - (Optional) Check passcode history for reparations. 
* `passcode_history` - (Optional) Number of passcodes that will be kept in history. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
