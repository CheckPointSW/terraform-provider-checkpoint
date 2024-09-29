---
layout: "checkpoint"
page_title: "checkpoint_management_network_probe"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-network-probe"
description: |-
This resource allows you to execute Check Point Network Probe.
---

# checkpoint_management_network_probe

This resource allows you to execute Check Point Network Probe.

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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `http_options` - (Optional) Additional options when [protocol] is set to "http".http_options blocks are documented below.
* `icmp_options` - (Optional) Additional options when [protocol] is set to "icmp".icmp_options blocks are documented below.
* `install_on` - (Required) Collection of Check Point Security Gateways that generate the probe, identified by name or UID.install_on blocks are documented below.
* `protocol` - (Optional) The probing protocol to use. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `interval` - (Optional) The time interval (in seconds) between each probe request.<br>Best Practice - The interval value should be lower than the timeout value. 
* `timeout` - (Optional) The probe expiration timeout (in seconds). If there is not a single reply within this time, the status of the probe changes to "Down". 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`http_options` supports the following:
* `destination` - (Optional) The destination URL. 


`icmp_options` supports the following:
* `destination` - (Optional) One of these:Name or UID of an existing object with a unicast IPv4 address (Host, Security Gateway, and so on). A unicast IPv4 address string (if you do not want to create such an object). 
* `source` - (Optional) One of these: The string "main-ip" (the probe uses the main IPv4 address of the Security Gateway objects you specified in the parameter [install-on]). Name or UID of an existing object of type 'Host' with a unicast IPv4 address. A unicast IPv4 address string (if you do not want to create such an object). 
