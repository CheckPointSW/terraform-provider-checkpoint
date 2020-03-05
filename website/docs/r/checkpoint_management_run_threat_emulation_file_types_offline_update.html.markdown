---
layout: "checkpoint"
page_title: "checkpoint_management_run_threat_emulation_file_types_offline_update"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-run-threat-emulation-file-types-offline-update"
description: |-
This resource allows you to execute Check Point Run Threat Emulation File Types Offline Update.
---

# checkpoint_management_run_threat_emulation_file_types_offline_update

This resource allows you to execute Check Point Run Threat Emulation File Types Offline Update.

## Example Usage


```hcl
resource "checkpoint_management_run_threat_emulation_file_types_offline_update" "example" {
  file_path = "/tmp/FileTypeUpdate.xml"
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) File path for offline update of Threat Emulation file types, the file path should be on the management machine. 
* `file_raw_data` - (Required) The contents of a file containing the Threat Emulation file types. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

