---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsm-gateway"
description: |-
Use this data source to get information on an existing Check Point Lsm Gateway.
---

# Data Source: checkpoint_management_lsm_gateway

Use this data source to get information on an existing Check Point Lsm Gateway.

## Example Usage


```hcl
data "checkpoint_management_lsm_gateway" "data_lsm" {
 name = "${checkpoint_management_lsm_gateway.lsm_gw.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object uid.
* `security_profile` - LSM profile.
* `device_id` - Device ID. 
* `dynamic_objects` -  Dynamic Objects.dynamic_objects blocks are documented below.
* `provisioning_settings` - Provisioning settings.provisioning_settings blocks are documented below.
* `provisioning_state` - Provisioning state. By default the state is 'manual'- enable provisioning but not attach to profile.
If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.
* `sic_name` - Secure Internal Communication name.
* `sic_state`- Secure Internal Communication state.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `topology` -  Topology. topology blocks are documented below.
* `version` - Device platform version.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 
* `ignore_warnings` -  Apply changes ignoring warnings. 
* `ignore_errors` -  Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `gateway_status` - The current status of the Gateway. Shown only when the 'show-statuses' parameter is set to 'true'.
* `last_applied_provisioning_settings_time` - The last time when the Provisioning Settings were changed. Shown only when the 'show-statuses' parameter is set to 'true'.last_applied_provisioning_settings_time blocks are documented below.
* `last_policy_fetch_time` - The last time when the Security Policy was fetched. Shown only when the 'show-statuses' parameter is set to 'true'.last_policy_fetch_time blocks are documented below.
* `last_provisioning_settings_sync_time` - The last time of Provisioning Settings synchronization with the Gateway. Shown only when the 'show-statuses' parameter is set to 'true'.last_provisioning_settings_sync_time blocks are documented below.
* `policy_status` - The current status of the Security Policy. Shown only when the 'show-statuses' parameter is set to 'true'.
* `provisioning_settings_status` - The current status of the Provisioning Settings. Shown only when the 'show-statuses' parameter is set to 'true'.


`dynamic_objects` supports the following:

* `name` - Object name. Must be unique in the domain. 
* `uid`- Dynamic Object UID.
* `comments` - Comments.
* `resolved_ip_addresses` - Single IP-address or a range of addresses.resolved_ip_addresses blocks are documented below.


`provisioning_settings` supports the following:

* `provisioning_profile` -  Provisioning profile. 

`topology` supports the following:

* `manual_vpn_domain` -  A list of IP-addresses ranges, defined the VPN community network.
This field is relevant only when 'manual' option of vpn-domain is checked.manual_vpn_domain blocks are documented below.
* `vpn_domain` -  VPN Domain type.
 'external-interfaces-only' is relevnt only for Gaia devices.
'hide-behind-gateway-external-ip-address' is relevant only for SMB devices. 


`resolved_ip_addresses` supports the following:

* `ipv4_address` -  IPv4 Address. 
* `ipv4_address_range` -  IPv4 Address range.ipv4_address_range blocks are documented below.


`manual_vpn_domain` supports the following:

* `comments` -  Comments string. 
* `from_ipv4_address` - First IPv4 address of the IP address range. 
* `to_ipv4_address` -  Last IPv4 address of the IP address range. 


`ipv4_address_range` supports the following:

* `from_ipv4_address` - (Optional) First IPv4 address of the IP address range. 
* `to_ipv4_address` - (Optional) Last IPv4 address of the IP address range. 


`last_applied_provisioning_settings_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`last_policy_fetch_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`last_provisioning_settings_sync_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
