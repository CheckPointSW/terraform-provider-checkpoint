---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_initial_setup"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-initial-setup"
description: |-
This resource allows you to execute Check Point Set Initial Setup.
---

# checkpoint_gaia_command_set_initial_setup

This resource allows you to execute Check Point Set Initial Setup.

## Example Usage


```hcl
resource "checkpoint_gaia_command_set_initial_setup" "example" {
  password = "user-password"
  grub_password = "grub-password"
}
```

## Argument Reference

The following arguments are supported:

* `password` - (Optional) Password of user admin. Required in case default initial password has not been changed before 
* `security_gateway` - (Optional) Install Security Gateway. security_gateway blocks are documented below.
* `security_management` - (Optional) Install Security Management or Multi-Domain Server security_management blocks are documented below.
* `grub_password` - (Optional) Password of the GRUB maintenence. Required in case default initial GRUB password has not been changed before 


`security_gateway` supports the following:

* `activation_key` - (Optional) Secure Internal Communication key 
* `dynamically_assigned_ip` - (Optional) Enable DAIP (Dynamic IP) gateway. Should be false if cluster-member or security-management enabled 
* `cluster_member` - (Optional) Enable/Disable ClusterXL. 
* `vsnext` - (Optional) Enable/Disable VSNext. To use VSNext, elastic-xl must be true 
* `elastic_xl` - (Optional) Enable/Disable ElasticXL. Cannot be enabled in combination with cluster-member 


`security_management` supports the following:

* `type` - (Optional) Type of security management or Multi-Domain Server 
* `multi_domain` - (Optional) Install Security Multi-Domain Server, it can be primary or secondary or Log Server according to type parameter 
* `gui_clients` - (Optional) Choose which GUI clients can log into the Security Management. Fill one of the parameters (range/network/Single IP), for Multi-Domain Server it can be only Single IP or can keep the default value gui_clients blocks are documented below.
* `activation_key` - (Optional) Secure Internal Communication key, relevant in case of secondary or Log Server 
* `leading_interface` - (Optional) Leading Multi-Domain Server interface, relevant in case of Multi-Domain Server enabled 


`gui_clients` supports the following:

* `range` - (Optional) Range of IPs allowed to connect to management range blocks are documented below.
* `network` - (Optional) IPs from specific network allowed to connect to management network blocks are documented below.
* `single_ip` - (Optional) In case of a single IP which allowed to connect to management 


`range` supports the following:

* `first_ipv4_range` - (Optional) First IP in range 
* `last_ipv4_range` - (Optional) Last IP in range 


`network` supports the following:

* `address` - (Optional) IPv4 address of network 
* `mask_length` - (Optional) Mask length of network 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

