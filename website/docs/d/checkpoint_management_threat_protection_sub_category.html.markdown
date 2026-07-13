---
layout: "checkpoint"
page_title: "checkpoint_management_threat_protection_sub_category"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-protection-sub-category"
description: |-
Use this data source to get information on an existing Check Point Threat Protection Sub Category.
---

# Data Source: checkpoint_management_threat_protection_sub_category

Use this data source to get information on an existing Check Point Threat Protection Sub Category.

## Example Usage

```hcl
data "checkpoint_management_threat_protection_sub_category" "data_test" {
    name = "threat-protection-sub-category1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The sub-category's name.
* `category_id` - (Optional) The sub-category's unique identifier.
* `show_profiles` - (Optional) Indicates whether to calculate and show "profiles" field in reply.
* `category` - (Computed) Parent category reference. category blocks are documented below.
* `blade` - (Computed) The blade this protection belongs to.
* `engine` - (Computed) The engine that handles this protection.
* `known_today` - (Computed) The current number of protection items available in the latest update.
* `last_update` - (Computed) The date in which the protection was updated by Check Point. last_update blocks are documented below.
* `confidence_level` - (Computed) Confidence level of the protection.
* `performance_impact` - (Computed) Performance impact of the protection.
* `description` - (Computed) Detailed description.
* `profiles` - (Computed) Protection settings per profile.


`category` supports the following:

* `id` - (Computed) Parent category unique identifier.
* `name` - (Computed) Parent category name.


`last_update` supports the following:

* `iso_8601` - (Computed) Date and time represented in international ISO 8601 format.
* `posix` - (Computed) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
