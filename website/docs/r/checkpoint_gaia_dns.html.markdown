---
layout: "checkpoint"
page_title: "checkpoint_gaia_dns"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-dns"
description: |-
This resource allows you to execute Check Point Dns.
---

# checkpoint_gaia_dns

This resource allows you to execute Check Point Dns.

## Example Usage


```hcl
resource "checkpoint_gaia_dns" "example" {
  primary              = "1.2.3.4"
  secondary            = "2.3.4.5"
  tertiary             = "3.4.5.6"
  suffix               = "checkpoint.com"
  listening_interfaces = ["all"]
  forwarding_domains {
    suffix    = "google.com"
    primary   = "1.1.1.1"
    secondary = "2.2.2.2"
    tertiary  = "4.4.2.2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `primary` - (Optional) Use empty-string in order to remove the setting 
* `secondary` - (Optional) Use empty-string in order to remove the setting 
* `tertiary` - (Optional) Use empty-string in order to remove the setting 
* `suffix` - (Optional) Use empty-string in order to remove the setting 
* `forwarding_domains` - (Optional) DNS proxy forwarding domains forwarding_domains blocks are documented below.
* `listening_interfaces` - (Optional) DNS proxy listening interfaces 
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`forwarding_domains` supports the following:

* `suffix` - (Optional)  
* `primary` - (Optional)  
* `secondary` - (Optional)  
* `tertiary` - (Optional)  


## Import

Use `settings` to import the existing configuration on the connected Gaia machine.

```shell
terraform import checkpoint_gaia_dns.example 'settings'
```

To specify optional read context or filters (`virtual_system_id`, `member_id`), use a JSON object with Terraform attribute names and include all required lookup fields.

`member_id` selects the member used for reads; it does not restrict later writes to that member. Gaia set operations retain the scope documented by the resource.

Create a matching resource block before importing, and review `terraform plan` afterward. The importer first validates and records the lookup fields; Terraform then calls the resource reader to fetch existing configuration. Import does not create the resource or recover write-only values such as passwords.
