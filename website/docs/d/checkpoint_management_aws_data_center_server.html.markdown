---
layout: "checkpoint"
page_title: "checkpoint_management_aws_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-aws-data-center-server"
description: |- Use this data source to get information on an existing AWS data center server.
---

# Resource: checkpoint_management_aws_data_center_server

Use this data source to get information on an existing AWS Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_aws_data_center_server" "testAws" {
  authenticationMethod = "user-authentication"
  accessKeyId          = "MY-KEY-ID"
  secretAccessKey      = "MY-SECRET-KEY"
  region               = "us-east-1"
}

data "checkpoint_management_aws_data_center_server" "data_aws_data_center_server" {
  name = "${checkpoint_management_aws_data_center_server.testAws.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
