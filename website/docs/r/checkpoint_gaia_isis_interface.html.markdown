---
layout: "checkpoint"
page_title: "checkpoint_gaia_isis_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-isis-interface"
description: |-
This resource allows you to execute Check Point Isis Interface.
---

# checkpoint_gaia_isis_interface

This resource allows you to execute Check Point Isis Interface.

## Example Usage


```hcl
# Step 1: configure the ISIS instance with a system-id
resource "checkpoint_gaia_command_set_isis" "isis_setup" {
  system_id = "0101.0101.0101"
}

# Step 2: add the ISIS interface
resource "checkpoint_gaia_isis_interface" "example" {
  interface      = "eth0"
  address_family = "ipv4"
  circuit_type   = "level-1-2"
  lsp_interval   = "default"
  passive_mode   = true

  depends_on = [checkpoint_gaia_command_set_isis.isis_setup]
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the interface 
* `address_family` - (Required) Address family that the interface will run on 
* `name` - (Computed) The interface name of the interface to be queried 
* `advertise` - (Optional) Advertise this interfaces IP address 
* `circuit_type` - (Optional) Set level for the interface to run on 
* `csnp_interval` - (Optional) Configure IS-IS Interface Csnp Interval csnp_interval blocks are documented below.
* `hello` - (Optional) Configure ISIS interface hello hello blocks are documented below.
* `ip_reachability` - (Optional) Configure bidirectional forwarding detection (BFD) for interface 
* `ipv6` - (Optional) Configure IS-IS ipv6 options. Note that ipv6 multi topology must be enabled ipv6 blocks are documented below.
* `lsp_interval` - (Optional) Configure delay between sending LSPs 
* `mesh_group` - (Optional) Configure this interface as a member of a mesh group 
* `metric` - (Optional) Set the metric (cost) of this interface metric blocks are documented below.
* `passive_mode` - (Optional) Enable or disable passive operation 
* `point_to_point` - (Optional) Configure point to point options point_to_point blocks are documented below.
* `priority` - (Optional) Set DIS priority priority blocks are documented below.
* `protocol_instance` - (Computed) The instance to be queried 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `limit` - (Computed) No description available. 
* `offset` - (Computed) No description available. 
* `order` - (Computed) No description available. 


`csnp_interval` supports the following:

* `interval` - (Optional) Set the csnp interval configuration 
* `level` - (Optional) Set the level for the csnp configuration 


`hello` supports the following:

* `padding` - (Optional) Set hello padding for interface 
* `timers` - (Optional) Set level 1 configuration timers blocks are documented below.


`ipv6` supports the following:

* `advertise` - (Optional) Advertise this interfaces IP address 
* `ip_reachability` - (Optional) Configure bidirectional forwarding detection (BFD) for interface 
* `metric` - (Optional) Set the metric (cost) of this interface metric blocks are documented below.


`metric` supports the following:

* `metric` - (Optional) Set the interface metric interval configuration 
* `level` - (Optional) Set the level for this metric configuration 


`point_to_point` supports the following:

* `toggle` - (Optional) Configure toggle 
* `retransmit_interval` - (Optional) Configure retransmit interval 
* `retransmit_throttle_interval` - (Optional) Configure retransmit Throttle interval 


`priority` supports the following:

* `value` - (Optional) Set the level 1 interface priority interval configuration 
* `level` - (Optional) Set the level 2 csnp configuration 


`timers` supports the following:

* `holdtime` - (Optional) Set holdtime 
* `interval` - (Optional) Set interval 
* `level` - (Optional) The IS-IS level that this entry belongs to 


`metric` supports the following:

* `metric` - (Optional) Set the interface metric interval configuration 
* `level` - (Optional) Set the level for this metric configuration 
