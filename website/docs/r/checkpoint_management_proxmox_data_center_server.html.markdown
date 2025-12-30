---
layout: "checkpoint"
page_title: "checkpoint_management_proxmox_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-proxmox-data-center-server"
description: |- 
  This resource allows you to execute Check Point Proxmox data center server.
---

# Resource: checkpoint_management_proxmox_data_center_server

This resource allows you to execute Check Point Proxmox Data Center Server.

### Note:
Proxmox is supported from R82.10 and above


## Example Usage

```hcl
resource "checkpoint_management_proxmox_data_center_server" "testProxmox" {
  name     = "MyProxmox"
  token_id = "USER@REALM!TOKEN_NAME"
  secret   = "SECRET"
  hostname = "HOSTNAME"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name.
* `hostname` - (**Required**) IP Address or hostname of the Proxmox server.
* `token_id` - (**Required**) API Token Id. In the format of `<Username>@<Realm>!<Token-Name>`.
* `secret` - (**Required**) Secret token API.
* `certificate_fingerprint` - (Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - (Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `data_center_type` - Data center type.