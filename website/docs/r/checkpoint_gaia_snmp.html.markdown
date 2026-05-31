---
layout: "checkpoint"
page_title: "checkpoint_gaia_snmp"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-snmp"
description: |-
This resource allows you to execute Check Point Snmp.
---

# checkpoint_gaia_snmp

This resource allows you to execute Check Point Snmp.

## Example Usage


```hcl
resource "checkpoint_gaia_snmp" "example" {
  enabled    = true
  contact    = "ops-team"
  location   = "datacenter-1"
  version    = "any"
  interfaces = "all"
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Enables/Disables the SNMP Agent 
* `version` - (Optional) Configures the supported SNMP version: all - support SNMP v1, v2 and v3 v3-Only - support SNMP v3 only 
* `trap_usm` - (Optional) The user which will generate the SNMP traps, should be existed usm user 
* `contact` - (Optional) SNMP contact string 
* `location` - (Optional) SNMP location string: Specifies a string that contains the location for the device 
* `read_only_community` - (Optional) SNMP read-only community password, Where: * read-only: lets you only read the values of SNMP objects 
* `read_write_community` - (Optional) SNMP read-write community password, Where: * read-write: read and set the values as well 
* `interfaces` - (Optional) Adds a local interface to the list of local interfaces, on which the SNMP daemon listens 
* `pre_defined_traps_settings` - (Optional) Pre-defined traps settings pre_defined_traps_settings blocks are documented below.
* `custom_traps_settings` - (Optional) Custom traps settings custom_traps_settings blocks are documented below.
* `vsx_settings` - (Optional) VSX settings vsx_settings blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`pre_defined_traps_settings` supports the following:

* `polling_frequency` - (Optional) Polling interval in seconds 


`custom_traps_settings` supports the following:

* `clear_trap_interval` - (Optional) Interval in second between clear traps 
* `clear_trap_amount` - (Optional) Number of clear traps that is sent after custom trap termination 


`vsx_settings` supports the following:

* `enabled` - (Optional) True if SNMP is in vsx mode 
* `vs_access` - (Optional) SNMP vs-access type direct/indirect queries on Virtual-Devices direct: SNMP direct queries on Virtual-Devices indirect: SNMP direct queries via VS0 
* `sysname` - (Optional) This command is relevant only for VSX with SNMP VS mode, Where: False (default) = the sysname OID for all Virtual Devices will return the same result: VS0 hostname True = * VS0 sysname OID returns the VSX hostname * Virtual Device sysname OID returns the Check Point object name of the Virtual Device 
