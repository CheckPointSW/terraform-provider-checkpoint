---
layout: "checkpoint"
page_title: "checkpoint_management_opsec_trusted_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-opsec-trusted-ca"
description: |-
Use this data source to get information on an existing Check Point Opsec Trusted Ca.
---

# checkpoint_management_opsec_trusted_ca

Use this data source to get information on an existing Check Point Opsec Trusted Ca.

## Example Usage


```hcl
resource "checkpoint_management_opsec_trusted_ca" "example" {
  name = "opsec_ca"
  base64_certificate = "MIICwjCCAaqgAwIBAgIILdexblpVEMIwDQYJKoZIhvcNAQELBQAwGDEWMBQGA1UEAxMNd3d3Lm9wc2VjLmNvbTAeFw0yMzA2MjUwOTE3MDBaFw0yNTAzMzExNjAwMDBaMBgxFjAUBgNVBAMTDXd3dy5vcHNlYy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCjpqCxDaVg+I1b+wqnmjjYtL3v7Tlu/YpMbsKnv+M1gRz6QFUOoSVnxKLo0A7Y4kCqa1OPcHO/LtXuok43F1YZPVKm3xWpY8FmqGqf5ZuGmSwm1HPObcMjwGOyFgwpwEDF5e0UMZ7xtJF8BZ5KKBh3ZfQ1FbmbVqSUPcmOi+NE4JspPlHxX+m6es/yeSGR1A2ezKY7KePTlwVtDe8hiLrYyKG92nka5rkD1QyEIVJ0W5wrnU4nGEDIHeOfT09zroQxaNLkb51sl4Tog/qw+EraVGIBe/iFnSJoDF37i2mLJqI/t8bel+aGDAxgMx1pO85OClgjPSWL0UIXGI2xrR+JAgMBAAGjEDAOMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAHTs1AutAmSLHF2KRLJtrRNkso0lMyA7XI7k1TNpTk7TCZLNY0VbUliGbcl+POH4EG8ARUrftnwRDCTBd2BdJTqG2CyNADi+bw8aLvbxok7KH0GlQvGjyfq+sHK12wTl4ULNyYoAPZ01GhXOvkobROdSyjxvBVhxdVo90kj7mHFv3N83huNhfstDFUBcQCmMkbLuzDUZrl2a1OtqlOdNC6mNvb7Jq9W9vRxGA514e7jqyoM+PwHu5fILx/jmGT8suOUnvbtcDdFhjqixAPer6uSPR0CSbiJvuDy72DPH5mjZK5dQKewNYOZ/BQEsRIBe+Q6eGAoJqi+cD63cwlw0DCc="
}

data "checkpoint_management_opsec_trusted_ca" "data1" {
  name = "${checkpoint_management_opsec_trusted_ca.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier. 
* `base64_certificate` -  Certificate file encoded in base64. 
* `automatic_enrollment` - Certificate automatic enrollment.automatic_enrollment blocks are documented below.
* `retrieve_crl_from_http_servers` - Whether to retrieve Certificate Revocation List from http servers. 
* `retrieve_crl_from_ldap_servers` - Whether to retrieve Certificate Revocation List from ldap servers. 
* `cache_crl` - Cache Certificate Revocation List on the Security Gateway. 
* `crl_cache_method` - Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period. 
* `crl_cache_timeout` -  When to fetch new Certificate Revocation List (in minutes). 
* `allow_certificates_from_branches` - Allow only certificates from listed branches. 
* `branches` - Branches to allow certificates from. Required only if "allow-certificates-from-branches" set to "true".branches blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 

`automatic_enrollment` supports the following:

* `automatically_enroll_certificate` -  Whether to automatically enroll certificate. 
* `protocol` -  Protocol that communicates with the certificate authority. Available only if "automatically-enroll-certificate" parameter is set to true. 
* `scep_settings` -  Scep protocol settings. Available only if "protocol" is set to "scep".scep_settings blocks are documented below.
* `cmpv1_settings` -  Cmpv1 protocol settings. Available only if "protocol" is set to "cmpv1".cmpv1_settings blocks are documented below.
* `cmpv2_settings` - Cmpv2 protocol settings. Available only if "protocol" is set to "cmpv1".cmpv2_settings blocks are documented below.


`scep_settings` supports the following:

* `ca_identifier` - Certificate authority identifier. 
* `url` - Certificate authority URL. 


`cmpv1_settings` supports the following:

* `direct_tcp_settings` - Direct tcp transport layer settings.direct_tcp_settings blocks are documented below.


`cmpv2_settings` supports the following:

* `transport_layer` - Transport layer. 
* `direct_tcp_settings` - Direct tcp transport layer settings.direct_tcp_settings blocks are documented below.
* `http_settings` - Http transport layer settings.http_settings blocks are documented below.


`direct_tcp_settings` supports the following:
* `ip_address` - IP Address
* `port` - Port number. 


`direct_tcp_settings` supports the following:
* `ip_address` - IP Address
* `port` - Port number. 


`http_settings` supports the following:

* `url` -  Certificate authority URL. 
