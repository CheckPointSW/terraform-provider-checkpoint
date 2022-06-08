---
layout: "checkpoint"
page_title: "checkpoint_management_interoperable_device"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-interoperable-device"
description: |-
Use this data source to get information on an existing Check Point Interoperable Device.
---

# Data Source: checkpoint_management_interoperable_device

Use this data source to get information on an existing Check Point Interoperable Device.

## Example Usage


```hcl
resource "checkpoint_management_interoperable_device" "example" {
  name = "NewInteroperableDevice"
  ipv4_address = "192.168.1.6"
}

data "checkpoint_management_interoperable_device" "data_interoperable_device" {
  name = "${checkpoint_management_interoperable_device.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.