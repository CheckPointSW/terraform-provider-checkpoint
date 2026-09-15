---
layout: "checkpoint"
page_title: "checkpoint_management_kubernetes_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-kubernetes-data-center-server"
description: |- Use this data source to get information on an existing Kubernetes Data Center Server.
---

# Resource: checkpoint_management_kubernetes_data_center_server

Use this data source to get information on an existing Kubernetes Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_kubernetes_data_center_server" "testKubernetes" {
  name       = "MyKubernetes"
  hostname   = "MY_HOSTNAME"
  token_file = "MY_TOKEN"
}

data "checkpoint_management_kubernetes_data_center_server" "data_kubernetes_data_center_server" {
  name = "${checkpoint_management_kubernetes_data_center_server.testKubernetes.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `data_center_type` - Data Center type.
* `properties` - Data Center properties.properties blocks are documented below.


`properties` supports the following:

* `name` - N/A.
* `value` - N/A.
