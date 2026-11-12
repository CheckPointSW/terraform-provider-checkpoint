---
layout: "checkpoint"
page_title: "checkpoint_management_threat_emulation_file_type"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-emulation-file-type"
description: |-
Use this data source to get information on an existing Check Point Threat Emulation File Type.
---

# Data Source: checkpoint_management_threat_emulation_file_type

Use this data source to get information on an existing Check Point Threat Emulation File Type.

## Example Usage

```hcl
data "checkpoint_management_threat_emulation_file_type" "data_test" {
    name = "threat-emulation-file-type1"
}
```

## Argument Reference

The following arguments are supported:

* `file_type_id` - (Optional) File type id.
* `file_type` - (Optional) File type extension.
* `description` - (Computed) File type description.
* `enabled` - (Computed) Enable support for Threat Emulation.
* `icon` - (Computed) File type icon.
* `supported_platforms` - (Computed) Supported platforms for Threat Emulation.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
