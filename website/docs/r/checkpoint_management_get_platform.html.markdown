---
layout: "checkpoint"
page_title: "checkpoint_management_get_platform"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-get-platform"
description: |-
This resource allows you to execute Check Point Get Platform.
---

# checkpoint_management_get_platform

This resource allows you to execute Check Point Get Platform.

## Example Usage


```hcl
resource "checkpoint_management_get_platform" "example" {
  name = "gw1"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Gateway, cluster or Check Point host unique identifier.
* `name` - (Optional) Gateway, cluster or Check Point host name. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

