---
layout: "checkpoint"
page_title: "checkpoint_management_threat_rule"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-threat-rule"
description: |-
  This resource allows you to add/update/delete Check Point Threat Rule.
---

# Resource: checkpoint_management_threat_rule

This resource allows you to add/update/delete Check Point Threat Rule.

## Example Usage


```hcl
resource "checkpoint_management_threat_rule" "test" {
	name = "threat rule"
    layer = "Standard Threat Prevention"
	position = {top = "top"}
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `name` - (Optional) Rule name.
* `action` - (Optional) Action-the enforced profile.
* `enabled` - (Optional) Enable/Disable the rule.
* `install_on` - (Optional) Which Gateways identified by the name or UID to install the policy on.
* `source` - (Optional) Collection of Network objects identified by the name or UID.
* `source_negate` - (Optional) True if negate is set for source.
* `destination` - (Optional) Collection of Network objects identified by the name or UID.
* `destination_negate` - (Optional) True if negate is set for destination.
* `protected_scope` - (Optional) Collection of objects defining Protected Scope identified by the name or UID.
* `protected_scope_negate` - (Optional) True if negate is set for Protected Scope.
* `service` - (Optional) Collection of Network objects identified by the name or UID.
* `service_negate` - (Optional) True if negate is set for service.
* `track` - (Optional) Packet tracking.
* `track_settings` - (Optional) Threat rule track settings. track_settings block are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.
* `exceptions` - (Computed) Collection of the rule's exceptions identified by UID.


`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.

`track_settings` supports the following:

* `packet_capture` - (Optional) Packet capture.


## Import

`checkpoint_management_threat_rule` can be imported by using the following format: LAYER_NAME;RULE_UID

```
$ terraform import checkpoint_management_threat_rule.example Layer_Name;9423d36f-2d66-4754-b9e2-e9f4493751d3
```