---
layout: "checkpoint"
page_title: "checkpoint_management_threat_emulation_image"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-emulation-image"
description: |-
Use this data source to get information on an existing Check Point Threat Emulation Image.
---

# Data Source: checkpoint_management_threat_emulation_image

Use this data source to get information on an existing Check Point Threat Emulation Image.

## Example Usage

```hcl
data "checkpoint_management_threat_emulation_image" "data_test" {
    name = "threat-emulation-image1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Image name.
* `image_id` - (Optional) Image id.
* `description` - (Computed) Image description.
* `display_name` - (Computed) Image display name.
* `image_type` - (Computed) Image type.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
