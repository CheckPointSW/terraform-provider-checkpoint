---
layout: "checkpoint"
page_title: "checkpoint_management_resource_mms"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-resource-mms"
description: |- Use this data source to get information on an existing MMS resource.
---


# checkpoint_management_resource_mms

Use this data source to get information on an existing MMS resource.

## Example Usage


```hcl
data "checkpoint_management_resource_mms" "data_mms" {
  name = "resource_mms_example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `track` - Logs the activity when a packet matches on a Firewall Rule with the Resource.
* `action` - Accepts or Drops traffic that matches a Firewall Rule using the Resource.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
