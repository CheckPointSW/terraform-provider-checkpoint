---
layout: "checkpoint"
page_title: "checkpoint_management_threat_exception"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-exception"
description: |-
  This resource allows you to execute Check Point Threat Exception.
---

# Data Source: checkpoint_management_threat_exception

This resource allows you to execute Check Point Threat Exception.

## Example Usage


```hcl
resource "checkpoint_management_threat_rule" "threat_rule" {
    name = "threat rule"
    layer = "Standard Threat Prevention"
    position = {top = "top"}
}

resource "checkpoint_management_threat_exception" "threat_exception" {
    name = "threat exception"
    layer = "Standard Threat Prevention"
    position = {top = "top"}
    rule_name = "${checkpoint_management_threat_rule.threat_rule.name}"
}

data "checkpoint_management_threat_exception" "data_threat_exception" {
    name = "${checkpoint_management_threat_exception.threat_exception.name}"
    layer = "${checkpoint_management_threat_exception.threat_exception.layer}"
    rule_name = "${checkpoint_management_threat_exception.threat_exception.rule_name}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `name` - (Optional) 	The name of the exception.
* `uid` - (Optional) Object unique identifier.
* `exception_group_uid` - (Optional) The UID of the exception-group.
* `exception_group_name` - (Optional) The name of the exception-group.
* `rule_uid` - (Optional) The UID of the parent rule.
* `rule_name` - (Optional) The name of the parent rule.
* `action` - Action-the enforced profile.
* `enabled` - Enable/Disable the rule.
* `install_on` - Which Gateways identified by the name or UID to install the policy on.
* `source` - Collection of Network objects identified by the name or UID.
* `source_negate` - True if negate is set for source.
* `destination` - Collection of Network objects identified by the name or UID.
* `destination_negate` - True if negate is set for destination.
* `protected_scope` - Collection of objects defining Protected Scope identified by the name or UID.
* `protected_scope_negate` - True if negate is set for Protected Scope.
* `service` - Collection of Network objects identified by the name or UID.
* `service_negate` - True if negate is set for service.
* `protection_or_site` - Collection of protection or site objects identified by the name or UID.
* `track` - Packet tracking.
* `comments` - Comments string.
* `owner` - Owner UID.