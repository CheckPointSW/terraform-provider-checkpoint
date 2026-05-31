---
layout: "checkpoint"
page_title: "checkpoint_gaia_maestro_security_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-maestro-security-group"
description: |-
This resource allows you to execute Check Point Maestro Security Group.
---

# checkpoint_gaia_maestro_security_group

This resource allows you to execute Check Point Maestro Security Group.

## Example Usage


```hcl
resource "checkpoint_gaia_maestro_security_group" "example" {
  interfaces {
    resource_id = "1/1/1"
  }
  gateways {
    resource_id = "2108BA1058"
    description = "GW 2108BA1058 Description"
  }
  ftw_configuration {
    hostname          = "My_Host_Name"
    is_vsx            = true
    one_time_password = "otp_pass"
    admin_password    = "admin_pass"
  }
  mgmt_connectivity {
    ipv4_address     = "1.1.1.1"
    ipv4_mask_length = 24
    default_gateway  = "1.1.1.4"
  }
  description = "New Security Group Description"
}
```

## Argument Reference

The following arguments are supported:

* `interfaces` - (Required) Orchestrator port, or list of Orchestrator ports, that will be assigned to this Security Group. At least one of ‘id’ or ‘interface-name’ parameters must be provided interfaces blocks are documented below.
* `gateways` - (Required) Single Gateway or list of Gateways to be assigned to new Security Group gateways blocks are documented below.
* `ftw_configuration` - (Required) First Time Wizard configuration ftw_configuration blocks are documented below.
* `mgmt_connectivity` - (Required) The IP addresses that will be used to manage this Security Group mgmt_connectivity blocks are documented below.
* `resource_id` - (Optional, Computed) Security Group ID 
* `sites` - (Optional) List of Site descriptions. The security group will be assigned to sites automatically according to gateways associated with the Security Group sites blocks are documented below.
* `description` - (Optional) New Security Group description 
* `mgmt_interface_settings` - (Optional) Management interface settings of this Security Group. By default, values are create-mgmt-as-bond == True and bond-mode == 'active-backup'. mgmt_interface_settings blocks are documented below.
* `include_pending_changes` - (Computed) If true, show pending Security Groups changes. If false, show deployed topology 
* `id` - (Computed) Computed field, returned in the response. 


`interfaces` supports the following:

* `resource_id` - (Optional) Interface ID (e.g. "1/13/1") 
* `name` - (Optional) Interface name (e.g. "eth1-05") 
* `description` - (Optional) Description of the interface 


`gateways` supports the following:

* `resource_id` - (Optional) ID of this Gateway 
* `description` - (Optional) Description of this GW 


`ftw_configuration` supports the following:

* `hostname` - (Optional) Hostname for Security Group 
* `is_vsx` - (Optional) Determines if this Security Group is a VSX 
* `one_time_password` - (Optional) One time password for Secure Internal Communication (SIC) 
* `admin_password` - (Optional) Admin password for Security Group 


`mgmt_connectivity` supports the following:

* `ipv4_address` - (Optional) IPv4 address for Security Group 
* `ipv6_address` - (Optional) IPv6 address for Security Group. Supported starting from Gaia version R82.10 
* `ipv4_mask_length` - (Optional) IPv4 mask length for Security Group 
* `ipv6_mask_length` - (Optional) IPv6 mask length for Security Group. Supported starting from Gaia version R82.10 
* `default_gateway` - (Optional) Default Gateway address for Security Group 
* `ipv6_default_gateway` - (Optional) Default Gateway IPv6 address for Security Group. Supported starting from Gaia version R82.10 


`sites` supports the following:

* `resource_id` - (Optional) ID of this site 
* `description` - (Optional) Description of this site 


`mgmt_interface_settings` supports the following:

* `create_mgmt_as_bond` - (Optional) If True, a magg interface will be created for MGMT traffic. Every assigned MGMT interface will be enslaved to this magg. If False, only one of the assigned MGMT interfaces will be used for MGMT traffic. 
* `bond_mode` - (Optional) If create-mgmt-as-bond is true, this field determines the magg bond type. If create-mgmt-as-bond is false, this field will be ignored.Note that using "xor" or "8023AD" entails configuring a bond on the device this Maestro environment is connected to. 
