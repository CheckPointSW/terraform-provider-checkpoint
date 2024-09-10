---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-rule"
description: |-
This resource allows you to execute Check Point Mobile Access Rule.
---

# checkpoint_management_mobile_access_rule

This resource allows you to execute Check Point Mobile Access Rule.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_rule" "example" {
  name = "Rule 1"
  applications = ["N", "e", "w", " ", "A", "p", "p", "l", "i", "c", "a", "t", "i", "o", "n",]
  user_groups = ["my_group",]
  position = {top = "top"}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `user_groups` - (Optional) User groups that will be associated with the apps - identified by the name or UID.user_groups blocks are documented below.
* `applications` - (Optional) Available apps that will be associated with the user groups - identified by the name or UID.applications blocks are documented below.
* `enabled` - (Optional) Enable/Disable the rule. 
* `install_on` - (Optional) Which Gateways identified by the name or UID to install the policy on.install_on blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.

