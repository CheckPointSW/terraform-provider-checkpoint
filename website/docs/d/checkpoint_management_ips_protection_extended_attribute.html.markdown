---
layout: "checkpoint"
page_title: "checkpoint_management_ips_protection_extended_attribute"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-ips-protection-extended-attribute"
description: |-
This resource allows you to execute Check Point Ips Protection Extended Attribute.
---

# Data Source: checkpoint_management_ips_protection_extended_attribute

This resource allows you to execute Check Point Ips Protection Extended Attribute.

## Example Usage

```hcl
data "checkpoint_management_ips_protection_extended_attribute" "data_ips_protection_extended_attribute" {
	name = "File Type"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `values` - The Object content. Values blocks are documented below.

`values` supports the following:
* `name` - Object name.
* `uid` - Object unique identifier.