---
layout: "checkpoint"
page_title: "checkpoint_management_data_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-group"
description: |-
  Use this data source to get information on an existing Check Point Group.
---

# Data Source: checkpoint_management_data_group

Use this data source to get information on an existing Check Point Group.

## Example Usage


```hcl
resource "checkpoint_management_group" "group" {
    name = "My Group"
}

data "checkpoint_management_data_group" "data_group" {
    name = "${checkpoint_management_group.group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `members` - Collection of Network objects identified by the name or UID.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.
* `ranges` - Displays the group's matched content as ranges of IP addresses, in case 'show-as-ranges' is set to true.<br />In this case, the 'members' parameter is...ranges blocks are documented below.


`ranges` supports the following:

* `excluded_others` - Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...
* `ipv4` - Range of IPv4 addresses that match in the given rule.ipv4 blocks are documented below.
* `ipv6` - Range of IPv6 addresses that match in the given rule.ipv6 blocks are documented below.
* `others` - Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...


`ipv4` supports the following:

* `end` - N/A.
* `start` - N/A.


`ipv6` supports the following:

* `end` - N/A.
* `start` - N/A.
