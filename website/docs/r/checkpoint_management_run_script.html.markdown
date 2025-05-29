---
layout: "checkpoint"
page_title: "checkpoint_management_run_script"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-run-script"
description: |-
  This resource allows you to execute Check Point Run Script.
---

# Resource: checkpoint_management_run_script

This command resource allows you to execute Check Point Run Script.

## Example Usage


```hcl
resource "checkpoint_management_run_script" "example" {
  script_name = "Script Example: List files under / dir"
  script = "ls -l /"
  targets = ["corporate-gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `script_name` - (Required) Script name. 
* `script` - (Required) Script body. 
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.
* `args` - (Optional) Script arguments. 
* `comments` - (Optional) Comments string. 
* `tasks` - (Computed) Collection of asynchronous task unique identifiers.
* `response` - Response message in JSON format.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

