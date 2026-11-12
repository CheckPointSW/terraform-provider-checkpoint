---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_put_file"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-put-file"
description: |-
This resource allows you to execute Check Point Put File.
---

# checkpoint_gaia_command_put_file

This resource allows you to execute Check Point Put File.

## Example Usage


```hcl
resource "checkpoint_gaia_command_put_file" "example" {
  file_name       = "file.txt"
  text_content    = <<-EOT
    first line
    second line
  EOT
  override        = true
  group_ownership = "config"
  user_ownership  = "admin"
  permissions     = 644
}
```

## Argument Reference

The following arguments are supported:

* `file_name` - (Required) Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided 
* `text_content` - (Optional) File content as string, for new line use \n 
* `override` - (Optional) overwrite file content 
* `group_ownership` - (Optional) Group file owner 
* `user_ownership` - (Optional) User file owner 
* `permissions` - (Optional) File permissions, provided in octal mode 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

