---
layout: "checkpoint"
page_title: "checkpoint_management_data_access_section"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-section"
description: |-
  Use this data source to get information on an existing Check Point Access Section.
---

# Data Source: checkpoint_management_data_access_section

Use this data source to get information on an existing Check Point Access Section.

## Example Usage


```hcl
resource "checkpoint_management_access_layer" "access_layer" {
        name = "myaccesslayer"
        detect_using_x_forward_for = false
        applications_and_url_filtering = true
}

resource "checkpoint_management_access_section" "access_section" {
    name = "myaccesssection"
	layer = "${checkpoint_management_access_layer.access_layer.name}"
	position = {top = "top"}
}

data "checkpoint_management_data_access_section" "data_access_section" {
    name = "${checkpoint_management_access_section.access_section.name}"
    layer = "${checkpoint_management_access_section.access_section.layer}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID. 
* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier. 