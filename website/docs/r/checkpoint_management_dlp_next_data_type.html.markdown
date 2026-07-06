---
layout: "checkpoint"
page_title: "checkpoint_management_dlp_next_data_type"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-dlp-next-data-type"
description: |-
This resource allows you to execute Check Point Dlp Next Data Type.
---

# checkpoint_management_dlp_next_data_type

This resource allows you to execute Check Point Dlp Next Data Type.

## Example Usage


```hcl
resource "checkpoint_management_dlp_next_data_type" "example" {
  name = "ASIC or FPGA Designs"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `external_id` - (Optional) DLP Next Data Type unique identifier in Infinity Portal.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `description` - (Computed) Description of the data type.
