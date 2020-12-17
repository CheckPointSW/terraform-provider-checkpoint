---
layout: "checkpoint"
page_title: "checkpoint_management_nat_section"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-nat-section"
description: |-
  This resource allows you to execute Check Point NAT section.
---

# Data Source: checkpoint_management_nat_section

This resource allows you to execute Check Point NAT section.

## Example Usage


```hcl
resource "checkpoint_management_nat_section" "test" {
    name = "nat section"
    package = "Standard"
    position = {top = "top"}
}

data "checkpoint_management_nat_section" "nat_section" {
    package = "${checkpoint_management_nat_section.test.package}"
    name = "${checkpoint_management_nat_section.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `package` - (Required) Name of the package.
* `uid` - (Optional) Object unique identifier.   
* `name` - (Optional) Object name.