---
layout: "checkpoint"
page_title: "checkpoint_management_set_cloud_license_pool"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-cloud-license-pool"
description: |-
This resource allows you to execute Check Point Set Cloud License Pool.
---

# checkpoint_management_set_cloud_license_pool

This resource allows you to execute Check Point Set Cloud License Pool.

## Example Usage


```hcl
resource "checkpoint_management_set_cloud_license_pool" "example" {
  pool = "VE-NGTX"
}
```

## Argument Reference

The following arguments are supported:

* `pool` - (Required) Pool name. 
* `ck` - (Optional) Contract Key. Required to identify a specific pool when multiple pools share the same name. 
* `default_pool` - (Optional) Set pool to default. This value can only be changed from false to true. To disable the current default pool, you must set a different pool as the default. 
* `migrate_gateways` - (Optional) Move gateways from current default pool to the new default pool. Required when default-pool parameter is set to true. 
* `assigned_gateways` - (Optional) Attach security gateways to the pool. The attached gateways will use licenses from this pool.assigned_gateways blocks are documented below.


`assigned_gateways` supports the following:

* `gateway` - (Optional) Gateway name or uid. 
* `domain` - (Optional) Domain name or uid. Required when running from MDS context.

