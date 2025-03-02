---
layout: "checkpoint"
page_title: "checkpoint_management_ise_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-ise-data-center-server"
description: |- This resource allows you to execute Check Point Cisco ISE data center server.
---

# Resource: checkpoint_management_ise_data_center_server

This resource allows you to execute Check Point Cisco ISE Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_ise_data_center_server" "testIse" {
  name      = "MyIse"
  username  = "USERNAME"
  password  = "PASSWORD"
  hostnames = ["host1", "host2"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `hostnames` - (**Required**) IP address or hostname of the Cisco ISE administration Node(s).
* `username` - (**Required**) Username of the ISE administrator.
* `password` - (Optional)  Password of the ISE administrator.
* `password_base64` - (Optional) Password of the ISE administrator encoded in Base64.
* `certificate_fingerprint` - (Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
