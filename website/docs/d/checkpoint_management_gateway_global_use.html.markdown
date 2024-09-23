---
layout: "checkpoint"
page_title: "checkpoint_managemen_gateway_global_use"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-gateway-global-use"
description: |-
Use this data source to get information on an existing Check Point Set Gateway Global Use.
---

# checkpoint_management_set_gateway_global_use

Use this data source to get information on an existing Check Point Set Gateway Global Use.

## Example Usage


```hcl
resource "checkpoint_management_set_gateway_global_use" "example" {
  target = "vpn_gw"
  enabled = true
}
data "checkpoint_management_gateway_global_use" "data" {
  target = "${checkpoint_management_set_gateway_global_use.example.target}"
}
```

## Argument Reference

The following arguments are supported:

* `target` - (Required) On what target to execute this command. Target may be identified by its object name, or object unique identifier.
* `enabled` -  Indicates whether global use is enabled on the target. 
* `uid` - Object Identifier.
* `name` - Object Name.
* `domain` - data about doamin

`domain` supports the following:

* `uid` - Object Identifier.
* `name` - Object Name.
* `domain_type` - domain type.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

