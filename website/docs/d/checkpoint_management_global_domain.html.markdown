---
layout: "checkpoint"
page_title: "checkpoint_management_global_domain"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-global-domain"
description: |-
Use this data source to get information on an existing Check Point Global Domain.
---

# Data Source: checkpoint_management_global_domain

Use this data source to get information on an existing Check Point Global Domain.

## Example Usage


```hcl
data "checkpoint_management_global_domain" "data_global_domain" {
    name = "Global"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optioanl) Object name.
* `type` - Object type.
* `domain_type` - The domain type.
* `global_domain_assignments` - The assignments. global_domain_assignments blocks are documented below.
* `servers` - Domain Servers. servers blocks are documented below.
* `tags` - Collection of tag objects identified by the name or UID.
* `color` - Color of the object.
* `comments` - Comments string.

`global_domain_assignments` supports the following:

* `name` - Object name.
* `uid` - Object unique identifier.
* `type` - Object type.
* `assignment_status` - The status of the assignment.
* `assignment_up_to_date` - The time when the assignment was assigned. assignment_up_to_date blocks are documented below.
* `dependent_domain` - Dependent domain. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. dependent_domain blocks are documented below.
* `global_access_policy` - Global domain access policy that is assigned to a dependent domain.
* `global_threat_prevention_policy` - Global domain threat prevention policy that is assigned to a dependent domain.
* `manage_protection_actions`
* `tags` - Collection of tag objects identified by the name or UID.
* `color` - Object color.
* `comments` - Coemmnet string.


`assignment_up_to_date` supports the following:

* `iso_9601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`dependent_domain` supports the following:

* `name` - Object name.
* `uid` - Object unique identifier.


`servers` supports the following:

* `name` - Object name.
* `active` - Domain server status.
* `ipv4_address` - IPv4 address.
* `ipv6_address` - IPv6 address.
* `multi_domain_server` - Multi Domain server name or UID.
* `skip_start_domain_server` - Set this value to be true to prevent starting the new created domain.
* `type` - Domain server type.