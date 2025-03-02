---
layout: "checkpoint"
page_title: "checkpoint_management_generic_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-generic-data-center-server"
description: |- This resource allows you to execute Check Point generic data center server.
---

# Resource: checkpoint_management_generic_data_center_server

This resource allows you to execute Check Point Generic Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_generic_data_center_server" "generic_test" {
  name     = "test"
  url      = "MY_URL"
  interval = "60"
  comments = "testing generic data center"
  color    = "crete blue"
  tags     = ["terraform"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `url` - (**Required**) URL of the JSON feed (e.g. https://example.com/file.json).
* `interval` - (**Required**)    Update interval of the feed in seconds.
* `custom_header` - (Optional) When set to false, The admin is not using Key and Value for a Custom Header in order to
  connect to the feed server. When set to true, The admin is using Key and Value for a Custom Header in order to connect
  to the feed server.
* `custom_key` - (Optional) Key for the Custom Header, relevant and required only when custom_header set to true.
* `custom_value` - (Optional)    Value for the Custom Header, relevant and required only when custom_header set to true.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If
  ignore-warnings flag was omitted - warnings will also be ignored.
