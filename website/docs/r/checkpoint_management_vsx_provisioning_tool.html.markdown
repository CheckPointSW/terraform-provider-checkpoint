---
layout: "checkpoint"
page_title: "checkpoint_management_vsx_provisioning_tool"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-vsx-provisioning-tool"
description: |-
This resource allows you to execute Check Point VSX Provisioning Tool.
---

# checkpoint_management_vsx_provisioning_tool

This resource allows you to execute Check Point VSX Provisioning Tool.

## Example Usage


```hcl
resource "checkpoint_management_vsx_provisioning_tool" "example" {
  operation = "add-vsx-cluster"
  add_vsx_cluster_params  {
    vsx_name = "VSX_CLUSTER"
    cluster_type = "vsls"
    ipv4_address = "192.168.0.0"
    version = "R82"
    sync_if_name = "eth3"
    sync_netmask = "255.255.255.0"
    rule_ping = "enable"
    rule_drop = "enable"
    members  {
      name = "mem1"
      ipv4_address = "4.4.4.4"
      sync_ip = "192.168.1.1"
      sic_otp = "sicotp123"
    }
    members {
      name = "mem2"
      ipv4_address = "8.4.4.4"
      sync_ip = "192.168.2.2"
      sic_otp = "sicotp456"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `operation` - (Required) The name of the provisioning operation to run. Each operation has its own specific parameters.<br> The available operations are:<ul><li><i>add-vsx-gateway</i> - Adds a new VSX gateway</li><li><i>add-vsx-cluster</i> - Adds a new VSX cluster*</li><li><i>add-vsx-cluster-member</i> - Adds a new VSX cluster member*</li><li><i>add-vd</i> - Adds a new Virtual Device (VS/VSB/VSW/VR) to a VSX gateway or VSX cluster</li><li><i>add-vd-interface</i> - Adds a new virtual interface to a Virtual Device</li><li><i>add-physical-interface</i> - Adds a physical interface to a VSX gateway or VSX cluster</li><li><i>add-route</i> - Adds a route to a Virtual Device</li><li><i>attach-bridge</i> - Attaches a bridge interface to a Virtual System</li><li><i>remove-vsx</i> - Removes a VSX gateway or VSX cluster</li><li><i>remove-vd</i> - Removes a Virtual Device</li><li><i>remove-vd-interface</i> - Removes an interface from a Virtual Device</li><li><i>remove-physical-interface</i> - Removes a physical interface from a VSX gateway or VSX cluster</li><li><i>remove-route</i> - Removes a route from a Virtual Device</li><li><i>set-vd</i> - Modifies a Virtual Device</li><li><i>set-vd-interface</i> - Modifies an interface on a Virtual Device</li><li><i>set-physical-interface</i> - Modifies a physical interface on a VSX cluster or VSX gateway</li></ul><br> * When adding a VSX Cluster, you must also add at least 2 cluster members<br> * Adding cluster members is only allowed when adding a new VSX cluster<br> * To add members to an existing cluster, use vsx-run-operation. 
* `add_physical_interface_params` - (Optional) Parameters for the operation to add a physical interface to a VSX gateway or VSX Cluster.add_physical_interface_params blocks are documented below.
* `add_route_params` - (Optional) Parameters for the operation to add a route to a Virtual System or Virtual Router.add_route_params blocks are documented below.
* `add_vd_interface_params` - (Optional) Parameters for the operation to add a new interface to a Virtual Device.add_vd_interface_params blocks are documented below.
* `add_vd_params` - (Optional) Parameters for the operation to add a new Virtual Device (VS/VSB/VSW/VR).add_vd_params blocks are documented below.
* `add_vsx_cluster_params` - (Optional) Parameters for the operation to add a new VSX Cluster.add_vsx_cluster_params blocks are documented below.
* `add_vsx_gateway_params` - (Optional) Parameters for the operation to add a new VSX Gateway.add_vsx_gateway_params blocks are documented below.
* `attach_bridge_params` - (Optional) Parameters for the operation to attach a new bridge interface to a Virtual System.attach_bridge_params blocks are documented below.
* `remove_physical_interface_params` - (Optional) Parameters for the operation to remove a physical interface from a VSX (Gateway or Cluster).remove_physical_interface_params blocks are documented below.
* `remove_route_params` - (Optional) Parameters for the operation to remove a route from a Virtual System or Virtual Router.remove_route_params blocks are documented below.
* `remove_vd_interface_params` - (Optional) Parameters for the operation to remove a logical interface from a Virtual Device.remove_vd_interface_params blocks are documented below.
* `remove_vd_params` - (Optional) Parameters for the operation to remove a Virtual Device.remove_vd_params blocks are documented below.
* `remove_vsx_params` - (Optional) Parameters for the operation to remove a VSX Gateway or VSX Cluster.remove_vsx_params blocks are documented below.
* `set_physical_interface_params` - (Optional) Parameters for the operation to change the configuration of a physical interface.set_physical_interface_params blocks are documented below.
* `set_vd_interface_params` - (Optional) Parameters for the operation to change the configuration of a logical interface.set_vd_interface_params blocks are documented below.
* `set_vd_params` - (Optional) Parameters for the operation to change the configuration of a Virtual Device.set_vd_params blocks are documented below.
* `task_id` - Operation Task UID.


`add_physical_interface_params` supports the following:

* `name` - (Required) Name of the interface. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 
* `vlan_trunk` - (Optional) True if this interface is a VLAN trunk. 


`add_route_params` supports the following:

* `destination` - (Required) Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6. 
* `next_hop` - (Optional) Next hop IP address. 
* `leads_to` - (Optional) Virtual Router for this route<br/>This VD must have an existing connection to the VR. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `netmask` - (Optional) Subnet mask for this route. 
* `prefix` - (Optional) CIDR prefix for this route. 
* `propagate` - (Optional) Propagate this route to adjacent virtual devices. 


`add_vd_interface_params` supports the following:

* `leads_to` - (Optional) Virtual Switch or Virtual Router for this interface. 
* `name` - (Optional) Name of the interface. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `anti_spoofing` - (Optional) The anti-spoofing enforcement setting of this interface. 
* `anti_spoofing_tracking` - (Optional) The anti-spoofing tracking setting of this interface. 
* `ipv4_address` - (Optional) IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `ipv4_netmask` - (Optional) IPv4 Subnet mask of this interface. 
* `ipv4_prefix` - (Optional) IPv4 CIDR prefix of this interface. 
* `ipv6_address` - (Optional) IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `ipv6_netmask` - (Optional) IPv6 Subnet mask of this interface. 
* `ipv6_prefix` - (Optional) IPv6 CIDR prefix of this interface. 
* `mtu` - (Optional) MTU of this interface. 
* `propagate` - (Optional) Propagate IPv4 route to adjacent virtual devices. 
* `propagate6` - (Optional) Propagate IPv6 route to adjacent virtual devices. 
* `specific_group` - (Optional) Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'. 
* `topology` - (Optional) Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS. 
* `vti_settings` - (Optional) VTI settings for this interface. This Virtual System must have VPN blade enabled.vti_settings blocks are documented below.


`add_vd_params` supports the following:

* `interfaces` - (Required) The list of interfaces for this new Virtual Device.<br/>Optional if this new VD is a Virtual Switch.interfaces blocks are documented below.
* `type` - (Required) Type of the Virtual Device <br><br>vs - Virtual Firewall<br>vr - Virtual Router<br>vsw - Virtual Switch<br>vsbm - Virtual Firewall in bridge mode. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 
* `calc_topology_auto` - (Optional) Calculate interface topology automatically based on routes.<br/>Relevant only for Virtual Systems.<br/>Do not use for virtual devices. 
* `ipv4_address` - (Optional) Main IPv4 Address.<br/>Required if this device is a Virtual System.<br/>Do not use for other virtual devices. 
* `ipv4_instances` - (Optional) Number of IPv4 instances for the Virtual System.<br/>Must be greater or equal to 1.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode. 
* `ipv6_address` - (Optional) Main IPv6 Address.<br/>Required if this device is a Virtual System.<br/>Do not use for other virtual devices. 
* `ipv6_instances` - (Optional) Number of IPv6 instances for the Virtual System.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode. 
* `routes` - (Optional) The list of routes for this new Virtual Device (VS or VR only).routes blocks are documented below.
* `vs_mtu` - (Optional) MTU of the Virtual System.<br/>Only relevant for Virtual Systems in bridge mode.<br/>Do not use for other virtual devices. 


`add_vsx_cluster_params` supports the following:

* `cluster_type` - (Required) Cluster type for the VSX Cluster Object.<br/>Starting in R81.10, only VSLS can be configured during cluster creation.<br/>To use High Availability ('ha'), first create the cluster as VSLS and then run vsx_util on the Management. 
* `ipv4_address` - (Optional) Main IPv4 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv6 Address is defined. 
* `ipv6_address` - (Optional) Main IPv6 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv4 Address is defined. 
* `members` - (Required) The list of cluster members for this new VSX Cluster. Minimum: 2.members blocks are documented below.
* `sync_if_name` - (Required) Sync interface name for the VSX Cluster. 
* `sync_netmask` - (Required) Sync interface netmask for the VSX Cluster. 
* `version` - (Required) Version of the VSX Gateway or Cluster object. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 
* `rule_drop` - (Optional) Add a default drop rule to the VSX Gateway or Cluster initial policy. 
* `rule_https` - (Optional) Add a rule to allow HTTPS traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ping` - (Optional) Add a rule to allow ping traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ping6` - (Optional) Add a rule to allow ping6 traffic to the VSX Gateway or Cluster initial policy. 
* `rule_snmp` - (Optional) Add a rule to allow SNMP traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ssh` - (Optional) Add a rule to allow SSH traffic to the VSX Gateway or Cluster initial policy. 


