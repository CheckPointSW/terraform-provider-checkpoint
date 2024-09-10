---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_profile_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-profile-rule"
description: |-
Use this data source to get information on an existing Mobile Access Profile Rule.
---

# Data Source: checkpoint_management_mobile_access_profile_rule

Use this data source to get information on an existing Check Point Mobile Access Profile Rule.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_profile_rule" "example" {
  name = "Rule 1"
  mobile_profile = "Default_Profile"
  user_groups = ["my_group",]
  position = {top = "top"}
}

data "checkpoint_management_mobile_access_profile_rule" "data" {
  name = "${checkpoint_management_mobile_access_profile_rule.example.name}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `mobile_profile` - Profile configuration for User groups - identified by the name or UID. 
* `user_groups` -  User groups that will be configured with the profile object - identified by the name or UID.user_groups blocks are documented below.
* `enabled` - Enable/Disable the rule. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `comments` - Comments string.
 
