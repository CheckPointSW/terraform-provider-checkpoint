---
layout: "checkpoint"
page_title: "checkpoint_management_domain"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-domain"
description: |- Use this data source to get information on an existing Check Point Domain.
---

# Data Source: checkpoint_management_domain

Use this data source to get information on an existing Check Point Domain.

## Example Usage

```hcl
resource "checkpoint_management_domain" "example" {
    name = "domain1"
    servers {
      name = "domain1_ManagementServer_1"
      ipv4_address = "192.0.2.1"
      multi_domain_server = "MDM_Server"
    }
}

data "checkpoint_management_domain" "data_domain" {
  name = "${checkpoint_management_domain.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Must be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `domain_type` - N/A.
* `global_domain_assignments` - N/A.global_domain_assignments blocks are documented below.


`global_domain_assignments` supports the following:

* `assignment_status` - N/A.
* `assignment_up_to_date` - The time when the assignment was assigned.assignment_up_to_date blocks are documented below.
* `global_access_policy` - Global domain access policy that is assigned to a dependent domain.
* `global_domain` - Global domain. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard l...
* `global_threat_prevention_policy` - Global domain threat prevention policy that is assigned to a dependent domain.
* `manage_protection_actions` - N/A.


`assignment_up_to_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
