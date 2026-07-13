---
layout: "checkpoint"
page_title: "checkpoint_management_firewall_best_practice"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-firewall-best-practice"
description: |-
Use this data source to get information on an existing Check Point Firewall Best Practice.
---

# Data Source: checkpoint_management_firewall_best_practice

Use this data source to get information on an existing Check Point Firewall Best Practice.

## Example Usage

```hcl
data "checkpoint_management_firewall_best_practice" "data_test" {
    name = "firewall-best-practice1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Best Practice Name.
* `uid` - (Optional) Object unique identifier.
* `best_practice_id` - (Optional) Best Practice ID.
* `show_regulations` - (Optional) Show the applicable regulations of the Best Practice.
* `show_relevant_objects` - (Optional) Show the relevant objects of the Best Practice.
* `action_item` - (Computed) Required action item to comply with the Best Practice.
* `description` - (Computed) Description of the Best Practice.
* `enabled` - (Computed) The activation status of the best practice.
* `expiration` - (Computed) Deactivation expiration settings. Present only when the best practice is disabled. expiration blocks are documented below.
* `policy_range_percentage` - (Computed) Percentage of the rule base to scan, 0-100.
* `policy_range_position` - (Computed) The direction of the scan.
* `poor_condition` - (Computed) Visibility of poor-result rules in the Relevant Objects pane.
* `regulations` - (Computed) The applicable regulations of the Best Practice. Appears only when the value of the 'show-regulations' parameter is set to 'true'. regulations blocks are documented below.
* `relevant_objects` - (Computed) The applicable objects of the Best Practice. Appears only when the value of the 'show-relevant-objects' parameter is set to 'true'. relevant_objects blocks are documented below.
* `rule` - (Computed) The rule criteria the firewall best practice evaluates against the rule base. rule blocks are documented below.
* `secure_condition` - (Computed) Visibility of secure-result rules in the Relevant Objects pane.
* `status` - (Computed) The current status of the Best Practice.
* `tolerance` - (Computed) Number of matches allowed before a violation is created. Relevant only when violation-condition is set to 'Rule found'.
* `violation_condition` - (Computed) Define when a violation occurs: 'Rule found' means the criteria match a rule; 'Rule not found' means no rule matches.


`expiration` supports the following:

* `comment` - (Computed) The reason the best practice was deactivated.
* `expire_on` - (Computed) When the deactivation expires. Date and time represented in international ISO 8601 format. expire_on blocks are documented below.
* `mode` - (Computed) Whether the deactivation never expires or expires on a specific date.


`regulations` supports the following:

* `regulation_name` - (Computed) The name of the regulation.
* `requirement_description` - (Computed) The description of the requirement.
* `requirement_id` - (Computed) The id of the requirement.
* `requirement_status` - (Computed) The status of the requirement.
* `requirement_uid` - (Computed) The unique identifier of the requirement.


`relevant_objects` supports the following:

* `access_rules_info` - (Computed) The information about the relevant access rules. Appears only when the value of the 'relevant-objects-type' parameter is 'access-rule'. access_rules_info blocks are documented below.
* `cpm_relevant_objects_info` - (Computed) The information about the relevant objects. Appears only when the value of the 'relevant-objects-type' parameter is 'cpm-relevant-object'.
* `ips_protections_info` - (Computed) The information about the relevant ips-protection objects. Appears only when the value of the 'relevant-objects-type' parameter is 'ips-protection'. ips_protections_info blocks are documented below.
* `relevant_objects_type` - (Computed) The type of the relevant object.


`rule` supports the following:

* `source` - (Computed) Network objects to match in the rule Source column.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_source` - (Computed) Shows if the source values are negated.
* `destination` - (Computed) Network objects to match in the rule Destination column.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_destination` - (Computed) Shows if the destination values are negated.
* `vpn` - (Computed) VPN communities to match.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_vpn` - (Computed) Shows if the vpn values are negated.
* `services_and_applications` - (Computed) Services, applications, categories or sites to match.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_services_and_applications` - (Computed) Shows if the services and applications values are negated.
* `install_on` - (Computed) Security Gateways or Clusters the rule applies to.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_install_on` - (Computed) Shows if the install-on values are negated.
* `time` - (Computed) Time objects the rule applies to.Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `negate_time` - (Computed) Shows if the time values are negated.
* `action` - (Computed) Rule actions to match.
* `negate_action` - (Computed) Shows if the action values are negated.
* `track` - (Computed) Tracking methods to match.
* `negate_track` - (Computed) Shows if the track values are negated.
* `hit_count` - (Computed) Hit-count levels to match.
* `negate_hit_count` - (Computed) Shows if the hit-count values are negated.
* `name_condition` - (Computed) Match the rule name against a text condition. name_condition blocks are documented below.
* `comment_condition` - (Computed) Match the rule comment against a text condition. comment_condition blocks are documented below.


`expire_on` supports the following:

* `iso_8601` - (Computed) Date and time represented in international ISO 8601 format.
* `posix` - (Computed) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`access_rules_info` supports the following:

* `enabled` - (Computed) Shows if the Compliance scan is enabled or not for this object.
* `layer_name` - (Computed) The name of the relevant policy layer.
* `layer_uid` - (Computed) The UID of the relevant policy layer.
* `policy_name` - (Computed) The name of the relevant policy.
* `rule_indexes` - (Computed) Comma-separated indexes of the relevant rules in the relevant policy and policy layer.
* `status` - (Computed) The status of the relevant object.


`ips_protections_info` supports the following:

* `action` - (Computed) The current action of the Threat Prevention profile.
* `enabled` - (Computed) Shows if the Compliance scan is enabled or not for this object.
* `profile_name` - (Computed) The name of the relevant Threat Prevention profile.
* `profile_uid` - (Computed) The UID of the relevant Threat Prevention profile.
* `protection_name` - (Computed) The name of the relevant IPS protection.
* `status` - (Computed) The status of the relevant object.


`name_condition` supports the following:

* `condition_type` - (Computed) The condition type.
* `value` - (Computed) The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.


`comment_condition` supports the following:

* `condition_type` - (Computed) The condition type.
* `value` - (Computed) The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.
