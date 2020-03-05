---
layout: "checkpoint"
page_title: "checkpoint_management_exception_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-exception-group"
description: |-
This resource allows you to execute Check Point Exception Group.
---

# checkpoint_management_exception_group

This resource allows you to execute Check Point Exception Group.

## Example Usage


```hcl
resource "checkpoint_management_exception_group" "example" {
  name = "exception_group_2"
  apply_on = "manually-select-threat-rules"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `applied_profile` - (Optional) The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile. 
* `applied_threat_rules` - (Optional) The threat rules to apply this group on in the case of apply-on manually-select-threat-rules.applied_threat_rules blocks are documented below.
* `apply_on` - (Optional) An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`applied_threat_rules` supports the following:

* `layer` - (Optional) The layer of the threat rule to which the group is to be attached. 
* `name` - (Optional) The name of the threat rule to which the group is to be attached. 
* `rule_number` - (Optional) The rule-number of the threat rule to which the group is to be attached. 
* `position` - (Optional) Position in the rulebase. 
