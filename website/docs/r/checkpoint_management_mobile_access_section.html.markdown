---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_section"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-section"
description: |-
This resource allows you to execute Check Point Mobile Access Section.
---

# checkpoint_management_mobile_access_section

This resource allows you to execute Check Point Mobile Access Section.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_section" "example" {
  name = "New Section 1"
  position = {top = "top"}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.
