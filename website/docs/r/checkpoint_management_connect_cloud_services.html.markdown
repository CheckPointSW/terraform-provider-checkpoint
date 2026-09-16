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
* `gateways_onboarding_settings` - (Optional) Gateways on-boarding to Infinity Portal settings.gateways_onboarding_settings blocks are documented below.

`connected_at` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

`gateways_onboarding_settings` supports the following:

* `connection_method` - (Optional) Indicate whether Gateways will be connected to Infinity Portal automatically or only after policy installation.
* `details_level` - (Optional) The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation o...
* `enabled` - (Optional) Enable/Disable automatic connection of Security Gateways to Infinity Portal.
* `participant_gateways` - (Optional) Which Gateways will be connected to Infinity Portal.
* `specific_gateways` - (Optional) Selection of targets identified by the name or UID which will be on-boarded to the cloud. Configuration will be applied only when 'participant-gateway...


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  
