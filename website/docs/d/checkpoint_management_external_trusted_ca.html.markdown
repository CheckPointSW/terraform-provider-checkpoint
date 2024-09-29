---
layout: "checkpoint"
page_title: "checkpoint_management_external_trusted_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-external-trusted-ca"
description: |-
Use this data source to get information on an existing Check Point External Trusted Ca.
---

# checkpoint_management_external_trusted_ca

Use this data source to get information on an existing Check Point External Trusted Ca.

## Example Usage


```hcl
resource "checkpoint_management_external_trusted_ca" "example" {
  name = "external_ca"
  base64_certificate = "MIICujCCAaKgAwIBAgIIP1+IHWHbl0EwDQYJKoZIhvcNAQELBQAwFDESMBAGA1UEAxMJd3d3LnouY29tMB4XDTIzMTEyOTEyMzAwMFoXDTI0MTEyMDE2MDAwMFowFDESMBAGA1UEAxMJd3d3LnouY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoBreRGuq8u43GBog+ZaAnaR8ZF8cT2ppvtd3JoFmzTOQivLIt9sNtFYqEgHCtnNkKn9TRrxN14YscHgKIxfDSVlC9Rh0rrBvWgFqcm715Whr99Ogx6JbYFkusFWJarSejIFx4n6MM48MJxLdtCP6Hy1G2cj1BCiCHj4i3VIVaDE/aMkSqJbYEvf+vFqUWxY8/uEuKI/HGhI7mhUPW4NSGL0Oafz5eEFVsxqV5NA19/JJZ9NajSkyANnaNL5raxGV0oeqaE3JB3lSEZfWbH6mQsToUxxwIQfsZiIBozajDdTgP3Kn4SMY0b+I/WAWgfigMSDTAIR8J1sdzGXy2w2kqQIDAQABoxAwDjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBUgrHztHwC1E0mU5c4reMrHg+z+YRHrgJNHVIYQbL5I2TJHk9S3UZsynoMa1CO86rReOtR5xoGv4PCkyyOW+PNlWUtXF3tNgqWj/21+XzG4RBHPw89TaTxRCdo+MHX58fi07SIzKjmxfdkEi+7+HQEQluDZGViolrGBAw2rXq/SZ3q/11mNqlb5ZyqyOa2u1sBF1ApvG5a/FBRTaO8gaiNelRf0PGYkuV+1HhF2XyP8Qk565d+uxUH5M7eHF2PNyVk/r/36T+x+UMql9y9iizA0ekuAjXLok1xYl3Vw4S5zXCXYtNZLOVrs+plJb7IrlElyTOAbDFuPugh0medz7uZ"
}

data "checkpoint_management_external_trusted_ca" "data1" {
  uid = "${checkpoint_management_external_trusted_ca.obj1.id}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `base64_certificate` -  Certificate file encoded in base64. 
* `retrieve_crl_from_http_servers` - Whether to retrieve Certificate Revocation List from http servers. 
* `crl_cache_method` -  Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period. 
* `crl_cache_timeout` - When to fetch new Certificate Revocation List (in minutes). 
* `allow_certificates_from_branches` - Allow only certificates from listed branches. 
* `branches` -  Branches to allow certificates from. Required only if "allow-certificates-from-branches" set to "true".branches blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 
