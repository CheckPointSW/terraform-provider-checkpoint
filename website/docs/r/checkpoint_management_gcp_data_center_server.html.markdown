---
layout: "checkpoint"
page_title: "checkpoint_management_gcp_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-gcp-data-center-server"
description: |- This resource allows you to execute Check Point gcp data center server.
---

# Resource: checkpoint_management_gcp_data_center_server

This resource allows you to execute Check Point Google Cloud Platform Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_gcp_data_center_server" "testGcp" {
  name                  = "myGcp"
  authentication_method = "key-authentication"
  private_key           = "MYKEY.json"
  ignore_warnings       = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `authentication_method` - (Required) key-authentication Uses the Service Account private key file to authenticate. vm-instance-authentication Uses the Service Account VM Instance to authenticate. This option requires the Security Management Server deployed in a GCP, and runs as a Service Account with the required permissions.
* `private_key` - (Required for authentication-method: key-authentication) A Service Account Key JSON file, encoded in base64. Required for authentication-method:key-authentication.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
