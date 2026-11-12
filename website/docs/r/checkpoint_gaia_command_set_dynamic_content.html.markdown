---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_dynamic_content"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-dynamic-content"
description: |-
This resource allows you to execute Check Point Set Dynamic Content.
---

# checkpoint_gaia_command_set_dynamic_content

This resource allows you to execute Check Point Set Dynamic Content.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_dynamic_content" "example" {
  comments = "testing the API"
  dry_run  = true

  objects {
    hosts {
      name       = "test-host"
      ip_address = "1.2.3.4"
    }
  }

  access_layers_content {
    name      = "dynamic_layer"
    operation = "replace"
    rulebase {
      name        = "rule1"
      action      = "Accept"
      source      = ["any"]
      destination = ["any"]
      service     = ["any"]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `objects` - (Required) List of objects to create. objects blocks are documented below.
* `access_layers_content` - (Required) List of layers to apply. Supported layers : Layers created using this API, externally referenced layers (layers marked as 'dynamic layers' in the SmartConsole) access_layers_content blocks are documented below.
* `comments` - (Optional) Comments for this operation. 
* `tags` - (Optional) List of tags for this operation. tags blocks are documented below.
* `custom_fields` - (Optional) List of custom fields for this operation. custom_fields blocks are documented below.
* `dry_run` - (Optional) Perform validation without applying changes. 
* `referenced_objects` - (Optional) List of object names defined externally ("internet" , "any" , "_GW_" are already referenced). referenced_objects blocks are documented below.
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 


`objects` supports the following:

* `hosts` - (Optional)  hosts blocks are documented below.
* `networks` - (Optional)  networks blocks are documented below.
* `network_groups` - (Optional)  network_groups blocks are documented below.
* `services_tcp` - (Optional)  services_tcp blocks are documented below.
* `services_udp` - (Optional)  services_udp blocks are documented below.
* `services_other` - (Optional)  services_other blocks are documented below.
* `service_groups` - (Optional)  service_groups blocks are documented below.
* `application_sites` - (Optional)  application_sites blocks are documented below.
* `application_site_categories` - (Optional)  application_site_categories blocks are documented below.
* `application_site_groups` - (Optional)  application_site_groups blocks are documented below.
* `dynamic_objects` - (Optional)  dynamic_objects blocks are documented below.
* `dns_domains` - (Optional)  dns_domains blocks are documented below.
* `address_ranges` - (Optional)  address_ranges blocks are documented below.
* `groups_with_exclusion` - (Optional)  groups_with_exclusion blocks are documented below.
* `wildcards` - (Optional)  wildcards blocks are documented below.
* `identity_tags` - (Optional)  identity_tags blocks are documented below.
* `access_roles` - (Optional)  access_roles blocks are documented below.
* `access_layers` - (Optional)  access_layers blocks are documented below.


`access_layers_content` supports the following:

* `name` - (Optional) Layer name. 
* `rulebase` - (Optional) Rules of the layer. rulebase blocks are documented below.
* `operation` - (Optional) Layer operation. 


`custom_fields` supports the following:

* `field_1` - (Optional) First Custom Field 
* `field_2` - (Optional) Second Custom Field 
* `field_3` - (Optional) Third Custom Field 


`referenced_objects` supports the following:

* `application_sites` - (Optional) List of Application/Site objects as configured in SmartConsole and identified by the name. application_sites blocks are documented below.
* `application_site_categories` - (Optional) List of Application/Site Category objects as configured in SmartConsole and identified by the name. application_site_categories blocks are documented below.
* `services_tcp` - (Optional) List of TCP service objects as configured in SmartConsole and identified by the name. services_tcp blocks are documented below.
* `services_udp` - (Optional) List of UDP service objects as configured in SmartConsole and identified by the name. services_udp blocks are documented below.
* `services_icmp` - (Optional) List of ICMP service objects as configured in SmartConsole and identified by the name. services_icmp blocks are documented below.
* `updatable_objects` - (Optional) List of Updatable objects as configured in SmartConsole and identified by the name. updatable_objects blocks are documented below.
* `access_layers` - (Optional) List of Policy Layers in the Access Control Policy as configured in SmartConsole and identified by the name. access_layers blocks are documented below.


`hosts` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ip_address` - (Optional) IPv4 or IPv6 address. If both addresses are required then use the 'ipv4-address' and 'ipv6-address' fields explicitly. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv6_address` - (Optional) IPv6 address. 


`networks` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `subnet` - (Optional) IPv4 or IPv6 network address. If both addresses are required then use the 'subnet4' and 'subnet6' fields explicitly. 
* `subnet4` - (Optional) IPv4 address. 
* `subnet6` - (Optional) IPv6 address. 
* `subnet_mask` - (Optional) IPv4 network mask. 
* `mask_length` - (Optional) IPv4 or IPv6 network mask length.                     If both masks are required then use the 'mask-length4' and 'mask-length6' fields explicitly.                    Instead of the IPv4 mask length it is possible to specify the IPv4 mask itself in 'subnet-mask' field. 
* `mask_length4` - (Optional) IPv4 network mask length. 
* `mask_length6` - (Optional) IPv6 network mask length. 
* `broadcast` - (Optional) Allow broadcast address inclusion. 


`network_groups` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `members` - (Optional) Collection of Network objects identified by the name. members blocks are documented below.


`services_tcp` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `port` - (Optional) The number of the port used to provide this service.                    To specify a port range, place a hyphen between the lowest and highest port numbers,                        for example 44-55. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been                   installed even if they are not allowed under the new policy.                       This overrides the settings on the Connection Persistence page of the Security Gateway object.                           If you change this property, the change will not affect open connections,                               but only future connections. 
* `session_timeout` - (Optional) Time (in seconds) before the session times out. 
* `sync_connections_on_cluster` - (Optional) Enables the state synchronization in a ClusterXL or OPSEC-certified cluster. 
* `source_port` - (Optional) Port number for the client-side service.                     If specified, only packets with these source port numbers will be accepted,                     dropped, or rejected during packet inspection.                     Otherwise, the packets are not matched to this service. 
* `use_delayed_sync` - (Optional) Enable this option to delay notifying the Security Gateway about a connection,                    so that the connection will only be synchronized if it still exists x seconds after the connection is initiated.                     This feature uses SecureXL that is enabled by default. 
* `delayed_sync_value` - (Optional) Specify the delay (in seconds) when the synchronization will start after connection initiation. Relevant only if "use-delayed-sync" was set to "true". 


`services_udp` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `port` - (Optional) The number of the port used to provide this service.         To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been                   installed even if they are not allowed under the new policy.                       This overrides the settings on the Connection Persistence page of the Security Gateway object.                           If you change this property, the change will not affect open connections,                               but only future connections. 
* `session_timeout` - (Optional) Time (in seconds) before the session times out. 
* `sync_connections_on_cluster` - (Optional) Enables onthe state synchronization in a ClusterXL or OPSEC-certified cluster. 
* `source_port` - (Optional) Port number for the client-side service.                     If specified, only packets with these source port numbers will be accepted,                     dropped, or rejected during packet inspection.                     Otherwise, the packets are not matched to this service. 
* `accept_replies` - (Optional) Specifies whether to accept UDP replies for this service. 


`services_other` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ip_protocol` - (Optional) IP protocol number. 
* `keep_connections_open_after_policy_installation` - (Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy.                     This overrides the settings on the Connection Persistence page in the Security Gateway object.                     If you change this property, the change will not affect open connections, but only future connections. 
* `session_timeout` - (Optional) Time (in seconds) before the session times out. 
* `sync_connections_on_cluster` - (Optional) Enables the state synchronization in a ClusterXL or OPSEC-certified cluster. 
* `accept_replies` - (Optional) Specifies whether to accept replies for this service. 


`service_groups` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `members` - (Optional) Collection of Service objects identified by the name. members blocks are documented below.


`application_sites` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `clone_of` - (Optional) Name of existing Application/Category to be cloned. 
* `services` - (Optional) Collection of Service objects identified by the name. You can specify any service with the value 'any' or 'Any'. services blocks are documented below.
* `negate` - (Optional) Specifies whether this object is negated. 


`application_site_categories` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `clone_of` - (Optional) Name of existing Application/Category to be cloned. 
* `services` - (Optional) Collection of Service objects identified by the name. You can specify any service with the value 'any' or 'Any'. services blocks are documented below.
* `negate` - (Optional) Specifies whether this object is negated. 


`application_site_groups` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `members` - (Optional) Collection of Application/Site objects identified by the name. members blocks are documented below.


`dynamic_objects` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 


`dns_domains` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `is_sub_domain` - (Optional) Specifies whether to match sub-domains in addition to the domain itself. 


`address_ranges` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ip_address_first` - (Optional) First IP address in the range.                     If both IPv4 and IPv6 address ranges are required,                         then use the 'ipv4-address-first' and the 'ipv6-address-first' fields instead. 
* `ipv4_address_first` - (Optional) First IPv4 address in the range. 
* `ipv6_address_first` - (Optional) First IPv6 address in the range. 
* `ip_address_last` - (Optional) Last IP address in the range.                     If both IPv4 and IPv6 address ranges are required,                         then use the 'ipv4-address-last' and the 'ipv6-address-last' fields instead. 
* `ipv4_address_last` - (Optional) Last IPv4 address in the range. 
* `ipv6_address_last` - (Optional) Last IPv6 address in the range. 


`groups_with_exclusion` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `include` - (Optional) Name of an object which the group includes. 
* `except` - (Optional) Name of an object which the group excludes. 


`wildcards` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv4_mask_wildcard` - (Optional) IPv4 mask wildcard. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv6_mask_wildcard` - (Optional) IPv6 mask wildcard. 


`identity_tags` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `external_identifier` - (Optional) External identifier. For example: Cisco ISE security group tag. 


`access_roles` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `ip_spoofing_protection` - (Optional) Enforce IP spoofing protection. 
* `networks` - (Optional) Collection of Network objects identified by the name that can access the system.                         Level of details in the output corresponds to the number of details for search.                             You can specify any Access Role with the value 'any' or 'Any'. networks blocks are documented below.
* `users` - (Optional) Users that can access the system.                     Level of details in the output corresponds to the number of details for search.                     Valid options: 'any', 'all identified'. users blocks are documented below.
* `machines` - (Optional) Machines that can access the system.                         Level of details in the output corresponds to the number of details for search.                         Valid options: 'any', 'all identified'. machines blocks are documented below.


`access_layers` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `firewall` - (Optional) Whether to enable the Firewall blade on the layer. 
* `applications_and_url_filtering` - (Optional) Whether to enable the Application Control & URL Filtering blades on the layer. 
* `content_awareness` - (Optional) Whether to enable the Content Awareness blade on the layer. 
* `detect_using_x_forward_for` - (Optional) Whether to use the 'X-Forward-For' HTTP header,                     which is added by the proxy server to keep track of the original source IP address. 
* `mobile_access` - (Optional) Whether to enable the Mobile Access blade on the layer. 
* `implicit_cleanup_action` - (Optional) Specifies the default 'catch-all' action for traffic that does not match any explicit or implied rules in the layer. 


`rulebase` supports the following:

* `name` - (Optional) Rule name, Must be unique in the layer. 
* `action` - (Optional) Action. Valid options: "Accept", "Drop", "Ask", "Drop with Block message", "Inform", "Reject", "Apply Layer". 
* `action_settings` - (Optional) Action settings. action_settings blocks are documented below.
* `inline_layer` - (Optional) Inline Layer identified by the name.  Relevant only if "the action" was set to "Apply Layer". 
* `track` - (Optional) Track Settings. track blocks are documented below.
* `source` - (Optional) Collection of network objects identified by their name or 'any'. source blocks are documented below.
* `source_negate` - (Optional) Specifies whether to negate the source. 
* `destination` - (Optional) Collection of network objects identified by their name or 'any'. destination blocks are documented below.
* `destination_negate` - (Optional) Specifies whether to negate the destination. 
* `user_check` - (Optional) UserCheck settings. user_check blocks are documented below.
* `service` - (Optional) Collection of Service and Application objects identified by the name. You can specify any object with the value 'any' or 'Any'. service blocks are documented below.
* `service_negate` - (Optional) Specifies whether to negate this service. 


`users` supports the following:

* `source` - (Optional) Active Directory name or 'Identity Tag' or 'Guests'. 
* `selection` - (Optional) Distinguished Name (DN) or Identity Tag name or 'Unauthenticated Guests'. 
* `ad_entity_type` - (Optional) Active directory entity type. 


`machines` supports the following:

* `source` - (Optional) Active Directory name or 'Identity Tag'. 
* `selection` - (Optional) Distinguished Name (DN) or Identity Tag name. 
* `ad_entity_type` - (Optional) Active directory entity type. 


`action_settings` supports the following:

* `enable_identity_captive_portal` - (Optional) Redirect HTTP traffic to an authentication (Captive Portal). After the user is authenticated, new connections from this source are inspected without requiring authentication. 


`track` supports the following:

* `type` - (Optional) Track type. Valid options: "Log", "Extended Log", "Detailed Log", "None". 
* `accounting` - (Optional) Turns the Accounting on and off. 
* `alert` - (Optional) Type of alert for the track. Valid options: "None", "Alert", "Snmp", "Mail", "User Alert 1", "User Alert 2", "User Alert 3". 
* `enable_firewall_session` - (Optional) Specifies whether to generate a session log for connections that are inspected only by the Firewall blade. 
* `per_connection` - (Optional) Specifies whether to generate a log for each connection. If set to 'true', may decrease the Security Gateway performance because of the number of generated logs. 
* `per_session` - (Optional) Specifies whether to generate a log for each session. If set to 'true', may decrease the Security Gateway performance because of the number of generated logs. 


`user_check` supports the following:

* `confirm` - (Optional) Valid options: "per rule", "per category", "per application/site", "per data type" . 
* `custom_frequency` - (Optional) Configure how often the user sees the configured message when the action is "ask", "inform", or "block". Relevant only if "frequency" was set to "custom frequency". custom_frequency blocks are documented below.
* `frequency` - (Optional) Configure how often the user sees the configured message when the action is "ask", "inform", or "block". Valid options: "once a day", "once a week", "once a month", "custom frequency" . 
* `interaction` - (Optional) Add the relevant interaction text. Need to be relevant to the rule action. 


`custom_frequency` supports the following:

* `every` - (Optional) Valid values: 1 - 999 . 
* `unit` - (Optional) Valid options: hours, days, weeks, months . 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

