---
layout: "checkpoint"
page_title: "checkpoint_management_guardicore_data_center_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-guardicore-data-center-server"
description: |- This resource allows you to execute Check Point Akamai Guardicore data center server.
---

# Resource: checkpoint_management_guardicore_data_center_server

This resource allows you to execute Check Point Akamai Guardicore Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_guardicore_data_center_server" "testGuardicore" {
  name = "MY-GUARDICORE"
  hostname = "127.0.0.1"
  username = "admin"
  password = "MY-PASSWORD"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `hostname` - (Required) IP Address or hostname of the Guardicore Centra management server.
* `username` - (Required) Username for Guardicore Centra.
* `password` - (Required) Password for Guardicore Centra.
* `certificate_fingerprint` - (Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
