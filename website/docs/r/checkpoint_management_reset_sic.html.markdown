---
layout: "checkpoint"
page_title: "checkpoint_management_reset_sic"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-reset-sic"
description: |-
This resource allows you to execute Check Point Reset Sic.
---

# checkpoint_management_reset_sic

This resource allows you to execute Check Point Reset Sic.

## Example Usage


```hcl
resource "checkpoint_management_reset_sic" "example" {
  name = "gw1"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Gateway, cluster member or Check Point host unique identifier. 
* `name` - (Optional) Gateway, cluster member or Check Point host name. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

