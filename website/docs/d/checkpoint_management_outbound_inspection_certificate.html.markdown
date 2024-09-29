---
layout: "checkpoint"
page_title: "checkpoint_management_outbound_inspection_certificate"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-outbound-inspection-certificate"
description: |-
Use this data source to get information on an existing Check Point Outbound Inspection Certificate.
---

# Data Source: checkpoint_management_outbound_inspection_certificate

Use this data source to get information on an existing Check Point Outbound Inspection Certificate.

## Example Usage


```hcl
resource "checkpoint_management_outbound_inspection_certificate" "example" {
  name            = "cert2"
  issued_by       = "www.checkpoint.com"
  base64_password = "bXlfcGFzc3dvcmQ="
  valid_from      = "2021-04-17"
  valid_to        = "2028-04-17"
}
data "checkpoint_management_outbound_inspection_certificate" "data" {
  uid = "${checkpoint_management_outbound_inspection_certificate.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `issued_by` -  The DN (Distinguished Name) of the certificate. 
* `base64_certificate` -  Certificate file encoded in base64.
* `base64_public_certificate` - Public Certificate file encoded in base64 (pem format).
* `valid_from` -  The date, from which the certificate is valid. Format: YYYY-MM-DD. 
* `valid_to` - The certificate expiration date. Format: YYYY-MM-DD. 
* `is_default` -  Is the certificate the default certificate. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

