---
layout: "checkpoint"
page_title: "checkpoint_management_show_objects"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-show-objects"
description: |-
  This resource allows you to execute Check Point Show Objects.
---

# Data Source: checkpoint_management_show_objects

This resource allows you to execute Check Point Show Objects.

## Example Usage


```hcl
data "checkpoint_management_show_objects" "query" {
    type = "service-tcp"
    filter = "13+"
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the "ip-only" parameter to true.
* `ip_only` - (Optional) If using "filter", use this field to search objects by their IP address only, without involving the textual search.
* `type` - (Optional) The objects' type, e.g.: host, service-tcp, network, address-range.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional) Number of the results to initially skip.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.
* `from` - From which element number the query was done.
* `to` - To which element number the query was done.
* `total` - Total number of elements returned by the query.
* `objects` - Collection of retrieved objects. objects blocks blocks are documented below.

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.

`objects` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.
* `type` - Object type.
* `domain` - Information about the domain that holds the Object. domain blocks are documented below.

`domain` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.
* `domain_type` - Domain type.