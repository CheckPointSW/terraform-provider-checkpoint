---
layout: "checkpoint"
page_title: "checkpoint_management_nutanix_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-nutanix-data-center-server"
description: |- This resource allows you to execute Check Point nutanix data center server.
---

# Resource: checkpoint_management_nutanix_data_center_server

This resource allows you to execute Check Point Nutanix Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_nutanix_data_center_server" "testNutanix" {
  name = "MY-NUTANIX"
  hostname = "127.0.0.1"
  username = "admin"
  password = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `hostname` - (**Required**) IP Address or hostname of the Nutanix Prism Central server.
* `username` - (**Required**) Username of the Nutanix Prism Central server.
* `password` - (**Required**) Password of the Nutanix Prism Central server.
* `certificate_fingerprint` - (Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `data_center_type` - Data center type.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `properties` - Data center properties. properties blocks are documented below.


`properties` supports the following:

* `name`
* `value`