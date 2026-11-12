---
layout: "checkpoint"
page_title: "checkpoint_management_azure_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-azure-data-center-server"
description: |- Use this data source to get information on an existing azure data center server.
---

# Resource: checkpoint_management_azure_data_center_server

Use this data source to get information on an existing Microsoft Azure Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_azure_data_center_server" "testAzure" {
  name                  = "myAzure"
  authentication_method = "user-authentication"
  username              = "MY-KEY-ID"
  password              = "MY-SECRET-KEY"
}

data "checkpoint_management_azure_data_center_server" "data_azure_data_center_server" {
  name = "${checkpoint_management_azure_data_center_server.testAzure.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
