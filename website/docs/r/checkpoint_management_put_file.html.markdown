---
layout: "checkpoint"
page_title: "checkpoint_management_put_file"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-put-file"
description: |-
This resource allows you to execute Check Point Put File.
---

# checkpoint_management_put_file

This resource allows you to execute Check Point Put File.

## Example Usage


```hcl
resource "checkpoint_management_put_file" "put_file" {
  file_path = "/home/admin/"
  file_name = "myfile.txt"
  file_content = "first line\n second line"
  targets = ["corporate-gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.
* `file_content` - (Optional) Text file content.
* `file_name` - (Optional) Text file name.
* `file_path` - (Optional) Text file target path. 
* `comments` - (Optional) Comments string. 


## How To Use
Make sure this command will be executed in the right execution order.