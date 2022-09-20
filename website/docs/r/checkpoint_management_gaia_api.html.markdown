---
layout: "checkpoint"
page_title: "checkpoint_management_command_gaia_api"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-gaia-api"
description: |-
This resource allows you to execute Check Point Gaia Api.
---

# Resource: checkpoint_management_command_gaia_api

This resource allows you to execute Check Point Gaia Api.

## Example Usage


```hcl
resource "checkpoint_management_command_gaia_api" "example" {
  target = "my_gateway"
  command_name = "show-hostname"
}
```

## Argument Reference

The following arguments are supported:

* `target` - (Required) Gateway-object-name or gateway-ip-address or gateway-UID. 
* `command_name` - (Required) Target's api command.
* `other_parameter` - (Optional) Other input parameters that gateway needs it. 
* `response_message` - Response's object from the target in json format.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

