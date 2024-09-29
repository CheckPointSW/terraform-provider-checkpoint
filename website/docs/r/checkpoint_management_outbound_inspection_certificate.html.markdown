---
layout: "checkpoint"
page_title: "checkpoint_management_outbound_inspection_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-outbound-inspection-certificate"
description: |-
This resource allows you to execute Check Point Outbound Inspection Certificate.
---

# checkpoint_management_outbound_inspection_certificate

This resource allows you to execute Check Point Outbound Inspection Certificate.

## Example Usage


```hcl
resource "checkpoint_management_outbound_inspection_certificate" "example" {
  name            = "cert2"
  issued_by       = "www.checkpoint.com"
  base64_password = "bXlfcGFzc3dvcmQ="
  valid_from      = "2021-04-17"
  valid_to        = "2028-04-17"
  
}
```

## Argument Reference

The following arguments are supported:

* `issued_by` - (Required) The DN (Distinguished Name) of the certificate. 
* `base64_password` - (Required) Password (encoded in Base64 with padding) for the certificate file. 
* `valid_from` - (Required) The date, from which the certificate is valid. Format: YYYY-MM-DD. 
* `valid_to` - (Required) The certificate expiration date. Format: YYYY-MM-DD. 
* `name` - (Optional) Object name.
* `is_default` - (Optional) Is the certificate the default certificate. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `base64_certificate` -  Certificate file encoded in base64.
* `base64_public_certificate` - Public Certificate file encoded in base64 (pem format).