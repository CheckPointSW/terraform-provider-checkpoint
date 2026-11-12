---
layout: "checkpoint"
page_title: "checkpoint_management_servicenow_cmdb_data_center_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-servicenow-cmdb-data-center-server"
description: |- Use this data source to get information on an existing Check Point ServiceNow CMDB data center server.
---

# Data Source: checkpoint_management_servicenow_cmdb_data_center_server

Use this data source to get information on an existing Check Point ServiceNow CMDB Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_servicenow_cmdb_data_center_server" "testServicenowCmdb" {
  name = "MY-SERVICENOW-CMDB"
  hostname = "instance.service-now.com"
  username = "admin"
  password = "MY-PASSWORD"
}

data "checkpoint_management_servicenow_cmdb_data_center_server" "data_servicenow_cmdb_data_center_server" {
  name = "${checkpoint_management_servicenow_cmdb_data_center_server.testServicenowCmdb.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `hostname` - (Computed) Instance hostname of the ServiceNow instance (e.g. instance.service-now.com).
* `username` - (Computed) User name for ServiceNow instance.
* `certificate_fingerprint` - (Computed) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Computed) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Computed) Collection of tag objects identified by the name or UID.
* `color` - (Computed) Color of the object.
* `comments` - (Computed) Comments string.
