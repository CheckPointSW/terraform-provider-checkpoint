---
layout: "checkpoint"
page_title: "checkpoint_management_vmware_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-vmware-data-center-server"
description: |- Use this data source to get information on an existing VMware Data Center Server.
---

# Resource: checkpoint_management_vmware_data_center_server

Use this data source to get information on an existing VMware Data Center Server.
### Note:
* NSX-V (nsx type) is deprecated from R82 and above
* Global NSX-T supported from R82 and above

## Example Usage

```hcl
resource "checkpoint_management_vmware_data_center_server" "testVMware" {
  name     = "MyVMware"
  type     = "vcenter" # "nsx" or "nsxt" or "globalnsxt"
  username = "USERNAME"
  password = "PASSWORD"
  hostname = "HOSTNAME"
}


data "checkpoint_management_vmware_data_center_server" "data_vmware_data_center_server" {
  name = "${checkpoint_management_vmware_data_center_server.testVMware.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