`add_vsx_gateway_params` supports the following:

* `ipv4_address` - (Optional) Main IPv4 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv6 Address is defined. 
* `ipv6_address` - (Optional) Main IPv6 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv4 Address is defined. 
* `sic_otp` - (Required) SIC one-time-password of the VSX Gateway or Cluster member.<br/>Password must be between 4-127 characters in length. 
* `version` - (Required) Version of the VSX Gateway or Cluster object. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 
* `rule_drop` - (Optional) Add a default drop rule to the VSX Gateway or Cluster initial policy. 
* `rule_https` - (Optional) Add a rule to allow HTTPS traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ping` - (Optional) Add a rule to allow ping traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ping6` - (Optional) Add a rule to allow ping6 traffic to the VSX Gateway or Cluster initial policy. 
* `rule_snmp` - (Optional) Add a rule to allow SNMP traffic to the VSX Gateway or Cluster initial policy. 
* `rule_ssh` - (Optional) Add a rule to allow SSH traffic to the VSX Gateway or Cluster initial policy. 


`attach_bridge_params` supports the following:

* `ifs1` - (Required) Name of the first interface for the bridge. 
* `ifs2` - (Required) Name of the second interface for the bridge. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 


`remove_physical_interface_params` supports the following:

* `name` - (Required) Name of the interface. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 


