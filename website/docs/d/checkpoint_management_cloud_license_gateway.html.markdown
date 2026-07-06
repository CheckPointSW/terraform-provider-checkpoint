---
layout: "checkpoint"
page_title: "checkpoint_management_cloud_license_gateway"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cloud-license-gateway"
description: |-
Use this data source to get information on an existing Check Point Cloud License Gateway.
---

# Data Source: checkpoint_management_cloud_license_gateway

Use this data source to get information on an existing Check Point Cloud License Gateway.

## Example Usage

```hcl
data "checkpoint_management_cloud_license_gateway" "data_test" {
  gateway = "cloud-license-gateway1"
}
```

## Argument Reference

The following arguments are supported:

* `gateway` - (Optional) Security gateway name or UID.
* `domain` - (Optional) Domain name or UID of security gateway. Required when running from MDS context.
* `used_quota` - (Computed) The number of licenses used by the gateway.
* `enable_auto_distribution` - (Computed) Whether automatic distribution is enabled for this gateway.
* `cks` - (Computed) List of CKs assigned to the gateway.
* `pool` - (Computed) Pool information including name, total quota, and available quota. pool blocks are documented below.
* `scheduled_auto_distribution` - (Computed) Time until the next scheduled automatic distribution.


`pool` supports the following:

* `pool` - (Computed) A group of CloudGuard Central Licenses with the same valid contract blades.
* `available_quota` - (Computed) The difference between the pool's total quota and the total cores quantity of the pool's subscribed Security Gateways.
* `total_quota` - (Computed) A license pool total quota is the total quantity of all the virtual cores provided by all the Central Licenses in this pool.
