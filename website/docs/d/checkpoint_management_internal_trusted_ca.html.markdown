---
layout: "checkpoint"
page_title: "checkpoint_management_internal_trusted_ca"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-internal-trusted-ca"
description: |-
Use this data source to get information on an existing Check Point Internal Trusted Ca.
Use this data source to get information on an existing Check Point Internal Trusted Ca.
Use this data source to get information on an existing Check Point Internal Trusted Ca.
---

# checkpoint_management_internal_trusted_ca

Use this data source to get information on an existing Check Point Internal Trusted Ca.

## Example Usage


```hcl
resource "checkpoint_management_command_set_internal_trusted_ca" "internal_ca" {

  cache_crl = "false"
  crl_cache_timeout = 1200
  
}
data "checkpoint_management_internal_trusted_ca" "data3" {
  depends_on = [checkpoint_management_command_set_internal_trusted_ca.internal_ca]
}
```

## Argument Reference

The following arguments are supported:

* `name` -  Object name. Must be unique in the domain.
* `uid` -  Object unique identifier.
* `retrieve_crl_from_http_servers` -  Whether to retrieve Certificate Revocation List from http servers. 
* `retrieve_crl_from_ldap_servers` -  Whether to retrieve Certificate Revocation List from ldap servers.
* `cache_crl` -  Cache Certificate Revocation List on the Security Gateway. 
* `crl_cache_method` -  Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period. 
* `crl_cache_timeout` -  When to fetch new Certificate Revocation List (in minutes). 
* `allow_certificates_from_branches` -  Allow only certificates from listed branches. 
* `branches` -  Branches to allow certificates from. Required only if "allow-certificates-from-branches" set to "true".branches blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 



## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

