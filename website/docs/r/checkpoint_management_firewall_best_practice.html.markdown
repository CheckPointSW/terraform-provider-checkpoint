---
layout: "checkpoint"
page_title: "checkpoint_management_firewall_best_practice"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-firewall-best-practice"
description: |-
This resource allows you to execute Check Point Firewall Best Practice.
---

# checkpoint_management_firewall_best_practice

This resource allows you to execute Check Point Firewall Best Practice.

## Example Usage


```hcl
resource "checkpoint_management_firewall_best_practice" "example" {
  name = "Clean-up rule defined in Access Policy"
  action_item = "Define a clean-up rule at the end of the policy."
  description = "Checks that the rule base ends with a clean-up rule."
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Best Practice Name. 
* `action_item` - (Required) To comply with Best Practice, do this action item. 
* `description` - (Required) Description of the Best Practice. 
* `enabled` - (Optional) The activation status of the best practice. 
* `expiration` - (Optional) Deactivation expiration settings. Required only if enabled is set to false. expiration blocks are documented below.
* `policy_range_percentage` - (Optional) The percentage of the Rule Base to scan (0-100). 
* `policy_range_position` - (Optional) The direction of the scan. 
* `poor_condition` - (Optional) Visibility of poor-result rules in the Relevant Objects pane. 
* `rule` - (Optional) The rule criteria the firewall best practice evaluates against the rule base.rule blocks are documented below.
* `secure_condition` - (Optional) Visibility of secure-result rules in the Relevant Objects pane. 
* `tolerance` - (Optional) Number of matches allowed before a violation is created. Valid values: between 0 and 100.<br><font color="red">Required only if</font> violation-condition is set to 'Rule found'. 
* `violation_condition` - (Optional) Define when a violation occurs: 'Rule found' means the criteria match a rule; 'Rule not found' means no rule matches. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `best_practice_id` - (Optional) Best Practice ID. Relevant only for updating object

`expiration` supports the following:

* `comment` - (Optional) The reason for deactivating the best practice. 
* `expire_on` - (Optional) When the deactivation expires. Date and time represented in international ISO 8601 format. Relevant only if mode is set to 'expire-on'. 
* `mode` - (Optional) Whether the deactivation never expires or expires on a specific date. 


`rule` supports the following:

* `source` - (Optional) Network objects to match in the rule Source column. Identified by name or UID.
* `negate_source` - (Optional) Shows if the source values are negated. 
* `destination` - (Optional) Network objects to match in the rule Destination column. Identified by name or UID.
* `negate_destination` - (Optional) Shows if the destination values are negated. 
* `vpn` - (Optional) VPN communities to match. Identified by name or UID.
* `negate_vpn` - (Optional) Shows if the vpn values are negated. 
* `services_and_applications` - (Optional) Services, applications, categories or sites to match. Identified by name or UID.
* `negate_services_and_applications` - (Optional) Shows if the services and applications values are negated. 
* `install_on` - (Optional) Security Gateways or Clusters the rule applies to. Identified by name or UID.
* `negate_install_on` - (Optional) Shows if the install-on values are negated. 
* `time` - (Optional) Time objects the rule applies to. Identified by name or UID.
* `negate_time` - (Optional) Shows if the time values are negated. 
* `action` - (Optional) Rule actions to match.
* `negate_action` - (Optional) Shows if the action values are negated. 
* `track` - (Optional) Tracking methods to match.
* `negate_track` - (Optional) Shows if the track values are negated. 
* `hit_count` - (Optional) Hit-count levels to match.
* `negate_hit_count` - (Optional) Shows if the hit-count values are negated. 
* `name_condition` - (Optional) Match the rule name against a text condition.name_condition blocks are documented below.
* `comment_condition` - (Optional) Match the rule comment against a text condition.comment_condition blocks are documented below.


`name_condition` supports the following:

* `condition_type` - (Optional) The condition type. 
* `value` - (Optional) The condition match string. Relevant only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'. 


`comment_condition` supports the following:

* `condition_type` - (Optional) The condition type. 
* `value` - (Optional) The condition match string. Relevant only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'. 
