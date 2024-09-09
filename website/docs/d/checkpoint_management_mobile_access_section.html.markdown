---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_section"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-section"
description: |-
Use this data source to get information on an existing Check Point Mobile Access Section.
---

# Data Source: checkpoint_management_mobile_access_section

Use this data source to get information on an existing Check Point Mobile Access Section.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_section" "example" {
  name = "New Section 1"
  position = {top = "top"}
}
data "checkpoint_management_mobile_access_section" "data" {
  name = "${checkpoint_management_mobile_access_section.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.

