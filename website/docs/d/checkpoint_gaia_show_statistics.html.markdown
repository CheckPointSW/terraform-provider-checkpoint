---
layout: "checkpoint"
page_title: "checkpoint_gaia_show_statistics"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-show-statistics"
description: |-
This resource allows you to execute Check Point Show Statistics.
---

# checkpoint_gaia_show_statistics

This resource allows you to execute Check Point Show Statistics.

## Example Usage


```hcl
data "checkpoint_gaia_show_statistics" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `new_query` - (Optional) Create new query. new_query blocks are documented below.
* `use_cursor` - (Optional) Request for data on existing query. use_cursor blocks are documented below.
* `virtual_system_id` - (Optional) Virtual System ID. Relevant for VSNext setups 
* `member_id` - (Optional) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`new_query` supports the following:

* `filter` - (Optional) Filter the results by a list of labels and stat-IDs filter blocks are documented below.
* `show_info` - (Optional) Data response contains additional information. 
* `ignore_warnings` - (Optional) Ignore all warnings. 
* `historical_data` - (Optional) Data response should contains historical data.. Supported starting from Gaia version R82.10 
* `from_date` - (Optional) The starting date (posix/iso-8601) for the query.. Supported starting from Gaia version R82.10 
* `to_date` - (Optional) The ending date (posix/iso-8601) for the query.. Supported starting from Gaia version R82.10 
* `records_limit` - (Optional) Limit the amount of records in the result.. Supported starting from Gaia version R82.10 
* `override_db_name` - (Optional) Override the historical database name (Notice you can only provide a name - You must upload the DB as <Name>.dat).. Supported starting from Gaia version R82.10 


`use_cursor` supports the following:

* `cursor_id` - (Optional) The cursor ID for an existing query. For cursor based pagination. 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

