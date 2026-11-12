---
layout: "checkpoint"
page_title: "checkpoint_management_threat_rule_exception_rulebase"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-rule-exception-rulebase"
description: |-
Use this data source to get information on an existing Check Point threat-rule-exception-rulebase
---

# Data Source: checkpoint_management_threat_rule_exception_rulebase

Use this data source to get information on an existing Check Point threat-rule-exception-rulebase

## Example Usage


```hcl
data "checkpoint_management_threat_rule_exception_rulebase" "base1" {
  name = "Standard Threat Prevention"
  rule_number  = 1
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Optional) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `filter` - (Optional) Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. If an operator is not used, the default OR operator applies.
* `filter_settings` -(Optional) Enable enforce end user domain. filter_settings blocks are documented below.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional)  Number of the results to initially skip.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.
* `package` - (Optional) Name of the package.
* `use_object_dictionary` - (Optional) boolean flag. indicate whether to use object dictionary in the response (default true).
* `name` - 	The name of the exception.
* `uid` - Object unique identifier.
* `from` - From which element number the query was done.
* `rulebase` - Array that contain rulebase for each group of the matched rule.
* `to` - To which element number the query was done.
* `total` - Total number of elements returned by the query.
* `objects_dictionary` - List of object that are part of the rulebase as services,sources etc..


`filter_settings` supports the following:

* `search_mode` - When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.
* `expand_group_members` - (Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.
* `expand_group_with_exclusion_members` - (Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.
* `match_on_any` - (Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.
* `match_on_group_with_exclusion` - (Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.
* `match_on_negate` - (Optional), can only be used when search_mode is set to "packet") Whether to match on a negated cell.

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.


`rulebase` supports the following:

* `name` - 	The name of the exception group.
* `uid` - Object unique identifier.
* `type` - Object type.
* `from` - From which element number the query was done.
* `rulebase` - Array that contain threat exception for a specific exception group.
* `to` - To which element number the query was done.

`rulebase` supports the following:

* `action` - Action-the enforced profile.
* `install_on` - Which Gateways identified by the name or UID to install the policy on.
* `source` - Collection of Network objects identified by the name or UID.
* `source_negate` - True if negate is set for source.
* `destination` - Collection of Network objects identified by the name or UID.
* `destination_negate` - True if negate is set for destination.
* `protected_scope` - Collection of objects defining Protected Scope identified by the name or UID.
* `protected_scope_negate` - True if negate is set for Protected Scope.
* `service` - Collection of Network objects identified by the name or UID.
* `service_negate` - True if negate is set for service.
* `track` - Packet tracking.

`objects_dictionary` supports the following:

* `name` - 	The name of the Object.
* `uid` - Object unique identifier.
* `type` - Object type.