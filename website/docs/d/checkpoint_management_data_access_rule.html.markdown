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
  position { top = "top" }
  source = ["Any"]
  destination = ["Any"]
  service = ["Any"]
  content = ["Any"]
  time = ["Any"]
  install_on = ["Policy Targets"]
  track {
    type = "Log"
    accounting = false
    alert = "none"
    enable_firewall_session = false
    per_connection = true
    per_session = false
  }
  custom_fields {}
  vpn = "Any"
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
* `vpn` - VPN community identified by name or "Any" or "All_GwToGw".
* `vpn_communities` - Collection of VPN communities identified by name.
* `vpn_directional` - Collection of VPN directional. VPN directional block documented below.
* `comments` - Comments string.
* `fields_with_uid_identifier` - (Optional) List of resource fields that will use object UIDs as object identifiers. Default is object name.
* `destination_ranges` - Displays the destination as ranges of IP addresses, in case show-as-ranges is set to true.<br />In this case, 'destination' and 'destination-negate' p...destination_ranges blocks are documented below.
* `expiration_settings` - Displays the expiration date settings.expiration_settings blocks are documented below.
* `hits` - Hits count object.hits blocks are documented below.
* `service_ranges` - Displays the services and applications as ranges of port numbers, in case show-as-ranges is set to true.<br />In this case, 'service' and 'service-neg...service_ranges blocks are documented below.
* `service_resource` - Resource of the service.
* `source_ranges` - Displays the source as ranges of IP addresses, in case show-as-ranges is set to true.<br />In this case, 'source' and 'source-negate' parameters are o...source_ranges blocks are documented below.

`action_settings` supports the following:

* `enable_identity_captive_portal`
* `limit`
* `client_auth_settings` - N/A.client_auth_settings blocks are documented below.
* `user_auth_settings` - N/A.user_auth_settings blocks are documented below.

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

`vpn_directional` supports the following:

* `from` - From VPN community.
* `to` - To VPN community.


`client_auth_settings` supports the following:

* `destination` - How destination hosts are matched against the user database.
* `require_desktop_config_verification` - When true, the rule only applies if the client's desktop security policy configuration has been verified by the gateway.
* `sessions_limit` - Maximum number of concurrent sessions allowed per user. Only used when unlimited-sessions is false. Must be 1 or greater.
* `sign_on_method` - Mechanism used to authenticate the client.
* `sign_on_type` - Determines whether a standard or specific sign-on policy is applied. Allowed values: standard, specific.
* `source` - How source users are matched against the user database.
* `timeout` - Controls how long an authenticated session remains valid.timeout blocks are documented below.
* `tracking` - Action to take in the log when a user successfully authenticates. Allowed values: none, log, alert.
* `unlimited_sessions` - When true, a single user may have any number of concurrent authenticated sessions. When false, the number of concurrent sessions is capped at sessions...


`timeout` supports the following:

* `enable` - When true, authenticated sessions expire after the duration defined by minutes. When false, authenticated sessions never expire.
* `minutes` - Number of minutes before an authenticated session expires. Only used when enable is true. Must be 0 or greater.
* `refreshable` - When true, the session timeout timer resets each time the authenticated user performs a network action.


`user_auth_settings` supports the following:

* `allowed_http_servers` - Restricts which HTTP servers authenticated users may access.
* `destination` - How destination hosts are matched against the user database.
* `source` - How source users are matched against the user database.


`destination_ranges` supports the following:

* `excluded_others` - Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...
* `ipv4` - Range of IPv4 addresses that match in the given rule.ipv4 blocks are documented below.
* `ipv6` - Range of IPv6 addresses that match in the given rule.ipv6 blocks are documented below.
* `others` - Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...


`ipv4` supports the following:

* `end` - N/A.
* `start` - N/A.


`ipv6` supports the following:

* `end` - N/A.
* `start` - N/A.


`expiration_settings` supports the following:

* `expiration_date` - Expiration date.expiration_date blocks are documented below.
* `expired` - Expired rule.
* `has_expiration_date` - Rule has expiration date.


`expiration_date` supports the following:

* `iso_8601` - N/A.
* `posix` - N/A.


`hits` supports the following:

* `first_date` - N/A.first_date blocks are documented below.
* `last_date` - N/A.last_date blocks are documented below.
* `level` - N/A.
* `percentage` - N/A.
* `value` - N/A.


`first_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`last_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`service_ranges` supports the following:

* `excluded_others` - Objects which are not represented as port numbers and are negated in the given rule - for example if negate is set for the service of this rule. The d...
* `others` - Objects which are not represented as port numbers and match the given rule. The details-level parameter of the request determines whether they are dis...
* `tcp` - Range of TCP ports that match in the given rule.tcp blocks are documented below.
* `udp` - Range of UDP ports that match in the given rule.udp blocks are documented below.


`tcp` supports the following:

* `end` - N/A.
* `start` - N/A.


`udp` supports the following:

* `end` - N/A.
* `start` - N/A.


`source_ranges` supports the following:

* `excluded_others` - Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...
* `ipv4` - Range of IPv4 addresses that match in the given rule.ipv4 blocks are documented below.
* `ipv6` - Range of IPv6 addresses that match in the given rule.ipv6 blocks are documented below.
* `others` - Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...


`ipv4` supports the following:

* `end` - N/A.
* `start` - N/A.


`ipv6` supports the following:

* `end` - N/A.
* `start` - N/A.
