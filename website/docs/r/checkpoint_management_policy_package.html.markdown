---
layout: "checkpoint"
page_title: "checkpoint_management_policy_package"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-policy-package"
description: |-
This resource allows you to execute Check Point Policy Package.
---

# checkpoint_management_policy_package

This resource allows you to execute Check Point Policy Package.

## Example Usage


```hcl
resource "checkpoint_management_policy_package" "example" {
  name = "New_Standard_Package_1"
  comments = "My Comments"
  color = "green"
  threat_prevention = false
  access = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `access` - (Optional) True - enables, False - disables access & NAT policies, empty - nothing is changed. 
* `access_layers` - (Optional) Access policy layers.access_layers blocks are documented below.
* `desktop_security` - (Optional) True - enables, False - disables Desktop security policy, empty - nothing is changed. 
* `https_layer` - (Optional) HTTPS inspection policy layer. 
* `installation_targets` - (Optional) Which Gateways identified by the name or UID to install the policy on.installation_targets blocks are documented below.
* `qos` - (Optional) True - enables, False - disables QoS policy, empty - nothing is changed. 
* `qos_policy_type` - (Optional) QoS policy type. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `threat_layers` - (Optional) Threat policy layers.threat_layers blocks are documented below.
* `threat_prevention` - (Optional) True - enables, False - disables Threat policy, empty - nothing is changed. 
* `vpn_traditional_mode` - (Optional) True - enables, False - disables VPN traditional mode, empty - nothing is changed. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`access_layers` supports the following:

* `add` - (Optional) Collection of Access layer objects to be added identified by the name or UID.add blocks are documented below.
* `remove` - (Optional) Collection of Access layer objects to be removed identified by the name or UID.remove blocks are documented below.
* `value` - (Optional) Collection of Access layer objects to be set identified by the name or UID. Replaces existing Access layers.value blocks are documented below.


`threat_layers` supports the following:

* `add` - (Optional) Collection of Threat layer objects to be added identified by the name or UID.add blocks are documented below.
* `remove` - (Optional) Collection of Threat layer objects to be removed identified by the name or UID.remove blocks are documented below.
* `value` - (Optional) Collection of Threat layer objects to be set identified by the name or UID. Replaces existing Threat layers.value blocks are documented below.


`add` supports the following:

* `name` - (Optional) Layer name or UID. 
* `position` - (Optional) Layer position. 


`add` supports the following:

* `name` - (Optional) Layer name or UID. 
* `position` - (Optional) Layer position. 
