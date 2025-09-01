---
layout: "checkpoint"
page_title: "checkpoint_management_ldap_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-ldap-group"
description: |- Use this data source to get information on an existing LDAP Group.
---


# checkpoint_management_ldap_group

Use this data source to get information on an existing LDAP Group.

## Example Usage


```hcl
data "checkpoint_management_ldap_group" "data_ldap_group" {
  name = "ldap_group_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `account_unit` - LDAP account unit of the group. Identified by name or UID.
* `scope` - Group's scope. There are three possible ways of defining a group, based on the users defined on the Account Unit.
* `account_unit_branch` - Branch of the selected LDAP Account Unit.
* `sub_tree_prefix` - Sub tree prefix of the selected branch.
* `group_prefix` - Group name in the selected branch.
* `apply_filter_for_dynamic_group` - Indicate whether to apply LDAP filter for dynamic group.
* `ldap_filter` - LDAP filter for the dynamic group.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
