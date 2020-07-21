---
layout: "checkpoint"
page_title: "checkpoint_physical_interface"
sidebar_current: "docs-checkpoint-gaia-resource-checkpoint-physical-interface"
description: |-
  This resource allows you to set a Physical interface.
---

# checkpoint_physical_interface

This resource allows you to set a Physical interface.

## Example Usage


```hcl
resource "checkpoint_physical_interface" "physical_interface1" {
      name = "eth1"
      enabled = "true"
      ipv4_address = "20.30.1.10"
      ipv4_mask_length = 24
}

resource "checkpoint_physical_interface" "physical_interface2" {
      name = "eth2"
      enabled = "true"
      speed = "100M"
      duplex = "full"
}

resource "checkpoint_physical_interface" "physical_interface3" {
      name = "eth3"
      monitor_mode = "true"
      enabled = "true"
      ipv4_address = "1.2.3.4"
      ipv4_mask_length = 24
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Interface name.
* `enabled` - (Optional) Interface state.
* `ipv4_address` - (Optional) IPv4 address to set for the interface.
* `ipv4_mask_length` - (Optional) Interface IPv4 address mask length.
* `ipv6_address` - (Optional) IPv6 address to set for the interface.
* `ipv6_mask_length` - (Optional) Interface IPv6 address mask length.
* `ipv6_autoconfig` - (Optional) Configure IPv6 auto-configuration true/false.
* `mac_addr` - (Optional) Configure hardware address.
* `mtu` - (Optional) Interface Mtu.
* `rx_ringsize` - (Optional) Set receive buffer size for the interface.
* `tx_ringsize` - (Optional) Set transmit buffer size for the interface.
* `monitor_mode` - (Optional) Set monitor mode for the interface true/false.
* `auto_negotiation` - (Optional) Configure auto-negotiation. Activating Auto-Negotiation will skip the speed and duplex configuration.
* `duplex` - (Optional) duplex for the interface. Duplex is not relevant when 'auto_negotiation' is enabled.
* `speed` - (Optional) Interface link speed. Speed is not relevant when 'auto_negotiation' is enabled.
* `comments` - (Optional) interface Comments.













