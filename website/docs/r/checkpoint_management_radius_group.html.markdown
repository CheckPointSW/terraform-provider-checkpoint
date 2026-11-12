---
layout: "checkpoint"
page_title: "checkpoint_management_radius_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-radius-group"
description: |-
This resource allows you to add/update/delete Check Point Radius Group.
---

# Resource: checkpoint_management_radius_group

This resource allows you to add/update/delete Check Point Radius Group.

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
  name = "New Radius Group"
  members = ["${checkpoint_management_radius_server.radius_server.name}"]
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Must be unique in the domain.
* `members` - (Optional) Collection of radius servers identified by the name or UID.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
