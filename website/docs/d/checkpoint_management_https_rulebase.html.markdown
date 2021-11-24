---
layout: "checkpoint"
page_title: "checkpoint_management_https_rulebase"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-https-rulebase"
description: |- This resource allows you to execute Check Point Https Rule Base.
---

# Data Source: checkpoint_management_https_rulebase

Use this data source to get information on an existing https RuleBase.

## Example Usage

```hcl
data "checkpoint_management_rulebase" "https_rulebase" {
  name = "Default Layer"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `filter` - Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. If an operator is not used, the default OR operator applies.
* `filter_settings` - Enable enforce end user domain. filter_settings blocks are documented below.
* `limit` - The maximal number of returned results.
* `offset` - Number of the results to initially skip.
* `order` - Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.
* `package` - Name of the package.
* `dereference_group_members` - Indicates whether to dereference "members" field by details level for every object in reply.
* `show_membership` - Indicates whether to calculate and show "groups" field for every object in reply.

`filter_settings` supports the following:

* `search_mode` - When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.
* `expand_group_members` - (Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.
* `expand_group_with_exclusion_members` - (Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.
* `match_on_any` - (Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.
* `match_on_group_with_exclusion` - (Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.
* `match_on_negate` - (Optional, can only be used when search_mode is set to "packet") Whether to match on a negated cell.

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.
