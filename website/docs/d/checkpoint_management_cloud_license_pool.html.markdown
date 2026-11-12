---
layout: "checkpoint"
page_title: "checkpoint_management_cloud_license_pool"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cloud-license-pool"
description: |-
Use this data source to get information on an existing Check Point Cloud License Pool.
---

# Data Source: checkpoint_management_cloud_license_pool

Use this data source to get information on an existing Check Point Cloud License Pool.

## Example Usage

```hcl
data "checkpoint_management_cloud_license_pool" "data_test" {
    pool = "cloud-license-pool1"
}
```

## Argument Reference

The following arguments are supported:

* `pool` - (Required) Pool name.
* `ck` - (Optional) Certificate Key. Required to identify a specific pool when multiple pools share the same name.
* `cks` - (Computed) List of the licenses CKs (Certificate Keys) that belong to this license pool. cks blocks are documented below.
* `available_quota` - (Computed) The difference between the pool's total quota and the total cores quantity of the pool's subscribed Security Gateways.
* `default_pool` - (Computed) All new CloudGuard Gateways are automatically subscribed to the default license pool.
* `total_quota` - (Computed) A license pool total quota is the total quantity of all the virtual cores provided by all the Central Licenses in this pool.
* `subscribed_gateways` - (Computed) List of the subscribed CloudGuard Gateways of this license pool. subscribed_gateways blocks are documented below.


`cks` supports the following:

* `ck` - (Computed) The license CK (Certificate Key).
* `expired` - (Computed) Whether this CK is expired.

`subscribed_gateways` supports the following:

* `name` - (Computed) Gateway name.
* `uid` - (Computed) Gateway's unique identifier.
* `domain` - (Computed) Domain name.
* `used-quota` - (Computed) Cores quantity used by the gateway.