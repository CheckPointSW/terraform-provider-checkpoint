---
layout: "checkpoint"
page_title: "checkpoint_management_log_exporter"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-log-exporter"
description: |- Use this data source to get information on an existing Log Exporter.
---


# checkpoint_management_log_exporter

Use this data source to get information on an existing Log Exporter.

## Example Usage


```hcl
data "checkpoint_management_log_exporter" "data_log_exporter" {
name = "example_log_exporter"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `target_server` - Target server port to which logs are exported.
* `target_port` - Port number of the target server.
* `protocol` - Protocol used to send logs to the target server.
* `enabled` - Indicates whether to enable export.
* `attachments` - Log exporter attachments. attachments blocks are documented below.
* `data_manipulation` - Log exporter data manipulation. data_manipulation blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.


`attachments` supports the following:

* `add_link_to_log_attachment` - Indicates whether to add link to log attachment in SmartView.
* `add_link_to_log_details` - Indicates whether to add link to log details in SmartView.
* `add_log_attachment_id` - Indicates whether to add log attachment ID.


`data_manipulation` supports the following:

* `aggregate_log_updates` - Indicates whether to aggregate log updates.
* `format` - Logs format. 
