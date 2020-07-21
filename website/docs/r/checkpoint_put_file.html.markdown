---
layout: "checkpoint"
page_title: "checkpoint_put_file"
sidebar_current: "docs-checkpoint-gaia-resource-checkpoint-put-file"
description: |-
  This resource allows you to add a new file to a Check Point machine.
---

# checkpoint_put_file

This resource allows you to add a new file to a Check Point machine.

## Example Usage


```hcl
resource "checkpoint_put_file" "put_file1" {
      file_name = "/path/to/file1/terrafile1.txt"
      text_content = "It's a terrafile!"
      override = true
}

resource "checkpoint_put_file" "put_file1" {
      file_name = "/path/to/file2/terrafile2.txt"
      text_content = "It's a terrafile!"
}
```

## Argument Reference

The following arguments are supported:

* `file_name` - (Required) Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided.
* `text_content` - (Required) Content to add to the new file. 
* `override` - (Optional) If the file already exists, indicates whether to overwrite it.













