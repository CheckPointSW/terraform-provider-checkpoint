---
layout: "checkpoint"
page_title: "checkpoint_management_test_sic_status"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-test-sic-status"
description: |-
This resource allows you to execute Check Point Test Sic Status.
---

# checkpoint_management_test_sic_status

This resource allows you to execute Check Point Test Sic Status.

## Example Usage


```hcl
resource "checkpoint_management_test_sic_status" "example" {
  name = "gw1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Gateway, cluster member or Check Point host name. 
* `uid` - (Optional) Gateway, cluster member or Check Point host name.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

