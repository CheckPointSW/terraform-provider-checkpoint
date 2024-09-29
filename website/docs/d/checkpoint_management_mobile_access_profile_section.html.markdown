---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_profile_section"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-profile-section"
description: |-
Use this data source to get information on an existing Check Point Mobile Access Profile Section.
---

# Data Source: checkpoint_management_mobile_access_profile_section

Use this data source to get information on an existing Check Point Mobile Access Profile Section.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_profile_section" "example" {
  name = "New Section 1"
  position = {top = "top"}
}
data "checkpoint_management_mobile_access_profile_section" "data" {
  uid = "${checkpoint_management_mobile_access_profile_section.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
