---
layout: "checkpoint"
page_title: "checkpoint_management_disconnect"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-disconnect"
description: |-
This resource allows you to execute Check Point Disconnect.
---

# checkpoint_management_disconnect

This resource allows you to execute Check Point Disconnect.

## Example Usage


```hcl
resource "checkpoint_management_disconnect" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `discard` - (Optional) Discard all changes committed during the session. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

