---
layout: "checkpoint"
page_title: "checkpoint_management_radius_group"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-radius-group"
description: |-
Use this data source to get information on an existing Check Point Radius Group.
---

# Data Source: checkpoint_management_radius_group

Use this data source to get information on an existing Check Point Radius Group.

## Example Usage


```hcl
resource "checkpoint_management_host" "host" {
  name = "My Host"
  ipv4_address = "1.2.3.4"
}

resource "checkpoint_management_radius_server" "radius_server" {
  name = "New Radius Server"
  server = "${checkpoint_management_host.host.name}"
  shared_secret = "123"
}

resource "checkpoint_management_radius_group" "radius_group" {
  name    = "New Radius Group"
  members = ["${checkpoint_management_radius_server.radius_server.name}"]
}

data "checkpoint_management_radius_group" "data_radius_group" {
  name = "${checkpoint_management_radius_group.radius_group.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `members` - Collection of radius servers identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
