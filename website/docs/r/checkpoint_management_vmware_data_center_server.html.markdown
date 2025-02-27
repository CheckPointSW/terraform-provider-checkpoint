---
layout: "checkpoint"
page_title: "checkpoint_management_vmware_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-vmware-data-center-server"
description: |- This resource allows you to execute Check Point vmware data center server.
---

# Resource: checkpoint_management_vmware_data_center_server

This resource allows you to execute Check Point VMware Data Center Server.

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
```
<br><br>
```hcl
resource "checkpoint_management_vmware_data_center_server" "testNsxt" {
  name     = "MyNSXT"
  type     = "nsxt"
  username = "USERNAME"
  password = "PASSWORD"
  hostname = "HOSTNAME"
  policy_mode = false
  import_vms = false
}
```
## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `type` - (**Required**) Object type. ~~nsx~~ or nsxt or globalnsxt or vcenter.
* `hostname` - (**Required**) IP Address or hostname of the VMware server.
* `username` - (**Required**) Username of the VMware server.
* `password` - (Optional)  Password of the VMware server.
* `password_base64` - (Optional) Password of the VMware server encoded in Base64.
* `policy_mode` - (Optional) for **NSX-T** only and false at default.<br>When set to false, the Data Center Server will use Manager Mode APIs. <br>When set to true, the Data Center Server will use Policy Mode APIs.
* `import_vms`  - (Optional) for **NSX-T** only and false at default.<br>When set to true, the Data Center Server will import Virtual Machines as well.<br>This feature will create additional API requests toward NSX-T manager.<br><u>Note</u>: importing Virtual Machines can only be enabled while using Policy Mode APIs.
* `certificate_fingerprint` - (Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
