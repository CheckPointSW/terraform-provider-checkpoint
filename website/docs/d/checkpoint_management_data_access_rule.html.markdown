---
layout: "checkpoint"
page_title: "checkpoint_management_data_access_rule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-rule"
description: |- Use this data source to get information on an existing Check Point Access Rule.
---

# Data Source: checkpoint_management_data_access_rule

Use this data source to get information on an existing Check Point Access Rule.

## Example Usage

```hcl
resource "checkpoint_management_access_rule" "access_rule" {
  name = "My Rule"
  layer = "Network"
  position = { top = "top" }
  source = ["Any"]
  destination = ["Any"]
  service = ["Any"]
  content = ["Any"]
  time = ["Any"]
  install_on = ["Policy Targets"]
  track = {
    type = "Log"
    accounting = false
    alert = "none"
    enable_firewall_session = false
    per_connection = true
    per_session = false
  }
  custom_fields = {}
}

data "checkpoint_management_data_access_rule" "data_access_rule" {
  name = "${checkpoint_management_access_rule.access_rule.name}"
  layer = "${checkpoint_management_access_rule.access_rule.layer}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Rule name.
* `action` - \"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".
* `action_settings` - Action settings. Action settings blocks are documented below.
* `content` - List of processed file types that this rule applies on.
* `content_direction` - On which direction the file types processing is applied.
* `content_negate` - True if negate is set for data.
* `custom_fields` - Custom fields. Custom fields blocks are documented below.
* `destination` - Collection of Network objects identified by the name or UID.
* `destination_negate` - True if negate is set for destination.
* `enabled` - Enable/Disable the rule.
* `inline_layer` - Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".
* `install_on` - Which Gateways identified by the name or UID to install the policy on.
* `service` - Collection of Network objects identified by the name or UID.
* `service_negate` - True if negate is set for service.
* `source` - Collection of Network objects identified by the name or UID.
* `source_negate` - True if negate is set for source.
* `time` - List of time objects. For example: "Weekend", "Off-Work", "Every-Day".
* `track` - Track Settings. Track Settings blocks are documented below.
* `user_check` - User check settings. User check settings blocks are documented below.
* `vpn` - VPN community identified by name or UID or "Any" or "All_GwToGw".
* `comments` - Comments string.
* `fields_with_uid_identifier` - (Optional) List of resource fields that will use object UIDs as object identifiers. Default is object name.

`action_settings` supports the following:

* `enable_identity_captive_portal`
* `limit`

`custom_fields` supports the following:

* `field_1` - First custom field.
* `field_2` - Second custom field.
* `field_3` - Third custom field.

`track` supports the following:

* `accounting` - Turns accounting for track on and off.
* `alert` - Type of alert for the track.
* `enable_firewall_session` - Determine whether to generate session log to firewall only connections.
* `per_connection` - Determines whether to perform the log per connection.
* `per_session` - Determines whether to perform the log per session.
* `type` - \"Log\", \"Extended Log\", \"Detailed Log\", \"None\".

`user_check` supports the following:

* `confirm`
* `custom_frequency` - Custom Frequency blocks are documented below.
* `frequency`
* `interaction`

`custom_frequency` supports the following:

* `every`
* `unit`











