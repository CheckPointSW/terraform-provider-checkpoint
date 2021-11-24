---
layout: "checkpoint"
page_title: "checkpoint_management_azure_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-azure-data-center-server"
description: |- This resource allows you to execute Check Point azure data center server.
---

# Resource: checkpoint_management_azure_data_center_server

This resource allows you to execute Check Point Azure Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_azure_data_center_server" "testAzure" {
  name = "myAzure"
  authentication_method = "user-authentication"
  username         = "MY-KEY-ID"
  password     = "MY-SECRET-KEY"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `authentication_method` - (Required) user-authentication Uses the Azure AD User to authenticate. service-principal-authentication Uses the Service Principal to authenticate.
* `username` - (Required for authentication-method: user-authentication) An Azure Active Directory user Format <username>@<domain>. Required for authentication-method: user-authentication.
* `password` - (Optional)  Password of the Azure account. Required for authentication-method: user-authentication.
* `password_base64` - (Optional) Password of the Azure account encoded in Base64. Required for authentication-method: user-authentication.
* `application_id` - (Required for authentication-method: service-principal-authentication) The Application ID of the Service Principal, in UUID format. Required for authentication-method: service-principal-authentication.
* `application_key` - (Required for authentication-method: service-principal-authentication) The key created for the Service Principal. Required for authentication-method: service-principal-authentication.
* `directory_id` - (Required for authentication-method: service-principal-authentication) The Directory ID of the Azure AD, in UUID format. Required for authentication-method: service-principal-authentication.
* `environment` - (Optional) Select the Azure Cloud Environment.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
