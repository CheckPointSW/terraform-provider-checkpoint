---
layout: "checkpoint"
page_title: "checkpoint_management_regulation"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-regulation"
description: |-
This resource allows you to execute Check Point Regulation.
---

# checkpoint_management_regulation

This resource allows you to execute Check Point Regulation.

## Example Usage


```hcl
resource "checkpoint_management_regulation" "example" {
  name = "MyReg"
  full_name = "My New Regulation"
  comments = "My compliance regulation"
}
```

## Argument Reference

The following arguments are supported:

* `full_name` - (Required) Regulation full name. 
* `name` - (Required) Regulation name. 
* `enabled` - (Optional) Determines if the regulation is enabled. 
* `new_full_name` - (Optional) Object new full name. Must be unique in the domain. 
* `show_requirements` - (Optional) Show the requirements of the regulation.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments about this regulation.
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `requirements` - (Computed) The requirements of the regulation, identified by name. Appears only when `show_requirements` is set to `true`.
