---
layout: "checkpoint"
page_title: "checkpoint_management_repository_script"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-repository-script"
description: |-
This resource allows you to execute Check Point Repository Script.
---

# checkpoint_management_repository_script

This resource allows you to execute Check Point Repository Script.

## Example Usage


```hcl
resource "checkpoint_management_repository_script" "example" {
  name = "New Script 1"
  script_body = "ls -l /"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `script_body` - (Optional) The entire content of the script. 
* `script_body_base64` - (Optional) The entire content of the script encoded in Base64. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
