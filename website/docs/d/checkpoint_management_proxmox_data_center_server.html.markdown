---
layout: "checkpoint"
page_title: "checkpoint_management_proxmox_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-proxmox-data-center-server"
description: |- 
  Use this data source to get information on an existing Proxmox Data Center Server.
---

# Resource: checkpoint_management_proxmox_data_center_server

Use this data source to get information on an existing Proxmox Data Center Server.

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

data "checkpoint_management_proxmox_data_center_server" "data_proxmox_data_center_server" {
  name = "${checkpoint_management_proxmox_data_center_server.testProxmox.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `data_center_type` - Data center type.
* `properties` - Data center properties. properties blocks are documented below.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `certificate_fingerprint` - Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.
* `unsafe_auto_accept` - When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`properties` supports the following:
* `name`
* `value`