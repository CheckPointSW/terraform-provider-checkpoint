---
layout: "checkpoint"
page_title: "checkpoint_management_network_probe"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-network-probe"
description: |-
Use this data source to get information on an existing Check Point Network Probe.
---

# Data Source: checkpoint_management_network_probe

Use this data source to get information on an existing Check Point Network Probe.

## Example Usage


```hcl
resource "checkpoint_management_network_probe" "example" {
  name = "network1"
  icmp_options = {
    source = "host1"
    destination = "host2"
  }
  install_on = ["gw1","gw2"]
  interval  = "20"
  protocol = "icmp"
}
data "checkpoint_management_network_probe" "data" {
  uid = "${checkpoint_management_network_probe.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.
* `http_options` - Additional options when [protocol] is set to "http".http_options blocks are documented below.
* `icmp_options` -  Additional options when [protocol] is set to "icmp".icmp_options blocks are documented below.
* `install_on` -  Collection of Check Point Security Gateways that generate the probe, identified by name or UID.install_on blocks are documented below.
* `protocol` - The probing protocol to use. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `interval` -  The time interval (in seconds) between each probe request.<br>Best Practice - The interval value should be lower than the timeout value. 
* `timeout` - The probe expiration timeout (in seconds). If there is not a single reply within this time, the status of the probe changes to "Down". 
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string.

`http_options` supports the following:
* `destination` -  The destination URL. 


`icmp_options` supports the following:
* `destination` -  Name  of an existing object with a unicast IPv4 address (Host, Security Gateway, and so on). A unicast IPv4 address string (if you do not want to create such an object). 
* `source` -  One of these: The string "main-ip" (the probe uses the main IPv4 address of the Security Gateway objects you specified in the parameter [install-on]). Name  of an existing object of type 'Host' with a unicast IPv4 address.
