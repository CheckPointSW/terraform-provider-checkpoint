---
layout: "checkpoint"
page_title: "checkpoint_gaia_maestro_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-maestro-gateway"
description: |-
This resource allows you to execute Check Point Maestro Gateway.
---

# checkpoint_gaia_maestro_gateway

This resource allows you to execute Check Point Maestro Gateway.

## Example Usage


```hcl
resource "checkpoint_gaia_maestro_gateway" "example" {
  resource_id = ""
  security_group = 2
  description = "New Gateway Description"
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required) ID of Gateway to modify 
* `description` - (Optional) New Gateway description 
* `security_group` - (Optional) ID of a Security Group. If specified, the Gateway will be assigned to this Security Group,regardless of it's current assignment status. In case you want to unassign Gateway from Security Group, use 0 
* `include_pending_changes` - (Computed) If true, show pending topology. If false, show deployed topology 


## Import

Use the existing resource's `resource_id` as the import ID.

```shell
terraform import checkpoint_gaia_maestro_gateway.example '1'
```

IDs beginning with `{` must use the JSON-object form; JSON syntax is reserved for structured import IDs.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
