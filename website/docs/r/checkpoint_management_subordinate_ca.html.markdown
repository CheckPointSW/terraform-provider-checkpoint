---
layout: "checkpoint"
page_title: "checkpoint_management_subordinate_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-subordinate-ca"
description: |-
 This resource allows you to execute Check Point Subordinate CA.
---

# checkpoint_management_subordinate_ca

This resource allows you to execute Check Point Subordinate CA.

## Example Usage


```hcl
resource "checkpoint_management_subordinate_ca" "example" {
  name = "TestSubordinateCa"
  base64_certificate = "MIICwjCCAaqgAwIBAgIILdexblpVEMIwDQYJKoZIhvcNAQELBQAwGDEWMBQGA1UEAxMNd3d3Lm9wc2VjLmNvbTAeFw0yMzA2MjUwOTE3MDBaFw0yNTAzMzExNjAwMDBaMBgxFjAUBgNVBAMTDXd3dy5vcHNlYy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCjpqCxDaVg+I1b+wqnmjjYtL3v7Tlu/YpMbsKnv+M1gRz6QFUOoSVnxKLo0A7Y4kCqa1OPcHO/LtXuok43F1YZPVKm3xWpY8FmqGqf5ZuGmSwm1HPObcMjwGOyFgwpwEDF5e0UMZ7xtJF8BZ5KKBh3ZfQ1FbmbVqSUPcmOi+NE4JspPlHxX+m6es/yeSGR1A2ezKY7KePTlwVtDe8hiLrYyKG92nka5rkD1QyEIVJ0W5wrnU4nGEDIHeOfT09zroQxaNLkb51sl4Tog/qw+EraVGIBe/iFnSJoDF37i2mLJqI/t8bel+aGDAxgMx1pO85OClgjPSWL0UIXGI2xrR+JAgMBAAGjEDAOMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAHTs1AutAmSLHF2KRLJtrRNkso0lMyA7XI7k1TNpTk7TCZLNY0VbUliGbcl+POH4EG8ARUrftnwRDCTBd2BdJTqG2CyNADi+bw8aLvbxok7KH0GlQvGjyfq+sHK12wTl4ULNyYoAPZ01GhXOvkobROdSyjxvBVhxdVo90kj7mHFv3N83huNhfstDFUBcQCmMkbLuzDUZrl2a1OtqlOdNC6mNvb7Jq9W9vRxGA514e7jqyoM+PwHu5fILx/jmGT8suOUnvbtcDdFhjqixAPer6uSPR0CSbiJvuDy72DPH5mjZK5dQKewNYOZ/BQEsRIBe+Q6eGAoJqi+cD63cwlw0DCc="
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `base64_certificate` - (Required) Certificate file encoded in base64.
* `automatic_enrollment` - (Optional) Certificate automatic enrollment. automatic_enrollment blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`automatic_enrollment` supports the following:

* `automatically_enroll_certificate` - (Optional) Whether to automatically enroll certificate. 
* `protocol` - (Optional) Protocol that communicates with the certificate authority. Available only if "automatically-enroll-certificate" parameter is set to true. 
* `scep_settings` - (Optional) Scep protocol settings. Available only if "protocol" is set to "scep". scep_settings blocks are documented below.
* `cmpv1_settings` - (Optional) Cmpv1 protocol settings. Available only if "protocol" is set to "cmpv1". cmpv1_settings blocks are documented below.
* `cmpv2_settings` - (Optional) Cmpv2 protocol settings. Available only if "protocol" is set to "cmpv1". cmpv2_settings blocks are documented below.


`scep_settings` supports the following:

* `ca_identifier` - (Optional) Certificate authority identifier. 
* `url` - (Optional) Certificate authority URL. 


`cmpv1_settings` supports the following:

* `direct_tcp_settings` - (Optional) Direct tcp transport layer settings. direct_tcp_settings blocks are documented below.


`cmpv2_settings` supports the following:

* `transport_layer` - (Optional) Transport layer. 
* `direct_tcp_settings` - (Optional) Direct tcp transport layer settings. direct_tcp_settings blocks are documented below.
* `http_settings` - (Optional) Http transport layer settings. http_settings blocks are documented below.


`direct_tcp_settings` supports the following:

* `ip_address` - (Optional) Certificate authority IP address.
* `port` - (Optional) Port number.


`http_settings` supports the following:

* `url` - (Optional) Certificate authority URL. 
