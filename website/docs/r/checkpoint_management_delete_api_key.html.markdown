---
layout: "checkpoint"
page_title: "checkpoint_management_delete_api_key"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-api-key"
description: |-
This resource allows you to execute Check Point Delete Api Key.
---

# checkpoint_management_delete_api_key

This resource allows you to execute Check Point Delete Api Key.

## Example Usage


```hcl
resource "checkpoint_management_delete_api_key" "example" {
  api_key = "eea3be76f4a8eb740ee872bcedc692748ff256a2d21c9ffd2754facbde046d00"
}
```

## Argument Reference

The following arguments are supported:

* `api_key` - (Required) API key to be deleted. 
* `admin_uid` - (Required) Administrator uid to generate API key for. 
* `admin_name` - (Required) Administrator name to generate API key for. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

