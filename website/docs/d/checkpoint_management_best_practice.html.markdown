---
layout: "checkpoint"
page_title: "checkpoint_management_best_practice"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-best-practice"
description: |-
 Use this data source to get information on an existing Check Point Best Practice.
---

# Data Source: checkpoint_management_best_practice

Use this data source to get information on an existing Check Point Best Practice.

## Example Usage

```hcl
# Get best practice by name
data "checkpoint_management_best_practice" "test" {
  name = "Firewall Best Practice"
  show_regulations = true
}

# Get best practice by UID
data "checkpoint_management_best_practice" "test" {
  uid = "ed997ff8-6709-4d71-a713-99bf01711cd5"
}

# Get best practice by best practice ID
data "checkpoint_management_best_practice" "test" {
  best_practice_id = "BP001"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `best_practice_id` - (Optional) Best Practice ID.
* `show_regulations` - (Optional) Show the applicable regulations of the Best Practice.
* `action_item` - Required action item to comply with the Best Practice.
* `active` - Shows if the Best Practice is active.
* `blade` - The Software Blade name of the Best Practice.
* `description` - Description of the Best Practice.
* `due_date` - Shows if there is a due date for the action item of this Best Practice.
* `status` - The current status of the Best Practice.
* `user_defined` - Shows if the Best Practice is a user-defined Best Practice.
* `comments` - Comments string.
* `regulations` - The applicable regulations of the Best Practice. Appears only when `show_regulations` is set to `true`. regulations blocks are documented below.
* `relevant_objects` - The applicable objects of the Best Practice. relevant_objects blocks are documented below.
* `user_defined_firewall` - The definitions of the user-defined Firewall Best Practice. Relevant only for Firewall Best Practices created by the user. user_defined_firewall blocks are documented below.
* `user_defined_gaia_os` - The definitions of the user-defined Gaia OS Best Practice. Relevant only for Gaia OS Best Practices created by the user. user_defined_gaia_os blocks are documented below.

`regulations` supports the following:

* `regulation_name` - The name of the regulation.
* `requirement_description` - The description of the requirement.
* `requirement_id` - The id of the requirement.
* `requirement_status` - The status of the requirement.
* `requirement_uid` - The unique identifier of the requirement.

`relevant_objects` supports the following:

* `relevant_objects_type` - The type of the relevant object.
* `access_rules_info` - The information about the relevant access rules. Appears only when `relevant_objects_type` is 'access-rule'. access_rules_info blocks are documented below.
* `cpm_relevant_objects_info` - The information about the relevant objects. Appears only when `relevant_objects_type` is 'cpm-relevant-object'. cpm_relevant_objects_info blocks are documented below.
* `ips_protections_info` - The information about the relevant ips-protection objects. Appears only when `relevant_objects_type` is 'ips-protection'. ips_protections_info blocks are documented below.

`access_rules_info` supports the following:

* `enabled` - Shows if the Compliance scan is enabled or not for this object.
* `layer_name` - The name of the relevant policy layer.
* `layer_uid` - The UID of the relevant policy layer.
* `policy_name` - The name of the relevant policy.
* `rule_indexes` - Comma-separated indexes of the relevant rules in the relevant policy and policy layer.
* `status` - The status of the relevant object.

`cpm_relevant_objects_info` supports the following:

* `cpm_relevant_object_type` - The type of the relevant object.
* `enabled` - Shows if the Compliance scan is enabled or not for this object.
* `name` - The name of the relevant object.
* `status` - The status of the relevant object.

`ips_protections_info` supports the following:

* `action` - The current action of the Threat Prevention profile.
* `enabled` - Shows if the Compliance scan is enabled or not for this object.
* `profile_name` - The name of the relevant Threat Prevention profile.
* `profile_uid` - The UID of the relevant Threat Prevention profile.
* `protection_name` - The name of the relevant IPS protection.
* `status` - The status of the relevant object.

`user_defined_firewall` supports the following:

* `policy_range_percentage` - User-defined policy range percentage to test.
* `policy_range_position` - User-defined policy range position.
* `poor_condition` - User-defined poor condition.
* `secure_condition` - User-defined secure condition.
* `tolerance` - User-defined tolerance. Appears only when the value of the 'violation-condition' parameter is 'Rule found'.
* `violation_condition` - User-defined violation condition.
* `user_defined_rules` - User-defined Firewall rules. user_defined_rules blocks are documented below.

`user_defined_rules` supports the following:

* `action` - User-defined actions. action blocks are documented below.
* `comment` - User-defined comment. comment blocks are documented below.
* `destination` - User-defined destination objects. destination blocks are documented below.
* `hit_count` - User-defined hit count value. hit_count blocks are documented below.
* `install_on` - User-defined "Install On" objects. install_on blocks are documented below.
* `name` - User-defined name. name blocks are documented below.
* `services_and_applications` - User-defined service and application objects. services_and_applications blocks are documented below.
* `source` - User-defined source objects. source blocks are documented below.
* `time` - User-defined time. time blocks are documented below.
* `track` - User-defined track actions. track blocks are documented below.
* `vpn` - User-defined VPN objects. vpn blocks are documented below.

`action` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`comment` supports the following:

* `condition_type` - The condition type.
* `value` - The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.

`destination` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`hit_count` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`install_on` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`name` supports the following:

* `condition_type` - The condition type.
* `value` - The condition match string. Appears only when the value of the 'condition-type' parameter is: 'Equals', 'Starts with', 'Ends with', 'Contains'.

`services_and_applications` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`source` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`time` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`track` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`vpn` supports the following:

* `negate` - Shows if the rule is negated.
* `reference_objects` - The reference objects. reference_objects blocks are documented below.

`reference_objects` supports the following:

* `name` - The name of the reference object.
* `reference_object_type` - The type of the reference object.
* `uid` - The UID of the reference object.

`user_defined_gaia_os` supports the following:

* `expected_output_base64` - The expected output of the script in the Base64.
* `practice_script_base64` - The script in Base64 to run on Gaia Security Gateways during the Compliance scans. 
