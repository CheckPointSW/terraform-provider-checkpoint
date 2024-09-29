---
layout: "checkpoint"
page_title: "checkpoint_management_delete_custom_trusted_ca_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-delete-custom-trusted-ca-certificate"
description: |-
This resource allows you to execute Check Point Delete Custom Trusted Ca Certificate.
---

# checkpoint_management_delete_custom_trusted_ca_certificate

This resource allows you to execute Check Point Delete Custom Trusted Ca Certificate.

## Example Usage


```hcl
resource "checkpoint_management_delete_custom_trusted_ca_certificate" "del" {
  name = "custom-trusted-ca-cert-object"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object name. 
* `name` - (Optional) Object name.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

