---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_profile_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-profile-rule"
description: |-
This resource allows you to execute Check Point Mobile Access Profile Rule.
---

# checkpoint_management_mobile_access_profile_rule

This resource allows you to execute Check Point Mobile Access Profile Rule.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_profile_rule" "example" {
  name = "Rule 1"
  mobile_profile = "Default_Profile"
  user_groups = ["my_group"]
  position = {top = "top"}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `mobile_profile` - (Optional) Profile configuration for User groups - identified by the name or UID. 
* `user_groups` - (Optional) User groups that will be configured with the profile object - identified by the name or UID.user_groups blocks are documented below.
* `enabled` - (Optional) Enable/Disable the rule. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.