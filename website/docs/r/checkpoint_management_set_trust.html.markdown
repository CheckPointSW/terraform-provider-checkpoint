---
layout: "checkpoint"
page_title: "checkpoint_management_set_trust"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-trust"
description: |-
 This resource allows you to execute Check Point Set Trust.
---

# checkpoint_management_set_trust

This resource allows you to execute Check Point Set Trust.

## Example Usage


```hcl
resource "checkpoint_management_set_trust" "example" {
  name = "smb_daip"
  one_time_password = "aaaa"
  trust_settings {
    initiation_phase = "when_gateway_connects"
    identification_method = "mac_address"
    gateway_mac_address = "ffff:45:0000:0000:0000"
  }
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Minimum Check Point password length.
* `ipv4_address` - (Optional) IP address of the object, for establishing trust with dynamic gateways.
* `one_time_password` - (Optional) Shared password to establish SIC between the Security Management and the Security Gateway.
* `trust_method` - (Optional) Establish the trust communication method.
* `trust_settings` - (Optional) Settings for the trusted communication establishment. trust settings blocks are documented below.

`trust_settings` supports the following:

* `gateway_mac_address` - (Optional) Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.
* `identification_method` - (Optional) How to identify the gateway (relevant for Spark DAIP gateways only).
* `initiation_phase` - (Optional) Push the certificate to the Security Gateway immediately, or wait for the Security Gateway to pull the certificate. Default value for Spark Gateway is 'when_gateway_connects'.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

