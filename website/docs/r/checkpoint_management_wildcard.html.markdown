---
layout: "checkpoint"
page_title: "checkpoint_management_wildcard"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-wildcard"
description: |-
This resource allows you to execute Check Point Wildcard.
---

# checkpoint_management_wildcard

This resource allows you to execute Check Point Wildcard.

## Example Usage


```hcl
resource "checkpoint_management_wildcard" "example" {
  name = "New Wildcard 1"
  ipv4_address = "192.168.2.1"
  ipv4_mask_wildcard = "0.0.0.128"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `ipv4_address` - (Optional) IPv4 address. 
* `ipv4_mask_wildcard` - (Optional) IPv4 mask wildcard. 
* `ipv6_address` - (Optional) IPv6 address. 
* `ipv6_mask_wildcard` - (Optional) IPv6 mask wildcard. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
