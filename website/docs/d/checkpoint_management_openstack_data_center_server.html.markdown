---
layout: "checkpoint"
page_title: "checkpoint_management_openstack_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-openstack-data-center-server"
description: |- Use this data source to get information on an existing OpenStack Data Center Server.
---

# Resource: checkpoint_management_openstack_data_center_server

Use this data source to get information on an existing OpenStack Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_openstack_data_center_server" "testOpenStack" {
  name     = "MyOpenStack"
  username = "USERNAME"
  password = "PASSWORD"
  hostname = "HOSTNAME"
}

data "checkpoint_management_openstack_data_center_server" "data_openstack_data_center_server" {
  name = "${checkpoint_management_openstack_data_center_server.testOpenStack.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
