---
layout: "checkpoint"
page_title: "checkpoint_management_default_administrator_settings"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-default-administrator-settings"
description: |-
  Use this data source to get information on an existing Check Point Default Administrator Settings.
---

# Data Source: checkpoint_management_default_administrator_settings

Use this data source to get information on an existing Check Point Default Administrator Settings.

## Example Usage


```hcl
data "checkpoint_management_default_administrator_settings" "data_test" {
}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object unique identifier.
* `authentication_method` - Authentication method for new administrator. 
* `expiration_type` - Expiration type for new administrator. 
* `expiration_date` - Expiration date for new administrator in YYYY-MM-DD format. expiration_date blocks are documented below.
* `expiration_period` - Expiration period for new administrator.
* `expiration_period_time_units` - Expiration period time units for new administrator.
* `indicate_expiration_in_admin_view` - Indicates whether to notify administrator about expiration. 
* `notify_expiration_to_admin` - Indicates whether to show 'about to expire' indication in administrator view. 
* `days_to_indicate_expiration_in_admin_view` - Number of days in advanced to show 'about to expire' indication in administrator view. 
* `days_to_notify_expiration_to_admin` - Number of days in advanced to notify administrator about expiration. 

`expiration_date` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

