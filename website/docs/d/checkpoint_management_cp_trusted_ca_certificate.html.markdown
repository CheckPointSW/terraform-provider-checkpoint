---
layout: "checkpoint"
page_title: "checkpoint_management_cp_trusted_ca_certificate"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cp-trusted-ca-certificate"
description: |-
Use this data source to get information on an existing Check Point  Cp Trusted Ca Certificate.
---

# checkpoint_management_cp_trusted_ca_certificate

**This resource allows you to execute** Check Point Set Cp Trusted Ca Certificate.

## Example Usage


```hcl
resource "checkpoint_management_set_cp_trusted_ca_certificate" "example" {
  uid  = "24283511-e2a9-46da-a9eb-d22ec839ab1a"
  status = "disabled"
}

data "checkpoint_management_cp_trusted_ca_certificate" "data" {
  uid = "${checkpoint_management_command_set_cp_trusted_ca_certificate.cert1.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `status` -  Indicates whether the trusted CP CA certificate is enabled/disabled. 
* `added_by` - By whom the certificate was added.
* `base64_certificate` - The certificate in base64.
* `base64-public_certificate` - Public Certificate file encoded in base64 (pem format).
* `issued_by` - Trusted CA certificate issued by.
* `issued_to` - Trusted CA certificate issued to.
* `status` - Indicates whether the trusted CP CA certificate is enabled/disabled.
* `tags` - Collection of tag identifiers.
* `valid_from` - Trusted CA certificate valid from date.
* `valid_to` - Trusted CA certificate valid to date.

`valid_from` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970

`valid_to` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

