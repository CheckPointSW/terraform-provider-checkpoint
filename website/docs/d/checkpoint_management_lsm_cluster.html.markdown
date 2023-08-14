---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_cluster"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsm-cluster"
description: |-
Use this data source to get information on an existing Check Point Lsm Cluster.
---

# Data Source: checkpoint_management_lsm_cluster

Use this data source to get information on an existing Check Point Lsm Cluster.

## Example Usage


```hcl
data "checkpoint_management_lsm_cluster" "data_cluster"{
 uid = "${checkpoint_management_lsm_cluster.cluster1.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object uid.
* `main_ip_address` -  Main IP address.
* `security_profile` -  LSM profile.
* `os_name` - Device platform operating system.
* `dynamic_objects` -  Dynamic Objects.dynamic_objects blocks are documented below.
* `interfaces` - Interfaces.interfaces blocks are documented below.
* `members` - Cluster members.members blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `topology` -  Topology.topology blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 




`dynamic_objects` supports the following:

* `name` - Object name. Must be unique in the domain. 
* `resolved_ip_addresses` -  Single IP-address or a range of addresses.resolved_ip_addresses blocks are documented below.


`interfaces` supports the following:

* `name` -  Interface name. 
* `ip-address-override` Cluster IP address override.
* `member_network_override` -  Member network override. Net mask is defined by the attached LSM profile. 


`members` supports the following:

* `name` - Member Name. Consists of the member name in the LSM profile and the name or prefix or suffix of the cluster. 
* `device_id` -  Device ID. 
* `provisioning_settings` -  Provisioning settings. This field is relevant just for SMB clusters.provisioning_settings blocks are documented below.
* `provisioning_state` -  Provisioning state. This field is relevant just for SMB clusters. By default the state is 'manual'- enable provisioning but not attach to profile.If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings. 
* `sic_name` - Secure Internal Communication name.
* `sic_state`- Secure Internal Communication state.



`topology` supports the following:

* `manual_vpn_domain` -  A list of IP-addresses ranges, defined the VPN community network.
This field is relevant only when 'manual' option of vpn-domain is checked.manual_vpn_domain blocks are documented below.
* `vpn_domain` -  VPN Domain type.
 'external-interfaces-only' is relevnt only for Gaia devices.
'hide-behind-gateway-external-ip-address' is relevant only for SMB devices. 


`resolved_ip_addresses` supports the following:

* `ipv4_address` -  IPv4 Address. 
* `ipv4_address_range` -  IPv4 Address range.ipv4_address_range blocks are documented below.


`provisioning_settings` supports the following:

* `provisioning_profile` -  Provisioning profile. 

`manual_vpn_domain` supports the following:

* `comments` -  Comments string. 
* `from_ipv4_address` -  First IPv4 address of the IP address range. 
* `to_ipv4_address` -  Last IPv4 address of the IP address range. 


`ipv4_address_range` supports the following:

* `from_ipv4_address` -  First IPv4 address of the IP address range. 
* `to_ipv4_address` -  Last IPv4 address of the IP address range. 
