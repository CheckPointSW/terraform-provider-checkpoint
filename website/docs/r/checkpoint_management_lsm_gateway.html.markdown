---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsm-gateway"
description: |-
This resource allows you to execute Check Point Lsm Gateway.
---

# Resource: checkpoint_management_lsm_gateway

This resource allows you to execute Check Point Lsm Gateway.

## Example Usage


```hcl
resource "checkpoint_management_lsm_gateway" "lsm_gw" {
  name = "lsm_gateway"
  security_profile = "lsm_profile"
  provisioning_state = "using-profile"
  provisioning_settings = {
  "provisioning_profile" = "my_proviosioning_profile"
  }
  topology {
    vpn_domain = "manual"
    manual_vpn_domain {
      comments          = "domain1"
      from_ipv4_address = "192.168.10.0"
      to_ipv4_address   = "192.168.10.255"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `security_profile` - (Required) LSM profile.
* `device_id` - (Optional) Device ID. 
* `dynamic_objects` - (Optional) Dynamic Objects.dynamic_objects blocks are documented below.
* `provisioning_settings` - (Optional) Provisioning settings.provisioning_settings blocks are documented below.
* `provisioning_state` - (Optional) Provisioning state. By default the state is 'manual'- enable provisioning but not attach to profile.
If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.
* `sic` - (Optional) Secure Internal Communication.sic blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) Topology.topology blocks are documented below.
* `version` - (Optional) Device platform version.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`dynamic_objects` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `resolved_ip_addresses` - (Optional) Single IP-address or a range of addresses.resolved_ip_addresses blocks are documented below.


`provisioning_settings` supports the following:

* `provisioning_profile` - (Optional) Provisioning profile. 


`sic` supports the following:

* `one_time_password` - (Optional) One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time. 
* `ip_address` - (Optional) IP address. When IP address is provided- initiate trusted communication immediately using this IP address.

`topology` supports the following:

* `manual_vpn_domain` - (Optional) A list of IP-addresses ranges, defined the VPN community network.
This field is relevant only when 'manual' option of vpn-domain is checked.manual_vpn_domain blocks are documented below.
* `vpn_domain` - (Optional) VPN Domain type.
 'external-interfaces-only' is relevnt only for Gaia devices.
'hide-behind-gateway-external-ip-address' is relevant only for SMB devices. 


`resolved_ip_addresses` supports the following:

* `ipv4_address` - (Optional) IPv4 Address. 
* `ipv4_address_range` - (Optional) IPv4 Address range.ipv4_address_range blocks are documented below.


`manual_vpn_domain` supports the following:

* `comments` - (Optional) Comments string. 
* `from_ipv4_address` - (Optional) First IPv4 address of the IP address range. 
* `to_ipv4_address` - (Optional) Last IPv4 address of the IP address range. 


`ipv4_address_range` supports the following:

* `from_ipv4_address` - (Optional) First IPv4 address of the IP address range. 
* `to_ipv4_address` - (Optional) Last IPv4 address of the IP address range. 
