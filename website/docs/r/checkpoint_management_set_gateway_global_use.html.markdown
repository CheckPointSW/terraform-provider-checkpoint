---
layout: "checkpoint"
page_title: "checkpoint_management_set_gateway_global_use"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-gateway-global-use"
description: |-
This resource allows you to execute Check Point Set Gateway Global Use.
---

# checkpoint_management_set_gateway_global_use

This resource allows you to execute Check Point Set Gateway Global Use.

## Example Usage


```hcl
resource "checkpoint_management_set_gateway_global_use" "example" {
  target = "vpn_gw"
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Required) Indicates whether global use is enabled on the target. 
* `target` - (Required) On what target to execute this command. Target may be identified by its object name, or object unique identifier. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

