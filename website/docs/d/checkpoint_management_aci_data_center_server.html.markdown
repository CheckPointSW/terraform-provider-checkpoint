---
layout: "checkpoint"
page_title: "checkpoint_management_aci_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-aci-data-center-server"
description: |- Use this data source to get information on an existing Cisco APIC data center server.
---

# Resource: checkpoint_management_aci_data_center_server

Use this data source to get information on an existing Cisco APIC Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_aci_data_center_server" "testAci" {
  name     = "MyAci"
  username = "USERNAME"
  password = "PASSWORD"
  urls     = ["url1", "url2"]
}

data "checkpoint_management_aci_data_center_server" "data_aci_data_center_server" {
  name = "${checkpoint_management_aci_data_center_server.testAci.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
