---
layout: "checkpoint"
page_title: "checkpoint_gaia_maestro_site"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-maestro-site"
description: |-
This resource allows you to execute Check Point Maestro Site.
---

# checkpoint_gaia_maestro_site

This resource allows you to execute Check Point Maestro Site.

## Example Usage


```hcl
resource "checkpoint_gaia_maestro_site" "example" {
  site_id = 1
  descriptions {
    security_group = 1
    description = "Site 1 description for Security Group 1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `site_id` - (Required)  
* `descriptions` - (Optional) Provide optional site description per Security Group descriptions blocks are documented below.
* `include_pending_changes` - (Computed)  


`descriptions` supports the following:

* `security_group` - (Optional) The Site Security Group 
* `description` - (Optional) Site description 


## Import

Use the existing resource's `site_id` as the import ID.

```shell
terraform import checkpoint_gaia_maestro_site.example '1'
```

To specify optional read context or filters (`include_pending_changes`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
