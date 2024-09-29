---
layout: "checkpoint"
page_title: "checkpoint_management_add_custom_trusted_ca_certificate"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-add-custom-trusted-ca-certificate"
description: |-
Use this data source to get information on an existing Check Point  Custom Trusted Ca Certificate.
---

# checkpoint_management_add_custom_trusted_ca_certificate

Use this data source to get information on an existing Check Point  Custom Trusted Ca Certificate.

## Example Usage


```hcl
resource "checkpoint_management_add_custom_trusted_ca_certificate" "example" {
  base64_certificate = "MIIEkzCCAnugAwIBAgIVAO5SRZQELwNNhWF+8st6ox9uXYgeMA0GCSqGSIb3DQEBCwUAMIGrMQswCQYDVQQGEwJJTDEPMA0GA1UECBMGSXNyYWVsMS4wLAYDVQQKEyVDaGVja1BvaW50IFNvZnR3YXJlIFRlY2hub2xvZ2llcyBMVEQuMQwwCgYDVQQLEwNNSVMxIjAgBgNVBAMTGUNoZWNrUG9pbnQtU1NMLUluc3BlY3Rpb24xKTAnBgkqhkiG9w0BCQEWGmlsX3NlY3VyaXR5QGNoZWNrcG9pbnQuY29tMB4XDTIzMDMxMzAwMDAwMFoXDTIzMDYxMTIzNTk1OVowbzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExEzARBgNVBAcTCk1lbmxvIFBhcmsxHTAbBgNVBAoTFE1ldGEgUGxhdGZvcm1zLCBJbmMuMRcwFQYDVQQDDA4qLndoYXRzYXBwLm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPjo05vRHAJYYWx55SOu2b1ZIQPOOtJNipSBXf1BFBDQhrkp20YTA296MzKii2j3TgVi/1t44cW5mD1RWobfAQujgbMwgbAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMHQGA1UdEQRtMGuCDioud2hhdHNhcHAubmV0ghIqLmNkbi53aGF0c2FwcC5uZXSCEiouc25yLndoYXRzYXBwLm5ldIIOKi53aGF0c2FwcC5jb22CBXdhLm1lggx3aGF0c2FwcC5jb22CDHdoYXRzYXBwLm5ldDAOBgNVHQ8BAf8EBAMCBaAwCQYDVR0TBAIwADANBgkqhkiG9w0BAQsFAAOCAgEAA/sIadLr9ahEVq8h9HuofHODUuzxVFulAZu8uSiyY4ACbaHcvm36MYQCzYV56t4fe+I++ls8KAESZgdE0KoD5/6efzK05Ufok+y15QexAR5AxZlJqtoHIuc7iOolPbkLW77GKrbgfEgmwOCX9/86Pug4ZSrrBUPPt9i3accNkAP+SH9Lft1geS2E/q+xcRhbhDcYTYD56X0MiEv0UaAzwS3adWAZbD7R42u+xNCpX8iUyiwp2UvMf0l/+Q8CAtw4D5s/8hD7Vqvrv4H/ZfV7SrZ+rPrihi01t6LlcpZ2YMucX/tSgDzkjYWmT26V2OgRklM0aQWvHD3DVpghIJfI2swAAJJ5wvqwcJeAWHAQb3aQZgHXjGF/LyBYCQsohTHUL7rhL8CxNlDTNhN2e+NRFGYGer157RCmM8xKroe3/X9pYifbzyEWInqQ+ycmLsQyAd7pPW+W1K1tlk9Niqk3dNQ10daYGau3IPWF5+iHtOlWjLcQrSj60Uv7Ebi0E+bOe0tDabunCj6SEauGFxeJhM9xUZnOwb5wqIt+uGqPQ9WRJLehqwdFhiWOqwUfNcksn7l0M6e9Mnkh1J2kGxamQ0bvK7ftpm5O8MTAft0y882IfC++Zuk4gLhQoeE3s6877/rrHRJB/H8ZUaaBxAi2qH0NZ+ParXUxOkil5rVgFqI="
}
data "checkpoint_management_custom_trusted_ca_certificate" "data" {
  uid = "${checkpoint_management_add_custom_trusted_ca_certificate.example.id}"
}
```

## Argument Reference

The following arguments are supported:


* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `added_by` - By whom the certificate was added.
* `base64_certificate` -  Certificate file encoded in base64.<br/>Valid file formats: x509. 
* `issued_by` - Trusted CA certificate issued by.
* `issued_to` - Trusted CA certificate issued to.
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

