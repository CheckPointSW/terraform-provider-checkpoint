---
layout: "checkpoint"
page_title: "checkpoint_management_set_cp_trusted_ca_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-cp-trusted-ca-certificate"
description: |-
This resource allows you to execute Check Point Set Cp Trusted Ca Certificate.
---

# checkpoint_management_set_cp_trusted_ca_certificate

This resource allows you to execute Check Point Set Cp Trusted Ca Certificate.

## Example Usage


```hcl
resource "checkpoint_management_set_cp_trusted_ca_certificate" "example" {
  uid  = "24283511-e2a9-46da-a9eb-d22ec839ab1a"
  status = "disabled"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `status` - (Optional) Indicates whether the trusted CP CA certificate is enabled/disabled. 



## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

