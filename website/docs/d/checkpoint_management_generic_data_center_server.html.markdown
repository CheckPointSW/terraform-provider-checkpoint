---
layout: "checkpoint"
page_title: "checkpoint_management_generic_data_center_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-generic-data-center-server"
description: |- Use this data source to get information on an existing generic data center server.
---

# Data Source: checkpoint_management_generic_data_center_server

Use this data source to get information on an existing Generic Data Center Server.

## Example Usage


```hcl
resource "checkpoint_management_generic_data_center_server" "generic_test" {
  name            = "test"
  url             = "MY_URL"
  interval        = "60"
  comments        = "testing generic data center"
  color          = "crete blue"
  tags            = ["terraform"]
}

data "checkpoint_management_generic_data_center_server" "data_generic_data_center_server" {
    name = "${checkpoint_management_generic_data_center_server.generic_test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required if uid is not given) Object name.
* `uid` - (Required if name is not given) Object unique identifier.
* `details-level` - The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
