---
layout: "checkpoint"
page_title: "checkpoint_management_package "
sidebar_current: "docs-checkpoint-resource-checkpoint-management-package"
description: |-
  This resource allows you to add/update/delete Check Point Package Object.
---

# checkpoint_management_package

This resource allows you to add/update/delete Check Point Package Object.

## Example Usage

```hcl
resource "checkpoint_management_package" "example" {
  name = "New_Standard_Package_1"
  comments = "My Comments"
  color = "green"
  threat_prevention = false
  access = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `access` - (Optional) True - enables, False - disables access & NAT policies, empty - nothing is changed.
* `desktop_security` - (Optional) True - enables, False - disables Desktop security policy, empty - nothing is changed.
* `qos` - (Optional) True - enables, False - disables QoS policy, empty - nothing is changed.
* `qos_policy_type` - (Optional) QoS policy type.
* `threat_prevention` - (Optional) True - enables, False - disables Threat policy, empty - nothing is changed.
* `vpn_traditional_mode` - (Optional) True - enables, False - disables VPN traditional mode, empty - nothing is changed.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `comments` - (Optional) Comments string.
* `tags` - (Optional) Collection of tag identifiers.

