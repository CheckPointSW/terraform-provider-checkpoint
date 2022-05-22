---
layout: "checkpoint"
page_title: "checkpoint_management_set_outbound_inspection_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-outbound-inspection-certificate"
description: |-
This resource allows you to execute Check Point Set Outbound Inspection Certificate.
---

# checkpoint_management_set_outbound_inspection_certificate

This resource allows you to execute Check Point Set Outbound Inspection Certificate.

## Example Usage


```hcl
resource "checkpoint_management_set_outbound_inspection_certificate" "example" {
  issued_by = "www.checkpoint.com"
  base64_password = "bXlfcGFzc3dvcmQ="
  valid_from = "2021-04-17"
  valid_to = "2028-04-17"
}
```

## Argument Reference

The following arguments are supported:

* `issued_by` - (Required) The DN (Distinguished Name) of the certificate. 
* `base64_password` - (Required) Password (encoded in Base64 with padding) for the certificate file. 
* `valid_from` - (Required) The date, from which the certificate is valid. Format: YYYY-MM-DD. 
* `valid_to` - (Required) The certificate expiration date. Format: YYYY-MM-DD. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