`remove_route_params` supports the following:

* `destination` - (Required) Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `netmask` - (Optional) Subnet mask for this route. 
* `prefix` - (Optional) CIDR prefix for this route. 


`remove_vd_interface_params` supports the following:

* `leads_to` - (Optional) Virtual Switch or Virtual Router for this interface. 
* `name` - (Optional) Name of the interface. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 


`remove_vd_params` supports the following:

* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 


`remove_vsx_params` supports the following:

* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 


`set_physical_interface_params` supports the following:

* `name` - (Required) Name of the interface. 
* `vlan_trunk` - (Required) True if this interface is a VLAN trunk. 
* `vsx_name` - (Required) Name of the VSX Gateway or Cluster object. 


`set_vd_interface_params` supports the following:

* `leads_to` - (Optional) Virtual Switch or Virtual Router for this interface. 
* `name` - (Optional) Name of the interface. 
* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `anti_spoofing` - (Optional) The anti-spoofing enforcement setting of this interface. 
* `anti_spoofing_tracking` - (Optional) The anti-spoofing tracking setting of this interface. 
* `ipv4_address` - (Optional) IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `ipv6_address` - (Optional) IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `mtu` - (Optional) MTU of this interface. 
* `new_leads_to` - (Optional) New Virtual Switch or Virtual Router for this interface. 
* `propagate` - (Optional) Propagate IPv4 route to adjacent virtual devices. 
* `propagate6` - (Optional) Propagate IPv6 route to adjacent virtual devices. 
* `specific_group` - (Optional) Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'. 
* `topology` - (Optional) Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS. 


