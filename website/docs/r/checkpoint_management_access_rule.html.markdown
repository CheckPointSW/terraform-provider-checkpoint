---
layout: "checkpoint"
page_title: "checkpoint_management_access_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-access-rule"
description: |-
  This resource allows you to add/update/delete Check Point Access Rule.
---

# checkpoint_management_access_rule

This resource allows you to add/update/delete Check Point Access Rule.

## Example Usage


```hcl
resource "checkpoint_management_access_rule" "rule1" {
  layer = "Network"
  position = {top = "top"}
  name = "test1"
}

resource "checkpoint_management_access_rule" "rule2" {
  layer = "Network"
  position = {below = checkpoint_management_access_rule.rule1.name}
  name = "test2"
  enabled = true
}

resource "checkpoint_management_access_rule" "rule3" {
  layer = "Network"
  position = {below = checkpoint_management_access_rule.rule2.name}
  name = "test3"
  action = "Accept"
  action_settings = {
    enable_identity_captive_portal = true
  }
  source = ["DMZNet", "DMZZone", "WirelessZone"]
  enabled = true
  destination = ["InternalNet", "CPDShield"]
  destination_negate = true
}

resource "checkpoint_management_access_rule" "rule4" {
  layer = "Network"
  position = {below = checkpoint_management_access_rule.rule3.name}
  name = "test4"
  track = {
    type = "Log"
  }
  enabled = false
}

resource "checkpoint_management_access_rule" "rule5" {
  layer = "Network"
  position = {below = checkpoint_management_access_rule.rule4.name}
  name = "test5"
  action = "Accept"
}

resource "checkpoint_management_access_rule" "rule6" {
  layer = "Network"
  position = {below = checkpoint_management_access_rule.rule5.name}
  name = "test6"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `name` - (Required) Rule name.
* `action` - (Optional) \"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".
* `action_settings` - (Optional) Action settings. Action settings blocks are documented below.
* `content` - (Optional) List of processed file types that this rule applies on.
* `content_direction` - (Optional) On which direction the file types processing is applied.
* `content_negate` - (Optional) True if negate is set for data.
* `custom_fields` - (Optional) Custom fields. Custom fields blocks are documented below.
* `destination` - (Optional) Collection of Network objects identified by the name or UID.
* `destination_negate` - (Optional) True if negate is set for destination.
* `enabled` - (Optional) Enable/Disable the rule.
* `inline_layer` - (Optional) Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".
* `install_on` - (Optional) Which Gateways identified by the name or UID to install the policy on.
* `service` - (Optional) Collection of Network objects identified by the name or UID.
* `service_negate` - (Optional) True if negate is set for service.
* `source` - (Optional) Collection of Network objects identified by the name or UID.
* `source_negate` - (Optional) True if negate is set for source.
* `time` - (Optional) List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".
* `track` - (Optional) Track Settings. Track Settings blocks are documented below.
* `user_check` - (Optional) User check settings. User check settings blocks are documented below.
* `vpn` - (Optional) Communities or Directional.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.

`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.

`action_settings` supports the following:

* `enable_identity_captive_portal` - (Optional) N/A.
* `limit` - (Optional) N/A.

`custom_fields` supports the following:

* `field_1` - (Optional) First custom field.
* `field_2` - (Optional) Second custom field.
* `field_3` - (Optional) Third custom field.

`track` supports the following:

* `accounting` - (Optional) Turns accounting for track on and off.
* `alert` - (Optional) Type of alert for the track.
* `enable_firewall_session` - (Optional) Determine whether to generate session log to firewall only connections.
* `per_connection` - (Optional) Determines whether to perform the log per connection.
* `per_session` - (Optional) Determines whether to perform the log per session.
* `type` - (Optional) \"Log\", \"Extended Log\", \"Detailed Log\", \"None\".

`user_check` supports the following:

* `confirm` - (Optional) N/A.
* `custom_frequency` - (Optional) N/A. Custom Frequency blocks are documented below.
* `frequency` - (Optional) N/A.
* `interaction` - (Optional) N/A.

`custom_frequency` supports the following:

* `every` - (Optional) N/A.
* `unit` - (Optional) N/A. 













