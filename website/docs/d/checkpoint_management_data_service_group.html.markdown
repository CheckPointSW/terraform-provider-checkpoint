---
layout: "checkpoint"
page_title: "checkpoint_management_data_service_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-service-group"
description: |-
  Use this data source to get information on an existing Check Point Service Group.
---

# Data Source: checkpoint_management_data_service_group

Use this data source to get information on an existing Check Point Service Group.

## Example Usage


```hcl
resource "checkpoint_management_service_group" "service_group" {
    name = "service group"
}

data "checkpoint_management_data_service_group" "data_service_group" {
    name = "${checkpoint_management_service_group.service_group.name}"
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
* `ranges` - Displays the service group's matched content as ranges of port numbers, in case 'show-as-ranges' is set to true.<br />In this case, the 'members' para...ranges blocks are documented below.


`ranges` supports the following:

* `excluded_others` - Objects which are not represented as port numbers and are negated in the given rule - for example if negate is set for the service of this rule. The d...
* `others` - Objects which are not represented as port numbers and match the given rule. The details-level parameter of the request determines whether they are dis...
* `tcp` - Range of TCP ports that match in the given rule.tcp blocks are documented below.
* `udp` - Range of UDP ports that match in the given rule.udp blocks are documented below.


`tcp` supports the following:

* `end` - N/A.
* `start` - N/A.


`udp` supports the following:

* `end` - N/A.
* `start` - N/A.
