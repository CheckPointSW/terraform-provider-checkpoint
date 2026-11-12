---
layout: "checkpoint"
page_title: "checkpoint_management_illumio_data_center_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-illumio-data-center-server"
description: |- Use this data source to get information on an existing Illumio data center server.
---

# Data Source: checkpoint_management_illumio_data_center_server

Use this data source to get information on an existing Illumio Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_illumio_data_center_server" "testIllumio" {
  name          = "MY-ILLUMIO"
  hostname      = "127.0.0.1"
  org_id        = 1234567
  auth_username = "api_user"
  secret        = "secret"
}

data "checkpoint_management_illumio_data_center_server" "data_illumio_data_center_server" {
  name = "${checkpoint_management_illumio_data_center_server.testIllumio.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `hostname` - IP address or hostname of the Illumio PCE server.
* `org_id` - Organization ID in the Illumio PCE.
* `auth_username` - Authentication username.
* `certificate_fingerprint` - SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - Whether the Data Center Server's certificate is trusted as-is.
* `tags` - Collection of tag objects identified by the name or UID.
* `color` - Color of the object.
* `comments` - Comments string.
