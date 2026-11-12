---
layout: "checkpoint"
page_title: "checkpoint_gaia_license"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-license"
description: |-
This resource allows you to execute Check Point License.
---

# checkpoint_gaia_license

This resource allows you to execute Check Point License.

## Example Usage


```hcl
resource "checkpoint_gaia_license" "example" {
  license = ""
}
```

## Argument Reference

The following arguments are supported:

* `license` - (Required) The license string received from the User Center - without 'cplic put' 
* `signature` - (Computed) The license signature to show details for 
* `target` - (Optional) The remote target to deploy the license on - used for central licenses only 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `ip_addr` - (Computed) Computed field, returned in the response. 
* `expiration` - (Computed) Computed field, returned in the response. 
* `sku` - (Computed) Computed field, returned in the response. 
* `ck` - (Computed) Computed field, returned in the response. 
* `central` - (Computed) Computed field, returned in the response. 


## Import

Use the existing resource's `signature` as the import ID.

```shell
terraform import checkpoint_gaia_license.example '<license-signature>'
```

To specify optional read context or filters (`member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
