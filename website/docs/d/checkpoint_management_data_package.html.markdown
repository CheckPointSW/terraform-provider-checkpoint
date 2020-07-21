---
layout: "checkpoint"
page_title: "checkpoint_management_data_package "
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-package"
description: |-
  Use this data source to get information on an existing Check Point Package Object.
---

# checkpoint_management_data_package

Use this data source to get information on an existing Check Point Package Object.

## Example Usage

```hcl
resource "checkpoint_management_package" "package" {
    name = "My Package"
}

data "checkpoint_management_data_package" "data_package" {
    name = "${checkpoint_management_package.package.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `access` - True - enables, False - disables access & NAT policies, empty - nothing is changed.
* `desktop_security` - True - enables, False - disables Desktop security policy, empty - nothing is changed.
* `qos` - True - enables, False - disables QoS policy, empty - nothing is changed.
* `qos_policy_type` - QoS policy type.
* `threat_prevention` - True - enables, False - disables Threat policy, empty - nothing is changed.
* `vpn_traditional_mode` - True - enables, False - disables VPN traditional mode, empty - nothing is changed.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.