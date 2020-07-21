---
layout: "checkpoint"
page_title: "checkpoint_management_data_wildcard"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-wildcard"
description: |-
  Use this data source to get information on an existing Check Point Wildcard.
---

# checkpoint_management_data_wildcard

Use this data source to get information on an existing Check Point Wildcard.

## Example Usage


```hcl
resource "checkpoint_management_wildcard" "wildcard" {
    name = "%s"
	ipv4_address = "192.168.2.1"
 	ipv4_mask_wildcard = "0.0.0.128"
}

data "checkpoint_management_data_wildcard" "data_wildcard" {
    name = "${checkpoint_management_wildcard.wildcard.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `ipv4_address` - IPv4 address. 
* `ipv4_mask_wildcard` - IPv4 mask wildcard. 
* `ipv6_address` - IPv6 address. 
* `ipv6_mask_wildcard` - IPv6 mask wildcard. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
* `groups` - Collection of group identifiers.