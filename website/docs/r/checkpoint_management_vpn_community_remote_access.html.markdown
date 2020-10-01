---
layout: "checkpoint"
page_title: "checkpoint_management_vpn_community_remote_access"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-vpn-community-remote-access"
description: |-
This resource allows you to execute Check Point VPN Community Remote Access.
---

# Resource: checkpoint_management_vpn_community_remote_access

This resource allows you to execute Check Point VPN Community Remote Access.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_remote_access" "example" {
    name = "RemoteAccess"
    user_groups = ["All Users"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `gateways` - (Optional) Collection of Gateway objects identified by the name or UID.
* `user_groups` - (Optional) Collection of User group objects identified by the name or UID.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.