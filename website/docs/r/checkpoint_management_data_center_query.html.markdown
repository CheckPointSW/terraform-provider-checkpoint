---
layout: "checkpoint"
page_title: "checkpoint_management_data_center_query"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-center-query"
description: |- This resource allows you to execute Check Point Data Center Query.
---

# checkpoint_management_data_center_query

This resource allows you to execute Check Point Data Center Query.

## Example Usage

```hcl
resource "checkpoint_management_data_center_query" "example" {
  name         = "myQuery"
  data_centers = ["All"]
  query_rules {
    key_type = "predefined"
    key      = "name-in-data-center"
    values   = ["firstVal", "secondVal"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `data_centers` - (Optional) Collection of Data Center servers identified by the name or UID. Use "All" to select all data centers.data_centers blocks are documented below.
* `query_rules` - (Optional) Data Center Query Rules.<br>There is an 'AND' operation between multiple Query Rules.query_rules blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`query_rules` supports the following:

* `key_type` - (Optional) The type of the "key" parameter.<br>Use "predefined" for these keys: type-in-data-center, name-in-data-center, and ip-address.<br>Use "tag" to query the Data Center tag�s property.
* `key` - (Optional) Defines in which Data Center property to query.<br>For key-type "predefined", use these keys:type-in-data-center, name-in-data-center, and ip-address.<br>For key-type "tag", use the Data Center tag key to query.<br>Keys are case-insensitive.
* `values` - (Optional) The value(s) of the Data Center property to match the Query Rule.<br>Values are case-insensitive.<br>There is an 'OR' operation between multiple values.<br>For key-type "predefined" and key 'ip-address', the values must be an IPv4 or IPv6 address.<br>For key-type "tag", the values must be the Data Center tag values.values blocks are documented below.
