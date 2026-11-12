---
---
layout: "checkpoint"
page_title: "checkpoint_management_set_internal_trusted_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-internal-trusted-ca"
description: |-
This resource allows you to execute Check Point Set Internal Trusted Ca.
---

# checkpoint_management_set_internal_trusted_ca

This resource allows you to execute Check Point Set Internal Trusted Ca.

## Example Usage


```hcl
resource "checkpoint_management_set_internal_trusted_ca" "example" {
  retrieve_crl_from_http_servers = false
  cache_crl = false
}
```

## Argument Reference

The following arguments are supported:

* `retrieve_crl_from_http_servers` - (Optional) Whether to retrieve Certificate Revocation List from http servers. 
* `cache_crl` - (Optional) Cache Certificate Revocation List on the Security Gateway. 
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


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

