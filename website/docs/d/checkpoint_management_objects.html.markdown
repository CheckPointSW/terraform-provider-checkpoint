---
layout: "checkpoint"
page_title: "checkpoint_management_objects"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-objects"
description: |-
Use this data source to get information on an existing Check Point Objects.
---

# Data Source: checkpoint_management_objects

Use this data source to get information on an existing Check Point Objects.

## Example Usage


```hcl
data "checkpoint_management_objects" "data_objects" {
  limit = 10
}
```

## Argument Reference

The following arguments are supported:

* `uids` - (Optional) List of UIDs of the objects to retrieve.
* `filter` - (Optional) Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the "ip-only" parameter to true.
* `ip_only` - (Optional) If using "filter", use this field to search objects by their IP address only, without involving the textual search.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional) The maximal number of returned results.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.
* `type` - (Optional) The objects' type, e.g.: host, service-tcp, network, address-range...
* `dereference_group_members` - (Optional) Indicates whether to dereference "members" field by details level for every object in reply.
* `show_membership` - (Optional) Indicates whether to calculate and show "groups" field for every object in reply.
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.