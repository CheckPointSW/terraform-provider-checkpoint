---
layout: "checkpoint"
page_title: "checkpoint_management_service_group"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-service-group"
description: |-
  This resource allows you to add/update/delete Check Point Service Group.
---

# checkpoint_management_service_group

This resource allows you to add/update/delete Check Point Service Group.

## Example Usage


```hcl
resource "checkpoint_management_service_group" "example" {
  name = "New Service Group 1"
  members = [ "https", "bootp", "nisplus", "HP-OpCdistm" ]
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `members` - (Optional) Collection of Network objects identified by the name or UID.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `tags` - (Optional) Collection of tag identifiers.
