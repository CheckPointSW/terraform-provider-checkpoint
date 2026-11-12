---
layout: "checkpoint"
page_title: "checkpoint_management_set_threat_extraction_file_type"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-threat-extraction-file-type"
description: |-
This resource allows you to execute Check Point Set Threat Extraction File Type.
---

# checkpoint_management_set_threat_extraction_file_type

This resource allows you to execute Check Point Set Threat Extraction File Type.

## Example Usage


```hcl
resource "checkpoint_management_set_threat_extraction_file_type" "example" {
  file_type = "pdf"
  enabled = false
}
```

## Argument Reference

The following arguments are supported:

* `file_type_id` - (Optional) File type id. 
* `file_type` - (Optional) File type extension. 
* `enabled` - (Optional) Enable support for Threat Extraction.