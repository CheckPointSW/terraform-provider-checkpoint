---
layout: "checkpoint"
page_title: "checkpoint_management_connect_cloud_services"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-connect-cloud-services"
description: |-
This resource allows you to execute Check Point Connect Cloud Services.
---

# checkpoint_management_connect_cloud_services

This resource allows you to execute Check Point Connect Cloud Services.

## Example Usage

```hcl
resource "checkpoint_management_connect_cloud_services" "example" {
  auth_token = "aHR0cHM6Ly9kZXYtY2xvdWRpbmZyYS1ndy5rdWJlMS5pYWFzLmNoZWNrcG9pbnQuY29tL2FwcC9tYWFzL2FwaS92Mi9tYW5hZ2VtZW50cy9hZmJlYWRlYS04Y2U2LTRlYTUtOTI4OS00ZTQ0N2M0ZjgyMTMvY2xvdWRBY2Nlc3MvP290cD02ZWIzNThlOS1hMzkxLTQxOGQtYjlmZi0xOGIxOTQwOGJlN2Y="
}
```

## Argument Reference

The following arguments are supported:

* `auth_token` - (Required) Copy the authentication token from the Smart-1 cloud service hosted in the Infinity Portal. 
* `status` - Status of the connection to the Infinity Portal.
* `connected_at` - The time of the connection between the Management Server and the Infinity Portal. connected_at is documented below.
* `management_url` - The Management Server's public URL.

`connected_at` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

