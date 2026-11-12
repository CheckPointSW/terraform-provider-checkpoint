---
layout: "checkpoint"
page_title: "checkpoint_management_set_threat_protection_category"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-threat-protection-category"
description: |-
This resource allows you to execute Check Point Set Threat Protection Category.
---

# checkpoint_management_set_threat_protection_category

This resource allows you to execute Check Point Set Threat Protection Category.

## Example Usage


```hcl
resource "checkpoint_management_set_threat_protection_category" "example" {
  name = "Reputation IPs"
  all_profiles = true
  action = "detect"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The Category name. 
* `category_id` - (Optional) The Category unique identifier. 
* `blade` - (Optional) The blade this category belongs to. Required when using 'name'. 
* `show_profiles` - (Optional) Indicates whether to calculate and show "profiles" field in reply. 
* `all_profiles` - (Optional) Apply action to all profiles. Default: true. 
* `action` - (Optional) Action to apply to all profiles. Required when all-profiles is true. 
* `overrides` - (Optional) Overrides per profile for this protection. Required when all-profiles is false.overrides blocks are documented below.
* `engine` - (Computed) The engine that handles this category.
* `known_today` - (Computed) The current number of protection items available in the latest update.
* `description` - (Computed) Description of the category.
* `domain` - (Computed) Information about the domain that holds the Object.
* `last_update` - (Computed) The date in which the protection was updated by Check Point. last_update blocks are documented below.
* `confidence_level` - (Computed) Confidence levels. confidence_level blocks are documented below.
* `performance_impact` - (Computed) Performance impacts. performance_impact blocks are documented below.
* `profiles` - (Computed) Protection settings per profile. profiles blocks are documented below.


`overrides` supports the following:

* `action` - (Optional) Action to apply for the specified profile.
* `profile` - (Optional) Profile name or UID.


`last_update` supports the following:

* `iso_8601` - (Computed) Date and time represented in international ISO 8601 format.
* `posix` - (Computed) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`confidence_level` supports the following:

* `high` - (Computed) Count of protections classified with high level.
* `low` - (Computed) Count of protections classified with low level.
* `medium` - (Computed) Count of protections classified with medium level.


`performance_impact` supports the following:

* `high` - (Computed) Count of protections classified with high level.
* `low` - (Computed) Count of protections classified with low level.
* `medium` - (Computed) Count of protections classified with medium level.


`profiles` supports the following:

* `uid` - (Computed) Profile UID.
* `name` - (Computed) Profile name.
* `default_action` - (Computed) Default action applied for this profile. default_action blocks are documented below.
* `override_action` - (Computed) Override action applied for this profile. override_action blocks are documented below.


`default_action` supports the following:

* `name` - (Computed) Object name. Must be unique in the domain.
* `uid` - (Computed) Object unique identifier.


`override_action` supports the following:

* `name` - (Computed) Object name. Must be unique in the domain.
* `uid` - (Computed) Object unique identifier.