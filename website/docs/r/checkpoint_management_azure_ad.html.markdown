---
layout: "checkpoint"
page_title: "checkpoint_management_azure_ad"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-azure-ad"
description: |-
This resource allows you to execute Check Point Azure Ad.
---

# Resource: checkpoint_management_azure_ad

This resource allows you to execute Check Point Azure Ad.

## Example Usage


```hcl
resource "checkpoint_management_azure_ad" "example" {
  name = "example"
  password = "123"
  user_authentication = "user-authentication"
  username = "example"
  application_id = "a8662b33-306f-42ba-9ffb-a0ac27c8903f"
  application_key = "EjdJ2JcNGpw3[GV8:PMN_s2KH]JhtlpO"
  directory_id = "19c063a8-3bee-4ea5-b984-e344asds37f7"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `authentication_method` - (Required) <b>user-authentication</b><br>Uses the Azure AD User to authenticate.<br><b>service-principal-authentication</b><br>Uses the Service Principal to authenticate. 
* `password` - (Required) Password of the Azure account.<br><p><font color="red">Required for authentication-method:</font></p>user-authentication. 
* `username` - (Required) An Azure Active Directory user Format<br>&lt;username&gt;@&lt;domain&gt;.<br><p><font color="red">Required for authentication-method:</font></p>user-authentication. 
* `application_id` - (Required) The Application ID of the Service Principal, in UUID format.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication. 
* `application_key` - (Required) The key created for the Service Principal.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication. 
* `directory_id` - (Required) The Directory ID of the Azure AD, in UUID format.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `properties` - Azure AD connection properties. properties blocks are documented below.

`properties` supports the following:

* `name`
* `value`