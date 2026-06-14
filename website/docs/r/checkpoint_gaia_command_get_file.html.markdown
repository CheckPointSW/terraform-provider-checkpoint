---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_get_file"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-get-file"
description: |-
This resource allows you to execute Check Point Get File.
---

# checkpoint_gaia_command_get_file

This resource allows you to execute Check Point Get File.

## Example Usage


```hcl
# Step 1: create the file on the remote machine
resource "checkpoint_gaia_command_put_file" "put_setup" {
  file_name    = "/home/admin/file.txt"
  text_content = "example content"
  override     = true
}

# Step 2: retrieve the file
resource "checkpoint_gaia_command_get_file" "example" {
  file_name = "/home/admin/file.txt"

  depends_on = [checkpoint_gaia_command_put_file.put_setup]
}
```

## Argument Reference

The following arguments are supported:

* `file_name` - (Required) Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

