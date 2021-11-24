---
layout: "checkpoint"
page_title: "checkpoint_management_gcp_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-gcp-data-center-server"
description: |- Use this data source to get information on an existing Google Cloud Platform Data Center Server.
---

# Resource: checkpoint_management_gcp_data_center_server

Use this data source to get information on an existing Google Cloud Platform Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_gcp_data_center_server" "testGcp" {
  name                  = "myGcp"
  authentication_method = "key-authentication"
  private_key           = "MYKEY.json"
  ignore_warnings       = true
}

data "checkpoint_management_gcp_data_center_server" "data_gcp_data_center_server" {
  name = "${checkpoint_management_gcp_data_center_server.testGcp.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
