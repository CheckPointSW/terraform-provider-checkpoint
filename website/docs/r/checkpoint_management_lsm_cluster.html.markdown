---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_cluster"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsm-cluster"
description: |-
This resource allows you to execute Check Point Lsm Cluster.
---

# Resource: checkpoint_management_lsm_cluster

This resource allows you to execute Check Point Lsm Cluster.

## Example Usage

```hcl
resource "checkpoint_management_lsm_cluster" "cluster" {
 name = "Gaia_"
 main_ip_address = "192.168.8.0"
 security_profile = "my_security_profile"
 interfaces {
  name= "et0"
  ip_address_override = "170.150.0.1"
  member_network_override = "192.168.8.0"

 }
 members {
  name = "Gaia_mem1"
  provisioning_state = "off"
  provisioning_settings {
   provisioning_profile = "No Provisioning Profile"
  }
 }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `main_ip_address` - (Required) Main IP address.
* `security_profile` - (Required) LSM profile.
* `dynamic_objects` - (Optional) Dynamic Objects.dynamic_objects blocks are documented below.
* `interfaces` - (Optional) Interfaces.interfaces blocks are documented below.
* `members` - (Optional) Cluster members.members blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `topology` - (Optional) Topology.topology blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.



`dynamic_objects` supports the following:

* `name` - (Optional) Object name. Must be unique in the domain. 
* `resolved_ip_addresses` - (Optional) Single IP-address or a range of addresses.resolved_ip_addresses blocks are documented below.


`interfaces` supports the following:

* `name` - (Optional) Interface name. 
* `member_network_override` - (Optional) Member network override. Net mask is defined by the attached LSM profile. 


`members` supports the following:

* `name` - (Optional) Member Name. Consists of the member name in the LSM profile and the name or prefix or suffix of the cluster. 
* `device_id` - (Optional) Device ID. 
* `provisioning_settings` - (Optional) Provisioning settings. This field is relevant just for SMB clusters.provisioning_settings blocks are documented below.
* `provisioning_state` - (Optional) Provisioning state. This field is relevant just for SMB clusters. By default the state is 'manual'- enable provisioning but not attach to profile.If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings. 
* `sic` - (Optional) Secure Internal Communication.sic blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`topology` supports the following:

* `manual_vpn_domain` - (Optional) A list of IP-addresses ranges, defined the VPN community network.
This field is relevant only when 'manual' option of vpn-domain is checked.manual_vpn_domain blocks are documented below.
* `vpn_domain` - (Optional) VPN Domain type.
 'external-interfaces-only' is relevnt only for Gaia devices.
'hide-behind-gateway-external-ip-address' is relevant only for SMB devices. 


`resolved_ip_addresses` supports the following:

* `ipv4_address` - (Optional) IPv4 Address. 
* `ipv4_address_range` - (Optional) IPv4 Address range.ipv4_address_range blocks are documented below.


`provisioning_settings` supports the following:

* `provisioning_profile` - (Optional) Provisioning profile. 


`sic` supports the following:

* `one_time_password` - (Optional) One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time. 
* `ip_address` - (Optional) IP address. When IP address is provided- initiate trusted communication immediately using this IP address.


`manual_vpn_domain` supports the following:

* `comments` - (Optional) Comments string. 
* `from_ipv4_address` - (Optional) First IPv4 address of the IP address range. 
* `to_ipv4_address` - (Optional) Last IPv4 address of the IP address range. 


`ipv4_address_range` supports the following:

* `from_ipv4_address` - (Optional) First IPv4 address of the IP address range. 
* `to_ipv4_address` - (Optional) Last IPv4 address of the IP address range. 
