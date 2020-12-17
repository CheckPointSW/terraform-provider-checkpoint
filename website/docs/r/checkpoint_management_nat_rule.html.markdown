---
layout: "checkpoint"
page_title: "checkpoint_management_nat_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-nat-rule"
description: |-
  This resource allows you to add/update/delete Check Point NAT Rule.
---

# Resource: checkpoint_management_nat_rule

This resource allows you to add/update/delete Check Point NAT Rule.

## Example Usage


```hcl
resource "checkpoint_management_nat_rule" "rule1" {
  package = "Standard"
  position = {top = "top"}
  name = "rule1"
}

resource "checkpoint_management_nat_rule" "rule2" {
  package = "Standard"
  position = {below = checkpoint_management_access_rule.rule1.name}
  name = "rule2"
}

resource "checkpoint_management_nat_rule" "rule3" {
  package = "Standard"
  position = {below = checkpoint_management_access_rule.rule2.name}
  name = "rule3"
}
```

## Argument Reference

The following arguments are supported:

* `package` - (Required) Name of the package.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `name` - (Optional) Rule name.
* `enabled` - (Optional) Enable/Disable the rule.
* `method` - (Optional) Nat method.
* `install_on` - (Optional) Which Gateways identified by the name or UID to install the policy on.
* `original_destination` - (Optional) Original destination.
* `original_service` - (Optional) Original service.
* `original_source` - (Optional) Original source.
* `translated_destination` - (Optional) Translated destination.
* `translated_service` - (Optional) Translated service.
* `translated_source` - (Optional) Translated source.
* `auto_generated` - (Computed) Auto generated.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.

## Import

`checkpoint_management_nat_rule` can be imported by using the following format: PACKAGE_NAME;RULE_UID

```
$ terraform import checkpoint_management_nat_rule.example Standard;9423d36f-2d66-4754-b9e2-e9f4493751d3
```