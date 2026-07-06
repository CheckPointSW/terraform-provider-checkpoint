---
layout: "checkpoint"
page_title: "checkpoint_management_nozomi_data_center_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-nozomi-data-center-server"
description: |- Use this data source to get information on an existing Check Point Nozomi data center server.
---

# Data Source: checkpoint_management_nozomi_data_center_server

Use this data source to get information on an existing Check Point Nozomi Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_nozomi_data_center_server" "testNozomi" {
  name = "MY-NOZOMI"
  hostname = "127.0.0.1"
  key_name = "MY-KEY-NAME"
  key_token = "MY-KEY-TOKEN"
}

data "checkpoint_management_nozomi_data_center_server" "data_nozomi_data_center_server" {
  name = "${checkpoint_management_nozomi_data_center_server.testNozomi.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `hostname` - (Computed) IP address or hostname of the Nozomi Guardian or CMC server.
* `key_name` - (Computed) API key name for Nozomi.
* `certificate_fingerprint` - (Computed) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Computed) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Computed) Collection of tag objects identified by the name or UID.
* `color` - (Computed) Color of the object.
* `comments` - (Computed) Comments string.
