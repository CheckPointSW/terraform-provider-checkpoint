---
layout: "checkpoint"
page_title: "checkpoint_management_nutanix_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-nutanix-data-center-server"
description: |- Use this data source to get information on an existing Nutanix data center server.
---

# Data Source: checkpoint_management_nutanix_data_center_server

Use this data source to get information on an existing Nutanix Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_nutanix_data_center_server" "testNutanix" {
  name = "MY-NUTANIX"
  hostname = "127.0.0.1"
  username = "admin"
  password = "admin"
}

data "checkpoint_management_nutanix_data_center_server" "data_nutanix_data_center_server" {
  name = "${checkpoint_management_nutanix_data_center_server.testNutanix.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `data_center_type` - Data center type.
* `properties` - Data center properties. properties blocks are documented below.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`properties` supports the following:

* `name`
* `value`