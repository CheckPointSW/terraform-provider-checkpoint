---
layout: "checkpoint"
page_title: "checkpoint_management_update_provisioned_satellites"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-update-provisioned-satellites"
description: |-
This resource allows you to execute Check Point Update Provisioned Satellites.
---

# checkpoint_management_update_provisioned_satellites

This resource allows you to execute Check Point Update Provisioned Satellites.

## Example Usage


```hcl
resource "checkpoint_management_update_provisioned_satellites" "example" {
  vpn_center_gateways = ["co_gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `vpn_center_gateways` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier. The targets should be a corporate gateways.vpn_center_gateways blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

