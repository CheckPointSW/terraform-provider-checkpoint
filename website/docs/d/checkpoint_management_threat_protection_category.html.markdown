---
layout: "checkpoint"
page_title: "checkpoint_management_threat_protection_category"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-protection-category"
description: |-
Use this data source to get information on an existing Check Point Threat Protection Category.
---

# Data Source: checkpoint_management_threat_protection_category

Use this data source to get information on an existing Check Point Threat Protection Category.

## Example Usage

```hcl
data "checkpoint_management_threat_protection_category" "data_test" {
    name = "threat-protection-category1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The Category name.
* `category_id` - (Optional) The Category unique identifier.
* `blade` - (Optional) The blade this category belongs to. Required when using 'name'.
* `show_profiles` - (Optional) Indicates whether to calculate and show "profiles" field in reply.
* `engine` - (Computed) The engine that handles this category.
* `known_today` - (Computed) The current number of protection items available in the latest update.
* `last_update` - (Computed) The date in which the protection was updated by Check Point. last_update blocks are documented below.
* `confidence_level` - (Computed) Confidence levels. confidence_level blocks are documented below.
* `performance_impact` - (Computed) Performance impacts. performance_impact blocks are documented below.
* `description` - (Computed) Description of the category.
* `profiles` - (Computed) Protection settings per profile.


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