`set_vd_params` supports the following:

* `vd` - (Required) Name of the Virtual System, Virtual Switch, or Virtual Router. 
* `calc_topology_auto` - (Optional) Calculate interface topology automatically based on routes.<br/>Relevant only for Virtual Systems.<br/>Do not use for virtual devices. 
* `ipv4_address` - (Optional) Main IPv4 Address.<br/>Relevant only if this device is a Virtual System.<br/>Do not use for other virtual devices. 
* `ipv4_instances` - (Optional) Number of IPv4 instances for the Virtual System.<br/>Must be greater or equal to 1.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode. 
* `ipv6_address` - (Optional) Main IPv6 Address.<br/>Relevant only if this device is a Virtual System.<br/>Do not use for other virtual devices. 
* `ipv6_instances` - (Optional) Number of IPv6 instances for the Virtual System.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode. 
* `vs_mtu` - (Optional) MTU of the Virtual System.<br/>Only relevant for Virtual Systems in bridge mode.<br/>Do not use for other virtual devices. 


`vti_settings` supports the following:

* `local_ipv4_address` - (Optional) The IPv4 address of the VPN tunnel on this Virtual System. 
* `peer_name` - (Optional) The name of the remote peer object as defined in the VPN community. 
* `remote_ipv4_address` - (Optional) The IPv4 address of the VPN tunnel on the remote VPN peer. 
* `tunnel_id` - (Optional) Optional unique Tunnel ID.<br/>Automatically assigned by the system if empty. 


`interfaces` supports the following:

* `leads_to` - (Optional) Virtual Switch or Virtual Router for this interface. 
* `name` - (Optional) Name of the interface. 
* `anti_spoofing` - (Optional) The anti-spoofing enforcement setting of this interface. 
* `anti_spoofing_tracking` - (Optional) The anti-spoofing tracking setting of this interface. 
* `ipv4_address` - (Optional) IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `ipv4_netmask` - (Optional) IPv4 Subnet mask of this interface. 
* `ipv4_prefix` - (Optional) IPv4 CIDR prefix of this interface. 
* `ipv6_address` - (Optional) IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router. 
* `ipv6_netmask` - (Optional) IPv6 Subnet mask of this interface. 
* `ipv6_prefix` - (Optional) IPv6 CIDR prefix of this interface. 
* `mtu` - (Optional) MTU of this interface. 
* `propagate` - (Optional) Propagate IPv4 route to adjacent virtual devices. 
* `propagate6` - (Optional) Propagate IPv6 route to adjacent virtual devices. 
* `specific_group` - (Optional) Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'. 
* `topology` - (Optional) Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS. 


`routes` supports the following:

* `destination` - (Optional) Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6. 
* `next_hop` - (Optional) Next hop IP address. 
* `leads_to` - (Optional) Virtual Router for this route<br/>This VD must have an existing connection to the VR. 
* `netmask` - (Optional) Subnet mask for this route. 
* `prefix` - (Optional) CIDR prefix for this route. 
* `propagate` - (Optional) Propagate this route to adjacent virtual devices. 


`members` supports the following:

* `ipv4_address` - (Optional) Main IPv4 Address of the VSX Cluster member.<br/>Mandatory if the VSX Cluster has an IPv4 Address. 
* `ipv6_address` - (Optional) Main IPv6 Address of the VSX Cluster member.<br/>Mandatory if the VSX Cluster has an IPv6 Address. 
* `name` - (Required) Name of the new VSX Cluster member. 
* `sic_otp` - (Required) SIC one-time-password of the VSX Gateway or Cluster member.<br/>Password must be between 4-127 characters in length. 
* `sync_ip` - (Required) Sync IP address for the VSX Cluster member. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

