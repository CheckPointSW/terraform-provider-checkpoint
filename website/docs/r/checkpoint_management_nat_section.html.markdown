---
layout: "checkpoint"
page_title: "checkpoint_management_nat_section"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-nat-section"
description: |-
  This resource allows you to add/update/delete Check Point NAT section.
---

# Resource: checkpoint_management_nat_section

This resource allows you to add/update/delete Check Point NAT section.

## Example Usage


```hcl
resource "checkpoint_management_nat_section" "nat_section" {
    name = "nat section"
    package = "Standard"
    position = { "top": "top" }
}
```

## Argument Reference

The following arguments are supported:

* `package` - (Required) Name of the package.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `name` - (Optional) Object name.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.