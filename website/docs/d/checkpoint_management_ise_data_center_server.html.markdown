---
layout: "checkpoint"
page_title: "checkpoint_management_ise_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-ise-data-center-server"
description: |- Use this data source to get information on an existing Cisco ISE data center server.
---

# Resource: checkpoint_management_ise_data_center_server

Use this data source to get information on an existing Cisco ISE Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_ise_data_center_server" "testIse" {
  name      = "MyIse"
  username  = "USERNAME"
  password  = "PASSWORD"
  hostnames = ["host1", "host2"]
}

data "checkpoint_management_ise_data_center_server" "data_ise_data_center_server" {
  name = "${checkpoint_management_ise_data_center_server.testIse.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
