---
layout: "checkpoint"
page_title: "checkpoint_management_kubernetes_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-kubernetes-data-center-server"
description: |- This resource allows you to execute Check Point kubernetes data center server.
---

# Resource: checkpoint_management_kubernetes_data_center_server

This resource allows you to execute Check Point Kubernetes Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_kubernetes_data_center_server" "testKubernetes" {
  name       = "MyKubernetes"
  hostname   = "MY_HOSTNAME"
  token_file = "MY_TOKEN"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `hostname` - (**Required**) IP address or hostname of the Kubernetes server.
* `token_file` - (**Required**) Kubernetes access token encoded in base64.
* `ca_certificate` - (Optional) The Kubernetes public certificate key encoded in base64.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
