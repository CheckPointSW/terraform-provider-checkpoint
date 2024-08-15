---
layout: "checkpoint"
page_title: "checkpoint_management_trusted_ca_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-trusted-ca-settings"
description: |-
Use this data source to get information on an existing Check Point  Trusted Ca Settings.
---

# checkpoint_management_trusted_ca_settings

Use this data source to get information on an existing Check Point Trusted Ca Settings.

## Example Usage
```hcl
resource "checkpoint_management_command_set_trusted_ca_settings" "settings" {
  automatic_update = "false"
}

data "checkpoint_management_trusted_ca_settings" "data1" {
  depends_on = [checkpoint_management_command_set_trusted_ca_settings.settings]
}
```

## Argument Reference

The following arguments are supported:

* `automatic_update` -  Whether the trusted CAs package should be updated automatically. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

