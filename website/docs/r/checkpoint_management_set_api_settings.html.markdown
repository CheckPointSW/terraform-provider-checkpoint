---
layout: "checkpoint"
page_title: "checkpoint_management_set_api_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-api-settings"
description: |-
This resource allows you to execute Check Point Set Api Settings.
---

# checkpoint_management_set_api_settings

This resource allows you to execute Check Point Set Api Settings.

## Example Usage


```hcl
resource "checkpoint_management_set_api_settings" "example" {
  accepted_api_calls_from = "All IP addresses"
}
```

## Argument Reference

The following arguments are supported:

* `accepted_api_calls_from` - (Optional) Clients allowed to connect to the API Server. 
* `automatic_start` - (Optional) MGMT API will start after server will start. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

