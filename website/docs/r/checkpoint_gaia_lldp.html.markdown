---
layout: "checkpoint"
page_title: "checkpoint_gaia_lldp"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-lldp"
description: |-
This resource allows you to execute Check Point Lldp.
---

# checkpoint_gaia_lldp

This resource allows you to execute Check Point Lldp.

## Example Usage


```hcl
resource "checkpoint_gaia_lldp" "example" {
  enabled = false
  interfaces {
    interface_name = "eth0"
    mode           = "transmit"
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) LLDP State 
* `timers` - (Optional) LLDP Timers timers blocks are documented below.
* `tlv` - (Optional) LLDP Tlv tlv blocks are documented below.
* `interfaces` - (Optional) LLDP per-interface configuration. Each block sets the LLDP mode for one interface. Valid `mode` values: `transmit-and-receive`, `transmit`, `receive`. 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`timers` supports the following:

* `hold_time_multiplier` - (Optional) Define LLDP hold time multiplier interval to cache learned information before discarding. Range: 2-10, Default: 4 (giving 120 seconds LLDP cache lifetime with other defaults). 
* `transmit_interval` - (Optional) Define LLDP packet transmitting interval (Seconds). Range: 8-32768 seconds, Default: 30 seconds. 


`tlv` supports the following:

* `management_address` - (Optional) Define Gaia to send the Management Address information in the LLDP packets. management_address blocks are documented below.
* `port_description` - (Optional) Define Gaia to send the Port Description information in the LLDP packets. 
* `system_capabilities` - (Optional) Define Gaia to send the System Capabilities information in the LLDP packets. 
* `system_description` - (Optional) Define Gaia to send the System Description information in the LLDP packets. 
* `system_name` - (Optional) Define Gaia to send the System Name information in the LLDP packets. 


`management_address` supports the following:

* `enabled` - (Optional) Define Gaia to send the Management Address information in the LLDP packets. 
* `ip_from` - (Optional) configured-interface - Send Configured interface IP within the LLDP packets, mgmt-interface - Send Management interface IP within the LLDP packets, Default is configured-interface. (supported from version R81.20) 
