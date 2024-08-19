---
layout: "checkpoint"
page_title: "checkpoint_management_set_trusted_ca_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-trusted-ca-settings"
description: |-
This resource allows you to execute Check Point Set Trusted Ca Settings.
---

# checkpoint_management_set_trusted_ca_settings

This resource allows you to execute Check Point Set Trusted Ca Settings.

## Example Usage
```hcl
resource "checkpoint_management_command_set_trusted_ca_settings" "settings" {
  automatic_update = "false"
}
```

## Argument Reference

The following arguments are supported:

* `automatic_update` - (Optional) Whether the trusted CAs package should be updated automatically. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

