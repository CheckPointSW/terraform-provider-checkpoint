---
layout: "checkpoint"
page_title: "checkpoint_management_ldap_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-ldap-group"
description: |-
This resource allows you to execute Check Point Ldap Group.
---

# checkpoint_management_ldap_group

This resource allows you to execute Check Point Ldap Group.

## Example Usage


```hcl
resource "checkpoint_management_ldap_group" "example" {
  name = "TestLdapGroup"
  account_unit = "TestLdapAccountUnit"
  scope = "only_sub_tree"
  account_unit_branch = "OU=Test"
  sub_tree_prefix = "CA=AC"
  group_prefix = "N=TestGroup"
  apply_filter_for_dynamic_group = true
  ldap_filter = "N=AnotherGroup"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `account_unit` - (Required) LDAP account unit of the group. 
Identified by name or UID. 
* `scope` - (Optional) Group's scope. There are three possible ways of defining a group, based on the users defined on the Account Unit. 
* `account_unit_branch` - (Optional) Branch of the selected LDAP Account Unit. 
* `sub_tree_prefix` - (Optional) Sub tree prefix of the selected branch. <font color="red">Relevant only when</font> 'scope' is set to 'only_sub_prefix'. Must be in DN syntax. 
* `group_prefix` - (Optional) Group name in the selected branch. <font color="red">Required only when</font> 'scope' is set to 'only_group_in_branch'. Must be in DN syntax. 
* `apply_filter_for_dynamic_group` - (Optional) Indicate whether to apply LDAP filter for dynamic group. <font color="red">Relevant only when</font> 'scope' is not set to 'only_group_in_branch'. 
* `ldap_filter` - (Optional) LDAP filter for the dynamic group. <font color="red">Relevant only when</font> 'apply-filter-for-dynamic-group' is set to 'true'. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
