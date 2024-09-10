---
layout: "checkpoint"
page_title: "checkpoint_management_mobile_access_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-mobile-access-rule"
description: |-
Use this data source to get information on an existing Check Point Mobile Access Rule.
---

# Data Source: checkpoint_management_mobile_access_rule

Use this data source to get information on an existing Check Point Mobile Access Rule.

## Example Usage


```hcl
resource "checkpoint_management_mobile_access_rule" "example" {
  name = "Rule 1"
  applications = ["N", "e", "w", " ", "A", "p", "p", "l", "i", "c", "a", "t", "i", "o", "n",]
  user_groups = ["my_group",]
  position = {top = "top"}
}
data "checkpoint_management_mobile_access_rule" "data" {
  uid = "${checkpoint_management_mobile_access_rule.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `user_groups` -  User groups that will be associated with the apps - identified by the name or UID.user_groups blocks are documented below.
* `applications` - Available apps that will be associated with the user groups - identified by the name or UID.applications blocks are documented below.
* `enabled` - Enable/Disable the rule. 
* `install_on` -  Which Gateways identified by the name or UID to install the policy on.install_on blocks are documented below.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `comments` - Comments string. 

