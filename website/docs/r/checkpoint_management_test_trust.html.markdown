---
layout: "checkpoint"
page_title: "checkpoint_management_test_trust"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-test-trust"
description: |-
 This resource allows you to execute Check Point Test Trust.
---

# checkpoint_management_test_trust

This resource allows you to execute Check Point Test Trust.

## Example Usage


```hcl
resource "checkpoint_management_test_trust" "example" {
  name = "smb_daip"
  trust_method = "all"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Minimum Check Point password length.
* `trust_method` - (Optional) Trust method to use for establishing communication.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

