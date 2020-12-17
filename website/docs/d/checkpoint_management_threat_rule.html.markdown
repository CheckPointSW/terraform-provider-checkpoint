---
layout: "checkpoint"
page_title: "checkpoint_management_threat_rule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-rule"
description: |-
  This resource allows you to execute Check Point Threat Rule.
---

# Data Source: checkpoint_management_threat_rule

This resource allows you to execute Check Point Threat Rule.

## Example Usage


```hcl
resource "checkpoint_management_threat_rule" "test" {
    name = "threat rule"
    layer = "Standard Threat Prevention"
    position = {top = "top"}
}

data "checkpoint_management_threat_rule" "test" {
    layer = "${checkpoint_management_threat_rule.test.layer}"
    name = "${checkpoint_management_threat_rule.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that the rule belongs to identified by the name or UID.
* `name` - (Optional) Rule name.
* `uid` - (Optional) Object unique identifier.
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
* `track` - Packet tracking.
* `track_settings` - Threat rule track settings. track_settings block are documented below.
* `comments` - Comments string.
* `exceptions` - Collection of the rule's exceptions identified by UID.

`track_settings` supports the following:

* `packet_capture` - Packet capture.