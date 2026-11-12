---
layout: "checkpoint"
page_title: "Check Point: checkpoint_gaia_management_interface"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-management-interface"
description: |-
  Manage the Gaia management interface.
---

# Resource: checkpoint_gaia_management_interface

## Argument Reference

* `name` - (Required) Management interface name.
* `virtual_system_id` - (Optional) Virtual system used for read requests.
* `member_id` - (Optional) Member used for read requests on supported platforms.


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_management_interface.example 'settings'
```

To specify optional read context or filters (`virtual_system_id`, `member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
