---
layout: "checkpoint"
page_title: "checkpoint_management_add_api_key"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-add-api-key"
description: |-
This resource allows you to execute Check Point Add Api Key.
---

# checkpoint_management_add_api_key

This resource allows you to execute Check Point Add Api Key.

## Example Usage


```hcl
resource "checkpoint_management_add_api_key" "example" {
  admin_name = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `admin_uid` - (Required) Administrator uid to generate API key for. 
* `admin_name` - (Required) Administrator name to generate API key for. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

