---
layout: "checkpoint"
page_title: "checkpoint_management_app_control_status"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-app-control-status"
description: |-
Use this data source to get information on an existing Check Point App Control status.
---

# checkpoint_management_app_control_status

Use this data source to get information on an existing Check Point App Control status.

## Example Usage


```hcl
data "checkpoint_management_app_control_status" "data" {
}
```

## Argument Reference

The following arguments are supported:

* `uid` - Object Identifier.
* `last_updated` - The last time Application Control & URL Filtering was updated on the management server. last_updated blocks are documented below.
* `installed_version` - Installed Application Control & URL Filtering version.
* `installed_version_creation_time` - Installed Application Control & URL Filtering version creation time. installed_version_creation_time blocks are documented below.

`last_updated` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

`installed_version_creation_time` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

