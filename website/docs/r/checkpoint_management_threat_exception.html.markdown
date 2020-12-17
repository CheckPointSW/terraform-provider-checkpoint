---
layout: "checkpoint"
page_title: "checkpoint_management_threat_exception"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-threat-exception"
description: |-
  This resource allows you to add/update/delete Check Point Threat Exception.
---

# Resource: checkpoint_management_threat_exception

This resource allows you to add/update/delete Check Point Threat Exception.

## Example Usage


```hcl
resource "checkpoint_management_threat_rule" "test" {
    name = "threat rule"
    layer = "Standard Threat Prevention"
    position = {top = "top"}
}

resource "checkpoint_management_threat_exception" "test" {
    name = "threat exception"
    layer = "Standard Threat Prevention"
    position = {top = "top"}
    rule_name = "${checkpoint_management_threat_rule.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `position` - (Required) Position in the rulebase. Position blocks are documented below.
* `name` - (Required) 	The name of the exception.
* `exception_group_uid` - (Optional) The UID of the exception-group.
* `exception_group_name` - (Optional) The name of the exception-group.
* `rule_uid` - (Optional) The UID of the parent rule.
* `rule_name` - (Optional) The name of the parent rule.
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
* `protection_or_site` - (Optional) Collection of protection or site objects identified by the name or UID.
* `track` - (Optional) Packet tracking.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.
* `owner` - (Computed) Owner UID.


`position` supports the following:

* `top` - (Optional) Add rule at the top of the rulebase.
* `above` - (Optional) Add rule above specific section/rule identified by uid or name.
* `below` - (Optional) Add rule below specific section/rule identified by uid or name.
* `bottom` - (Optional) Add rule at the bottom of the rulebase.


## Import

`checkpoint_management_threat_exception` can be imported by using the following format: LAYER_UID;exception_group_uid or rule_uid;EXCEPTION_GROUP_UID or PARENT_RULE_UID;EXCEPTION_GROUP_UID

```
$ terraform import checkpoint_management_threat_exception.example "Standard Threat Prevention;exception_group_uid;9423d36f-2d66-4754-b9e2-e9f4493751e5;9423d36f-2d66-4754-b9e2-e9f4493751d3"
```