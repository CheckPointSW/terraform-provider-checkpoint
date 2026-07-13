---
layout: "checkpoint"
page_title: "checkpoint_management_threat_extraction_file_type"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-extraction-file-type"
description: |-
Use this data source to get information on an existing Check Point Threat Extraction File Type.
---

# Data Source: checkpoint_management_threat_extraction_file_type

Use this data source to get information on an existing Check Point Threat Extraction File Type.

## Example Usage

```hcl
data "checkpoint_management_threat_extraction_file_type" "data_test" {
    name = "threat-extraction-file-type1"
}
```

## Argument Reference

The following arguments are supported:

* `file_type_id` - (Optional) File type id.
* `file_type` - (Optional) File type extension.
* `description` - (Computed) File type description.
* `enabled` - (Computed) Enable support for Threat Extraction.
* `icon` - (Computed) File type icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
