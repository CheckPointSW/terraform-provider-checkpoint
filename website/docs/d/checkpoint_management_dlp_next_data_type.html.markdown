---
layout: "checkpoint"
page_title: "checkpoint_management_dlp_next_data_type"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-dlp-next-data-type"
description: |-
Use this data source to get information on an existing Check Point Dlp Next Data Type.
---

# Data Source: checkpoint_management_dlp_next_data_type

Use this data source to get information on an existing Check Point Dlp Next Data Type.

## Example Usage

```hcl
data "checkpoint_management_dlp_next_data_type" "data_test" {
    name = "dlp-next-data-type1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `description` - (Computed) DLP Next Data Type description in Infinity Portal.
* `external_id` - (Computed) DLP Next Data Type unique identifier in Infinity Portal.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `comments` - (Computed) Comments string.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
