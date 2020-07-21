---
layout: "checkpoint"
page_title: "checkpoint_management_data_https_section"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-https-section"
description: |-
  Use this data source to get information on an existing Check Point Https Section.
---

# checkpoint_management_data_https_section

Use this data source to get information on an existing Check Point Https Section.

## Example Usage


```hcl
resource "checkpoint_management_https_section" "https_section" {
        name = "HTTPS section"
		layer = "Default Layer"
        position = {top = "top"}
}

data "checkpoint_management_data_https_section" "data_https_section" {
    name = "${checkpoint_management_https_section.https_section.name}"
    layer = "${checkpoint_management_https_section.https_section.layer}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that holds the Object. Identified by the Name or UID. 
* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
