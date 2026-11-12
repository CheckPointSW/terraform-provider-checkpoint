---
layout: "checkpoint"
page_title: "checkpoint_management_command_lock_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-lock-object"
description: |-
This resource allows you to execute Check Point Lock Object.
---

# checkpoint_management_command_lock_object

This resource allows you to execute Check Point Lock Object.

## Example Usage


```hcl
resource "checkpoint_management_host" "example" {
  name = "test"
  ipv4_address = "1.1.1.1"
}

resource "checkpoint_management_command_lock_object" "lock" {
  name = "test"
  type = "host"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Opitioal) Object name.
* `uid` - (Optional) Object unique identifier. When using uid, there is no need to use name/type parameters.
* `type` - (Optional) Object type.
* `layer` - (Optional) Object layer, need to specify the layer if the object is rule/section and uid is not supplied.