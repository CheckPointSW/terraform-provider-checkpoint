---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_simulate_packet"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-simulate-packet"
description: |-
This resource allows you to execute Check Point Simulate Packet.
---

# checkpoint_gaia_command_simulate_packet

This resource allows you to execute Check Point Simulate Packet.

## Example Usage


```hcl
resource "checkpoint_gaia_command_simulate_packet" "example" {
  source_ip          = "1.1.1.1"
  destination_ip     = "8.8.8.8"
  ip_protocol        = 6
  protocol_options {
    tcp {
      destination_port = "80"
    }
  }
  incoming_interface = "eth0"
}
```

## Argument Reference

The following arguments are supported:

* `source_ip` - (Required) Source IP, should match selected "ip-version" (which defaults to 4) 
* `destination_ip` - (Required) Destination IP, should match selected "ip-version" (which defaults to 4) 
* `ip_protocol` - (Required) ip-protocol either in integer form: based on IANA Protocol Number in decimal format
                     or a string of one of the following protocols: [UDP, TCP, ICMP] 
* `protocol_options` - (Required) Protocol options required for the selected ip-protocol.
please note, only the relevant protocol's options should be filled. protocol_options blocks are documented below.
* `incoming_interface` - (Required) Incoming interface name for the packet, identified by the name.
The simulated connection is inbound, in order to simulate a local outgoing connection, set incoming-interface to localhost 
* `ip_version` - (Optional) IP version of the source and destinations IPs, can be either 4 or 6 
* `application` - (Optional) Name of the Application/Category as defined in SmartConsole.
 You can specify multiple applications. application blocks are documented below.
* `protocol` - (Optional) Protocol to match for services that have "Protocol Signature" enabled 
* `check_access_rule_uid` - (Optional) Rule uid to check why the packet didn't match this rule 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 


`protocol_options` supports the following:

* `tcp` - (Optional) TCP specific required options. required if ip-protocol is "TCP" or its IANA Protocol Number "6". tcp blocks are documented below.
* `udp` - (Optional) UDP specific required options. required if ip-protocol is "UDP" or its IANA Protocol Number "17". udp blocks are documented below.
* `icmp` - (Optional) icmp specific required options. required if ip-protocol is "icmp" or its IANA Protocol Number in IPv4: "1" or in IPv6: "58". icmp blocks are documented below.


`TCP` supports the following:

* `destination_port` - (Optional) Destination port in the Decimal format.
This parameter is mandatory for the TCP (6) and UDP (17) protocols. 
* `source_port` - (Optional) Source port in the Decimal format. if not specified will default to 12345 


`UDP` supports the following:

* `destination_port` - (Optional) Destination port in the Decimal format.
This parameter is mandatory for the TCP (6) and UDP (17) protocols. 
* `source_port` - (Optional) Source port in the Decimal format. if not specified will default to 12345 


`icmp` supports the following:

* `type` - (Optional) a string of the desired icmp type in decimal format as seen in IANA icmp parameters 
* `code` - (Optional) a string of the desired icmp code in decimal format as seen in IANA icmp parameters,
 will defualt to 0. 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

