---
layout: "checkpoint"
page_title: "checkpoint_management_set_threat_extraction_file_types"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-threat-extraction-file-types"
description: |-
This resource allows you to execute Check Point Set Threat Extraction File Types.
---

# checkpoint_management_set_threat_extraction_file_types

This resource allows you to execute Check Point Set Threat Extraction File Types.

## Example Usage


```hcl
resource "checkpoint_management_set_threat_extraction_file_types" "example" {
  file_types {
    file_type = "pdf"
    enabled   = false
  }
  file_types {
    file_type = "docx"
    enabled   = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `file_types` - (Required) List of Threat Extraction file type updates. Each entry sets 'enabled' on the file type identified by 'file-type-id' or 'file-type'.file_types blocks are documented below.
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`file_types` supports the following:

* `file_type_id` - (Optional) File type id. 
* `file_type` - (Optional) File type extension. 
* `enabled` - (Optional) Enable support for Threat Extraction.
