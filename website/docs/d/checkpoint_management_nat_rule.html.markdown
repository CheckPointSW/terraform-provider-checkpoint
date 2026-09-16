---
layout: "checkpoint"
page_title: "checkpoint_management_nat_rule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-nat-rule"
description: |-
  This resource allows you to execute Check Point NAT Rule.
---

# Data Source: checkpoint_management_nat_rule

This resource allows you to execute Check Point NAT Rule.

## Example Usage


```hcl
resource "checkpoint_management_nat_rule" "test" {
    name = "natrule"
    package = "Standard"
    position {top = "top"}
}

data "checkpoint_management_nat_rule" "test" {
    package = "${checkpoint_management_nat_rule.test.package}"
    name = "${checkpoint_management_nat_rule.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `package` - (Required) Name of the package.
* `name` - (Optional) Rule name.
* `uid` - (Optional) Object unique identifier.   
* `enabled` - Enable/Disable the rule.
* `method` - Nat method.
* `install_on` - Which Gateways identified by the name or UID to install the policy on.
* `original_destination` - Original destination.
* `original_service` - Original service.
* `original_source` - Original source.
* `translated_destination` - Translated destination.
* `translated_service` - Translated service.
* `translated_source` - Translated source.
* `auto_generated` - Auto generated.
* `comments` - Comments string.
* `hits` - Hits count object.hits blocks are documented below.


`hits` supports the following:

* `first_date` - N/A.first_date blocks are documented below.
* `last_date` - N/A.last_date blocks are documented below.
* `level` - N/A.
* `percentage` - N/A.
* `value` - N/A.


`first_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`last_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
