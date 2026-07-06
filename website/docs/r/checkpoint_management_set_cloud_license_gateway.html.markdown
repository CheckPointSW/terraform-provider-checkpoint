---
layout: "checkpoint"
page_title: "checkpoint_management_set_cloud_license_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-cloud-license-gateway"
description: |-
This resource allows you to execute Check Point Set Cloud License Gateway.
---

# checkpoint_management_set_cloud_license_gateway

This resource allows you to execute Check Point Set Cloud License Gateway.

## Example Usage


```hcl
resource "checkpoint_management_set_cloud_license_gateway" "example" {
  gateway = "Gateway_0.0.0.0"
  enable_auto_distribution = true
}
```

## Argument Reference

The following arguments are supported:

* `gateway` - (Required) Security gateway name or UID to set. 
* `enable_auto_distribution` - (Required) Enable or disable auto distribution of cloud licenses for the specified gateway. 
* `domain` - (Optional) Domain name or UID for the gateway. Required when running from MDS context.

