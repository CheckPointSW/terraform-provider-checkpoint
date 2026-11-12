---
layout: "checkpoint"
page_title: "checkpoint_management_install_lsm_policy"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-install-lsm-policy"
description: |-
This resource allows you to execute Check Point Install Lsm Policy.
---

# checkpoint_management_install_lsm_policy

This resource allows you to execute Check Point Install Lsm Policy.

## Example Usage


```hcl
resource "checkpoint_management_install_lsm_policy" "example" {
  targets = ["lsm_gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

