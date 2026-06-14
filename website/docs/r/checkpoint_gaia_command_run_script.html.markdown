---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_run_script"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-run-script"
description: |-
This resource allows you to execute Check Point Run Script.
---

# checkpoint_gaia_command_run_script

This resource allows you to execute Check Point Run Script.

## Example Usage


```hcl
resource "checkpoint_gaia_command_run_script" "example" {
  script = "date > /home/admin/run_script_example;echo Done"
}
```

## Argument Reference

The following arguments are supported:

* `script` - (Required) Script body 
* `description` - (Optional) Script description 
* `args` - (Optional) Script arguments, separated by space character. Note: don't send sensitive data on this parameter. 
* `environment_variables` - (Optional) Define environment variables to be used in the script, it's better to send sensitive data on environment variables since it's not stored. environment_variables blocks are documented below.


`environment_variables` supports the following:

* `name` - (Optional) Variable's name 
* `value` - (Optional) Variable's value 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

