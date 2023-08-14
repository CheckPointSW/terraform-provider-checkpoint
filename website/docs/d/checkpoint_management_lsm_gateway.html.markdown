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
