---
layout: "checkpoint"
page_title: "checkpoint_management_lsm_run_script"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-lsm-run-script"
description: |-
This resource allows you to execute Check Point Lsm Run Script.
---

# checkpoint_management_lsm_run_script

This resource allows you to execute Check Point Lsm Run Script.

## Example Usage


```hcl
resource "checkpoint_management_lsm_run_script" "example" {
  targets = ["lsm_gateway"]
  script = "ls -l /"
}
```

## Argument Reference

The following arguments are supported:

* `script_base64` - (Optional) The entire content of the script encoded in Base64. 
* `script` - (Optional) The entire content of the script. 
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

