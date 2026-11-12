---
layout: "checkpoint"
page_title: "checkpoint_management_external_trusted_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-external-trusted-ca"
description: |-
This resource allows you to execute Check Point External Trusted Ca.
---

# checkpoint_management_external_trusted_ca

This resource allows you to execute Check Point External Trusted Ca.

## Example Usage


```hcl
resource "checkpoint_management_external_trusted_ca" "example" {
  name = "external_ca"
  base64_certificate = "MIICujCCAaKgAwIBAgIIP1+IHWHbl0EwDQYJKoZIhvcNAQELBQAwFDESMBAGA1UEAxMJd3d3LnouY29tMB4XDTIzMTEyOTEyMzAwMFoXDTI0MTEyMDE2MDAwMFowFDESMBAGA1UEAxMJd3d3LnouY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoBreRGuq8u43GBog+ZaAnaR8ZF8cT2ppvtd3JoFmzTOQivLIt9sNtFYqEgHCtnNkKn9TRrxN14YscHgKIxfDSVlC9Rh0rrBvWgFqcm715Whr99Ogx6JbYFkusFWJarSejIFx4n6MM48MJxLdtCP6Hy1G2cj1BCiCHj4i3VIVaDE/aMkSqJbYEvf+vFqUWxY8/uEuKI/HGhI7mhUPW4NSGL0Oafz5eEFVsxqV5NA19/JJZ9NajSkyANnaNL5raxGV0oeqaE3JB3lSEZfWbH6mQsToUxxwIQfsZiIBozajDdTgP3Kn4SMY0b+I/WAWgfigMSDTAIR8J1sdzGXy2w2kqQIDAQABoxAwDjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBUgrHztHwC1E0mU5c4reMrHg+z+YRHrgJNHVIYQbL5I2TJHk9S3UZsynoMa1CO86rReOtR5xoGv4PCkyyOW+PNlWUtXF3tNgqWj/21+XzG4RBHPw89TaTxRCdo+MHX58fi07SIzKjmxfdkEi+7+HQEQluDZGViolrGBAw2rXq/SZ3q/11mNqlb5ZyqyOa2u1sBF1ApvG5a/FBRTaO8gaiNelRf0PGYkuV+1HhF2XyP8Qk565d+uxUH5M7eHF2PNyVk/r/36T+x+UMql9y9iizA0ekuAjXLok1xYl3Vw4S5zXCXYtNZLOVrs+plJb7IrlElyTOAbDFuPugh0medz7uZ"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `base64_certificate` - (Required) Certificate file encoded in base64. 
* `retrieve_crl_from_http_servers` - (Optional) Whether to retrieve Certificate Revocation List from http servers. 
* `crl_cache_method` - (Optional) Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period. 
* `crl_cache_timeout` - (Optional) When to fetch new Certificate Revocation List (in minutes). 
* `allow_certificates_from_branches` - (Optional) Allow only certificates from listed branches. 
* `branches` - (Optional) Branches to allow certificates from. Required only if "allow-certificates-from-branches" set to "true".branches blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
