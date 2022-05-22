---
layout: "checkpoint"
page_title: "checkpoint_management_trusted-client"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-trusted-client"
description: |-
Use this data source to get information on an existing Check Point Trusted Client.
---

# Data source: checkpoint_management_trusted_client

Use this data source to get information on an existing Check Point Trusted Client.

## Example Usage


```hcl
resource "checkpoint_management_trusted_client" "trustedClient" {
  name = "New TrustedClient 1"
  ipv4_address = "192.168.2.1"
}

data "checkpoint_management_trusted_client" "data_trusted_client" {
  name = "${checkpoint_management_trusted_client.trustedClient.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
