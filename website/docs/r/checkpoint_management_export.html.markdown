---
layout: "checkpoint"
page_title: "checkpoint_management_export"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-export"
description: |-
This resource allows you to execute Check Point Export.
---

# checkpoint_management_export

This resource allows you to execute Check Point Export.

## Example Usage


```hcl
resource "checkpoint_management_export" "example" {
  export_files_by_class = true
}
```

## Argument Reference

The following arguments are supported:

* `exclude_classes` - (Optional) N/Aexclude_classes blocks are documented below.
* `exclude_topics` - (Optional) N/Aexclude_topics blocks are documented below.
* `export_files_by_class` - (Optional) N/A 
* `include_classes` - (Optional) N/Ainclude_classes blocks are documented below.
* `include_topics` - (Optional) N/Ainclude_topics blocks are documented below.
* `query_limit` - (Optional) N/A 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

