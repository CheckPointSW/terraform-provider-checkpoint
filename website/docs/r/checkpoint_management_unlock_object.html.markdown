---
layout: "checkpoint"
page_title: "checkpoint_management_command_unlock_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-unlock-object"
description: |-
This resource allows you to execute Check Point Unlock Object.
---

# checkpoint_management_command_unlock_object

This resource allows you to execute Check Point Unlock Object.

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

resource "checkpoint_management_command_unlock_object" "unlock" {
  name = "${checkpoint_management_command_lock_object.lock.name}"
  type = "${checkpoint_management_command_lock_object.lock.type}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Opitioal) Object name.
* `uid` - (Optional) Object unique identifier. When using uid, there is no need to use name/type parameters.
* `type` - (Optional) Object type.
* `layer` - (Optional) Object layer, need to specify the layer if the object is rule/section and uid is not supplied.