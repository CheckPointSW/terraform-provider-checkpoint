---
layout: "checkpoint"
page_title: "checkpoint_management_cloud_services"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cloud-services"
description: |-
This resource allows you to execute Check Point Show Cloud Services.
---

# checkpoint_management_cloud_services

This resource allows you to execute Check Point Show Cloud Services.

## Example Usage

```hcl
data "checkpoint_management_cloud_services" "example" {}
```

## Argument Reference

The following arguments are supported:
* `status` - Status of the connection to the Infinity Portal.
* `connected_at` - The time of the connection between the Management Server and the Infinity Portal. connected_at is documented below.
* `management_url` - The Management Server's public URL.

`connected_at` supports the following:
* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

