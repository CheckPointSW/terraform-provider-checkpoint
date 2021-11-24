---
layout: "checkpoint"
page_title: "checkpoint_management_data_center_content"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-data_center_content"
description: |- Use this data source to get information on an existing Check Point data center content.
---

# Data Source: checkpoint_management_data_center_content

Use this data source to get information on an existing Check Point data center content.

## Example Usage

```hcl
data "checkpoint_management_data_center_content" "test" {
  name   = "Network"
  filter = {
    text = "TEXT_TO_FIND"
    uri  = "DATA_CENTER_URI"
  }
  limit  = 100
}
```

## Argument Reference

The following arguments are supported:

* `data_center_name` - (Optional) Name of the Data Center Server where to search for objects.
* `data_center_uid` - (Optional) Unique identifier of the Data Center Server where to search for objects.
* `limit` - The maximal number of returned results.
* `offset` - Number of the results to initially skip.
* `order` - Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order.
  orders blocks are documented below.
* `uid_in_data_center` - Return result matching the unique identifier of the object on the Data Center Server.
* `filter` - Return results matching the specified filter.

`filter` supports the following:

* `text` - (Optional) Return results containing the specified text value.
* `uri` - (Optional) Return results under the specified Data Center Object (identified by URI).
* `parent_uid_in_data_center` - (Optional) Return results under the specified Data Center Object (identified by UID).

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.
