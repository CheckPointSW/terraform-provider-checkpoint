---
layout: "checkpoint"
page_title: "checkpoint_management_set_threat_protection_sub_category"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-threat-protection-sub-category"
description: |-
This resource allows you to execute Check Point Set Threat Protection Sub Category.
---

# checkpoint_management_set_threat_protection_sub_category

This resource allows you to execute Check Point Set Threat Protection Sub Category.

## Example Usage


```hcl
resource "checkpoint_management_set_threat_protection_sub_category" "example" {
  name = "Backdoor.WIN32.FoggyWeb.B"
  action = "prevent"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The sub-category's name. 
* `category_id` - (Optional) The sub-category's unique identifier. 
* `all_profiles` - (Optional) Apply action to all profiles. Default: true. 
* `show_profiles` - (Optional) Indicates whether to calculate and show "profiles" field in reply. 
* `action` - (Optional) Action to apply to all profiles. Required when all-profiles is true. 
* `overrides` - (Optional) Overrides per profile for this protection. Required when all-profiles is false.overrides blocks are documented below.
* `blade` - (Computed) The blade this protection belongs to.
* `engine` - (Computed) The engine that handles this protection.
* `known_today` - (Computed) The current number of protection items available in the latest update.
* `confidence_level` - (Computed) Confidence level of the protection.
* `performance_impact` - (Computed) Performance impact of the protection.
* `description` - (Computed) Detailed description.
* `profiles` - (Computed) Protection settings per profile.
* `domain` - (Computed) Information about the domain that holds the Object.
* `category` - (Computed) Parent category reference. category blocks are documented below.
* `last_update` - (Computed) The date in which the protection was updated by Check Point. last_update blocks are documented below.


`overrides` supports the following:

* `action` - (Optional) Action to apply for the specified profile.
* `profile` - (Optional) Profile name or UID.


`category` supports the following:

* `name` - (Computed) Object name. Must be unique in the domain.
* `uid` - (Computed) Object unique identifier.


`last_update` supports the following:

* `iso_8601` - (Computed) Date and time represented in international ISO 8601 format.
* `posix` - (Computed) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

